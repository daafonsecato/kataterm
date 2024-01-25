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
if [ "$branch" != "master" ]; then
    echo "Branch is not master"
else
    if [[ "$sactualoutput" == "$actual_output" ]]; then
        src_files=$(echo "$actual_output" | grep -v "^\./src" | grep -v "^src" | grep -v "^\.:" | grep -v "^\README.txt" | grep -v "^\myapp.c" | grep -v "^\mylib.c" | grep -v "^\myapp.h")
        if [[ -z "$src_files" ]]; then
            echo "success"
        else
            echo "You have unexpected files in the src folder"
        fi
    else
        echo "You dont have the expected files in your repository." 
    fi
fi