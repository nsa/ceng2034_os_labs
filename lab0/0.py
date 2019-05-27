import os

myfolder = os.getcwd()

print(os.system("clear"))
print(myfolder)

print(os.listdir())

print("second index", os.listdir()[1])

#print("_______________________________")
#print("ls -la ", os.system('ls -la'))

print("is it directory? ",os.path.isdir(myfolder))

mypath = os.path.join(myfolder , "another")

print(mypath)

print("is it dir ?", os.path.isdir(mypath))
