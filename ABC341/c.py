result=0

H,W,N=map(int,input().split())
T=input()
S = [input() for _ in range(H)]

for i in range(1,H-1):
    for j in range(1,W-1):
        if S[i][j] != ".":
            continue
        x=i
        y=j
        r=True
        for k in range(N):
            if T[k] =="L":
                x-=1
            elif T[k]=="R":
                x+=1
            elif T[k] == "U":
                y-=1
            elif T[k] == "D":
                y+=1
            if S[x][y]=="#":
                r=False
                break
        if r==True:
            result+=1
        
print(result)


        

