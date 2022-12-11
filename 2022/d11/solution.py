import time

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
        return eval(f"{worry_level} {self.arith}")
        

def p1(data):
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
                print(line.split("=")[-1][4:])
                monkeys[m].set_arithmetic("".join(line.split("=")[-1][4:]).replace("old", str(worry_level)))
            elif "Test" in line:
                monkeys[m].set_test(f"/ {line.split(' ')[-1]}")
            elif "true" in line:
                monkeys[m].set_test_true(line.split(" ")[-1])
            elif "false" in line:
                monkeys[m].set_test_false(line.split(" ")[-1])

        start += 7
        end += 7

    for monkey in monkeys:
        print(monkey.__dict__)

def p2(data):
    pass

with open("test.txt") as t:
    data = t.read().splitlines()
    start = time.time()
    p1(data)
    p1_time = time.time() - start
    p2_start = time.time()
    p2(data)
    p2_time = time.time() - p2_start

    print("P1 took " + str(round(p1_time * 1000)) + " ms")
    print("P2 took " + str(round(p2_time * 1000)) + " ms")