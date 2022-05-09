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

```
### Bad Way
```go
package main

func main() {
	firstMethod()
}

func firstMethod() {
    fmt.Println("I am the first method")
	secondMethod()
}
func secondMethod() {
    fmt.Println("I am the second method")
	thirdMethod()
}
func thirdMethod() {
    fmt.Println("I am the third method")
	forthMethod()
}
func forthMethod() {
    fmt.Println("I am the forth method")
}
```
### Good Way
```go
package main

import (
	"fmt"
	"os"
)

var list = [...]string{"first", "second", "third", "forth"}

func main() {
	for i, v := range list {
		simpleRecursionExample(i, v)
	}
}

func simpleRecursionExample(i int, s string) {
	if i >= (len(list) - 1) {
		fmt.Println("I'm the forth method")
		os.Exit(0)
	} else {
		fmt.Printf("I'm the %s method\n", s)
		i++
		s = list[i]
		simpleRecursionExample(i, s)
	}
}
```

```java
public class recursion {
  static void simpleRecursionExample(int n) {
    if (n == 1) {
        System.out.println(n);
    } else {
        System.out.println(n);
        simpleRecursionExample(n-1);
    }
  }

  public static void Recursion(int[] args) {
    simpleRecursionExample(10);
  }
}
```

```javascript
function simpleRecursionExample(n) {
    if (n == 1) {
        alert(n)
    } else {
        alert(n)
        recursionExample(n-1)
    }
}

simpleRecursionExample(10)
```

```php
<?php
function simpleRecursionExample($n) {
    if ($n == 1) {
        echo "$n\n";
    } else {
        echo "$n\n";
        simpleRecursionExample($n-1);
    }
}

simpleRecursionExample(10);

?>
```
## When to use Recursion?
- When we can break the big probelm into smaller easier to solve sub-problems.
- When we can accept the extra overhead (both time and space) that comes along with recursion.
- When we need a quick solution instead of an efficient solution.
- When we need to travese a tree graph.
- When we use memorization to solve a problem.
## When to avoid using Recursion?
- When we can't break the problem into small sub-problems.
- When time and space complexity actually matter to solving the problem.
- When we require a low-memory consumption solution.
- When we require a time-sensitive solution.
## Challenge
How to find the Fibonacci numbers using Recursion?
```python
```
