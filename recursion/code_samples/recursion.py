def simple_recursion_example(n):
    '''Example of Recursion Function - it will keep calling the function until it has
    reached 1'''
    if n==1:
        print("n is equal to 1")
    else:
        print(n)
        simple_recursion_example(n-1)

def power_of_two_recursive(n):
    '''An example of a recursive function that returns a number to the power of two.'''
    if n == 0:
        return 1
    else:
        power = power_of_two_recursive(n-1)
        return power * 2

def power_of_two_iterative(n):
    '''An example of an iterative function that returns a number to the power of two.'''
    i = 0
    power = 1
    while i < n:
        power = power * 2
        i = i + 1
    return power

def recursive_factorial_number(n):
    '''An example of a recursive factorial function written in python.'''
    if n < 0:
        return -1
    elif n < 2:
        return 1
    else:
        return (n * recursive_factorial_number(n-1))

def iterative_factorial_number(n):
    '''An example of iterative factorial function written in python'''
    if n < 0:
        return -1
    elif n < 2:
        return 1
    else:
        answer = 1
        for i in range(answer, n+1):
            answer *= i
        return answer

if __name__ == "__main__":
    recursive_answer = recursive_factorial_number(5)
    print(recursive_answer)
    iterative_answer = iterative_factorial_number(5)
    print(iterative_answer)
    simple_recursion_example(10)
    x = power_of_two_recursive(10)
    print(x)
    y = power_of_two_iterative(10)
    print(y)