num = int(input("enter a number:"))

x = num & 1

if x == 1:
    print(f"number {num} is odd: {x}")
else:
    print(f"number {num} is even: {x}")
