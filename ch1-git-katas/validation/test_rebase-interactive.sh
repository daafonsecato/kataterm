#!/bin/bash
cd /exercise

if git reflog | grep -q "rebase (finish):"; then
        if git reflog | grep -q "rebase (start):"; then
            echo "success"
        else
            echo "No rebase detected."
        fi
    else
        echo "Rebase on master was not finished."
fi