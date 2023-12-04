import time
import re
from functools import reduce

def p1(data):
    answer = 0
    
    for line in data:
        card_number, rest = line.split(": ")
        win_nums, own_nums = rest.split("|")
        win_nums = [x for x in win_nums.replace("  ", " ").split(" ") if x.isnumeric()]
        own_nums = [x for x in own_nums.replace("  ", " ").split(" ") if x.isnumeric()]

        points_in_card = [x for x in own_nums if x in win_nums]
        if len(points_in_card) > 1:
            answer += (1*2) ** (len(points_in_card) - 1)
        elif len(points_in_card) == 1:
            answer += 1
        else:
            pass
   
    return answer

def p2(data):
    pass

if __name__ == "__main__":
    with open("../input") as t:
    # with open("2023/d3/input") as t:
        data = t.read().splitlines()

        start_p1 = time.time()
        answer_1 = p1(data)
        print(f"P1 answer = {answer_1}")
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        # start_p2 = time.time()
        # answer_2 = p2(data, True)
        # print(f"P2 answer = {answer_2}")
        # p2_time = time.time() - start_p2
        # print("P2 took " + str(round(p2_time * 1000)) + " ms")