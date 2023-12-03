import time
import re
from functools import reduce

def solution(data, p2=False):
    m = []
    parts_list = []

    for line in data:
        m.append(line)

    directions = [
        (0, 1),     # left
        (1, 1),   # top left
        (1, 0),    # top
        (1, -1),    # top right
        (0, -1),     # right 
        (-1, -1),     # bottom right
        (-1, 0),     # bottom
        (-1, 1),    # bottom left
    ]

    gears = {}
    for r_index, row in enumerate(m):
        # print(row)
        for c_index, col in enumerate(row):
            if not col.isnumeric() and col != ".":
                pos = (r_index, c_index)
                parts_found = []
                row_used = []
                gears[f"{col}_{str(pos)}"] = []
                for d in directions:
                    # if d[0] in row_used:
                    #     continue
                    new_pos = tuple(map(lambda i, j: i - j, pos, d))
                    val = m[new_pos[0]][new_pos[1]]
                    # print(val)
                    if val.isnumeric():
                        # print(m[new_pos[0]])
                        part_num = []
                        # Found value, check sides to get the whole number
                        part_num.append(val)
                        left = m[new_pos[0]][new_pos[1]-1]
                        right = m[new_pos[0]][new_pos[1]+1]
                        if not left.isnumeric() and right.isnumeric():
                            try:
                                print("right")
                                # Then go right
                                i = 1
                                while val.isnumeric() and val != ".":
                                    val = m[new_pos[0]][new_pos[1]+i]
                                    print(val)
                                    part_num.append(val)
                                    i += 1
                                    val = m[new_pos[0]][new_pos[1]+i]
                            except IndexError:
                                part_number = int("".join(part_num))
                                if part_number not in parts_found:
                                    row_used.append(d[0])
                                    parts_found.append(part_number)
                                    gears[f"{col}_{str(pos)}"].append(part_number)
                                    parts_list.append(part_number)
                                continue

                        elif left.isnumeric() and not right.isnumeric():
                            try:
                                print("left")
                                i = 1
                                val = m[new_pos[0]][new_pos[1]-i]
                                while val.isnumeric() and val != ".":
                                    print(val)
                                    part_num.append(val)
                                    i += 1
                                    val = m[new_pos[0]][new_pos[1]-i]
                            except IndexError:
                                continue
                            finally:
                                part_num.reverse()

                        elif left.isnumeric() and right.isnumeric():
                            print("middle")
                            # You are in the middle
                            part_num.insert(0, left)
                            part_num.append(right)

                        part_number = int("".join(part_num))
                        if part_number not in parts_found:
                            row_used.append(d[0])
                            parts_found.append(part_number)
                            gears[f"{col}_{str(pos)}"].append(part_number)
                            parts_list.append(part_number)
                        # break
                            
    if p2:
        answer = 0
        for x in gears:
            if x.startswith("*"):
                g = gears[x]
                if len(g) == 2:
                    answer += (gears[x][0] * gears[x][1])
        return answer

    answer = sum(parts_list)
    return answer

def p2(data):
    pass

if __name__ == "__main__":
    # with open("../test") as t:
    with open("2023/d3/input") as t:
        data = t.read().splitlines()

        # start_p1 = time.time()
        # answer_1 = solution(data)
        # print(f"P1 answer = {answer_1}")
        # p1_time = time.time() - start_p1
        # print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        start_p2 = time.time()
        answer_2 = solution(data, True)
        print(f"P2 answer = {answer_2}")
        p2_time = time.time() - start_p2
        print("P2 took " + str(round(p2_time * 1000)) + " ms")