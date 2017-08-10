#!/bin/bash

set -e
trap 'error ${LINENO}' ERR

rm -rf workload_dir
mkdir workload_dir

rm -rf traces_dir
mkdir traces_dir

rm -rf bin
mkdir bin

go build -o bin/workload workload.go

cd workload_dir

alice-record --workload_dir . \
	--traces_dir ../traces_dir \
	../bin/workload

cd ..
go build -o bin/checker checker.go
alice-check --traces_dir=traces_dir --checker=./bin/checker
