#!/bin/bash
cd /exercise
# Check if greeting.txt exists
if [ -f "greeting.txt" ]; then
    # Check if greeting.txt contains your favorite greeting
    favorite_greeting="Hello, world!"
    if grep -q "$favorite_greeting" "greeting.txt"; then
        # Check if README.md exists
        if [ -f "README.md" ]; then
            if git reflog | grep -q "merge greeting: Merge made by the 'ort' strategy"; then
                echo "success"
            else
                echo "Merge from greeting into master not detected."
            fi
        else
            echo "README.md does not exist"
        fi
    else
        echo "greeting.txt does not contain your favorite greeting Hello, World!"
    fi
else
    echo "greeting.txt does not exist"
fi
