<?php
function recursion($n) {
    if ($n == 1) {
        return $n;
    }
    echo "$n\n";
    recursion($n-1);
}



$x = recursion(10); // should count down from 10 to 1
echo "$x\n";
?>