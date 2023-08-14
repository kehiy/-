import os
import time

# generates a seed from unix time stamp and the pid
def gen_seed():
    pid = os.getppid()
    seed = int(str(time.time()).replace('.', '')[-8:])  + pid * (pid - 10)
    return seed

# return a pseudorandom number
# seed : your seed number
# length : your random number length
def random(seed : int, length : int = 10):
    return str(int((int(str(seed)[-2:]) * seed / int(str(seed)[-8:])) * (seed * 7)))[-length:]


