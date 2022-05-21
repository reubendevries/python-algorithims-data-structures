<?php
function factorial($n){
    if  ($n < 2) {
        return 1;
    }
    return ($n * factorialNumber($n-1));
}

$x = factorial(10);
echo "$x\n";
?>