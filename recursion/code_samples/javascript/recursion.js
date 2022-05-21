const recursion = function(n) {
    if (n == 1) {
        return n;
    } 
    console.log(n);
    recursion(n-1);
}

const x = recursion(10);
console.log(x);