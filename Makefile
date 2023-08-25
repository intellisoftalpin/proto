regen:
	protoc --proto_path=$(shell pwd) --go_out=. --go-grpc_out=. wallet/*.proto
	go mod tidy

clean:
	rm -r proto-gen

.PHONY: regen clean
