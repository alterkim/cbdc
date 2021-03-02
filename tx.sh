#!/bin/bash

~/go/bin/cbdccli query account $(~/go/bin/cbdccli keys show user1 -a) | jq ".value.coins[0]"
~/go/bin/cbdccli query account $(~/go/bin/cbdccli keys show user2 -a) | jq ".value.coins[0]"

# Buy your first name using your coins from the genesis file
~/go/bin/cbdccli tx cbdc buy-name user1.id 5cbdctoken --from user1 -y | jq ".txhash" |  xargs $(sleep 6) ~/go/bin/cbdccli q tx

# Set the value for the name you just bought
~/go/bin/cbdccli tx cbdc set-name user1.id 8.8.8.8 --from user1 -y | jq ".txhash" |  xargs $(sleep 6) ~/go/bin/cbdccli q tx

# Try out a resolve query against the name you registered
~/go/bin/cbdccli query cbdc resolve user1.id | jq ".value"
# > 8.8.8.8

# Try out a whois query against the name you just registered
~/go/bin/cbdccli query cbdc get-whois user1.id
# > {"value":"8.8.8.8","owner":"cosmos1l7k5tdt2qam0zecxrx78yuw447ga54dsmtpk2s","price":[{"denom":"nametoken","amount":"5"}]}

# user2 buys name from user1
~/go/bin/cbdccli tx cbdc buy-name user1.id 10cbdctoken --from user2 -y | jq ".txhash" |  xargs $(sleep 6) ~/go/bin/cbdccli q tx

# user2 decides to delete the name she just bought from user1
~/go/bin/cbdccli tx cbdc delete-name user1.id --from user2 -y | jq ".txhash" |  xargs $(sleep 6) ~/go/bin/cbdccli q tx

# Try out a whois query against the name you just deleted
~/go/bin/cbdccli query cbdc get-whois user1.id
# > could not resolve whois user1.id 