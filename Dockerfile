FROM golang:1.22 as builder
WORKDIR /build
COPY src/ .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -v -o remoteaddr

FROM scratch
COPY --from=builder /build/remoteaddr remoteaddr
EXPOSE 8080
ENTRYPOINT ["/remoteaddr"]
