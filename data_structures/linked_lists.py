from utils.exeptions import Empty

#TODO:  - size of node and whole list
#       - node use __slots__
#       - make pythonic

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

    def get_data(self):
        return self.data

    def get_next(self):
        return self.next

    def set_data(self, data):
        self.data = data

    def set_next(self, next_node):
        self.next = next_node


class LinkedList:
    # TODO: refactor (pop, insert)
    def __init__(self):
        self.size = 0
        self.head = None

    def add(self, item):
        temp = Node(item)
        temp.set_next(self.head)
        self.head = temp
        self.size += 1

    def search(self, item):
        curr = self.head
        found = False
        while curr and not found:
            if curr.get_data() == item:
                found = True
            else:
                curr = curr.get_next()
        return found

    def remove(self, item):
        current = self.head
        previous = None
        found = False
        while not found and current:
            if current.get_data() == item:
                found = True
            else:
                previous = current
                current = current.get_next()

        if found:
            if previous is None:
                self.head = current.get_next()
            else:
                previous.set_next(current.get_next())
            self.size -= 1
        else:
            raise ValueError('Item {} not in LinkedList'.format(item))

    def append(self, item):
        self.add(item)

    def index(self, item):
        """Returns index of item in list. If no such item raises ValueError"""
        index = 0
        curr = self.head
        while curr:
            if curr.get_data() == item:
                return index
            curr = curr.get_next()
            index +=1
        raise ValueError('Item {} not in LinkedList'.format(item))

    def insert(self, pos, item):
        """Insert item in pos. If pos not in [0, size_of_list) raises IndexError"""
        if 0 <= pos < self.size:
            index = 0
            curr = self.head
            prev = None
            while index != pos:
                prev = curr
                curr = curr.get_next()
                index += 1
            temp = Node(item)
            if prev:
                prev.set_next(temp)
            else:
                # insert in head
                self.head = temp
            temp.set_next(curr)
            self.size += 1
        else:
            raise IndexError('Position {} out of range'.format(pos))

    def pop(self, pos=None):
        """Remove and return the element at pos (default from the head of list)
        Raises IndexError if list is empty or index is out of range.
        """
        if pos is None:
            pos = 0
        if 0 <= pos < self.size:
            index = 0
            curr = self.head
            prev = None
            while index != pos:
                prev = curr
                curr = curr.get_next()
                index += 1
            if prev is not None:
                prev.set_next(curr.get_next())
            else:
                self.head = curr.get_next()
            self.size -= 1
            return curr.get_data()
        else:
            raise IndexError('Position {} out of range'.format(pos))

    def __repr__(self):
        # DEBUG
        res = []
        curr = self.head
        while curr:
            res.append(curr.get_data())
            curr = curr.get_next()
        return ' '.join((map(str,res)))

class CircularQueue:

    class _Node:
        __slots__ = '_data', '_next'

        def __init__(self, data, next):
            self._data = data
            self._next = next

    def __init__(self):
        self._tail = None
        self._size = 0

    def __len__(self):
        return self._size

    def is_empty(self):
        return not bool(self._size)

    def first(self):
        if self.is_empty():
            raise Empty('Queue empty')
        head = self._tail._next
        return head._data

    def dequeue(self):
        if self.is_empty():
            raise Empty('Queue empty')
        oldhead = self._tail._next
        if self._size == 1:
            self._tail = None
        else:
            self._tail._next = oldhead._next
        self._size -= 1
        return oldhead._data

    def enqueue(self, item):
        new_item = self._Node(item, None)
        if self.is_empty():
            new_item._next = new_item
        else:
            new_item._next = self._tail._next
            self._tail._next = new_item
        self._tail = new_item
        self._size += 1

    def rotate(self):
        if self._size > 0:
            self._tail = self._tail._next


if __name__=="__main__":
    l = LinkedList()
