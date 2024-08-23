import argparse
import secrets

from pactus.crypto import CryptoConfig
from pactus.crypto.address import AddressType, Address
from pactus.crypto.bls.private_key import PrivateKey, PublicKey


def main() -> None:
    parser = argparse.ArgumentParser(description="Create a Key Pair")

    parser.add_argument(
        "--testnet",
        action="store_true",
        help="Specify if the key should be created for the testnet",
    )

    parser.add_argument(
        "--have",
        type=str
    )
    args = parser.parse_args()

    if args.testnet:
        CryptoConfig.use_testnet()

    while 1:    
        ikm = secrets.token_bytes(32)
        sec = PrivateKey.key_gen(ikm)
        pub = sec.public_key()
        addr = pub.account_address()
        print(addr.string())
        if args.have in addr.string():
            show(sec, pub, addr)
            break

def show(sec: PrivateKey, pub: PublicKey, addr: Address):
    print(f"Your PrivateKey: {sec.string()}")
    print(f"Your PublicKey: {pub.string()}")
    print(f"Your Address: {addr.string()}")


if __name__ == "__main__":
    main()
