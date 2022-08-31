# \InteractionsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InteractionsGetRestrictionsForAuthenticatedUser**](InteractionsApi.md#InteractionsGetRestrictionsForAuthenticatedUser) | **Get** /user/interaction-limits | Get interaction restrictions for your public repositories
[**InteractionsGetRestrictionsForOrg**](InteractionsApi.md#InteractionsGetRestrictionsForOrg) | **Get** /orgs/{org}/interaction-limits | Get interaction restrictions for an organization
[**InteractionsGetRestrictionsForRepo**](InteractionsApi.md#InteractionsGetRestrictionsForRepo) | **Get** /repos/{owner}/{repo}/interaction-limits | Get interaction restrictions for a repository
[**InteractionsRemoveRestrictionsForAuthenticatedUser**](InteractionsApi.md#InteractionsRemoveRestrictionsForAuthenticatedUser) | **Delete** /user/interaction-limits | Remove interaction restrictions from your public repositories
[**InteractionsRemoveRestrictionsForOrg**](InteractionsApi.md#InteractionsRemoveRestrictionsForOrg) | **Delete** /orgs/{org}/interaction-limits | Remove interaction restrictions for an organization
[**InteractionsRemoveRestrictionsForRepo**](InteractionsApi.md#InteractionsRemoveRestrictionsForRepo) | **Delete** /repos/{owner}/{repo}/interaction-limits | Remove interaction restrictions for a repository
[**InteractionsSetRestrictionsForAuthenticatedUser**](InteractionsApi.md#InteractionsSetRestrictionsForAuthenticatedUser) | **Put** /user/interaction-limits | Set interaction restrictions for your public repositories
[**InteractionsSetRestrictionsForOrg**](InteractionsApi.md#InteractionsSetRestrictionsForOrg) | **Put** /orgs/{org}/interaction-limits | Set interaction restrictions for an organization
[**InteractionsSetRestrictionsForRepo**](InteractionsApi.md#InteractionsSetRestrictionsForRepo) | **Put** /repos/{owner}/{repo}/interaction-limits | Set interaction restrictions for a repository



## InteractionsGetRestrictionsForAuthenticatedUser

> InteractionsGetRestrictionsForOrg200Response InteractionsGetRestrictionsForAuthenticatedUser(ctx).Execute()

Get interaction restrictions for your public repositories



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
    resp, r, err := apiClient.InteractionsApi.InteractionsGetRestrictionsForAuthenticatedUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsGetRestrictionsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InteractionsGetRestrictionsForAuthenticatedUser`: InteractionsGetRestrictionsForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `InteractionsApi.InteractionsGetRestrictionsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsGetRestrictionsForAuthenticatedUserRequest struct via the builder pattern


### Return type

[**InteractionsGetRestrictionsForOrg200Response**](InteractionsGetRestrictionsForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InteractionsGetRestrictionsForOrg

> InteractionsGetRestrictionsForOrg200Response InteractionsGetRestrictionsForOrg(ctx, org).Execute()

Get interaction restrictions for an organization



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
    resp, r, err := apiClient.InteractionsApi.InteractionsGetRestrictionsForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsGetRestrictionsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InteractionsGetRestrictionsForOrg`: InteractionsGetRestrictionsForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `InteractionsApi.InteractionsGetRestrictionsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsGetRestrictionsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InteractionsGetRestrictionsForOrg200Response**](InteractionsGetRestrictionsForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InteractionsGetRestrictionsForRepo

> InteractionsGetRestrictionsForOrg200Response InteractionsGetRestrictionsForRepo(ctx, owner, repo).Execute()

Get interaction restrictions for a repository



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionsApi.InteractionsGetRestrictionsForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsGetRestrictionsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InteractionsGetRestrictionsForRepo`: InteractionsGetRestrictionsForOrg200Response
    fmt.Fprintf(os.Stdout, "Response from `InteractionsApi.InteractionsGetRestrictionsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsGetRestrictionsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InteractionsGetRestrictionsForOrg200Response**](InteractionsGetRestrictionsForOrg200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InteractionsRemoveRestrictionsForAuthenticatedUser

> InteractionsRemoveRestrictionsForAuthenticatedUser(ctx).Execute()

Remove interaction restrictions from your public repositories



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
    resp, r, err := apiClient.InteractionsApi.InteractionsRemoveRestrictionsForAuthenticatedUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsRemoveRestrictionsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsRemoveRestrictionsForAuthenticatedUserRequest struct via the builder pattern


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


## InteractionsRemoveRestrictionsForOrg

> InteractionsRemoveRestrictionsForOrg(ctx, org).Execute()

Remove interaction restrictions for an organization



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
    resp, r, err := apiClient.InteractionsApi.InteractionsRemoveRestrictionsForOrg(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsRemoveRestrictionsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsRemoveRestrictionsForOrgRequest struct via the builder pattern


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


## InteractionsRemoveRestrictionsForRepo

> InteractionsRemoveRestrictionsForRepo(ctx, owner, repo).Execute()

Remove interaction restrictions for a repository



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionsApi.InteractionsRemoveRestrictionsForRepo(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsRemoveRestrictionsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsRemoveRestrictionsForRepoRequest struct via the builder pattern


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


## InteractionsSetRestrictionsForAuthenticatedUser

> InteractionLimitResponse InteractionsSetRestrictionsForAuthenticatedUser(ctx).InteractionLimit(interactionLimit).Execute()

Set interaction restrictions for your public repositories



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
    interactionLimit := *openapiclient.NewInteractionLimit(openapiclient.interaction-group("existing_users")) // InteractionLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionsApi.InteractionsSetRestrictionsForAuthenticatedUser(context.Background()).InteractionLimit(interactionLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsSetRestrictionsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InteractionsSetRestrictionsForAuthenticatedUser`: InteractionLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `InteractionsApi.InteractionsSetRestrictionsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsSetRestrictionsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactionLimit** | [**InteractionLimit**](InteractionLimit.md) |  | 

### Return type

[**InteractionLimitResponse**](InteractionLimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InteractionsSetRestrictionsForOrg

> InteractionLimitResponse InteractionsSetRestrictionsForOrg(ctx, org).InteractionLimit(interactionLimit).Execute()

Set interaction restrictions for an organization



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
    interactionLimit := *openapiclient.NewInteractionLimit(openapiclient.interaction-group("existing_users")) // InteractionLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionsApi.InteractionsSetRestrictionsForOrg(context.Background(), org).InteractionLimit(interactionLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsSetRestrictionsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InteractionsSetRestrictionsForOrg`: InteractionLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `InteractionsApi.InteractionsSetRestrictionsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsSetRestrictionsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interactionLimit** | [**InteractionLimit**](InteractionLimit.md) |  | 

### Return type

[**InteractionLimitResponse**](InteractionLimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InteractionsSetRestrictionsForRepo

> InteractionLimitResponse InteractionsSetRestrictionsForRepo(ctx, owner, repo).InteractionLimit(interactionLimit).Execute()

Set interaction restrictions for a repository



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
    interactionLimit := *openapiclient.NewInteractionLimit(openapiclient.interaction-group("existing_users")) // InteractionLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InteractionsApi.InteractionsSetRestrictionsForRepo(context.Background(), owner, repo).InteractionLimit(interactionLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InteractionsApi.InteractionsSetRestrictionsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InteractionsSetRestrictionsForRepo`: InteractionLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `InteractionsApi.InteractionsSetRestrictionsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInteractionsSetRestrictionsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **interactionLimit** | [**InteractionLimit**](InteractionLimit.md) |  | 

### Return type

[**InteractionLimitResponse**](InteractionLimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

