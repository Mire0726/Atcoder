

from string import ascii_lowercase

N = int(input())
S = input()
Q = int(input())
mapping_from = ascii_lowercase
mapping_to = ascii_lowercase

for _ in range(Q):
  C,D = input().split()  # splitメソッドを正しく呼び出す
  mapping_to=mapping_to.replace(C,D)

print(S.translate(str.maketrans(mapping_from, mapping_to)))
