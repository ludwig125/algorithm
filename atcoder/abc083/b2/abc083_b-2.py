N, A, B = [int(x) for x in input().split()]

def sumEachOrder(n):
    s = 0
#    print(n)
    while n > 0:
        a = n % 10
#        print("  {}".format(a))
        s += a
        n //= 10
    return s

ans=0
for n in range(1, N+1):
    #print(sumEachOrder(n))
    s = sumEachOrder(n)
    if s >= A and s <= B:
        ans += n
print(ans)
