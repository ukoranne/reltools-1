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

            dbFd.write('\nimport (\n\t"database/sql"\n\t"fmt"\n\t"strings"\n)\n')
            generate_gosqllite_funcs(dbFd, directory, gofilename, obj)

            dbFd.close()
            executeGoFmtCommand(dbFd, ["gofmt -w %s" % dbFd.name], GO_MODEL_BASE_PATH)


def createDBTable(fd, structName, goMemberTypeDict):

    createfuncline = '\nfunc (obj %s) CreateDBTable(dbHdl *sql.DB) error {\n' % structName
    fd.write(createfuncline)
    fd.write('\tdbCmd := "CREATE TABLE IF NOT EXISTS %s " +\n' % structName)
    fd.write('\t\t"( " +\n')
    keyList = []
    # loop through member and type
    for m, (t, key) in goMemberTypeDict[structName].iteritems():
        #print "createDBTable key:", m, "value:", t, "key:", key
        if 'Key' in m or key > 0:
            keyList.append((m, key))
        if "LIST" in t:
            fd.write('\t\t" %s TEXT " +\n' %(m,))
        else:
            fd.write('\t\t" %s %s " +\n' %(m, t))

    keyList = sorted(keyList, key=lambda i: i[1])
    fd.write('\t\t "PRIMARY KEY(')
    for i, (k, l) in enumerate(keyList):
        if i == len(keyList) - 1:
            fd.write('%s)' % k)
        else:
            fd.write('%s, ' % k)
    fd.write(' ) "\n')

    fd.write('\n\t_, err := ExecuteSQLStmt(dbCmd, dbHdl)\n')
    fd.write('\treturn err\n')
    fd.write('}\n')

def createStoreObjInDb(fd, structName, goMemberTypeDict):
    storefuncline = "\nfunc (obj %s) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {\n" % structName
    fd.write(storefuncline)
    fd.write('\tvar objectId int64\n')
    fd.write('\tdbCmd := fmt.Sprintf("INSERT INTO %s (' % structName)
    # loop through member and type
    for i, (m, (t, key)) in enumerate(goMemberTypeDict[structName].iteritems()):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""%s) VALUES (""" % m)
        else:
            fd.write("""%s, """ % m )

    for i in range(len(goMemberTypeDict[structName])):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""%v);\",\n\t\t""")
        else:
            fd.write("""%v, """ )

    for i, (m, (t, key)) in enumerate(goMemberTypeDict[structName].iteritems()):
        if i == len(goMemberTypeDict[structName]) - 1:
            fd.write("""obj.%s)\n""" % m )
        else:
            fd.write("""obj.%s, """ % m )

    fd.write("""\tfmt.Println("**** Create Object called with ", obj)

	result, err := ExecuteSQLStmt(dbCmd, dbHdl)
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
    fd.write('\tsqlKey, err := obj.GetSqlKey(objKey)\n')
    fd.write('\tif err != nil {\n')
    fd.write('\t\tfmt.Println("GetSqlKeyStr for %s with key", objKey, "failed with error", err)\n' % (structName))
    fd.write('\t\treturn nil, err\n')
    fd.write('\t}\n\n')
    fd.write('\tdbCmd := "delete from %s where " + sqlKey\n' %(structName))
    fd.write('\tfmt.Println("### DB Deleting %s\\n")\n' % structName)
    fd.write('\t_, err := ExecuteSQLStmt(dbCmd, dbHdl)\n\treturn err\n}\n')

def createGetObjFromDb(fd, structName, goMemberTypeDict):
    storefuncline = "\nfunc (obj %s) GetObjectFromDb(objKey string, dbHdl *sql.DB) (%s, error) {\n" % (structName, structName)
    fd.write(storefuncline)
    fd.write('\tsqlKey, err := obj.GetSqlKey(objKey)\n')
    fd.write('\tif err != nil {\n')
    fd.write('\t\tfmt.Println("GetSqlKeyStr for object key", objKey, "failed with error", err)\n')
    fd.write('\t\treturn nil, err\n')
    fd.write('\t}\n\n')
    fd.write('\tdbCmd := "query from %s where " + sqlKey\n' % (structName))
    fd.write('\tfmt.Println("### DB Get %s\\n")\n' % structName)
    fd.write('\tobj, err := ExecuteSQLStmt(dbCmd, dbHdl)\n\treturn obj, err\n}\n')

    
def createGetKey(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) GetKey() (string, error) {" % structName)
    keys = sorted([(m, key) for m, (t, key) in goMemberTypeDict[structName].iteritems() if key], key=lambda i: i[1])
    objKey = ' + "#" + '.join(['string(obj.%s)' % (m) for m, key in keys])
    if objKey:
        fd.write('\n\tkey := ')
        fd.write(objKey)

    fd.write("\n\treturn key, nil\n}\n")

def createGetSqlKey(fd, structName, goMemberTypeDict):
    fd.write("\nfunc (obj %s) GetSqlKey(objKey string) (string, error) {\n" % structName)
    fd.write('\tkeys := strings.Split(objKey, "#")')
    #print "struct dict =", goMemberTypeDict[structName]
    keys = sorted([(m, key) for m, (t, key) in goMemberTypeDict[structName].iteritems() if key], key=lambda i: i[1])

    firstKey = ['" = + \\\" + "'.join(['"%s"' % (m), 'keys[%d]' % (i)]) for i, (m, key) in enumerate(keys)]
    #print "firstKey =", firstKey
    sqlKey = ' + '.join(['+ "\\\"" + '.join(['"%s = "' % (m), 'keys[%d] + "\\\""' % (i)]) for i, (m, key) in enumerate(keys)])
    fd.write('\n\tsqlKey := ')
    fd.write(sqlKey)

    fd.write("\n\treturn sqlKey, nil\n}\n")

def createCommonDbFunc(generatePath):

    fd = open(generatePath + "common_db.go", "w")

    fd.write("package models\n")

    fd.write("""import (
             "database/sql"
             "database/sql/driver"
             "fmt"
             )\n""")

    fd.write("""func ExecuteSQLStmt(dbCmd string, dbHdl *sql.DB) (driver.Result, error) {
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
                goMemberTypeDict[currentStruct] = {}
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
                createGetSqlKey(fd, currentStruct, goMemberTypeDict)

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
                    goMemberTypeDict[currentStruct].update({lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'):
                                                            (nativetype, key)})
                else:
                    if elemtype in goToSqlliteTypeMap.keys():
                        goMemberTypeDict[currentStruct].update({lineSplit[0].lstrip(' ').rstrip(' ').lstrip('\t'):
                                                                    (goToSqlliteTypeMap[elemtype]["native_type"], key)})

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
