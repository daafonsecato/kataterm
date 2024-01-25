#!/bin/bash
cd /exercise
# Check if all commits have a different author
if ! git log --format="%ae" | grep -q "git-katas@workplace.com"; then
    # Check if master and origin/master are in sync
    if git diff --quiet master origin/master; then
        # Check if local master and remote master have the same commit hash
        local_commit=$(git rev-parse master)
        remote_commit=$(git rev-parse origin/master)
        if [ "$local_commit" = "$remote_commit" ]; then
            echo "success"
        else
            echo "Error: Local master and remote master have different commit hashes"
        fi
    else
        echo "Error: master and origin/master are not in sync"
    fi
else
    echo "Error: Some commits have the same author as git-katas working bot"
fi



