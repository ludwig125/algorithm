# https://beta.atcoder.jp/contests/abc044/submissions/2441947
w = input()
for i in w:
    if w.count(i)%2!=0:
        print('No')
        break
else:
    print('Yes')
