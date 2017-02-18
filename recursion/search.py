import random

def bin_search(seq, item):
    mid = len(seq)//2

    if len(seq) == 0:
        return False
    elif seq[mid] == item:
        return True
    elif item < seq[mid]:
        return bin_search(seq[:mid], item)
    elif item > seq[mid]:
        return bin_search(seq[mid+1:], item)

seq = sorted(random.sample(range(1,100), 16))
item = random.choice(seq)

print("Seq: ", seq)
print("Item: ", item)

print('res: ', bin_search(seq, item))