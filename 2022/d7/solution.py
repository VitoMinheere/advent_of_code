def p1(data):
    MAX_SIZE = 100_000
    dirs = {}
    dir_contains = {}
    sum_dirs = {}
    dir_depth = ""
    cur_dir = ""
    cur_files = []
    output = False
    
    for i, line in enumerate(data):
        if line.startswith("$"): # It's a command
            if "cd" in line:
                output = False
                if cur_dir not in dirs:
                    # print("cur dir in cd = "+ cur_dir)
                    dirs[cur_dir] = cur_files
                    # print(dirs[cur_dir])

                if line[-1] == ".":
                    # print(dir_depth)
                    dir_depth = dir_depth[:-1]
                    # print(dir_depth)
                    cur_dir = dir_depth[-1]
                    # print("cur dir = " + cur_dir)
                else:
                    cur_dir = line[-1]
                    # print("cur dir = " + cur_dir)
                    dir_depth += line[-1]
            elif "ls" in line:
                # print("cur dir in ls = " + cur_dir)
                # print(cur_files)
                output = True
                # print(cur_files)
                # if cur_dir in dirs:
                #     dirs[cur_dir].extend([f for f in cur_files if f not in dirs[cur_dir]])
                # else:
                #     dirs[cur_dir] = cur_files

                cur_files = []
        if output:
            if line[0].isdigit():
                # print(line)
                cur_files.append(int(line.split(" ")[0])) # Add file size
            elif "dir" in line:
                if cur_dir in dir_contains:
                    dir_contains[cur_dir] += line[-1]
                else:
                    dir_contains[cur_dir] = line[-1]
        if i == len(data) -1 and cur_dir not in dirs: # Also add last dir
            # print("cur dir in cd = "+ cur_dir)
            dirs[cur_dir] = cur_files
            # print(dirs[cur_dir])

    # print(dirs)
    # print(dir_contains)
    for k, v in dir_contains.items():
        for x in v:
            print("DIR " + k)
            print(dirs[x])
            dirs[k].extend(dirs[x])

    dirs_in_size = []
    for k, v in dirs.items():
        s = sum(v)
        print(f"dir {k} size = {str(s)}")
        sum_dirs[k] = s
        if s <= MAX_SIZE:
            dirs_in_size.append(s)

    print("P1 = " + str(sum(dirs_in_size)))
            
def p2(data):
    pass

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    # p2(data)