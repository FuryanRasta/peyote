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

create_bond_multisig() {
  peycli tx peyote create-bond \
    --token=abc \
    --name="A B C" \
    --description="Description about A B C" \
    --function-type=power_function \
    --function-parameters="m:12,n:2,c:100" \
    --reserve-tokens=res \
    --tx-fee-percentage=0.5 \
    --exit-fee-percentage=0.1 \
    --fee-address="$FEE" \
    --max-supply=1000000abc \
    --order-quantity-limits="" \
    --sanity-rate="0" \
    --sanity-margin-percentage="0" \
    --allow-sells \
    --signers="$(peycli keys show francesco --keyring-backend=test -a),$(peycli keys show shaun --keyring-backend=test -a)" \
    --batch-blocks=1 \
    --from="$MIGUEL" -y --broadcast-mode block --generate-only >multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=francesco --output-document=multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=shaun --output-document=multisig.json
  peycli tx broadcast multisig.json
  rm multisig.json
}

edit_bond_multisig_incorrect_signers_1() {
  peycli tx peyote edit-bond \
    --token=abc \
    --name="(1) New A B C" \
    --description="(1) New description about A B C" \
    --signers="$(peycli keys show shaun --keyring-backend=test -a),$(peycli keys show francesco --keyring-backend=test -a)" \
    --from="$MIGUEL" -y --broadcast-mode block --generate-only >multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=shaun --output-document=multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=francesco --output-document=multisig.json
  peycli tx broadcast multisig.json
  rm multisig.json
}

edit_bond_multisig_incorrect_signers_2() {
  peycli tx peyote edit-bond \
    --token=abc \
    --name="(2) New A B C" \
    --description="(2) New description about A B C" \
    --signers="$FRANCESCO" \
    --from="$MIGUEL" -y --broadcast-mode block --generate-only >multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=francesco --output-document=multisig.json
  peycli tx broadcast multisig.json
  rm multisig.json
}

edit_bond_multisig_correct_signers() {
  peycli tx peyote edit-bond \
    --token=abc \
    --name="(3) New A B C" \
    --description="(3) New description about A B C" \
    --signers="$(peycli keys show francesco --keyring-backend=test -a),$(peycli keys show shaun --keyring-backend=test -a)" \
    --from="$MIGUEL" -y --broadcast-mode block --generate-only >multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=francesco --output-document=multisig.json
  yes $PASSWORD | peycli tx sign multisig.json --from=shaun --output-document=multisig.json
  peycli tx broadcast multisig.json
  rm multisig.json
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
create_bond_multisig
echo "Waiting a bit..."
sleep 5
echo "Created bond..."
peycli q peyote bond abc

echo "Editing bond with incorrect signers..."
edit_bond_multisig_incorrect_signers_1
echo "Waiting a bit..."
sleep 5
peycli q peyote bond abc
echo "Bond was NOT edited!"

echo "Editing bond with incorrect signers again..."
edit_bond_multisig_incorrect_signers_2
echo "Waiting a bit..."
sleep 5
peycli q peyote bond abc
echo "Bond was NOT edited!"

echo "Editing bond with correct..."
edit_bond_multisig_correct_signers
echo "Waiting a bit..."
sleep 5
peycli q peyote bond abc
echo "Bond was edited!"
