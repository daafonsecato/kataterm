#!/bin/bash
cd /exercise

if git reflog | grep -q "Merge branch 'master' into incremental-change-merge"; then
    if git reflog | grep -q "rebase (finish): returning to refs/heads/incremental-change-rebase"; then
        if git reflog | grep -q "rebase (continue): change 3" && git reflog | grep -q "rebase (continue): change 2" && git reflog | grep -q "rebase (continue): change 1"; then
            echo "success"
        else
            echo "Merge conflicts or rebase conflicts were not resolved."
        fi
    else
        echo "Rebase from master into incremental-change-rebase not detected."
    fi
else
    echo "Merge from master into incremental-change-merge not detected."
fi
