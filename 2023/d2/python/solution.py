import time
import re

MAX_CUBES = {
   "red": 12,
   "green": 13,
   "blue": 14
}

def p1(data):
    answer = 0

    for line in data:
        game_id, revealed = line.split(": ")
        game_id = int(game_id.split(" ")[-1])
        revealed = revealed.split(";")

        possible = True

        for part in revealed:
            cubes = [x.lstrip() for x in part.split(",")]
            for item in cubes:
                amount, color = item.split(" ")
                if int(amount) > MAX_CUBES[color]:
                    print(f"{game_id} is impossible")
                    possible = False
                    break

        if possible:
            answer += game_id

    return answer


def p2(data):
    pass

if __name__ == "__main__":
    with open("../input") as t:
        data = t.read().splitlines()

        start_p1 = time.time()
        answer_1 = p1(data)
        print(f"P1 answer = {answer_1}")
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        # start_p2 = time.time()
        # answer_2 = p2(data)
        # print(f"P2 answer = {answer_2}")
        # p2_time = time.time() - start_p2
        # print("P2 took " + str(round(p2_time * 1000)) + " ms")