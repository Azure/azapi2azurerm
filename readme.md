# azapi2azurerm

## Introduction
This tool is used to migrate resources from terraform `azapi` provider to `azurerm` provider.

## How to setup?
1. Clone this repo to local.
2. `go install` under project directory.
   
## Command Usage
```
PS C:\Users\henglu\go\src\github.com\Azure\azapi2azurerm> azapi2azurerm.exe            
Usage: azapi2azurerm [--version] [--help] <command> [<args>]

Available commands are:
    migrate    Migrate azapi resources to azurerm resources in current working directory
    plan       Show azapi resources which can migrate to azurerm resources in current working directory
    version    Displays the version of the migration tool
```

1. Run `azapi2azurerm plan` under your terraform working directory, 
   it will list all resources that can be migrated from `azapi` provider to `azurerm` provider.
   The Terraform addresses listed in file `azapi2azurerm.ignore` will be ignored during migration.
```
2022/01/25 14:34:46 [INFO] searching azapi_resource & azapi_update_resource...
2022/01/25 14:34:55 [INFO]

The tool will perform the following actions:

The following resources will be migrated:
azapi_resource.test2 will be replaced with azurerm_automation_account
azapi_update_resource.test will be replaced with azurerm_automation_account

The following resources can't be migrated:
azapi_resource.test: input properties not supported: [], output properties not supported: [identity.principalId, identity.type, identity.tenantId]

The following resources will be ignored in migration:
   ```
2. Run `azapi2azurerm migrate` under your terraform working directory, 
   it will migrate above resources from `azapi` provider to `azurerm` provider, 
   both terraform configuration and state.
   The Terraform addresses listed in file `azapi2azurerm.ignore` will be ignored during migration.
   
## Examples
There're some examples to show the migration results.
1. [case1 - basic](https://github.com/Azure/azapi2azurerm/tree/master/examples/case1%20-%20basic)
2. [case2 - for_each](https://github.com/Azure/azapi2azurerm/tree/master/examples/case2%20-%20for_each)
3. [case3 - nested block](https://github.com/Azure/azapi2azurerm/tree/master/examples/case3%20-%20nested%20block)
4. [case4 - count](https://github.com/Azure/azapi2azurerm/tree/master/examples/case4%20-%20count)
5. [case5 - nested block patch](https://github.com/Azure/azapi2azurerm/tree/master/examples/case5%20-%20nested%20block%20patch)
6. [case6 - meta argument](https://github.com/Azure/azapi2azurerm/tree/master/examples/case6%20-%20meta%20arguments)
7. [case7 - ignore](https://github.com/Azure/azapi2azurerm/tree/master/examples/case7%20-%20ignore)
   
## Features
- [x] Support resource `azapi_resource` migration
- [x] Support resource `azapi_update_resource` migration
- [x] Support meta-argument `for_each`
- [x] Support meta-argument `count`
- [x] Support meta-argument `depends_on`, `lifecycle` and `provisioner`
- [x] Support dependency injection in array and primitive value.
- [x] Support dependency injection in Map and other complicated struct value.
- [x] Support user input when there're multiple/none `azurerm` resource match for the resource id
- [x] Support migration based on `azurerm` provider's property coverage
- [x] Support ignore terraform addresses listed in file `azapi2azurerm.ignore`
- [ ] Support data source `azapi_resource` migration.

## Known limitations
1. References to local variables can't be migrated.
2. Usage of `dynamic` can't be migrated.
3. Update resource used to manage CMK can't be migrated.

## Credits

We wish to thank HashiCorp for the use of some MPLv2-licensed code from their open source project [terraform-plugin-sdk](https://github.com/hashicorp/terraform-plugin-sdk).