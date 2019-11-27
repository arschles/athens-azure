AZURE_ATHENS_RESOURCE_GROUP ?= athens
AZURE_ATHENS_CONTAINER_NAME ?= athens
IMAGE_TAG ?= v0.3.0
AZURE_ATHENS_DNS_NAME ?= athens

GIT_SHA?=$(shell git rev-parse --short HEAD)


deploy:
	AZURE_ATHENS_RESOURCE_GROUP=${AZURE_ATHENS_RESOURCE_GROUP} \
	AZURE_ATHENS_CONTAINER_NAME=${AZURE_ATHENS_CONTAINER_NAME} \
	AZURE_ATHENS_CONTAINER_NAME=${AZURE_ATHENS_CONTAINER_NAME} \
	IMAGE_TAG=${IMAGE_TAG} \
	AZURE_ATHENS_DNS_NAME=${AZURE_ATHENS_DNS_NAME} \
	./aci.sh

build-flanders:
	go build -o flanders ./cmd/flanders

.PHONY: install-flanders
install-flanders:
	go install ./cmd/flanders

# ----
# lathens
####
LATHENS_DOCKER_TAG?=${GIT_SHA}
LATHENS_DOCKER_IMAGE=arschles/lathens:${LATHENS_DOCKER_TAG}
.PHONY: lathens
lathens:
	go build -o bin/lathens ./lathens

lathens-docker-build:
	docker build -t ${LATHENS_DOCKER_IMAGE} -f Dockerfile.lathens .

lathens-docker-push: lathens-docker-build
	docker push ${LATHENS_DOCKER_IMAGE}
