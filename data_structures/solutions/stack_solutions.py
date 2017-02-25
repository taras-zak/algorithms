from data_structures.stack import ArrayStack

def transfer(s,t):
    """R-6.3 solution"""
    while not s.is_empty():
        item = s.pop()
        t.push(item)

def rec_remove(stack):
    """R-6.4 solution"""
    if stack.is_empty():
        return
    stack.pop()
    return rec_remove(stack)

if __name__ == "__main__":
    s = ArrayStack()
    t = ArrayStack()
    for i in range(15):
        s.push(i)
    transfer(s,t)
    print(s,t)
    rec_remove(t)
    print(s,t)