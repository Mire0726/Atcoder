n,k=map(int,input().split())
A=(x for x in set(map(int,input().split())) if x <=k)

print(((k+1)*k)//2-sum(A))


  