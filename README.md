# VeChain Thor

A general purpose blockchain highly compatible with Ethereum's ecosystem.

This is the first implementation written in golang.

[![Go](https://img.shields.io/badge/golang-%3E%3D1.13-orange.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/vechain/thor)](https://goreportcard.com/report/github.com/vechain/thor)
[![Travis](https://travis-ci.org/vechain/thor.svg?branch=master)](https://travis-ci.org/vechain/thor)
[![License](https://img.shields.io/badge/License-LGPL%20v3-blue.svg)](https://github.com/vechain/thor/blob/master/LICENSE)
&nbsp;&nbsp; [![TG](https://img.shields.io/badge/chat-on%20telegram-blue)](https://t.me/VeChainDevCommunity)

## Table of contents

* [Installation](#installation)
    * [Requirements](#requirements)
    * [Getting the source](#getting-the-source)
    * [Dependency management](#dependency-management)
    * [Building](#building)
* [Running Thor](#running-thor)
    * [Sub-commands](#sub-commands)
* [Docker](#docker)
* [Explorers](#explorers)
* [Faucet](#testnet-faucet)
* [RESTful API](#api)
* [Acknowledgement](#acknowledgement)
* [Contributing](#contributing)

## Installation

### Requirements

Thor requires `Go` 1.13+ and `C` compiler to build. To install `Go`, follow this [link](https://golang.org/doc/install). 

### Getting the source

Clone the Thor repo:

```
git clone https://github.com/vechain/thor.git
cd thor
```

### Dependency management

Simply run:
```
make dep
```

If you keep getting network error, it is suggested to use [Go Module Proxy](https://golang.org/cmd/go/#hdr-Module_proxy_protocol). [https://proxy.golang.org/](https://proxy.golang.org/) is one option.

### Building

To build the main app `thor`, just run

```
make
```

or build the full suite:

```
make all
```

If no error reported, all built executable binaries will appear in folder *bin*.

## Running Thor

Connect to VeChain's mainnet:

```
bin/thor --network main
```


Connect to VeChain's testnet:

```
bin/thor --network test
```

or startup a custom network
```
bin/thor --network <custom-net-genesis.json>
```
exmaple genesis config file can be found at [genesis/example.json](https://raw.githubusercontent.com/vechain/thor/master/genesis/example.json).


To find out usages of all command line options:

```
bin/thor -h
```

- `--network value`             the network to join (main|test) or path to genesis file
- `--data-dir value`            directory for block-chain databases
- `--cache value`               megabytes of ram allocated to internal caching (default: 2048)
- `--beneficiary value`         address for block rewards
- `--target-gas-limit value`    target block gas limit (adaptive if set to 0) (default: 0)
- `--api-addr value`            API service listening address (default: "localhost:8669")
- `--api-cors value`            comma separated list of domains from which to accept cross origin requests to API
- `--api-timeout value`         API request timeout value in milliseconds (default: 10000)
- `--api-call-gas-limit value`  limit contract call gas (default: 50000000)
- `--api-backtrace-limit value` limit the distance between 'position' and best block for subscriptions APIs (default: 1000)
- `--verbosity value`           log verbosity (0-9) (default: 3)
- `--max-peers value`           maximum number of P2P network peers (P2P network disabled if set to 0) (default: 25)
- `--p2p-port value`            P2P network listening port (default: 11235)
- `--nat value`                 port mapping mechanism (any|none|upnp|pmp|extip:<IP>) (default: "none")
- `--bootnode value`            comma separated list of bootnode IDs
- `--skip-logs`                 skip writing event|transfer logs (/logs API will be disabled)
- `--pprof`                     turn on go-pprof
- `--disable-pruner`            disable state pruner to keep all history
- `--help, -h`                  show help
- `--version, -v`               print the version

### Sub-commands

- `solo`                client runs in solo mode for test & dev

```
bin/thor solo --on-demand               # create new block when there is pending transaction
bin/thor solo --persist                 # save blockchain data to disk(default to memory)
bin/thor solo --persist --on-demand     # two options can work together
```

- `master-key`          master key management

```
# print the master address
bin/thor master-key

# export master key to keystore
bin/thor master-key --export > keystore.json


# import master key from keystore
cat keystore.json | bin/thor master-key --import
```

## Docker

Docker is one quick way for running a vechain node:

```
docker run -d\
  -v {path-to-your-data-directory}/.org.vechain.thor:/root/.org.vechain.thor\
  -p 127.0.0.1:8669:8669 -p 11235:11235 -p 11235:11235/udp\
  --name thor-node vechain/thor --network test
```

Do not forget to add the `--api-addr 0.0.0.0:8669` flag if you want other containers and/or hosts to have access to the RESTful API. `Thor`binds to `localhost` by default and it will not accept requests outside the container itself without the flag.

The [Dockerfile](https://raw.githubusercontent.com/vechain/thor/master/Dockerfile) is designed to build the last release of the source code and will publish docker images to [dockerhub](https://hub.docker.com/r/vechain/thor/) by release, feel free to fork and build Dockerfile for your own purpose.

## Explorers

- [VeChain Explorer (Official)](https://explore.vechain.org)
- [VeChainStats](https://vechainstats.com/) 
- [VechainThorScan](https://vechainthorscan.com/)
- [Vexplorer](https://vexplorer.io/)
- [Insight](https://insight.vecha.in/)

## Testnet faucet

- [faucet.vecha.in](https://faucet.vecha.in) by *VeChain Foundation*

## API

Once `thor` started, online *OpenAPI* doc can be accessed in your browser. e.g. http://localhost:8669/ by default.

[![Thorest](https://raw.githubusercontent.com/vechain/thor/master/thorest.png)](http://localhost:8669/)

## Acknowledgement

A Special shout out to following projects:

- [Ethereum](https://github.com/ethereum)

- [Swagger](https://github.com/swagger-api)

## Contributing

Thanks you so much for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes!

Please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base.

### Forking Thor
When you "Fork" the project, GitHub will make a copy of the project that is entirely yours; it lives in your namespace, and you can push to it.

### Getting ready for a pull request
Please check the following:

- Code must be adhere to the official Go Formatting guidelines.
- Get the branch up to date, by merging in any recent changes from the master branch.

### Making the pull request
- On the GitHub site, go to "Code". Then click the green "Compare and Review" button. Your branch is probably in the "Example Comparisons" list, so click on it. If not, select it for the "compare" branch.
- Make sure you are comparing your new branch to master. It probably won't be, since the front page is the latest release branch, rather than master now. So click the base branch and change it to master.
- Press Create Pull Request button.
- Give a brief title.
- Explain the major changes you are asking to be code reviewed. Often it is useful to open a second tab in your browser where you can look through the diff yourself to remind yourself of all the changes you have made.

## License

VeChain Thor is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.html), also included
in *LICENSE* file in repository.
