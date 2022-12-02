score_points = {"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
win_lose = {
        "AY": 6,
        "AZ": 0,
        "BX": 0,
        "BZ": 6,
        "CX": 6,
        "CY": 0
        }
# Not in dict = draw

def part_1(data):
    points = 0

    for round in data:
        s = round[0] + round[1]
        points += win_lose.get(s, 3)
        points += score_points[round[1]]

    print("P1 = " + str(points))


def part_2(data):
    to_play = {
            "AX": "Z",
            "AY": "X",
            "AZ": "Y",
            "BX": "X",
            "BY": "Y",
            "BZ": "Z",
            "CX": "Y",
            "CY": "Z",
            "CZ": "X"
            }
    points = 0

    for round in data:
        s = round[0] + round[1]
        pick = to_play[s]
        points += win_lose.get(round[0] + pick, 3)
        points += score_points[pick]

    print("P2 = " + str(points))


with open("input.txt") as t:
    data = [x.split(" ") for x in t.read().splitlines()]
    part_1(data)
    part_2(data)
