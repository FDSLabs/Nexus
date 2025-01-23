# Nexus Validator Setup Guide

This guide explains how to become a validator on the Nexus Mainnet (*nexus_1234*).

> Genesis file [Published](https://github.com/FDSLabs/Nexus/raw/main/Mainnet/genesis.json)
> Peers list [Published](https://github.com/FDSLabs/Nexus/blob/main/Mainnet/peers.txt)

## Hardware Requirements

### Minimum

* 16 GB RAM
* 100 GB NVME SSD
* 3.2 GHz x4 CPU

### Recommended

* 32 GB RAM
* 500 GB NVME SSD
* 4.2 GHz x6 CPU

### Operating System

* Linux (x86_64) or Linux (amd64)
* Recommended Ubuntu

## Installation Guide

### 1. Install Dependencies

**If using Ubuntu:**

Install all dependencies:

```bash
sudo snap install go --classic && sudo apt-get install git && sudo apt-get install gcc && sudo apt-get install make
```

Or install individually:

* go1.18+: `sudo snap install go --classic`
* git: `sudo apt-get install git`
* gcc: `sudo apt-get install gcc`
* make: `sudo apt-get install make`

### 2. Install `nexusd`

Clone git repository:

```bash
git clone https://github.com/FDSLabs/Nexus.git
cd nexus/cmd/nexusd
go install -tags ledger ./...
sudo mv $HOME/go/bin/nexusd /usr/bin/
```

### 3. Generate and Store Keys

Replace `<key_name>` with your preferred key name:

```bash
nexusd keys add <key_name>
# OR to recover existing keys:
nexusd keys add <key_name> --recover
# OR to use a ledger device:
nexusd keys add <key_name> --ledger
```

Store your keys and mnemonic securely offline.

Save the public key config in the main Nexus directory as `<key_name>.info`:

```
pubkey: {
  "@type":" ethermint.crypto.v1.ethsecp256k1.PubKey",
  "key":"############################################"
}
```

### 4. Set Up Validator Node

Install nexusd binary:

```bash
sudo make install
```

Initialize the node (replace `<moniker>` with your validator name):

```bash
nexusd init <moniker> --chain-id nexus_nexus_1234
```

Download Genesis file:

```bash
wget https://raw.githubusercontent.com/Nexus-Network/Nexus/genesis/Networks/Mainnet/genesis.json -P $HOME/.nexusd/config/
```

Configure minimum gas prices:

```bash
sed -i 's/minimum-gas-prices = "0anexus"/minimum-gas-prices = "0.0001anexus"/g' $HOME/.nexusd/config/app.toml
```

Add persistent peers:

```bash
sed -i 's/persistent_peers = ""/persistent_peers = "ec770ae4fd0fb4871b9a7c09f61764a0b010b293@164.90.134.106:26656"/g' $HOME/.nexusd/config/config.toml
```

### 5. Configure System Service

Create systemd service file:

```bash
sudo nano /etc/systemd/system/nexusd.service
```

Add the following content:

```ini
[Unit]
Description=Nexus Node
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/
ExecStart=/root/go/bin/nexusd start --trace --log_level info --json-rpc.api eth,txpool,net,debug,web3 --api.enable
Restart=on-failure
StartLimitInterval=0
RestartSec=3
LimitNOFILE=65535
LimitMEMLOCK=209715200

[Install]
WantedBy=multi-user.target
```

### 6. Start the Node

```bash
sudo systemctl daemon-reload
sudo systemctl enable nexusd.service
sudo systemctl start nexusd && journalctl -u nexusd -f
```

### 7. Create Validator Transaction

Modify and run the following command with your details:

```bash
nexusd tx staking create-validator \
--from <KEY_NAME> \
--chain-id nexus_nexus_1234 \
--moniker="<VALIDATOR_NAME>" \
--commission-max-change-rate=0.01 \
--commission-max-rate=1.0 \
--commission-rate=0.05 \
--details="<DESCRIPTION>" \
--security-contact="<SECURITY_CONTACT_EMAIL>" \
--website="<YOUR_WEBSITE>" \
--pubkey $(nexusd tendermint show-validator) \
--min-self-delegation="1" \
--amount <TOKEN_DELEGATION>anexus \
--fees 20anexus
```

Replace the following:

* `<KEY_NAME>`: Your key name from step 3
* `<VALIDATOR_NAME>`: Your validator node name
* `<DESCRIPTION>`: Description of your validator
* `<SECURITY_CONTACT_EMAIL>`: Contact email for security issues
* `<YOUR_WEBSITE>`: Your validator's website
* `<TOKEN_DELEGATION>`: Amount of tokens to stake (minimum 1anexus)
