from pprint import pprint
import collections

#TODO: clean code, implement Dijkstra and A* algorithms


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

    def draw(self, path_tracking=None, start=None, target=None, debug=True):
        if debug:
            print('   ', end='')
            print(''.join("{:<2}".format(i) for i in  range(self.cols)))

        for row in range(self.rows):
            if debug:
                print("{:2} ".format(row), end='')
            for col in range(self.cols):
                cell = ". "
                if self.edges[row][col]:
                    cell = '# '
                if path_tracking and path_tracking.get(Coord(row,col), None) is not None:
                    to = path_tracking[Coord(row,col)]
                    if to.col == col + 1: cell = "\u2192 "
                    if to.col == col - 1: cell = "\u2190 "
                    if to.row == row + 1: cell = "\u2193 "
                    if to.row == row - 1: cell = "\u2191 "
                if start and start == Coord(row,col): cell = "S "
                if target and target == Coord(row, col): cell = "T "
                print(cell, end='')
            print()

Coord = collections.namedtuple('Node_coord', 'row,col')

def breadth_first_search(graph, start, end):
    frontier = Queue()
    frontier.put(start)
    came_from = {}
    came_from[start] = None

    while not frontier.empty():
        current = frontier.get()
        #print("Visititng {}".format(current))

        if current == end:
            break

        n = graph.neighbors(current)
        for next_ in n:
            if next_ not in came_from:
                frontier.put(next_)
                came_from[next_] = current
    return came_from

def reconstruct_path(points, end):
    path = []
    current_cell = end
    while points[current_cell]:
        next = points[current_cell]
        if current_cell.col + 1 == next.col: cell = "W"
        if current_cell.col - 1 == next.col: cell = "E"
        if current_cell.row + 1 == next.row: cell = "N"
        if current_cell.row - 1 == next.row: cell = "S"
        current_cell = next
        path.append(cell)
    return path

def main(maize_map):
    graph = SimpleGraph(12, 12)
    graph.edges = maize_map

    start = Coord(1,1)
    target = Coord(10,10)
    path = breadth_first_search(graph, start, target)
    graph.draw(path, start, target)

    print(list(reversed(reconstruct_path(path, Coord(10, 10)))))

if __name__ == '__main__':
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

        [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
    ]
    main(labyrinth_example)



