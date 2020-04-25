from zlib import crc32
from hashlib import md5

def get_hash(value, hash_func = md5):
    digest = hash_func(bytes(value)).digest()

    return float(crc32(digest) & 0xffffffff) / (2**32 - 1)