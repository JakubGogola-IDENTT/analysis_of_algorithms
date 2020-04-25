from threading import Thread
from queue import Queue
from data import generate_data
from unique_sum import unique_sum

if __name__ == '__main__':
    multiset_queue = Queue()
    m = 1000

    t1 = Thread(target=generate_data, args=(multiset_queue, m))
    t2 = Thread(target=unique_sum, args=(multiset_queue, m))

    t1.start()
    t2.start()

    t1.join()
    t2.join()
