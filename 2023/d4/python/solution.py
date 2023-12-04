import time

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
    answer = 0

    hand = {}
    cards = {} 
    
    for i, line in enumerate(data):
        card_number, rest = line.split(": ")
        win_nums, own_nums = rest.split("|")
        win_nums = [x for x in win_nums.replace("  ", " ").split(" ") if x.isnumeric()]
        own_nums = [x for x in own_nums.replace("  ", " ").split(" ") if x.isnumeric()]

        points_in_card = len([x for x in own_nums if x in win_nums])

        cards[i] = {"copies": 1, "wins": points_in_card}
    
    for i, v in cards.items():
        for y in range(v["copies"]):
            answer += 1
            for w in range(i, i+v["wins"]):
                cards[w+1]["copies"] += 1

   
    return answer

if __name__ == "__main__":
    with open("../input") as t:
        data = t.read().splitlines()

        start_p1 = time.time()
        answer_1 = p1(data)
        print(f"P1 answer = {answer_1}")
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        start_p2 = time.time()
        answer_2 = p2(data)
        print(f"P2 answer = {answer_2}")
        p2_time = time.time() - start_p2
        print("P2 took " + str(round(p2_time * 1000)) + " ms")