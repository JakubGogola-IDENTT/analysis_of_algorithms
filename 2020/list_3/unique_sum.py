import math
from hashlib import sha256
from helpers import get_hash
from data import _sentinel

def unique_sum(multiset_queue, m):
    M = [math.inf] * m

    while True:
        element = multiset_queue.get()

        if element is _sentinel:
            multiset_queue.put(_sentinel)
            break

        for k, hash in enumerate(M):
            value_to_hash = bytes(element[0]) + bytes(k + 1)

            u = get_hash(value_to_hash)
            
            estimation = -(math.log(u) / element[1])

            M[k] = min(hash, estimation)
    
    sum_of_hashes = 0.0

    for hash in M:
        sum_of_hashes += hash

    print(float(m-1) / sum_of_hashes)



    

