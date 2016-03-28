package main

import (
	"fmt"
	"go/ast"
	"os"
	"strconv"
	"strings"
)

var fileHeader = `package models
import (
   "database/sql"
   "fmt"
   "reflect"
   "strings"
   "utils/dbutils"
)

`

var fileHeaderForState = `package models
import (
   "fmt"
   "strings"
)

`
var goTypesToSqliteMap = map[string]string{
	"bool":    "bool",
	"uint8":   "INTEGER",
	"uint16":  "INTEGER",
	"uint32":  "INTEGER",
	"uint64":  "INTEGER",
	"string":  "TEXT",
	"float64": "REAL",
	"int8":    "INTEGER",
	"int16":   "INTEGER",
	"int32":   "INTEGER",
	"int64":   "INTEGER",
}

func (obj *ObjectSrcInfo) WriteStoreObjectInDBFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") StoreObjectInDb(dbHdl *sql.DB)  (int64, error) {\n")
	lines = append(lines, "var objectId int64\n")
	stmt := "dbCmd := fmt.Sprintf(\" INSERT INTO " + obj.ObjName + " ( "
	attrNamesList := ""
	valuesList := "VALUES ("
	argsList := ""
	first := true
	for _, fld := range str.Fields.List {
		if fld.Names != nil {
			switch fld.Type.(type) {
			case *ast.ArrayType:
				//typ := fld.Type.(*ast.ArrayType)
				//fmt.Printf("++++++++ WriteCreateTableFcn: FieldsList=%+v, fld=%+v, type=%+v, tag=%+v, elt=%+v\n", str.Fields.List, fld,
				//	typ, fld.Tag, typ.Elt)
				//typ := spec.(*ast.TypeSpec)
				//str, ok := typ.Type.(*ast.StructType)
				//fmt.Printf("WriteCreateTableFcn: typ=%+v, str=%+v\n", typ, str)
				//listMembers = append(listMembers, fld.Names[0].String())
			case *ast.Ident:
				if !first {
					attrNamesList = attrNamesList + ", "
					argsList = argsList + ", "
					valuesList = valuesList + ", "
				}

				first = false
				attrNamesList = attrNamesList + fld.Names[0].String() + " "
				argsList = argsList + "obj." + fld.Names[0].String() + " "
				valuesList = valuesList + "'%v' "
			}
		}
	}
	valuesList = valuesList + ");\",\n"
	attrNamesList = attrNamesList + ")"
	argsList = argsList + " )"
	lines = append(lines, stmt+attrNamesList+valuesList+argsList+"\n")

	fcnClosure :=
		`result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
   	 if err != nil {
      fmt.Println("**** Failed to Create table", err)
   	} else {
    objectId, err = result.LastInsertId()
    if err != nil {
        fmt.Println("### Failed to return last object id", err)
    }

    }` + "\n"

	lines = append(lines, fcnClosure)
	// Write Secondary table lines
	secondaryLines := obj.WriteSecondaryTableInsertIntoDBFcn(str, fd, attrMap, objMap)
	if len(secondaryLines) > 0 {
		lines = append(lines, secondaryLines...)
	}
	lines = append(lines, `return objectId, err
						    }`+"\n")

	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteSecondaryTableInsertIntoDBFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) []string {
	var lines []string

	if strings.HasPrefix(obj.ObjName, "Vxlan") { // Temporary hack. Need to fix it. Hari. TODO
		return lines
	}
	for _, attrInfo := range attrMap {
		if attrInfo.IsArray == true {
			attrs := make([]string, 0)
			if _, ok := goTypesToSqliteMap[attrInfo.VarType]; !ok {
				memberAttrMap := getObjectMemberInfo(objMap, attrInfo.VarType)
				for memAttrName, _ := range memberAttrMap {
					attrs = append(attrs, memAttrName)
				}
			} else {
				attrs = append(attrs, attrInfo.MemberName)
			}

			for _, info := range attrMap {
				if info.IsKey == true {
					lines = append(lines,
						"for i:= 0; i < len (obj."+attrInfo.MemberName+"); i++ {\n")
					lines = append(lines, "dbCmd = fmt.Sprintf(\" INSERT INTO "+obj.ObjName+attrInfo.MemberName+"("+obj.ObjName+info.MemberName)
					for _, attr := range attrs {
						lines = append(lines, ", "+attr)
					}
					lines = append(lines, ") VALUES ('%v'")
					for _, _ = range attrs {
						lines = append(lines, ", '%v'")
					}
					lines = append(lines, ") ;\",\n")
					lines = append(lines, "obj."+info.MemberName)
					if _, ok := goTypesToSqliteMap[attrInfo.VarType]; !ok {
						for _, attr := range attrs {
							lines = append(lines, ", obj."+attrInfo.MemberName+"[i]."+attr)
						}
					} else {
						lines = append(lines, ", obj."+attrInfo.MemberName+"[i]")
					}
					lines = append(lines, ")\n")
				}
			}
			lines = append(lines,
				`	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
				if err != nil {
				fmt.Println("**** Failed to Create table", err)
				} else {
				_, err = result.LastInsertId()
				if err != nil {
					fmt.Println("### Failed to return last object id", err)
				}
				}
				}
				`)
		}
	}
	return lines
}

