package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// This structure represents the json layout for config objects
type ObjectSrcInfo struct {
	Access     string `json:"access"`
	Owner      string `json:"owner"`
	SrcFile    string `json:"srcfile"`
	UsesRedis  bool   `json:"usesRedis"`
	ObjName    string
	DbFileName string
	AttrList   []string
}

type ObjectInfoJson struct {
	Access       string `json:"access"`
	Owner        string `json:"owner"`
	SrcFile      string `json:"srcfile"`
	Multiplicity string `json:"multiplicity"`
	Accelerated  bool   `json:"accelerated"`
	UsesRedis    bool   `json:"usesRedis"`
}

// This structure represents the a golang Structure for a config object
type ObjectMembersInfo struct {
	VarType     string `json:"type"`
	IsKey       bool   `json:"isKey"`
	IsArray     bool   `json:"isArray"`
	Description string `json:"description"`
	DefaultVal  string `json:"default"`
	Position    int    `json:"position"`
	Selections  string `json:"selections"`
	QueryParam  string `json:"queryparam"`
	Accelerated bool   `json:"accelerated"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Len         int    `json:"len"`
	UsesRedis   bool   `json:"usesRedis"`
}

type ObjectMemberAndInfo struct {
	ObjectMembersInfo
	MemberName string
}

// This structure represents the objects that are generated directly from go files instead of yang models
type RawObjSrcInfo struct {
	Owner string `json:"owner"`
}

func main() {
	fset := token.NewFileSet() // positions are relative to fset
	base := os.Getenv("SR_CODE_BASE")
	if len(base) <= 0 {
		fmt.Println(" Environment Variable SR_CODE_BASE has not been set")
		return
	}
	jsonFile := base + "/snaproute/src/models/genObjectConfig.json"
	fileBase := base + "/snaproute/src/models/"
	var objMap map[string]ObjectSrcInfo

	//
	// Create a directory to store all the temporary files
	//
	dirStore := base + "/reltools/codegentools/._genInfo/"
	//os.Mkdir(dirStore, 0777)
	listingFile := dirStore + "generatedGoFiles.txt"

	//
	// Files generated from yang models are already listed in right format in genObjectConfig.json
	// However in some cases we have only go objects. Read the goObjInfo.json file and generate a similar
	// structure here.
	//

	goObjSources := base + "/snaproute/src/models/goObjInfo.json"

	listingsFd, err := os.OpenFile(listingFile, os.O_RDWR|os.O_APPEND+os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Failed to open the file", listingFile)
		return
	}
	defer listingsFd.Close()
	var goSrcsMap map[string]RawObjSrcInfo
	bytes, err := ioutil.ReadFile(goObjSources)
	if err != nil {
		fmt.Println("Error in reading Object configuration file", goObjSources)
		return
	}
	err = json.Unmarshal(bytes, &goSrcsMap)
	if err != nil {
		fmt.Printf("Error in unmarshaling data from ", goObjSources, err)
	}
	for goSrcFile, ownerName := range goSrcsMap {
		generateHandCodedObjectsInformation(listingsFd, fileBase, goSrcFile, ownerName.Owner)
	}

	bytes, err = ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error in reading Object configuration file", jsonFile)
		return
	}
	err = json.Unmarshal(bytes, &objMap)
	if err != nil {
		fmt.Printf("Error in unmarshaling data from ", jsonFile, err)
	}

	for name, obj := range objMap {
		obj.ObjName = name
		srcFile := fileBase + obj.SrcFile
		f, err := parser.ParseFile(fset,
			srcFile,
			nil,
			parser.ParseComments)

		if err != nil {
			fmt.Println("Failed to parse input file ", srcFile, err)
			return
		}

		for _, dec := range f.Decls {
			tk, ok := dec.(*ast.GenDecl)
			if ok {
				for _, spec := range tk.Specs {
					switch spec.(type) {
					case *ast.TypeSpec:
						typ := spec.(*ast.TypeSpec)
						str, ok := typ.Type.(*ast.StructType)
						if ok && name == typ.Name.Name {
							membersInfo := generateMembersInfoForAllObjects(str, dirStore+typ.Name.Name+"Members.json")
							for _, val := range membersInfo {
								if val.UsesRedis == true {
									obj.UsesRedis = true
								}
							}
							obj.DbFileName = fileBase + "gen_" + typ.Name.Name + "dbif.go"
							if strings.ContainsAny(obj.Access, "rw") {
								//fmt.Println("Creating DBIF for", obj.ObjName)
								listingsFd.WriteString(obj.DbFileName + "\n")
								obj.WriteDBFunctions(str, membersInfo, objMap)
							}
						}
					}
				}
			}
		}
	}
}

func getObjectMemberInfo(objMap map[string]ObjectSrcInfo, objName string) (membersInfo map[string]ObjectMembersInfo) {
	fset := token.NewFileSet() // positions are relative to fset
	base := os.Getenv("SR_CODE_BASE")
	if len(base) <= 0 {
		fmt.Println(" Environment Variable SR_CODE_BASE has not been set")
		return membersInfo
	}
	fileBase := base + "/snaproute/src/models/"
	for name, obj := range objMap {
		if objName == name {
			obj.ObjName = name
			srcFile := fileBase + obj.SrcFile
			f, err := parser.ParseFile(fset,
				srcFile,
				nil,
				parser.ParseComments)

			if err != nil {
				fmt.Println("Failed to parse input file ", srcFile, err)
				return
			}

			for _, dec := range f.Decls {
				tk, ok := dec.(*ast.GenDecl)
				if ok {
					for _, spec := range tk.Specs {
						switch spec.(type) {
						case *ast.TypeSpec:
							typ := spec.(*ast.TypeSpec)
							str, ok := typ.Type.(*ast.StructType)
							if ok && name == typ.Name.Name {
								membersInfo = generateMembersInfoForAllObjects(str, "")
								return membersInfo
							}
						}
					}
				}
			}
		}
	}
	return membersInfo
}

func getSpecialTagsForAttribute(attrTags string, attrInfo *ObjectMembersInfo) {
	reg, err := regexp.Compile("[`\"]")
	if err != nil {
		fmt.Println("Error in regex ", err)
	}
	tags := reg.ReplaceAllString(attrTags, "")
	splits := strings.Split(tags, ",")
	for _, part := range splits {
		keys := strings.Split(part, ":")
		for idx, key := range keys {
			alphas, err := regexp.Compile("[^A-Za-z]")
			if err != nil {
				fmt.Println("Error in regex ", err)
			}
			key = alphas.ReplaceAllString(key, "")
			switch key {
			case "SNAPROUTE":
				attrInfo.IsKey = true
			case "DESCRIPTION":
				attrInfo.Description = keys[idx+1]
			case "SELECTION":
				attrInfo.Selections = keys[idx+1]
			case "DEFAULT":
				attrInfo.DefaultVal = keys[idx+1]
			case "ACCELERATED":
				attrInfo.Accelerated = true
			case "MIN":
				attrInfo.Min = 0 //strconv.Atoi(keys[idx+1])
			case "MAX":
				attrInfo.Min = 10 //strconv.Atoi(keys[idx+1])
			case "RANGE":
				attrInfo.Min = 0  //keys[idx+1]
				attrInfo.Max = 10 //keys[idx+1]
			case "LEN":
				attrInfo.Len, _ = strconv.Atoi(strings.TrimSpace(keys[idx+1]))
			case "QPARAM":
				attrInfo.QueryParam = keys[idx+1]
			case "USESREDIS":
				attrInfo.UsesRedis = true
			}
		}
	}
	return
}
func generateMembersInfoForAllObjects(str *ast.StructType, jsonFileName string) map[string]ObjectMembersInfo {
	// Write Skeleton of the structure in json.
	//This would help later python scripts to understand the structure
	var objMembers map[string]ObjectMembersInfo
	objMembers = make(map[string]ObjectMembersInfo, 1)
	var fdHdl *os.File
	var err error
	if jsonFileName != "" {
		fdHdl, err = os.Create(jsonFileName)
		if err != nil {
			fmt.Println("Failed to open the file", jsonFileName)
			return nil
		}
		defer fdHdl.Close()
	}

	for idx, fld := range str.Fields.List {
		if fld.Names != nil {
			varName := fld.Names[0].String()
			switch fld.Type.(type) {
			case *ast.ArrayType:
				arrayInfo := fld.Type.(*ast.ArrayType)
				info := ObjectMembersInfo{}
				info.IsArray = true
				info.Position = idx
				objMembers[varName] = info
				idntType := arrayInfo.Elt.(*ast.Ident)
				varType := idntType.String()
				info.VarType = varType
				objMembers[varName] = info
				if fld.Tag != nil {
					getSpecialTagsForAttribute(fld.Tag.Value, &info)
				}
				objMembers[varName] = info
			case *ast.Ident:
				info := ObjectMembersInfo{}
				if fld.Tag != nil {
					getSpecialTagsForAttribute(fld.Tag.Value, &info)
				}
				idntType := fld.Type.(*ast.Ident)
				varType := idntType.String()
				info.VarType = varType
				info.Position = idx
				objMembers[varName] = info
			}
		}
	}
	lines, err := json.MarshalIndent(objMembers, "", " ")
	if err != nil {
		fmt.Println("Error in converting to Json", err)
	} else {
		if fdHdl != nil {
			fdHdl.WriteString(string(lines))
		}
	}
	return objMembers
}

