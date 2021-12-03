from collections import Counter

def main():
    bit_list = open("input.txt").read().splitlines()
    # bit_list = [
    #     "00100",
    #     "11110",
    #     "10110",
    #     "10111",
    #     "10101",
    #     "01111",
    #     "00111",
    #     "11100",
    #     "10000",
    #     "11001",
    #     "00010",
    #     "01010",
    # ]
    # find_oxygen_value(bit_list)
    find_co2_scrubber(bit_list)
    

def bits_equal(order):
    if order[0][1] == order[1][1]:
        return True
    return False

def find_oxygen_value(bit_list):
    while len(bit_list) > 1:
        for x in range(0, len(bit_list[0])):
            freq = Counter([i[x] for i in bit_list]).most_common()
            most_freq = freq[0][0]
            if bits_equal(freq):
                most_freq = "1"
            bit_list = [item for item in bit_list if item[x] == most_freq]
            print(len(bit_list))
    print(int(bit_list[0],2))

def find_co2_scrubber(bit_list):
    while len(bit_list) > 1:
        for x in range(0, len(bit_list[0])):
            test = [i[x] for i in bit_list]
            print(test)
            freq = Counter(test).most_common()
            print(freq)
            least_freq = freq[1][0]
            if bits_equal(freq):
                least_freq = "0"
            print("least_freq: %s", least_freq)
            bit_list = [item for item in bit_list if item[x] == least_freq]
            print(bit_list)
    print(int(bit_list[0],2))



main()