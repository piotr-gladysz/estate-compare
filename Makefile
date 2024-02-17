IMAGE ?= estate-compare
VERSION ?= 0.0.1
ARCH  ?= $(shell $(GO) env GOARCH)

WORKER_IMAGE ?= $(IMAGE)/worker
CLI_IMAGE ?= $(IMAGE)/cli

gen-grpc:
	docker run \
		-v .:/app \
		--workdir /app \
		--rm \
		bufbuild/buf generate proto

build-worker:
	mkdir -p bin
	go build -o bin/worker ./cmd/worker/main.go

build-cli:
	mkdir -p bin
	go build -o bin/cli ./cmd/cli/main.go

build-test-plugins:
	mkdir -p bin/plugins/condition/test
	tinygo build -o bin/plugins/condition/test/valid.wasm -scheduler=none -target=wasi ./plugin/condition/test/valid/main.go
	tinygo build -o bin/plugins/condition/test/invalid-export.wasm -scheduler=none -target=wasi ./plugin/condition/test/invalid-export/main.go
	tinygo build -o bin/plugins/condition/test/invalid-input.wasm -scheduler=none -target=wasi ./plugin/condition/test/invalid-input/main.go
	tinygo build -o bin/plugins/condition/test/invalid-output.wasm -scheduler=none -target=wasi ./plugin/condition/test/invalid-output/main.go

build-plugins:
	mkdir -p bin/plugins/condition
	tinygo build -o bin/plugins/condition/example.wasm -scheduler=none -target=wasi ./plugin/condition/example/main.go

test:
#	docker-compose -f docker-compose.test.yml up -d
#	while ! nc -z 127.0.0.1 27017; do sleep 1; echo "Waiting for port..."; done;
	go test -v -count=1 ./pkg/...
#	docker-compose -f docker-compose.test.yml down

benchmark:
	go test -bench=. -benchtime=10s ./pkg/...

benchmark-memory:
	go test -benchmem -bench=. -benchtime=10s ./pkg/...

docker-run-dev:
	docker-compose -f docker-compose.dev.yml up -d

docker-build-worker:
	docker build \
		-f dockerfile/worker/Dockerfile \
		--build-arg TARGETARCH=$(ARCH) \
		-t $(WORKER_IMAGE):$(VERSION) \
		-t $(WORKER_IMAGE):latest \
		.

docker-build-cli:
	docker build \
		-f dockerfile/cli/Dockerfile \
		--build-arg TARGETARCH=$(ARCH) \
		-t $(CLI_IMAGE):$(VERSION) \
		-t ($CLI_IMAGE):latest \
		.