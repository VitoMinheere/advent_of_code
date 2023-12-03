import time
import re
from functools import reduce

CORRECT = [733, 520, 161, 462, 450, 183, 707, 352, 333, 484, 635, 287, 42, 131, 913, 634, 440, 404, 272, 547, 344, 689, 589, 150, 382, 168, 433, 253, 102, 78, 69, 37, 510, 797, 596, 946, 602, 175, 100, 681, 110, 396, 858, 381, 246, 637, 391, 973, 274, 551, 576, 21, 176, 883, 223, 649, 701, 936, 17, 482, 80, 210, 563, 222, 373, 532, 707, 956, 927, 698, 763, 275, 27, 498, 719, 540, 24, 39, 514, 126, 974, 952, 308, 965, 450, 4, 343, 772, 437, 309, 58, 840, 34, 189, 186, 657, 670, 22, 406, 880, 939, 58, 387, 587, 564, 599, 831, 891, 552, 534, 802, 975, 960, 346, 792, 828, 807, 6, 326, 324, 722, 334, 895, 764, 611, 79, 461, 546, 940, 247, 473, 288, 932, 590, 918, 791, 621, 129, 10, 166, 11, 831, 302, 832, 159, 506, 229, 39, 587, 523, 610, 805, 559, 229, 961, 139, 411, 554, 26, 90, 463, 615, 422, 109, 247, 231, 56, 279, 159, 126, 569, 906, 20, 990, 960, 364, 812, 812, 176, 15, 346, 908, 828, 650, 842, 771, 189, 9, 42, 897, 479, 99, 842, 356, 849, 336, 201, 412, 205, 956, 319, 383, 438, 62, 348, 54, 302, 298, 318, 264, 28, 259, 824, 306, 249, 392, 444, 889, 418, 496, 610, 160, 699, 987, 125, 292, 514, 165, 672, 551, 194, 53, 500, 599, 275, 192, 854, 190, 40, 662, 706, 774, 840, 638, 82, 520, 190, 719, 34, 620, 532, 521, 423, 75, 258, 605, 354, 453, 248, 561, 481, 727, 437, 969, 243, 183, 250, 312, 150, 624, 454, 710, 30, 308, 381, 941, 461, 341, 19, 297, 213, 522, 442, 962, 324, 775, 290, 301, 15, 780, 562, 438, 806, 396, 508, 700, 16, 230, 443, 617, 447, 336, 424, 618, 276, 260, 452, 824, 916, 145, 55, 620, 773, 116, 384, 972, 974, 872, 494, 768, 195, 108, 440, 187, 135, 860, 923, 966, 136, 423, 185, 761, 217, 157, 292, 268, 789, 375, 919, 425, 542, 691, 69, 184, 578, 740, 619, 788, 324, 531, 82, 962, 409, 981, 652, 32, 53, 761, 353, 43, 6, 589, 190, 362, 961, 917, 165, 718, 633, 781, 565, 324, 339, 964, 427, 260, 919, 828, 469, 585, 803, 365, 406, 190, 2, 803, 417, 955, 75, 149, 293, 768, 678, 675, 449, 347, 301, 490, 751, 888, 348, 747, 318, 888, 160, 335, 455, 761, 564, 926, 907, 857, 639, 285, 567, 862, 731, 9, 513, 694, 233, 721, 806, 101, 729, 292, 565, 329, 137, 849, 773, 32, 359, 277, 849, 399, 274, 535, 348, 950, 553, 607, 564, 818, 438, 343, 85, 162, 67, 374, 593, 740, 553, 795, 664, 493, 971, 510, 843, 791, 277, 958, 489, 657, 213, 610, 193, 12, 531, 659, 521, 157, 896, 503, 241, 257, 171, 245, 959, 109, 454, 239, 203, 359, 330, 846, 751, 382, 985, 646, 624, 870, 890, 662, 371, 311, 739, 817, 250, 54, 182, 237, 261, 250, 946, 943, 414, 80, 159, 448, 91, 529, 683, 370, 569, 131, 34, 736, 434, 929, 127, 110, 644, 467, 92, 673, 718, 697, 745, 503, 288, 827, 911, 79, 191, 230, 517, 551, 192, 917, 597, 479, 316, 743, 122, 529, 461, 793, 633, 604, 320, 393, 599, 717, 567, 144, 545, 6, 591, 216, 73, 215, 842, 783, 855, 909, 537, 435, 874, 902, 657, 449, 510, 28, 605, 468, 418, 701, 748, 487, 521, 934, 992, 160, 477, 916, 476, 651, 462, 464, 425, 531, 735, 853, 965, 754, 657, 92, 838, 116, 469, 498, 537, 666, 622, 237, 204, 242, 599, 283, 283, 919, 638, 794, 204, 326, 541, 544, 183, 67, 903, 75, 605, 860, 455, 730, 143, 366, 59, 447, 916, 127, 60, 572, 658, 891, 300, 674, 733, 768, 522, 870, 89, 764, 606, 798, 298, 342, 430, 668, 485, 270, 265, 893, 524, 880, 88, 907, 239, 509, 406, 264, 829, 991, 850, 913, 302, 228, 601, 63, 156, 310, 184, 260, 748, 149, 895, 697, 66, 483, 500, 970, 88, 941, 25, 623, 436, 278, 522, 885, 848, 260, 579, 950, 948, 736, 245, 936, 512, 790, 634, 764, 488, 46, 995, 344, 59, 135, 851, 949, 214, 610, 185, 383, 218, 355, 804, 687, 751, 476, 145, 991, 314, 123, 917, 434, 634, 805, 15, 884, 526, 20, 328, 672, 245, 392, 423, 83, 484, 581, 440, 759, 124, 397, 779, 752, 303, 764, 376, 806, 976, 75, 765, 430, 796, 82, 474, 351, 287, 731, 960, 381, 192, 808, 959, 322, 243, 534, 134, 828, 95, 241, 388, 511, 825, 742, 888, 53, 191, 933, 454, 649, 993, 430, 624, 817, 157, 845, 466, 692, 987, 381, 746, 744, 722, 340, 526, 74, 810, 257, 941, 898, 769, 480, 658, 338, 107, 563, 550, 388, 542, 112, 572, 553, 588, 641, 759, 880, 740, 247, 310, 3, 154, 790, 24, 847, 767, 933, 790, 669, 626, 862, 77, 627, 809, 922, 913, 243, 372, 878, 155, 36, 656, 917, 935, 184, 888, 357, 917, 427, 665, 766, 986, 671, 247, 641, 598, 459, 922, 295, 664, 188, 35, 806, 107, 32, 261, 113, 456, 916, 480, 942, 848, 781, 3, 146, 168, 222, 137, 783, 940, 25, 902, 217, 460, 140, 170, 283, 647, 501, 359, 365, 491, 869, 103, 941, 75, 282, 786, 741, 686, 24, 527, 109, 79, 575, 810, 373, 444, 37, 331, 714, 62, 983, 446, 79, 886, 40, 506, 421, 726, 171, 676, 964, 651, 500, 497, 601, 263, 675, 136, 56, 836, 264, 156, 495, 769, 597, 355, 365, 379, 539, 933, 730, 718, 12, 63, 368, 691, 772, 116, 501, 777, 673, 467, 688, 526, 387, 726, 568, 260, 456, 403, 311, 387, 901, 744, 929, 825, 867, 994, 346, 260, 892, 294, 172, 730, 743, 610, 954, 44, 46, 606, 718, 970, 863, 772, 284, 950, 425, 625, 479, 305, 331, 139, 976, 655, 233, 373, 870, 17, 399, 820, 844, 744, 353, 135, 700, 364, 266, 295, 486, 932, 98, 311, 319, 599, 98, 202, 768, 982, 413, 185, 48, 132, 239, 551, 38, 111, 778, 856, 25, 355, 230, 879, 264, 783, 839, 682, 756, 415, 589, 438, 428, 465, 194, 803, 100, 955, 238, 836, 767, 555, 86, 702, 363, 630, 737, 256, 57, 806, 591, 348, 460, 244, 6, 789, 687] 

