from stack import ArrayStack, Empty

class ArrayQueue:
    """Queue implementation using a Python list circular fashion,
       by avoiding pop(0) call for O(n) time.
    """
    DEFAULT_CAPACITY = 10

    def __init__(self):
        self._data = [None] * ArrayQueue.DEFAULT_CAPACITY
        self._size = 0
        self._front = 0

    def __len__(self):
        return self._size

    def is_empty(self):
        return not bool(self._size)

    def first(self):
        if self.is_empty():
            raise Empty('Queue is empty')
        return self._data[self._front]

    def enqueue(self, item):
        """Add an element to the back of queue."""
        if self._size == len(self._data):
            self._resize(2 * len(self._data))

        back = (self._front + self._size) % len(self._data)
        self._data[back] = item
        self._size += 1

    def dequeue(self):
        """Remove and return the first element of the queue (i.e., FIFO)

        Raise Empty exception if the queue is empty
        """
        if self.is_empty():
            raise Empty('Queue is empty')

        item = self._data[self._front]
        self._data[self._front] = None
        self._front = (self._front + 1) % len(self._data)
        self._size -= 1

        if 0 < self._size < len(self._data) // 4:
            self._resize(len(self._data) // 2)
        return item

    def _resize(self, cap):
        """Resize to a new list of capacity >= len(self)."""

        old = self._data
        self._data = [None] * cap
        front = self._front
        for i in range(self._size):
            self._data[i] = old[front]
            front = (front + 1) % len(old)
        self._front = 0

class QueueOnStacks:
    pass

if __name__=="__main__":
    q = ArrayQueue()
    for i in range(11):
        q.enqueue(i)
    print(q.first())
