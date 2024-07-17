a = ord('k')
key = 0x7

cipher = a ^ key

print(f"Cipher is: {chr(cipher)}")

plain = cipher ^ 0x7

print(f"Plain is: {chr(plain)}")
