import math
import random
from io import StringIO

class Heap:
    def __init__(self, _heap_array=None):
        if _heap_array is None:
            _heap_array = []
        self._heap = _heap_array

    def root(self):
        return self._heap[0]

    def parent(self, i):
        return 0 if i == 0 else (i-1)//2

    def right(self, i):
        return 2*i+1

    def left(self, i):
        return 2*i+2

    def max_heapify(self, i, n=None):
        """Push element i onto heap with maintaining the heap invariant. Recursive implementation O(log(n))"""
        if n is None:
            n = len(self._heap)
        l = self.left(i)
        r = self.right(i)
        largest = i
        if l < n and self._heap[i] < self._heap[l]:
            largest = l
        if r < n and self._heap[largest] < self._heap[r]:
            largest = r
        if largest != i:
            self._heap[i], self._heap[largest] = self._heap[largest], self._heap[i]
            self.max_heapify(largest, n)

    def heapify(self):
        """Build heap from unsorted array in O(n)"""
        for i in range(len(self._heap)//2, -1, -1):
            self.max_heapify(i)

    def heapsort(self):
        """Make sorted array from heap in O(n*log(n))"""
        res = []
        self.heapify() # TODO: move to init
        for i in range(len(self._heap), 0, -1):
            self._heap[-1], self._heap[0] = self._heap[0], self._heap[-1]
            res.append(self._heap.pop(-1))
            self.max_heapify(0, len(self._heap))
        return res

    def iter_heapify(self, i):
        """Iterative implementation of push method of heap. O(log(n))"""
        while i < len(self._heap):
            l = self.left(i)
            r = self.right(i)
            value = self._heap[i]
            largest = i
            if l < len(self._heap) and self._heap[i] < self._heap[l]:
                largest = l
            if r < len(self._heap) and self._heap[i] < self._heap[r]:
                largest = r
            if largest != i:
                self._heap[i], self._heap[largest] = self._heap[largest], self._heap[i]
                i = largest
            else:
                break

    def __str__(self, i=0, depth=0):
        """Pretty-print a tree."""
        ret = ""
        value = self._heap[i]
        right = self.right(i)
        left = self.left(i)
        if left < len(self._heap):
            ret += self.__str__(left, depth + 1)
        ret += "\n" + ("    "*depth) + str(value)
        if right < len(self._heap):
            ret += self.__str__(right, depth + 1)
        return ret
    def __str__(self):
        return show_tree(self._heap)

class PriorityQueue(Heap):

    def __init__(self, arr):
        super().__init__(arr)
        self.heapify()

    def maximum(self):
        return self._heap[0]

    def extract_max(self):
        """Pop max element in O(log(n)) with maintaining the heap invariant"""
        if len(self._heap):
            self._heap[-1], self._heap[0] = self._heap[0], self._heap[-1]
            res = self._heap.pop(-1)
            self.max_heapify(0, len(self._heap))
            return res
        else:
            raise Exception("Heap is empty")

    def increase_key(self, key, new_value):
        if new_value < self._heap[key]:
            raise Exception("New value less than old")
        self._heap[key] = new_value
        while key > 0 and self._heap[key] > self._heap[self.parent(key)]:
            self._heap[key], self._heap[self.parent(key)] = self._heap[self.parent(key)], self._heap[key]
            key = self.parent(key)

    def insert(self, key):
        self._heap.append(-99999)
        self.increase_key(len(self._heap)-1, key)

    def _heapify(self, arr):
        self._heap = []
        for i in range(len(arr)):
            self.insert(arr[i])


def show_tree(tree, total_width=36*2, fill=' '):
    """Pretty-print a tree."""
    output = StringIO()
    last_row = -1
    for i, n in enumerate(tree):
        if i:
            row = int(math.floor(math.log(i + 1, 2)))
        else:
            row = 0
        if row != last_row:
            output.write('\n\n')
        columns = 2 ** row
        col_width = int(math.floor(total_width / columns))
        output.write(str(n).center(col_width, fill))
        last_row = row
    output.write('\n' + '-' * total_width)
    return output.getvalue()


if __name__ == '__main__':
    # TODO: min-heap (operator to heapify method ?)
    # TODO: tests for heap and priority queue

    def test_parent():
        h = Heap([16,14,10,8,7,9,3,2,4,1])
        show_tree(h._heap)
        for i,_ in enumerate(h._heap):
            print("Element: ", h._heap[i], "Parent: ", h._heap[h.parent(i)], "Index: ", h.parent(i))

    def test_extract():
        h = PriorityQueue([16,4,10,14,7,9,3,2,8,1])
        print(h.extract_max())
        print(h.extract_max())
        print(h.extract_max())
        print(h)

    def test_increment():
        h = PriorityQueue([16, 4, 11, 14, 7, 9, 3, 2, 8, 1])
        print(h)
        print("start incr")
        print(h._heap[4])
        print("Parent:", h._heap[h.parent(4)])
        print("Arr:", h._heap)
        h.increase_key(4, 20)

    def test_insert():
        h = PriorityQueue([16, 4, 11, 14, 7, 9, 3, 2, 8, 1])
        print(h)
        h.insert(15)
        print(h)

    def test_heapify_vs_insert():
        arr = generate_array(16)
        arr = list(range(16))
        # arr = list(reversed(range(16)))
        print(arr)

        h = PriorityQueue([])
        h._heapify(arr)
        print(len(arr), len(h._heap))
        print(h)
        print(h._heap)

        h = Heap(arr)
        h.heapify()
        print(h)
        print(h._heap)
        print(len(arr), len(h._heap))


    test_heapify_vs_insert()