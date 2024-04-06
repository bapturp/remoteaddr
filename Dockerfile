ARG GO_VERSION=1.22.1

# Build the executable
FROM golang:${GO_VERSION}-alpine as builder

RUN addgroup -S remoteaddr && \
    adduser -G remoteaddr -S remoteaddr

WORKDIR /usr/src/remoteaddr

COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -v -o /usr/local/bin/remoteaddr

# Build the final container
FROM scratch AS final

LABEL maintainer="bapturp"

COPY --from=builder /usr/local/bin/remoteaddr /remoteaddr

COPY --from=builder /etc/passwd /etc/passwd

USER remoteaddr

EXPOSE 8080

ENTRYPOINT ["/remoteaddr"]
