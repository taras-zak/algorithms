from utils.generate import generate_array


def counting_sort(arr):
    max_value = max(arr)

    c = [0] * (max_value + 1)
    res = [0] * (len(arr))

    for i in arr:
        c[i] += 1

    for j in range(1, len(c)):
        c[j] += c[j-1]

    for k in arr[::-1]:
        res[c[k]-1] = k
        print('res: ', res)
        c[k] -= 1
    return res




if __name__ == "__main__":
    # TODO: radix sort
    # TODO: bucket sort

    arr = [2,2,5,2,2,3,2,8,2,4]
    print(arr)
    print(counting_sort(arr))
