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

docker-run-dev:
	docker-compose up mongo selenium -d

docker-build-worker:
	docker build \
		-f docker/worker/Dockerfile \
		--build-arg TARGETARCH=$(ARCH) \
		-t $(WORKER_IMAGE):$(VERSION) \
		-t $(WORKER_IMAGE):latest \
		.

docker-build-cli:
	docker build \
		-f docker/cli/Dockerfile \
		--build-arg TARGETARCH=$(ARCH) \
		-t $(CLI_IMAGE):$(VERSION) \
		-t ($CLI_IMAGE):latest \
		.