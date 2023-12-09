compile:
	protoc api/v1/*.proto \
		--plugin=protoc-gen-go=/Users/jeremy/go/bin/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=/Users/jeremy/go/bin/protoc-gen-go-grpc \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
	
test:
	go test -race ./...

