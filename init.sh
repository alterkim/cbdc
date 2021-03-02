#!/usr/bin/env bash

rm -rf ~/.cbdcd
rm -rf ~/.cbdccli

cbdcd init test --chain-id=namechain

cbdccli config output json
cbdccli config indent true
cbdccli config trust-node true
cbdccli config chain-id namechain
cbdccli config keyring-backend test

cbdccli keys add user1
cbdccli keys add user2

cbdcd add-genesis-account $(cbdccli keys show user1 -a) 1000cbdctoken,100000000stake
cbdcd add-genesis-account $(cbdccli keys show user2 -a) 1000cbdctoken,100000000stake

cbdcd gentx --name user1 --keyring-backend test

echo "Collecting genesis txs..."
cbdcd collect-gentxs

echo "Validating genesis file..."
cbdcd validate-genesis