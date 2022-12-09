def p1(data):
    matrix = [[0 for x in range(0, 1000)] for _ in range(0, 1000)]

    H = {"row": -1, "col": 0, "pos": matrix[-1][0]}
    T = {"row": -1, "col": 0, "pos": matrix[-1][0]}

    matrix[-1][0] = 1

    # data = data[:2]
    for line in data:
        # print(line)
        # print("########")
        instructions = line.split(" ")

        # Move Head first
        for _ in range(0, int(instructions[1])):
            last_h_pos = (H["row"], H["col"])
        
            if instructions[0] == "R":
                H["col"] += 1 #int(instructions[1])
            elif instructions[0] == "L":
                H["col"] -= 1 #int(instructions[1])
            elif instructions[0] == "U":
                H["row"] -= 1 #int(instructions[1])
            else: # Down
                H["row"] += 1 #int(instructions[1])
                
            H["pos"] = matrix[H["row"]][H["col"]]
            # matrix[H["row"]][H["col"]] = "H" # Mark spot for HEad
            # print(f"HEAD is at {H['row']} {H['col']}")
            # print(f"TAIL is at {T['row']} {T['col']}")
            row_dif = H["row"] - T["row"]
            col_dif = H["col"] - T["col"]
            # print("row dif " + str(row_dif))
            # print("col dif " + str(col_dif))

            if abs(row_dif) > 1 or abs(col_dif) > 1:
                T["row"], T["col"] = last_h_pos
                # matrix[T["row"]][T["col"]] = "T" # Mark spot for Tail
                matrix[T["row"]][T["col"]] = 1 # Mark spot for Tail
                
            # if abs(row_dif) > 1:
            #     step = (-1,1)[row_dif > 0]
            #     # for _ in range(0, abs(row_dif)):
            #     T["row"] += step
            #     T["pos"] = matrix[T["row"]][T["col"]]
            #     matrix[T["row"]][T["col"]] = "T" # Mark spot for Tail
            # if abs(col_dif) > 1:
            #     # for _ in range(0, abs(col_dif)):
            #     step = (-1,1)[col_dif > 0]
            #     T["col"] += 1
            #     T["pos"] = matrix[T["row"]][T["col"]]
            #     matrix[T["row"]][T["col"]] = "T" # Mark spot for Tail

            # for row in matrix:
            #     print(row)
            # m = matrix[::-1]
            # for row in m:
            #     r = row[::-1]
            #     print(r)
            # print("                 ")

    p1 = 0
    for row in matrix:
        # print(row)
        p1 += sum(row)
    # matrix.reverse()
    # for row in matrix:
    #     row.reverse()
    #     print(row)
    #     p1 += sum(row)

    print("P1 = " + str(p1))

             
                

def p2(data):
    pass


with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    # p2(data)