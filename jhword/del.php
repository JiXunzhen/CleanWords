<?php

$cont = file_get_contents("./result.csv");

$rows = explode("\n", $cont);

$cnt = 0;
$total = 0;
$idx = 0;

foreach ($rows as $row) {
    $kwd = explode("&", $row)[0];
    $node = \MJuheKeyword::fetchItemBySid($kwd);
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

    usleep(1000);
}

echo "over.\n";
