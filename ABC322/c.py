import bisect

n,m=map(int,input().split())
a = list(map(int, input().split()))
result=0

for i in range (1,n+1):
  result = bisect.bisect_left(a, i)
  print(a[result]-i)


