#arr = list(map(int, input().split(" ")))
# なぜかこれだとEOFError: EOF when reading a lineがabcoderの環境ででる
import sys
input_line  = sys.stdin.readline()
#arr = list(map(int, input_line.rstrip().split(" ")))  
#なぜかこれは  invalid literal for int()

arr = input_line.rstrip().split(" ")

count5 = 0
count7 = 0

for i in arr:
    if int(i) == 5:
        count5 += 1
    if int(i) == 7:
        count7 += 1

if count5 == 2 and count7 == 1:
    print("YES")
else:
    print("NO")

