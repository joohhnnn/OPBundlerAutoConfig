#!/bin/bash

# User-provided path to the Bundler directory
BUNDLER_DIR=$1

# Check if the Bundler directory exists
if [ ! -d "$BUNDLER_DIR" ]; then
  echo "The provided Bundler directory does not exist."
  exit 1
fi

# Define the target files
DEPLOY_FILE="$BUNDLER_DIR/packages/bundler/deploy/2-deploy-entrypoint.ts"
HARDHAT_CONFIG="$BUNDLER_DIR/packages/bundler/hardhat.config.ts"
LOCAL_CONFIG="$BUNDLER_DIR/packages/bundler/localconfig/bundler.config.json"

# Check if the target files exist
if [ ! -f "$DEPLOY_FILE" ] || [ ! -f "$HARDHAT_CONFIG" ] || [ ! -f "$LOCAL_CONFIG" ]; then
  echo "One or more target files do not exist in the provided directory."
  exit 1
fi

# Modify 2-deploy-entrypoint.ts
sed -i 's/net.chainId !== 1337 && net.chainId !== 31337/net.chainId !== 1337 && net.chainId !== 31337 && net.chainId !== 901/g' $DEPLOY_FILE

# Modify hardhat.config.ts
sed -i 's/url: '\''http:\/\/localhost:8545'\''/url: '\''http:\/\/localhost:9545'\''/g' $HARDHAT_CONFIG

# Modify bundler.config.json
sed -i 's/"network": "http:\/\/127.0.0.1:8545"/"network": "http:\/\/127.0.0.1:9545"/g' $LOCAL_CONFIG

echo "Successfully modified Bundler-related files."