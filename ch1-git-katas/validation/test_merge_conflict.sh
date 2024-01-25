#!/bin/bash
cd /exercise
# Check if merge-conflict-branch1 was merged into master
if git branch --merged | grep -q "merge-conflict-branch1"; then
    # Check if there are any unresolved conflicts
    if git diff --check | grep -q "conflict"; then
        echo "Error: There are unresolved merge conflicts."
    else
        echo "Success: Merge was completed successfully."
    fi
else
    echo "Error: merge-conflict-branch1 was not merged into master."
fi
