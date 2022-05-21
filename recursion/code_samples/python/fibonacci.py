def fibonacci_function(n):
    '''An example of a recursive fibonacci function written in python.'''
    if n == 1 or n == 0:
        return n
    return fibonacci_function(n-2) + fibonacci_function(n-1)

if __name__ == "__main__":
    x = fibonacci_function(10)
    print(x)