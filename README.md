# concrete-template

A simple template for an application-specific rollup built with the [Concrete framework](https://github.com/concrete-eth/concrete-geth).

## Requirements

- Go 1.19 or higher
- [Concrete CLI](https://github.com/concrete-eth/concrete-geth#installing-the-concrete-cli) for code generation
- Yarn for running scripts and managing dependencies
- Foundry for compiling Solidity contracts
- Docker and Docker Compose for running a devnet

Installing the concrete CLI:

Clone the concrete github repo and install the CLI

```
https://github.com/concrete-eth/concrete-geth
cd concrete-geth
go install ./concrete/cmd/concrete
```

## Getting started

```bash
yarn install
yarn test
```

To get familiar with how concrete works, see [Quickstart.md](Quickstart.md).

## Running a devnet

Run a normal Optimism Bedrock devnet replacing op-geth for `ghcr.io/concrete-eth/concrete-template-geth:latest` in `ops-bedrock/Dockerfile.l2`. Build the image with `yarn build:docker`.

Alternatively, clone and run an already modified version of Bedrock:

```bash
# Clone repo
git clone -b concrete-template https://github.com/concrete-eth/optimism.git
cd optimism
# Start devnet
make devnet-up
# Stop and clean devnet
make devnet-down && make devnet-clean
```
