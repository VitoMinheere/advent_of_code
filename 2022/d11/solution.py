import time
from math import prod

class Monkey:

    def __init__(self, items: list) -> None:
        self.items = items
        self.items_inspected = 0
        self.arith = ""
        self.test = ""
        self.test_true = 0
        self.test_false = 0

    def append_items(self, items):
        for item in items:
            self.items.append(item)

    def set_arithmetic(self, opp: str):
        self.arith = opp

    def set_test(self, test):
        self.test = test

    def set_test_true(self, monkey):
        self.test_true = int(monkey)

    def set_test_false(self, monkey):
        self.test_false = int(monkey)

    def operation(self, worry_level: int):
        self.items_inspected += 1
        if "old" in self.arith:
            arith = self.arith.replace("old", str(worry_level))
        else:
            arith = self.arith
        return eval(f"{worry_level} {arith}")
        
    def get_next_monkey(self, worry_level):
        res = eval(f"{worry_level} {self.test} == 0")
        return (self.test_false, self.test_true)[res]



def p1(data, rounds):
    rounds = rounds
    worry_level = 0
    am_monkeys = int(data[-6].split(" ")[-1][:1])+1

    monkeys = [Monkey([]) for _ in range(0, am_monkeys)]

    start = 0
    end = 6
    for m in range(0, am_monkeys):
        for line in data[start:end]:
            if "Starting" in line:
                items = [int(x) for x in line.split(":")[-1].split(",")]
                monkeys[m].append_items(items)
            elif "Operation" in line:
                monkeys[m].set_arithmetic("".join(line.split("=")[-1][4:]))
            elif "Test" in line:
                monkeys[m].set_test(f"% {line.split(' ')[-1]}")
            elif "true" in line:
                monkeys[m].set_test_true(line.split(" ")[-1])
            elif "false" in line:
                monkeys[m].set_test_false(line.split(" ")[-1])

        start += 7
        end += 7

    mod = prod([int(x.test.split(" ")[-1]) for x in monkeys])

    for r in range(1, rounds+1):
        for i, m in enumerate(monkeys):
            # print("Monkey " + str(i))
            for _, item in enumerate(m.items):
                worry_level = item
                # print(f"Monkey inspects an item with a worry level of {item}")
                worry_level = m.operation(item)
                # print(f"Worry level = " + str(worry_level))
                if rounds == 20:
                    worry_level = worry_level // 3
                else:
                    worry_level = worry_level % mod
                # print(f"Monkey gets bored with item. Worry level is divided by 3 to {worry_level}.")
                next = m.get_next_monkey(worry_level)
                # print(f"Item with worry level {worry_level} is thrown to monkey {str(next)}.")
                monkeys[next].append_items([worry_level])
            m.items = []

    active = sorted(monkeys, key=lambda x: x.items_inspected, reverse=True)
    for i, x in enumerate(active):
        print(f"Monkey {str(i)} inspected items {str(x.items_inspected)} times.")

    p1 = 1
    for x in active[:2]:
        print(x.items_inspected)
        p1 *= x.items_inspected

    print("P1 = " + str(p1))


with open("input.txt") as t:
    data = t.read().splitlines()
    start = time.time()
    p1(data, 20)
    p1_time = time.time() - start
    p2_start = time.time()
    p1(data, 10000)
    p2_time = time.time() - p2_start

    print("P1 took " + str(round(p1_time * 1000)) + " ms")
    print("P2 took " + str(round(p2_time * 1000)) + " ms")