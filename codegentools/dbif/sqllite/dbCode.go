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

func (obj *ObjectSrcInfo) WriteStoreObjectInDBFcn(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") StoreObjectInDb(dbHdl *sql.DB)  (int64, error) {\n")
	lines = append(lines, "var objectId int64\n")
	stmt := "dbCmd := fmt.Sprintf(\" INSERT INTO " + obj.ObjName + " ( "
	attrNamesList := ""
	valuesList := "VALUES ("
	argsList := ""
	for idx, fld := range str.Fields.List {
		if fld.Names != nil {
			attrNamesList = attrNamesList + fld.Names[0].String() + " "
			argsList = argsList + "obj." + fld.Names[0].String() + " "
			valuesList = valuesList + "'%v' "

			if idx != len(str.Fields.List)-1 {
				attrNamesList = attrNamesList + ", "
				argsList = argsList + ", "
				valuesList = valuesList + ", "
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
                                                                                                                           
    }                                                                                                                       
    return objectId, err                                                                                                    
    }` + "\n"
	lines = append(lines, fcnClosure)
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteCreateTableFcn(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") CreateDBTable(dbHdl *sql.DB) error {\n")
	lines = append(lines, "dbCmd := \"CREATE TABLE IF NOT EXISTS "+obj.ObjName+" \"+ \n")
	lines = append(lines, "\"( \" + \n")
	keys := make([]string, 0)
	for _, fld := range str.Fields.List {
		if fld.Names != nil {
			switch fld.Type.(type) {
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
		
	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)                                                                          
   	return err
	}` + "\n"

	lines = append(lines, fcnClosure)

	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteDeleteObjectFromDbFcn(str *ast.StructType, fd *os.File) {
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

func (obj *ObjectSrcInfo) WriteGetObjectFromDbFcn(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetObjectFromDb (objKey string, dbHdl *sql.DB) (ConfigObj, error) {\n")
	lines = append(lines, "var object "+obj.ObjName+"\n")
	lines = append(lines, "sqlKey, err := obj.GetSqlKeyStr(objKey)\n")
	lines = append(lines, "dbCmd := \"select * from "+obj.ObjName+" where \" + sqlKey\n")
	attrNamesList := "err = dbHdl.QueryRow(dbCmd).Scan("
	for _, fld := range str.Fields.List {
		if fld.Names != nil {
			attrNamesList = attrNamesList + "&obj." + fld.Names[0].String() + ", "
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
	return false
	switch typeVal {
	case "uint8", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "complex64", "complex128", "byte", "rune":
		return true
	default:
		return false
	}
	return false
}
func (obj *ObjectSrcInfo) WriteKeyRelatedFcns(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetKey () (string, error) {\n")

	multipleKeys := 0
	keyStr := "key := "
	reverseKeyStr := "sqlKey := \""
	for idx, fld := range str.Fields.List {
		if fld.Names != nil {
			switch fld.Type.(type) {
			case *ast.Ident:
				varName := fld.Names[0].String()
				if fld.Tag != nil {
					if strings.Contains(fld.Tag.Value, "SNAPROUTE") {
						idntType := fld.Type.(*ast.Ident)
						varType := idntType.String()
						if multipleKeys == 0 {
							if obj.IsNumericType(varType) {
								keyStr = keyStr + " string (strconv.Atoi(int(obj." + varName + "))) "
							} else {
								keyStr = keyStr + " string (obj." + varName + ") "
							}
							reverseKeyStr = reverseKeyStr + varName + " = \" + \"\\\"\" + keys [" + strconv.Itoa(idx-1) + "]"

							multipleKeys = 1
						} else {
							if obj.IsNumericType(varType) {
								keyStr = keyStr + "+ \"#\" + string (strconv.Atoi(int(obj." + varName + "))) "
							} else {
								keyStr = keyStr + "+ \"#\" + string (obj." + varName + ") "
							}

							reverseKeyStr = reverseKeyStr + " + " + "\"\\\"\"" + " +  \" and \" + " + "\"" + varName + " = \"  + \"\\\"\"  +  keys [" + strconv.Itoa(idx-1) + "]" + " + " + "\"\\\"\""
						}
					}
				}
			}
		}
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

func (obj *ObjectSrcInfo) WriteGetAllObjFromDbFcn(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") GetAllObjFromDb(dbHdl *sql.DB) (objList []* "+obj.ObjName+", err error) { \n")
	lines = append(lines, "dbCmd :=  \"select * from "+obj.ObjName+"\"\n")
	lines = append(lines, `
						rows, err := dbHdl.Query(dbCmd)
						if err != nil {
						 return objList, err
						 }
						defer rows.Close()
						for rows.Next() {`+"\n")

	lines = append(lines, "object := new("+obj.ObjName+")\n")
	stmt := "if err = rows.Scan("
	for idx, fld := range str.Fields.List {
		if fld.Names != nil {
			if idx != len(str.Fields.List)-1 {
				stmt = stmt + "&obj." + fld.Names[0].String() + ", "
			} else {
				stmt = stmt + "&obj." + fld.Names[0].String() + "); err != nil {\n"
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

func (obj *ObjectSrcInfo) WriteCompareObjectsAndDiffFcn(str *ast.StructType, fd *os.File) {
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

func (obj *ObjectSrcInfo) WriteUpdateObjectInDbFcn(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") UpdateObjectInDb(inObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {\n")
	lines = append(lines, "var fieldSqlStr string\n")
	lines = append(lines, "dbObj := inObj.("+obj.ObjName+")\n")
	lines = append(lines, "dbCmd := \"update "+obj.ObjName+" set\"\n")
	lines = append(lines, `
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
								} else if objVal.Kind() == reflect.Bool {
									fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, dbutils.ConvertBoolToInt(bool(fieldVal.Bool())))
								} else {
									fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
								}
								dbCmd += fieldSqlStr
							}
							idx++
						}
						dbCmd += " where " + objSqlKey
						_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
						return err
					}
				`)
	for _, line := range lines {
		fd.WriteString(line)
	}
	fd.Sync()
}

func (obj *ObjectSrcInfo) WriteMergeDbAndConfigObjFcn(str *ast.StructType, fd *os.File) {
	var lines []string
	lines = append(lines, "\nfunc (obj "+obj.ObjName+") MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {\n")
	lines = append(lines, "var mergedObject  "+obj.ObjName)
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
								} else {
									mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
								}
							}
							idx++
					
						}
						return mergedObject , nil
					}
					
					`)
}

func (obj *ObjectSrcInfo) WriteDBFunctions(str *ast.StructType) {
	fmt.Println("Generating dbIf file ", obj.DbFileName)
	dbFile, err := os.Create(obj.DbFileName)
	if err != nil {
		fmt.Println("Failed to open the file", obj.DbFileName)
		return
	}
	defer dbFile.Close()
	dbFile.WriteString(fileHeader)
	obj.WriteCreateTableFcn(str, dbFile)
	obj.WriteStoreObjectInDBFcn(str, dbFile)
	obj.WriteDeleteObjectFromDbFcn(str, dbFile)
	obj.WriteGetObjectFromDbFcn(str, dbFile)
	obj.WriteKeyRelatedFcns(str, dbFile)
	obj.WriteGetAllObjFromDbFcn(str, dbFile)
	obj.WriteCompareObjectsAndDiffFcn(str, dbFile)
	obj.WriteUpdateObjectInDbFcn(str, dbFile)
	dbFile.Sync()
}
