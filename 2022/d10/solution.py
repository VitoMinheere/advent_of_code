import time

def p1(data):
    cycles = 0
    register = 0

    check_cycles = [20, 60, 100, 140, 180, 220]
    check_values = []

    def update_cycles(amount, cycles, register):
        for _ in range(0, amount):
            cycles += 1
            if cycles in check_cycles:
                print(f"Cycle {str(cycles)} multiplying by {str(register)}")
                check_values.append(cycles * (register + 1)) # ugly hack
        return cycles
                

    for line in data:
        if line == "noop":
            cycles = update_cycles(1, cycles, register)
            continue
        else:
            val = int(line.split(" ")[-1])
            cycles = update_cycles(2, cycles, register)
            register += val

    p1 = sum(check_values)
    print("P1 = " + str(p1))
        

def p2(data):
    pass

with open("input.txt") as t:
    data = t.read().splitlines()
    start = time.time()
    p1(data)
    p1_time = time.time() - start
    # p2_start = time.time()
    # p2(data)
    # p2_time = time.time() - p2_start

    print("P1 took " + str(round(p1_time * 1000)) + " ms")
    # print("P2 took " + str(round(p2_time * 1000)) + " ms")