# Algorithms
___

#### Notes

- The point of alogrithms is to understand the different methods used to attack a specific coding problem. 
    These are following different types of alogrithms we will use to attack different coding problems given 
    to us.

## Types of Algorithms
```mermaid
flowchart TB
ta["Different Types of Alogrithms"]
ta-->rs[Simple Recursive Alogrithms]
ta-->dca[Divide & Conquer Alogrithms]
ta-->dpa[Dynamic Programming Alogrithms]
ta-->ga[Greedy Alogrithms]
ta-->bfa[Brute Force Algorithms]
ta-->ra[Randomized Alogrithms]
```
### Simple Recursive Algorithm
___
#### Notes

- Recursion Algorithm is a function that calls itself with smaller or simpler input values.
- It's a problem that can be solved by solving smaller instances of the same problem unless 
    that problem is so small that we can just solve it directly.

**Example of a simple Recursive Algorithm**

```
Algorithm Sum(A, n)
    if n = 1
        return A[0]
    s = Sum[A, n-1] /* recurse on all but last */
    s = s + A[n-1] /* add last element */
return s
```

### Divide & Conquer Algorithm
___
#### Notes

- We should divide the problem into small sub-problems of the same type, then then solve these 
    problems recursively.
- We then should combine the solutions of the sub-problems into a final solution to the original problem.

**Example of Divide & Conquer Algorithm**

```

```
### Dynamic Programming Algorithm
___
#### Notes

- Dynanmic Programming Algorithm works based on memorization. It will then use whatever it's memorized 
    to find the best solution among multiple solutions.

**Example of Dynamic Programing Algorithm**

```
```
### Greedy Algorithm
___
#### Notes

- Greedy algorithms will take the best answer we can without worrying about future consequences.
- We hope when using a greedy algorithm that by choosing a local optimum soluion at each step, we
    will end up at a global optimum solution.

**Example of Greedy Algorithm**

```
```
### Brute Force Algorithm
___
#### Notes

- Brute force Alogrithim will use all possiblities to until a satisfacory solution is found.

**Example of Brute Force Algorithm**

```
```
### Randomized Algorithm
___
#### Notes

- Randomized Algorithms will use a random number at least once during the computation to make a decision.

**Example of Randomized Algorithm**

```
```