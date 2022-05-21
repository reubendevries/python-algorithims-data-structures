const factorialNumber = function(n) {
    if (n < 2){
        return 1;
    } 
    return (n * factorialNumber(n-1));
}

const x = factorialNumber(10);
console.log(x);