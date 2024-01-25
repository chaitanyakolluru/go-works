provider azurerm {
features {}
}
module "appservice-mod" {
	source = "git@ghe.aa.com:AA/terraform.git//azure-modules/appservice?ref=appservice-v2.2.1"
	appservice_plan_kind = "Linux"
	appservice_plan_name = "appserv-chai-plan"
	appservice_plan_reserved = true
	appservice_plan_size = "S1"
	appservice_plan_tier = "Standard"
	resource_group_name = "aa-aot-sais-nonprod-eastus-rg"
	appservice = {
		"appservice-1": {
			"appservice_name": "appserv-chai-1"
		},
		"appservice-2": {
			"appservice_name": "appserv-chai-2"
		},
		"appservice-3": {
			"appservice_name": "appserv-chai-3"
		}
	}
}

