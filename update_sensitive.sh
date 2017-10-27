#!/bin/sh

haojing ./sensitive_words/getWords.php
python3 ./sensitive_words/distinct.py
