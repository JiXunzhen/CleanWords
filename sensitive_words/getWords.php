<?php

$dir = dirname(__FILE__);
use Chibi\Keyword\StoreManager;
use Chibi\Text\Keyword;

$words = StoreManager::getWords(['tags' => ['公安', '全站各入口', '屏蔽']]);

$res = [];
foreach ($words['keyword'] as $word => $type) {
    if (Keyword::hasBanWord(Keyword::TYPE_POLICE, $word)) {
        echo "$word\n";
        $res[] = $word;
    }
}

file_put_contents($dir . "/files/chibi.csv", implode("\n", $res));
