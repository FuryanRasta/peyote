#!/usr/bin/env bash

# Uncomment the below to broadcast REST endpoint
# Do not forget to comment the bottom lines !!
# peyote start --pruning "everything" &
# peycli rest-server --chain-id peyotechain-1 --laddr="tcp://0.0.0.0:1317" --trust-node && fg

peyote start --pruning "everything" &
peycli rest-server --chain-id peyotechain-1 --trust-node && fg
