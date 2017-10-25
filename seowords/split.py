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
    for line in smart_open.smart_open('./seowords.csv'):
        row = str(line, encoding='utf8')
        word = row.split('&')[0]
        count += 1
        total += 1
        if count == size:
            index += 1
            count = 0
            curfile.close()
            curfile = open(((DIR + '/files/words%d.csv') % index), 'w')

        curfile.write(word + "\n")

    curfile.close()

    print(index)
    print(total)
