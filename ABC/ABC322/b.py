n,m=map(int,input().split())
s=input()
t=input()

tf=t[0:len(s)]
tl=t[-len(s):]
c=0
if s==tf and s==tl:
  print(0)
elif s==tf and s!=tl:
  print(1)
elif s!=tf and s==tl:
  print(2)
else:
  print(3)

