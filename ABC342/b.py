N=int(input())
A=input().split()
Q=int(input())

for i in range(Q):
  f,s=input().split()
  a=A.index(f)
  b=A.index(s)
  if a<b:
    print(f)
  else:
    print(s)
  