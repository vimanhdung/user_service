gen:
	protoc -I=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go