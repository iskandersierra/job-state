gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/job-state.proto

build:
	go build -o bin/ ./cmd/api/.