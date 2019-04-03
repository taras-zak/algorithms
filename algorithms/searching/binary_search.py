import random
import timeit


def bin_search(seq, item, start=0, stop=None):
    if stop is None:
        stop = len(seq)

    mid = start + (stop - start) // 2
    if (stop - start) == 0:
        return False
    elif seq[mid] == item:
        return True, mid
    elif item < seq[mid]:
        return bin_search(seq, item, start=start, stop=mid)
    elif item > seq[mid]:
        return bin_search(seq, item, start=mid+1, stop=stop)

def linear_search(seq, item):
    for pos, i in enumerate(seq):
        if i == seq:
            return (True, pos)
    return False

def test_timings():
    for n in range(18):
        print("N = ", 2 ** n)
        print(
            "\tLinear time: ",
            timeit.timeit(
                "linear_search(data, item)",
                setup=f"from __main__ import linear_search, random; data = sorted(random.sample(range(1,{2**n + 1}), {2**n})); item = random.choice(data)",
                number=1000,
            ),
        )
        print(
            "\tBin search time: ",
            timeit.timeit(
                "bin_search(data, item)",
                setup=f"from __main__ import bin_search, random; data = sorted(random.sample(range(1,{2**n + 1}), {2**n})); item = random.choice(data)",
                number=1000,
            ),
        )


if __name__=="__main__":
    seq = sorted(random.sample(range(1,100), 16))
    item = random.choice(seq)
    print("Seq: ", seq)
    print("Item: ", item)
    print('res: ', bin_search(seq, item))
    test_timings()


