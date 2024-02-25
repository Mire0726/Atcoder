s=input()

for i in range(len(s)-1):
  if s[i] != s[i+1] and i==len(s)-2:
    print(len(s))
    exit()
  if s[i] != s[i+1] and s[i] != s[i+2]:
    print (i+1)
    exit()
  if s[i] != s[i+1] and s[i] == s[i+2]:
    print(i+2)
    exit()
    