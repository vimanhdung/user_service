gen:
	protoc -I=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

.PHONY: gen clean server client