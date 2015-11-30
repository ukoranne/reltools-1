import os
import json

HOME = os.getenv("HOME")
MODEL_NAME = 'genmodels'
GO_MODEL_BASE_PATH = HOME + "/git/generated/src/%s/" % MODEL_NAME

def scan_dir_for_go_files(dir):
    for name in os.listdir(dir):
        #print "x", dir, name
        path = os.path.join(dir, name)
        if name.endswith('.go'):
            if os.path.isfile(path) and "_enum" not in path and "_func" not in path and "_db" not in path:
                yield (dir, name)
        elif not "." in name:
            for d, f  in scan_dir_for_go_files(path):
                yield (d, f)

def build_json_object_from_go():
    # generate thrift code from go code
    for dir, gofilename in scan_dir_for_go_files(GO_MODEL_BASE_PATH):
        jsonFileName = gofilename.split('.')[0] + ".json"

        data = {}
        with open(jsonFileName, 'w') as f:

            path = os.path.join(dir, gofilename)
            gofd = open(path, 'r')
            deletingComment = False
            for line in gofd.readlines():
                if line.startswith("type") and \
                    "struct" in line and \
                    "Config" in line:
                    structName = line.split(" ")[1]
                    data.update({structName : {"Owner": "",
                                               "Listeners": []}})
            json.dump(data, f, indent=4)

if __name__ == "__main__":


    build_json_object_from_go()
