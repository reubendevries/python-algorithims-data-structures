
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

simpleRecursionExample(10);
const x = powerOfTwoRecursive(10);
console.log(x);
const y = powerOfTwoIterative(10);
console.log(y);
