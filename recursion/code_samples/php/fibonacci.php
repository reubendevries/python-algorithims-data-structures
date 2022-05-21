<?php

function fibonacciNumber($n) {
    if ($n == 0 || $n == 1) {
        return $n;
    }
    return (fibonacciNumber($n-2) + fibonacciNumber($n-1));
}

$x = fibonacci(10);
echo "$x\n";
?>