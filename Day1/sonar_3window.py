depth_check = 0
depth = 150 # Starting depth because I was lazy
window = []
nxt1 = 1
nxt2 = 2

with open("input") as file_in:
    lines = []
    for line in file_in:
        lines.append(line)
numlines = len(lines)-1
for x in range(numlines): 
    if (x+1) > numlines:
        nxt1=0
    else:
        nxt1 = lines[x+1]
    if (x+2) > numlines:
        nxt2=0
    else:
        nxt2=lines[x+2]
    sum = int(lines[x]) + int(nxt1) + int(nxt2)
    window.append(sum)

depth=window[0]
for x in window:
    if int(x)>depth:
        depth_check +=1
    depth=int(x)

print(depth_check)