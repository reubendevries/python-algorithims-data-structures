def power_of_two(n):
    '''An example of a recursive function that returns a number to the power of two.'''
    if n == 0:
        return 1
    else:
        power = power_of_two(n-1)
        return power * 2

if __name__ == "__main__":
    x = power_of_two(10)
    print(x)