# \ScimApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScimDeleteUserFromOrg**](ScimApi.md#ScimDeleteUserFromOrg) | **Delete** /scim/v2/organizations/{org}/Users/{scim_user_id} | Delete a SCIM user from an organization
[**ScimGetProvisioningInformationForUser**](ScimApi.md#ScimGetProvisioningInformationForUser) | **Get** /scim/v2/organizations/{org}/Users/{scim_user_id} | Get SCIM provisioning information for a user
[**ScimListProvisionedIdentities**](ScimApi.md#ScimListProvisionedIdentities) | **Get** /scim/v2/organizations/{org}/Users | List SCIM provisioned identities
[**ScimProvisionAndInviteUser**](ScimApi.md#ScimProvisionAndInviteUser) | **Post** /scim/v2/organizations/{org}/Users | Provision and invite a SCIM user
[**ScimSetInformationForProvisionedUser**](ScimApi.md#ScimSetInformationForProvisionedUser) | **Put** /scim/v2/organizations/{org}/Users/{scim_user_id} | Update a provisioned organization membership
[**ScimUpdateAttributeForUser**](ScimApi.md#ScimUpdateAttributeForUser) | **Patch** /scim/v2/organizations/{org}/Users/{scim_user_id} | Update an attribute for a SCIM user



## ScimDeleteUserFromOrg

> ScimDeleteUserFromOrg(ctx, org, scimUserId).Execute()

Delete a SCIM user from an organization



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimApi.ScimDeleteUserFromOrg(context.Background(), org, scimUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimApi.ScimDeleteUserFromOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimDeleteUserFromOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetProvisioningInformationForUser

> ScimUser ScimGetProvisioningInformationForUser(ctx, org, scimUserId).Execute()

Get SCIM provisioning information for a user



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimApi.ScimGetProvisioningInformationForUser(context.Background(), org, scimUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimApi.ScimGetProvisioningInformationForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScimGetProvisioningInformationForUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `ScimApi.ScimGetProvisioningInformationForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetProvisioningInformationForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScimUser**](ScimUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimListProvisionedIdentities

> ScimUserList ScimListProvisionedIdentities(ctx, org).StartIndex(startIndex).Count(count).Filter(filter).Execute()

List SCIM provisioned identities



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    startIndex := int32(56) // int32 | Used for pagination: the index of the first result to return. (optional)
    count := int32(56) // int32 | Used for pagination: the number of results to return. (optional)
    filter := "filter_example" // string | Filters results using the equals query parameter operator (`eq`). You can filter results that are equal to `id`, `userName`, `emails`, and `external_id`. For example, to search for an identity with the `userName` Octocat, you would use this query:  `?filter=userName%20eq%20\\\"Octocat\\\"`.  To filter results for the identity with the email `octocat@github.com`, you would use this query:  `?filter=emails%20eq%20\\\"octocat@github.com\\\"`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimApi.ScimListProvisionedIdentities(context.Background(), org).StartIndex(startIndex).Count(count).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimApi.ScimListProvisionedIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScimListProvisionedIdentities`: ScimUserList
    fmt.Fprintf(os.Stdout, "Response from `ScimApi.ScimListProvisionedIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimListProvisionedIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Used for pagination: the index of the first result to return. | 
 **count** | **int32** | Used for pagination: the number of results to return. | 
 **filter** | **string** | Filters results using the equals query parameter operator (&#x60;eq&#x60;). You can filter results that are equal to &#x60;id&#x60;, &#x60;userName&#x60;, &#x60;emails&#x60;, and &#x60;external_id&#x60;. For example, to search for an identity with the &#x60;userName&#x60; Octocat, you would use this query:  &#x60;?filter&#x3D;userName%20eq%20\\\&quot;Octocat\\\&quot;&#x60;.  To filter results for the identity with the email &#x60;octocat@github.com&#x60;, you would use this query:  &#x60;?filter&#x3D;emails%20eq%20\\\&quot;octocat@github.com\\\&quot;&#x60;. | 

### Return type

[**ScimUserList**](ScimUserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimProvisionAndInviteUser

> ScimUser ScimProvisionAndInviteUser(ctx, org).ScimProvisionAndInviteUserRequest(scimProvisionAndInviteUserRequest).Execute()

Provision and invite a SCIM user



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    scimProvisionAndInviteUserRequest := *openapiclient.NewScimProvisionAndInviteUserRequest("someone@example.com", *openapiclient.NewScimProvisionAndInviteUserRequestName("GivenName_example", "FamilyName_example"), []openapiclient.ScimProvisionAndInviteUserRequestEmailsInner{*openapiclient.NewScimProvisionAndInviteUserRequestEmailsInner("Value_example")}) // ScimProvisionAndInviteUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimApi.ScimProvisionAndInviteUser(context.Background(), org).ScimProvisionAndInviteUserRequest(scimProvisionAndInviteUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimApi.ScimProvisionAndInviteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScimProvisionAndInviteUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `ScimApi.ScimProvisionAndInviteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimProvisionAndInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimProvisionAndInviteUserRequest** | [**ScimProvisionAndInviteUserRequest**](ScimProvisionAndInviteUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/scim+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimSetInformationForProvisionedUser

> ScimUser ScimSetInformationForProvisionedUser(ctx, org, scimUserId).ScimSetInformationForProvisionedUserRequest(scimSetInformationForProvisionedUserRequest).Execute()

Update a provisioned organization membership



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.
    scimSetInformationForProvisionedUserRequest := *openapiclient.NewScimSetInformationForProvisionedUserRequest("someone@example.com", *openapiclient.NewScimProvisionAndInviteUserRequestName("GivenName_example", "FamilyName_example"), []openapiclient.ScimSetInformationForProvisionedUserRequestEmailsInner{*openapiclient.NewScimSetInformationForProvisionedUserRequestEmailsInner("Value_example")}) // ScimSetInformationForProvisionedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimApi.ScimSetInformationForProvisionedUser(context.Background(), org, scimUserId).ScimSetInformationForProvisionedUserRequest(scimSetInformationForProvisionedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimApi.ScimSetInformationForProvisionedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScimSetInformationForProvisionedUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `ScimApi.ScimSetInformationForProvisionedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimSetInformationForProvisionedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scimSetInformationForProvisionedUserRequest** | [**ScimSetInformationForProvisionedUserRequest**](ScimSetInformationForProvisionedUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/scim+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimUpdateAttributeForUser

> ScimUser ScimUpdateAttributeForUser(ctx, org, scimUserId).ScimUpdateAttributeForUserRequest(scimUpdateAttributeForUserRequest).Execute()

Update an attribute for a SCIM user



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    scimUserId := "scimUserId_example" // string | The unique identifier of the SCIM user.
    scimUpdateAttributeForUserRequest := *openapiclient.NewScimUpdateAttributeForUserRequest([]openapiclient.ScimUpdateAttributeForUserRequestOperationsInner{*openapiclient.NewScimUpdateAttributeForUserRequestOperationsInner("Op_example")}) // ScimUpdateAttributeForUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimApi.ScimUpdateAttributeForUser(context.Background(), org, scimUserId).ScimUpdateAttributeForUserRequest(scimUpdateAttributeForUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimApi.ScimUpdateAttributeForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScimUpdateAttributeForUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `ScimApi.ScimUpdateAttributeForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**scimUserId** | **string** | The unique identifier of the SCIM user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimUpdateAttributeForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scimUpdateAttributeForUserRequest** | [**ScimUpdateAttributeForUserRequest**](ScimUpdateAttributeForUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/scim+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

