#!/bin/bash

# Define the file name
file="CHANGELOG.md"

# Check if the file exists
if [ -e "$file" ]; then
    # Append the demo line to the end of the file
    echo "## Demo" >> "$file"
    echo "- Added demo line" >> "$file"
    echo "Demo complete" >> "$file"
    echo "Demo changes documented in changelog.md"
else
    # If the file doesn't exist, display an error message
    echo "Error: File $file not found."
fi
