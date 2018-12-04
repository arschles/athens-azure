#!/bin/sh

set -eou pipefail

az container create \
-g "${AZURE_ATHENS_RESOURCE_GROUP}" \
-n "${AZURE_ATHENS_CONTAINER_NAME}" \
--image gomods/athens:v0.2.0 \
-e "ATHENS_STORAGE_TYPE=mongo" "ATHENS_MONGO_STORAGE_URL=${AZURE_ATHENS_MONGO_URL}" \
--ip-address=Public \
--dns-name="${AZURE_ATHENS_DNS_NAME}" \
--ports="3000"


az container show -n "${AZURE_ATHENS_CONTAINER_NAME}" -g "${AZURE_ATHENS_RESOURCE_GROUP}"
az container logs -n "${AZURE_ATHENS_CONTAINER_NAME}" -g "${AZURE_ATHENS_RESOURCE_GROUP}"
