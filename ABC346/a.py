n=int(input())
a=list(map(int, input().split()))
b=[]
for i in range(1,n):
  b.append(a[i-1]*a[i])

print(*b)