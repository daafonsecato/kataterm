#!/bin/bash
cd /exercise
# Check if the branches have diverged by one commit
if [[ $(git rev-list --count --left-right mybranch...master | tr -d '[:space:]') == "11" ]]; then
    # Check if file1.txt was added to mybranch
    if git diff --name-only mybranch~1 mybranch | grep -q "file1.txt"; then
        # Check if file2.txt was added to master branch
        if git diff --name-only master~1 master | grep -q "file2.txt"; then
            echo "success"
        else
            echo "file2.txt was not added to master branch"
        fi
    else
        echo "file1.txt was not added to mybranch"
    fi
else
    echo "Both branches should have diverged by one commit from the initial commit named 'dummy commit'"
fi
