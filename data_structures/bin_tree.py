class Node:
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

    def __str__(self, depth=0):
        """Pretty-print a tree."""
        ret = ""
        if self.left:
            ret += self.left.__str__(depth + 1)
        ret += "\n" + ("    " * depth) + str(self.data)
        if self.right:
            ret += self.right.__str__(depth + 1)
        return ret


tree = Node(1)
tree.left = Node(2)
tree.right = Node(3)
tree.left.left = Node(4)
tree.left.right = Node(5)
tree.right.left = Node(6)
tree.right.right = Node(7)


def dfs(tree, acum=None):
    if acum is None:
        acum = []
    acum.append(tree.data)
    if tree.left:
        dfs(tree.left, acum)
    if tree.right:
        dfs(tree.right, acum)
    return acum

def bfs(tree):
    res = []
    res.append(tree.data)
    stack = []
    stack.append(tree.left)
    stack.append(tree.right)
    while stack:
        node = stack.pop()
        res.append(node.data)
        if node.left:
            stack.append(node.left)
        if node.right:
            stack.append(node.right)
    return res
print(tree)
print(dfs(tree))
print(bfs(tree))