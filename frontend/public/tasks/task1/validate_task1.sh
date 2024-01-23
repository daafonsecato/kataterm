#!/bin/bash
# validate_task1.sh

# Perform validation logic here
# For example, check if a certain file exists with the correct content
if [ -f "solution.txt" ] && grep -q "correct content" "solution.txt"; then
    echo "TASK_VALIDATION_FAILURE: Task completed successfully."
    exit 0
else
    echo "TASK_VALIDATION_FAILURE: Task validation failed. Please review your solution."
    exit 1
fi
