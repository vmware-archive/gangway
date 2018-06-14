
gangway
=======

_(noun): An opening in the bulwark of the ship to allow passengers to board or leave the ship._

An application that can be used to easily enable authentication flows via OIDC for a kubernetes cluster.

## Deployment

Instructions for deploying gangway for common cloud providers can be found [here](docs/README.md).

## API-Server flags

gangway requires that the Kubernetes API server is configured for OIDC:

https://kubernetes.io/docs/admin/authentication/#configuring-the-api-server

```
kube-apiserver
...
--oidc-issuer-url "https://example.auth0.com"
--oidc-client-id 3YM4ue8MoXgBkvCIHh00000000000
--oidc-username-claim sub
--oidc-groups-claim "https://example.auth0.com/groups"
```

## Build

Requirements for building

- Go (built with 1.10)
- [go-bindata](https://github.com/jteeuwen/go-bindata)
- [dep](https://github.com/golang/dep)

A Makefile is provided for building tasks. The options are as follows

Getting started is as simple as:
```
$ go get github.com/heptiolabs/gangway
$ cd $GOPATH/src/github.com/heptiolabs/gangway
$ make setup
$ make
```