func (obj *ObjectSrcInfo) WriteCreateTableFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var prelines []string
	var lines []string
	var listMembers []string
	var pragmaline []string
	var finalLines []string
	//var userDefinedObjectMembersInfo map[string]ObjectMembersInfo
	prelines = append(prelines, "\nfunc (obj "+obj.ObjName+") CreateDBTable(dbHdl *sql.DB) error {\n")
	prelines = append(prelines, "var dbCmd string\n")
	prelines = append(prelines, "var err error\n")
	lines = append(lines, "dbCmd = \"CREATE TABLE IF NOT EXISTS "+obj.ObjName+" \"+ \n")
	lines = append(lines, "\"( \" + \n")
	keys := make([]string, 0)
	for _, fld := range str.Fields.List {
		if fld.Names != nil {
			switch fld.Type.(type) {
			case *ast.ArrayType:
				listMembers = append(listMembers, fld.Names[0].String())
			case *ast.Ident:
				varName := fld.Names[0].String()
				if fld.Tag != nil {
					if strings.Contains(fld.Tag.Value, "SNAPROUTE") {
						keys = append(keys, varName)
					}
				}
				idntType := fld.Type.(*ast.Ident)
				varType := idntType.String()

				if sqlType, ok := goTypesToSqliteMap[varType]; ok {
					lines = append(lines, "\""+varName+"  "+sqlType+", \" +\n")
				} else {
					fmt.Println("No matching SQL Type for golang type ", varType)
					panic("Undefined SQL Type")
				}
			}
		}
	}
	keyStr := "\"PRIMARY KEY ( "
	for idx, key := range keys {
		if idx == 0 {
			keyStr = keyStr + key
		} else {
			keyStr = keyStr + ", " + key
		}
	}
	keyStr = keyStr + ")\" +\n"

	lines = append(lines, keyStr)
	fcnClosure :=
		`")"

	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)` + "\n"
	lines = append(lines, fcnClosure)
	secondaryTblLines := obj.WriteSecondaryTableCreateFcn(str, fd, attrMap, objMap)
	if len(secondaryTblLines) > 0 {
		lines = append(lines, secondaryTblLines...)
		pragmaline = append(pragmaline, `
						    dbCmd = "PRAGMA foreign_keys = ON;"
						    _, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
						    if err != nil {
							    fmt.Println("Failed to SET Foreignkey pragma", err)
						    }
						`+"\n")
	}
	lines = append(lines, "return err "+"\n"+"}  \n")
	// Append Pre Db Cmd Lines... like func name, var define
	finalLines = append(finalLines, prelines[:]...)
	// Append PRAGMA foreign key line if the secondary table is present
	if len(pragmaline) > 0 {
		finalLines = append(finalLines, pragmaline[:]...)
	}
	// Append remaining lines from the function..
	finalLines = append(finalLines, lines[:]...)
	for _, line := range finalLines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteSecondaryTableCreateFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) []string {
	var lines []string
	var frnKeyLine string

	for _, attrInfo := range attrMap {
		comma := ""
		frnKeyLine = ""
		frnKeyRef := ""
		conditionsLine := make([]string, 0)
		if attrInfo.IsArray == true {
			for _, info := range attrMap {
				if info.IsKey == true {
					conditionsLine = append(conditionsLine,
						"\""+obj.ObjName+info.MemberName+" "+goTypesToSqliteMap[info.VarType]+" NOT NULL, \\n \" +\n ")
					frnKeyLine = frnKeyLine + obj.ObjName + info.MemberName
					frnKeyLine = frnKeyLine + comma
					frnKeyRef = frnKeyRef + info.MemberName + comma
					comma = ","
				}
			}
			lines = append(lines, "\ndbCmd = \"CREATE TABLE IF NOT EXISTS "+obj.ObjName+attrInfo.MemberName+" \" + \n")
			lines = append(lines, " \" ( \" + \n")
			lines = append(lines, conditionsLine...)
			if _, ok := goTypesToSqliteMap[attrInfo.VarType]; !ok {
				memberAttrMap := getObjectMemberInfo(objMap, attrInfo.VarType)
				for memAttrName, memAttrInfo := range memberAttrMap {
					lines = append(lines, "\""+memAttrName)
					lines = append(lines, " "+goTypesToSqliteMap[memAttrInfo.VarType]+", \\n \" +\n")
				}
			} else {
				lines = append(lines, "\""+attrInfo.MemberName)
				lines = append(lines, " "+goTypesToSqliteMap[attrInfo.VarType]+", \\n \" +\n")
			}
			lines = append(lines, "\"FOREIGN KEY ( "+frnKeyLine+" ) "+"REFERENCES"+" "+obj.ObjName+"("+frnKeyRef+") ON DELETE CASCADE\"+\n")
			lines = append(lines, "\");\"\n")
			lines = append(lines, `_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)`+"\n")
		}
	}
	return lines
}

