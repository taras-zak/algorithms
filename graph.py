from pprint import pprint
import collections


class Queue:
    def __init__(self):
        self.elements = collections.deque()

    def empty(self):
        return len(self.elements) == 0

    def put(self, x):
        self.elements.append(x)

    def get(self):
        return self.elements.popleft()

class SimpleGraph:
    def __init__(self, rows, cols):
        self.rows = rows
        self.cols = cols
        self.edges = []
        for row in range(rows):
            self.edges.append(list('-'.join(map(str, (row, col))) for col in range(cols)))

    def neighbors(self, coord):
        dirs = (Coord(1, 0), Coord(0, 1), Coord(-1, 0), Coord(0, -1))
        result = []
        for _dir in dirs:
            row = coord.row + _dir.row
            col = coord.col + _dir.col
            if 0 <= row < self.rows and 0 <= col < self.cols and self.edges[row][col]:
                result.append(Coord(row, col))
        return result

Coord = collections.namedtuple('Node_coord', 'row,col')

def breadth_first_search(graph, start):
    frontier = Queue()
    frontier.put(start)
    visited = {}
    visited[start] = True

    while not frontier.empty():
        current = frontier.get()
        print("Visititng {}".format(current))
        n = graph.neighbors(current)
        for next_ in n:
            if next_ not in visited:
                frontier.put(next_)
                visited[next_] = True


graph = SimpleGraph(5, 5)
pprint(graph.edges)
print(graph.neighbors(Coord(2,4)),end='\n\n')

breadth_first_search(graph, Coord(0,0))




