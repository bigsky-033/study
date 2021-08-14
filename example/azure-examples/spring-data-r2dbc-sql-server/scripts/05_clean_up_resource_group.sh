#!/bin/bash

source ./00_user_info.sh

az group delete \
    --name $AZ_RESOURCE_GROUP \
    --yes
