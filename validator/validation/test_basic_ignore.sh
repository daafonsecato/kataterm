#!/bin/bash
cd /exercise
# Check .gitignore file
if grep -q "*.s" .gitignore && grep -q "*.txt" .gitignore && grep -q "!file3.txt" .gitignore; then
    if [ -e foo.s ] && ! git ls-files --error-unmatch foo.s >/dev/null 2>&1; then
        if [ -e file2.txt ] && ! git ls-files --error-unmatch file2.txt >/dev/null 2>&1; then
            if git status | grep -q "deleted:    file1.txt"; then
                echo "success"
            else
                echo "file1.txt is not deleted"
            fi
        else
            echo "file2.txt is not untracked"
        fi
    else
        echo "file1.s is not untracked"
    fi
else
    echo ".gitignore contents are incorrect"
fi
