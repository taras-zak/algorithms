from pprint import pprint
import collections

labyrinth_example = [

    [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],

    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],

    [1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1],

    [1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1],

    [1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1],

    [1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1],

    [1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1],

    [1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1],

    [1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1],

    [1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1],

    [1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1],

    [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]]

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
            if 0 <= row < self.rows and 0 <= col < self.cols and not self.edges[row][col]:
                result.append(Coord(row, col))
        return result

Coord = collections.namedtuple('Node_coord', 'row,col')

def breadth_first_search(graph, start, end):
    frontier = Queue()
    frontier.put(start)
    came_from = {}
    came_from[start] = None

    while not frontier.empty():
        current = frontier.get()
        print("Visititng {}".format(current))

        if current == end:
            break

        n = graph.neighbors(current)
        for next_ in n:
            if next_ not in came_from:
                frontier.put(next_)
                came_from[next_] = current


graph = SimpleGraph(5, 5)
graph.edges = labyrinth_example
graph.cols = 12
graph.rows = 12
pprint(graph.edges)
print(graph.neighbors(Coord(3,3)),end='\n\n')

breadth_first_search(graph, Coord(1,1), Coord(10,10))




