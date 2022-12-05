def p1(data):
    test_matrix = [
        ["Z", "N"],
        ["M", "C", "D"],
        ["P"]
    ]

    matrix = [
        ["Q", "M", "G", "C", "L"],
        ["R", "D", "L", "C", "T", "F", "H", "G"],
        ["V", "J", "F", "N", "M", "T", "W", "R"],
        ["J", "F", "D", "V", "Q", "P"],
        ["N", "F", "M", "S", "L", "B", "T"],
        ["R", "N", "V", "H", "C", "D", "P"],
        ["H", "C", "T"],
        ["G", "S", "J", "V", "Z", "N", "H", "P"],
        ["Z", "F", "H", "G"],
    ]

    code = ""

    for line in data:
        s = line.split(" ")
        am_crates = int(s[1])
        stack = int(s[3]) - 1
        dest = int(s[5]) - 1

        for _ in range(0, am_crates):
            origin = matrix[stack]
            destination = matrix[dest]
            destination.append(origin.pop(-1)) 

    for x in matrix:
        code += x[-1]

    print("P1 = " + code)
            

def p2(data):
    pass

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)