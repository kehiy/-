import time

# generates a seed from unix time stamp
def gen_seed():
    seed = int(str(time.time()).replace('.', '')[-8:])
    return seed

# return a pseudorandom number
# seed : your seed number
# length : your random number length
def random(seed : int, length : int = 10):
    return str(int((int(str(seed)[-4:]) * seed / int(str(seed)[5:])) * (seed ** 8)))[-length:]
    


print(random(gen_seed(),1))