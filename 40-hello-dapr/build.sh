#!/bin/bash

set -e 
dir=$(pwd)
name=$(basename "$dir" | sed 's/^...//g')
ver=1.0

echo -e "\n\n ===> build ===> :\n\n\n"
docker build -t "${name}":${ver} .

echo -e "\n\n ===> show ===> :\n\n\n"
docker images "${name}"


exit 0
