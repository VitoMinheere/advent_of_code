import time
from functools import reduce



def p1(data):
    times = []
    distances = []
    for line in data:
        if line.startswith("Time"):
            times = [int(x) for x in line.split(": ")[-1].split(" ") if x.isnumeric()]
        elif line.startswith("Distance"):
            distances = [int(x) for x in line.split(": ")[-1].split(" ") if x.isnumeric()]

    options_per_race = []
    for i in range(0, len(times)):

        winning_options = 0
        for s in range(1, times[i]+1):
            speed = 1 * s
            time_to_move = times[i] - s
            dist = speed * time_to_move
            if dist > distances[i]:
                winning_options = times[i] - s - (s-1)
                break
                
        options_per_race.append(winning_options)
            
    answer = reduce(lambda x, y: x*y, options_per_race)
    return answer

def p2(data):
    time = 0
    distance = 0
    for line in data:
        if line.startswith("Time"):
            time = int("".join([x for x in line.split(": ")[-1].split(" ") if x.isnumeric()]))
        elif line.startswith("Distance"):
            distance = int("".join([x for x in line.split(": ")[-1].split(" ") if x.isnumeric()]))

    
    winning_options = 0
    for s in range(1, time+1):
        speed = 1 * s
        time_to_move = time - s
        dist = speed * time_to_move
        if dist > distance:
            winning_options += 1
            answer = time - s - (s-1)
            break

    return answer


if __name__ == "__main__":
    with open("2023/d6/input") as t:
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