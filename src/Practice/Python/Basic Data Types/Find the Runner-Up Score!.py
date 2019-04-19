input()  # Discard number of elements.
print(sorted({int(score) for score in input().split()})[-2])
