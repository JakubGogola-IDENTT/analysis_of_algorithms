from threading import Thread
from queue import Queue

# object which marks end of data
_sentinel = object()

def generate_data(multiset_queue, size):
    for i in range(size):
        multiset_queue.put((i, 1.0))

    multiset_queue.put(_sentinel)