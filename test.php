<?php

use Chibi\Text\Keyword;

$dir = "./seowords/files/";

$files = scandir($dir);

foreach ($files as $idx => $file) {
    $res = [];
    if ($file == "." || $file == "..") {
        continue;
    }

    $content = file_get_contents($dir . "/". $file);
    $words = explode("\n", $content);
    print "$word\n";
    foreach ($words as $word) {
        if (Keyword::hasBanWord(Keyword::TYPE_POLICE, $word)) {
            print "$word\n";
            $res[] = $word;
        }
    }
    $cur = $idx - 2;
    file_put_contents("./seowords/res/php_res$cur.csv", implode("\n", $res), FILE_APPEND);
    print "$file over.\n";
}


$cur = $idx - 2;
file_put_contents("./seowords/res/php_res$cur.csv", implode("\n", $res), FILE_APPEND);
print "$file over.\n";
