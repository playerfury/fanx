# FURYX Network

The Sports, Gaming & Entertainment Network (FURYX Network), is a blockchain
designed to support the future of sports betting & related gaming by
leveraging the modular Cosmos design. We believe the future will be heavily shaped by many of the values driving the recent wave of crypto and blockchain development: transparency, increased decentralization, and utility that benefits all stakeholders, especially the user-base.

Utilizing a sovereign blockchain uniquely enables:

- An adaptable framework to design custom applications.
- Enablement of features, tools and economic models where users can directly benefit from the value they help create.
- An unparalleled level of transparency.
- An efficiency of settlement and immediate payout to participants.

At launch, the FURYX Network will be optimized to deploy an inaugural application: Six Sigma Sports, which is re-imagining the sports betting landscape and bringing a unique user experience with the benefit of blockchain technology.[Please visit to learn more about Six Sigma Sports.](https://sixsigmasports.io/)

---

## Hardware Requirements

- **Minimal**
  - 1 GB RAM
  - 25 GB SSD
  - 1.4 GHz CPU
- **Recommended**
  - 2 GB RAM
  - 100 GB SSD
  - 2.0 GHz x2 CPU

## Operating System

- Linux/Windows/MacOS(x86)
- **Recommended**
  - Linux(x86_64)

## Installation Steps
>
>Prerequisite: go1.18+ required. [ref](https://golang.org/doc/install)

Furyx could be installed by two ways - downloading binary from releases page or build from source.

### Download from releases page

- Download from release required binary

- Check sha256 hash sum

- Place fanx into /usr/local/sbin

```shell
sudo mv fanx /usr/local/sbin/fanx
```

### Building from source
>
>Optional requirement: git. [ref](https://github.com/git/git) and GNU make. [ref](https://www.gnu.org/software/make/manual/html_node/index.html)

- Clone git repository

```shell
git clone https://github.com/playerfury/fanx.git
```

- Checkout release tag

```shell
cd fanx
git fetch --tags
git checkout [vX.X.X]
```

- Install

```shell
go mod tidy
make install
```

### Install system.d service file

```shell
nano /etc/systemd/system/fanx.service
```

Please following contents(working dir may be changed as needed)

```systemd
[Unit]
Description=Furyx Network node
After=network.target

[Service]
Type=simple
User=ubuntu
WorkingDirectory=/home/ubuntu
ExecStart=/usr/local/sbin/fanx start
Restart=on-failure
RestartSec=10
LimitNOFILE=40960

[Install]
WantedBy=multi-user.target
```

Reload unit files in systemd

```shell
sudo systemctl daemon-reload
```

### Generate keys

`fanx keys add [key_name]`

or

`fanx keys add [key_name] --recover` to regenerate keys with your [BIP39](https://github.com/bitcoin/bips/tree/master/bip-0039) mnemonic

### Connect to a chain and start node

- [Install](#installation-steps) fanx application
- Initialize node

```shell
fanx init {{NODE_NAME}} --chain-id playerfury-1
```

Select network to join

- Replace `${HOME}/.fanx/config/genesis.json` with the genesis file of the chain.
- Add `persistent_peers` or `seeds` in `${HOME}/.fanx/config/config.toml`
- Start node

```shell
fanx start
```

## Network Compatibility Matrix

| Version | Mainnet | Testnet | SDK Version |
|:-------:|:-------:|:-------:|:-----------:|
|  v0.0.1 |    ✗    |    ✓    |   v0.45.4   |

## Active Networks

### Mainnet

Coming Soon!!

### Testnet

- [playerfury-1](https://github.com/playerfury/networks/playerfury-1)

- Place the genesis file  with the genesis file of the chain.

```shell
wget https://github.com/playerfury/networks/blob/master/playerfury-1/genesis.json -O ~/.fanx/config/genesis.json
```

Verify genesis hash sum

```shell
sha256sum ~/.fanx/config/genesis.json
```

Correct sha256 sum for playerfury-1 is - 2bea72699f9c1afd6217f7e76f14f07c1fbe849d090fc37cd008a42d14d5d30c
Genesis file sha sum is published in according repository.

- Add `persistent_peers` or `seeds` in `${HOME}/.fanx/config/config.toml`

```shell
sed -i '/s/persistent_peers = ""/persistent_peers = "4980b478f91de9be0564a547779e5c6cb07eb995@3.239.15.80:26656,0e7042be1b77707aaf0597bb804da90d3a606c08@3.88.40.53:26656/g' $HOME/.fanx/config/config.toml
```

- Start node

```shell
fanx start
```

### Initialize a new chain and start node

- Initialize: `fanx init [node_name] --chain-id [chain_name]`
- Add key for genesis account `fanx keys add [genesis_key_name]`
- Add genesis account `fanx add-genesis-account [genesis_key_name] 10000000000000000000ufanx`
- Create a validator at genesis `fanx gentx [genesis_key_name] 10000000ufanx --chain-id [chain_name]`
- Collect genesis transactions `fanx collect-gentxs`
- Start node `fanx start`

### Reset chain

```shell
rm -rf ~/.fanx
```

### Shutdown node

```shell
killall fanx
```

### Check version

```shell
fanx version
```

### Documentations

For the most up to date documentation please visit [Gitbook](https://six-sigma-sports.gitbook.io/documentation/)
