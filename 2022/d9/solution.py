import copy 
import time

def p1(data):
    matrix = [[0 for x in range(0, 1000)] for _ in range(0, 1000)]

    H = {"row": -1, "col": 0, "pos": matrix[-1][0]}
    T = {"row": -1, "col": 0, "pos": matrix[-1][0]}

    matrix[-1][0] = 1

    for line in data:
        instructions = line.split(" ")

        for _ in range(0, int(instructions[1])):
            last_h_pos = (H["row"], H["col"])
        
            if instructions[0] == "R":
                H["col"] += 1 
            elif instructions[0] == "L":
                H["col"] -= 1 
            elif instructions[0] == "U":
                H["row"] -= 1 
            else: # Down
                H["row"] += 1 
                
            H["pos"] = matrix[H["row"]][H["col"]]
            row_dif = H["row"] - T["row"]
            col_dif = H["col"] - T["col"]

            if abs(row_dif) > 1 or abs(col_dif) > 1:
                T["row"], T["col"] = last_h_pos
                matrix[T["row"]][T["col"]] = 1 # Mark spot for Tail

    p1 = 0
    for row in matrix:
        p1 += sum(row)

    print("P1 = " + str(p1))
                

def p2(data):

    def move(knot, instructions):
        # print("Instruction = " + instructions)
        if instructions == "R":
            knot["col"] += 1
        elif instructions == "L":
            knot["col"] -= 1
        elif instructions == "U":
            knot["row"] -= 1
        else: # Down
            knot["row"] += 1
        return knot
        
    matrix = [[0 for x in range(0, 1000)] for _ in range(0, 1000)]
    rope = [ {"row": 17, "col": 12, "prev_col": 17, "prev_row": 12, "pos": matrix[17][12]} for _ in range(0, 10)]

    for line in data:
        instructions = line.split(" ")

        # Move Head first
        for _ in range(0, int(instructions[1])):

            rope[0] = move(rope[0], instructions[0])
            rope[0]["pos"] = matrix[rope[0]["row"]][rope[0]["col"]]

            for i, knot in enumerate(rope):
                if i == 0:
                    continue

                row_dif = knot["row"] - rope[i-1]["row"]
                col_dif = knot["col"] - rope[i-1]["col"]

                if abs(row_dif) <= 1 and abs(col_dif) <= 1:
                    pass
                elif knot["col"] == rope[i-1]["col"]:
                    if knot["row"] < rope[i-1]["row"]:
                        knot = move(knot, "D")
                    else:
                        knot = move(knot, "U")

                elif knot["row"] == rope[i-1]["row"]:
                    if knot["col"] < rope[i-1]["col"]:
                        knot = move(knot, "R")
                    else:
                        knot = move(knot, "L")
                else:
                    if knot["col"] < rope[i-1]["col"]:
                        knot = move(knot, "R")
                    else:
                        knot = move(knot, "L")
                    
                    if knot["row"] < rope[i-1]["row"]:
                        knot = move(knot, "D")
                    else:
                        knot = move(knot, "U")
                    
         
                if (i) == 9:
                    matrix[knot["row"]][knot["col"]] = 1 
     
    p2 = 0
    for row in matrix:
        p2 += sum(row)

    print("P2 = " + str(p2))


with open("input.txt") as t:
    data = t.read().splitlines()
    start = time.time()
    p1(data)
    p1_time = time.time() - start
    p2_start = time.time()
    p2(data)
    p2_time = time.time() - p2_start

    print("P1 took " + str(round(p1_time * 1000)) + " ms")
    print("P2 took " + str(round(p2_time * 1000)) + " ms")