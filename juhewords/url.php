<?php


use Listing\RefashionJuheListing;

$dir = dirname(__FILE__);

$content = file_get_contents($dir . "/result.csv");

$rows = explode("\n", $content);

$cnt = 0;
$idx = 0;
$file = fopen($dir . "/urls/url0.txt", 'w');

foreach ($rows as $row) {
    $info = explode("&", $row);

    if (count($info) == 3) {
        $url = "http://{$info[2]}.baixing.com/{$info[1]}/" . RefashionJuheListing::SYMBOL . RefashionJuheListing::encode($info[0]) . '/';

        fwrite($file, "$url\n");
        $cnt++;
        if ($cnt == 48000) {
            $cnt = 0;
            fclose($file);
            echo "$idx * 48000 over\n";
            $idx++;

            $file = fopen($dir . "/urls/url{$idx}.txt", 'w');
        }
    }
}

fclose($file);
