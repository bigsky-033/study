#!/bin/bash

source ./00_user_info.sh

az sql db create \
  --resource-group $AZ_RESOURCE_GROUP \
  --name demo \
  --server $AZ_DATABASE_NAME |
  jq

# Response Sample

#{
#  "autoPauseDelay": null,
#  "catalogCollation": "SQL_Latin1_General_CP1_CI_AS",
#  "collation": "SQL_Latin1_General_CP1_CI_AS",
#  "createMode": null,
#  "creationDate": "2021-08-14T05:29:44.220000+00:00",
#  "currentBackupStorageRedundancy": "Geo",
#  "currentServiceObjectiveName": "GP_Gen5_2",
#  "currentSku": {
#    "capacity": 2,
#    "family": "Gen5",
#    "name": "GP_Gen5",
#    "size": null,
#    "tier": "GeneralPurpose"
#  },
#  "databaseId": "***",
#  "defaultSecondaryLocation": "koreasouth",
#  "earliestRestoreDate": null,
#  "edition": "GeneralPurpose",
#  "elasticPoolId": null,
#  "elasticPoolName": null,
#  "failoverGroupId": null,
#  "highAvailabilityReplicaCount": null,
#  "id": "/subscriptions/*****/resourceGroups/database-workshop/providers/Micro
#soft.Sql/servers/workshop-db/databases/demo",
#  "isInfraEncryptionEnabled": false,
#  "kind": "v12.0,user,vcore",
#  "ledgerOn": false,
#  "licenseType": "LicenseIncluded",
#  "location": "koreacentral",
#  "longTermRetentionBackupResourceId": null,
#  "maintenanceConfigurationId": "/subscriptions/*****/providers/Microsoft.Main
#tenance/publicMaintenanceConfigurations/SQL_Default",
#  "managedBy": null,
#  "maxLogSizeBytes": 68719476736,
#  "maxSizeBytes": 34359738368,
#  "minCapacity": null,
#  "name": "demo",
#  "pausedDate": null,
#  "readScale": "Disabled",
#  "recoverableDatabaseId": null,
#  "recoveryServicesRecoveryPointId": null,
#  "requestedBackupStorageRedundancy": "Geo",
#  "requestedServiceObjectiveName": "GP_Gen5_2",
#  "resourceGroup": "database-workshop",
#  "restorableDroppedDatabaseId": null,
#  "restorePointInTime": null,
#  "resumedDate": null,
#  "sampleName": null,
#  "secondaryType": null,
#  "sku": {
#    "capacity": 2,
#    "family": "Gen5",
#    "name": "GP_Gen5",
#    "size": null,
#    "tier": "GeneralPurpose"
#  },
#  "sourceDatabaseDeletionDate": null,
#  "sourceDatabaseId": null,
#  "status": "Online",
#  "tags": null,
#  "type": "Microsoft.Sql/servers/databases",
#  "zoneRedundant": false
#}
