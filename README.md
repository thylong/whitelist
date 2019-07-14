# Whitelist [![GoDoc](https://godoc.org/github.com/thylong/whitelist?status.svg)](https://godoc.org/github.com/thylong/whitelist) [![Go Report Card](https://goreportcard.com/badge/github.com/thylong/whitelist)](https://goreportcard.com/report/github.com/thylong/whitelist)

Whitelist is a very simple IP whitelisting microservice acting as a reverse proxy.  Its purpose is to highlight the importance of using the right algorithm/datastructure in the right context.

## How does the service work?



## Installation

```sh
go get -u github.com/thylong/whitelist
```

## Usage

To get a sense of the difference of the time complexity of each algorithm, you can simply run the benchmark tests.
These tests compare the basic operations possible on the service with small and large sets of whitelisted IPs.

If you want to experience the service under real conditions, replace the **dummy** service content in the docker-compose
by the service of your choice.

Then start all services :

```sh
docker-compose up
```

Then seed the whitelist :

```sh
docker-compose exec -it tester ash
# Once attached to the tester container (required to be in Docker network space)
curl -X POST http://localhost:8081/ip/<ip_address>
```

And then query the backend service :

```sh
# Still attached to the tester container
curl --header "X-Forwarded-For: <ip_address>" -X POST http://localhost:8080/
```

## Running tests

```sh
# Run all tests
make test

# Run only functionnal tests
make test-func

# Run only benchmark tests
make test-bench

# Run only end-to-end tests (requires Docker + docker-compose)
make test-e2e
```

## Contribute

Contributions are welcome and actually encouraged !
Don't hesitate to open issues, pull requests, we'll take a look as soon as possible to improve the project.

### How to contribute?

Every little effort counts.
Feel free to suggest documentation improvements, fix typos or add new features, even copy a subset of the project !

Just make sure :
- You're aware of the Licence boundaries (as small as they actually are).
- For new features: you test the code you'd like to introduce
- If you'd like to introduce new dependencies, open first an issue

### It's your first open source contribution?

Awesome, welcome !
Have a look to this guide [here](https://github.com/github/opensource.guide/blob/master/CONTRIBUTING.md) first.

## License

[MIT License](https://github.com/BackMarket/whitelist/blob/master/LICENSE)

## Credits

`ip.go` & `ip_test.go` were written by [tomasen](https://github.com/tomasen) in its package [realip](https://github.com/tomasen/realip)