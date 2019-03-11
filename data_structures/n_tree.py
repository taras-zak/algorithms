class Node:
    def __init__(self, data, parent=None, left_child=None, right_sibling=None):
        if right_sibling is None:
            right_sibling = []
        self.data = data
        self.parent = parent
        self.left_child = left_child
        self.right_siblings = right_sibling

    def add_children(self, node):
        self.left_child = node
        node.parent = self

    def add_sibling(self, node):
        node.parent = self.parent
        self.right_siblings.append(node)

tree = Node(1)
tree.add_children(Node(2))
tree.left_child.add_sibling(Node(3))
n = Node(4)
tree.left_child.add_sibling(n)
tree.left_child.add_sibling(Node(5))
n.add_children(Node(6))
n.left_child.add_sibling(Node(7))
n.left_child.add_sibling(Node(8))


def dfs(tree, acum=None):
    if acum is None:
        acum = []
    acum.append(tree.data)
    if tree.left_child:
        dfs(tree.left_child, acum)
    if tree.right_siblings:
        for sibling in tree.right_siblings:
            dfs(sibling, acum)
    return acum
// TODO: pprint ntree

print(dfs(tree))