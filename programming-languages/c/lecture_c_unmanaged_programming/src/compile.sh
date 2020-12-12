#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "[ERROR] Illegal number of parameters"
    exit 1
fi

echo "[INFO] Compile $1 to $2"
clang -std=c89 -W -Wall -o $2 -pedantic-errors $1
