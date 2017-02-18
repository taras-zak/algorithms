import unittest
from dequeue import ArrayDeque
from stack import Empty


class TestDequeueMethods(unittest.TestCase):

    def testFirstMethod(self):
        d = ArrayDeque()
        self.assertRaises(Empty, d.first)
        d.add_first(3)
        self.assertEquals(d.first(), 3)
        d.add_first(10)
        self.assertEquals(d.first(), 10)

    def testLastMethod(self):
        d = ArrayDeque()
        self.assertRaises(Empty, d.last)
        d.add_first(3)
        self.assertEquals(d.last(), 3)
        d.add_first(10)
        self.assertEquals(d.last(), 3)

    def testAddMethods(self):
        d = ArrayDeque()
        d.add_first(3)
        self.assertEquals(d.last(), 3) and self.assertEquals(d.first(), 3)
        d.add_last(10)
        self.assertEquals(d.last(), 10)
        d.add_last(155)
        self.assertEquals(d.last(), 155)

    def testDeleteMethods(self):
        d = ArrayDeque()
        for i in range(3):
            d.add_last(i)
        self.assertEquals(d.delete_first(), 0)
        self.assertEquals(d.delete_last(), 2)
        d.delete_last()
        self.assertRaises(Empty, d.delete_last)
        self.assertRaises(Empty, d.delete_first)

#TODO:
#  - test Empty messages
#  - test resize deque

if __name__ == "__main__":
    unittest.main()