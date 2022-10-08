# https://atcoder.jp/contests/abc082/tasks/abc082_b
# do this. "python3 <>.py"

sl = list(input())
tl = list(input())

#print(''.join(sorted(sl)))
#print(''.join(sorted(tl, reverse=True)))

print('Yes') if ''.join(sorted(sl)) < ''.join(sorted(tl, reverse=True)) else print('No')
