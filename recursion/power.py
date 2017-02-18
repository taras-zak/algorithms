import time

def pow_(x, n):
    """Calculate power n of x for O(n)"""
    if n == 0:
        return 1
    else:
        return x * pow_(x, n-1)

def pow2(x,n):
    """Calculate power of x for O(log(x))"""
    if n == 0:
        return 1
    else:
        partial = pow2(x, n//2)
        result = partial*partial
        if n%2 == 1:
            result *= x
        return result
        
start = time.time()
res1 = pow_(2, 995)
cp = time.time()
res2 = pow2(2, 995)
print(res1)
print(res2)
print(cp - start)
print(time.time() - cp)
