from data_structures.stack import Empty
from data_structures.queue_custom import ArrayQueue


class ArrayDeque(ArrayQueue):

    DEFAULT_CAPACITY = 5

    def __init__(self):
        super().__init__()
        self._data = [None] * ArrayDeque.DEFAULT_CAPACITY

    def add_first(self, item):
        """Add an element to the front of deque."""
        if self._size == len(self._data):
            self._resize(2 * len(self._data))

        self._front = (self._front - 1) % len(self._data)
        self._data[self._front] = item
        self._size += 1

    def add_last(self, item):
        """Add an element to the rear of deque."""
        super().enqueue(item)

    def delete_last(self):
        """Remove and return the rear element of the deque

        Raise Empty exception if the deque is empty
        """
        if self.is_empty():
            raise Empty('Deque is empty')

        back = (self._front + self._size - 1) % len(self._data)
        item = self._data[back]
        self._data[back] = None
        self._size -= 1

        if 0 < self._size < len(self._data) // 4:
            self._resize(len(self._data) // 2)
        return item

    def delete_first(self):
        if self.is_empty():
            raise Empty('Deque is empty')
        return super().dequeue()

    def last(self):
        if self.is_empty():
            raise Empty('Deque is empty')

        back = (self._front + self._size - 1) % len(self._data)
        return self._data[back]
