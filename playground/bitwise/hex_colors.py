a = 255 # alpha value
r = 222 # red value
g = 192 # green value
b = 222 # blue value

# unpaked raw code
print((r, g, b, a))


# pack code
rgba_code = (a << 24) | (r << 16) | (g << 8) | b
print(hex(rgba_code))


# unpack code
b = rgba_code & 255
g = (rgba_code >> 8) & 255
r = (rgba_code >> 16) & 255
a = (rgba_code >> 24) & 255

print((r, g, b, a))
