def sigma(n):
    if int(n) > 0:
        return int(n) + sigma(int(n)-1)
    return 0


#arr = list(map(int, input().split(" ")))
# なぜかこれだとEOFError: EOF when reading a lineがabcoderの環境ででる
import sys
N  = sys.stdin.readline().rstrip()

print(sigma(N))
