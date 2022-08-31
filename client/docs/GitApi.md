# \GitApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GitCreateBlob**](GitApi.md#GitCreateBlob) | **Post** /repos/{owner}/{repo}/git/blobs | Create a blob
[**GitCreateCommit**](GitApi.md#GitCreateCommit) | **Post** /repos/{owner}/{repo}/git/commits | Create a commit
[**GitCreateRef**](GitApi.md#GitCreateRef) | **Post** /repos/{owner}/{repo}/git/refs | Create a reference
[**GitCreateTag**](GitApi.md#GitCreateTag) | **Post** /repos/{owner}/{repo}/git/tags | Create a tag object
[**GitCreateTree**](GitApi.md#GitCreateTree) | **Post** /repos/{owner}/{repo}/git/trees | Create a tree
[**GitDeleteRef**](GitApi.md#GitDeleteRef) | **Delete** /repos/{owner}/{repo}/git/refs/{ref} | Delete a reference
[**GitGetBlob**](GitApi.md#GitGetBlob) | **Get** /repos/{owner}/{repo}/git/blobs/{file_sha} | Get a blob
[**GitGetCommit**](GitApi.md#GitGetCommit) | **Get** /repos/{owner}/{repo}/git/commits/{commit_sha} | Get a commit
[**GitGetRef**](GitApi.md#GitGetRef) | **Get** /repos/{owner}/{repo}/git/ref/{ref} | Get a reference
[**GitGetTag**](GitApi.md#GitGetTag) | **Get** /repos/{owner}/{repo}/git/tags/{tag_sha} | Get a tag
[**GitGetTree**](GitApi.md#GitGetTree) | **Get** /repos/{owner}/{repo}/git/trees/{tree_sha} | Get a tree
[**GitListMatchingRefs**](GitApi.md#GitListMatchingRefs) | **Get** /repos/{owner}/{repo}/git/matching-refs/{ref} | List matching references
[**GitUpdateRef**](GitApi.md#GitUpdateRef) | **Patch** /repos/{owner}/{repo}/git/refs/{ref} | Update a reference



## GitCreateBlob

> ShortBlob GitCreateBlob(ctx, owner, repo).GitCreateBlobRequest(gitCreateBlobRequest).Execute()

Create a blob



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
    gitCreateBlobRequest := *openapiclient.NewGitCreateBlobRequest("Content_example") // GitCreateBlobRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitCreateBlob(context.Background(), owner, repo).GitCreateBlobRequest(gitCreateBlobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitCreateBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitCreateBlob`: ShortBlob
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitCreateBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitCreateBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitCreateBlobRequest** | [**GitCreateBlobRequest**](GitCreateBlobRequest.md) |  | 

### Return type

[**ShortBlob**](ShortBlob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitCreateCommit

> GitCommit GitCreateCommit(ctx, owner, repo).GitCreateCommitRequest(gitCreateCommitRequest).Execute()

Create a commit



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
    gitCreateCommitRequest := *openapiclient.NewGitCreateCommitRequest("Message_example", "Tree_example") // GitCreateCommitRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitCreateCommit(context.Background(), owner, repo).GitCreateCommitRequest(gitCreateCommitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitCreateCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitCreateCommit`: GitCommit
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitCreateCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitCreateCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitCreateCommitRequest** | [**GitCreateCommitRequest**](GitCreateCommitRequest.md) |  | 

### Return type

[**GitCommit**](GitCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitCreateRef

> GitRef GitCreateRef(ctx, owner, repo).GitCreateRefRequest(gitCreateRefRequest).Execute()

Create a reference



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
    gitCreateRefRequest := *openapiclient.NewGitCreateRefRequest("Ref_example", "Sha_example") // GitCreateRefRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitCreateRef(context.Background(), owner, repo).GitCreateRefRequest(gitCreateRefRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitCreateRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitCreateRef`: GitRef
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitCreateRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitCreateRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitCreateRefRequest** | [**GitCreateRefRequest**](GitCreateRefRequest.md) |  | 

### Return type

[**GitRef**](GitRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitCreateTag

> GitTag GitCreateTag(ctx, owner, repo).GitCreateTagRequest(gitCreateTagRequest).Execute()

Create a tag object



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
    gitCreateTagRequest := *openapiclient.NewGitCreateTagRequest("Tag_example", "Message_example", "Object_example", "Type_example") // GitCreateTagRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitCreateTag(context.Background(), owner, repo).GitCreateTagRequest(gitCreateTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitCreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitCreateTag`: GitTag
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitCreateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitCreateTagRequest** | [**GitCreateTagRequest**](GitCreateTagRequest.md) |  | 

### Return type

[**GitTag**](GitTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitCreateTree

> GitTree GitCreateTree(ctx, owner, repo).GitCreateTreeRequest(gitCreateTreeRequest).Execute()

Create a tree



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
    gitCreateTreeRequest := *openapiclient.NewGitCreateTreeRequest([]openapiclient.GitCreateTreeRequestTreeInner{*openapiclient.NewGitCreateTreeRequestTreeInner()}) // GitCreateTreeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitCreateTree(context.Background(), owner, repo).GitCreateTreeRequest(gitCreateTreeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitCreateTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitCreateTree`: GitTree
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitCreateTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitCreateTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitCreateTreeRequest** | [**GitCreateTreeRequest**](GitCreateTreeRequest.md) |  | 

### Return type

[**GitTree**](GitTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitDeleteRef

> GitDeleteRef(ctx, owner, repo, ref).Execute()

Delete a reference



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitDeleteRef(context.Background(), owner, repo, ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitDeleteRef``: %v\n", err)
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
**ref** | **string** | ref parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitDeleteRefRequest struct via the builder pattern


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


## GitGetBlob

> Blob GitGetBlob(ctx, owner, repo, fileSha).Execute()

Get a blob



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
    fileSha := "fileSha_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitGetBlob(context.Background(), owner, repo, fileSha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitGetBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitGetBlob`: Blob
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitGetBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**fileSha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitGetBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Blob**](Blob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitGetCommit

> GitCommit GitGetCommit(ctx, owner, repo, commitSha).Execute()

Get a commit



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
    commitSha := "commitSha_example" // string | The SHA of the commit.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitGetCommit(context.Background(), owner, repo, commitSha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitGetCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitGetCommit`: GitCommit
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitGetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commitSha** | **string** | The SHA of the commit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GitCommit**](GitCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitGetRef

> GitRef GitGetRef(ctx, owner, repo, ref).Execute()

Get a reference



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitGetRef(context.Background(), owner, repo, ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitGetRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitGetRef`: GitRef
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitGetRef`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGitGetRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GitRef**](GitRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitGetTag

> GitTag GitGetTag(ctx, owner, repo, tagSha).Execute()

Get a tag



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
    tagSha := "tagSha_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitGetTag(context.Background(), owner, repo, tagSha).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitGetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitGetTag`: GitTag
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitGetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**tagSha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GitTag**](GitTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitGetTree

> GitTree GitGetTree(ctx, owner, repo, treeSha).Recursive(recursive).Execute()

Get a tree



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
    treeSha := "treeSha_example" // string | 
    recursive := "recursive_example" // string | Setting this parameter to any value returns the objects or subtrees referenced by the tree specified in `:tree_sha`. For example, setting `recursive` to any of the following will enable returning objects or subtrees: `0`, `1`, `\"true\"`, and `\"false\"`. Omit this parameter to prevent recursively returning objects or subtrees. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitGetTree(context.Background(), owner, repo, treeSha).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitGetTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitGetTree`: GitTree
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitGetTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**treeSha** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitGetTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **recursive** | **string** | Setting this parameter to any value returns the objects or subtrees referenced by the tree specified in &#x60;:tree_sha&#x60;. For example, setting &#x60;recursive&#x60; to any of the following will enable returning objects or subtrees: &#x60;0&#x60;, &#x60;1&#x60;, &#x60;\&quot;true\&quot;&#x60;, and &#x60;\&quot;false\&quot;&#x60;. Omit this parameter to prevent recursively returning objects or subtrees. | 

### Return type

[**GitTree**](GitTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitListMatchingRefs

> []GitRef GitListMatchingRefs(ctx, owner, repo, ref).Execute()

List matching references



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitListMatchingRefs(context.Background(), owner, repo, ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitListMatchingRefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitListMatchingRefs`: []GitRef
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitListMatchingRefs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGitListMatchingRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]GitRef**](GitRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitUpdateRef

> GitRef GitUpdateRef(ctx, owner, repo, ref).GitUpdateRefRequest(gitUpdateRefRequest).Execute()

Update a reference



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
    ref := "refs/head/master" // string | The name of the fully qualified reference to update. For example, `refs/heads/master`. If the value doesn't start with `refs` and have at least two slashes, it will be rejected.
    gitUpdateRefRequest := *openapiclient.NewGitUpdateRefRequest("Sha_example") // GitUpdateRefRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitApi.GitUpdateRef(context.Background(), owner, repo, ref).GitUpdateRefRequest(gitUpdateRefRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitApi.GitUpdateRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GitUpdateRef`: GitRef
    fmt.Fprintf(os.Stdout, "Response from `GitApi.GitUpdateRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**ref** | **string** | The name of the fully qualified reference to update. For example, &#x60;refs/heads/master&#x60;. If the value doesn&#39;t start with &#x60;refs&#x60; and have at least two slashes, it will be rejected. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitUpdateRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gitUpdateRefRequest** | [**GitUpdateRefRequest**](GitUpdateRefRequest.md) |  | 

### Return type

[**GitRef**](GitRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

