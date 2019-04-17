n = int(input())
s = set(map(int, input().split()))
N = int(input())
for _ in range(N):
    data = input().split(' ')
    if data[0] == 'pop':
        s.pop()
    elif data[0] == 'remove':
        s.remove(int(data[1]))
    elif data[0] == 'discard':
        s.discard(int(data[1]))
print(sum(s))
