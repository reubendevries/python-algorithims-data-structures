def fibonacci_function(n):
    '''An example of a recursive fibonacci function written in python.'''
    try:
        assert n >= 0 and int(n) == n
        if n == 1 or n == 0:
            return n
        return fibonacci_function(n-2) + fibonacci_function(n-1)
    except AssertionError:
        print(f"{0} must be a positive interger.".format(n))
        return exit(1)

if __name__ == "__main__":
    x = fibonacci_function(10)
    print(x)