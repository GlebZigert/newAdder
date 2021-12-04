https://www.youtube.com/watch?v=z-mHhobE0Pw

protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto

go build -v ./cmd/server

go build -v ./cmd/client

https://www.youtube.com/watch?v=_CjzNp67LVg