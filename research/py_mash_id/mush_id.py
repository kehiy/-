import binascii
import base64
import time
import os
from uuid import getnode as getmac
import uuid

t1 = time.time()

pid = os.getpid()
mac_address = getmac()

def new_id(s: int):
    t = time.time()

    # mush = f"{mac_address}-{pid}-{t}-{s}"
    # s += 1

    mush = uuid.uuid4().__str__()

    #print(f"mac: {mac_address}, pid: {pid}, t: {t}")

    mush_crc32 = binascii.crc32(mush.encode('utf8')).to_bytes(4, 'big')
    mush_id = base64.b32encode(mush_crc32)

    print(str(mush_id, encoding='utf8'))

s = 0
for i in range(200_000):
    new_id(s)

print(time.time()-t1)
