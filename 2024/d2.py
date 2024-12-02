import time

def p1(data):
    result = 0
    for row in data:
        nums = [int(x) for x in row.split(" ")]
        #line = iter(nums)
        if nums[0] < nums[1]:
            increasing = True
        else:
            increasing = False

        correct = True
        for i, n in enumerate(nums):
            try:
                nxt = nums[i+1]
            except IndexError: 
                # at last number
                break

            if increasing and nxt <= n:
                correct = False
                break
            if not increasing and nxt >= n:
                correct = False
                break

            diff = abs(n - nxt)
            if diff < 1 or diff > 3:
                correct = False
                break

        if correct:
            result +=1

    return result


def p2(data):
    pass

if __name__ == "__main__":
    with open("test.txt") as t:
        data = t.read().splitlines()

        start_p1 = time.time()
        answer_1 = p1(data)
        print(f"P1 answer = {answer_1}")
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
        # x ms
    
        start_p2 = time.time()
        answer_2 = p2(data)
        print(f"P2 answer = {answer_2}")
        p2_time = time.time() - start_p2
        print("P2 took " + str(round(p2_time * 1000)) + " ms")
        # x ms
