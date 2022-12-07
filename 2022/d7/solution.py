def p1(data):
    MAX_SIZE = 100_000
    dirs = {}
    dir_depth = []
    cur_dir = ""
    file_cnt = 0
    
    for i, line in enumerate(data):
        if line.startswith("$"): # It's a command
            if "cd" in line:
                if line[-1] == ".":
                    cur_dir = dir_depth.pop()
                else:
                    cur_dir = line.split(" ")[-1] + str(i) # Make each dir unique
                    if cur_dir not in dirs:
                        dirs[cur_dir] = 0
                    dir_depth.append(cur_dir)

        elif line[0].isdigit():
            file_cnt += 1
            size = int(line.split(" ")[0])
            for dir in dir_depth:
                dirs[dir] += size
        else:
            continue
                        
    p1 = sum([v for k, v in dirs.items() if v <= MAX_SIZE])
    print("P1 = " + str(p1))
    return dirs
            
def p2(dirs):
    DISK_SPACE = 70_000_000
    USED = dirs["/0"]
    UNUSED = DISK_SPACE - USED
    NEEDED = 30_000_000
    
    options = []
    for k, v in dirs.items():
        if (v + UNUSED) > NEEDED:
            options.append(v)
    print(f"P2 = {sorted(options)[0]}") 


with open("input.txt") as t:
    data = t.read().splitlines()
    dirs = p1(data)
    p2(dirs)