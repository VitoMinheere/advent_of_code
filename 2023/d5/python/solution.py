import time

seeds = []

def parse_values(dest, src, ran, val):
    new_val = val
    num = val - src 
    if num >= 0 and num <= ran:
        new_val = dest + num

    return new_val

def solution(data, p2=False):
    blocks = data.split("\n\n")
    
    if not p2: 
        seed_numbers = blocks[0].split(": ")[-1].split(" ")
        for num in seed_numbers:
            seeds.append(int(num))

    calculations = {} 
    for line in blocks[1:]:
        name = line.split("\n")[0]
        calculations[name] = []
        parts = line.split("\n")[1:] 
        for p in parts:
            calculations[name].append([int(x) for x in p.split(" ")])

    for i, s in enumerate(seeds):
        changes = [s]
        val = s
        for k, v in calculations.items():
            for c in v:
                new_val = parse_values(c[0], c[1], c[2], val)
                if new_val != val:
                    val = new_val
                    changes.append(new_val)
                    break
        seeds[i] = val
        
    answer = sorted(seeds)[0]
    return answer
            
            

def p2(data):
    pass

if __name__ == "__main__":
    # with open("../input") as t:
    with open("2023/d5/input") as t:
        data = t.read()

        start_p1 = time.time()
        answer_1 = solution(data)
        print(f"P1 answer = {answer_1}")
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        # start_p2 = time.time()
        # answer_2 = p2(data)
        # print(f"P2 answer = {answer_2}")
        # p2_time = time.time() - start_p2
        # print("P2 took " + str(round(p2_time * 1000)) + " ms")