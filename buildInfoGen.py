import os
import subprocess
from optparse import OptionParser
import sys
import json

SNAP_ROUTE_SRC  = '/snaproute/src/'
SRC_INFO_FILE   = 'srcInfo.json'
BUILD_INFO_FILE = 'buildInfo.json'

def executeCommand (command) :
    out = ''
    if type(command) != list:
        command = [ command]
    for cmd in command:
        process = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE)
        out,err = process.communicate()
    return out

class repo (object):
    def __init__ (self, repo):
        self.name = repo

    def writeRepoInfo (self) :
        op = executeCommand('git show')
        gitHash = ''
        timeStamp = ''
        branch = ''
        for line in op.split('\n'):
            if 'commit' in line.split():
                gitHash = line.split()[1]
            elif 'Date:' in line.split():
                timeStamp = ' '.join(line.split()[1:])

        branch = executeCommand('git rev-parse --abbrev-ref HEAD')
        repoInfo = {'Name' : self.name,
                    'Sha1' : gitHash,
                    'Time' : timeStamp,
                    'Branch' : branch.rstrip('\n')
                    }
        return repoInfo

if __name__ == '__main__':
    repos = []
    with open(SRC_INFO_FILE) as infoFd:
        info = json.load(infoFd)
        for rp in info ['repos'] ['snaproute'] ['package']:
            repos.append(repo(rp))
    reposInfoList = []
    with open(BUILD_INFO_FILE, 'w') as bldFile: 
        for rp in repos:
            reposInfoList.append(rp.writeRepoInfo())
        json.dump(reposInfoList, bldFile, indent=4, separators=(',', ': '), sort_keys=False)

