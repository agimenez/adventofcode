#!/bin/bash

test -z ${AOC_SESSION} && {
	echo "Need to set AOC_SESSION environment variable"
	exit 1
}

year=$(date +%Y)
day=$(date +%d)
dir="${year}/${day}"

if [[ -n "$1" ]]; then
	dir=$1
fi

year=${dir%%/*}
day=${dir##*/}
day=${day#0}

mkdir -p ${dir}
test -f ${dir}/main.go || cp main.go ${dir}
test -f ${dir}/main_test.go || cp main_test.go ${dir}
test -f ${dir}/Justfile || cp Justfile ${dir}
cd ${dir}

curl https://adventofcode.com/${year}/day/${day}/input \
	--silent --cookie "session=${AOC_SESSION}" \
	--output input.txt

cat <<EOF > README.md
## $year - Day ${day}

\`\`\`
curl https://adventofcode.com/${year}/day/${day}
\`\`\`
EOF

touch test.txt
