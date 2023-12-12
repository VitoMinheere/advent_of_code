import time
from functools import reduce
from itertools import product, permutations

def foo(l):
     yield from product(*([l] * len(l)))

def find(s, ch):
    return [i for i, ltr in enumerate(s) if ltr == ch]

def p1(data):
    answer = 0
    for i, line in enumerate(data):
        if i == 1:
            springs, records = line.split(" ")
            print(springs)
            records = [int(x) for x in records if x.isnumeric()]
            total_working = sum(records)

            # prods = list(product([False, True], repeat=7))
            # for p in prods:
            #     print(p)

            option_string = "#"*total_working + "."*(len(springs)-total_working)
            print(option_string)
            combs = set(permutations(option_string, len(option_string)))
            dot_indexes = find(springs, ".")
            print(dot_indexes)
            answer = 0
            for c in combs:
                c_dot_indexes = find(c, ".")
                c_str = "".join(c)
                if set(dot_indexes).issubset(c_dot_indexes):
                    w_springs = c_str.split(".")
                    if all([len(x) == records[i] for i, x in enumerate(w_springs)]):
                        print(c)
                        answer += 1
    return answer
def p2(data):
    pass


if __name__ == "__main__":
    with open("../test") as t:
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
