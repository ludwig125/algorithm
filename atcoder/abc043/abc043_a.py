def sigma(n):
    if int(n) > 0:
        return int(n) + sigma(int(n)-1)
    return 0


#arr = list(map(int, input().split(" ")))
# $B$J$<$+$3$l$@$H(BEOFError: EOF when reading a line$B$,(Babcoder$B$N4D6-$G$G$k(B
import sys
N  = sys.stdin.readline().rstrip()

print(sigma(N))
