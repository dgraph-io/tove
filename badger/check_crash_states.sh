#!/bin/bash

# Runs the checker script over set of crash states. Usage: set `base` below to
# the directory that contains the various crash states. The checker will run
# for each. Look at the output to see which crash states cause problems with
# badger starting up (they will have output). The checker script is silent when
# there are no problems.

base=/dev/shm/alice-12692

for dir in $base/*/
do
	echo $dir
	bin/checker $dir <(echo dummy stdout)
done
