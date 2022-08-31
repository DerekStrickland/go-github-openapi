# \BillingApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingGetGithubActionsBillingGhe**](BillingApi.md#BillingGetGithubActionsBillingGhe) | **Get** /enterprises/{enterprise}/settings/billing/actions | Get GitHub Actions billing for an enterprise
[**BillingGetGithubActionsBillingOrg**](BillingApi.md#BillingGetGithubActionsBillingOrg) | **Get** /orgs/{org}/settings/billing/actions | Get GitHub Actions billing for an organization
[**BillingGetGithubActionsBillingUser**](BillingApi.md#BillingGetGithubActionsBillingUser) | **Get** /users/{username}/settings/billing/actions | Get GitHub Actions billing for a user
[**BillingGetGithubAdvancedSecurityBillingGhe**](BillingApi.md#BillingGetGithubAdvancedSecurityBillingGhe) | **Get** /enterprises/{enterprise}/settings/billing/advanced-security | Get GitHub Advanced Security active committers for an enterprise
[**BillingGetGithubAdvancedSecurityBillingOrg**](BillingApi.md#BillingGetGithubAdvancedSecurityBillingOrg) | **Get** /orgs/{org}/settings/billing/advanced-security | Get GitHub Advanced Security active committers for an organization
[**BillingGetGithubPackagesBillingGhe**](BillingApi.md#BillingGetGithubPackagesBillingGhe) | **Get** /enterprises/{enterprise}/settings/billing/packages | Get GitHub Packages billing for an enterprise
[**BillingGetGithubPackagesBillingOrg**](BillingApi.md#BillingGetGithubPackagesBillingOrg) | **Get** /orgs/{org}/settings/billing/packages | Get GitHub Packages billing for an organization
[**BillingGetGithubPackagesBillingUser**](BillingApi.md#BillingGetGithubPackagesBillingUser) | **Get** /users/{username}/settings/billing/packages | Get GitHub Packages billing for a user
[**BillingGetSharedStorageBillingGhe**](BillingApi.md#BillingGetSharedStorageBillingGhe) | **Get** /enterprises/{enterprise}/settings/billing/shared-storage | Get shared storage billing for an enterprise
[**BillingGetSharedStorageBillingOrg**](BillingApi.md#BillingGetSharedStorageBillingOrg) | **Get** /orgs/{org}/settings/billing/shared-storage | Get shared storage billing for an organization
[**BillingGetSharedStorageBillingUser**](BillingApi.md#BillingGetSharedStorageBillingUser) | **Get** /users/{username}/settings/billing/shared-storage | Get shared storage billing for a user



## BillingGetGithubActionsBillingGhe

> ActionsBillingUsage BillingGetGithubActionsBillingGhe(ctx, enterprise).Execute()

Get GitHub Actions billing for an enterprise



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
    resp, r, err := apiClient.BillingApi.BillingGetGithubActionsBillingGhe(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubActionsBillingGhe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubActionsBillingGhe`: ActionsBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubActionsBillingGhe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubActionsBillingGheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsBillingUsage**](ActionsBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubActionsBillingOrg

> ActionsBillingUsage BillingGetGithubActionsBillingOrg(ctx, org).Execute()

Get GitHub Actions billing for an organization



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
    resp, r, err := apiClient.BillingApi.BillingGetGithubActionsBillingOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubActionsBillingOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubActionsBillingOrg`: ActionsBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubActionsBillingOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubActionsBillingOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsBillingUsage**](ActionsBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubActionsBillingUser

> ActionsBillingUsage BillingGetGithubActionsBillingUser(ctx, username).Execute()

Get GitHub Actions billing for a user



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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.BillingGetGithubActionsBillingUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubActionsBillingUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubActionsBillingUser`: ActionsBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubActionsBillingUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubActionsBillingUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsBillingUsage**](ActionsBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubAdvancedSecurityBillingGhe

> AdvancedSecurityActiveCommitters BillingGetGithubAdvancedSecurityBillingGhe(ctx, enterprise).PerPage(perPage).Page(page).Execute()

Get GitHub Advanced Security active committers for an enterprise



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
    resp, r, err := apiClient.BillingApi.BillingGetGithubAdvancedSecurityBillingGhe(context.Background(), enterprise).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubAdvancedSecurityBillingGhe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubAdvancedSecurityBillingGhe`: AdvancedSecurityActiveCommitters
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubAdvancedSecurityBillingGhe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubAdvancedSecurityBillingGheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**AdvancedSecurityActiveCommitters**](AdvancedSecurityActiveCommitters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubAdvancedSecurityBillingOrg

> AdvancedSecurityActiveCommitters BillingGetGithubAdvancedSecurityBillingOrg(ctx, org).PerPage(perPage).Page(page).Execute()

Get GitHub Advanced Security active committers for an organization



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.BillingGetGithubAdvancedSecurityBillingOrg(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubAdvancedSecurityBillingOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubAdvancedSecurityBillingOrg`: AdvancedSecurityActiveCommitters
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubAdvancedSecurityBillingOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubAdvancedSecurityBillingOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**AdvancedSecurityActiveCommitters**](AdvancedSecurityActiveCommitters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubPackagesBillingGhe

> PackagesBillingUsage BillingGetGithubPackagesBillingGhe(ctx, enterprise).Execute()

Get GitHub Packages billing for an enterprise



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
    resp, r, err := apiClient.BillingApi.BillingGetGithubPackagesBillingGhe(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubPackagesBillingGhe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubPackagesBillingGhe`: PackagesBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubPackagesBillingGhe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubPackagesBillingGheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PackagesBillingUsage**](PackagesBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubPackagesBillingOrg

> PackagesBillingUsage BillingGetGithubPackagesBillingOrg(ctx, org).Execute()

Get GitHub Packages billing for an organization



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
    resp, r, err := apiClient.BillingApi.BillingGetGithubPackagesBillingOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubPackagesBillingOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubPackagesBillingOrg`: PackagesBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubPackagesBillingOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubPackagesBillingOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PackagesBillingUsage**](PackagesBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetGithubPackagesBillingUser

> PackagesBillingUsage BillingGetGithubPackagesBillingUser(ctx, username).Execute()

Get GitHub Packages billing for a user



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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.BillingGetGithubPackagesBillingUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetGithubPackagesBillingUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetGithubPackagesBillingUser`: PackagesBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetGithubPackagesBillingUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetGithubPackagesBillingUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PackagesBillingUsage**](PackagesBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetSharedStorageBillingGhe

> CombinedBillingUsage BillingGetSharedStorageBillingGhe(ctx, enterprise).Execute()

Get shared storage billing for an enterprise



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
    resp, r, err := apiClient.BillingApi.BillingGetSharedStorageBillingGhe(context.Background(), enterprise).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetSharedStorageBillingGhe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetSharedStorageBillingGhe`: CombinedBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetSharedStorageBillingGhe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetSharedStorageBillingGheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedBillingUsage**](CombinedBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetSharedStorageBillingOrg

> CombinedBillingUsage BillingGetSharedStorageBillingOrg(ctx, org).Execute()

Get shared storage billing for an organization



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
    resp, r, err := apiClient.BillingApi.BillingGetSharedStorageBillingOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetSharedStorageBillingOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetSharedStorageBillingOrg`: CombinedBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetSharedStorageBillingOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetSharedStorageBillingOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedBillingUsage**](CombinedBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingGetSharedStorageBillingUser

> CombinedBillingUsage BillingGetSharedStorageBillingUser(ctx, username).Execute()

Get shared storage billing for a user



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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.BillingGetSharedStorageBillingUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.BillingGetSharedStorageBillingUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BillingGetSharedStorageBillingUser`: CombinedBillingUsage
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.BillingGetSharedStorageBillingUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingGetSharedStorageBillingUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CombinedBillingUsage**](CombinedBillingUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

