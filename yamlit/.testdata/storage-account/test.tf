provider azurerm {
features {}
}
module "storage_account-mod" {
	source = "git@ghe.aa.com:AA/terraform.git//azure-modules/storage_account?ref=storage-account-v1.1.2"
	account_tier = "Standard"
	enable_static_website = true
	error_404_document = "error.html"
	index_document = "index.html"
	storage_account_name = "demochaisa2"
	access_tier = "Hot"
	account_kind = "StorageV2"
	account_replication_type = "grs"
	enable_https_traffic_only = true
	network_rules_bypass = [
		"AzureServices"
	]
	resource_group_name = "aa-aot-sais-nonprod-eastus-rg"
}

