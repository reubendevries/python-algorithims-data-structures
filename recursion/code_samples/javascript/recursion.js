
const simpleRecursionExample = function(n) {
    if (n == 1) {
        console.log(n);
    } else {
        console.log(n)
        simpleRecursionExample(n-1);
    }
}
const powerOfTwoRecursive = function(n) {
    if (n == 0) {
        return 1;
    } else {
        power = powerOfTwoRecursive(n - 1);
		return power * 2;
    }
}
const powerOfTwoIterative = function(n) {
    i = 0;
    power = 1;
    while (i < n) {
        power = (power * 2);
        i = i + 1;
    }
    return power;
}

const factorialNumber = function(n) {
    if (n <= 0){
        return -1;
    } else if (n < 2){
        return 1;
    } else{
        return (n * factorialNumber(n-1));
    }
}

const fibonacciFunction = function(n) {
    if (n == 0 || n == 1) {
        return n;
    } else {
      return (fibonacciFunction(n-2) + fibonacciFunction(n-1));
    }
}

simpleRecursionExample(10);
const a = powerOfTwoRecursive(10);
console.log(a);
const b = powerOfTwoIterative(10);
console.log(b);
const c = factorialNumber(5);
console.log(c);
const d = fibonacciFunction(10);
console.log(d);