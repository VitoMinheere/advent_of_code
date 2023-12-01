import time
import re

def p1(data):
    answer = 0
    for line in data:
        nums = [x for x in line if x.isnumeric()]
        nums = [nums[0], nums[-1]]
        num = "".join(nums)
        answer += int(num)
    return answer

def p2(data):
    numbers = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9
        }
    
    answer = 0
    for line in data:
        num_list = []
        split = re.split("(\d+)", line)
        for part in split:
            if part in numbers:
                num_list.append(numbers[part])

            elif part.isnumeric():
                if len(part) > 1:
                    for x in part:
                       num_list.append(int(x)) 
                else:
                    num_list.append(int(part))

            if found := [x for x in numbers if part.find(x) >= 0]:
                index = 0
                while index <= len(part):
                    for i in range(5, 2, -1):
                        word = part[index:index+i] 
                        if not word:
                            break
                        if word in numbers:
                            num_list.append(numbers[word])
                            index += (i-2) 
                            continue
                    index+=1

        nums = [str(num_list[0]), str(num_list[-1])]
        num = "".join(nums)
        answer += int(num)

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