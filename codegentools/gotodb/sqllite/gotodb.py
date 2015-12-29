import os
import subprocess
import json
import pprint
import argparse
import itertools


IGNORE_GO_FILE_LIST = ["objectmap.go"]

#HOME = os.getenv("HOME")
srBase = os.environ.get('SR_CODE_BASE', None)
GO_MODEL_BASE_PATH = srBase + "/generated/src/models/"
CODE_GENERATION_PATH = srBase + "/generated/src/models/db/"

goToSqlliteTypeMap = {
  'bool':          {"native_type": "bool"},
  'uint8':            {"native_type": "INTEGER", "unsigned": True},
  'uint16':           {"native_type": "INTEGER", "unsigned": True},
  'uint32':           {"native_type": "INTEGER", "unsigned": True},
  'uint64':           {"native_type": "INTEGER", "unsigned": True},
  'string':           {"native_type": "TEXT","unsigned": None },
  'float64':          {"native_type": "REAL", "unsigned": False},
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
            #nfd = open(fmt_name_with_dir, 'w+')
            #nfd.write(out)
            #nfd.close()

            #process = subprocess.Popen("ls".split(), stdout=subprocess.PIPE)
            #out,err = process.communicate()
            #print out, err

            #renameCmd = "mv %s %s" %(fmt_name_with_dir, fd.name)
            #process = subprocess.Popen(renameCmd.split(), stdout=subprocess.PIPE)
            #out,err = process.communicate()
            #print out, err

            #out = executeCopyCommand(fd.name, dstPath)

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

def build_gosqllite_from_go(files, generatePath, objects):
    # generate thrift code from go code
    goStructToListersDict = {}
    
    # lets determine from the json file the structs and associated listeners
    #for directory, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH):
    for directory, gofilename in files:
        if '_func' in gofilename and '_enum' in gofilename and '_db' in gofilename:
            continue

        if gofilename in IGNORE_GO_FILE_LIST:
            continue

        for obj in (itertools.combinations(objects, 1) if objects else [[]]):
            dbFileName = generatePath + (obj[0].lower() if obj else gofilename.rstrip('.go')) + "dbif.go"
            if not os.path.exists(generatePath):
                os.makedirs(generatePath)

            #print "generate file name =", dbFileName, "path =", generatePath

            dbFd = open(dbFileName, 'w')
            dbFd.write("package models\n")

            dbFd.write('\nimport (\n\t"database/sql"\n\t"fmt"\n\t"strings"\n\t"utils/dbutils"\n\t"reflect"\n)\n')
            generate_gosqllite_funcs(dbFd, directory, gofilename, obj)

            dbFd.close()
            executeGoFmtCommand(dbFd, ["gofmt -w %s" % dbFd.name], GO_MODEL_BASE_PATH)


def createDBTable(fd, structName, goMemberTypeDict):

    createfuncline = '\nfunc (obj %s) CreateDBTable(dbHdl *sql.DB) error {\n' % structName
    fd.write(createfuncline)
    fd.write('\tdbCmd := "CREATE TABLE IF NOT EXISTS %s " +\n' % structName)
    fd.write('\t\t"( " +')
    keyList = []
    # loop through member and type
    for i, (m, t, key) in enumerate(goMemberTypeDict[structName]):
        #print "createDBTable key:", m, "value:", t, "key:", key
        if 'Key' in m or key > 0:
            keyList.append((m, key))
        if "LIST" in t:
            fd.write('\n\t\t"%s TEXT, " +' %(m,))
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
    for i, (m, t, key) in enumerate(goMemberTypeDict[structName]):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""%s) VALUES (""" % m)
        else:
            fd.write("""%s, """ % m )

    for i in range(len(goMemberTypeDict[structName])):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""'%v') ;\",\n\t\t""")
        else:
            fd.write("""'%v', """ )

    for i, (m, t, key) in enumerate(goMemberTypeDict[structName]):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""obj.%s)\n""" % m )
        else:
            fd.write("""obj.%s, """ % m )

    fd.write("""\tfmt.Println("**** Create Object called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	}\n
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to return last object id", err)
	}\n
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
    fd.write('\tif err != nil {\n')
    fd.write('\t\tfmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)\n')
    fd.write('\t\treturn object, err\n')
    fd.write('\t}\n\n')
    fd.write('\tdbCmd := "SELECT * from %s where " + sqlKey\n' % (structName))
    fd.write('\tfmt.Println("### DB Get %s\\n")\n' % structName)
    fd.write('\terr = dbHdl.QueryRow(dbCmd).Scan(%s)\n' % (', '.join(['&object.%s' % (m) for m, t, key in goMemberTypeDict[structName]])))
    fd.write('\treturn object, err\n}\n')

    
