s=input().split()
cnt=0

for i in range (int(s[1]),int(s[0])+1,int(s[2])):
  if i:
    cnt=cnt+1
  
print(cnt)
