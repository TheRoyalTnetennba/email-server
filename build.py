import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye/go/src/EmailServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('**/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/EmailServer')

build = "go build EmailServer"

processBuild = subprocess.Popen(build.split(), stdout=subprocess.PIPE)
output, error = processBuild.communicate()
