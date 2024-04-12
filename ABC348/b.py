import math
# import itertools
n=int(input())

def point(a,b,c,d):
  r=(a-c)**2 + (b-d)**2
  return math.sqrt(r)
  

A=list()
B=list()
for i in range(n):
  A.append(list(map(int, input().split())))

for j in range(n):
  for k in range(n):
    B.append(point(A[j][0],A[j][1],A[k][0],A[k][1]))
  print(B.index(max(B)) + 1)
  B.clear()
