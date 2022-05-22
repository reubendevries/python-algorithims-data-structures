def factorial_number(n):
    '''An example of a recursive factorial function written in python.'''
    assert n >= 0 and int(n) == n, "Factorial Sums must be calaculated from a positive integer"
    if n <= 1:
        return 1
    return n * factorial_number(n-1)
    

if __name__ == "__main__":
    x = factorial_number(-10)
    print(x)
