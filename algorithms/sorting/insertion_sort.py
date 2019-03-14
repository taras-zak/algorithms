import random

def insertion_sort(arr):
    for i in range(1, len(arr)):
        item = arr[i]
        j = i
        while j > 0 and arr[j-1] > item:
            arr[j] = arr[j-1]
            j -= 1
        arr[j] = item
    return arr


def gen_rand_array(n):
    arr = range(1, n+1)
    random.shuffle(arr)
    return arr

if __name__ == '__main__':

    arr = list(range(1,120))
    random.shuffle(arr)
    print(arr)
    print(insertion_sort(arr))