import os

def du(path):
    total = os.path.getsize(path)
    if os.path.isdir(path):
        for filename in os.listdir(path):
            childpath = os.path.join(path, filename)
            total += du(childpath)

    print("{0:<7} {1}".format(total, path))
    return total

du("courses")