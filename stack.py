class Empty(Exception):
    pass

class ArrayStack:
    """Stack implementation using a Python list as storage"""

    def __init__(self):
        self._data = []

    def __len__(self):
        return len(self._data)

    def __repr__(self):
        return "<Stack {}>".format(self._data)

    def is_empty(self):
        return not bool(self._data)

    def top(self):
        if self.is_empty():
            raise Empty('Stack is empty')
        return self._data[-1]

    def push(self, item):
        self._data.append(item)

    def pop(self):
        if self.is_empty():
            raise Empty('Stack is empty')
        return self._data.pop()

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
    return rec_rem(stack)


if __name__== "__main__":
    s = ArrayStack()
    t = ArrayStack()
    for i in range(15):
        s.push(i)
    transfer(s,t)
    print(s,t)
    rec_rem(t)
    print(s,t)

