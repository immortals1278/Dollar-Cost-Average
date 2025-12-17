#!/bin/bash

# This script retrieves various attributes of a ZRC20 contract on a blockchain.
# It takes the contract address and an RPC URL as inputs and calls several 
# functions on the contract to obtain key information. These attributes 
# are typically required when a clone of the contract is to be deployed.
#
# Usage:
#   ./get_zrc20_info.sh <contract_address> <rpc_url>

# Check if cast is installed
if ! command -v cast &> /dev/null
then
    echo "cast CLI could not be found. Please install it first."
    exit 1
fi

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <contract_address> <rpc_url>"
    exit 1
fi

# Get the contract address and RPC URL from the script arguments
CONTRACT_ADDRESS=$1
RPC_URL=$2

# Run the cast commands and capture the outputs
SYSTEM_CONTRACT=$(cast call "$CONTRACT_ADDRESS" "SYSTEM_CONTRACT_ADDRESS()(address)" --rpc-url "$RPC_URL" | tr -d '\n')
ZRC20_NAME=$(cast call "$CONTRACT_ADDRESS" "name()(string)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/"//g')
ZRC20_SYMBOL=$(cast call "$CONTRACT_ADDRESS" "symbol()(string)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/"//g')
ZRC20_DECIMALS=$(cast call "$CONTRACT_ADDRESS" "decimals()(uint8)" --rpc-url "$RPC_URL" | tr -d '\n')
ZRC20_CHAIN_ID=$(cast call "$CONTRACT_ADDRESS" "CHAIN_ID()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/\[.*\]//g')
ZRC20_COIN_TYPE=$(cast call "$CONTRACT_ADDRESS" "COIN_TYPE()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n')
ZRC20_GAS_LIMIT=$(cast call "$CONTRACT_ADDRESS" "GAS_LIMIT()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/\[.*\]//g')

# Output the final result string
echo "SYSTEM_CONTRACT=$SYSTEM_CONTRACT"
echo "ZRC20_NAME=$ZRC20_NAME"
echo "ZRC20_SYMBOL=$ZRC20_SYMBOL"
echo "ZRC20_DECIMALS=$ZRC20_DECIMALS"
echo "ZRC20_CHAIN_ID=$ZRC20_CHAIN_ID"
echo "ZRC20_COIN_TYPE=$ZRC20_COIN_TYPE"
echo "ZRC20_GAS_LIMIT=$ZRC20_GAS_LIMIT"