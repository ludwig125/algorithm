a, b = [int(x) for x in input().split()]

mul = a * b

if mul % 2 == 0:
    print("Even")
else:
    print("Odd")
