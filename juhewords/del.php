<?php

$cont = file_get_contents("./result.csv");

$rows = explode("\n", $cont);


$cnt = 0;
$total = 0;
$idx = 0;

foreach ($rows as $row) {
    $node = \MJuheListingCache::fetchItemBySid($row);
    if ($node) {
        $node->delete();
    }


    $cnt++;
    $total++;
    if ($cnt == 10000) {
        $cnt=0;
        $idx++;
        echo "{$idx}0000 over.";
    }

    usleep(10 * 1000);
}

echo "over.\n";
