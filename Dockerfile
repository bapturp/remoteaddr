ARG GO_VERSION=1.22.1

# STAGE 1: build the executable
FROM golang:${GO_VERSION}-alpine as builder

# add user here since addgroup and adduser are not available in scratch
RUN addgroup -S remoteaddr && adduser -S -u 10000 -g remoteaddr remoteaddr

WORKDIR /src
COPY . .
RUN go mod download && go mod verify
RUN go test -timeout 1m
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o remoteaddr

# STAGE 2: build the final container
FROM scratch AS final
LABEL maintainer="bapturp"
COPY --from=builder /src/remoteaddr remoteaddr

# copy user from builder
COPY --from=builder /etc/passwd /etc/passwd

USER remoteaddr

EXPOSE 8080
ENTRYPOINT ["/remoteaddr"]
