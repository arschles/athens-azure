#!/bin/bash

docker login quay.io \
-u ${CENSUSAI_DOCKER_USERNAME} \
-p ${CENSUSAI_DOCKER_PASSWORD}
docker push ${CENSUSAI_IMAGE}
