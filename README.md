
## Installation
```bash
brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get google.golang.org/grpc
```

## Generate API File
`*.proto` to `*.pb.go`
```bash
cd proto
protoc --go_out=plugins=grpc:. *.proto
```

## Conenct
Only one `main` can be declared in same file and the same package name.

#### Pure Case:

```bash
# client/
go run client.go
# server/
go run server.go
```

#### Streaming Case:

```bash
# client/stream
go run clientStream.go
# server/stream
go run serverStream.go
```


#### Pure Case with TLS:
1. Use `{GOROOT}\src\crypto\tls\generate_cert.go` to generate `cert.pem` and `key.pem`
2. Import them

