#!/bin/bash

# Check if enough arguments are provided
if [ "$#" -ne 1 ]; then
    echo "Usage: ./start_bundler.sh <path_to_bundler>"
    exit 1
fi

# Extract paths from command line arguments
BUNDLER_PATH=$1

# Navigate to the bundler directory and run the yarn commands
cd "$BUNDLER_PATH" || { echo "Failed to navigate to bundler directory"; exit 1; }
yarn hardhat-deploy --network localhost
yarn run bundler > bundler.log 2>&1 &
