# Setzer MCD

Query USD price feeds

## Usage

```
Usage: setzer <command> [<args>]
   or: setzer <command> --help

Commands:

   help            Print help about setzer or one of its subcommands
   pairs           List all supported pairs
   price           Show price(s) for a given asset or pair
   sources         Show price sources for a given asset or pair
   test            Test all price feeds
```

## Installation

Dependencies:

* GNU [bc](https://www.gnu.org/software/bc/)
* [curl](https://curl.haxx.se/download.html)
* GNU [datamash](https://www.gnu.org/software/datamash/)
* GNU `date`
* [jshon](http://kmkeen.com/jshon/)
* GNU `timeout`
* [htmlq](https://github.com/mgdm/htmlq)

Install via make:

* `make link` -  link setzer into `/usr/local`
* `make install` -  copy setzer into `/usr/local`
* `make uninstall` -  remove setzer from `/usr/local`

## Configuration

* `SETZER_CACHE` - Cache directory (default: ~/.setzer)
* `SETZER_CACHE_EXPIRY` - Cache expiry (default: 60) seconds
* `SETZER_TIMEOUT` - HTTP request timeout (default: 10) seconds

## wstETH pair requirement

Due to process of pulling details from mainnet for getting price information.
You need to set `ETH_RPC_URL` environemnt variable. By default it will point to `http://127.0.0.1:8545`.

Example of usage: 

```bash
export ETH_RPC_URL="https://mainnet.infura.io/v3/fac98e56ea7e49608825dfc726fab703"
```

### Fx/Exchangerates API Key
Since latest changes in Exchangerates API, now it requires API key.
To set API Key for this exchange you can use `EXCHANGERATES_API_KEY` env variable. 

Example:

```bash
$ EXCHANGERATES_API_KEY=your_api_key setzer fx krwusd
```

### E2E tests
E2E tests for setzer are written in Go language and relies on [Smocker](https://smocker.dev) for API manipulation.

#### How to setup tests environment
You have to install Docker on your machine first.
Then you will have to start [smocker](https://smocker.dev) container.

```bash
$ docker run -d \
  --restart=always \
  -p 8080:8080 \
  -p 8081:8081 \
  --name smocker \
  thiht/smocker
```

Next step will be to build setzer E2E docker container:

```bash
$ docker build -t setzer -f e2e/Dockerfile .
```

Run newly created container:

```bash
$ docker run -i --rm --link smocker setzer
```

If you need to write tests or want to continuesly run them while doing something you might use docker interactive mode.

```bash
$ docker run -it --rm -v $(pwd):/app --link smocker setzer /bin/bash
```

It will start docker in interactove mode and you will be able to run E2E tests using command: 

```bash
$ go test -v -parallel 1 -cpu 1 ./...
```