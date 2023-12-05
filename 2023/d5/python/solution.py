import time

seeds = {}

def parse_values(line, lookup, name):
    _map = line.split("\n")[1:]
    for s in _map:
        dest, src, ran = [int(x) for x in s.split(" ")]

        for k, v in seeds.items():
                name_num = v[lookup] - src 
                if name_num > 0 and name_num <= ran:
                    new_val = dest + name_num
                    seeds[k][name] = new_val

    s = [k for k, v in seeds.items() if name not in v]
    for i in s:
        seeds[i][name] = seeds[i][lookup]

def update_seed(num):
    if num not in seeds:
        seeds[int(num)] = {
            "seed": int(num),
        }
    
def p1(data):
    blocks = data.split("\n\n")

    for line in blocks:
        if "seeds:" in line:
            seed_numbers = line.split(": ")[-1].split(" ")
            for num in seed_numbers:
                update_seed(int(num))

        elif "seed-to-soil" in line:
            parse_values(line, "seed", "soil")

        elif "soil-to-fertilizer" in line:
            parse_values(line, "soil", "fertilizer")

        elif "fertilizer-to-water" in line:
            parse_values(line, "fertilizer", "water")

        elif "water-to-light" in line:
            parse_values(line, "water", "light")

        elif "light-to-temperature" in line:
            parse_values(line, "light", "temperature")

        elif "temperature-to-humidity" in line:
            parse_values(line, "temperature", "humidity")

        elif "humidity-to-location" in line:
            parse_values(line, "humidity", "location")

    answer = sorted([v["location"] for k, v in seeds.items()])[0]
    return answer
            
            

def p2(data):
    pass

if __name__ == "__main__":
    # with open("../input") as t:
    with open("2023/d5/input") as t:
        data = t.read()

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