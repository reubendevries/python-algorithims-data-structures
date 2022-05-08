def simple_recursion_example(n):
    '''Example of Recursion Function - it will keep calling the function until it has
    reached 1'''
    if n==1:
        print("n is equal to 1")
    else:
        print(n)
        simple_recursion_example(n-1)

if __name__ == "__main__":
    simple_recursion_example(10)
