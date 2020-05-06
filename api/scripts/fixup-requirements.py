#!/usr/bin/env python

import os
import sys

def fixup_target(target):
    if 'git+' in target:
        return target

    if '../utils' in target:
        return 'notifications-utils'


if __name__ == '__main__':
    for line in sys.stdin:
        if not line.startswith('-e'):
            sys.stdout.write(line)
            continue

        target = line[2:].strip()
        replaced_target = fixup_target(target)
        if replaced_target:
            sys.stdout.write('%s\n' % replaced_target)

