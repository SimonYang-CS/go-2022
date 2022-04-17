#!/bin/bash

dir=$(pwd)
name=$(basename "$dir")
ver=1.0

echo -e "\n\n ===> build ===> :\n\n\n"
docker build -t "${name}":${ver} .

echo -e "\n\n ===> show ===> :\n\n\n"
docker images "${name}"

echo -e "\n\n ===> run ===> :\n\n\n"
docker run --rm "${name}":${ver}

exit 0
