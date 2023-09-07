#!/bin/sh
set -exu

VERBOSITY=${GETH_VERBOSITY:-3}
GETH_DATA_DIR=/db
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GETH_KEYSTORE_DIR="$GETH_DATA_DIR/keystore"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
CHAIN_ID=$(cat "$GENESIS_FILE_PATH" | jq -r .config.chainId)
BLOCK_SIGNER_PRIVATE_KEY="ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
BLOCK_SIGNER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"

# Check and create GETH_CHAINDATA_DIR
if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
  echo "Creating $GETH_CHAINDATA_DIR..."
  mkdir -p "$GETH_CHAINDATA_DIR"
else
  echo "$GETH_CHAINDATA_DIR exists."
fi

# Check and create GETH_KEYSTORE_DIR
if [ ! -d "$GETH_KEYSTORE_DIR" ]; then
  echo "Creating $GETH_KEYSTORE_DIR..."
  mkdir -p "$GETH_KEYSTORE_DIR"
else
  echo "$GETH_KEYSTORE_DIR exists."
fi

# Check if GETH_KEYSTORE_DIR is empty and run account import
if [ ! "$(ls -A $GETH_KEYSTORE_DIR)" ]; then
    echo "$GETH_KEYSTORE_DIR is empty, running account import"
    echo -n "pwd" > "$GETH_DATA_DIR"/password
    echo -n "$BLOCK_SIGNER_PRIVATE_KEY" | sed 's/0x//' > "$GETH_DATA_DIR"/block-signer-key
    geth account import \
        --datadir="$GETH_DATA_DIR" \
        --password="$GETH_DATA_DIR"/password \
        "$GETH_DATA_DIR"/block-signer-key
else
    echo "$GETH_KEYSTORE_DIR is not empty, skipping account import"
fi

# Initialize genesis if GETH_CHAINDATA_DIR is missing
if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
    echo "$GETH_CHAINDATA_DIR missing, running init"
    echo "Initializing genesis."
    geth --verbosity="$VERBOSITY" init \
        --datadir="$GETH_DATA_DIR" \
        "$GENESIS_FILE_PATH"
else
    echo "$GETH_CHAINDATA_DIR exists."
fi

# Start geth
exec geth \
    --datadir="$GETH_DATA_DIR" \
    --verbosity="$VERBOSITY" \
    --http \
    --http.corsdomain="*" \
    --http.vhosts="*" \
    --http.addr=0.0.0.0 \
    --http.port="$RPC_PORT" \
    --http.api=web3,debug,eth,txpool,net,engine \
    --ws \
    --ws.addr=0.0.0.0 \
    --ws.port="$WS_PORT" \
    --ws.origins="*" \
    --ws.api=debug,eth,txpool,net,engine \
    --syncmode=full \
    --nodiscover \
    --maxpeers=0 \
    --networkid=$CHAIN_ID \
    --unlock=$BLOCK_SIGNER_ADDRESS \
    --allow-insecure-unlock \
    --password="$GETH_DATA_DIR"/password \
    --rpc.allow-unprotected-txs \
    --authrpc.addr="0.0.0.0" \
    --authrpc.port="8551" \
    --authrpc.vhosts="*" \
    --authrpc.jwtsecret=/config/jwt-secret.txt \
    --gcmode=archive \
    --metrics \
    --metrics.addr=0.0.0.0 \
    --metrics.port=6060 \
    "$@"