import os
import subprocess
import json
import pprint
import argparse
import itertools


IGNORE_GO_FILE_LIST = ["objectmap.go"]

#HOME = os.getenv("HOME")
srBase = os.environ.get('SR_CODE_BASE', None)
# file holds all objects we are interested in
OBJECT_FILE_MAP_FILE = srBase + "/snaproute/src/models/objectconfig.json"
GO_MODEL_BASE_PATH = srBase + "/snaproute/src/models/"
CODE_GENERATION_PATH = srBase + "/generated/src/models/db/"

goToSqlliteTypeMap = {
  'bool':             {"native_type": "bool"},
  'uint':             {"native_type": "INTEGER", "unsigned": True},
  'uint8':            {"native_type": "INTEGER", "unsigned": True},
  'uint16':           {"native_type": "INTEGER", "unsigned": True},
  'uint32':           {"native_type": "INTEGER", "unsigned": True},
  'uint64':           {"native_type": "INTEGER", "unsigned": True},
  'string':           {"native_type": "TEXT","unsigned": None },
  'float64':          {"native_type": "REAL", "unsigned": False},
  'int':              {"native_type": "INTEGER", "unsigned": False},
  'int8':             {"native_type": "INTEGER", "unsigned": False},
  'int16':            {"native_type": "INTEGER", "unsigned": False},
  'int32':            {"native_type": "INTEGER", "unsigned": False},
  'int64':            {"native_type": "INTEGER", "unsigned": False},
}

gDryRun =  False
def executeGoFmtCommand (fd, command, dstPath) :
    out = ''
    if type(command) != list:
        command = [ command]
    for cmd in command:
        if gDryRun :
            print cmd
        else:
            process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
            out,err = process.communicate()
            # create a go format version, at this point the fd is still
            # open so this is a .tmp file, lets strip this for the new
            # file
            #print out, err
            directory = CODE_GENERATION_PATH
            fmt_name_with_dir = fd.name
            print fmt_name_with_dir
            if not os.path.exists(directory):
              os.makedirs(directory)

        return out

def executeCopyCommand (name, dstPath) :
    directory = dstPath
    if not os.path.exists(directory):
      os.makedirs(directory)

    copyCmd = "cp %s %s" %(name, dstPath,)
    process = subprocess.Popen(copyCmd.split(), stdout=subprocess.PIPE)
    out,err = process.communicate()

    print out, err

    return out

def executeLocalCleanup():

    for name in os.listdir(CODE_GENERATION_PATH):
        if name.endswith(".go"):
            cmd = "rm %s" %(CODE_GENERATION_PATH+name,)
            process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
            out,err = process.communicate()
            print out, err

def scan_dir_for_go_files(directory):
    for name in os.listdir(directory):
        #print "x", directory, name
        path = os.path.join(directory, name)
        if name.endswith('.go'):
            if os.path.isfile(path) and "_enum" not in path and "_func" not in path and "_db" not in path:
                yield (directory, name)
        elif not "." in name:
            for d, f  in scan_dir_for_go_files(path):
                yield (d, f)

def get_dir_file_names(file):
    if file and os.path.isdir(file) and file[-1] != '/':
        file = file + '/'

    if not file or os.path.isdir(file):
        files = scan_dir_for_go_files(GO_MODEL_BASE_PATH)
    elif os.path.isfile(file):
        files = [(os.path.dirname(file), os.path.basename(file))]
    else:
        print "file name [%s] is not a file or a directory" % (file)
        files = None

    if file:
        generatePath = os.path.dirname(file) + '/dbifs/'
    else:
        generatePath = CODE_GENERATION_PATH

    return files, generatePath


def get_all_object_names():

    with open( OBJECT_FILE_MAP_FILE, 'r') as object_data:
        d = json.load(object_data)
        return [ k for k in  d.keys() if 'State' not in k ]

