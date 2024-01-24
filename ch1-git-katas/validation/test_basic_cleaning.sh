#!/bin/bash

expected_output=".:
README.txt
src

./src:
myapp.c
mylib.c
myapp.h" 

actual_output=$(cd /exercise; ls -R . | sort)

sactualoutput=$(echo "$actual_output" | sort)

if [[ "$sactualoutput" == "$actual_output" ]]; then
    echo "success"
else
    echo "Denied" 
fi
