# remoteaddr

![GitHub Actions Workflow Test Status](https://github.com/bapturp/remoteaddr/actions/workflows/test.yaml/badge.svg) ![GitHub Actions Workflow Release Status](https://github.com/bapturp/remoteaddr/actions/workflows/release.yaml/badge.svg)

Simple web server that replies to requests by echoing the information it receives from the client.

Useful to troubleshoot reverse proxies, load balancers, service mesh and API Gateways.

```
➜  ~ curl http://127.0.0.1:8080/
2024-03-14 22:59:57
GET / HTTP/1.1
Remote address: 192.168.127.1:30693
Host:           127.0.0.1:8080
Headers:
  User-Agent: curl/8.4.0
  Accept: */*
```

## Usage

```sh
docker run --rm -p 8080:8080 ghcr.io/bapturp/remoteaddr:latest
curl http://127.0.0.1:8080/
```
