#!/bin/bash
cd /exercise
# Check if cherry-pick of Commit F was performed
if git reflog | grep -q "cherry-pick: Commit F:"; then
    # Check if cherry-pick of Commit G was performed
    if git reflog | grep -q "cherry-pick: Commit G:"; then
        # Check if reset was performed
        if git reflog | grep -q "reset: moving to HEAD~1"; then
            echo "success"
        else
            echo "Error: Reset operation not found"
        fi
    else
        echo "Error: Cherry-pick of Commit G not found"
    fi
else
    echo "Error: Cherry-pick of Commit F not found"
fi
