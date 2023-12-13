
all: client server

protoc:
	@echo "Generating Go files"
	cd src/proto && protoc --go_out=plugins=grpc:. *.proto

server: protoc
	@echo "Building server"
	go build -o server \
		./src/server
## // github.com/pahanini/go-grpc-bidirectional-streaming-example/src/server

client: protoc
	@echo "Building client"
	go build -o client \
		./src/client
##// github.com/pahanini/go-grpc-bidirectional-streaming-example/src/client

clean:
	go clean github.com/pahanini/go-grpc-bidirectional-streaming-example/...
	rm -f server client

.PHONY: client server protoc
