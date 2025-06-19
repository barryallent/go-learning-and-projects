#!/bin/bash

# Run All Go Learning Examples
# This script runs all your Go examples in the recommended learning order

echo "🚀 Running all Go learning examples..."
echo ""

# Change to learning directory
cd learning

# Array of directories in learning order
directories=(
    "01-variables"
    "02-fmt"
    "03-booleans"
    "04-strings"
    "05-loops"
    "06-arrays"
    "07-maps"
    "08-functions"
    "09-receivers"
    "10-pointers"
    "11-structs"
    "12-pass-by-value"
    "13-packages"
)

# Run each example
for dir in "${directories[@]}"; do
    if [ -d "$dir" ]; then
        echo "=========================================="
        echo "📁 Running: $dir"
        echo "=========================================="
        
        cd "$dir"
        
        # Special handling for packages (run both files)
        if [ "$dir" = "13-packages" ]; then
            if [ -f "main.go" ] && [ -f "helper.go" ]; then
                go run main.go helper.go
            elif [ -f "main.go" ]; then
                go run main.go
            else
                echo "❌ No Go files found in $dir"
            fi
        else
            # Check if main.go exists
            if [ -f "main.go" ]; then
                go run main.go
            elif [ -f "*.go" ]; then
                go run *.go
            else
                echo "❌ No Go files found in $dir"
            fi
        fi
        
        cd ..
        echo ""
    else
        echo "⚠️  Directory $dir not found"
    fi
done

# Go back to root directory
cd ..

echo "🎉 All examples completed!"
echo ""
echo "💡 Tips:"
echo "   - Each example is now in its own directory"
echo "   - You can run individual examples: cd learning/01-variables && go run main.go"
echo "   - Follow the README.md for detailed explanations"
echo "   - Experiment by modifying the code in each directory"
