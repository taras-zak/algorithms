class HeapWithRemoval:

    def __init__(self):
        self.heap = []
        # Делаем не set, а словарь со счётчиком на случай повторов
        self.deleted = collections.defaultdict(int)

    def add(self, value):
        if self.deleted[value]:
            self.deleted[value] -= 1
        else:
            heapq.heappush(self.heap, value)

    def remove(self, value):
        self.deleted[value] += 1
        while self.heap and self.deleted[self.heap[0]]:
            self.deleted[self.heap[0]] -= 1
            heapq.heappop(self.heap)

    def minimum(self):
        return self.heap[0] if self.heap else 0


class Point:
    def __init__(self, x, h, t):
        self.x = x
        self.h = h
        self.t = t

    def __lt__(self, point):
        if self.x == point.x:
            return (self.t * self.h) > (point.t * point.h)
        return self.x < point.x

    def __repr__(self):
        return "<x:{}, h:{}, t:{}>".format(self.x, self.h, self.t)


class Solution:
    def getSkyline(self, buildings):
        """
        :type buildings: List[List[int]]
        :rtype: List[List[int]]
        """
        res = []
        pq = HeapWithRemoval()
        points = []
        for building in buildings:
            points.append(Point(building[0], building[2], 1))
            points.append(Point(building[1], building[2], -1))
        points.sort()
        print("All points: ", points)
        for point in points:
            print("point: ", point)
            # m1 = pq.minimum()
            if point.t == 1:
                if point.h > -pq.minimum(): res.append([point.x, point.h])
                pq.add(-point.h)
            else:
                pq.remove(-point.h)
                if point.h > -pq.minimum(): res.append([point.x, -pq.minimum()])
            # TODO: check if needs when points array doesnt have duplicates
            # if m1 == pq.minimum():
            # continue
            # res.append([point.x, -pq.minimum()])
            print("res: ", res)
        return res