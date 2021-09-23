#!/bin/sh

set -eu

COIN="${1:?eth}"

# 1: ganache, 2:go-ethereum, 3: something else
CLIENT=1

###############################################################################
# keygen wallet
###############################################################################
if [ $CLIENT -eq 1 ]; then
  echo import ganache keys
  docker compose exec btc-keygen-db mysql -u root -proot  -e "$(cat ./sql/ganache_key.sql)"
else
  # create seed
  keygen -coin ${COIN} create seed
  # create hdkey for client, deposit, payment account
  keygen -coin ${COIN} create hdkey -account client -keynum 10
  keygen -coin ${COIN} create hdkey -account deposit -keynum 1
  keygen -coin ${COIN} create hdkey -account payment -keynum 1
  keygen -coin ${COIN} create hdkey -account stored -keynum 1
fi

# import generated private key into keygen wallet (this command should run on ethereum server)
# create key files on key store directory
keygen -coin ${COIN} import privkey -account client
keygen -coin ${COIN} import privkey -account deposit
keygen -coin ${COIN} import privkey -account payment
keygen -coin ${COIN} import privkey -account stored

# export address
file_address_client=$(keygen -coin "${COIN}" export address -account client)
file_address_deposit=$(keygen -coin "${COIN}" export address -account deposit)
file_address_payment=$(keygen -coin "${COIN}" export address -account payment)
file_address_stored=$(keygen -coin "${COIN}" export address -account stored)

###############################################################################
# watch only wallet
###############################################################################
# import addresses generated by keygen wallet
watch -coin ${COIN} import address -file ${file_address_client##*\[fileName\]: }
watch -coin ${COIN} import address -file ${file_address_deposit##*\[fileName\]: }
watch -coin ${COIN} import address -file ${file_address_payment##*\[fileName\]: }
watch -coin ${COIN} import address -file ${file_address_stored##*\[fileName\]: }
