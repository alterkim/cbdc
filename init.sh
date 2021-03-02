#!/usr/bin/env bash

rm -rf ~/.cbdcd
rm -rf ~/.cbdccli

~/go/bin/cbdcd init test --chain-id=namechain

~/go/bin/cbdccli config output json
~/go/bin/cbdccli config indent true
~/go/bin/cbdccli config trust-node true
~/go/bin/cbdccli config chain-id namechain
~/go/bin/cbdccli config keyring-backend test

~/go/bin/cbdccli keys add user1
~/go/bin/cbdccli keys add user2

~/go/bin/cbdcd add-genesis-account $(~/go/bin/cbdccli keys show user1 -a) 1000nametoken,100000000stake
~/go/bin/cbdcd add-genesis-account $(~/go/bin/cbdccli keys show user2 -a) 1000nametoken,100000000stake

~/go/bin/cbdcd gentx --name user1 --keyring-backend test

echo "Collecting genesis txs..."
~/go/bin/cbdcd collect-gentxs

echo "Validating genesis file..."
~/go/bin/cbdcd validate-genesis