def build_gosqllite_from_go(files, generatePath, objects):
    # generate thrift code from go code
    goStructToListersDict = {}

    dirFileList = []
    for d, g in files:
        if '_func' in g and '_enum' in g and '_db' in g:
            continue

        if g in IGNORE_GO_FILE_LIST:
            continue
        dirFileList.append((d,g))

    # lets determine from the json file the structs and associated listeners
    #for directory, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH):
    for obj in itertools.combinations(objects, 1):
        #print obj
        found = False
        dbFileName = generatePath + obj[0].lower()  + "dbif.go"
        if not os.path.exists(generatePath):
            os.makedirs(generatePath)

        dbFd = open(dbFileName, 'w')
        dbFd.write("package models\n")
        dbFd.write('\nimport (\n\t"database/sql"\n\t"fmt"\n\t"strings"\n\t"utils/dbutils"\n\t"reflect"\n)\n')

        goFuncFileName = generatePath + obj[0].lower()  + "Func.go"
        if not os.path.exists(generatePath):
            os.makedirs(generatePath)

        goFd = open(goFuncFileName, 'w')
        goFd.write("package models\n")
        goFd.write("""import (\n\t "encoding/json"\n\t\"fmt\"\n)\n""")

        #print "generate file name =", dbFileName, "path =", generatePath

        for directory, gofilename in dirFileList:

            found = generate_go_sqllite_funcs(dbFd, directory, gofilename, obj, goFd)
            if found:
                print 'Found structure ', obj, ' in file ', directory, gofilename
                dbFd.close()
                #executeGoFmtCommand(dbFd, ["gofmt -w %s" % dbFd.name], GO_MODEL_BASE_PATH)
                break

        if not found:
            dbFd.close()


def createDBTable(fd, structName, goMemberTypeDict):

    createfuncline = '\nfunc (obj %s) CreateDBTable(dbHdl *sql.DB) error {\n' % structName
    fd.write(createfuncline)
    fd.write('\tdbCmd := "CREATE TABLE IF NOT EXISTS %s " +\n' % structName)
    fd.write('\t\t"( " +')
    keyList = []
    # loop through member and type
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        #print "createDBTable key:", m, "value:", t, "key:", key
        if key > 0:
            keyList.append((m, key))
        if "LIST" in t:
            fd.write('\n\t\t"%s TEXT, " +' %(m,))
        elif 'bool' in t:
            fd.write('\n\t\t"%s INTEGER, " +' %(m,))
        else:
            fd.write('\n\t\t"%s %s, " +' %(m, t))

    keyList = sorted(keyList, key=lambda i: i[1])
    if keyList:
        fd.write('\n\t\t"PRIMARY KEY(')
        for i, (k, l) in enumerate(keyList):
            if i == len(keyList) - 1:
                fd.write('%s) " +' % k)
            else:
                fd.write('%s, ' % k)
    fd.write('\n\t")"\n')

    fd.write('\n\t_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)\n')
    fd.write('\treturn err\n')
    fd.write('}\n')

def createStoreObjInDb(fd, structName, goMemberTypeDict):
    storefuncline = "\nfunc (obj %s) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {\n" % structName
    fd.write(storefuncline)
    fd.write('\tvar objectId int64\n')
    fd.write('\tdbCmd := fmt.Sprintf("INSERT INTO %s (' % structName)
    # loop through member and type
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""%s) VALUES (""" % m)
        else:
            fd.write("""%s, """ % m )

    for i in range(len(goMemberTypeDict[structName])):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""'%v') ;\",\n\t\t""")
        else:
            fd.write("""'%v', """ )

    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t != "bool":
            if i == len(goMemberTypeDict[structName]) - 1:
                fd.write("""obj.%s)\n""" % m )
            else:
                fd.write("""obj.%s, """ % m )
        else:
            if i == len(goMemberTypeDict[structName]) - 1:
                fd.write("""dbutils.ConvertBoolToInt(obj.%s))\n""" % m )
            else:
                fd.write("""dbutils.ConvertBoolToInt(obj.%s), """ % m )

    fd.write("""\tfmt.Println("**** Create Object called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	} else {
	    objectId, err = result.LastInsertId()
	    if err != nil {
		    fmt.Println("### Failed to return last object id", err)
	    }\n
	}
	return objectId, err
}\n""")

def createDeleteObjFromDb(fd, structName, goMemberTypeDict):
    storefuncline = "\nfunc (obj %s) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {\n" % structName
    fd.write(storefuncline)
    fd.write('\tsqlKey, err := obj.GetSqlKeyStr(objKey)\n')
    fd.write('\tif err != nil {\n')
    fd.write('\t\tfmt.Println("GetSqlKeyStr for %s with key", objKey, "failed with error", err)\n' % (structName))
    fd.write('\t\treturn err\n')
    fd.write('\t}\n\n')
    fd.write('\tdbCmd := "delete from %s where " + sqlKey\n' %(structName))
    fd.write('\tfmt.Println("### DB Deleting %s\\n")\n' % structName)
    fd.write('\t_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)\n\treturn err\n}\n')

