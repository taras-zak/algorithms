def pal(n):
    res = set()
    for i in range(100,1000):
        for j in range(100,1000):
            t = str(i*j)
            if len(t) == 6 and t == t[::-1] and int(t) <= n:
                res.add(t)
    return max(res)

def pal(n):
    res = set()
    for i in range(100,1000):
        for j in range(100,1000):
            t = str(i*j)
            if len(t) == 6 and t == t[::-1]:
                res.add(t)
    return res

def prepare_palindroms():
    return sorted(pal(999999))

def get_max_pal(n, pals):
    for p in pals[::-1]:
        if int(p) <= n:
            return p


pals = prepare_palindroms()
print(pals)
print([get_max_pal(int(i), pals) for i in pals])