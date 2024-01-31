#!/bin/bash
cd /exercise
# Validate git log
if git reflog | grep -q "Revert \"6\""; then
    if git log --oneline --graph --all | awk '{print $3}' | grep -q "7"; then
        if ! git log --oneline --graph --all | awk '{print $3}' | grep -q "10"; then
            if ! git log --oneline --graph --all | awk '{print $3}' | grep -q "9"; then
                if ! git log --oneline --graph --all | awk '{print $3}' | grep -q "8"; then
                    if git status | grep -qz "Untracked files:.*10.txt"; then
                        if git status | grep -qz "Untracked files:.*9.txt"; then
                            echo "success"
                        else
                            echo "Error: 9.txt is not untracked"
                        fi
                    else
                        echo "Error: 10.txt is not untracked"
                    fi
                else
                    echo "Error: Commit 8 is still present"
                fi
            else
                echo "Error: Commit 9 is still present"
            fi
        else
            echo "Error: Commit 10 is still present"
        fi
    else
        echo "Error: Commit 7 is not present"
    fi
else
    echo "Error: Revert \"6\" is not present"
fi