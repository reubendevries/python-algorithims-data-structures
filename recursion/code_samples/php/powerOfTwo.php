<?php

function powerOfTwo($n) {
    if ($n == 0) {
        return 1;
    } 
    $power = powerOfTwo($n - 1);
	return $power * 2;
}

$x = powerOfTwo(10);
echo "$x\n";

?>