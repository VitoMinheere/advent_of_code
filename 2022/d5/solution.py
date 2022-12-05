def p1(data):
    stacks = {}
    matrix = []

    for line in data:
        # print(line)
        items = line.split("  ")
        if "move" not in items[0]:
            cur_stack = 0
            for crate in items:
                if crate not in matrix and "[" in "".join(crate):
                    matrix.insert(0, crate)
                # for x in crate.split(" "):
                #     if x == "":
                #         # print("1234")
                #         cur_stack += 1
                #     elif "[" in crate:
                #         stack = cur_stack // 2
                #         if crate not in matrix:
                #             matrix.insert(0, crate)
                #         if (stack) not in stacks:
                #             stacks[stack] = []
                #         # print(cur_stack)
                #         # print(crate)
                #         if crate not in stacks[stack]:
                #             stacks[stack].insert(0, crate)
                        
                # print(crate.split(" "))
    print(stacks)
    print(matrix)
            

def p2(data):
    pass

with open("test.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)