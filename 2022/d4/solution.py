def contains(a, b):
    a_start, a_end =  a.split("-")
    b_start, b_end =  b.split("-")

    a_start, a_end, b_start, b_end = int(a_start), int(a_end), int(b_start), int(b_end)

    if a_start >= b_start and a_end <= b_end:
        return True
    if b_start >= a_start and b_end <= a_end:
        return True
    return False

def p1(data):
    p1 = 0

    for line in data:
        first, second = line.split(",")
        if contains(first, second):
            p1 += 1
        
    print("P1 is " + str(p1))

def p2(data):
    p2 = 0
    
    for line in data:
        first, second = line.split(",")
        
        range_first = [str(x) for x in range(int(first.split("-")[0]),
                                                     int(first.split("-")[-1])+1)]
        range_second = [str(x) for x in range(int(second.split("-")[0]), 
                                                      int(second.split("-")[-1])+1)]

        overlap = [x for x in range_first if x in range_second]
        if overlap:
            print(overlap[0])
            p2 += 1

    print("P2 is " + str(p2))

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)