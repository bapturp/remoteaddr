# remoteaddr

Simple webserver that replies request information as they are seen by the 
webserver.

Usefull to troubleshoot reverse proxy and load balancer especially in 
kubernetes.

```
➜  ~ curl http://127.0.0.1:8080/
2024-03-12 12:03:16
GET / HTTP/1.1
Remote address: 192.168.127.1:53651
Host: 127.0.0.1:8080
---- HEADERS ----
Accept: */*
User-Agent: curl/8.4.0
```

## Usage

```sh
docker run --rm -p 8080:8080 ghcr.io/bapturp/remoteaddr:latest
curl http://127.0.0.1:8080/
```
