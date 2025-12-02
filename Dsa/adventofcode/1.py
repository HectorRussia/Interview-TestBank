
from pathlib import Path


p = Path("D:/TestFullstack/for_test_every/se-interviews/Dsa/adventofcode/my_input.txt")

with p.open("r",encoding="utf-8") as f:
    arr = [line.strip() for line in f if line.strip()]

# print (f"Loaded {len(arr)} item")

# print(arr)


pos = 50
count = 0
for move in arr:

    dir = move[0]
    steps = int(move[1:])

    if dir == "L":
        pos = (pos + steps) % 100
        count+1
    elif dir == "R":
        pos = (pos - steps) % 100
        count+1
    if pos == 0:
        count += 1

print(count)