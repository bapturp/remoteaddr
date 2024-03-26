ARG GO_VERSION=1.22.1

# STAGE 1: build the executable
FROM golang:${GO_VERSION}-alpine as builder
# RUN apk add --no-cache git ca-certificates

# add user here since addgroup and adduser are not available in scratch
RUN addgroup -S remoteaddr && \
    adduser -G remoteaddr -S remoteaddr

WORKDIR /usr/src/remoteaddr
COPY . .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -v -o /usr/local/bin/remoteaddr

# STAGE 2: build the final container
FROM scratch AS final
LABEL maintainer="bapturp"
COPY --from=builder /usr/local/bin/remoteaddr /remoteaddr

# copy user from builder
COPY --from=builder /etc/passwd /etc/passwd

# copy ca-certificates.crt from builder
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER remoteaddr

EXPOSE 8080

ENTRYPOINT ["/remoteaddr"]
