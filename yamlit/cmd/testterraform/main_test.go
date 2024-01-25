package testterraform

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sync"
	"testing"
)

func TestNowProcessAndTest(t *testing.T) {
	cwd, _ := os.Getwd()

	// test by mocking testcondition and terraform.tfstate files on the local filesystem.
	testConditionContent := `
ap-yamlit-test:
    kind: linux
    size: S1
    tier: Standard
    reserved: true
apps-yamlit-test:
`

	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"testcondition.yaml", []byte(testConditionContent), 0644); err != nil {
		t.Error(err)
	}

	terraformStateContent := `
	{
		"version": 4,
		"terraform_version": "0.12.25",
		"serial": 4,
		"lineage": "6199cefb-9544-687c-9dee-69ec2e97c702",
		"outputs": {},
		"resources": [
		  {
			"module": "module.appservice-mod",
			"mode": "data",
			"type": "azurerm_app_service_plan",
			"name": "appserviceplan",
			"each": "list",
			"provider": "provider.azurerm",
			"instances": []
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "data",
			"type": "azurerm_client_config",
			"name": "current",
			"provider": "provider.azurerm",
			"instances": [
			  {
				"schema_version": 0,
				"attributes": {
				  "client_id": "3bcbd362-c13f-448a-b478-1f61eaff336f",
				  "id": "2020-07-06 17:22:56.6550763 +0000 UTC",
				  "object_id": "dfb940f8-5a50-4e9e-81ed-ebba2d36d62a",
				  "subscription_id": "0de836dc-fea8-4805-b578-5d70b530d15e",
				  "tenant_id": "49793faf-eb3f-4d99-a0cf-aef7cce79dc1",
				  "timeouts": null
				}
			  }
			]
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "data",
			"type": "azurerm_key_vault",
			"name": "kv",
			"each": "list",
			"provider": "provider.azurerm",
			"instances": []
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "data",
			"type": "azurerm_resource_group",
			"name": "rg_info",
			"provider": "provider.azurerm",
			"instances": [
			  {
				"schema_version": 0,
				"attributes": {
				  "id": "/subscriptions/0de836dc-fea8-4805-b578-5d70b530d15e/resourceGroups/aa-aot-sais-nonprod-eastus-rg",
				  "location": "eastus",
				  "name": "aa-aot-sais-nonprod-eastus-rg",
				  "tags": {
					"aa-app-id": "188020",
					"aa-application": "sais",
					"aa-costcenter": "0900/1894",
					"aa-location": "eastus",
					"aa-sdlc-environment": "nonprod",
					"aa-security": "reserved",
					"aa-vertical": "aot"
				  },
				  "timeouts": null
				}
			  }
			]
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "managed",
			"type": "azurerm_app_service",
			"name": "appservice",
			"each": "map",
			"provider": "provider.azurerm",
			"instances": [
			  {
				"index_key": "appservice-1",
				"schema_version": 0,
				"attributes": {
				  "app_service_plan_id": "/subscriptions/0de836dc-fea8-4805-b578-5d70b530d15e/resourceGroups/aa-aot-sais-nonprod-eastus-rg/providers/Microsoft.Web/serverfarms/ap-yamlit-test",
				  "app_settings": {},
				  "auth_settings": [
					{
					  "active_directory": [],
					  "additional_login_params": {},
					  "allowed_external_redirect_urls": [],
					  "default_provider": "",
					  "enabled": false,
					  "facebook": [],
					  "google": [],
					  "issuer": "",
					  "microsoft": [],
					  "runtime_version": "",
					  "token_refresh_extension_hours": 0,
					  "token_store_enabled": false,
					  "twitter": [],
					  "unauthenticated_client_action": ""
					}
				  ],
				  "backup": [],
				  "client_affinity_enabled": true,
				  "client_cert_enabled": false,
				  "connection_string": [],
				  "default_site_hostname": "apps-yamlit-test.azurewebsites.net",
				  "enabled": true,
				  "https_only": false,
				  "id": "/subscriptions/0de836dc-fea8-4805-b578-5d70b530d15e/resourceGroups/aa-aot-sais-nonprod-eastus-rg/providers/Microsoft.Web/sites/apps-yamlit-test",
				  "identity": [
					{
					  "identity_ids": null,
					  "principal_id": "1a12a0c1-a0c7-4206-ab13-8f183a59f57e",
					  "tenant_id": "49793faf-eb3f-4d99-a0cf-aef7cce79dc1",
					  "type": "SystemAssigned"
					}
				  ],
				  "location": "eastus",
				  "logs": [
					{
					  "application_logs": [
						{
						  "azure_blob_storage": []
						}
					  ],
					  "http_logs": [
						{
						  "azure_blob_storage": [],
						  "file_system": []
						}
					  ]
					}
				  ],
				  "name": "apps-yamlit-test",
				  "outbound_ip_addresses": "40.71.11.140,40.121.154.115,13.82.228.43,40.121.158.167,40.117.44.182",
				  "possible_outbound_ip_addresses": "40.71.11.140,40.121.154.115,13.82.228.43,40.121.158.167,40.117.44.182,168.61.50.107,40.121.80.139,40.117.44.94,23.96.53.166,40.121.152.91",
				  "resource_group_name": "aa-aot-sais-nonprod-eastus-rg",
				  "site_config": [
					{
					  "always_on": false,
					  "app_command_line": "",
					  "auto_swap_slot_name": "",
					  "cors": [],
					  "default_documents": [
						"Default.htm",
						"Default.html",
						"Default.asp",
						"index.htm",
						"index.html",
						"iisstart.htm",
						"default.aspx",
						"index.php",
						"hostingstart.html"
					  ],
					  "dotnet_framework_version": "v4.0",
					  "ftps_state": "AllAllowed",
					  "http2_enabled": false,
					  "ip_restriction": [],
					  "java_container": "",
					  "java_container_version": "",
					  "java_version": "",
					  "linux_fx_version": "",
					  "local_mysql_enabled": false,
					  "managed_pipeline_mode": "Integrated",
					  "min_tls_version": "1.2",
					  "php_version": "",
					  "python_version": "",
					  "remote_debugging_enabled": false,
					  "remote_debugging_version": "",
					  "scm_type": "None",
					  "use_32_bit_worker_process": true,
					  "websockets_enabled": false,
					  "windows_fx_version": ""
					}
				  ],
				  "site_credential": [
					{
					  "password": "it7wJqRriFiyS2fHSM6pwjpnf3Jflcm0WxbslEMW28lW604cXm8h8ZtwM8s3",
					  "username": "$apps-yamlit-test"
					}
				  ],
				  "source_control": [
					{
					  "branch": "master",
					  "repo_url": ""
					}
				  ],
				  "storage_account": [],
				  "tags": {
					"aa-app-id": "188020",
					"aa-application": "sais",
					"aa-costcenter": "0900/1894",
					"aa-sdlc-environment": "nonprod",
					"aa-security": "reserved",
					"aa-vertical": "aot"
				  },
				  "timeouts": null
				},
				"private": "eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjoxODAwMDAwMDAwMDAwLCJkZWxldGUiOjE4MDAwMDAwMDAwMDAsInJlYWQiOjMwMDAwMDAwMDAwMCwidXBkYXRlIjoxODAwMDAwMDAwMDAwfX0=",
				"dependencies": [
				  "module.appservice-mod.azurerm_app_service_plan.appsrvplan"
				]
			  }
			]
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "managed",
			"type": "azurerm_app_service_plan",
			"name": "appsrvplan",
			"each": "list",
			"provider": "provider.azurerm",
			"instances": [
			  {
				"index_key": 0,
				"schema_version": 0,
				"attributes": {
				  "app_service_environment_id": "",
				  "id": "/subscriptions/0de836dc-fea8-4805-b578-5d70b530d15e/resourceGroups/aa-aot-sais-nonprod-eastus-rg/providers/Microsoft.Web/serverfarms/ap-yamlit-test",
				  "is_xenon": false,
				  "kind": "linux",
				  "location": "eastus",
				  "maximum_elastic_worker_count": 1,
				  "maximum_number_of_workers": 10,
				  "name": "ap-yamlit-test",
				  "per_site_scaling": false,
				  "reserved": true,
				  "resource_group_name": "aa-aot-sais-nonprod-eastus-rg",
				  "sku": [
					{
					  "capacity": 1,
					  "size": "S1",
					  "tier": "Standard"
					}
				  ],
				  "tags": {
					"aa-app-id": "188020",
					"aa-application": "sais",
					"aa-costcenter": "0900/1894",
					"aa-sdlc-environment": "nonprod",
					"aa-security": "reserved",
					"aa-vertical": "aot"
				  },
				  "timeouts": null
				},
				"private": "eyJlMmJmYjczMC1lY2FhLTExZTYtOGY4OC0zNDM2M2JjN2M0YzAiOnsiY3JlYXRlIjozNjAwMDAwMDAwMDAwLCJkZWxldGUiOjM2MDAwMDAwMDAwMDAsInJlYWQiOjMwMDAwMDAwMDAwMCwidXBkYXRlIjozNjAwMDAwMDAwMDAwfX0="
			  }
			]
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "managed",
			"type": "azurerm_app_service_slot",
			"name": "appservice_slot",
			"each": "map",
			"provider": "provider.azurerm",
			"instances": []
		  },
		  {
			"module": "module.appservice-mod",
			"mode": "managed",
			"type": "azurerm_key_vault_access_policy",
			"name": "azaccesspolicy",
			"each": "list",
			"provider": "provider.azurerm",
			"instances": []
		  }
		]
	  }	  
	`

	if err := os.Mkdir(cwd+string(os.PathSeparator)+".temp", 0755); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+".temp"+string(os.PathSeparator)+"terraform.tfstate", []byte(terraformStateContent), 0644); err != nil {
		t.Error(err)
	}

	outGot := captureStdout(NowProcessAndTest)
	outWant := "TESTS PASSED"
	if m, _ := regexp.MatchString(outWant, outGot); !m {
		t.Errorf("Test command output is incorrect, got: %s, want: %s", outGot, outWant)
	}

	if err := os.Remove(cwd + string(os.PathSeparator) + "testcondition.yaml"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(cwd + string(os.PathSeparator) + ".temp"); err != nil {
		t.Fatal(err)
	}

}

func captureStdout(f func(string, string)) string {
	cwd, _ := os.Getwd()

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = w
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	wg.Wait()

	f(cwd, "testcondition.yaml")

	w.Close()

	return <-out

}
