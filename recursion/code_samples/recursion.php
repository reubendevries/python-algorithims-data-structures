<?php
function recursion($n) {
    if ($n == 1) {
        echo "$n\n";
    } else {
        echo "$n\n";
        recursion($n-1);
    }
}

recursion(10);

?>