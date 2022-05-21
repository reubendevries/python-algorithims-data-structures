def recursion(n):
    '''An example of basic recusion.'''
    while n > 1:
        print(n)
        recursion(n-1)
    return n

if __name__ == "__main__":
    x = recursion(10)
    print(x)
    