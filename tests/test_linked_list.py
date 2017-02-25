import unittest

from data_structures.linked_lists import LinkedList


class TestLinkedListMethods(unittest.TestCase):

    def test_append_method(self):
        l = LinkedList()
        l.append(3)
        self.assertEqual(l.size, 1)
        l.append(4)
        self.assertEqual(l.size, 2)
        self.assertEqual(l.head.data, 4)

    def test_delete_method(self):
        l = LinkedList()
        for i in range(10, 0, -1):
            l.add(i)
        #remove middle
        l.remove(5)
        self.assertEqual(l.size, 9)
        #remove tail
        l.remove(10)
        self.assertEqual(l.size, 8)
        #remove head
        l.remove(1)
        self.assertEqual(l.size, 7)
        self.assertEqual(l.head.get_data(), 2)

    def test_search_digit(self):
        l = LinkedList()
        for i in range(5):
            l.add(i)
        self.assertEqual(l.search(4), True)
        self.assertEqual(l.search(9), False)

    def test_search_string(self):
        l = LinkedList()
        l.append('test')
        self.assertEqual(l.search('test'), True)

    def test_index_method(self):
        l = LinkedList()
        for i in range(5):
            l.add(i)
        self.assertEqual(l.index(4),0)
        self.assertEqual(l.index(2), 2)
        self.assertRaises(ValueError, l.index, 10)

    def test_insert_method(self):
        l = LinkedList()
        for i in range(5, 0, -1):
            l.add(i)

        # insert head
        l.insert(0, 99)
        self.assertEqual(l.index(99), 0)
        # insert middle
        l.insert(3, 88)
        self.assertEqual(l.index(88), 3)
        # insert pos==size-1
        l.insert(l.size-1, 77)
        self.assertEqual(l.index(77), 6)

        self.assertRaises(IndexError, l.insert, -1, 22)
        self.assertRaises(IndexError, l.insert, 100, 22)

    def test_pop_method(self):
        l = LinkedList()
        for i in range(5, 0, -1):
            l.add(i)

        self.assertEqual(l.pop(0), 1)
        self.assertEqual(l.pop(2), 4)
        self.assertEqual(l.pop(l.size-1), 5)

        self.assertRaises(IndexError, l.pop, -1)
        self.assertRaises(IndexError, l.pop, 100)

if __name__ == "__main__":
    unittest.main()