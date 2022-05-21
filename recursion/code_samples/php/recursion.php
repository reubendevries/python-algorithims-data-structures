<?php
function simpleRecursionExample($n) {
    if ($n == 1) {
        echo "$n\n";
    } else {
        echo "$n\n";
        simpleRecursionExample($n-1);
    }
}

function powerOfTwoRecursive($n) {
    if ($n == 0) {
        return 1;
    } else {
        $power = powerOfTwoRecursive($n - 1);
		return $power * 2;
    }
}

function powerOfTwoIterative($n) {
    $i = 0;
    $power = 1;
    while ($i < $n) {
        $power = $power * 2;
        $i = $i+ + 1;
    }
    return $power;
}

function factorialNumber($n) {
    if ($n <= 0) {
        return -1;
    } else if  ($n < 2) {
        return 1;
    } else {
        return ($n * factorialNumber($n-1));
    }
}

function fibonacciNumber($n) {
    if ($n == 0 || $n == 1) {
        return $n;
    } else {
        return (fibonacciNumber($n-2) + fibonacciNumber($n-1));
    }
}

simpleRecursionExample(10); // should count down from 10 to 1
$a = powerOfTwoRecursive(10);
echo "$a\n"; // should return 1024
$b = powerOfTwoIterative(10);
echo "$b\n"; // should return 1024
$c = factorialNumber(5);
echo "$c\n"; // should return 120
$d = fibonacciNumber(10);
echo "$d\n"; // should return 55
?>