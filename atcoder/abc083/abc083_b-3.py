N, A, B = [int(x) for x in input().split()]

def sumEachOrder(n):
    s = 0
    while n > 0:
        s += n % 10
        n //= 10
    return s

ans=0
print(sum([n for n in range(1, N+1) if sumEachOrder(n) >= A and sumEachOrder(n) <= B]))
#for n in range(1, N+1):
#    s = sumEachOrder(n)
#    if s >= A and s <= B:
#        ans += n
#print(ans)
