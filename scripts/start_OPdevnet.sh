#!/bin/bash

# Check if enough arguments are provided
if [ "$#" -ne 1 ]; then
    echo "Usage: ./start_OPdevnet.sh <path_to_optimism>"
    exit 1
fi

# Extract paths from command line arguments
OPTIMISM_PATH=$1

# Navigate to the optimism directory and start the devnet
cd "$OPTIMISM_PATH" || { echo "Failed to navigate to optimism directory"; exit 1; }
make devnet-up

