#!/usr/bin/env bash

PASSWORD="12345678"
GAS_PRICES="0.025stake"

peyote init local --chain-id peyotechain-1

yes $PASSWORD | peycli keys delete miguel --keyring-backend=test --force
yes $PASSWORD | peycli keys delete francesco --keyring-backend=test --force
yes $PASSWORD | peycli keys delete shaun --keyring-backend=test --force
yes $PASSWORD | peycli keys delete fee --keyring-backend=test --force
yes $PASSWORD | peycli keys delete fee2 --keyring-backend=test --force
yes $PASSWORD | peycli keys delete fee3 --keyring-backend=test --force
yes $PASSWORD | peycli keys delete fee4 --keyring-backend=test --force
yes $PASSWORD | peycli keys delete fee5 --keyring-backend=test --force

yes $PASSWORD | peycli keys add miguel --keyring-backend=test
yes $PASSWORD | peycli keys add francesco --keyring-backend=test
yes $PASSWORD | peycli keys add shaun --keyring-backend=test
yes $PASSWORD | peycli keys add fee --keyring-backend=test
yes $PASSWORD | peycli keys add fee2 --keyring-backend=test
yes $PASSWORD | peycli keys add fee3 --keyring-backend=test
yes $PASSWORD | peycli keys add fee4 --keyring-backend=test
yes $PASSWORD | peycli keys add fee5 --keyring-backend=test

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | peyote add-genesis-account $(peycli keys show miguel --keyring-backend=test -a) 200000000stake,1000000res,1000000rez
yes $PASSWORD | peyote add-genesis-account $(peycli keys show francesco --keyring-backend=test -a) 100000000stake,1000000res,1000000rez
yes $PASSWORD | peyote add-genesis-account $(peycli keys show shaun --keyring-backend=test -a) 100000000stake,1000000res,1000000rez

# Set min-gas-prices
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025stake\""
sed -i "s/$FROM/$TO/" "$HOME"/.peyote/config/app.toml

peycli config chain-id peyotechain-1
peycli config output json
peycli config indent true
peycli config trust-node true
peycli config keyring-backend test

yes $PASSWORD | peyote gentx --name miguel --keyring-backend=test

peyote collect-gentxs
peyote validate-genesis

# Uncomment the below to broadcast node RPC endpoint
#FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
#TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
#sed -i "s/$FROM/$TO/" "$HOME"/.peyote/config/config.toml

# Uncomment the below to broadcast REST endpoint
# Do not forget to comment the bottom lines !!
# peyote start --pruning "everything" &
# peycli rest-server --chain-id peyotechain-1 --laddr="tcp://0.0.0.0:1317" --trust-node && fg

peyote start --pruning "everything" &
peycli rest-server --chain-id peyotechain-1 --trust-node && fg
