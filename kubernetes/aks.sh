PIP_RESOURCE_GROUP=K8S
AKS_RESOURCE_GROUP=K8S
AKS_CLUSTER_NAME=Million-Singer
# Do not change anything below this line
CLIENT_ID=$(az aks show --resource-group $AKS_RESOURCE_GROUP --name $AKS_CLUSTER_NAME --query "servicePrincipalProfile.clientId" --output tsv)
SUB_ID=$(az account show --query "id" --output tsv)
az role assignment create --assignee $CLIENT_ID --role "Network Contributor" --scope /subscriptions/$SUB_ID/resourceGroups/$PIP_RESOURCE_GROUP

# https://medium.com/microsoftazure/how-to-perform-role-assignments-on-azure-resources-from-an-azure-devops-pipeline-c9f4dc10d0a4