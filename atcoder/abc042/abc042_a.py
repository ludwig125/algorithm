#arr = list(map(int, input().split(" ")))
# $B$J$<$+$3$l$@$H(BEOFError: EOF when reading a line$B$,(Babcoder$B$N4D6-$G$G$k(B
import sys
input_line  = sys.stdin.readline()
#arr = list(map(int, input_line.rstrip().split(" ")))  
#$B$J$<$+$3$l$O(B  invalid literal for int()

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

