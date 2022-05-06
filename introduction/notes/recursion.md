# Recursion

## What is Recursion?
A way of solving a problem by having a function call itself. A real life metaphor of recursion might look like a nested Russian Doll, where you get one Russian doll, but when you open up the Russian doll you find a smaller one embedded inside it, you can keep on finding a new smaller Russian doll, until the smallest one is unable to be opened. When you reach the smallest doll, you can reverse the process, closing each doll one by one in reverse order.

- Performing the same operation multiple times with different inputs
- In every step we try smaller inputs to make the problem smaller
- We need to define a base condition so we know when we need to stop the recursion process. Otherwise we will create an infinite loop.

## Why do we need Recursion?
Recursive thinking is really important in programming when you need to break down bigger problems into smaller ones to help make the problem solving a bit more manageable. These are the reasons why you would want to learn about recursion algorithms.

- You can divide the problem into similar sub-problems.
- Prominate usage of recussion will be used when facing problems with both trees and graphs.
- Recursion is used in many other algorithms (i.e. Divide & Conquer, Greedy, and )

## How does Recursion work?
### Bad Way
```python
def firstMethod():
    secondMethod()
    print("I am the first method")
def secondMethod():
    thirdMethod()
    print("I am the second method")
def thirdMethod():
    fourthMethod()
    print("I am the third method")
def fourthMethod():
    print("I am the fourth method")
```
### Good Way
```python
    def method(str):
        print("I am the %s method")
    list = ["first", "second","third","forth"]
    for i in list:
        method(i)
```
## When to use Recursion?

## When to avoid using Recursion?

## Challenge
How to find the Fibonacci numbers using Recursion?
```python
```
