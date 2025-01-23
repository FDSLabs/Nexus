# Quickstart Guide

This guide helps you quickly set up and run a Nexus full node.

## Prerequisites

- **Operating System**: Linux (x86_64/amd64), Ubuntu recommended
- **Hardware**:
  - 8GB+ RAM
  - 100GB+ NVMe SSD
  - 4 CPU cores

## Installation Steps

### 1. Install Dependencies

```bash
# Install required packages
sudo snap install go --classic
sudo apt-get update
sudo apt-get install -y git gcc make
```

### 2. Install Nexus

```bash
# Clone repository
git clone https://github.com/FDSLabs/Nexus.git
cd Nexus

# Install binary
make install

# Verify installation
nexusd version
```

### 3. Initialize Node

```bash
# Initialize node
nexusd init <your-node-name> --chain-id nexus_1234

# Download genesis file
wget https://raw.githubusercontent.com/Nexus-Network/Nexus/genesis/Networks/Mainnet/genesis.json -P $HOME/.nexusd/config/

# Set minimum gas price
sed -i 's/minimum-gas-prices = "0anexus"/minimum-gas-prices = "0.0001anexus"/g' $HOME/.nexusd/config/app.toml

# Add persistent peers
sed -i 's/persistent_peers = ""/persistent_peers = "ec770ae4fd0fb4871b9a7c09f61764a0b010b293@164.90.134.106:26656"/g' $HOME/.nexusd/config/config.toml
```

### 4. Start Your Node

```bash
# Start the node
nexusd start
```

Your node will begin syncing with the network. This process may take several hours depending on your hardware and network connection.

## Next Steps

- [Become a Validator](validator-guide.md)
- [Join the Network](join-network.md)

## Common Operations

### Check Node Status

```bash
# Check sync status
nexusd status | jq .SyncInfo

# Check node info
nexusd status | jq .NodeInfo
```

### Reset Node

```bash
# Stop the node
nexusd unsafe-reset-all
```

### Common Issues

1. **Sync Issues**: If your node is stuck syncing, check:
   - Network connectivity
   - Disk space
   - System resources

2. **Connection Issues**: If you can't connect to peers:
   - Verify your firewall settings
   - Check your config.toml
   - Ensure your ports are open (26656, 26657)
