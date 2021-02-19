#!/usr/bin/env bash

rm -rf ~/.nameserviced
rm -rf ~/.nameservicecli

~/go/bin/nameserviced init test --chain-id=namechain

~/go/bin/nameservicecli config output json
~/go/bin/nameservicecli config indent true
~/go/bin/nameservicecli config trust-node true
~/go/bin/nameservicecli config chain-id namechain
~/go/bin/nameservicecli config keyring-backend test

~/go/bin/nameservicecli keys add user1
~/go/bin/nameservicecli keys add user2

~/go/bin/nameserviced add-genesis-account $(~/go/bin/nameservicecli keys show user1 -a) 1000nametoken,100000000stake
~/go/bin/nameserviced add-genesis-account $(~/go/bin/nameservicecli keys show user2 -a) 1000nametoken,100000000stake

~/go/bin/nameserviced gentx --name user1 --keyring-backend test

echo "Collecting genesis txs..."
~/go/bin/nameserviced collect-gentxs

echo "Validating genesis file..."
~/go/bin/nameserviced validate-genesis