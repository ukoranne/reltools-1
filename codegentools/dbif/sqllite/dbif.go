package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
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
	bytes, err := ioutil.ReadFile(jsonFile)
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
		//fmt.Println("### struct Name ", name)
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
