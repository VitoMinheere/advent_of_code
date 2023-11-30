import time

def p1(data):
    print(data)

def p2(data):
    print(data)


if __name__ == "__main__":
    with open("../test") as t:
        data = t.read().splitlines()

        start_p1 = time.time()
        answer_1 = p1(data)
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        start_p2 = time.time()
        answer_2 = p2(data)
        p2_time = time.time() - start_p2
        print("P2 took " + str(round(p2_time * 1000)) + " ms")