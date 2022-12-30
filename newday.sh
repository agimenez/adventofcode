#!/bin/bash

test -z ${AOC_SESSION} && {
	echo "Need to set AOC_SESSION environment variable"
	exit 1
}

dir=$1
year=${dir%%/*}
day=${dir##*/}
day=${day#0}

mkdir -p ${dir}
cp main.go main_test.go ${dir}
cd ${dir}

curl https://adventofcode.com/${year}/day/${day}/input \
	--silent --cookie "session=${AOC_SESSION}" \
	--output input.txt

curl https://adventofcode.com/${year}/day/${day} \
	--silent --cookie "session=${AOC_SESSION}" \
	| html2markdown --mark-code --asterisk-emphasis > README.md

touch test.txt
