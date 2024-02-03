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
	go build -o bin/worker ./cmd/worker/main.go

build-cli:
	go build -o bin/cli ./cmd/cli/main.go

test:
	docker-compose -f docker-compose.test.yml up -d
	while ! nc -z 127.0.0.1 27017; do sleep 1; echo "Waiting for port..."; done;
	go test -v ./pkg/...
	docker-compose -f docker-compose.test.yml down

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