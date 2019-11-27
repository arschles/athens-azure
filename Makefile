AZURE_ATHENS_RESOURCE_GROUP ?= athens
AZURE_ATHENS_CONTAINER_NAME ?= athens
IMAGE_TAG ?= v0.3.0
AZURE_ATHENS_DNS_NAME ?= athens

GIT_SHA?=$(shell git rev-parse --short HEAD)

.PHONY: git-sha
git-sha:
	@echo ${GIT_SHA}

##### THIS IS DEPRECATED
deploy:
	AZURE_ATHENS_RESOURCE_GROUP=${AZURE_ATHENS_RESOURCE_GROUP} \
	AZURE_ATHENS_CONTAINER_NAME=${AZURE_ATHENS_CONTAINER_NAME} \
	AZURE_ATHENS_CONTAINER_NAME=${AZURE_ATHENS_CONTAINER_NAME} \
	IMAGE_TAG=${IMAGE_TAG} \
	AZURE_ATHENS_DNS_NAME=${AZURE_ATHENS_DNS_NAME} \
	./aci.sh
#####

build-flanders:
	go build -o flanders ./cmd/flanders

.PHONY: install-flanders
install-flanders:
	go install ./cmd/flanders

#####
# lathens
#####
LATHENS_DOCKER_TAG?=${GIT_SHA}
LATHENS_DOCKER_IMAGE=arschles/lathens:${LATHENS_DOCKER_TAG}
.PHONY: lathens
lathens:
	go build -o bin/lathens ./lathens

.PHONY: lathens-docker-build
lathens-docker-build:
	docker build -t ${LATHENS_DOCKER_IMAGE} -f Dockerfile.lathens .

.PHONY: lathens-docker-push
lathens-docker-push: lathens-docker-build
	docker push ${LATHENS_DOCKER_IMAGE}

#####
# crathens
#####
CRATHENS_DOCKER_TAG?=${GIT_SHA}
CRATHENS_DOCKER_IMAGE=arschles/crathens:${CRATHENS_DOCKER_TAG}

.PHONY: crathens
crathens:
	go build -o bin/crathens ./crathens

.PHONY: crathens-docker-build
crathens-docker-build:
	docker build -t ${CRATHENS_DOCKER_IMAGE} -f Dockerfile.crathens .

.PHONY: crathens-docker-push
crathens-docker-push: crathens-docker-build
	docker push ${CRATHENS_DOCKER_IMAGE}

.PHONY: auto-tfvars
auto-tfvars:
	@echo "lathens-image-tag = \"${LATHENS_DOCKER_TAG}\"" > tags.auto.tfvars