func generateHandCodedObjectsInformation(listingsFd *os.File, fileBase string, srcFile string, owner string) error {
	var objMap map[string]ObjectInfoJson
	objMap = make(map[string]ObjectInfoJson, 1)

	// First read the existing objects
	genObjInfoFile := fileBase + "genObjectConfig.json"

	bytes, err := ioutil.ReadFile(genObjInfoFile)
	if err != nil {
		fmt.Println("Error in reading Object configuration file", genObjInfoFile)
		return err
	}
	err = json.Unmarshal(bytes, &objMap)
	if err != nil {
		fmt.Printf("Error in unmarshaling data from ", genObjInfoFile, err)
	}

	fset := token.NewFileSet() // positions are relative to fset

	// Now read the contents of Hand coded Go structures
	f, err := parser.ParseFile(fset,
		fileBase+srcFile,
		nil,
		parser.ParseComments)

	if err != nil {
		fmt.Println("Failed to parse input file ", srcFile, err)
		return err
	}

	var marshalFcnsLine []string

	for _, dec := range f.Decls {
		tk, ok := dec.(*ast.GenDecl)
		if ok {
			for _, spec := range tk.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					obj := ObjectInfoJson{}
					obj.SrcFile = srcFile
					obj.Owner = owner
					typ := spec.(*ast.TypeSpec)
					str, ok := typ.Type.(*ast.StructType)
					if ok == true {
						marshalFcnFile := fileBase + "gen_" + strings.Split(srcFile, ".")[0] + "_serializer.go"
						listingsFd.WriteString(marshalFcnFile + "\n")
						for _, fld := range str.Fields.List {
							if fld.Names != nil {
								switch fld.Type.(type) {
								case *ast.Ident:
									if fld.Tag != nil {
										if strings.Contains(fld.Tag.Value, "SNAPROUTE") {
											for _, elem := range strings.Split(fld.Tag.Value, ",") {
												splits := strings.Split(elem, ":")
												switch strings.Trim(splits[0], " ") {
												case "ACCESS":
													obj.Access = strings.Trim(splits[1], "\"")

												case "MULTIPLICITY":
													tmpString := strings.Trim(splits[1], "`")
													obj.Multiplicity = strings.Trim(tmpString, "\"")

												case "ACCELERATED":
													obj.Accelerated = true

												case "USESREDIS":
													obj.UsesRedis = true
												}
											}
										}
									}
								}
							}
						}
						objMap[typ.Name.Name] = obj
						if strings.Contains(obj.Access, "w") || strings.Contains(obj.Access, "r") || strings.Contains(obj.Access, "x") {
							marshalFcnsLine = append(marshalFcnsLine, "\nfunc (obj "+typ.Name.Name+") UnmarshalObject(body []byte) (ConfigObj, error) {\n")
							marshalFcnsLine = append(marshalFcnsLine, `
													var err error
													if len(body) > 0 {
													    if err = json.Unmarshal(body, &obj); err != nil {
													         fmt.Println("###  called, unmarshal failed", obj, err)
													      }
													   }
													   return obj, err
													}
													`)

							marshalFcnFd, err := os.Create(marshalFcnFile)
							if err != nil {
								fmt.Println("Failed to open the file", marshalFcnFile)
								return err
							}
							defer marshalFcnFd.Close()
							marshalFcnFd.WriteString(`package models

													import (
													   "encoding/json"

													   "fmt"
													)`)

							for _, marshalLine := range marshalFcnsLine {
								marshalFcnFd.WriteString(string(marshalLine))
							}
							//marshalFcnFd.WriteString("}\n")
						}

					}
				}
			}
			lines, err := json.MarshalIndent(objMap, "", " ")
			if err != nil {
				fmt.Println("Error is ", err)
			} else {
				genFile, err := os.Create(genObjInfoFile)
				if err != nil {
					fmt.Println("Failed to open the file", genObjInfoFile)
					return err
				}
				defer genFile.Close()
				genFile.WriteString(string(lines))
			}
		}
	}
	return nil
}
