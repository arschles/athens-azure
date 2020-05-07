#!/bin/bash

# this is for one-time setup of the Terraform state
# azure blob store. taken from:
# https://docs.microsoft.com/en-us/azure/developer/terraform/store-state-in-azure-storage

RESOURCE_GROUP_NAME=athens
STORAGE_ACCOUNT_NAME=athensci
CONTAINER_NAME=tstate

# If you don't already have a resource group, 
# uncomment this command. It will create your resource
# group you specified in RESOURCE_GROUP_NAME
# in eastus
# az group create --name $RESOURCE_GROUP_NAME --location eastus

# Create storage account
az storage account create --resource-group $RESOURCE_GROUP_NAME --name $STORAGE_ACCOUNT_NAME --sku Standard_LRS --encryption-services blob

# Get storage account key
ACCOUNT_KEY=$(az storage account keys list --resource-group $RESOURCE_GROUP_NAME --account-name $STORAGE_ACCOUNT_NAME --query [0].value -o tsv)

# Create blob container
az storage container create --name $CONTAINER_NAME --account-name $STORAGE_ACCOUNT_NAME --account-key $ACCOUNT_KEY

echo "storage_account_name: $STORAGE_ACCOUNT_NAME"
echo "container_name: $CONTAINER_NAME"
echo "access_key: $ACCOUNT_KEY"