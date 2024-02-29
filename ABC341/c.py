Vec = {
    'D': (1, 0),
    'U': (-1, 0),
    'R': (0, 1),
    'L': (0, -1)
}

def isOK(i,j):
    if i+1

H,W,N=map(int,input().split())
T=input()
S = [input() for _ in range(H)]

for i in range(1,H-1):
    for j in range(1,W-1):
        if S[i][j] != ".":
            continue
        x=i
        y=j
        

