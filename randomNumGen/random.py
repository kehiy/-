import time

# get unix time stamp as seed
# return each of numbers in timestamp to power of  3 + i and add to result
def random():
    seed = str(time.time())
    result = 0
    for i in seed:
        if i == ".":
            continue
        result += int(i) ** 3 + int(i)
    return result


print(random()) # => 3504