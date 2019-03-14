import random
import timeit

data = [100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97, 141]


def subarrays(array):
    res = []
    for start in range(len(array)):
        for end in range(1, len(array) - start + 1):
            res.append(array[start:start + end])
    return res


def max_subarray1(array):
    """O(n^2)"""
    return max(list(map(sum, subarrays(array)))), sorted(
        map(sum, subarrays(array)), reverse=True
    )


def max_subarray(array):
    """O(n^2)"""
    max_s = float("-inf")
    si, mi = None, None
    for start in range(len(array)):
        for end in range(1, len(array) - start + 1):
            s = sum(array[start:start + end])
            if s > max_s:
                max_s, si, mi = s, start, start + end
    return max_s, si, mi, array[si:mi]


def max_subarray2(array, lo=0, hi=None):
    if hi is None:
        hi = len(array)

    def crossing(arr, lo, mid, hi):
        left_s = float("-inf")
        s = 0
        for i in reversed(range(lo, mid)):
            s = arr[i] + s
            if s > left_s:
                left_s = s
                left_idx = i
        right_s = float("-inf")
        s = 0
        for j in range(mid, hi):
            s = arr[j] + s
            if s > right_s:
                right_s = s
                right_idx = j
        return left_s + right_s, left_idx, right_idx + 1

    if hi - 1 == lo:
        return array[lo], lo, hi

    elif not len(array):
        return 0, lo, lo

    else:
        mid = (lo + hi) // 2
        left_sum, ll, lh = max_subarray2(array, lo, mid)
        right_sum, rl, rh = max_subarray2(array, mid, hi)
        cross, cl, ch = crossing(array, lo, mid, hi)
        if left_sum >= right_sum and left_sum >= cross:
            s, lo, hi = left_sum, ll, lh
        elif right_sum >= left_sum and right_sum >= cross:
            s, lo, hi = right_sum, rl, rh
        else:
            s, lo, hi = cross, cl, ch
        return s, lo, hi


def max_subarray_linear(data):
    max_sum = 0
    sum = 0
    for i in data:
        sum += i
        max_sum = max(max_sum, sum)
        if sum < 0:
            sum = 0
    return max_sum


def test_empty_array():
    data = []
    run(data)

def run(data=None, n=33):
    if data is None:
        data = generate_array(n)
    print(max_subarray(data))
    t = max_subarray2(data)
    print(t, data[t[1]:t[2]])
    print(max_subarray_linear(data))


def generate_array(lenth=33):
    data = [random.randint(-10, 10) for _ in range(lenth + 1)]
    data = list(e - data[i - 1] for i, e in enumerate(data[1:], 1))
    return data

def test_timings():
    for n in range(20):
        print("N = ", 2 ** n)
        print(
            "\tLinear time: ",
            timeit.timeit(
                "max_subarray_linear(data)",
                setup=f"from __main__ import generate_array, max_subarray_linear;data = generate_array(lenth={2**n})",
                number=100,
            ),
        )
        if 2 ** n <= 1024:
            print(
                "\tRecursive time: ",
                timeit.timeit(
                    "max_subarray2(data)",
                    setup=f"from __main__ import generate_array, max_subarray2;data = generate_array(lenth={2**n})",
                    number=100,
                ),
            )
        if 2 ** n < 128:
            print(
                "\tBute time: ",
                timeit.timeit(
                    "max_subarray(data)",
                    setup=f"from __main__ import generate_array, max_subarray;data = generate_array(lenth={2**n})",
                    number=100,
                ),
            )

if __name__ == "__main__":
    run()
    test_timings()