#!/bin/bash

source ./00_user_info.sh

az group create \
  --name $AZ_RESOURCE_GROUP \
  --location $AZ_LOCATION |
  jq

# Response Sample

#{
#  "id": "/subscriptions/**********/resourceGroups/database-workshop",
#  "location": "koreacentral",
#  "managedBy": null,
#  "name": "database-workshop",
#  "properties": {
#    "provisioningState": "Succeeded"
#  },
#  "tags": null,
#  "type": "Microsoft.Resources/resourceGroups"
#}
