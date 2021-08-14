#!/bin/bash

source ./00_user_info.sh

az sql server firewall-rule create \
  --resource-group $AZ_RESOURCE_GROUP \
  --name $AZ_DATABASE_NAME-database-allow-local-ip \
  --server $AZ_DATABASE_NAME \
  --start-ip-address $AZ_LOCAL_IP_ADDRESS \
  --end-ip-address $AZ_LOCAL_IP_ADDRESS |
  jq

# Response Sample

#{
#  "endIpAddress": "**.**.**.**",
#  "id": "/subscriptions/**/resourceGroups/database-workshop/providers/Micr
#osoft.Sql/servers/workshop-db/firewallRules/workshop-db-database-allow-local-ip",
#  "name": "workshop-db-database-allow-local-ip",
#  "resourceGroup": "database-workshop",
#  "startIpAddress": "*.**.**.**",
#  "type": "Microsoft.Sql/servers/firewallRules"
#}
