const fibonacci = function(n) {
    if (n == 0 || n == 1) {
        return n;
    } 
    return (fibonacci(n-2) + fibonacci(n-1));
}

const x = fibonacciFunction(10);
console.log(x);