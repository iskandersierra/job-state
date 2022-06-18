.PHONY: build
build:
	go build -o bin/job-state-api.exe ./cmd/api/.
	pwsh -c "Copy-Item '.env' -Destination './bin/'"

.PHONY: dev
dev:
	air

.PHONY: run
run:
	go run ./cmd/api/.

.PHONY: repl
repl:
	go run ./cmd/repl/.

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/job-state.proto
