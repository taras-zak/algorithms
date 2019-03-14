from random import randint
from utils.generate import generate_array


def partition(arr, start, stop):
    pivot = arr[stop]
    i = start - 1
    for j in range(start, stop):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i+1], arr[stop] = arr[stop], arr[i+1]
    return i+1

def quicksort(arr, start=0, stop=None, random=False):
    if stop is None:
        stop = len(arr)-1
    if start > stop:
        return
    if random:
        pivot = randint(start, stop)
        arr[stop], arr[pivot] = arr[pivot], arr[stop]
    mid = partition(arr, start, stop)
    quicksort(arr, start, mid-1, random)
    quicksort(arr, mid + 1, stop, random)

if __name__ == "__main__":
    # arr = generate_array(10)
    arr = [2,2,2,2,2,2,2,2,2,2]
    print(arr)
    print(": ", quicksort(arr, random=True))
    print(arr)