def p1(data):
    m = []
    parts_list = []

    for line in data:
        m.append(line)

    directions = [
        (0, 1),     # left
        (1, 1),   # top left
        (1, 0),    # top
        (1, -1),    # top right
        (0, -1),     # right 
        (-1, -1),     # bottom right
        (-1, 0),     # bottom
        (-1, 1),    # bottom left
    ]

    for r_index, row in enumerate(m):
        # print(row)
        for c_index, col in enumerate(row):
            if not col.isnumeric() and col != ".":
                pos = (r_index, c_index)
                parts_found = []
                row_used = []
                for d in directions:
                    # if d[0] in row_used:
                    #     continue
                    new_pos = tuple(map(lambda i, j: i - j, pos, d))
                    val = m[new_pos[0]][new_pos[1]]
                    # print(val)
                    if val.isnumeric():
                        # print(m[new_pos[0]])
                        part_num = []
                        # Found value, check sides to get the whole number
                        part_num.append(val)
                        left = m[new_pos[0]][new_pos[1]-1]
                        right = m[new_pos[0]][new_pos[1]+1]
                        if not left.isnumeric() and right.isnumeric():
                            try:
                                print("right")
                                # Then go right
                                i = 1
                                while val.isnumeric() and val != ".":
                                    val = m[new_pos[0]][new_pos[1]+i]
                                    print(val)
                                    part_num.append(val)
                                    i += 1
                                    val = m[new_pos[0]][new_pos[1]+i]
                            except IndexError:
                                part_number = int("".join(part_num))
                                if part_number not in parts_found:
                                    row_used.append(d[0])
                                    parts_found.append(part_number)
                                    parts_list.append(part_number)
                                continue

                        elif left.isnumeric() and not right.isnumeric():
                            try:
                                print("left")
                                i = 1
                                val = m[new_pos[0]][new_pos[1]-i]
                                while val.isnumeric() and val != ".":
                                    print(val)
                                    part_num.append(val)
                                    i += 1
                                    val = m[new_pos[0]][new_pos[1]-i]
                            except IndexError:
                                continue
                            finally:
                                part_num.reverse()

                        elif left.isnumeric() and right.isnumeric():
                            print("middle")
                            # You are in the middle
                            part_num.insert(0, left)
                            part_num.append(right)

                        part_number = int("".join(part_num))
                        if part_number not in parts_found:
                            row_used.append(d[0])
                            parts_found.append(part_number)
                            parts_list.append(part_number)
                        # break
                            
    print(parts_list)
    answer = sum(parts_list)
    # anser = [467, 35, 633, 617, 592, 755, 664, 598]
    return answer

def p2(data):
    pass

if __name__ == "__main__":
    # with open("../test") as t:
    with open("2023/d3/input") as t:
        data = t.read().splitlines()

        start_p1 = time.time()
        answer_1 = p1(data)
        print(f"P1 answer = {answer_1}")
        p1_time = time.time() - start_p1
        print("P1 took " + str(round(p1_time * 1000)) + " ms")
    
        # start_p2 = time.time()
        # answer_2 = p2(data)
        # print(f"P2 answer = {answer_2}")
        # p2_time = time.time() - start_p2
        # print("P2 took " + str(round(p2_time * 1000)) + " ms")