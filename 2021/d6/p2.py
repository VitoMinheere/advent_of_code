# fish = [3,4,3,1,2]
fish = [1,3,1,5,5,1,1,1,5,1,1,1,3,1,1,4,3,1,1,2,2,4,2,1,3,3,2,4,4,4,1,3,1,1,4,3,1,5,5,1,1,3,4,2,1,5,3,4,5,5,2,5,5,1,5,5,2,1,5,1,1,2,1,1,1,4,4,1,3,3,1,5,4,4,3,4,3,3,1,1,3,4,1,5,5,2,5,2,2,4,1,2,5,2,1,2,5,4,1,1,1,1,1,4,1,1,3,1,5,2,5,1,3,1,5,3,3,2,2,1,5,1,1,1,2,1,1,2,1,1,2,1,5,3,5,2,5,2,2,2,1,1,1,5,5,2,2,1,1,3,4,1,1,3,1,3,5,1,4,1,4,1,3,1,4,1,1,1,1,2,1,4,5,4,5,5,2,1,3,1,4,2,5,1,1,3,5,2,1,2,2,5,1,2,2,4,5,2,1,1,1,1,2,2,3,1,5,5,5,3,2,4,2,4,1,5,3,1,4,4,2,4,2,2,4,4,4,4,1,3,4,3,2,1,3,5,3,1,5,5,4,1,5,1,2,4,2,5,4,1,3,3,1,4,1,3,3,3,1,3,1,1,1,1,4,1,2,3,1,3,3,5,2,3,1,1,1,5,5,4,1,2,3,1,3,1,1,4,1,3,2,2,1,1,1,3,4,3,1,3]

x = {i:0 for i in range(0,9)}

# setup dict
for f in fish:
    x[f] += 1

print(x)

for i in range(1, 257):
    # create new dict
    n = {i:0 for i in range(0,9)}
    print(f"Day {i}")
    # Every day go through the loop/check
    # Every fish with 0 days left will spawn an equal amount of new fish with 8 days left
    n[8] = x[0]
    # The pre existing fish will be added to the ones with 6 days left
    day6_fish = x[0]

    for f in reversed(range(1,8)):
        # For all fishes with 1 to 7 days left
        # Move current tier to 1 lower
        # print(f"Moving day {f} ({x[f]} to {f-1} ({n[f-1]}")
        n[f-1] = x[f]

    print(f"fish to add to day 6 {day6_fish}")
    n[6] += day6_fish
    x = n

res = sum(n.values())
print(res)
