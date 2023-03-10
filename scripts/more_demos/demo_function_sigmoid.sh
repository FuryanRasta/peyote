#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(peycli status 2>&1)
    if [[ ($RET == ERROR*) || ($RET == *'"latest_block_height": "0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

tx_from_m() {
  cmd=$1
  shift
  yes $PASSWORD | peycli tx peyote "$cmd" --from miguel --keyring-backend=test -y --broadcast-mode block --gas-prices="$GAS_PRICES" "$@"
}

tx_from_f() {
  cmd=$1
  shift
  yes $PASSWORD | peycli tx peyote "$cmd" --from francesco --keyring-backend=test -y --broadcast-mode block --gas-prices="$GAS_PRICES" "$@"
}

RET=$(peycli status 2>&1)
if [[ ($RET == ERROR*) || ($RET == *'"latest_block_height": "0"'*) ]]; then
  wait
fi

PASSWORD="12345678"
GAS_PRICES="0.025stake"
MIGUEL=$(yes $PASSWORD | peycli keys show miguel --keyring-backend=test -a)
FRANCESCO=$(yes $PASSWORD | peycli keys show francesco --keyring-backend=test -a)
SHAUN=$(yes $PASSWORD | peycli keys show shaun --keyring-backend=test -a)
FEE=$(yes $PASSWORD | peycli keys show fee --keyring-backend=test -a)

echo "Creating bond..."
tx_from_m create-bond \
  --token=abc \
  --name="A B C" \
  --description="Description about A B C" \
  --function-type=sigmoid_function \
  --function-parameters="a:3,b:5,c:1" \
  --reserve-tokens=res \
  --tx-fee-percentage=0.5 \
  --exit-fee-percentage=0.1 \
  --fee-address="$FEE" \
  --max-supply=1000000abc \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells \
  --signers="$MIGUEL" \
  --batch-blocks=1
echo "Created bond..."
peycli q peyote bond abc

echo "Miguel buys 50abc..."
tx_from_m buy 50abc 1000000res
echo "Miguel's account..."
peycli q auth account "$MIGUEL"

echo "Francesco buys 50abc..."
tx_from_f buy 50abc 1000000res
echo "Francesco's account..."
peycli q auth account "$FRANCESCO"

echo "Miguel sells 50abc..."
tx_from_m sell 50abc
echo "Miguel's account..."
peycli q auth account "$MIGUEL"

echo "Francesco sells 50abc..."
tx_from_f sell 50abc
echo "Francesco's account..."
peycli q auth account "$FRANCESCO"
