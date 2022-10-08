def cost(N, K, X, Y):
    cost = 0
    for i in range(1,N+1):
        if i <= K:
            cost += X
        else:
            cost += Y
    return cost

N = int(input())
K = int(input())
X = int(input())
Y = int(input())

print(cost(N, K, X, Y))
