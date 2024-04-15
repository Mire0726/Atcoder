n=int(input())
a = list(map(int, input().split()))
r=0
m=[]

for i in a:
  if i>=0:
    r+=i
  else:
    r+=i
    if r<0:
      m.append(abs(r))

n = max(m) if m else 0

for j in a:
  n+=j
  
print(n)