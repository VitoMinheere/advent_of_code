import copy

def p1(data):
    matrix = [[] for _ in range(0, len(data))]
    for i, line in enumerate(data):
        matrix[i] = [int(x) for x in line]
    
    visible = copy.deepcopy(matrix)

    for i, row in enumerate(matrix):
        for j, cell in enumerate(row):
            # Check if edge
            if i == 0 or i == len(matrix)-1 or j == 0 or j == len(row)-1:
                visible[i][j] = 1
                continue
            
            # print("i = " + str(i) + " j = " + str(j))
            # all left and right without edge
            l_row = row[0:j]
            # print("full row")
            # print(row)
            # print("left of tree")
            # print(l_row)
            r_row = row[j+1:len(row)]
            # print("right of tree")
            # print(r_row)

            col = [r[j] for r in matrix] # all col values
            # print("Full column")
            # print(col)
            u_col = col[0:i] # up
            # print("u col")
            # print(u_col)
            d_col = col[i+1:len(col)] # down
            # print("d col")
            # print(d_col)

            # if all([cell <= x for x in l_row]) \
            # and all([cell <= x for x in r_row]) \
            # and all([cell <= x for x in u_col]) \
            # and all([cell <= x for x in d_col]):
            if all([x < cell for x in l_row]) \
            or all([x < cell for x in r_row]) \
            or all([x < cell for x in u_col]) \
            or all([x < cell for x in d_col]):
                visible[i][j] = 1 # invisible
                # print("visible")
            else:
                visible[i][j] = 0 # visible
                # print("inVisible")
                
                
            # check row has lower trees
            # for t in range(0, j):
            #     if matrix[i][t] < cell:
            #         visible[i][j] = 1 # visible
            #     else:
            #         visible[i][j] = 0 # not visible
                    
            # # check col has lower trees
            # for t in range(0, i):
            #     if matrix[t][j] < cell:
            #         visible[i][j] = 1 # visible
            #     else:
            #         visible[i][j] = 0 # not visible
            
            # elif matrix[i][j-1] < cell \
            # and matrix[i][j+1] < cell \
            # and matrix[i-1][j] < cell \
            # and matrix[i+1][j] < cell:
            #     visible[i][j] = 1
            # else:
            #     visible[i][j] = 0
                
    # print(matrix)
    # print(visible)
    # answer = [[1,1,1,1,1], [1,1,1,0,1], [1,1,0,1,1], [1,0,1,0,1], [1,1,1,1,1]]
    # print(answer)
    p1 = sum(map(sum, visible))
    print("P1 = " + str(p1))
def p2(data):
    pass
    

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)