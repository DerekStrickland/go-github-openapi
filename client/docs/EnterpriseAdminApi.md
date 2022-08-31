# \EnterpriseAdminApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise) | **Post** /enterprises/{enterprise}/actions/runners/{runner_id}/labels | Add custom labels to a self-hosted runner for an enterprise
[**EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterprise) | **Put** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id} | Add organization access to a self-hosted runner group in an enterprise
[**EnterpriseAdminAddSelfHostedRunnerToGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminAddSelfHostedRunnerToGroupForEnterprise) | **Put** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id} | Add a self-hosted runner to a group for an enterprise
[**EnterpriseAdminCreateRegistrationTokenForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminCreateRegistrationTokenForEnterprise) | **Post** /enterprises/{enterprise}/actions/runners/registration-token | Create a registration token for an enterprise
[**EnterpriseAdminCreateRemoveTokenForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminCreateRemoveTokenForEnterprise) | **Post** /enterprises/{enterprise}/actions/runners/remove-token | Create a remove token for an enterprise
[**EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise) | **Post** /enterprises/{enterprise}/actions/runner-groups | Create a self-hosted runner group for an enterprise
[**EnterpriseAdminDeleteScimGroupFromEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminDeleteScimGroupFromEnterprise) | **Delete** /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id} | Delete a SCIM group from an enterprise
[**EnterpriseAdminDeleteSelfHostedRunnerFromEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminDeleteSelfHostedRunnerFromEnterprise) | **Delete** /enterprises/{enterprise}/actions/runners/{runner_id} | Delete a self-hosted runner from an enterprise
[**EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterprise) | **Delete** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id} | Delete a self-hosted runner group from an enterprise
[**EnterpriseAdminDeleteUserFromEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminDeleteUserFromEnterprise) | **Delete** /scim/v2/enterprises/{enterprise}/Users/{scim_user_id} | Delete a SCIM user from an enterprise
[**EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterprise) | **Delete** /enterprises/{enterprise}/actions/permissions/organizations/{org_id} | Disable a selected organization for GitHub Actions in an enterprise
[**EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterprise) | **Put** /enterprises/{enterprise}/actions/permissions/organizations/{org_id} | Enable a selected organization for GitHub Actions in an enterprise
[**EnterpriseAdminGetAllowedActionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminGetAllowedActionsEnterprise) | **Get** /enterprises/{enterprise}/actions/permissions/selected-actions | Get allowed actions and reusable workflows for an enterprise
[**EnterpriseAdminGetAuditLog**](EnterpriseAdminApi.md#EnterpriseAdminGetAuditLog) | **Get** /enterprises/{enterprise}/audit-log | Get the audit log for an enterprise
[**EnterpriseAdminGetConsumedLicenses**](EnterpriseAdminApi.md#EnterpriseAdminGetConsumedLicenses) | **Get** /enterprises/{enterprise}/consumed-licenses | List enterprise consumed licenses
[**EnterpriseAdminGetGithubActionsPermissionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminGetGithubActionsPermissionsEnterprise) | **Get** /enterprises/{enterprise}/actions/permissions | Get GitHub Actions permissions for an enterprise
[**EnterpriseAdminGetLicenseSyncStatus**](EnterpriseAdminApi.md#EnterpriseAdminGetLicenseSyncStatus) | **Get** /enterprises/{enterprise}/license-sync-status | Get a license sync status
[**EnterpriseAdminGetProvisioningInformationForEnterpriseGroup**](EnterpriseAdminApi.md#EnterpriseAdminGetProvisioningInformationForEnterpriseGroup) | **Get** /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id} | Get SCIM provisioning information for an enterprise group
[**EnterpriseAdminGetProvisioningInformationForEnterpriseUser**](EnterpriseAdminApi.md#EnterpriseAdminGetProvisioningInformationForEnterpriseUser) | **Get** /scim/v2/enterprises/{enterprise}/Users/{scim_user_id} | Get SCIM provisioning information for an enterprise user
[**EnterpriseAdminGetSelfHostedRunnerForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminGetSelfHostedRunnerForEnterprise) | **Get** /enterprises/{enterprise}/actions/runners/{runner_id} | Get a self-hosted runner for an enterprise
[**EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise) | **Get** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id} | Get a self-hosted runner group for an enterprise
[**EnterpriseAdminGetServerStatistics**](EnterpriseAdminApi.md#EnterpriseAdminGetServerStatistics) | **Get** /enterprise-installation/{enterprise_or_org}/server-statistics | Get GitHub Enterprise Server statistics
[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise) | **Get** /enterprises/{enterprise}/actions/runners/{runner_id}/labels | List labels for a self-hosted runner for an enterprise
[**EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise) | **Get** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations | List organization access to a self-hosted runner group in an enterprise
[**EnterpriseAdminListProvisionedGroupsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListProvisionedGroupsEnterprise) | **Get** /scim/v2/enterprises/{enterprise}/Groups | List provisioned SCIM groups for an enterprise
[**EnterpriseAdminListProvisionedIdentitiesEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListProvisionedIdentitiesEnterprise) | **Get** /scim/v2/enterprises/{enterprise}/Users | List SCIM provisioned identities for an enterprise
[**EnterpriseAdminListRunnerApplicationsForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListRunnerApplicationsForEnterprise) | **Get** /enterprises/{enterprise}/actions/runners/downloads | List runner applications for an enterprise
[**EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise) | **Get** /enterprises/{enterprise}/actions/permissions/organizations | List selected organizations enabled for GitHub Actions in an enterprise
[**EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise) | **Get** /enterprises/{enterprise}/actions/runner-groups | List self-hosted runner groups for an enterprise
[**EnterpriseAdminListSelfHostedRunnersForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListSelfHostedRunnersForEnterprise) | **Get** /enterprises/{enterprise}/actions/runners | List self-hosted runners for an enterprise
[**EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise) | **Get** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners | List self-hosted runners in a group for an enterprise
[**EnterpriseAdminProvisionAndInviteEnterpriseGroup**](EnterpriseAdminApi.md#EnterpriseAdminProvisionAndInviteEnterpriseGroup) | **Post** /scim/v2/enterprises/{enterprise}/Groups | Provision a SCIM enterprise group and invite users
[**EnterpriseAdminProvisionAndInviteEnterpriseUser**](EnterpriseAdminApi.md#EnterpriseAdminProvisionAndInviteEnterpriseUser) | **Post** /scim/v2/enterprises/{enterprise}/Users | Provision and invite a SCIM enterprise user
[**EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise) | **Delete** /enterprises/{enterprise}/actions/runners/{runner_id}/labels | Remove all custom labels from a self-hosted runner for an enterprise
[**EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise) | **Delete** /enterprises/{enterprise}/actions/runners/{runner_id}/labels/{name} | Remove a custom label from a self-hosted runner for an enterprise
[**EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterprise) | **Delete** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id} | Remove organization access to a self-hosted runner group in an enterprise
[**EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterprise) | **Delete** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id} | Remove a self-hosted runner from a group for an enterprise
[**EnterpriseAdminSetAllowedActionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminSetAllowedActionsEnterprise) | **Put** /enterprises/{enterprise}/actions/permissions/selected-actions | Set allowed actions and reusable workflows for an enterprise
[**EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise) | **Put** /enterprises/{enterprise}/actions/runners/{runner_id}/labels | Set custom labels for a self-hosted runner for an enterprise
[**EnterpriseAdminSetGithubActionsPermissionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminSetGithubActionsPermissionsEnterprise) | **Put** /enterprises/{enterprise}/actions/permissions | Set GitHub Actions permissions for an enterprise
[**EnterpriseAdminSetInformationForProvisionedEnterpriseGroup**](EnterpriseAdminApi.md#EnterpriseAdminSetInformationForProvisionedEnterpriseGroup) | **Put** /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id} | Set SCIM information for a provisioned enterprise group
[**EnterpriseAdminSetInformationForProvisionedEnterpriseUser**](EnterpriseAdminApi.md#EnterpriseAdminSetInformationForProvisionedEnterpriseUser) | **Put** /scim/v2/enterprises/{enterprise}/Users/{scim_user_id} | Set SCIM information for a provisioned enterprise user
[**EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterprise) | **Put** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations | Set organization access for a self-hosted runner group in an enterprise
[**EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterprise) | **Put** /enterprises/{enterprise}/actions/permissions/organizations | Set selected organizations enabled for GitHub Actions in an enterprise
[**EnterpriseAdminSetSelfHostedRunnersInGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminSetSelfHostedRunnersInGroupForEnterprise) | **Put** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners | Set self-hosted runners in a group for an enterprise
[**EnterpriseAdminUpdateAttributeForEnterpriseGroup**](EnterpriseAdminApi.md#EnterpriseAdminUpdateAttributeForEnterpriseGroup) | **Patch** /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id} | Update an attribute for a SCIM enterprise group
[**EnterpriseAdminUpdateAttributeForEnterpriseUser**](EnterpriseAdminApi.md#EnterpriseAdminUpdateAttributeForEnterpriseUser) | **Patch** /scim/v2/enterprises/{enterprise}/Users/{scim_user_id} | Update an attribute for a SCIM enterprise user
[**EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise**](EnterpriseAdminApi.md#EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise) | **Patch** /enterprises/{enterprise}/actions/runner-groups/{runner_group_id} | Update a self-hosted runner group for an enterprise



## EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise(ctx, enterprise, runnerId).EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest(enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest).Execute()

Add custom labels to a self-hosted runner for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest := *openapiclient.NewEnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest([]string{"Labels_example"}) // EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise(context.Background(), enterprise, runnerId).EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest(enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest** | [**EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest**](EnterpriseAdminAddCustomLabelsToSelfHostedRunnerForEnterpriseRequest.md) |  | 

### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterprise

> EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterprise(ctx, enterprise, runnerGroupId, orgId).Execute()

Add organization access to a self-hosted runner group in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    orgId := int32(56) // int32 | The unique identifier of the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterprise(context.Background(), enterprise, runnerGroupId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**orgId** | **int32** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminAddOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminAddSelfHostedRunnerToGroupForEnterprise

> EnterpriseAdminAddSelfHostedRunnerToGroupForEnterprise(ctx, enterprise, runnerGroupId, runnerId).Execute()

Add a self-hosted runner to a group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminAddSelfHostedRunnerToGroupForEnterprise(context.Background(), enterprise, runnerGroupId, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminAddSelfHostedRunnerToGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminAddSelfHostedRunnerToGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminCreateRegistrationTokenForEnterprise

> AuthenticationToken EnterpriseAdminCreateRegistrationTokenForEnterprise(ctx, enterprise).Execute()

Create a registration token for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminCreateRegistrationTokenForEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminCreateRegistrationTokenForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminCreateRegistrationTokenForEnterprise`: AuthenticationToken
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminCreateRegistrationTokenForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminCreateRegistrationTokenForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationToken**](AuthenticationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminCreateRemoveTokenForEnterprise

> AuthenticationToken EnterpriseAdminCreateRemoveTokenForEnterprise(ctx, enterprise).Execute()

Create a remove token for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminCreateRemoveTokenForEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminCreateRemoveTokenForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminCreateRemoveTokenForEnterprise`: AuthenticationToken
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminCreateRemoveTokenForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminCreateRemoveTokenForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationToken**](AuthenticationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise

> RunnerGroupsEnterprise EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise(ctx, enterprise).EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest(enterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest).Execute()

Create a self-hosted runner group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    enterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest := *openapiclient.NewEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest("Name_example") // EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise(context.Background(), enterprise).EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest(enterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise`: RunnerGroupsEnterprise
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminCreateSelfHostedRunnerGroupForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest** | [**EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest**](EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest.md) |  | 

### Return type

[**RunnerGroupsEnterprise**](RunnerGroupsEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminDeleteScimGroupFromEnterprise

> EnterpriseAdminDeleteScimGroupFromEnterprise(ctx, enterprise, scimGroupId).Execute()

Delete a SCIM group from an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimGroupId := "scimGroupId_example" // string | Identifier generated by the GitHub SCIM endpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminDeleteScimGroupFromEnterprise(context.Background(), enterprise, scimGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminDeleteScimGroupFromEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimGroupId** | **string** | Identifier generated by the GitHub SCIM endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminDeleteScimGroupFromEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminDeleteSelfHostedRunnerFromEnterprise

> EnterpriseAdminDeleteSelfHostedRunnerFromEnterprise(ctx, enterprise, runnerId).Execute()

Delete a self-hosted runner from an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminDeleteSelfHostedRunnerFromEnterprise(context.Background(), enterprise, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminDeleteSelfHostedRunnerFromEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminDeleteSelfHostedRunnerFromEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterprise

> EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterprise(ctx, enterprise, runnerGroupId).Execute()

Delete a self-hosted runner group from an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterprise(context.Background(), enterprise, runnerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminDeleteSelfHostedRunnerGroupFromEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminDeleteUserFromEnterprise

> EnterpriseAdminDeleteUserFromEnterprise(ctx, enterprise, scimUserId).Execute()

Delete a SCIM user from an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminDeleteUserFromEnterprise(context.Background(), enterprise, scimUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminDeleteUserFromEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminDeleteUserFromEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterprise

> EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterprise(ctx, enterprise, orgId).Execute()

Disable a selected organization for GitHub Actions in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    orgId := int32(56) // int32 | The unique identifier of the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterprise(context.Background(), enterprise, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminDisableSelectedOrganizationGithubActionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**orgId** | **int32** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminDisableSelectedOrganizationGithubActionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterprise

> EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterprise(ctx, enterprise, orgId).Execute()

Enable a selected organization for GitHub Actions in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    orgId := int32(56) // int32 | The unique identifier of the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterprise(context.Background(), enterprise, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminEnableSelectedOrganizationGithubActionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**orgId** | **int32** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminEnableSelectedOrganizationGithubActionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetAllowedActionsEnterprise

> SelectedActions EnterpriseAdminGetAllowedActionsEnterprise(ctx, enterprise).Execute()

Get allowed actions and reusable workflows for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetAllowedActionsEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetAllowedActionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetAllowedActionsEnterprise`: SelectedActions
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetAllowedActionsEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetAllowedActionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SelectedActions**](SelectedActions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetAuditLog

> []AuditLogEvent EnterpriseAdminGetAuditLog(ctx, enterprise).Phrase(phrase).Include(include).After(after).Before(before).Order(order).Page(page).PerPage(perPage).Execute()

Get the audit log for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    phrase := "phrase_example" // string | A search phrase. For more information, see [Searching the audit log](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization#searching-the-audit-log). (optional)
    include := "include_example" // string | The event types to include:  - `web` - returns web (non-Git) events. - `git` - returns Git events. - `all` - returns both web and Git events.  The default is `web`. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. (optional)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. (optional)
    order := "order_example" // string | The order of audit log events. To list newest events first, specify `desc`. To list oldest events first, specify `asc`.  The default is `desc`. (optional)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetAuditLog(context.Background(), enterprise).Phrase(phrase).Include(include).After(after).Before(before).Order(order).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetAuditLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetAuditLog`: []AuditLogEvent
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phrase** | **string** | A search phrase. For more information, see [Searching the audit log](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization#searching-the-audit-log). | 
 **include** | **string** | The event types to include:  - &#x60;web&#x60; - returns web (non-Git) events. - &#x60;git&#x60; - returns Git events. - &#x60;all&#x60; - returns both web and Git events.  The default is &#x60;web&#x60;. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. | 
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. | 
 **order** | **string** | The order of audit log events. To list newest events first, specify &#x60;desc&#x60;. To list oldest events first, specify &#x60;asc&#x60;.  The default is &#x60;desc&#x60;. | 
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**[]AuditLogEvent**](AuditLogEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetConsumedLicenses

> GetConsumedLicenses EnterpriseAdminGetConsumedLicenses(ctx, enterprise).PerPage(perPage).Page(page).Execute()

List enterprise consumed licenses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetConsumedLicenses(context.Background(), enterprise).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetConsumedLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetConsumedLicenses`: GetConsumedLicenses
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetConsumedLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetConsumedLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**GetConsumedLicenses**](GetConsumedLicenses.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetGithubActionsPermissionsEnterprise

> ActionsEnterprisePermissions EnterpriseAdminGetGithubActionsPermissionsEnterprise(ctx, enterprise).Execute()

Get GitHub Actions permissions for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetGithubActionsPermissionsEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetGithubActionsPermissionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetGithubActionsPermissionsEnterprise`: ActionsEnterprisePermissions
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetGithubActionsPermissionsEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetGithubActionsPermissionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsEnterprisePermissions**](ActionsEnterprisePermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetLicenseSyncStatus

> GetLicenseSyncStatus EnterpriseAdminGetLicenseSyncStatus(ctx, enterprise).Execute()

Get a license sync status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetLicenseSyncStatus(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetLicenseSyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetLicenseSyncStatus`: GetLicenseSyncStatus
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetLicenseSyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetLicenseSyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLicenseSyncStatus**](GetLicenseSyncStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetProvisioningInformationForEnterpriseGroup

> ScimEnterpriseGroup EnterpriseAdminGetProvisioningInformationForEnterpriseGroup(ctx, enterprise, scimGroupId).ExcludedAttributes(excludedAttributes).Execute()

Get SCIM provisioning information for an enterprise group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimGroupId := "scimGroupId_example" // string | Identifier generated by the GitHub SCIM endpoint.
    excludedAttributes := "excludedAttributes_example" // string | Attributes to exclude. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetProvisioningInformationForEnterpriseGroup(context.Background(), enterprise, scimGroupId).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetProvisioningInformationForEnterpriseGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetProvisioningInformationForEnterpriseGroup`: ScimEnterpriseGroup
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetProvisioningInformationForEnterpriseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimGroupId** | **string** | Identifier generated by the GitHub SCIM endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetProvisioningInformationForEnterpriseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludedAttributes** | **string** | Attributes to exclude. | 

### Return type

[**ScimEnterpriseGroup**](ScimEnterpriseGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetProvisioningInformationForEnterpriseUser

> ScimEnterpriseUser EnterpriseAdminGetProvisioningInformationForEnterpriseUser(ctx, enterprise, scimUserId).Execute()

Get SCIM provisioning information for an enterprise user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetProvisioningInformationForEnterpriseUser(context.Background(), enterprise, scimUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetProvisioningInformationForEnterpriseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetProvisioningInformationForEnterpriseUser`: ScimEnterpriseUser
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetProvisioningInformationForEnterpriseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetProvisioningInformationForEnterpriseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScimEnterpriseUser**](ScimEnterpriseUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetSelfHostedRunnerForEnterprise

> Runner EnterpriseAdminGetSelfHostedRunnerForEnterprise(ctx, enterprise, runnerId).Execute()

Get a self-hosted runner for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetSelfHostedRunnerForEnterprise(context.Background(), enterprise, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetSelfHostedRunnerForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetSelfHostedRunnerForEnterprise`: Runner
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetSelfHostedRunnerForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetSelfHostedRunnerForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Runner**](Runner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise

> RunnerGroupsEnterprise EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise(ctx, enterprise, runnerGroupId).Execute()

Get a self-hosted runner group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise(context.Background(), enterprise, runnerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise`: RunnerGroupsEnterprise
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetSelfHostedRunnerGroupForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetSelfHostedRunnerGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RunnerGroupsEnterprise**](RunnerGroupsEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminGetServerStatistics

> []ServerStatisticsInner EnterpriseAdminGetServerStatistics(ctx, enterpriseOrOrg).DateStart(dateStart).DateEnd(dateEnd).Execute()

Get GitHub Enterprise Server statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterpriseOrOrg := "enterpriseOrOrg_example" // string | The slug version of the enterprise name or the login of an organization.
    dateStart := "dateStart_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. (optional)
    dateEnd := "dateEnd_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminGetServerStatistics(context.Background(), enterpriseOrOrg).DateStart(dateStart).DateEnd(dateEnd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminGetServerStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetServerStatistics`: []ServerStatisticsInner
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminGetServerStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterpriseOrOrg** | **string** | The slug version of the enterprise name or the login of an organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminGetServerStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateStart** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. | 
 **dateEnd** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. | 

### Return type

[**[]ServerStatisticsInner**](ServerStatisticsInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise(ctx, enterprise, runnerId).Execute()

List labels for a self-hosted runner for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise(context.Background(), enterprise, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListLabelsForSelfHostedRunnerForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise

> EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise(ctx, enterprise, runnerGroupId).PerPage(perPage).Page(page).Execute()

List organization access to a self-hosted runner group in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise(context.Background(), enterprise, runnerGroupId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise`: EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response**](EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListProvisionedGroupsEnterprise

> ScimGroupListEnterprise EnterpriseAdminListProvisionedGroupsEnterprise(ctx, enterprise).StartIndex(startIndex).Count(count).Filter(filter).ExcludedAttributes(excludedAttributes).Execute()

List provisioned SCIM groups for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    startIndex := int32(56) // int32 | Used for pagination: the index of the first result to return. (optional)
    count := int32(56) // int32 | Used for pagination: the number of results to return. (optional)
    filter := "filter_example" // string | filter results (optional)
    excludedAttributes := "excludedAttributes_example" // string | attributes to exclude (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListProvisionedGroupsEnterprise(context.Background(), enterprise).StartIndex(startIndex).Count(count).Filter(filter).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListProvisionedGroupsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListProvisionedGroupsEnterprise`: ScimGroupListEnterprise
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListProvisionedGroupsEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListProvisionedGroupsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Used for pagination: the index of the first result to return. | 
 **count** | **int32** | Used for pagination: the number of results to return. | 
 **filter** | **string** | filter results | 
 **excludedAttributes** | **string** | attributes to exclude | 

### Return type

[**ScimGroupListEnterprise**](ScimGroupListEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListProvisionedIdentitiesEnterprise

> ScimUserListEnterprise EnterpriseAdminListProvisionedIdentitiesEnterprise(ctx, enterprise).StartIndex(startIndex).Count(count).Filter(filter).Execute()

List SCIM provisioned identities for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    startIndex := int32(56) // int32 | Used for pagination: the index of the first result to return. (optional)
    count := int32(56) // int32 | Used for pagination: the number of results to return. (optional)
    filter := "filter_example" // string | filter results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListProvisionedIdentitiesEnterprise(context.Background(), enterprise).StartIndex(startIndex).Count(count).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListProvisionedIdentitiesEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListProvisionedIdentitiesEnterprise`: ScimUserListEnterprise
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListProvisionedIdentitiesEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListProvisionedIdentitiesEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Used for pagination: the index of the first result to return. | 
 **count** | **int32** | Used for pagination: the number of results to return. | 
 **filter** | **string** | filter results | 

### Return type

[**ScimUserListEnterprise**](ScimUserListEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListRunnerApplicationsForEnterprise

> []RunnerApplication EnterpriseAdminListRunnerApplicationsForEnterprise(ctx, enterprise).Execute()

List runner applications for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListRunnerApplicationsForEnterprise(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListRunnerApplicationsForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListRunnerApplicationsForEnterprise`: []RunnerApplication
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListRunnerApplicationsForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListRunnerApplicationsForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RunnerApplication**](RunnerApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise

> EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise(ctx, enterprise).PerPage(perPage).Page(page).Execute()

List selected organizations enabled for GitHub Actions in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise(context.Background(), enterprise).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise`: EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response**](EnterpriseAdminListSelectedOrganizationsEnabledGithubActionsEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise

> EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise200Response EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise(ctx, enterprise).PerPage(perPage).Page(page).VisibleToOrganization(visibleToOrganization).Execute()

List self-hosted runner groups for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    visibleToOrganization := "visibleToOrganization_example" // string | Only return runner groups that are allowed to be used by this organization. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise(context.Background(), enterprise).PerPage(perPage).Page(page).VisibleToOrganization(visibleToOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise`: EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListSelfHostedRunnerGroupsForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **visibleToOrganization** | **string** | Only return runner groups that are allowed to be used by this organization. | 

### Return type

[**EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise200Response**](EnterpriseAdminListSelfHostedRunnerGroupsForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListSelfHostedRunnersForEnterprise

> EnterpriseAdminListSelfHostedRunnersForEnterprise200Response EnterpriseAdminListSelfHostedRunnersForEnterprise(ctx, enterprise).PerPage(perPage).Page(page).Execute()

List self-hosted runners for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnersForEnterprise(context.Background(), enterprise).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnersForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListSelfHostedRunnersForEnterprise`: EnterpriseAdminListSelfHostedRunnersForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnersForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListSelfHostedRunnersForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**EnterpriseAdminListSelfHostedRunnersForEnterprise200Response**](EnterpriseAdminListSelfHostedRunnersForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise

> EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise(ctx, enterprise, runnerGroupId).PerPage(perPage).Page(page).Execute()

List self-hosted runners in a group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise(context.Background(), enterprise, runnerGroupId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise`: EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminListSelfHostedRunnersInGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response**](EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminProvisionAndInviteEnterpriseGroup

> ScimEnterpriseGroup EnterpriseAdminProvisionAndInviteEnterpriseGroup(ctx, enterprise).EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(enterpriseAdminProvisionAndInviteEnterpriseGroupRequest).Execute()

Provision a SCIM enterprise group and invite users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    enterpriseAdminProvisionAndInviteEnterpriseGroupRequest := *openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest([]string{"Schemas_example"}, "DisplayName_example") // EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminProvisionAndInviteEnterpriseGroup(context.Background(), enterprise).EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(enterpriseAdminProvisionAndInviteEnterpriseGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminProvisionAndInviteEnterpriseGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminProvisionAndInviteEnterpriseGroup`: ScimEnterpriseGroup
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminProvisionAndInviteEnterpriseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enterpriseAdminProvisionAndInviteEnterpriseGroupRequest** | [**EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest**](EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest.md) |  | 

### Return type

[**ScimEnterpriseGroup**](ScimEnterpriseGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminProvisionAndInviteEnterpriseUser

> ScimEnterpriseUser EnterpriseAdminProvisionAndInviteEnterpriseUser(ctx, enterprise).EnterpriseAdminProvisionAndInviteEnterpriseUserRequest(enterpriseAdminProvisionAndInviteEnterpriseUserRequest).Execute()

Provision and invite a SCIM enterprise user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    enterpriseAdminProvisionAndInviteEnterpriseUserRequest := *openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequest([]string{"Schemas_example"}, "UserName_example", *openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName("GivenName_example", "FamilyName_example"), []openapiclient.EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner{*openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner("Value_example", "Type_example", false)}) // EnterpriseAdminProvisionAndInviteEnterpriseUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminProvisionAndInviteEnterpriseUser(context.Background(), enterprise).EnterpriseAdminProvisionAndInviteEnterpriseUserRequest(enterpriseAdminProvisionAndInviteEnterpriseUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminProvisionAndInviteEnterpriseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminProvisionAndInviteEnterpriseUser`: ScimEnterpriseUser
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminProvisionAndInviteEnterpriseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminProvisionAndInviteEnterpriseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enterpriseAdminProvisionAndInviteEnterpriseUserRequest** | [**EnterpriseAdminProvisionAndInviteEnterpriseUserRequest**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequest.md) |  | 

### Return type

[**ScimEnterpriseUser**](ScimEnterpriseUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise(ctx, enterprise, runnerId).Execute()

Remove all custom labels from a self-hosted runner for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise(context.Background(), enterprise, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminRemoveAllCustomLabelsFromSelfHostedRunnerForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise(ctx, enterprise, runnerId, name).Execute()

Remove a custom label from a self-hosted runner for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    name := "name_example" // string | The name of a self-hosted runner's custom label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise(context.Background(), enterprise, runnerId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 
**name** | **string** | The name of a self-hosted runner&#39;s custom label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminRemoveCustomLabelFromSelfHostedRunnerForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterprise

> EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterprise(ctx, enterprise, runnerGroupId, orgId).Execute()

Remove organization access to a self-hosted runner group in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    orgId := int32(56) // int32 | The unique identifier of the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterprise(context.Background(), enterprise, runnerGroupId, orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**orgId** | **int32** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminRemoveOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterprise

> EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterprise(ctx, enterprise, runnerGroupId, runnerId).Execute()

Remove a self-hosted runner from a group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterprise(context.Background(), enterprise, runnerGroupId, runnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminRemoveSelfHostedRunnerFromGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetAllowedActionsEnterprise

> EnterpriseAdminSetAllowedActionsEnterprise(ctx, enterprise).SelectedActions(selectedActions).Execute()

Set allowed actions and reusable workflows for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    selectedActions := *openapiclient.NewSelectedActions() // SelectedActions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetAllowedActionsEnterprise(context.Background(), enterprise).SelectedActions(selectedActions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetAllowedActionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetAllowedActionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedActions** | [**SelectedActions**](SelectedActions.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise

> EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise(ctx, enterprise, runnerId).EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest(enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest).Execute()

Set custom labels for a self-hosted runner for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerId := int32(56) // int32 | Unique identifier of the self-hosted runner.
    enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest([]string{"Labels_example"}) // EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise(context.Background(), enterprise, runnerId).EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest(enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise`: EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerId** | **int32** | Unique identifier of the self-hosted runner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest** | [**EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest**](EnterpriseAdminSetCustomLabelsForSelfHostedRunnerForEnterpriseRequest.md) |  | 

### Return type

[**EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response**](EnterpriseAdminListLabelsForSelfHostedRunnerForEnterprise200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetGithubActionsPermissionsEnterprise

> EnterpriseAdminSetGithubActionsPermissionsEnterprise(ctx, enterprise).EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(enterpriseAdminSetGithubActionsPermissionsEnterpriseRequest).Execute()

Set GitHub Actions permissions for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    enterpriseAdminSetGithubActionsPermissionsEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(openapiclient.enabled-organizations("all")) // EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetGithubActionsPermissionsEnterprise(context.Background(), enterprise).EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(enterpriseAdminSetGithubActionsPermissionsEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetGithubActionsPermissionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enterpriseAdminSetGithubActionsPermissionsEnterpriseRequest** | [**EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest**](EnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetInformationForProvisionedEnterpriseGroup

> ScimEnterpriseGroup EnterpriseAdminSetInformationForProvisionedEnterpriseGroup(ctx, enterprise, scimGroupId).EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(enterpriseAdminProvisionAndInviteEnterpriseGroupRequest).Execute()

Set SCIM information for a provisioned enterprise group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimGroupId := "scimGroupId_example" // string | Identifier generated by the GitHub SCIM endpoint.
    enterpriseAdminProvisionAndInviteEnterpriseGroupRequest := *openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest([]string{"Schemas_example"}, "DisplayName_example") // EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetInformationForProvisionedEnterpriseGroup(context.Background(), enterprise, scimGroupId).EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(enterpriseAdminProvisionAndInviteEnterpriseGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetInformationForProvisionedEnterpriseGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminSetInformationForProvisionedEnterpriseGroup`: ScimEnterpriseGroup
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminSetInformationForProvisionedEnterpriseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimGroupId** | **string** | Identifier generated by the GitHub SCIM endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetInformationForProvisionedEnterpriseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminProvisionAndInviteEnterpriseGroupRequest** | [**EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest**](EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest.md) |  | 

### Return type

[**ScimEnterpriseGroup**](ScimEnterpriseGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetInformationForProvisionedEnterpriseUser

> ScimEnterpriseUser EnterpriseAdminSetInformationForProvisionedEnterpriseUser(ctx, enterprise, scimUserId).EnterpriseAdminProvisionAndInviteEnterpriseUserRequest(enterpriseAdminProvisionAndInviteEnterpriseUserRequest).Execute()

Set SCIM information for a provisioned enterprise user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.
    enterpriseAdminProvisionAndInviteEnterpriseUserRequest := *openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequest([]string{"Schemas_example"}, "UserName_example", *openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName("GivenName_example", "FamilyName_example"), []openapiclient.EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner{*openapiclient.NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner("Value_example", "Type_example", false)}) // EnterpriseAdminProvisionAndInviteEnterpriseUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetInformationForProvisionedEnterpriseUser(context.Background(), enterprise, scimUserId).EnterpriseAdminProvisionAndInviteEnterpriseUserRequest(enterpriseAdminProvisionAndInviteEnterpriseUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetInformationForProvisionedEnterpriseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminSetInformationForProvisionedEnterpriseUser`: ScimEnterpriseUser
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminSetInformationForProvisionedEnterpriseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetInformationForProvisionedEnterpriseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminProvisionAndInviteEnterpriseUserRequest** | [**EnterpriseAdminProvisionAndInviteEnterpriseUserRequest**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequest.md) |  | 

### Return type

[**ScimEnterpriseUser**](ScimEnterpriseUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterprise

> EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterprise(ctx, enterprise, runnerGroupId).EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest(enterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest).Execute()

Set organization access for a self-hosted runner group in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    enterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest([]int32{int32(123)}) // EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterprise(context.Background(), enterprise, runnerGroupId).EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest(enterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest** | [**EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest**](EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterprise

> EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterprise(ctx, enterprise).EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest(enterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest).Execute()

Set selected organizations enabled for GitHub Actions in an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    enterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest([]int32{int32(123)}) // EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterprise(context.Background(), enterprise).EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest(enterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest** | [**EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest**](EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminSetSelfHostedRunnersInGroupForEnterprise

> EnterpriseAdminSetSelfHostedRunnersInGroupForEnterprise(ctx, enterprise, runnerGroupId).EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest).Execute()

Set self-hosted runners in a group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest := *openapiclient.NewEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest([]int32{int32(123)}) // EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminSetSelfHostedRunnersInGroupForEnterprise(context.Background(), enterprise, runnerGroupId).EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminSetSelfHostedRunnersInGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest** | [**EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest**](EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminUpdateAttributeForEnterpriseGroup

> ScimEnterpriseGroup EnterpriseAdminUpdateAttributeForEnterpriseGroup(ctx, enterprise, scimGroupId).EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest(enterpriseAdminUpdateAttributeForEnterpriseGroupRequest).Execute()

Update an attribute for a SCIM enterprise group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimGroupId := "scimGroupId_example" // string | Identifier generated by the GitHub SCIM endpoint.
    enterpriseAdminUpdateAttributeForEnterpriseGroupRequest := *openapiclient.NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequest([]string{"Schemas_example"}, []openapiclient.EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner{*openapiclient.NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner("Op_example")}) // EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminUpdateAttributeForEnterpriseGroup(context.Background(), enterprise, scimGroupId).EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest(enterpriseAdminUpdateAttributeForEnterpriseGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminUpdateAttributeForEnterpriseGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminUpdateAttributeForEnterpriseGroup`: ScimEnterpriseGroup
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminUpdateAttributeForEnterpriseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimGroupId** | **string** | Identifier generated by the GitHub SCIM endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminUpdateAttributeForEnterpriseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminUpdateAttributeForEnterpriseGroupRequest** | [**EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest**](EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest.md) |  | 

### Return type

[**ScimEnterpriseGroup**](ScimEnterpriseGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminUpdateAttributeForEnterpriseUser

> ScimEnterpriseUser EnterpriseAdminUpdateAttributeForEnterpriseUser(ctx, enterprise, scimUserId).EnterpriseAdminUpdateAttributeForEnterpriseUserRequest(enterpriseAdminUpdateAttributeForEnterpriseUserRequest).Execute()

Update an attribute for a SCIM enterprise user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.
    enterpriseAdminUpdateAttributeForEnterpriseUserRequest := *openapiclient.NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequest([]string{"Schemas_example"}, []map[string]interface{}{map[string]interface{}(123)}) // EnterpriseAdminUpdateAttributeForEnterpriseUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminUpdateAttributeForEnterpriseUser(context.Background(), enterprise, scimUserId).EnterpriseAdminUpdateAttributeForEnterpriseUserRequest(enterpriseAdminUpdateAttributeForEnterpriseUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminUpdateAttributeForEnterpriseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminUpdateAttributeForEnterpriseUser`: ScimEnterpriseUser
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminUpdateAttributeForEnterpriseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminUpdateAttributeForEnterpriseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminUpdateAttributeForEnterpriseUserRequest** | [**EnterpriseAdminUpdateAttributeForEnterpriseUserRequest**](EnterpriseAdminUpdateAttributeForEnterpriseUserRequest.md) |  | 

### Return type

[**ScimEnterpriseUser**](ScimEnterpriseUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise

> RunnerGroupsEnterprise EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise(ctx, enterprise, runnerGroupId).EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest(enterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest).Execute()

Update a self-hosted runner group for an enterprise



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    runnerGroupId := int32(56) // int32 | Unique identifier of the self-hosted runner group.
    enterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest := *openapiclient.NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest() // EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseAdminApi.EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise(context.Background(), enterprise, runnerGroupId).EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest(enterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseAdminApi.EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise`: RunnerGroupsEnterprise
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseAdminApi.EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 
**runnerGroupId** | **int32** | Unique identifier of the self-hosted runner group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest** | [**EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest**](EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest.md) |  | 

### Return type

[**RunnerGroupsEnterprise**](RunnerGroupsEnterprise.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

