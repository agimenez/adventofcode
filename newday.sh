#!/bin/bash

dir=$1
mkdir -p ${dir}
cp main.go ${dir}
cd ${dir}
touch README.md input.txt test.txt

echo "cd ${dir}"
