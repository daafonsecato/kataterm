#!/bin/bash

expected_output=".:
README.txt
src

./src:
myapp.c
mylib.c
myapp.h"

actual_output=$(cd /exercise; ls -R .)

echo "$actual_output"
echo "$expected_output"
if [[ "$actual_output" == "$expected_output" ]]; then
    echo "success"
else
    echo "Denied" 
fi
