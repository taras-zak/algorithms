def unique(seq):
    """check sequence for unique for O(n**2)"""
    for i, item in enumerate(seq):
        for item_compare in seq[i+1:]:
            if item == item_compare:
                return False
    return True


def unique3(seq, start, stop):
    """check sequence for unique for O(2**n). inefficient example """
    if stop - start <= 1: return True
    elif not unique3(seq, start, stop-1): return False
    elif not unique3(seq, start+1, stop): return False
    else: return seq[start] != seq[stop-1]

seq = 'sdf'
print(unique3(seq,0,len(seq)))
