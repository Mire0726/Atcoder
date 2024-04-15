import math
a,b,c=map(int,input().split())

gcd=math.gcd(a,b,c)

count = (a//gcd + b//gcd +c//gcd) -3

print(count)