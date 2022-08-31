# \ChecksApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChecksCreate**](ChecksApi.md#ChecksCreate) | **Post** /repos/{owner}/{repo}/check-runs | Create a check run
[**ChecksCreateSuite**](ChecksApi.md#ChecksCreateSuite) | **Post** /repos/{owner}/{repo}/check-suites | Create a check suite
[**ChecksGet**](ChecksApi.md#ChecksGet) | **Get** /repos/{owner}/{repo}/check-runs/{check_run_id} | Get a check run
[**ChecksGetSuite**](ChecksApi.md#ChecksGetSuite) | **Get** /repos/{owner}/{repo}/check-suites/{check_suite_id} | Get a check suite
[**ChecksListAnnotations**](ChecksApi.md#ChecksListAnnotations) | **Get** /repos/{owner}/{repo}/check-runs/{check_run_id}/annotations | List check run annotations
[**ChecksListForRef**](ChecksApi.md#ChecksListForRef) | **Get** /repos/{owner}/{repo}/commits/{ref}/check-runs | List check runs for a Git reference
[**ChecksListForSuite**](ChecksApi.md#ChecksListForSuite) | **Get** /repos/{owner}/{repo}/check-suites/{check_suite_id}/check-runs | List check runs in a check suite
[**ChecksListSuitesForRef**](ChecksApi.md#ChecksListSuitesForRef) | **Get** /repos/{owner}/{repo}/commits/{ref}/check-suites | List check suites for a Git reference
[**ChecksRerequestRun**](ChecksApi.md#ChecksRerequestRun) | **Post** /repos/{owner}/{repo}/check-runs/{check_run_id}/rerequest | Rerequest a check run
[**ChecksRerequestSuite**](ChecksApi.md#ChecksRerequestSuite) | **Post** /repos/{owner}/{repo}/check-suites/{check_suite_id}/rerequest | Rerequest a check suite
[**ChecksSetSuitesPreferences**](ChecksApi.md#ChecksSetSuitesPreferences) | **Patch** /repos/{owner}/{repo}/check-suites/preferences | Update repository preferences for check suites
[**ChecksUpdate**](ChecksApi.md#ChecksUpdate) | **Patch** /repos/{owner}/{repo}/check-runs/{check_run_id} | Update a check run



## ChecksCreate

> CheckRun ChecksCreate(ctx, owner, repo).ChecksCreateRequest(checksCreateRequest).Execute()

Create a check run



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checksCreateRequest := openapiclient.checks_create_request{ChecksCreateRequestOneOf: openapiclient.NewChecksCreateRequestOneOf("Status_example")} // ChecksCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksCreate(context.Background(), owner, repo).ChecksCreateRequest(checksCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksCreate`: CheckRun
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checksCreateRequest** | [**ChecksCreateRequest**](ChecksCreateRequest.md) |  | 

### Return type

[**CheckRun**](CheckRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksCreateSuite

> CheckSuite ChecksCreateSuite(ctx, owner, repo).ChecksCreateSuiteRequest(checksCreateSuiteRequest).Execute()

Create a check suite



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checksCreateSuiteRequest := *openapiclient.NewChecksCreateSuiteRequest("HeadSha_example") // ChecksCreateSuiteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksCreateSuite(context.Background(), owner, repo).ChecksCreateSuiteRequest(checksCreateSuiteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksCreateSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksCreateSuite`: CheckSuite
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksCreateSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksCreateSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checksCreateSuiteRequest** | [**ChecksCreateSuiteRequest**](ChecksCreateSuiteRequest.md) |  | 

### Return type

[**CheckSuite**](CheckSuite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksGet

> CheckRun ChecksGet(ctx, owner, repo, checkRunId).Execute()

Get a check run



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkRunId := int32(56) // int32 | The unique identifier of the check run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksGet(context.Background(), owner, repo, checkRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksGet`: CheckRun
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkRunId** | **int32** | The unique identifier of the check run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CheckRun**](CheckRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksGetSuite

> CheckSuite ChecksGetSuite(ctx, owner, repo, checkSuiteId).Execute()

Get a check suite



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkSuiteId := int32(56) // int32 | The unique identifier of the check suite.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksGetSuite(context.Background(), owner, repo, checkSuiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksGetSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksGetSuite`: CheckSuite
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksGetSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkSuiteId** | **int32** | The unique identifier of the check suite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksGetSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CheckSuite**](CheckSuite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksListAnnotations

> []CheckAnnotation ChecksListAnnotations(ctx, owner, repo, checkRunId).PerPage(perPage).Page(page).Execute()

List check run annotations



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkRunId := int32(56) // int32 | The unique identifier of the check run.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksListAnnotations(context.Background(), owner, repo, checkRunId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksListAnnotations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksListAnnotations`: []CheckAnnotation
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksListAnnotations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkRunId** | **int32** | The unique identifier of the check run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksListAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]CheckAnnotation**](CheckAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksListForRef

> ChecksListForSuite200Response ChecksListForRef(ctx, owner, repo, ref).CheckName(checkName).Status(status).Filter(filter).PerPage(perPage).Page(page).AppId(appId).Execute()

List check runs for a Git reference



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    ref := "ref_example" // string | ref parameter
    checkName := "checkName_example" // string | Returns check runs with the specified `name`. (optional)
    status := "status_example" // string | Returns check runs with the specified `status`. (optional)
    filter := "filter_example" // string | Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs. (optional) (default to "latest")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    appId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksListForRef(context.Background(), owner, repo, ref).CheckName(checkName).Status(status).Filter(filter).PerPage(perPage).Page(page).AppId(appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksListForRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksListForRef`: ChecksListForSuite200Response
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksListForRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**ref** | **string** | ref parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksListForRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **checkName** | **string** | Returns check runs with the specified &#x60;name&#x60;. | 
 **status** | **string** | Returns check runs with the specified &#x60;status&#x60;. | 
 **filter** | **string** | Filters check runs by their &#x60;completed_at&#x60; timestamp. &#x60;latest&#x60; returns the most recent check runs. | [default to &quot;latest&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **appId** | **int32** |  | 

### Return type

[**ChecksListForSuite200Response**](ChecksListForSuite200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksListForSuite

> ChecksListForSuite200Response ChecksListForSuite(ctx, owner, repo, checkSuiteId).CheckName(checkName).Status(status).Filter(filter).PerPage(perPage).Page(page).Execute()

List check runs in a check suite



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkSuiteId := int32(56) // int32 | The unique identifier of the check suite.
    checkName := "checkName_example" // string | Returns check runs with the specified `name`. (optional)
    status := "status_example" // string | Returns check runs with the specified `status`. (optional)
    filter := "filter_example" // string | Filters check runs by their `completed_at` timestamp. `latest` returns the most recent check runs. (optional) (default to "latest")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksListForSuite(context.Background(), owner, repo, checkSuiteId).CheckName(checkName).Status(status).Filter(filter).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksListForSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksListForSuite`: ChecksListForSuite200Response
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksListForSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkSuiteId** | **int32** | The unique identifier of the check suite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksListForSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **checkName** | **string** | Returns check runs with the specified &#x60;name&#x60;. | 
 **status** | **string** | Returns check runs with the specified &#x60;status&#x60;. | 
 **filter** | **string** | Filters check runs by their &#x60;completed_at&#x60; timestamp. &#x60;latest&#x60; returns the most recent check runs. | [default to &quot;latest&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ChecksListForSuite200Response**](ChecksListForSuite200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksListSuitesForRef

> ChecksListSuitesForRef200Response ChecksListSuitesForRef(ctx, owner, repo, ref).AppId(appId).CheckName(checkName).PerPage(perPage).Page(page).Execute()

List check suites for a Git reference



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    ref := "ref_example" // string | ref parameter
    appId := int32(1) // int32 | Filters check suites by GitHub App `id`. (optional)
    checkName := "checkName_example" // string | Returns check runs with the specified `name`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksListSuitesForRef(context.Background(), owner, repo, ref).AppId(appId).CheckName(checkName).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksListSuitesForRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksListSuitesForRef`: ChecksListSuitesForRef200Response
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksListSuitesForRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**ref** | **string** | ref parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksListSuitesForRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **appId** | **int32** | Filters check suites by GitHub App &#x60;id&#x60;. | 
 **checkName** | **string** | Returns check runs with the specified &#x60;name&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ChecksListSuitesForRef200Response**](ChecksListSuitesForRef200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksRerequestRun

> map[string]interface{} ChecksRerequestRun(ctx, owner, repo, checkRunId).Execute()

Rerequest a check run



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkRunId := int32(56) // int32 | The unique identifier of the check run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksRerequestRun(context.Background(), owner, repo, checkRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksRerequestRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksRerequestRun`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksRerequestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkRunId** | **int32** | The unique identifier of the check run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksRerequestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksRerequestSuite

> map[string]interface{} ChecksRerequestSuite(ctx, owner, repo, checkSuiteId).Execute()

Rerequest a check suite



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkSuiteId := int32(56) // int32 | The unique identifier of the check suite.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksRerequestSuite(context.Background(), owner, repo, checkSuiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksRerequestSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksRerequestSuite`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksRerequestSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkSuiteId** | **int32** | The unique identifier of the check suite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksRerequestSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksSetSuitesPreferences

> CheckSuitePreference ChecksSetSuitesPreferences(ctx, owner, repo).ChecksSetSuitesPreferencesRequest(checksSetSuitesPreferencesRequest).Execute()

Update repository preferences for check suites



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checksSetSuitesPreferencesRequest := *openapiclient.NewChecksSetSuitesPreferencesRequest() // ChecksSetSuitesPreferencesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksSetSuitesPreferences(context.Background(), owner, repo).ChecksSetSuitesPreferencesRequest(checksSetSuitesPreferencesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksSetSuitesPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksSetSuitesPreferences`: CheckSuitePreference
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksSetSuitesPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksSetSuitesPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checksSetSuitesPreferencesRequest** | [**ChecksSetSuitesPreferencesRequest**](ChecksSetSuitesPreferencesRequest.md) |  | 

### Return type

[**CheckSuitePreference**](CheckSuitePreference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksUpdate

> CheckRun ChecksUpdate(ctx, owner, repo, checkRunId).ChecksUpdateRequest(checksUpdateRequest).Execute()

Update a check run



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    checkRunId := int32(56) // int32 | The unique identifier of the check run.
    checksUpdateRequest := *openapiclient.NewChecksUpdateRequest() // ChecksUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChecksApi.ChecksUpdate(context.Background(), owner, repo, checkRunId).ChecksUpdateRequest(checksUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChecksApi.ChecksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChecksUpdate`: CheckRun
    fmt.Fprintf(os.Stdout, "Response from `ChecksApi.ChecksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**checkRunId** | **int32** | The unique identifier of the check run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **checksUpdateRequest** | [**ChecksUpdateRequest**](ChecksUpdateRequest.md) |  | 

### Return type

[**CheckRun**](CheckRun.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

