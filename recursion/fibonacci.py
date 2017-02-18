def bad_fibo(n):
    if n <= 1:
        return n
    else:
        return bad_fibo(n-1) + bad_fibo(n-2)

def good_fibo(n):
    if n <= 1:
        return 0,1
    else:
        a,b = good_fibo(n-1)
        return a+b, a

print(good_fibo(10))