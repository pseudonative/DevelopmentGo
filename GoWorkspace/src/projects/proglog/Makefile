compile:
	protoc api/v1/*.proto \
		--plugin=protoc-gen-go=/Users/jeremy/go/bin/protoc-gen-go \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.
	
test:
	go test -race ./...

