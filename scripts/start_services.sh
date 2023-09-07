#!/bin/bash

# Check if enough arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: ./start_services.sh <path_to_optimism> <path_to_bundler>"
    exit 1
fi

# Extract paths from command line arguments
OPTIMISM_PATH=$1
BUNDLER_PATH=$2

# Navigate to the optimism directory and start the devnet
cd "$OPTIMISM_PATH" || { echo "Failed to navigate to optimism directory"; exit 1; }
make devnet-up

# Navigate to the bundler directory and run the yarn commands
cd "$BUNDLER_PATH" || { echo "Failed to navigate to bundler directory"; exit 1; }
yarn hardhat-deploy --network localhost
yarn run bundler
