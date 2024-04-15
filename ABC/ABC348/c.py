from collections import defaultdict
n=int(input())
dic=defaultdict(int)

for i in range(n):
  a,c= map(int,input().split())
  dica=dic.get(c,a)
  dic[c]= min(dica,a)

print(max(dic.values()))



  