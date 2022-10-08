# https://qiita.com/drken/items/fd4e5e3630d0f5859067

def isEven(num):
    if num % 2 == 0:
        return 0
    else:
        return 1

def isAllEven(l):
    allEven = 0
    for i in l:
        allEven += isEven(i)
    
    if allEven == 0:
        return True
    else:
        return False

def countDivide(n, l):
    if isAllEven(l):
        newl = list(map(lambda x: int(x/2), l))
        #print(newl)
        n += 1
        return countDivide(n, newl)
    else:
        return n

N = int(input())
A = list(map(int, input().split()))

#print(N)
#print(A)

#print(list(map(isEven, A)))

print(countDivide(0, A))
