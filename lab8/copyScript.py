import shutil, random

for i in range (10,150):
    name = "{0}{1}".format("file", i)
    origin = "{0}{1}".format("file", random.randint(1,4))
    shutil.copyfile(origin, name )
