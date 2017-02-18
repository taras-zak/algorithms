from time import time
import sys

#TODO: read https://www.quora.com/How-are-Python-lists-implemented-internally

def compute_avg_append(n):
    data = []
    start = time()
    for k in range(n):
        data.append(None)
    end = time()
    return (end-start)/n

def compute_avg_pop_first(n):
    data = [None for i in range(n)]
    start = time()
    for k in range(n):
        data.pop(0)
    end = time()
    return (end-start)/n


def compute_list_size_append(n):
    data = []
    size = sys.getsizeof(data)
    for k in range(n):
        a = len(data)
        b = sys.getsizeof(data)
        if size != b:
            size = b
            print('Length: {0:3d}; Size in bytes: {1:4d}'.format(a, b))
        data.append(None)

def compute_list_size_pop(n):
    data = [None for i in range(n)]
    size = sys.getsizeof(data)
    for k in range(n):
        a = len(data)
        b = sys.getsizeof(data)
        if size != b:
            size = b
            print('Length: {0:3d}; Size in bytes: {1:4d}'.format(a, b))
        data.pop()

print(compute_avg_pop_first(1000000))
