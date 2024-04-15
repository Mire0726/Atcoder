h,w=map(int,input().split())
A = [list(map(int, input().split())) for _ in range (h)]
c_sum=[0] * w
r_sum=[0] * h

for i in range(h):
  for j in range(w):
    c_sum[j] += A[i][j]
    r_sum[i] += A[i][j]

for i in range(h):
  ans=[]
  for j in range(w):
    cr_sum=c_sum[j]+r_sum[i]-A[i][j]
    ans.append(cr_sum)
  print(*ans)