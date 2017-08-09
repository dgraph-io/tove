#!/bin/bash
#./go_toy_workload.sh && alice-check --traces_dir=traces_dir --checker=./toy_checker.py

set -e
trap 'error ${LINENO}' ERR

rm -rf workload_dir
mkdir workload_dir
echo -n "hello" > workload_dir/file1

rm -rf traces_dir
mkdir traces_dir

rm -rf bin
mkdir bin

go build -o bin/toy toy.go

cd workload_dir

alice-record --workload_dir . \
	--traces_dir ../traces_dir \
	../bin/toy

cd ..
go build -o bin/checker checker.go
alice-check --traces_dir=traces_dir --checker=./bin/checker
