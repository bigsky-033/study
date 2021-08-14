#!/bin/bash

source ./00_user_info.sh

az sql server create \
  --resource-group $AZ_RESOURCE_GROUP \
  --name $AZ_DATABASE_NAME \
  --location $AZ_LOCATION \
  --admin-user $AZ_SQL_SERVER_USERNAME \
  --admin-password $AZ_SQL_SERVER_PASSWORD |
  jq
