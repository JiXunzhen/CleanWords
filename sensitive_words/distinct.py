#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys

if __name__ == "__main__":
    DIR = os.path.abspath(os.path.dirname(sys.argv[0]))
    worddir = DIR  + "/files"

    wordList = []
    for fileName in os.listdir(worddir):
        if fileName.endswith(".csv"):
            with open(DIR + "/files/" + fileName, 'r') as f:
                cont = f.read()

                rows = cont.split('\n')
                for row in rows:
                    word = row.split(',')[0]
                    if word != '':
                        wordList.append(row.split(',')[0])

    distinct = []
    for i in range(0, len(wordList)):
        distinct.append(1)

    for idx, word in enumerate(wordList):
        for idx2, word2 in enumerate(wordList):
            if idx != idx2 and distinct[idx] == 1 and distinct[idx2] == 1:
                if len(word) < len(word2) and word in word2:
                    print("%s %s" % (word, word2))
                    distinct[idx2] = 0
                elif word2 in word:
                    print("%s %s" % (word2, word))
                    distinct[idx] = 0

    # 写到文件中
    with open(DIR + "/word.csv", 'w') as f:
        for idx, d in enumerate(distinct):
            if d == 1:
                f.write(wordList[idx] +'\n')

