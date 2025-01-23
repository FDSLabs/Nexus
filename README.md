# Nexus Core

[![Build Status](https://github.com/FDSLabs/Nexus/workflows/Build%20Nexus%20Core/badge.svg)](https://github.com/FDSLabs/Nexus/actions/workflows/build.yml)
[![Tests Status](https://github.com/FDSLabs/Nexus/workflows/Tests%20%2F%20Code%20Coverage/badge.svg)](https://github.com/FDSLabs/Nexus/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/FDSLabs/Nexus/branch/main/graph/badge.svg)](https://codecov.io/gh/FDSLabs/Nexus)
[![Go Report Card](https://goreportcard.com/badge/github.com/FDSLabs/Nexus)](https://goreportcard.com/report/github.com/FDSLabs/Nexus)
[![license](https://img.shields.io/github/license/FDSLabs/Nexus.svg)](https://github.com/FDSLabs/Nexus/blob/main/LICENSE)
[![Lines of Code](https://tokei.rs/b1/github/FDSLabs/Nexus)](https://github.com/FDSLabs/Nexus)

Nexus is a scalable and interoperable blockchain platform built using the Cosmos SDK.

## Chain Specifications

- **Chain ID**: nexus_1234
- **Binary**: `nexusd`
- **Denom**: `anexus`
- **Version**: [Latest Release](https://github.com/FDSLabs/Nexus/releases)

### Key Features

- EVM compatibility
- IBC (Inter-Blockchain Communication) enabled
- Fast finality through Tendermint BFT consensus
- Scalable architecture
- Native token staking and delegation
- Governance system

## Quick Start

- [Run a Full Node](docs/quickstart.md)
- [Become a Validator](docs/validator-guide.md)
- [Join the Network](docs/join-network.md)

## Documentation

For detailed documentation, visit our [docs](docs/) directory.

## Development

### Prerequisites

- [Go](https://golang.org/doc/install) 1.21+
- [Make](https://www.gnu.org/software/make/)

### Build

```bash
make build
```

### Test

```bash
make test
```

### Install

```bash
make install
```

## Contributing

We welcome contributions from the community. Please read our [contributing guidelines](CONTRIBUTING.md) before submitting a pull request.

## Security

See our [security policy](SECURITY.md) for reporting security vulnerabilities.

## License

This project is licensed under the [GNU Lesser General Public License v3.0 (LGPL-3.0)](LICENSE).
