# concrete-template

A simple template for an application-specific rollup built with the [Concrete framework](https://github.com/concrete-eth/concrete-geth).

```bash
yarn
yarn test
```

## Running a devnet

Run a normal Optimism Bedrock devnet replacing op-geth for `github.com/concrete-eth/concrete-template-geth:latest` in `ops-bedrock/Dockerfile.l2`.

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