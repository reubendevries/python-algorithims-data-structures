# Recursion

## What is Recursion?
A way of solving a problem by having a function call itself. A real life metaphor of recursion might look 
like a nested Russian Doll, where you get one Russian doll, but when you open up the Russian doll you find 
a smaller one embedded inside it, you can keep on finding a new smaller Russian doll, until the smallest one 
is unable to be opened. When you reach the smallest doll, you can reverse the process, closing each doll one 
by one in reverse order.

- Performing the same operation multiple times with different inputs
- In every step we try smaller inputs to make the problem smaller
- We need to define a base condition so we know when we need to stop the recursion process. 
    Otherwise we will create an infinite loop.

## Why do we need Recursion?

Recursive thinking is really important in programming when you need to break down bigger problems 
into smaller ones to help make the problem solving a bit more manageable. These are the reasons 
why you would want to learn about recursion algorithms.

- You can divide the problem into similar sub-problems.
- Prominate usage of recussion will be used when facing problems with both trees and graphs.
- Recursion is used in many other algorithms (i.e. Divide & Conquer, Greedy, and )

## Recursion Code Examples (Python, Go, Java, Javascript, PHP)?

### Python Recursion Example: The Bad Way

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

if __name__ == "__main__":
    firstMethod()

```

### Python Recursion Example: The Good Way

```python

list = ["first", "second", "third", "forth"]

def recursion_method_example(n, s):
    if s == "forth":
        print("I am the forth method")
    else:
        print(f"I am the {s} method")
        n = (n + 1)
        s = list[n]
        recursion_example(n, s)

if __name__ == "__main__":
    for i,v in range list:
        recursion_example(i,v)

```

### Go Recursion Example: The Bad Way

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

### Go Recursion Example: The Good Way

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

func recursionMethodExample(i int, s string) {
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

### Java Recursion Example: The Bad Way

```java
public class recursion {
    
    static void firstMethod() {
        System.out.println("I am the first method");
        secondMethod();
    }

    static void secondMethod() {
        System.out.println("I am the second method");
        thirdMethod();
    }

    static void thirdMethod() {
        System.out.println("I am the third method");
        fourthMethod();
    }

    static void forthMethod() {
        System.out.println("I am the forth method");
    }

    public static void main(String[] args) {
        firstMethod();
    }
}

```

### Java Recursion Example: The Good Way

```java
public class recursion {

    final static ArrayList<String> list  = new ArrayList<String>(4);

    static void recursionMethodExample(int i) {
        if (i >= list.size()-1) {
          System.out.printf("I am the %s method.", list.get(i));
          System.exit(0);
        } else {
          System.out.printf("I am the %s method.%n", list.get(i));
          i++;
          recursionMethodExample(i);
        }
    }

    public static void main(String[] args) {
        list.add("First");
        list.add("Second");
        list.add("Third");
        list.add("Forth");
        for (int i = 0; i < list.size(); i++) {
            recursionMethodExample(i);
        }
    }
}
```
### Javascript Recursion Example: The Bad Way

```javascript
Recursion Example: The Good Way
const firstMethod = function() {
    console.log("I am the first method");
    secondMethod();
}
const secondMethod = function() {
    console.log("I am the second method");
    thirdMethod();
}
const thirdMethod = function() {
    console.log("I am the third method");
    forthMethod();
}

const fourthMethod = function() {
    console.log("I am the fourth method");
}

firstmethod();

```

### Javascript Recursion Example: The Good Way

```javascript

const list = ["first", "second", "third", "fourth"];
    
const recursionMethodExample = function(i) {
    if (i >= list.length-1) {
        document.write("I am the %s method", list[i]);
    } else {
        document.write("I am the %s method", list[i]);
        i++
        recursionMethodExample(i);
    }
}

let i = 0;
recursionMethodExample(i);

```

### PHP Recursion Example: The Bad Way

```php
<?php
    function firstMethod() {
        echo "This is the first method";
        secondMethod();
    }

    function secondMethod() {
        echo "This is the second method";
        thirdMethod();
    }

    function thirdMethod() {
        echo "This is the third method";
        forthMethod();
    }

    function fourthMethod() {
        echo "This is the fourth method";
    }
?>
```

### PHP Recursion Example: The Good Way

```php
<?php
$i = 0;

function recursionMethodExample($i) {
    $array = array("first", "second", "third", "fourth");
    $arrayLength = count($array);
    while ($i < ($arrayLength-1)) {
        echo "I am the $array[$i] method\n";
        $i++;
        recursionMethodExample($i);
    }
    echo "I am the fourth method\n";
    exit;
}

recursionMethodExample($i);

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

### How to find the Fibonacci numbers using Recursion?

```python
```

```go
```

```java
```

```javascript
```

```php
```

###