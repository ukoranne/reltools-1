import os

GO_MODEL_BASE_PATH = "/home/ccordes/git/snaproute/generated/src/gomodel/"

def scan_dir_for_go_files(dir):
    for name in os.listdir(dir):
        #print "x", dir, name
        path = os.path.join(dir, name)
        if name.endswith('.go'):
            if os.path.isfile(path) and "enum" not in path and "func" not in path:
                yield (dir, name)
        elif not "." in name:
            for d, f  in scan_dir_for_go_files(path):
                yield (d, f)

def build_json_object_from_go():
    # generate thrift code from go code
    for dir, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH):
        print dir, gofilename, dir.split('/')[-1]
        jsonFileName = gofilename.split('.')[0] + ".json"
        jsonfd = open(jsonFileName, 'w')

        path = os.path.join(dir, gofilename)
        print 'path', path
        gofd = open(path, 'r')
        jsonfd.write("{\n")
        deletingComment = False
        for line in gofd.readlines():
            if line.startswith("type") and \
                "struct" in line and \
                "Config" in line:
                structName = line.split(" ")[1]
                print "found struct", structName

                jsonfd.write(""" \"%s\"\t: {"Owner" : \"\",\n""" % structName)
                jsonfd.write("""            "Listeners" : []},\n""")
        jsonfd.write("}")
        jsonfd.close()

if __name__ == "__main__":


    build_json_object_from_go()
