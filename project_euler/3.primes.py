def is_prime(a):
    return not (a < 2 or any(a % x == 0 for x in range(3, int(a ** 0.5) + 1)))

def fak(n):
    f = []
    divider = 2
    while n > 1:
        while n % divider == 0:
            f.append(divider)
            n /= divider
        divider += 1
    return f

t = fak(600851475143)
print(t)
