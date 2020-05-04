#!/usr/bin/env python

import os
import sys


if __name__ == '__main__':
    for line in sys.stdin:
        if not line.startswith('-e'):
            sys.stdout.write(line)
            continue

        target = line[2:].strip()
        sys.stdout.write('%s\n' % target)
