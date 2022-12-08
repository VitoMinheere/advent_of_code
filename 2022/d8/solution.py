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
            
            l_row = row[0:j]
            r_row = row[j+1:len(row)]

            col = [r[j] for r in matrix] # all col values
            u_col = col[0:i] # up
            d_col = col[i+1:len(col)] # down

            if all([x < cell for x in l_row]) \
            or all([x < cell for x in r_row]) \
            or all([x < cell for x in u_col]) \
            or all([x < cell for x in d_col]):
                visible[i][j] = 1 # invisible
            else:
                visible[i][j] = 0 # visible

    p1 = sum(map(sum, visible))
    print("P1 = " + str(p1))


def p2(data):
    matrix = [[] for _ in range(0, len(data))]
    for i, line in enumerate(data):
        matrix[i] = [int(x) for x in line]
    
    highest_score = 0

    def get_score(l, cell):
        score = 0
        for t in l:
            if t < cell:
                score += 1
            elif t >= cell:
                score += 1
                break
        if score == 0:
            score = 1
        return score

    for i, row in enumerate(matrix):
        for j, cell in enumerate(row):
            score = 1
            # Check if edge
            if i == 0 or i == len(matrix)-1 or j == 0 or j == len(row)-1:
                continue
            
            l_row = row[0:j]
            l_row.reverse()
            score *= get_score(l_row, cell)

            r_row = row[j+1:len(row)]
            score *= get_score(r_row, cell)

            col = [r[j] for r in matrix] # all col values
            u_col = col[0:i] # up
            u_col.reverse()
            score *= get_score(u_col, cell)

            d_col = col[i+1:len(col)] # down
            score *= get_score(d_col, cell)

            if score > highest_score:
                highest_score = score

                
    print("P2 = " + str(highest_score))
    

with open("input.txt") as t:
    data = t.read().splitlines()
    p1(data)
    p2(data)