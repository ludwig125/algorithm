# https://atcoder.jp/contests/abc081/submissions/4789633
N = int(input())
A = [int(x) for x in input().split()]
#print(N)
#print(A)


n = 0
while True:
    a = list(map(lambda x: x % 2, A))
    if a.count(0) == N:
        b = list(map(lambda x: x / 2, A))
#        print(b)
        A = b
        n += 1
    else:
        break

print(n)
