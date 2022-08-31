# \CodesOfConductApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodesOfConductGetAllCodesOfConduct**](CodesOfConductApi.md#CodesOfConductGetAllCodesOfConduct) | **Get** /codes_of_conduct | Get all codes of conduct
[**CodesOfConductGetConductCode**](CodesOfConductApi.md#CodesOfConductGetConductCode) | **Get** /codes_of_conduct/{key} | Get a code of conduct



## CodesOfConductGetAllCodesOfConduct

> []CodeOfConduct CodesOfConductGetAllCodesOfConduct(ctx).Execute()

Get all codes of conduct



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodesOfConductApi.CodesOfConductGetAllCodesOfConduct(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodesOfConductApi.CodesOfConductGetAllCodesOfConduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodesOfConductGetAllCodesOfConduct`: []CodeOfConduct
    fmt.Fprintf(os.Stdout, "Response from `CodesOfConductApi.CodesOfConductGetAllCodesOfConduct`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCodesOfConductGetAllCodesOfConductRequest struct via the builder pattern


### Return type

[**[]CodeOfConduct**](CodeOfConduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodesOfConductGetConductCode

> CodeOfConduct CodesOfConductGetConductCode(ctx, key).Execute()

Get a code of conduct



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodesOfConductApi.CodesOfConductGetConductCode(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodesOfConductApi.CodesOfConductGetConductCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodesOfConductGetConductCode`: CodeOfConduct
    fmt.Fprintf(os.Stdout, "Response from `CodesOfConductApi.CodesOfConductGetConductCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodesOfConductGetConductCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CodeOfConduct**](CodeOfConduct.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

