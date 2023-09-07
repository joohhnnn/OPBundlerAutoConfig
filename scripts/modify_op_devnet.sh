#!/bin/bash

# User-provided path to the Optimism directory
OPTIMISM_DIR=$1

# Check if the Optimism directory exists
if [ ! -d "$OPTIMISM_DIR" ]; then
  echo "The provided Optimism directory does not exist."
  exit 1
fi

# Locate the entrypoint-l2.sh file under ops-bedrock
TARGET_FILE="$OPTIMISM_DIR/ops-bedrock/entrypoint-l2.sh"

# Check if the target file exists
if [ ! -f "$TARGET_FILE" ]; then
  echo "The target file entrypoint-l2.sh does not exist in the provided directory."
  exit 1
fi

# Replace the target file with the entrypoint-l2.sh file from the overrides directory
cp overrides/entrypoint-l2.sh $TARGET_FILE

echo "Successfully replaced entrypoint-l2.sh."
