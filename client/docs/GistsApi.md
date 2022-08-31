# \GistsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GistsCheckIsStarred**](GistsApi.md#GistsCheckIsStarred) | **Get** /gists/{gist_id}/star | Check if a gist is starred
[**GistsCreate**](GistsApi.md#GistsCreate) | **Post** /gists | Create a gist
[**GistsCreateComment**](GistsApi.md#GistsCreateComment) | **Post** /gists/{gist_id}/comments | Create a gist comment
[**GistsDelete**](GistsApi.md#GistsDelete) | **Delete** /gists/{gist_id} | Delete a gist
[**GistsDeleteComment**](GistsApi.md#GistsDeleteComment) | **Delete** /gists/{gist_id}/comments/{comment_id} | Delete a gist comment
[**GistsFork**](GistsApi.md#GistsFork) | **Post** /gists/{gist_id}/forks | Fork a gist
[**GistsGet**](GistsApi.md#GistsGet) | **Get** /gists/{gist_id} | Get a gist
[**GistsGetComment**](GistsApi.md#GistsGetComment) | **Get** /gists/{gist_id}/comments/{comment_id} | Get a gist comment
[**GistsGetRevision**](GistsApi.md#GistsGetRevision) | **Get** /gists/{gist_id}/{sha} | Get a gist revision
[**GistsList**](GistsApi.md#GistsList) | **Get** /gists | List gists for the authenticated user
[**GistsListComments**](GistsApi.md#GistsListComments) | **Get** /gists/{gist_id}/comments | List gist comments
[**GistsListCommits**](GistsApi.md#GistsListCommits) | **Get** /gists/{gist_id}/commits | List gist commits
[**GistsListForUser**](GistsApi.md#GistsListForUser) | **Get** /users/{username}/gists | List gists for a user
[**GistsListForks**](GistsApi.md#GistsListForks) | **Get** /gists/{gist_id}/forks | List gist forks
[**GistsListPublic**](GistsApi.md#GistsListPublic) | **Get** /gists/public | List public gists
[**GistsListStarred**](GistsApi.md#GistsListStarred) | **Get** /gists/starred | List starred gists
[**GistsStar**](GistsApi.md#GistsStar) | **Put** /gists/{gist_id}/star | Star a gist
[**GistsUnstar**](GistsApi.md#GistsUnstar) | **Delete** /gists/{gist_id}/star | Unstar a gist
[**GistsUpdate**](GistsApi.md#GistsUpdate) | **Patch** /gists/{gist_id} | Update a gist
[**GistsUpdateComment**](GistsApi.md#GistsUpdateComment) | **Patch** /gists/{gist_id}/comments/{comment_id} | Update a gist comment



## GistsCheckIsStarred

> GistsCheckIsStarred(ctx, gistId).Execute()

Check if a gist is starred



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
    gistId := "gistId_example" // string | The unique identifier of the gist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsCheckIsStarred(context.Background(), gistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsCheckIsStarred``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsCheckIsStarredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsCreate

> GistSimple GistsCreate(ctx).GistsCreateRequest(gistsCreateRequest).Execute()

Create a gist



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
    gistsCreateRequest := *openapiclient.NewGistsCreateRequest(map[string]GistsCreateRequestFilesValue{"key": *openapiclient.NewGistsCreateRequestFilesValue("Content_example")}) // GistsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsCreate(context.Background()).GistsCreateRequest(gistsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsCreate`: GistSimple
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGistsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gistsCreateRequest** | [**GistsCreateRequest**](GistsCreateRequest.md) |  | 

### Return type

[**GistSimple**](GistSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsCreateComment

> GistComment GistsCreateComment(ctx, gistId).GistsCreateCommentRequest(gistsCreateCommentRequest).Execute()

Create a gist comment



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    gistsCreateCommentRequest := *openapiclient.NewGistsCreateCommentRequest("Body of the attachment") // GistsCreateCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsCreateComment(context.Background(), gistId).GistsCreateCommentRequest(gistsCreateCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsCreateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsCreateComment`: GistComment
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsCreateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gistsCreateCommentRequest** | [**GistsCreateCommentRequest**](GistsCreateCommentRequest.md) |  | 

### Return type

[**GistComment**](GistComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsDelete

> GistsDelete(ctx, gistId).Execute()

Delete a gist



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
    gistId := "gistId_example" // string | The unique identifier of the gist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsDelete(context.Background(), gistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsDeleteComment

> GistsDeleteComment(ctx, gistId, commentId).Execute()

Delete a gist comment



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    commentId := int32(56) // int32 | The unique identifier of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsDeleteComment(context.Background(), gistId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsDeleteComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsDeleteCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsFork

> BaseGist GistsFork(ctx, gistId).Execute()

Fork a gist



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
    gistId := "gistId_example" // string | The unique identifier of the gist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsFork(context.Background(), gistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsFork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsFork`: BaseGist
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsFork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseGist**](BaseGist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsGet

> GistSimple GistsGet(ctx, gistId).Execute()

Get a gist



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
    gistId := "gistId_example" // string | The unique identifier of the gist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsGet(context.Background(), gistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsGet`: GistSimple
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GistSimple**](GistSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsGetComment

> GistComment GistsGetComment(ctx, gistId, commentId).Execute()

Get a gist comment



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    commentId := int32(56) // int32 | The unique identifier of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsGetComment(context.Background(), gistId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsGetComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsGetComment`: GistComment
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsGetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GistComment**](GistComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsGetRevision

> GistSimple GistsGetRevision(ctx, gistId, sha).Execute()

Get a gist revision



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    sha := "sha_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsGetRevision(context.Background(), gistId, sha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsGetRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsGetRevision`: GistSimple
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsGetRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 
**sha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsGetRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GistSimple**](GistSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsList

> []BaseGist GistsList(ctx).Since(since).PerPage(perPage).Page(page).Execute()

List gists for the authenticated user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsList(context.Background()).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsList`: []BaseGist
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGistsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]BaseGist**](BaseGist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsListComments

> []GistComment GistsListComments(ctx, gistId).PerPage(perPage).Page(page).Execute()

List gist comments



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsListComments(context.Background(), gistId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsListComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsListComments`: []GistComment
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsListComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsListCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]GistComment**](GistComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsListCommits

> []GistCommit GistsListCommits(ctx, gistId).PerPage(perPage).Page(page).Execute()

List gist commits



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsListCommits(context.Background(), gistId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsListCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsListCommits`: []GistCommit
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsListCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsListCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]GistCommit**](GistCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsListForUser

> []BaseGist GistsListForUser(ctx, username).Since(since).PerPage(perPage).Page(page).Execute()

List gists for a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | The handle for the GitHub user account.
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsListForUser(context.Background(), username).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsListForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsListForUser`: []BaseGist
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsListForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsListForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]BaseGist**](BaseGist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsListForks

> []GistSimple GistsListForks(ctx, gistId).PerPage(perPage).Page(page).Execute()

List gist forks



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsListForks(context.Background(), gistId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsListForks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsListForks`: []GistSimple
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsListForks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsListForksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]GistSimple**](GistSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsListPublic

> []BaseGist GistsListPublic(ctx).Since(since).PerPage(perPage).Page(page).Execute()

List public gists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsListPublic(context.Background()).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsListPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsListPublic`: []BaseGist
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsListPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGistsListPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]BaseGist**](BaseGist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsListStarred

> []BaseGist GistsListStarred(ctx).Since(since).PerPage(perPage).Page(page).Execute()

List starred gists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsListStarred(context.Background()).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsListStarred``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsListStarred`: []BaseGist
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsListStarred`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGistsListStarredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]BaseGist**](BaseGist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsStar

> GistsStar(ctx, gistId).Execute()

Star a gist



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
    gistId := "gistId_example" // string | The unique identifier of the gist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsStar(context.Background(), gistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsStar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsUnstar

> GistsUnstar(ctx, gistId).Execute()

Unstar a gist



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
    gistId := "gistId_example" // string | The unique identifier of the gist.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsUnstar(context.Background(), gistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsUnstar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsUnstarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsUpdate

> GistSimple GistsUpdate(ctx, gistId).GistsUpdateRequest(gistsUpdateRequest).Execute()

Update a gist



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    gistsUpdateRequest := *openapiclient.NewGistsUpdateRequest() // GistsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsUpdate(context.Background(), gistId).GistsUpdateRequest(gistsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsUpdate`: GistSimple
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gistsUpdateRequest** | [**GistsUpdateRequest**](GistsUpdateRequest.md) |  | 

### Return type

[**GistSimple**](GistSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GistsUpdateComment

> GistComment GistsUpdateComment(ctx, gistId, commentId).GistsCreateCommentRequest(gistsCreateCommentRequest).Execute()

Update a gist comment



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
    gistId := "gistId_example" // string | The unique identifier of the gist.
    commentId := int32(56) // int32 | The unique identifier of the comment.
    gistsCreateCommentRequest := *openapiclient.NewGistsCreateCommentRequest("Body of the attachment") // GistsCreateCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GistsApi.GistsUpdateComment(context.Background(), gistId, commentId).GistsCreateCommentRequest(gistsCreateCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GistsApi.GistsUpdateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GistsUpdateComment`: GistComment
    fmt.Fprintf(os.Stdout, "Response from `GistsApi.GistsUpdateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gistId** | **string** | The unique identifier of the gist. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGistsUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gistsCreateCommentRequest** | [**GistsCreateCommentRequest**](GistsCreateCommentRequest.md) |  | 

### Return type

[**GistComment**](GistComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

