#!/usr/bin/env bash

start_rpcsplitter() {
    git clone https://github.com/chronicleprotocol/oracle-suite.git
    cd oracle-suite
    go mod download

    timeout=2
    endpoints="$SMOCKER,$SMOCKER,http://10.255.255.1"
    (go run cmd/rpc-splitter/*.go run --log.format json --eth-rpc $endpoints -v debug -t $timeout)&
    sleep 3
    cd ..
}

start_rpcsplitter

#export DEBUG=true
printenv

go test -v -parallel 1 -cpu 1 ./...
