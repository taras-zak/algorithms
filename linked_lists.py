#TODO:  - size of node and whole list
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
        if pos is None:
            pos = self.size-1
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
            return curr
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

# TODO: refactor (pop, insert)

if __name__=="__main__":
    l = LinkedList()
