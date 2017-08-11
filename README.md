# Dev-Nilson Client

## Install GO

https://golang.org/doc/install#install

## Profile Add to your bash_profile:

    export GOPATH=$HOME/go
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOPATH/bin

## Organize your repositories as follows:

    $GOPATH/src/github.com/devnilson/nilson-cli

## Dependencies

    go get github.com/urfave/cli

## Installation

From root of nilson-cli

    go install

Then you can run commands

    nilson-cli