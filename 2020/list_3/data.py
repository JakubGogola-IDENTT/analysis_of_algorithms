from random import random

# object which marks end of data
_sentinel = object()

def generate_data(multiset_queue, size):
    for i in range(size):
        multiset_queue.put((i, 1.03221312131))

    multiset_queue.put(_sentinel)