# remoteaddr

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/bapturp/remoteaddr/build-and-push.yaml)


Simple webserver that replies request information as they are seen by the 
webserver.

Usefull to troubleshoot reverse proxies, load balancers and API Gateways.

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
