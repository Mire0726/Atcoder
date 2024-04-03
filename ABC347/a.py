n,k=map(int,input().split())
a=list(map(int,input().split()))
r=[]
for i in a:
  if i%k==0:
    r.append(int(i/k))

print(*r)