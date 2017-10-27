#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys

if __name__ == "__main__":
    DIR = os.path.abspath(os.path.dirname(sys.argv[0]))
    resdir = DIR + "/res/"

    cont = ""
    
    for fileName in os.listdir(resdir):
        with open(resdir + fileName, 'r') as f:
            cont += f.read() + "\n"

    resfile = DIR + "/result.csv"
    with open(resfile, 'w') as f:
        f.write(cont)
