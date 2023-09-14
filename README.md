# concrete-template

A simple template for an application-specific rollup built with the [Concrete framework](https://github.com/concrete-eth/concrete-geth).

## Requirements

- Go 1.19 or higher
- [Concrete CLI](https://github.com/concrete-eth/concrete-geth#installing-the-concrete-cli) for code generation
- Yarn for running scripts and managing dependencies
- Foundry for compiling Solidity contracts
- Docker and Docker Compose for running a devnet

## Getting started

```bash
yarn
yarn test
```

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
