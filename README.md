# OPBundlerAutoConfig

## Overview

OPBundlerAutoConfig is a Go-based utility tool designed to automate and simplify the configuration of Ethereum Bundler and Optimism's devnet. This project was inspired by the challenges developers face when setting up a local Ethereum bundler with Optimism's devnet, as discussed in this [GitHub issue](https://github.com/ethereum-optimism/optimism/issues/6450#issuecomment-1705973174).

## Features

- **Automated Setup**: One-click setup for both Bundler and Optimism devnet.
- **ChainID Validation**: Customizable ChainID validation to adapt to specific network settings.
- **Port Adjustment**: Easy port configurations for both Bundler and devnet.
- **Funds Management**: Automated funding of necessary addresses to avoid "Insufficient funds" errors.

## Prerequisites

- Go 1.x
- Docker
- Ethereum Node
- Optimism Network

## Installation and Setup

After cloning the repository, navigate to the project directory:

```bash
cd OPBundlerAutoConfig
```

Run the following command to build and install the project:

```bash
make all
```

This will build the project and install the binary into `/usr/local/bin`.

## Usage

### Running the Program

Once the binary is installed, you can run the program. Here's an example:

```bash
OPBundlerAutoConfig --bundlerPath=/path/to/bundler/config.json --optimismPath=/path/to/optimism/config.json
```

Replace `/path/to/bundler/config.json` and `/path/to/optimism/config.json` with the actual paths to your Bundler and Optimism configuration files.

## Troubleshooting

- For port-related issues, check the port configurations in the Bundler and devnet settings.
- Before you dive into using this repository, make sure that both the opdevnet and bundler are set up and running smoothly. You can find guidance on how to run opdevnet here[https://github.com/ethereum-optimism/optimism/issues/6976#issuecomment-1690628412]. After you've got them up and running for the first time, remember to shut them down to free up the 8545, 9545, and 3000 ports.

## Contributing

Feel free to open issues and pull requests to improve the middleware and make the integration between Bundler and Optimism devnet more seamless.

## License

[MIT](LICENSE)
