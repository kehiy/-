k = 7 # 111

third_bit = (k >> 2) & 1
assert third_bit == 1
print(third_bit)

k2 = 707 # 1011000011

second_bit = (k2 >> 3) & 1
assert second_bit == 0
print(second_bit)
