import time

def p1(data):
    list_1, list_2 = [], []
    for row in data:
        n1, n2 = row.split("   ")
        list_1.append(int(n1))
        list_2.append(int(n2))

    # Sort from small to large
    list_1.sort()
    list_2.sort()

    result = [abs(x - y) for x, y in zip(list_1, list_2)]
    return sum(result)

def p2(data):
    pass

if __name__ == "__main__":
    with open("2024/input.txt") as t:
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