with open("input.txt") as t:
    data = t.read().splitlines()

    cals_per_elf = []
    temp_cal = []
    for i, cal in enumerate(data):
        if cal != '':
            temp_cal.append(int(cal))
            if i == (len(data) - 1):
                cals_per_elf.append(sum(temp_cal))
                temp_cal = []
        else:
            cals_per_elf.append(sum(temp_cal))
            temp_cal = []
    
    p1 = max(cals_per_elf)
    print("P1 = " + str(p1))

    # Part 2. Get top 3 
    top_3 = sorted(cals_per_elf)[-3:]
    p2 = sum(top_3)
    print("P2 = " + str(p2))