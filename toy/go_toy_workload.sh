#!/bin/bash
set -e
trap 'error ${LINENO}' ERR

rm -rf workload_dir
mkdir workload_dir
echo -n "hello" > workload_dir/file1

rm -rf traces_dir
mkdir traces_dir

go build -o a.out toy.go

cd workload_dir

alice-record --workload_dir . \
	--traces_dir ../traces_dir \
	../a.out
