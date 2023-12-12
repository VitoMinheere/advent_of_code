with open("../input", "r") as t:
    data = t.read()
    lines = data.split("\n")
    instructions = lines[0]
    elements = {}
    e = lines[2:]
    e = e[:-1]
    for v in e:
        key = v.split(" = (")
        k = key[1].split(", ")
        elements[key[0]] = {"L": k[0].strip(), "R": k[1][:-1]}

    answer = 1
    cur_el = elements["AAA"]
    instructions = instructions * 30000
    print(instructions)
    for c in instructions:
        if cur_el[c] == "ZZZ":
            break
        else:
            next_el = elements[cur_el[c]]
            answer += 1
            cur_el = next_el
    print(answer)