def createGetObjFromDb(fd, structName, goMemberTypeDict):
    storefuncline = "\nfunc (obj %s) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {\n" % (structName)
    fd.write(storefuncline)
    fd.write('\tvar object %s\n' % (structName))
    fd.write('\tsqlKey, err := obj.GetSqlKeyStr(objKey)\n')
    fd.write('\tdbCmd := "select * from %s where " + sqlKey\n' % (structName))
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t == "bool" or 'LIST' in t:
            fd.write('\tvar tmp%s string\n' %(i))
    fd.write('\terr = dbHdl.QueryRow(dbCmd).Scan(')
    strList = ''
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t == "bool" or "LIST" in t:
            strList += '&tmp%s, ' %(i)
        else:
            strList += '&object.%s, ' %(m)

    strList = strList.rstrip(',')
    fd.write('%s)\n' %(strList))
    fd.write('\tfmt.Println("### DB Get %s\\n", err)\n' % structName)
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t == "bool":
            fd.write('\tobject.%s = dbutils.ConvertStrBoolIntToBool(tmp%s)\n' %(m, i))
        elif 'LIST' in t:
            fd.write("""convtmp%s := strings.Split(tmp%s, ",")
                        for _, x := range convtmp%s {
                            y := strings.Replace(x, " ", "", 1)
                     """ %(m, i, m))
            type = t[5:]
            if type not in ("TEXT", "bool"):
                fd.write(""" z, _ := strconv.Atoi(y)
                             object.%s = append(object.%s, %s(z))
                         """ %(m, m, gt))
            else:
                fd.write("""object.%s = append(object.%s, %s(y))
                     """ %(m, m, gt))
            fd.write("""}\n""""")

    fd.write('\treturn object, err\n}\n')
    
