def p1(data):
    for line in data:
        start = line[0:3]
        read = start
        for x in range(3, len(line)):
            if line[x] not in read and (len(read) == len(set(read))):
                print("Found mark ", line[x], " at char " + str(x + 1))
                print("P1 = " + str(x+1))
                read += line[x]
                break
            else:
                read += line[x]
                read = read[1:]
        

def p2(data):
    for line in data:
        start = line[0:13]
        read = start
        for x in range(13, len(line)):
            if line[x] not in read and (len(read) == len(set(read))):
                print("Found mark ", line[x], " at char " + str(x + 1))
                print("P2 = " + str(x+1))
                read += line[x]
                break
            else:
                read += line[x]
                read = read[1:]

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)