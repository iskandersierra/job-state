build:
	go build -o bin/job-state-api.exe ./cmd/api/.
	pwsh -c "Copy-Item '.env' -Destination './bin/'"

dev:
	air

run:
	go run ./cmd/api/.

fmt:
	gofmt -s -w .

gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/job-state.proto