def createGetKey(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) GetKey() (string, error) {" % structName)
    keys = sorted([(m, key) for m, t, gt, key in goMemberTypeDict[structName] if key], key=lambda i: i[1])
    objKey = ' + "#" + '.join(['string(obj.%s)' % (m) for m, key in keys])
    if objKey:
        fd.write('\n\tkey := ')
        fd.write(objKey)

    fd.write("\n\treturn key, nil\n}\n")

def createGetSqlKeyStr(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) GetSqlKeyStr(objKey string) (string, error) {\n" % structName)
    #print "struct dict =", goMemberTypeDict[structName]
    keys = sorted([(m, key) for m, t, gt, key in goMemberTypeDict[structName] if key], key=lambda i: i[1])
    if keys:
        fd.write('\tkeys := strings.Split(objKey, "#")')
        firstKey = ['" = + \\\" + "'.join(['"%s"' % (m), 'keys[%d]' % (i)]) for i, (m, key) in enumerate(keys)]
        #print "firstKey =", firstKey
        sqlKey = ' + " and " + '.join(['+ "\\\"" + '.join(['"%s = "' % (m), 'keys[%d] + "\\\""' % (i)]) for i, (m, key) in enumerate(keys)])
        fd.write('\n\tsqlKey := ')
        fd.write(sqlKey)
        fd.write("\n\treturn sqlKey, nil\n}\n")
    else:
        fd.write("""\n\treturn "", nil\n}\n""")

def createGetAllObjFromDb(fd, structName, goMemberTypeDict):
    fd.write("""\nfunc (obj *%s) GetAllObjFromDb(dbHdl *sql.DB) (objList []*%s, e error) {
	dbCmd := "select * from %s"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for '%s' with error %s", dbCmd, err))
		return objList, err
	}

	defer rows.Close()
    \n""" %(structName, structName, structName, structName, structName))
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t == "bool" or 'LIST' in t:
            fd.write('\tvar tmp%s string\n' %(i))

    fd.write("""\tfor rows.Next() {\n
             object := new(%s)
             if err = rows.Scan(""" %(structName))
    strList = ''
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t == "bool":
            strList += '&tmp%s, ' %(i)
        else:
            strList += '&object.%s, ' %(m)

    strList = strList.rstrip(',')
    fd.write("""%s); err != nil {\n
             fmt.Println("Db method Scan failed when interating over %s")
             }\n""" %(strList, structName))
    for i, (m, t, gt, key) in enumerate(goMemberTypeDict[structName]):
        if t == "bool":
            fd.write('\tobject.%s = dbutils.ConvertStrBoolIntToBool(tmp%s)\n' %(m, i))
        elif 'LIST' in t:
            fd.write("""convtmp%s := strings.Split(tmp%s, ",")
                        for _, x := range convtmp%s {
                            y := strings.Replace(x, " ", "", 1)
                     """ %(m, i, m))
            type = t[5:]
            if type not in ("TEXT", "bool"):
                fd.write(""" z, _ := strconv.Atoi(y)
                             object.%s = append(object.%s, %s(z))
                         """ %(m, m, gt))
            else:
                fd.write("""object.%s = append(object.%s, %s(y))
                     """ %(m, m, gt))
            fd.write("""}\n""""")
    fd.write("""\tobjList = append(objList, object)
    }
    return objList, nil
    }""")


def createUpdateObjInDb(fd, structName, goMemberTypeDict):
    fd.write("""
    func (obj %s) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(%s)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]bool, objTyp.NumField())
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
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
}\n""" %( structName, structName))

    fd.write("""
    func (obj %s) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var merged%s %s
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&merged%s)
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if  attrSet[idx] {
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
	return merged%s, nil
}\n""" %(structName, structName, structName, structName, structName))

    fd.write("""
    func (obj %s) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	db%s := dbObj.(%s)
	objKey, err := db%s.GetKey()
	objSqlKey, err := db%s.GetSqlKeyStr(objKey)
	dbCmd := "update " + "%s" + " set"\n""" %(structName, structName, structName, structName, structName, structName))

    fd.write("""objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
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
}\n""")
'''
def createCompareObjectsAndDiff(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) CompareObjectsAndDiff(dbObj ConfigObj) ([]byte, error) {\n" % (structName))
    fd.write("\tdbStruct := dbObj.(%s)\n" % (structName))
    fd.write("\tobjTyp := reflect.TypeOf(obj)\n")
    fd.write("\tobjVal := reflect.ValueOf(obj)\n")
    fd.write("\tdbObjVal := reflect.ValueOf(dbStruct)\n")
    fd.write("\tattrIds := make([]byte, objTyp.NumField())\n")
    fd.write("\tfor i:=0; i<objTyp.NumField(); i++ {\n")
    fd.write("\t\tobjVal := objVal.Field(i)\n")
    fd.write("\t\tdbObjVal := dbObjVal.Field(i)\n")
    fd.write("\t\tif objVal.Kind() == reflect.Int {\n")
    fd.write("\t\t\tif ((int(objVal.Int()) != 0) && (int(objVal.Int()) != int(dbObjVal.Int()))) {\n")
    fd.write("\t\t\t\tattrIds[i] = 1\n")
    fd.write("\t\t\t}\n")
    fd.write("\t\t} else {\n")
    fd.write('\t\t\tif objVal.String() != "" && objVal.String() != dbObjVal.String() {\n')
    fd.write("\t\t\t\tattrIds[i] = 1\n")
    fd.write("\t\t\t}\n")
    fd.write("\t\t}\n")
    fd.write("\t}\n")
    fd.write("\treturn attrIds, nil\n}\n")

def createMergeDbAndConfigObj(fd, structName, goMemberTypeDict):
    fd.write("func (obj %s) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []byte) (ConfigObj, error) {\n" % (structName))
    fd.write("\tvar mergedStruct %s\n" % (structName)) 
    fd.write("\tobjTyp := reflect.TypeOf(obj)\n")
    fd.write("\tobjVal := reflect.ValueOf(obj)\n")
    fd.write("\tdbObjVal := reflect.ValueOf(dbObj)\n")
    fd.write("\tmergedObjVal := reflect.ValueOf(&mergedStruct)\n")
    fd.write("\tfor i:=1; i<objTyp.NumField(); i++ {\n")
    fd.write("\t\tobjField := objVal.Field(i)\n")
    fd.write("\t\tdbObjField := dbObjVal.Field(i)\n")
    fd.write("\t\tif  attrSet[i] ==1 {\n")
    fd.write("\t\t\tif dbObjField.Kind() == reflect.Int {\n")
    fd.write("\t\t\t\tmergedObjVal.Elem().Field(i).SetInt(objField.Int())\n")
    fd.write("\t\t\t} else {\n")
    fd.write("\t\t\t\tmergedObjVal.Elem().Field(i).SetString(objField.String())\n")
    fd.write("\t\t\t}\n")
    fd.write("\t\t} else {\n")
    fd.write("\t\t\tif dbObjField.Kind() == reflect.Int {\n")
    fd.write("\t\t\t\tmergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())\n")
    fd.write("\t\t\t} else {\n")
    fd.write("\t\t\t\tmergedObjVal.Elem().Field(i).SetString(dbObjField.String())\n")
    fd.write("\t\t\t}\n")
    fd.write("\t\t}\n")
    fd.write("\t}\n")
    fd.write("\treturn mergedStruct, nil\n}\n")

def createUpdateObjectInDb(fd, structName, goMemberTypeDict):
    fd.write("func (obj %s) UpdateObjectInDb(dbObj ConfigObj, attrSet []byte, dbHdl *sql.DB) error {\n" % (structName))
    fd.write("\tvar fieldSqlStr string\n")
    fd.write("\tdbStruct := dbObj.(%s)\n" % (structName))
    fd.write("\tobjKey, err := dbStruct.GetKey()\n")
    fd.write("\tobjSqlKey, err := dbStruct.GetSqlKeyStr(objKey)\n")
    fd.write('\tdbCmd := "update " + "%s" + " set"\n' % (structName))
    fd.write("\tobjTyp := reflect.TypeOf(obj)\n")
    fd.write("\tobjVal := reflect.ValueOf(obj)\n")
    fd.write("\tfor i:=0; i<objTyp.NumField(); i++ {\n")
    fd.write("\t\tif attrSet[i] == 1 {\n")
    fd.write("\t\t\tfieldTyp := objTyp.Field(i)\n")
    fd.write("\t\t\tfieldVal := objVal.Field(i)\n")
    fd.write("\t\t\tif fieldVal.Kind() == reflect.Int {\n")
    fd.write('\t\t\t\tfieldSqlStr = fmt.Sprintf(" %s = %d ", fieldTyp.Name, int(fieldVal.Int()))\n')
    fd.write("\t\t\t} else {\n")
    fd.write('\t\t\t\tfieldSqlStr = fmt.Sprintf(" %s = %s ", fieldTyp.Name, fieldVal.String())\n')
    fd.write("\t\t\t}\n")
    fd.write("\t\t\tdbCmd += fieldSqlStr\n")
    fd.write("\t\t}\n")
    fd.write("\t}\n")
    fd.write('\tdbCmd += " where " + objSqlKey\n')
    fd.write("\t_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)\n")
    fd.write("\treturn err\n}\n")
'''
def createCommonDbFunc(generatePath):

    fd = open(generatePath + "common_db.go", "w")

    fd.write("package models\n")

    fd.write("""import (
             "database/sql"
             "database/sql/driver"
             "fmt"
             )\n""")

    fd.write("""func ConvertBoolToInt(val bool) int {
    if val {
        return 1
    }
    return 0
    }\n""")

    fd.write("""func ConvertStrBoolIntToBool(val string) bool {
    if val == "true" {
        return true
    } else if val == "True" {
        return true
    }
    return false
    }\n""")

    fd.write("""func Depreciated_ExecuteSQLStmt(dbCmd string, dbHdl *sql.DB) (driver.Result, error) {
	var result driver.Result
	txn, err := dbHdl.Begin()
	if err != nil {
		fmt.Println("### Failed to strart db transaction for command", dbCmd)
		return result, err
	}
	result, err = dbHdl.Exec(dbCmd)
	if err != nil {
		fmt.Println("### Failed to execute command ", dbCmd, err)
		return result, err
	}
	err = txn.Commit()
	if err != nil {
		fmt.Println("### Failed to Commit transaction for command", dbCmd, err)
		return result, err
	}
	return result, err
} """)
    fd.close()
    return fd

def createNewAndUnmarshal(fd, currentStruct, goMemberTypeDict):

    if fd:

        fd.write("func New%s() *%s {\n" % (currentStruct, currentStruct))

        # Generic NewFunc, set up the path_helper if asked to.
        fd.write("\tnewObj := &%s{}\n" % (currentStruct))
        fd.write("\treturn newObj\n}\n\n")

        # write unmarshalObject function
        fd.write("""func (obj %s) UnmarshalObject(body []byte) (ConfigObj, error) {
        var err error
        if len(body) > 0 {
            if err = json.Unmarshal(body, &obj); err != nil  {
                fmt.Println("### %s called, unmarshal failed", obj, err)
            }
        }
        return obj, err
        }\n""" %(currentStruct, currentStruct))

def generate_go_sqllite_funcs(fd, directory, gofilename, objectNames=[], goFd=None):

    goMemberTypeDict = {}

    path = os.path.join(directory, gofilename)
    gofd = open(path, 'r')
    deletingComment = False
    foundStruct = False
    currentStruct = None
    keyIdx = 0
    done = False
    for line in gofd.readlines():
        if '//' in line and gofilename == 'ospf.go':
            line = line.split('//')[0]
        if done:
            break
        if not deletingComment:
            if "//" in line:
                line = line.split("//")[0]
            if len(line) == 0:
                continue
            if "struct" in line:
                lineSplit = line.split(" ")
                if objectNames and len([obj for obj in objectNames if obj == lineSplit[1]]) == 0:
                    continue

                currentStruct = lineSplit[1]
                goMemberTypeDict[currentStruct] = []
                foundStruct = True
                keyIdx = 0

            elif "}" in line and foundStruct:
                #foundStruct = False
                keyIdx = 0
                # create the various functions for db
                createNewAndUnmarshal(goFd, currentStruct, goMemberTypeDict)
                createDBTable(fd, currentStruct, goMemberTypeDict)
                createStoreObjInDb(fd, currentStruct, goMemberTypeDict)
                createDeleteObjFromDb(fd, currentStruct, goMemberTypeDict)
                createGetObjFromDb(fd, currentStruct, goMemberTypeDict)
                createGetKey(fd, currentStruct, goMemberTypeDict)
                createGetSqlKeyStr(fd, currentStruct, goMemberTypeDict)
                createGetAllObjFromDb(fd, currentStruct, goMemberTypeDict)
                #createCompareObjectsAndDiff(fd, currentStruct, goMemberTypeDict)
                #createMergeDbAndConfigObj(fd, currentStruct, goMemberTypeDict)
                #createUpdateObjectInDb(fd, currentStruct, goMemberTypeDict)
                createUpdateObjInDb(fd, currentStruct, goMemberTypeDict)
                done = True
            # lets skip all blank lines
            # skip comments
            elif line == '\n' or \
                "#" in line or \
                "package" in line or \
                ("/*" in line and "*/" in line):
                continue
            elif "/*" in line:
                deletingComment = True
            elif foundStruct:  # found element in struct
                #print "found element line", line
                lineSplit = line.split(' ')
                #print lineSplit
                lineSplit = filter(lambda a: a, lineSplit)
                #print "after filtering, split =", lineSplit

                elemtype = lineSplit[-3].rstrip('\n') if 'KEY' in lineSplit[-1] else lineSplit[-1].rstrip('\n')
                key = 0
                if 'KEY' in lineSplit[-1]:
                    keyIdx += 1
                    key = keyIdx

                #print "elemtype:", lineSplit, elemtype
                if elemtype.startswith("[]"):
                    elemtype = elemtype.lstrip("[]")
                    # lets make all list an unordered list
                    nativetype = "LIST " + goToSqlliteTypeMap[elemtype]["native_type"]
                    goMemberTypeDict[currentStruct].append((lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'),
                                                            nativetype, elemtype, key))
                else:
                    if elemtype in goToSqlliteTypeMap.keys():
                        goMemberTypeDict[currentStruct].append((lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'),
                                                                    goToSqlliteTypeMap[elemtype]["native_type"], elemtype, key))

        else:
            if "*/" in line:
                deletingComment = False

    return foundStruct

if __name__ == "__main__":

    parser = argparse.ArgumentParser()
    parser.add_argument('--file', type=str)
    parser.add_argument('--objects', type=str, action='append')
    args = parser.parse_args()
    objects = args.objects
    if not objects:
        objects = get_all_object_names()
    files, generatePath = get_dir_file_names(args.file)
    build_gosqllite_from_go(files, generatePath, objects)
    fd = createCommonDbFunc(generatePath)
    #executeGoFmtCommand(fd, ["gofmt -w %s" % fd.name], GO_MODEL_BASE_PATH)
    #executeLocalCleanup()
