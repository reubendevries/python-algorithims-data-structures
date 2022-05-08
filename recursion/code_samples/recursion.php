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
        return (powerOfTwoRecursive($n-1) * 2);
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

simpleRecursionExample(10);
powerOfTwoRecursive(10);
powerOfTwoIterative(10);
?>