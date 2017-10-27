<?php

$content = file_get_contents("./result.csv");

$rows = explode("\n", $content);

$cnt = 0;

$file = fopen("./url.csv", 'w');

$cnt = 0;
$idx = 0;
foreach ($rows as $row) {
    $info = explode("&", $row);
    if (preg_match("/ObjectId\(([0-9a-zA-Z]+)\)/", $info[1], $m)) {
        $id = $m[1];
        $juheword = MJuheKeyword::fetchItemBySid($info[0]);

        if ($juheword) {
            foreach ($juheword->city as $city) {
                $urlArr = "http://{$city}.baixing.com/" . urlencode("jh_{$info[0]}") . '/';
                fwrite($file, $urlArr . "\n");
            }
        }
    }

    $cnt++;
    if ($cnt == 1000) {
        $idx++;
        $cnt = 0;
        echo "{$idx}000 over.";
    }

    usleep(100 * 1000);
}

fclose($file);
