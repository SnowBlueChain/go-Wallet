# Operation Example
## Scenario 1
- 3 nodes are running `btc-watch`,`btc-keygen`,`btc-sign` respectively.
- To use multisig functionality, multiple auth accounts are required.
- Set 5 auth accounts on [account.toml](https://github.com/hiromaily/go-crypto-wallet/blob/master/data/config/account.toml)

### Prerequisite
See [Installation](https://github.com/hiromaily/go-crypto-wallet/blob/master/docs/Installation.md#bitcoind-setup)
1. `bitcoin-cli` needs to be ready
2. [account.toml](https://github.com/hiromaily/go-crypto-wallet/blob/master/data/config/account.toml) configuration
3. create wallets

### 1. Generate Key for Bitcoin
```
$ ./scripts/operation/generate-btc-key.sh btc false 5
```

### 2. Reset Data if needed, when error occurred under key generating process
- Remove data from Database, remove wallet.dat which includes account's private key
```
$ ./scripts/operation/reset-data.sh
```
- After removing data,
```
# run database containers
$ docker compose up watch-db keygen-db sign-db

# run bitcoind nodes
$ docker compose up btc-watch btc-keygen btc-sign

# create wallets on bitcoind
$ ./scripts/operation/create-bitcoind-wallet.sh
```

### 3. Try Deposit Action
- There are 4 accounts and when `client` receives coin into their address, these coins are transferred into `deposit` account address.

1. Check `client` account address from `address` table on `watch-db`
2. Copy client's address and send coin into them by [Faucet](https://signet.bc-2.jp/)
3. Run script
```
./scripts/operation/create-btc-tx-deposit.sh
```