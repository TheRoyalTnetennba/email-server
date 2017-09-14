import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye/go/src/EmailServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('./*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/EmailServer')

pkg_test = False

commands = ["go build EmailServer"]

if pkg_test:
    commmands.insert(0, "git push")
    commmands.insert(0, "git commit -am.")
    commmands.insert(0, "git add -A")


for command in commands:
    processBuild = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = processBuild.communicate()
