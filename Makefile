export GODEBUG=http2debug=2

.PHONY:proto

proto:
	cd grpc && protoc --go_out=plugins=grpc:. proto/*.proto

clean:
	rm -f ./bin/client
	rm -f ./bin/server

build:clean
	go build  -v -o ./bin/client ./client
	go build  -v -o ./bin/server ./server

run-s:
	./bin/server

run-c:
	./bin/client

run-server:
	go run server/main.go 

run-client:
	go run client/main.go 