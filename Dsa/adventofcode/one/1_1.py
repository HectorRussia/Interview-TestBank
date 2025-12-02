
from math import floor
from pathlib import Path


p = Path("D:/TestFullstack/for_test_every/se-interviews/Dsa/adventofcode/one/my_input.txt")

with p.open("r",encoding="utf-8") as f:
    arr = [line.strip() for line in f if line.strip()]

pos = 50
count = 0   # นับแบบ Part 1 (จบแต่ละหมุน)
count_clicks  = 0  # นับแบบ Part 2 (ทุกคลิกที่โดน 0)
for move in arr:

    direction = move[0]
    steps = int(move[1:])
    start = pos

    if start == 0: # ไม่หมุนซ้ายขวา
        hits_this_move = steps // 100
    else:

        if direction == "R":
            distance = 100 - start
           
        elif direction == "L":
            distance = start

        if steps < distance:
            hits_this_move = 0
        else:
            hits_this_move = 1 + (steps - distance) // 100
    
    count_clicks += hits_this_move

    # PART 1

    if direction == "R":
        pos = (pos + steps) % 100
    else:  # "L"
        pos = (pos - steps) % 100

    if pos == 0:
        count += 1

print(count)
print(count_clicks )