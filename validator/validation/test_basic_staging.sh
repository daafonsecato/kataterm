#!/bin/bash
cd /exercise
# Check if the branch is master
branch=$(git rev-parse --abbrev-ref HEAD)
if [ "$branch" != "master" ]; then
    echo "Branch is not master"
else
    # Check if there are changes to be committed
    status=$(git status --porcelain)
    if [ -z "$status" ]; then
        echo "No changes to be committed"
    else
        # Check if file.txt is modified
        if ! git diff --quiet --exit-code file.txt; then
            echo "file.txt is not staged"
        else
            # Check if the content of file.txt is 3
            content=$(cat file.txt)
            if [ "$content" != "3" ]; then
                echo "Content of file.txt is not 3"
            else
                echo "success"
            fi
        fi
    fi
fi
