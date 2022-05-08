<?php
function simpleRecursionExample($n) {
    if ($n == 1) {
        echo "$n\n";
    } else {
        echo "$n\n";
        simpleRecursionExample($n-1);
    }
}

simpleRecursionExample(10);

?>