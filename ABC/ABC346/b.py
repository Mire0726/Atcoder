s="wbwbwwbwbwbw"*20
S=list(s)

w,b=map(int,input().split())
for i in range(200):
  l=S[i:i+w+b]
  L=''.join(sorted(l))

  if L==("b"*b+"w"*w):
    print("Yes")
    exit()
print("No")
