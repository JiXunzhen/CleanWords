<?php


use Listing\RefashionJuheListing;

$content = file_get_contents("./result.csv");

$rows = explode("\n", $content);

$count = 0;

$file = fopen("./url.csv", 'w');

foreach ($rows as $row) {
    $info = explode("&", $row);

    if (count($info) == 3) {
        $url = "http://{$info[2]}.baixing.com/{$info[1]}/" . RefashionJuheListing::SYMBOL . RefashionJuheListing::encode($info[0]) . '/';

        fprintf($file, "$url\n");
    }
}

fclose($file);
