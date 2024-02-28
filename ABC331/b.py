N,S,M,L=map(int,input().split)

ans = 10**9

for i in range(N + 1):
    for j in range(N+1):
        for k in range(N+1):
            if i*6 + j*8 +k*12>=N:
                ans=min(ans,i*S + j*M +k*L)

print(ans)
