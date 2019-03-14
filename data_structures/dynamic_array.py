import ctypes


class DynamicArray:
    """A dynamic array class"""
    def __init__(self):
        """Create an empty array"""
        self._n = 0
        self._capacity = 1
        self._array = self._make_array(self._capacity)

    def __len__(self):
        return self._n

    def __getitem__(self, k):
        if 0<= k < self._n:
            return self._array[k]
        elif -self._n <= k < 0:
            return self._array[k + self._n]
        else:
            raise IndexError('invalid index')


    def append(self, obj):
        if self._n == self._capacity:
            self._resize(self._capacity*2)
        self._array[self._n] = obj
        self._n += 1

    def _resize(self, c):
        """Resize internal array to capacity x"""
        new_array = self._make_array(c)
        for i in range(self._n):
            new_array[i] = self._array[i]
        self._array = new_array
        self._capacity = c

    def _make_array(self, c):
        return (c * ctypes.py_object)()


if __name__ == "__main__":
    arr = DynamicArray()
    arr.append(1)
    arr.append(2)
    arr.append(3)
    arr.append(4)
    arr.append(5)
    print(arr[-6])

