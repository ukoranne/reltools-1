package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

// This structure represents the json layout for config objects
type ObjectSrcInfo struct {
	Access     string `json:"access"`
	Owner      string `json:"owner"`
	SrcFile    string `json:"srcfile"`
	ObjName    string
	DbFileName string
	AttrList   []string
}

type ObjectInfoJson struct {
	Access       string `json:"access"`
	Owner        string `json:"owner"`
	SrcFile      string `json:"srcfile"`
	Multiplicity string `json:"multiplicity"`
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
	listingFile := "dbIffiles.txt"
	var objMap map[string]ObjectSrcInfo

	//
	// Files generated from yang models are already listed in right format in genObjectConfig.json
	// However in some cases we have only go objects. Read the goObjInfo.json file and generate a similar
	// structure here.
	//

	goObjSources := base + "/snaproute/src/models/goObjInfo.json"
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
		generateObjectsInformation(fileBase, goSrcFile, ownerName.Owner)
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

	listingsFd, err := os.Create(listingFile)
	if err != nil {
		fmt.Println("Failed to open the file", listingFile)
		return
	}
	defer listingsFd.Close()
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
							//fmt.Printf("%s \n", typ.Name.Name)
							obj.DbFileName = fileBase + typ.Name.Name + "dbif.go"
							listingsFd.WriteString(obj.DbFileName + "\n")
							obj.WriteDBFunctions(str)
						}
					}
				}
			}
		}
	}
}

func generateObjectsInformation(fileBase string, srcFile string, owner string) error {
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

	f, err := parser.ParseFile(fset,
		fileBase+srcFile,
		nil,
		parser.ParseComments)

	if err != nil {
		fmt.Println("Failed to parse input file ", srcFile, err)
		return err
	}

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
												}
											}
										}
									}
								}
							}
						}
						objMap[typ.Name.Name] = obj
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
