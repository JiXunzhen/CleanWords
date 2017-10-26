#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import smart_open
import os
import sys

if __name__ == "__main__":
    DIR = os.path.abspath(os.path.dirname(sys.argv[0]))
    size = 1000000
    count = 0
    total = 0
    index = 0
    curfile = open((DIR + '/files/words%d.csv' % index), 'w')
    for line in smart_open.smart_open('./juhelistingcache.csv'):
        row = str(line, encoding='utf8').strip('\n')
        words = row.split('&')

        # 统一数据格式，将kwd放到最前面
        if len(words) == 3:
            words[0], words[2] = words[2], words[0]
            row  = "&".join(words)

        count += 1
        total += 1
        if count == size:
            index += 1
            count = 0
            curfile.close()
            curfile = open(((DIR + '/files/words%d.csv') % index), 'w')
            print("file %d over." % index)

        curfile.write(row + "\n")

    curfile.close()

    print(index)
    print(total)