def createGetKey(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) GetKey() (string, error) {" % structName)
    keys = sorted([(m, key) for m, t, key in goMemberTypeDict[structName] if key], key=lambda i: i[1])
    objKey = ' + "#" + '.join(['string(obj.%s)' % (m) for m, key in keys])
    if objKey:
        fd.write('\n\tkey := ')
        fd.write(objKey)

    fd.write("\n\treturn key, nil\n}\n")

def createGetSqlKeyStr(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) GetSqlKeyStr(objKey string) (string, error) {\n" % structName)
    #print "struct dict =", goMemberTypeDict[structName]
    keys = sorted([(m, key) for m, t, key in goMemberTypeDict[structName] if key], key=lambda i: i[1])
    if keys:
        fd.write('\tkeys := strings.Split(objKey, "#")')
        firstKey = ['" = + \\\" + "'.join(['"%s"' % (m), 'keys[%d]' % (i)]) for i, (m, key) in enumerate(keys)]
        #print "firstKey =", firstKey
        sqlKey = ' + "and" + '.join(['+ "\\\"" + '.join(['"%s = "' % (m), 'keys[%d] + "\\\""' % (i)]) for i, (m, key) in enumerate(keys)])
        fd.write('\n\tsqlKey := ')
        fd.write(sqlKey)
        fd.write("\n\treturn sqlKey, nil\n}\n")
    else:
        fd.write("""\n\treturn "", nil\n}\n""")



def createUpdateObjInDb(fd, structName, goMemberTypeDict):
    fd.write("""
    func (obj %s) CompareObjectsAndDiff(dbObj ConfigObj) ([]byte, error) {
	dbV4Route := dbObj.(%s)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]byte, objTyp.NumField())
	for i:=0; i<objTyp.NumField(); i++ {
		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if objVal.Kind() == reflect.Int {
		    if int(objVal.Int()) != 0 && int(objVal.Int()) != int(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int8 {
		    if int8(objVal.Int()) != 0 && int8(objVal.Int()) != int8(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int16 {
		    if int16(objVal.Int()) != 0 && int16(objVal.Int()) != int16(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int32 {
		    if int32(objVal.Int()) != 0 && int32(objVal.Int()) != int32(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Int64 {
			if int64(objVal.Int()) != 0 && int64(objVal.Int()) != int64(dbObjVal.Int()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Uint {
			if uint(objVal.Uint()) != 0 && uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
        } else if objVal.Kind() == reflect.Uint8 {
			if uint8(objVal.Uint()) != 0 && uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
        } else if objVal.Kind() == reflect.Uint16 {
        	if uint16(objVal.Uint()) != 0 && uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
        } else if objVal.Kind() == reflect.Uint32 {
			if uint16(objVal.Uint()) != 0 && uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
        } else if objVal.Kind() == reflect.Uint64 {
			if uint16(objVal.Uint()) != 0 && uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
				attrIds[i] = 1
			}
		} else if objVal.Kind() == reflect.Bool {
		    if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
		        attrIds[i] = 1
		    }
		} else {
			if objVal.String() != "" && objVal.String() != dbObjVal.String() {
				attrIds[i] = 1
			}
		}
	}
	return attrIds, nil
}\n""" %( structName, structName))

    fd.write("""
    func (obj %s) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []byte) (ConfigObj, error) {
	var merged%s %s
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&merged%s)
	for i:=1; i<objTyp.NumField(); i++ {
		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if  attrSet[i] ==1 {
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
	}
	return merged%s, nil
}\n""" %(structName, structName, structName, structName, structName))

    fd.write("""
    func (obj %s) UpdateObjectInDb(dbObj ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	var fieldSqlStr string
	db%s := dbObj.(%s)
	objKey, err := db%s.GetKey()
	objSqlKey, err := db%s.GetSqlKeyStr(objKey)
	dbCmd := "update " + "%s" + " set"\n""" %(structName, structName, structName, structName, structName, structName))

    fd.write("""objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	for i:=0; i<objTyp.NumField(); i++ {
		if attrSet[i] == 1 {
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
			    fieldSqlStr = fmt.Sprintf(" %s = '%t' ", fieldTyp.Name, bool(fieldVal.Bool()))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}\n""")

def createCommonDbFunc(generatePath):

    fd = open(generatePath + "common_db.go", "w")

    fd.write("package models\n")

    fd.write("""import (
             "database/sql"
             "database/sql/driver"
             "fmt"
             )\n""")

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

def generate_gosqllite_funcs(fd, directory, gofilename, objectNames=[]):

    goMemberTypeDict = {}

    path = os.path.join(directory, gofilename)
    gofd = open(path, 'r')
    deletingComment = False
    foundStruct = False
    currentStruct = None
    keyIdx = 0
    for line in gofd.readlines():
        if not deletingComment:
            if "struct" in line:
                if objectNames and len([obj for obj in objectNames if obj in line]) == 0:
                    continue

                lineSplit = line.split(" ")
                currentStruct = lineSplit[1]
                goMemberTypeDict[currentStruct] = []
                foundStruct = True
                keyIdx = 0

            elif "}" in line and foundStruct:
                foundStruct = False
                keyIdx = 0
                # create the various functions for db
                createDBTable(fd, currentStruct, goMemberTypeDict)
                createStoreObjInDb(fd, currentStruct, goMemberTypeDict)
                createDeleteObjFromDb(fd, currentStruct, goMemberTypeDict)
                createGetObjFromDb(fd, currentStruct, goMemberTypeDict)
                createGetKey(fd, currentStruct, goMemberTypeDict)
                createGetSqlKeyStr(fd, currentStruct, goMemberTypeDict)
                createUpdateObjInDb(fd, currentStruct, goMemberTypeDict)

            # lets skip all blank lines
            # skip comments
            elif line == '\n' or \
                "//" in line or \
                "#" in line or \
                "package" in line:
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
                                                            nativetype, key))
                else:
                    if elemtype in goToSqlliteTypeMap.keys():
                        goMemberTypeDict[currentStruct].append((lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'),
                                                                    goToSqlliteTypeMap[elemtype]["native_type"], key))

        else:
            if "*/" in line:
                deletingComment = False


if __name__ == "__main__":

    parser = argparse.ArgumentParser()
    parser.add_argument('--file', type=str)
    parser.add_argument('--objects', type=str, action='append')
    args = parser.parse_args()

    files, generatePath = get_dir_file_names(args.file)
    build_gosqllite_from_go(files, generatePath, args.objects)
    fd = createCommonDbFunc(generatePath)
    executeGoFmtCommand(fd, ["gofmt -w %s" % fd.name], GO_MODEL_BASE_PATH)
    #executeLocalCleanup()
