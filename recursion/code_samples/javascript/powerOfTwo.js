const powerOfTwoRecursive = function(n) {
    if (n == 0) {
        return 1;
    } 
    power = powerOfTwoRecursive(n - 1);
	return power * 2;
}

const x = powerOfTwoRecursive(10);
console.log(x);