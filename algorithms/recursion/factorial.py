import math

def factorial(n):
    if n == 1:
        return 1
    else:
        return n*factorial(n-1)


if __name__ == '__main__':
    n = 50
    print(factorial(n))
    print(math.factorial(n))