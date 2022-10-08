N, A, B = [int(x) for x in input().split()]
#print(N)
#print(A)
#print(B)

ans = 0
for n in range(1, N+1):
    nl = list(map(int, list(str(n))))
    s = 0
    for i in nl:
        s += i
    #print(nl, s, n)
    if s >= A and s <= B:
        ans += n

print(ans)
