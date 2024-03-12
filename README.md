# remoteaddr

remoteaddr is a simple webserver that replies the remote address (client ip address).

## Usage

```sh
docker run --rm -p 8080:8080 ghcr.io/bapturp/remoteaddr:latest
curl http://127.0.0.1:8080/
```
