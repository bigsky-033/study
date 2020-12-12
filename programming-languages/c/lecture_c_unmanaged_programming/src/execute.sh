#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "[ERROR] Illegal number of parameters"
    exit 1
fi

out=$(date +%s)

./compile.sh $1 ${out}

exit_code=$?
if [ $exit_code -eq 1 ]; then
    echo "[ERROR] Error happens during compile $1"
else
    echo "[INFO] Execute $1"
    echo ""

    ./${out}

    echo ""
    echo "[INFO] Remove execute file for $1"
    rm ${out}
fi
