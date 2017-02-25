from data_structures.linked_lists import LinkedList

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

class ListStack:
    """Stack implementation using custom linked_list as storage"""

    def __init__(self):
        self._data = LinkedList()

    def __len__(self):
        return self._data.size

    def is_empty(self):
        return not bool(self._data.size)

    def top(self):
        if self.is_empty():
            raise Empty('Stack is empty')
        return self._data.head.data

    def push(self, item):
        self._data.append(item)

    def pop(self):
        if self.is_empty():
            raise Empty('Stack is empty')
        return self._data.pop()
