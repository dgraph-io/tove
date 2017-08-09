#!/usr/bin/env python2
import os
import sys

crashed_state_directory = sys.argv[1]
os.chdir(crashed_state_directory)
assert open('file1').read() in ['hello', 'world']
