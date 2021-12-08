depth_check = 0
depth = 150 # Starting depth because I was lazy
f = open("input", "r")
for x in f:
  if  int(x) > depth:
      depth_check += 1 
  depth = int(x)
f.close() 

print(depth_check)