func (obj *ObjectSrcInfo) WriteDeleteObjectFromDbFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") DeleteObjectFromDb (objKey string, dbHdl *sql.DB) error {\n")
	lines = append(lines,
		`sqlKey, err := obj.GetSqlKeyStr(objKey)
		if err != nil {
		fmt.Println("GetSqlKeyStr with key", objKey, "failed with error", err)
		return err
	}`)
	lines = append(lines, "\ndbCmd := \"delete from "+obj.ObjName+" where \" + sqlKey \n")
	lines = append(lines, "fmt.Println(\"### DB Deleting "+obj.ObjName+" \") \n")
	lines = append(lines, `_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
							return err
							}
							`)

	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteGetObjectFromDbFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetObjectFromDb (objKey string, dbHdl *sql.DB) (ConfigObj, error) {\n")
	lines = append(lines, "var object "+obj.ObjName+"\n")
	lines = append(lines, "sqlKey, err := obj.GetSqlKeyStr(objKey)\n")
	lines = append(lines, "dbCmd := \"select * from "+obj.ObjName+" where \" + sqlKey\n")
	attrNamesList := "err = dbHdl.QueryRow(dbCmd).Scan("
	first := true
	for _, fld := range str.Fields.List {
		if fld.Names != nil {
			switch fld.Type.(type) {
			case *ast.ArrayType:
				fmt.Println(fld.Names[0].String(), "Array type")
			default:
				if !first {
					attrNamesList = attrNamesList + ","
				}
				attrNamesList = attrNamesList + "&object." + fld.Names[0].String()
				first = false
			}
		}
	}
	attrNamesList = attrNamesList + ")\n"
	lines = append(lines, attrNamesList)
	lines = append(lines, "return object, err\n}\n")
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) IsNumericType(typeVal string) bool {
	switch typeVal {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "complex64", "complex128", "byte", "rune":
		return true
	default:
		return false
	}
	return false
}
func (obj *ObjectSrcInfo) WriteKeyRelatedFcns(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetKey () (string, error) {\n")

	lines = append(lines, "keyName := \""+obj.ObjName+"\"\n")
	lines = append(lines, "keyName = strings.TrimSuffix(keyName,"+"\"Config\")\n")
	lines = append(lines, "keyName = strings.TrimSuffix(keyName,"+"\"State\")\n")
	lines = append(lines, "fmt.Println(\"key is \", keyName)\n")

	prevKeyWasNum := false
	numKeys := 0
	keyStr := "key := keyName + \"#\" + "
	reverseKeyStr := "sqlKey := \""
	for _, fld := range str.Fields.List {
		if fld.Names != nil {
			switch fld.Type.(type) {
			case *ast.Ident:
				varName := fld.Names[0].String()
				if fld.Tag != nil {
					if strings.Contains(fld.Tag.Value, "SNAPROUTE") {
						idntType := fld.Type.(*ast.Ident)
						varType := idntType.String()
						if numKeys == 0 {
							if obj.IsNumericType(varType) {
								keyStr = keyStr + " string (fmt.Sprintf(\"%d\", obj." + varName + ")) "
								reverseKeyStr = reverseKeyStr + varName + " = \" + keys [" + strconv.Itoa(numKeys+1) + "]"
								prevKeyWasNum = true
							} else {
								keyStr = keyStr + " string (obj." + varName + ") "
								reverseKeyStr = reverseKeyStr + varName + " = \" + \"\\\"\" + keys [" + strconv.Itoa(numKeys+1) + "]"
								prevKeyWasNum = false
							}
						} else {
							if prevKeyWasNum == false {
								reverseKeyStr = reverseKeyStr + " + \"\\\"\""
							}
							if obj.IsNumericType(varType) {
								keyStr = keyStr + "+ \"#\" + string (fmt.Sprintf(\"%d\", obj." + varName + ")) "
								reverseKeyStr = reverseKeyStr + " + " + " \" and \" + " + "\"" + varName + " = \"  +  keys [" + strconv.Itoa(numKeys+1) + "]"
								prevKeyWasNum = true
							} else {
								keyStr = keyStr + "+ \"#\" + string (obj." + varName + ") "
								reverseKeyStr = reverseKeyStr + " +  \" and \" + " + "\"" + varName + " = \"  + \"\\\"\"  +  keys [" + strconv.Itoa(numKeys+1) + "]"
								prevKeyWasNum = false
							}
						}
						numKeys += 1
					}
				}
			}
		}
	}
	if prevKeyWasNum == false {
		reverseKeyStr = reverseKeyStr + " + \"\\\"\""
	}
	lines = append(lines, keyStr)
	lines = append(lines, `
						return key, nil
						}
						`)

	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetSqlKeyStr (objKey string) (string, error) { \n")
	lines = append(lines, "keys := strings.Split(objKey, \"#\")\n")

	lines = append(lines, reverseKeyStr)
	lines = append(lines, `
						return sqlKey, nil
						}
						`)
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteGetAllObjFromDbFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetAllObjFromDb(dbHdl *sql.DB) (objList []ConfigObj, err error) { \n")
	lines = append(lines, "var object "+obj.ObjName+"\n")
	lines = append(lines, "dbCmd :=  \"select * from "+obj.ObjName+"\"\n")
	lines = append(lines, `
						rows, err := dbHdl.Query(dbCmd)
						if err != nil {
						 return objList, err
						 }
						defer rows.Close()
						objList = make([]ConfigObj, 0)
						for rows.Next() {`+"\n")

	stmt := "if err = rows.Scan("
	for idx, fld := range str.Fields.List {
		if fld.Names != nil {
			if idx != len(str.Fields.List)-1 {
				stmt = stmt + "&object." + fld.Names[0].String() + ", "
			} else {
				stmt = stmt + "&object." + fld.Names[0].String() + "); err != nil {\n"
			}
		}
	}
	lines = append(lines, stmt)
	lines = append(lines, `fmt.Println("Db method Scan failed when interating over OspfAreaEntryConfig")
		}
		objList = append(objList, object)
		}
		return objList, nil
		}`+"\n")
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteCompareObjectsAndDiffFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") CompareObjectsAndDiff(updateKeys map[string]bool, inObj ConfigObj) ([]bool, error) {\n")
	lines = append(lines, "dbObj := inObj.("+obj.ObjName+")")
	lines = append(lines, `
			objTyp := reflect.TypeOf(obj)
			objVal := reflect.ValueOf(obj)
			dbObjVal := reflect.ValueOf(dbObj)
			attrIds := make([]bool, objTyp.NumField())
			idx := 0
			for i := 0; i < objTyp.NumField(); i++ {
				fieldTyp := objTyp.Field(i)
				if fieldTyp.Anonymous {
					continue
				}

				objVal := objVal.Field(i)
				dbObjVal := dbObjVal.Field(i)
				if _, ok := updateKeys[fieldTyp.Name]; ok {
					if objVal.Kind() == reflect.Int {
						if int(objVal.Int()) != int(dbObjVal.Int()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Int8 {
						if int8(objVal.Int()) != int8(dbObjVal.Int()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Int16 {
						if int16(objVal.Int()) != int16(dbObjVal.Int()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Int32 {
						if int32(objVal.Int()) != int32(dbObjVal.Int()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Int64 {
						if int64(objVal.Int()) != int64(dbObjVal.Int()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Uint {
						if uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Uint8 {
						if uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Uint16 {
						if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Uint32 {
						if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Uint64 {
						if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Bool {
						if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
							attrIds[idx] = true
						}
					} else if objVal.Kind() == reflect.Slice {
						attrIds[idx] = true
					} else {
						if objVal.String() != dbObjVal.String() {
							attrIds[idx] = true
						}
					}
					if attrIds[idx] {
						fmt.Println("attribute changed ", fieldTyp.Name)
					}
				}
				idx++

			}
			return attrIds[:idx], nil
		}

		`)
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteUpdateObjectInDbFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string

	lines = append(lines, "\nfunc (obj "+obj.ObjName+") UpdateObjectInDb(inObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {\n")
	lines = append(lines, "var fieldSqlStr string\n")
	lines = append(lines, "keys := make([]string, 0)\n")
	for _, attrInfo := range attrMap {
		if attrInfo.IsKey == true {
			lines = append(lines, "keys = append(keys, \""+attrInfo.MemberName+"\")\n")
		}
	}
	lines = append(lines, "dbObj := inObj.("+obj.ObjName+")\n")
	lines = append(lines, "dbCmd := \"update "+obj.ObjName+" set\"\n")

	lines = append(lines, `
						secondaryTableCommands := make([]string, 0)
						objKey, err := dbObj.GetKey()
						objSqlKey, err := dbObj.GetSqlKeyStr(objKey)
						objTyp := reflect.TypeOf(obj)
						objVal := reflect.ValueOf(obj)
						idx := 0
						for i := 0; i < objTyp.NumField(); i++ {
							if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
								continue
							}

							if attrSet[idx] {
								fieldTyp := objTyp.Field(i)
								fieldVal := objVal.Field(i)
								if fieldVal.Kind() == reflect.Int ||
									fieldVal.Kind() == reflect.Int8 ||
									fieldVal.Kind() == reflect.Int16 ||
									fieldVal.Kind() == reflect.Int32 ||
									fieldVal.Kind() == reflect.Int64 {
									fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
								} else if fieldVal.Kind() == reflect.Uint ||
									fieldVal.Kind() == reflect.Uint8 ||
									fieldVal.Kind() == reflect.Uint16 ||
									fieldVal.Kind() == reflect.Uint32 ||
									fieldVal.Kind() == reflect.Uint64 {
									fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Uint()))
								} else if fieldVal.Kind() == reflect.Bool {
									fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, dbutils.ConvertBoolToInt(bool(fieldVal.Bool())))
								} else if fieldVal.Kind() == reflect.Slice {
                						cmd := "delete from BGPPolicyDefinition"+fieldTyp.Name+" where "+objSqlKey
                						secondaryTableCommands = append(secondaryTableCommands, cmd)
                						for j := 0; j < fieldVal.Len(); j++ {
					                    cmd = "INSERT into BGPPolicyDefinition"+fieldTyp.Name+" ("
					                    attrNameList := make([]string, 0)
					                    valueList := make([]string, 0)
					                    for _, key := range keys {
					                        attrNameList = append(attrNameList, key)
					                        keyVal := objVal.FieldByName(key)
					                        valueList = append(valueList, fmt.Sprintf("'%v'", keyVal.Interface()))
					                    }
					                    secObjVal := fieldVal.Index(i)
					                    secObjTyp := fieldVal.Index(i).Type()
					                    if secObjVal.Kind() == reflect.String {
					                        attrNameList = append(attrNameList, secObjTyp.Name())
					                        valueList = append(valueList, fmt.Sprintf("'%v'", secObjVal.Interface()))
					                    } else if secObjVal.Kind() == reflect.Struct {
					                        for k := 0; k < secObjTyp.NumField(); k++ {
					                            if secFieldTyp := secObjTyp.Field(i); secFieldTyp.Anonymous {
					                                continue
					                            }

					                            secFieldTyp := secObjTyp.Field(i)
					                            secFieldVal := objVal.Field(i)
					                            if secFieldVal.Kind() == reflect.Int ||
					                                secFieldVal.Kind() == reflect.Int8 ||
					                                secFieldVal.Kind() == reflect.Int16 ||
					                                secFieldVal.Kind() == reflect.Int32 ||
					                                secFieldVal.Kind() == reflect.Int64 {
					                                attrNameList = append(attrNameList, secFieldTyp.Name)
					                                valueList = append(valueList, fmt.Sprintf("'%v'", secFieldVal.Interface()))
					                            } else if secFieldVal.Kind() == reflect.Uint ||
					                                secFieldVal.Kind() == reflect.Uint8 ||
					                                secFieldVal.Kind() == reflect.Uint16 ||
					                                secFieldVal.Kind() == reflect.Uint32 ||
					                                secFieldVal.Kind() == reflect.Uint64 {
					                                attrNameList = append(attrNameList, secFieldTyp.Name)
					                                valueList = append(valueList, fmt.Sprintf("'%v'", secFieldVal.Interface()))
					                            } else if secFieldVal.Kind() == reflect.Bool {
					                                attrNameList = append(attrNameList, secFieldTyp.Name)
					                                valueList = append(valueList, fmt.Sprintf("'%v'", dbutils.ConvertBoolToInt(bool(secFieldVal.Bool()))))
					                            } else {
					                                attrNameList = append(attrNameList, secFieldTyp.Name)
					                                valueList = append(valueList, fmt.Sprintf("'%v'", secFieldVal.Interface()))
					                            }
					                        }
					                    }
					                    cmd = cmd + strings.Join(attrNameList, " , ") + ") VALUES (" + strings.Join(valueList, " , ") + ");"
					                    secondaryTableCommands = append(secondaryTableCommands, cmd)
					                }
								} else {
									fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
								}
								dbCmd += fieldSqlStr
							}
							idx++
						}
						dbCmd += " where " + objSqlKey

						for _, cmd := range secondaryTableCommands {
							_, err = dbutils.ExecuteSQLStmt(cmd, dbHdl)
							if err != nil {
								fmt.Println("Failed to execute DB command", cmd, "with error", err)
							}
						}

						_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
						return err
					}
				`)
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteMergeDbAndConfigObjFcn(str *ast.StructType, fd *os.File, attrMap []ObjectMemberAndInfo, objMap map[string]ObjectSrcInfo) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {\n")
	lines = append(lines, "var mergedObject  "+obj.ObjName+"\n")
	lines = append(lines, `objTyp := reflect.TypeOf(obj)
						objVal := reflect.ValueOf(obj)
						dbObjVal := reflect.ValueOf(dbObj)
						mergedObjVal := reflect.ValueOf(&mergedObject)
						idx := 0
						for i := 0; i < objTyp.NumField(); i++ {
							if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
								continue
							}

							objField := objVal.Field(i)
							dbObjField := dbObjVal.Field(i)
							if attrSet[idx] {
								if dbObjField.Kind() == reflect.Int ||
									dbObjField.Kind() == reflect.Int8 ||
									dbObjField.Kind() == reflect.Int16 ||
									dbObjField.Kind() == reflect.Int32 ||
									dbObjField.Kind() == reflect.Int64 {
									mergedObjVal.Elem().Field(i).SetInt(objField.Int())
								} else if dbObjField.Kind() == reflect.Uint ||
									dbObjField.Kind() == reflect.Uint8 ||
									dbObjField.Kind() == reflect.Uint16 ||
									dbObjField.Kind() == reflect.Uint32 ||
									dbObjField.Kind() == reflect.Uint64 {
									mergedObjVal.Elem().Field(i).SetUint(objField.Uint())
								} else if dbObjField.Kind() == reflect.Bool {
									mergedObjVal.Elem().Field(i).SetBool(objField.Bool())
								} else if dbObjField.Kind() == reflect.Slice {
									reflect.Copy(mergedObjVal.Elem().Field(i), objField)
								} else {
									mergedObjVal.Elem().Field(i).SetString(objField.String())
								}
							} else {
								if dbObjField.Kind() == reflect.Int ||
									dbObjField.Kind() == reflect.Int8 ||
									dbObjField.Kind() == reflect.Int16 ||
									dbObjField.Kind() == reflect.Int32 ||
									dbObjField.Kind() == reflect.Int64 {
									mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
								} else if dbObjField.Kind() == reflect.Uint ||
									dbObjField.Kind() == reflect.Uint ||
									dbObjField.Kind() == reflect.Uint8 ||
									dbObjField.Kind() == reflect.Uint16 ||
									dbObjField.Kind() == reflect.Uint32 {
									mergedObjVal.Elem().Field(i).SetUint(dbObjField.Uint())
								} else if dbObjField.Kind() == reflect.Bool {
									mergedObjVal.Elem().Field(i).SetBool(dbObjField.Bool())
								} else if dbObjField.Kind() == reflect.Slice {
									reflect.Copy(mergedObjVal.Elem().Field(i), dbObjField)
								} else {
									mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
								}
							}
							idx++

						}
						return mergedObject , nil
					}

					`)
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) ConvertObjecMemberstMapToOrderedSlice(attrMap map[string]ObjectMembersInfo) (attrMapSlice []ObjectMemberAndInfo) {

	for i := 1; i < len(attrMap)+1; i++ {
		for attr, info := range attrMap {
			if i == info.Position {
				newMember := ObjectMemberAndInfo{
					ObjectMembersInfo: ObjectMembersInfo{
						VarType:     info.VarType,
						IsKey:       info.IsKey,
						IsArray:     info.IsArray,
						Description: info.Description,
						DefaultVal:  info.DefaultVal,
						Position:    info.Position,
					},
					MemberName: attr,
				}
				attrMapSlice = append(attrMapSlice, newMember)
			}
		}
	}
	return
}

func (obj *ObjectSrcInfo) WriteDBFunctions(str *ast.StructType, attrMap map[string]ObjectMembersInfo, objMap map[string]ObjectSrcInfo) {
	dbFile, err := os.Create(obj.DbFileName)
	if err != nil {
		fmt.Println("Failed to open the file", obj.DbFileName)
		return
	}
	defer dbFile.Close()
	attrMapSlice := obj.ConvertObjecMemberstMapToOrderedSlice(attrMap)

	if strings.Contains(obj.Access, "w") || strings.Contains(obj.Access, "rw") {
		dbFile.WriteString(fileHeader)
		obj.WriteCreateTableFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteStoreObjectInDBFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteDeleteObjectFromDbFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteGetObjectFromDbFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteKeyRelatedFcns(str, dbFile, attrMapSlice, objMap)
		obj.WriteGetAllObjFromDbFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteCompareObjectsAndDiffFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteUpdateObjectInDbFcn(str, dbFile, attrMapSlice, objMap)
		obj.WriteMergeDbAndConfigObjFcn(str, dbFile, attrMapSlice, objMap)
	} else {
		dbFile.WriteString(fileHeaderForState)
		obj.WriteKeyRelatedFcns(str, dbFile, attrMapSlice, objMap)
	}
	dbFile.Sync()
}
