<?php

ini_set('memory_limit', '2G');

$content = file_get_contents("./kwd-city.csv");
$rows = explode("\n", $content);

$file = fopen("./urls/url0.txt", 'w');
$cnt = 0;
$idx = 0;
$total = 0;

foreach ($rows as $row) {
    $info = explode("\t", $row);

    if (count($info) == 2) {
        $kwd = $info[0];
        $citys = explode(",", $info[1]);
        foreach ($citys as $city) {
            $urlArr = "http://{$city}.baixing.com/" . urlencode("jh_{$kwd}") . '/';
            fwrite($file, $urlArr . "\n");
            $cnt++;
            $total++;

            if ($cnt == 48000) {
                $cnt = 0;
                $idx++;
                echo "{$idx} * 48000 over.\n";
                fclose($file);
                $file = fopen("./urls/url{$idx}.txt", 'w');
            }
        }
    }
}

fclose($file);
