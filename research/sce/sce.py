# sce: stupid cryptography standards.
# just implementing stuff i lean about cryptography for fun.

import os
import bech32

# from: https://asecuritysite.com/encryption/nprimes?y=256
P = 90475314063646309186333911499545516020878393975510354007683035359907836935351
G = 97196008581366143336110500868632826688558438691823781666953410301647280664993

PRIVATE_KEY_PREFIX = "ss"
PUBLIC_KEY_PREFIX = "sp"
SHARED_SECRET_PREFIX = "sss"

def generate_key_pair():
    random_bytes = os.urandom(32)
    priv = int.from_bytes(random_bytes, 'big')
    pub = pow(G, priv, P)
    return priv, pub

def get_secret(pub, priv):
    return pow(pub, priv, P)

def int_to_bech32_data(n):
    byte_array = n.to_bytes((n.bit_length() + 7) // 8, 'big')
    return [b % 32 for b in byte_array]

alice_priv, alice_pub = generate_key_pair()
bob_priv, bob_pub = generate_key_pair()

print("Alice Private Key:", bech32.bech32_encode(PRIVATE_KEY_PREFIX, int_to_bech32_data(alice_priv)))
print("Alice Public Key:", bech32.bech32_encode(PUBLIC_KEY_PREFIX, int_to_bech32_data(alice_pub)))

print("----------------------------------\n")

print("Bob Private Key:", bech32.bech32_encode(PRIVATE_KEY_PREFIX, int_to_bech32_data(bob_priv)))
print("Bob Public Key:", bech32.bech32_encode(PUBLIC_KEY_PREFIX, int_to_bech32_data(bob_pub)))

print("----------------------------------\n")

alice_sec = get_secret(bob_pub, alice_priv)
bob_sec = get_secret(alice_pub, bob_priv)

print("Bob Secret  ", bech32.bech32_encode(SHARED_SECRET_PREFIX, int_to_bech32_data(bob_sec)))
print("Alice Secret", bech32.bech32_encode(SHARED_SECRET_PREFIX, int_to_bech32_data(alice_sec)))

assert alice_sec == bob_sec
