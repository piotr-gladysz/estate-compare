
gen-grpc:
	buf generate proto

build-worker:
	go build -o bin/worker ./cmd/worker/main.go

build-cli:
	go build -o bin/cli ./cmd/cli/main.go