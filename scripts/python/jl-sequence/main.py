def calc(n, a_n=0):
    if n == 0:
        return a_n
    return calc(n-1, a_n +37**n)

print(calc(3))