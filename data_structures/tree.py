class Tree:

    class Position:

        def element(self):
            """Return the element stored at this Position."""
            raise NotImplementedError

        def __eq__(self, other):
            """Return True if other Position represents the same location."""
            raise NotImplementedError

        def __ne__(self, other):
            """Return True if other Position not represents the same location ."""
            return not (self == other)

    def __len__(self):
        """Return the total number of elements in the tree."""
        raise NotImplementedError

    def root(self):
        """Return Position representing the tree's root (or None if empty)."""
        raise NotImplementedError

    def parent(self, p):
        """Return Position representing the p's parent (or None if p is root)."""
        raise NotImplementedError

    def num_children(self, p):
        """Return the number of children that Position p has."""
        raise NotImplementedError

    def children(self, p):
        """Generate an iteration of Positions representing p's children."""
        raise NotImplementedError

    def is_root(self, p):
        """Return True if Position p represents the root of tree."""
        return self.root() == p

    def is_leaf(self, p):
        """Return True if Position p does not have any children."""
        return self.num_children(p) == 0

    def is_empty(self):
        """Return True if the tree is empty."""
        return len(self) == 0

    def positions(self):
        """Generate an iteration of all positions stored within tree"""
        raise NotImplementedError

    def __iter__(self):
        """Generate an iteration of all elements stored within tree"""
        raise NotImplementedError

    def depth(self, p):
        """Return the number of levels separating p from the root"""
        if self.is_root(p):
            return 0
        else:
            return 1 + self.depth(self.parent(p))

    def _height(self, p):
        """Return the height of the subtree rooted at Position p."""
        if self.is_leaf(p):
            return 0
        else:
            return 1 + max(self._height(c) for c in self.children(p))

    def height(self, p=None):
        """Return the height of the subtree rooted at Position p.
        If p is None, return the height of the entire tree.
        """
        if p is None:
            p = self.root()
        return self._height(p)
