#!/bin/bash

# This script compares the attributes of two ZRC20 contracts on a blockchain.
# It takes two contract addresses and an RPC URL, fetches their values, and compares them.
#
# Usage:
#   ./compare_zrc20_info.sh <contract_address_1> <contract_address_2> <rpc_url>


# Check if cast is installed
if ! command -v cast &> /dev/null
then
    echo "cast CLI could not be found. Please install it first."
    exit 1
fi

# Check if the correct number of arguments is provided
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <contract_address_1> <contract_address_2> <rpc_url>"
    exit 1
fi

# Get the contract addresses and RPC URL from the script arguments
CONTRACT_ADDRESS_1=$1
CONTRACT_ADDRESS_2=$2
RPC_URL=$3

# Function to fetch ZRC20 attributes
fetch_zrc20_info() {
    local CONTRACT_ADDRESS=$1
    local RPC_URL=$2

    SYSTEM_CONTRACT=$(cast call "$CONTRACT_ADDRESS" "SYSTEM_CONTRACT_ADDRESS()(address)" --rpc-url "$RPC_URL" | tr -d '\n')
    ZRC20_NAME=$(cast call "$CONTRACT_ADDRESS" "name()(string)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/"//g')
    ZRC20_SYMBOL=$(cast call "$CONTRACT_ADDRESS" "symbol()(string)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/"//g')
    ZRC20_DECIMALS=$(cast call "$CONTRACT_ADDRESS" "decimals()(uint8)" --rpc-url "$RPC_URL" | tr -d '\n')
    ZRC20_CHAIN_ID=$(cast call "$CONTRACT_ADDRESS" "CHAIN_ID()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n')
    ZRC20_COIN_TYPE=$(cast call "$CONTRACT_ADDRESS" "COIN_TYPE()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n')
    ZRC20_GAS_LIMIT=$(cast call "$CONTRACT_ADDRESS" "GAS_LIMIT()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/\[.*\]//g')

    echo "$SYSTEM_CONTRACT,$ZRC20_NAME,$ZRC20_SYMBOL,$ZRC20_DECIMALS,$ZRC20_CHAIN_ID,$ZRC20_COIN_TYPE,$ZRC20_GAS_LIMIT"
}

# Fetch attributes for both contracts
INFO_1=$(fetch_zrc20_info $CONTRACT_ADDRESS_1 $RPC_URL)
INFO_2=$(fetch_zrc20_info $CONTRACT_ADDRESS_2 $RPC_URL)

# Split the results into arrays
IFS=',' read -r -a VALUES_1 <<< "$INFO_1"
IFS=',' read -r -a VALUES_2 <<< "$INFO_2"

# Attribute names for clarity
ATTRIBUTES=("SYSTEM_CONTRACT" "ZRC20_NAME" "ZRC20_SYMBOL" "ZRC20_DECIMALS" "ZRC20_CHAIN_ID" "ZRC20_COIN_TYPE" "ZRC20_GAS_LIMIT")

# Compare values and output the result
for i in "${!ATTRIBUTES[@]}"; do
    if [ "${VALUES_1[i]}" == "${VALUES_2[i]}" ]; then
        echo -e "✓ ${ATTRIBUTES[i]} is the same: ${VALUES_1[i]}"
    else
        echo -e "✗ ${ATTRIBUTES[i]} is different: ${VALUES_1[i]} (Contract 1) vs ${VALUES_2[i]} (Contract 2)"
    fi
done