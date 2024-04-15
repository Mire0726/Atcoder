n= int(input())

c1=[0]*(n+1)
c2=[0]*(n+1)

for i in range(n):
  c1[i+1]= c1[i]
  c2[i+1]= c2[i]
  #これで各項におけるそれまでの累積が出せるようになる！
  c,p=map(int,input().split())
  if c==1:
    c1[i+1]+= p
  else:
    c2[i+1]+= p
q=int(input())

for _ in range(q):
  l,r=map(int,input().split())
  l=l-1
  print(c1[r]-c1[l],c2[r]-c2[l])


