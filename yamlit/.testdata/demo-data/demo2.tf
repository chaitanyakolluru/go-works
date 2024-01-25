
provider azurerm {
features {}
}

module "storageaccount_demo" {
     source                       = "git@ghe.aa.com:AA/terraform.git//azure-modules/storage_account"
     storage_account_name         = "storageaccountnamedemo"
     resource_group_name          = "aa-aot-sais-nonprod-eastus-rg"
     account_tier                 = "Standard"
     account_kind                 = "StorageV2"
     account_replication_type     = "grs"
     access_tier                  = "Hot"
     network_rules_default_action = "Allow"
     network_rules_bypass         = ["AzureServices"]
     enable_https_traffic_only    = true
     blob_container = {
           container1 = {
                 container_name        = "democontainer1",
                 storage_account_name  = "storageaccountnamedemo",
                 container_access_type = "container"
            },
            container2 = {
                  container_name        = "democontainer2",
                  storage_account_name  = "storageaccountnamedemo",
                  container_access_type = "container"
             },
      }
}

module "appservice_demo" {
  source                   = "git@ghe.aa.com:AA/terraform.git//azure-modules/appservice"
  appservice_plan_name     = "appserv-chai-plan"
  appservice_plan_tier     = "Standard"
  appservice_plan_size     = "S1"
  appservice_plan_kind     = "Linux"
  resource_group_name      = "aa-aot-sais-nonprod-eastus-rg"
  appservice_plan_reserved = true
  appservice = {
    "appservice-1": {
      "appservice_name": "appserv-chai-1"
    },
    "appservice-2": {
      "appservice_name": "appserv-chai-2"
    }
  }
}
