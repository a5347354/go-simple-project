# grpc

## generate code by proto
```bash
protoc -I . hello.proto --go_out=plugins=grpc:.
```