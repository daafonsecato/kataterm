#!/bin/bash
cd /exercise

# Check if the repository has 3 commits
commit_count=$(git rev-list --count HEAD)
if [ $commit_count -eq 3 ]; then
    # Check if the branch feature/uppercase doesn't exist
    if ! git show-ref --quiet refs/heads/feature/uppercase 2>/dev/null; then
        # Check if there has been a fast-forward merge between master and feature/uppercase recently in the reflog
        if git reflog --date=iso | grep -q "merge feature/uppercase: Fast-forward"; then
            # Remove the feature/uppercase branch
            git branch -D feature/uppercase 2>/dev/null
            echo "success"
        else
            echo "No recent fast-forward merge found in the reflog"
        fi
    else
        echo "Branch feature/uppercase exists"
    fi
else
    echo "Repository does not have 3 commits"
fi
