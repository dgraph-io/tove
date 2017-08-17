#!/bin/bash

if [ "$#" -ne 1 ]; then
	echo "Usage: $0 <base_dir> (e.g. /dev/shm/alice-22936)"
	exit 1
fi

for dir in $(find $1 -type d); do
	echo $dir
	bin/checker $dir ${dir}.input_stdout
done
