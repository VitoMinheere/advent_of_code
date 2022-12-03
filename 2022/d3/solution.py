def p1(data):
    shared_items = []

    for bag in data:
        first = bag[slice(0, len(bag) // 2)]
        second = bag[slice(len(bag)// 2, len(bag))]
        shared_item = [x for x in first if x in second][0]
        shared_items.append(shared_item)

    sum_lower = sum([ord(char) - 96 for char in shared_items if char.islower()])
    sum_upper = sum([(ord(char) - 64) + 26 for char in shared_items if char.isupper()])

    print("P1 answer is " + str(sum_lower + sum_upper))

def p2(data):
    grouped = [data[x:x+3] for x in range(0, len(data), 3)]
    score = 0
    for bags in grouped:
        shared_item = [x for x in bags[0] if x in bags[1] and x in bags[2]][0]

        if shared_item.islower():
            score += ord(shared_item) - 96
        else:
            score += ord(shared_item) - 64 + 26
            
    print("P2 answer is " + str(score))

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)