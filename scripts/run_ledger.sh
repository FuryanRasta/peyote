#!/usr/bin/env bash

PASSWORD="12345678"
GAS_PRICES="0.025stake"

peyote init local --chain-id peyotechain-1

peycli keys add miguel --ledger
yes $PASSWORD | peycli keys add francesco
yes $PASSWORD | peycli keys add shaun
yes $PASSWORD | peycli keys add reserve
yes $PASSWORD | peycli keys add fee

peyote add-genesis-account "$(peycli keys show miguel -a)" 100000000stake,1000000res,1000000rez
peyote add-genesis-account "$(peycli keys show francesco -a)" 100000000stake,1000000res,1000000rez
peyote add-genesis-account "$(peycli keys show shaun -a)" 100000000stake,1000000res,1000000rez

peycli config chain-id peyotechain-1
peycli config output json
peycli config indent true
peycli config trust-node true

echo "$PASSWORD" | peyote gentx --name miguel

peyote collect-gentxs
peyote validate-genesis

peyote start --pruning "everything" &
peycli rest-server --chain-id peyotechain-1 --trust-node && fg
