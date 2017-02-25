import unittest

from data_structures.stack import ListStack, ArrayStack
from utils.exeptions import Empty


class CommonTests:

    def test_push_and_top_methods(self):
        self.assertRaises(Empty, self.stack.top)
        self.stack.push(2)
        self.assertEqual(self.stack.top(), 2)

        self.stack.push(55)
        self.assertEqual(self.stack.top(), 55)

    def test_len_method(self):
        self.assertEqual(len(self.stack), 0)
        self.stack.push(2)
        self.assertEqual(self.stack.top(), 2)
        self.assertEqual(len(self.stack), 1)

        self.stack.push(55)
        self.assertEqual(self.stack.top(), 55)
        self.assertEqual(len(self.stack), 2)

    def test_pop_method(self):
        self.stack.push(55)
        self.assertEqual(self.stack.pop(), 55)

        self.stack.push(100)
        self.assertEqual(self.stack.pop(), 100)

        self.assertRaises(Empty, self.stack.pop)


class TestListStackMethods(unittest.TestCase, CommonTests):

    def setUp(self):
        self.stack = ListStack()

    def tearDown(self):
        self.stack = None

class TestArrayStackMethods(unittest.TestCase, CommonTests):

    def setUp(self):
        self.stack = ArrayStack()

    def tearDown(self):
        self.stack = None

if __name__ == '__main__':
    unittest.main()