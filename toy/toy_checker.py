#!/usr/bin/env python2
import os
import sys

os.chdir(sys.argv[1])
assert open('file1').read() in ['hello', 'world']
