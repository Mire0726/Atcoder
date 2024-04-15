n=int(input())
S=set()

for i in range (n):
  s=input()
  if s in S:
    continue
  else:
    S.add(s)
    print(i+1)
