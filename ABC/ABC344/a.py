s=input()
s1=s.find("|")
s2=s.rfind("|")

print(s[:s1]+s[s2+1:])
