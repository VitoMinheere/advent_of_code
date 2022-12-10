import time

class Device:
    matrix = [[" " for x in range(0, 40)] for _ in range(0, 7)]
    matrix[0][0] = "#"
    val = 0
    cycles = 0
    register = 1
    check_cycles = [20, 60, 100, 140, 180, 220]
    check_values = []
    sprite = [0, 1, 2]

    crt = {"x": 0, "y": 0}

    def update_sprite(self):
        self.sprite = [self.register-1, self.register, self.register+1]

    def update_crt(self):
        if self.crt["x"] in self.sprite:
            self.matrix[self.crt["y"]][self.crt["x"]] = "#"
        else:
            self.matrix[self.crt["y"]][self.crt["x"]] = " "

        if self.crt["x"] == 39:
            self.crt["x"] = 0
            self.crt["y"] += 1
        else:
            self.crt["x"] += 1
        
            

    def update_cycles(self, amount):
        for _ in range(0, amount):
            self.cycles += 1
            self.update_crt()
            if self.cycles in self.check_cycles:
                self.check_values.append(self.cycles * (self.register)) # + 1)) # ugly hack
            if _ == 1:
                self.register += self.val
                self.update_sprite()

    def print_matrix(self):
        print('\n'.join([''.join([item for item in row]) for row in self.matrix]))

    def p1(self, data):
        for line in data:
            if line == "noop":
                self.update_cycles(1)
            else:
                self.val = int(line.split(" ")[-1])
                self.update_cycles(2)

        p1 = sum(self.check_values)
        print("P1 = " + str(p1))
        print("P2")
        self.print_matrix()
            

with open("input.txt") as t:
    data = t.read().splitlines()
    start = time.time()
    device = Device()
    device.p1(data)
    p1_time = time.time() - start

    print("P1 took " + str(round(p1_time * 1000)) + " ms")