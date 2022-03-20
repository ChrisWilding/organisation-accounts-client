# Organisation Accounts Client

This repository is an implementation of the Form3 take home exercise, details of which can be found [here](https://github.com/form3tech-oss/interview-accountapi).

My name is Chris Wilding. I'm fairly new to Go, I have experimented with it a few times but have not used it consistently or professionally.

## Prerequisites

1. [Docker](https://docs.docker.com/get-docker/)
1. [Go 1.18](https://go.dev/dl/)

## Setup

```sh
$ git clone git@github.com:ChrisWilding/organisation-accounts-client.git
$ cd organisation-accounts-client
```

## How To

### Test

```sh
$ docker compose up
```

The tests will run automatically. To run them again -

```sh
$ go test ./... -v
```
