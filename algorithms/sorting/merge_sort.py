import random


def merge(l,r):
    res = []
    i,j = 0,0
    while i < len(l) and j < len(r):
        if l[i] < r[j]:
            res.append(l[i])
            i += 1
        else:
            res.append(r[j])
            j += 1
    res += l[i:]
    res += r[j:]
    return res

def merge_sort(seq):
    if len(seq) == 1:
        return seq
    q = int(len(seq)/2)
    left = merge_sort(seq[:q])
    right = merge_sort(seq[q:])
    return merge(left, right)

if __name__ == '__main__':
    arr = range(1,10)
    random.shuffle(arr)
    print arr
    print merge_sort(arr)