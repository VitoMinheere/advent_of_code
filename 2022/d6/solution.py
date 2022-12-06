def p1(data):
    pass

def p2(data):
    pass

with open("test.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)