def factorial_number(n):
    '''An example of a recursive factorial function written in python.'''
    try:
        assert n >= 0 and int(n) == n
        if n == 1:
            return 1
        return n * factorial_number(n-1)
    except AssertionError:
        print (f"Factorial Sum must be a positive interger, not {n}.")
        exit(1)

if __name__ == "__main__":
    x = factorial_number(-10)
    print(x)
