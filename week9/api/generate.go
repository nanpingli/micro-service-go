package api

//go:generate protoc -I. -I$GOPATH/src --go_out=plugins=grpc:. --go_opt=paths=source_relative protocol/protocol.proto
