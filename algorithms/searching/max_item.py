import random

def max_(seq):
    if len(seq) == 1:
        return seq[0]
    else:
        m = max(seq[1:])
        return m if m > seq[0] else seq[0]


seq = random.sample(range(1,100), 16)
print(seq)
print(max_(seq), max(seq))