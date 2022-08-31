# \ServerStatisticsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseAdminGetServerStatistics**](ServerStatisticsApi.md#EnterpriseAdminGetServerStatistics) | **Get** /enterprise-installation/{enterprise_or_org}/server-statistics | Get GitHub Enterprise Server statistics



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
    resp, r, err := apiClient.ServerStatisticsApi.EnterpriseAdminGetServerStatistics(context.Background(), enterpriseOrOrg).DateStart(dateStart).DateEnd(dateEnd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerStatisticsApi.EnterpriseAdminGetServerStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseAdminGetServerStatistics`: []ServerStatisticsInner
    fmt.Fprintf(os.Stdout, "Response from `ServerStatisticsApi.EnterpriseAdminGetServerStatistics`: %v\n", resp)
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

