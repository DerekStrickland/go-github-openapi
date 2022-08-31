# \OidcApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OidcGetOidcCustomSubTemplateForOrg**](OidcApi.md#OidcGetOidcCustomSubTemplateForOrg) | **Get** /orgs/{org}/actions/oidc/customization/sub | Get the customization template for an OIDC subject claim for an organization
[**OidcUpdateOidcCustomSubTemplateForOrg**](OidcApi.md#OidcUpdateOidcCustomSubTemplateForOrg) | **Put** /orgs/{org}/actions/oidc/customization/sub | Set the customization template for an OIDC subject claim for an organization



## OidcGetOidcCustomSubTemplateForOrg

> OidcCustomSub OidcGetOidcCustomSubTemplateForOrg(ctx, org).Execute()

Get the customization template for an OIDC subject claim for an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcApi.OidcGetOidcCustomSubTemplateForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.OidcGetOidcCustomSubTemplateForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OidcGetOidcCustomSubTemplateForOrg`: OidcCustomSub
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.OidcGetOidcCustomSubTemplateForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOidcGetOidcCustomSubTemplateForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OidcCustomSub**](OidcCustomSub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OidcUpdateOidcCustomSubTemplateForOrg

> map[string]interface{} OidcUpdateOidcCustomSubTemplateForOrg(ctx, org).OidcCustomSub(oidcCustomSub).Execute()

Set the customization template for an OIDC subject claim for an organization



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
    oidcCustomSub := *openapiclient.NewOidcCustomSub([]string{"IncludeClaimKeys_example"}) // OidcCustomSub | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcApi.OidcUpdateOidcCustomSubTemplateForOrg(context.Background(), org).OidcCustomSub(oidcCustomSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcApi.OidcUpdateOidcCustomSubTemplateForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OidcUpdateOidcCustomSubTemplateForOrg`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OidcApi.OidcUpdateOidcCustomSubTemplateForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOidcUpdateOidcCustomSubTemplateForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcCustomSub** | [**OidcCustomSub**](OidcCustomSub.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

