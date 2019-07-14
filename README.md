# Whitelist [![GoDoc](https://godoc.org/github.com/thylong/whitelist?status.svg)](https://godoc.org/github.com/thylong/whitelist) [![Go Report Card](https://goreportcard.com/badge/github.com/thylong/whitelist)](https://goreportcard.com/report/github.com/thylong/whitelist)

Whitelist is a very simple IP whitelisting microservice acting as a reverse proxy.  Its purpose is educational, it helps to highlight the importance of using the right algorithm/datastructure in the right context.

Its storage can be either a [hashmap](https://en.wikipedia.org/wiki/Hash_table), a [list](https://en.wikipedia.org/wiki/Array_data_structure) or a [Radix trie](https://en.wikipedia.org/wiki/Radix_tree) depending on what performance you want have.

## How does the service work?

The service keeps in-memory a whitelist of IPs to let go through the reverse proxy.
It's possible to edit through HTTP calls that list without reloads.

2 HTTP servers are exposed for these interactions :

- The first one, on port `8080` is the reverse proxy that filters all incoming requests prior to fan out in the backend service.
- The second one, on port `8081` is a small API usable to insert/delete/lookup on the whitelist storage.

Storage can be either

### Features 

- Storage can be either a hashmap, a list or a Radix trie
- IP format validation
- Ipv4 and IPv6 are both supported
- 100% tests coverage
- Benchmarks
- Support TrueIP & X-Forwarded-For headers
- Hot whitelisting edition
- Speed/scalability oriented (use Radix tries by default)

### Schema

```sh
HTTP request    +-------------+       +-----------+
                |             |       |           |
         +----->+  Whitelist  +------>+  Backend  |
                |   service   |       |  service  |
                |             |       |           |
                +-------------+       +-----------+
```

### Interacting with the whitelist


```sh
# Insert an IP in the whitelist :
curl -X POST http://localhost:8081/ip/<ip_address>

# Lookup for an IP in the whitelist :
curl -X GET http://localhost:8081/ip/<ip_address>

# Delete an IP in the whitelist :
curl -X DELETE http://localhost:8081/ip/<ip_address>
```

### Using the proxy

All calls made to the port `8080` will be forwarded to the backend (dummy by default) after running through the internal
logic described in this schema :

```sh
      +---------------+          +----------------+   yes
      |               |          |                |          +-------------+
+----->  Detect the   +---------->  Is the IP     +----------> Forward to  |
      |  origin IP    |          |  whitelisted?  |          | the backend |
      |               |          |                |          +-------------+
      +---------------+          +-------+--------+
                                         |
                                         | no
                                         |
                                         |
                                  +------v--------+
                                  | Access Denied |
                                  +---------------+
```

## Benchmark results

The results of these tests highly depends on the platform it runs on. The following results are found on my machine:
```sh
# HTTP server
BenchmarkFilterWithRadixSingleIPWhitelist-4   	   10000	    152451 ns/op
BenchmarkFilterWithRadixLargeIPWhitelist-4    	   10000	    225666 ns/op
BenchmarkFilterWithListSingleIPWhitelist-4    	   20000	     98824 ns/op
BenchmarkFilterWithListLargeIPWhitelist-4     	    5000	    245444 ns/op
# Raw whitelisting functions
BenchmarkListInsert-4     	10000000	       172 ns/op
BenchmarkHMInsert-4       	100000000	        14.8 ns/op
BenchmarkRadixInsert-4    	50000000	        27.4 ns/op
BenchmarkListContain-4    	    2000	   1031950 ns/op
BenchmarkHMContain-4      	100000000	        14.8 ns/op
BenchmarkRadixContain-4   	20000000	        88.3 ns/op
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
curl --header "X-Forwarded-For: <ip_address>" http://localhost:8080/
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