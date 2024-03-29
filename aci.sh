#!/bin/sh

set -eou pipefail

for LOCATION in "westus2" "eastus2" "westeurope" "southeastasia"
do
    echo "Deploying in ${LOCATION}"
    az container create \
    -g "${AZURE_ATHENS_RESOURCE_GROUP}" \
    -n "${AZURE_ATHENS_CONTAINER_NAME}-${LOCATION}" \
    --image gomods/athens:${IMAGE_TAG} \
    -e "ATHENS_STORAGE_TYPE=mongo" "ATHENS_MONGO_STORAGE_URL=${AZURE_ATHENS_MONGO_URL}" "ATHENS_GO_GET_WORKERS=5" "GO_ENV=development" \
    --ip-address=Public \
    --dns-name="${AZURE_ATHENS_DNS_NAME}-${LOCATION}" \
    --ports="3000" \
    --location=${LOCATION}
done
