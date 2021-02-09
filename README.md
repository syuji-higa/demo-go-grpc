gRPC Demo for Golang (Client + Server)

```sh
# Compile for proto
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    sample/sample.proto

# Start Server
$ go run sample_server/main.go

# Run client
$ go run sample_client/main.go
```
