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

## Notes

I have tried to follow some of the conventions I have seen in other SDK's or within the Go community.

* I have used the functional options pattern described by Dave Cheney
* Implemented a String utility function like the AWS Go SDK for optional string parameters
* Accepted an http.Client as a parameter to NewClient similar to go-github

Given more time I would have liked to implement logging similar to stripe-go which has a logging interface that can accept either a Logrus logger or Zap's SugaredLogger.

I would normally except tests to run independent of any other services and not make network calls (particullary in CI). In Java I would often use Test Containers and WireMock. However, the instructions explicitly state to the provided fake API. As I have re-used test data between tests that has resulted in some dependencies between tests and on the ordering of tests.

I have used the initially provided model.go as is, and have not added mappings for any additional fields returned by the fake API e.g. created_on and modified_on.
