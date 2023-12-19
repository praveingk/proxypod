SW_VERSION ?= latest
IMAGE_ORG ?= praveingk/proxypod

IMAGE_TAG_BASE ?= $(IMAGE_ORG)
IMG ?= $(IMAGE_TAG_BASE):$(SW_VERSION)
build:
	@echo "Start go build phase"
	go build -o proxypod ./proxypod-process/proxypod.go
	go build -o proxypod-k8s proxypod-k8s.go

lint:  ; $(info running linters...)
	@golangci-lint run --config=./.golangci.yaml ./...

build-image:
	docker build --build-arg SW_VERSION="$(SW_VERSION)" -t ${IMG} .

push-image:
	docker push ${IMG}