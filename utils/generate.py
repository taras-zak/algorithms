import random


def generate_array(length=33, min=0, max=10):
    data = [random.randint(min, max) for _ in range(length + 1)]
    return data

def generate_array(lenth=33):
    data = list(range(lenth))
    random.shuffle(data)
    return data
