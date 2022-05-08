def simple_recursion_example(n):
    '''Example of Recursion Function - it will keep calling the function until it has
    reached 1'''
    if n==1:
        print("n is equal to 1")
    else:
        print(n)
        simple_recursion_example(n-1)

def power_of_two_recursive(n):
    '''Example of a recursive function that returns a number to the power of two.'''
    if n == 0:
        return 1
    else:
        power = power_of_two_recursive(n-1)
    return power * 2

def power_of_two_iterative(n):
    '''Example of an iterative function that returns a number to the power of two.'''
    i = 0
    power = 1
    while i < n:
        power = power * 2
        i = i + 1
    return power

if __name__ == "__main__":
    #simple_recursion_example(10)
    x = power_of_two_recursive(10)
    print(x)
    y = power_of_two_iterative(10)
    print(y)
