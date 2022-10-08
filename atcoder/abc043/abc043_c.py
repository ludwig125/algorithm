# https://beta.atcoder.jp/contests/abc043/submissions/2436112
N = int(input())
#print(N)


A = list(map(int, input().split()))
A.sort()
#print(A)
#print(A[0])
#print(A[len(A)-1])

min_cost = 10000000000
for i in range(A[0], A[len(A)-1]+1):
    cost = sum([(x - i)**2 for x in A])
    min_cost = min(cost, min_cost)

print(min_cost)
