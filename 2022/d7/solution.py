def p1(data):
    MAX_SIZE = 100_000
    dirs = {"/": 0}
    dir_contains = {}
    sum_dirs = {}
    dir_depth = []
    cur_dir = ""
    cur_files = []
    output = False
    
    for i, line in enumerate(data):
        if line.startswith("$"): # It's a command
            if "cd" in line:
                output = False
                if cur_dir not in dirs:
                    dirs[cur_dir] = 0 #cur_files

                if line[-1] == ".":
                    cur_dir = dir_depth.pop()
                else:
                    cur_dir = line.split(" ")[-1]
                    dir_depth.append(cur_dir)
            elif "ls" in line:
                output = True
                continue

        if output:
            if line[0].isdigit():
                size = int(line.split(" ")[0])
                # print(line)
                # print(dirs)
                # print(size)
                # print(dir_depth)
                for dir in dir_depth:
                    if dir in dirs:
                        # print(dir)
                        dirs[dir] += size
                    else:
                        dirs[dir] = size
                        
                # print(line)
                # cur_files.append(int(line.split(" ")[0])) # Add file size
            # elif "dir" in line:
            #     for dir in dir_depth:
            #         if dir in dir_contains:
            #             dir_contains[dir].append(line.split(" ")[-1])
            #         else:
            #             dir_contains[dir] = [line.split(" ")[-1]]

        # if (i == len(data)-1) and cur_dir not in dirs: # Also add last dir
            # print("cur dir in cd = "+ cur_dir)
            # dirs[cur_dir] = cur_files
            # print(dirs[cur_dir])

    # print(dirs)
    # print(dir_contains)
    # for k, v in dir_contains.items():
    #     for x in v:
    #         # print("DIR " + k)
    #         # print(dirs[x])
    #         dirs[k].extend(dirs[x])

    # dirs_in_size = []
    # for k, v in dirs.items():
    #     s = sum(v)
    #     print(f"dir {k} size = {str(s)}")
    #     sum_dirs[k] = s
    #     if s <= MAX_SIZE:
    #         dirs_in_size.append(s)
    print(dirs)
    p1 = sum([v for k, v in dirs.items() if v <= MAX_SIZE])

    # print("P1 = " + str(sum(dirs_in_size)))
    print("P1 = " + str(p1))
            
def p2(data):
    pass

with open("test.txt") as t:
    data = t.read().splitlines()
    p1(data)
    # p2(data)