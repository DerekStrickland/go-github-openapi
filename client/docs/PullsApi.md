# \PullsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PullsCheckIfMerged**](PullsApi.md#PullsCheckIfMerged) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/merge | Check if a pull request has been merged
[**PullsCreate**](PullsApi.md#PullsCreate) | **Post** /repos/{owner}/{repo}/pulls | Create a pull request
[**PullsCreateReplyForReviewComment**](PullsApi.md#PullsCreateReplyForReviewComment) | **Post** /repos/{owner}/{repo}/pulls/{pull_number}/comments/{comment_id}/replies | Create a reply for a review comment
[**PullsCreateReview**](PullsApi.md#PullsCreateReview) | **Post** /repos/{owner}/{repo}/pulls/{pull_number}/reviews | Create a review for a pull request
[**PullsCreateReviewComment**](PullsApi.md#PullsCreateReviewComment) | **Post** /repos/{owner}/{repo}/pulls/{pull_number}/comments | Create a review comment for a pull request
[**PullsDeletePendingReview**](PullsApi.md#PullsDeletePendingReview) | **Delete** /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id} | Delete a pending review for a pull request
[**PullsDeleteReviewComment**](PullsApi.md#PullsDeleteReviewComment) | **Delete** /repos/{owner}/{repo}/pulls/comments/{comment_id} | Delete a review comment for a pull request
[**PullsDismissReview**](PullsApi.md#PullsDismissReview) | **Put** /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/dismissals | Dismiss a review for a pull request
[**PullsGet**](PullsApi.md#PullsGet) | **Get** /repos/{owner}/{repo}/pulls/{pull_number} | Get a pull request
[**PullsGetReview**](PullsApi.md#PullsGetReview) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id} | Get a review for a pull request
[**PullsGetReviewComment**](PullsApi.md#PullsGetReviewComment) | **Get** /repos/{owner}/{repo}/pulls/comments/{comment_id} | Get a review comment for a pull request
[**PullsList**](PullsApi.md#PullsList) | **Get** /repos/{owner}/{repo}/pulls | List pull requests
[**PullsListCommentsForReview**](PullsApi.md#PullsListCommentsForReview) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/comments | List comments for a pull request review
[**PullsListCommits**](PullsApi.md#PullsListCommits) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/commits | List commits on a pull request
[**PullsListFiles**](PullsApi.md#PullsListFiles) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/files | List pull requests files
[**PullsListRequestedReviewers**](PullsApi.md#PullsListRequestedReviewers) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers | Get all requested reviewers for a pull request
[**PullsListReviewComments**](PullsApi.md#PullsListReviewComments) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/comments | List review comments on a pull request
[**PullsListReviewCommentsForRepo**](PullsApi.md#PullsListReviewCommentsForRepo) | **Get** /repos/{owner}/{repo}/pulls/comments | List review comments in a repository
[**PullsListReviews**](PullsApi.md#PullsListReviews) | **Get** /repos/{owner}/{repo}/pulls/{pull_number}/reviews | List reviews for a pull request
[**PullsMerge**](PullsApi.md#PullsMerge) | **Put** /repos/{owner}/{repo}/pulls/{pull_number}/merge | Merge a pull request
[**PullsRemoveRequestedReviewers**](PullsApi.md#PullsRemoveRequestedReviewers) | **Delete** /repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers | Remove requested reviewers from a pull request
[**PullsRequestReviewers**](PullsApi.md#PullsRequestReviewers) | **Post** /repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers | Request reviewers for a pull request
[**PullsSubmitReview**](PullsApi.md#PullsSubmitReview) | **Post** /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/events | Submit a review for a pull request
[**PullsUpdate**](PullsApi.md#PullsUpdate) | **Patch** /repos/{owner}/{repo}/pulls/{pull_number} | Update a pull request
[**PullsUpdateBranch**](PullsApi.md#PullsUpdateBranch) | **Put** /repos/{owner}/{repo}/pulls/{pull_number}/update-branch | Update a pull request branch
[**PullsUpdateReview**](PullsApi.md#PullsUpdateReview) | **Put** /repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id} | Update a review for a pull request
[**PullsUpdateReviewComment**](PullsApi.md#PullsUpdateReviewComment) | **Patch** /repos/{owner}/{repo}/pulls/comments/{comment_id} | Update a review comment for a pull request



## PullsCheckIfMerged

> PullsCheckIfMerged(ctx, owner, repo, pullNumber).Execute()

Check if a pull request has been merged



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsCheckIfMerged(context.Background(), owner, repo, pullNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsCheckIfMerged``: %v\n", err)
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
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsCheckIfMergedRequest struct via the builder pattern


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


## PullsCreate

> PullRequest PullsCreate(ctx, owner, repo).PullsCreateRequest(pullsCreateRequest).Execute()

Create a pull request



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
    pullsCreateRequest := *openapiclient.NewPullsCreateRequest("Head_example", "Base_example") // PullsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsCreate(context.Background(), owner, repo).PullsCreateRequest(pullsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsCreate`: PullRequest
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pullsCreateRequest** | [**PullsCreateRequest**](PullsCreateRequest.md) |  | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsCreateReplyForReviewComment

> PullRequestReviewComment PullsCreateReplyForReviewComment(ctx, owner, repo, pullNumber, commentId).PullsCreateReplyForReviewCommentRequest(pullsCreateReplyForReviewCommentRequest).Execute()

Create a reply for a review comment



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    commentId := int32(56) // int32 | The unique identifier of the comment.
    pullsCreateReplyForReviewCommentRequest := *openapiclient.NewPullsCreateReplyForReviewCommentRequest("Body_example") // PullsCreateReplyForReviewCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsCreateReplyForReviewComment(context.Background(), owner, repo, pullNumber, commentId).PullsCreateReplyForReviewCommentRequest(pullsCreateReplyForReviewCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsCreateReplyForReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsCreateReplyForReviewComment`: PullRequestReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsCreateReplyForReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsCreateReplyForReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pullsCreateReplyForReviewCommentRequest** | [**PullsCreateReplyForReviewCommentRequest**](PullsCreateReplyForReviewCommentRequest.md) |  | 

### Return type

[**PullRequestReviewComment**](PullRequestReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsCreateReview

> PullRequestReview PullsCreateReview(ctx, owner, repo, pullNumber).PullsCreateReviewRequest(pullsCreateReviewRequest).Execute()

Create a review for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsCreateReviewRequest := *openapiclient.NewPullsCreateReviewRequest() // PullsCreateReviewRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsCreateReview(context.Background(), owner, repo, pullNumber).PullsCreateReviewRequest(pullsCreateReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsCreateReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsCreateReview`: PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsCreateReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsCreateReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsCreateReviewRequest** | [**PullsCreateReviewRequest**](PullsCreateReviewRequest.md) |  | 

### Return type

[**PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsCreateReviewComment

> PullRequestReviewComment PullsCreateReviewComment(ctx, owner, repo, pullNumber).PullsCreateReviewCommentRequest(pullsCreateReviewCommentRequest).Execute()

Create a review comment for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsCreateReviewCommentRequest := *openapiclient.NewPullsCreateReviewCommentRequest("Body_example", "CommitId_example", "Path_example", int32(123)) // PullsCreateReviewCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsCreateReviewComment(context.Background(), owner, repo, pullNumber).PullsCreateReviewCommentRequest(pullsCreateReviewCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsCreateReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsCreateReviewComment`: PullRequestReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsCreateReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsCreateReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsCreateReviewCommentRequest** | [**PullsCreateReviewCommentRequest**](PullsCreateReviewCommentRequest.md) |  | 

### Return type

[**PullRequestReviewComment**](PullRequestReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsDeletePendingReview

> PullRequestReview PullsDeletePendingReview(ctx, owner, repo, pullNumber, reviewId).Execute()

Delete a pending review for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    reviewId := int32(56) // int32 | The unique identifier of the review.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsDeletePendingReview(context.Background(), owner, repo, pullNumber, reviewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsDeletePendingReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsDeletePendingReview`: PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsDeletePendingReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**reviewId** | **int32** | The unique identifier of the review. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsDeletePendingReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsDeleteReviewComment

> PullsDeleteReviewComment(ctx, owner, repo, commentId).Execute()

Delete a review comment for a pull request



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
    commentId := int32(56) // int32 | The unique identifier of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsDeleteReviewComment(context.Background(), owner, repo, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsDeleteReviewComment``: %v\n", err)
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
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsDeleteReviewCommentRequest struct via the builder pattern


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


## PullsDismissReview

> PullRequestReview PullsDismissReview(ctx, owner, repo, pullNumber, reviewId).PullsDismissReviewRequest(pullsDismissReviewRequest).Execute()

Dismiss a review for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    reviewId := int32(56) // int32 | The unique identifier of the review.
    pullsDismissReviewRequest := *openapiclient.NewPullsDismissReviewRequest("Message_example") // PullsDismissReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsDismissReview(context.Background(), owner, repo, pullNumber, reviewId).PullsDismissReviewRequest(pullsDismissReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsDismissReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsDismissReview`: PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsDismissReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**reviewId** | **int32** | The unique identifier of the review. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsDismissReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pullsDismissReviewRequest** | [**PullsDismissReviewRequest**](PullsDismissReviewRequest.md) |  | 

### Return type

[**PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsGet

> PullRequest PullsGet(ctx, owner, repo, pullNumber).Execute()

Get a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsGet(context.Background(), owner, repo, pullNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsGet`: PullRequest
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsGetReview

> PullRequestReview PullsGetReview(ctx, owner, repo, pullNumber, reviewId).Execute()

Get a review for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    reviewId := int32(56) // int32 | The unique identifier of the review.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsGetReview(context.Background(), owner, repo, pullNumber, reviewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsGetReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsGetReview`: PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsGetReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**reviewId** | **int32** | The unique identifier of the review. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsGetReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsGetReviewComment

> PullRequestReviewComment PullsGetReviewComment(ctx, owner, repo, commentId).Execute()

Get a review comment for a pull request



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
    commentId := int32(56) // int32 | The unique identifier of the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsGetReviewComment(context.Background(), owner, repo, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsGetReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsGetReviewComment`: PullRequestReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsGetReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsGetReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PullRequestReviewComment**](PullRequestReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsList

> []PullRequestSimple PullsList(ctx, owner, repo).State(state).Head(head).Base(base).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List pull requests



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
    state := "state_example" // string | Either `open`, `closed`, or `all` to filter by state. (optional) (default to "open")
    head := "head_example" // string | Filter pulls by head user or head organization and branch name in the format of `user:ref-name` or `organization:ref-name`. For example: `github:new-script-format` or `octocat:test-branch`. (optional)
    base := "base_example" // string | Filter pulls by base branch name. Example: `gh-pages`. (optional)
    sort := "sort_example" // string | What to sort results by. Can be either `created`, `updated`, `popularity` (comment count) or `long-running` (age, filtering by pulls updated in the last month). (optional) (default to "created")
    direction := "direction_example" // string | The direction of the sort. Can be either `asc` or `desc`. Default: `desc` when sort is `created` or sort is not specified, otherwise `asc`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsList(context.Background(), owner, repo).State(state).Head(head).Base(base).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsList`: []PullRequestSimple
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | Either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60; to filter by state. | [default to &quot;open&quot;]
 **head** | **string** | Filter pulls by head user or head organization and branch name in the format of &#x60;user:ref-name&#x60; or &#x60;organization:ref-name&#x60;. For example: &#x60;github:new-script-format&#x60; or &#x60;octocat:test-branch&#x60;. | 
 **base** | **string** | Filter pulls by base branch name. Example: &#x60;gh-pages&#x60;. | 
 **sort** | **string** | What to sort results by. Can be either &#x60;created&#x60;, &#x60;updated&#x60;, &#x60;popularity&#x60; (comment count) or &#x60;long-running&#x60; (age, filtering by pulls updated in the last month). | [default to &quot;created&quot;]
 **direction** | **string** | The direction of the sort. Can be either &#x60;asc&#x60; or &#x60;desc&#x60;. Default: &#x60;desc&#x60; when sort is &#x60;created&#x60; or sort is not specified, otherwise &#x60;asc&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]PullRequestSimple**](PullRequestSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListCommentsForReview

> []ReviewComment PullsListCommentsForReview(ctx, owner, repo, pullNumber, reviewId).PerPage(perPage).Page(page).Execute()

List comments for a pull request review



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    reviewId := int32(56) // int32 | The unique identifier of the review.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListCommentsForReview(context.Background(), owner, repo, pullNumber, reviewId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListCommentsForReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListCommentsForReview`: []ReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListCommentsForReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**reviewId** | **int32** | The unique identifier of the review. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListCommentsForReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]ReviewComment**](ReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListCommits

> []Commit PullsListCommits(ctx, owner, repo, pullNumber).PerPage(perPage).Page(page).Execute()

List commits on a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListCommits(context.Background(), owner, repo, pullNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListCommits`: []Commit
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Commit**](Commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListFiles

> []DiffEntry PullsListFiles(ctx, owner, repo, pullNumber).PerPage(perPage).Page(page).Execute()

List pull requests files



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListFiles(context.Background(), owner, repo, pullNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListFiles`: []DiffEntry
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]DiffEntry**](DiffEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListRequestedReviewers

> PullRequestReviewRequest PullsListRequestedReviewers(ctx, owner, repo, pullNumber).Execute()

Get all requested reviewers for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListRequestedReviewers(context.Background(), owner, repo, pullNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListRequestedReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListRequestedReviewers`: PullRequestReviewRequest
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListRequestedReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListRequestedReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PullRequestReviewRequest**](PullRequestReviewRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListReviewComments

> []PullRequestReviewComment PullsListReviewComments(ctx, owner, repo, pullNumber).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()

List review comments on a pull request



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    sort := "sort_example" // string | The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to. (optional) (default to "created")
    direction := "direction_example" // string | Can be either `asc` or `desc`. Ignored without `sort` parameter. (optional)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListReviewComments(context.Background(), owner, repo, pullNumber).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListReviewComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListReviewComments`: []PullRequestReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListReviewComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListReviewCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the repository was starred. &#x60;updated&#x60; means when the repository was last pushed to. | [default to &quot;created&quot;]
 **direction** | **string** | Can be either &#x60;asc&#x60; or &#x60;desc&#x60;. Ignored without &#x60;sort&#x60; parameter. | 
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]PullRequestReviewComment**](PullRequestReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListReviewCommentsForRepo

> []PullRequestReviewComment PullsListReviewCommentsForRepo(ctx, owner, repo).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()

List review comments in a repository



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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    sort := "sort_example" // string |  (optional)
    direction := "direction_example" // string | Can be either `asc` or `desc`. Ignored without `sort` parameter. (optional)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListReviewCommentsForRepo(context.Background(), owner, repo).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListReviewCommentsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListReviewCommentsForRepo`: []PullRequestReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListReviewCommentsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListReviewCommentsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** |  | 
 **direction** | **string** | Can be either &#x60;asc&#x60; or &#x60;desc&#x60;. Ignored without &#x60;sort&#x60; parameter. | 
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]PullRequestReviewComment**](PullRequestReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsListReviews

> []PullRequestReview PullsListReviews(ctx, owner, repo, pullNumber).PerPage(perPage).Page(page).Execute()

List reviews for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsListReviews(context.Background(), owner, repo, pullNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsListReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsListReviews`: []PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsListReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsListReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsMerge

> PullRequestMergeResult PullsMerge(ctx, owner, repo, pullNumber).PullsMergeRequest(pullsMergeRequest).Execute()

Merge a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsMergeRequest := *openapiclient.NewPullsMergeRequest() // PullsMergeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsMerge(context.Background(), owner, repo, pullNumber).PullsMergeRequest(pullsMergeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsMerge`: PullRequestMergeResult
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsMergeRequest** | [**PullsMergeRequest**](PullsMergeRequest.md) |  | 

### Return type

[**PullRequestMergeResult**](PullRequestMergeResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsRemoveRequestedReviewers

> PullRequestSimple PullsRemoveRequestedReviewers(ctx, owner, repo, pullNumber).PullsRemoveRequestedReviewersRequest(pullsRemoveRequestedReviewersRequest).Execute()

Remove requested reviewers from a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsRemoveRequestedReviewersRequest := *openapiclient.NewPullsRemoveRequestedReviewersRequest([]string{"Reviewers_example"}) // PullsRemoveRequestedReviewersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsRemoveRequestedReviewers(context.Background(), owner, repo, pullNumber).PullsRemoveRequestedReviewersRequest(pullsRemoveRequestedReviewersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsRemoveRequestedReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsRemoveRequestedReviewers`: PullRequestSimple
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsRemoveRequestedReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsRemoveRequestedReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsRemoveRequestedReviewersRequest** | [**PullsRemoveRequestedReviewersRequest**](PullsRemoveRequestedReviewersRequest.md) |  | 

### Return type

[**PullRequestSimple**](PullRequestSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsRequestReviewers

> PullRequestSimple PullsRequestReviewers(ctx, owner, repo, pullNumber).PullsRequestReviewersRequest(pullsRequestReviewersRequest).Execute()

Request reviewers for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsRequestReviewersRequest := *openapiclient.NewPullsRequestReviewersRequest() // PullsRequestReviewersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsRequestReviewers(context.Background(), owner, repo, pullNumber).PullsRequestReviewersRequest(pullsRequestReviewersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsRequestReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsRequestReviewers`: PullRequestSimple
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsRequestReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsRequestReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsRequestReviewersRequest** | [**PullsRequestReviewersRequest**](PullsRequestReviewersRequest.md) |  | 

### Return type

[**PullRequestSimple**](PullRequestSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsSubmitReview

> PullRequestReview PullsSubmitReview(ctx, owner, repo, pullNumber, reviewId).PullsSubmitReviewRequest(pullsSubmitReviewRequest).Execute()

Submit a review for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    reviewId := int32(56) // int32 | The unique identifier of the review.
    pullsSubmitReviewRequest := *openapiclient.NewPullsSubmitReviewRequest("Event_example") // PullsSubmitReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsSubmitReview(context.Background(), owner, repo, pullNumber, reviewId).PullsSubmitReviewRequest(pullsSubmitReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsSubmitReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsSubmitReview`: PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsSubmitReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**reviewId** | **int32** | The unique identifier of the review. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsSubmitReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pullsSubmitReviewRequest** | [**PullsSubmitReviewRequest**](PullsSubmitReviewRequest.md) |  | 

### Return type

[**PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsUpdate

> PullRequest PullsUpdate(ctx, owner, repo, pullNumber).PullsUpdateRequest(pullsUpdateRequest).Execute()

Update a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsUpdateRequest := *openapiclient.NewPullsUpdateRequest() // PullsUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsUpdate(context.Background(), owner, repo, pullNumber).PullsUpdateRequest(pullsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsUpdate`: PullRequest
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsUpdateRequest** | [**PullsUpdateRequest**](PullsUpdateRequest.md) |  | 

### Return type

[**PullRequest**](PullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsUpdateBranch

> ActivityMarkRepoNotificationsAsRead202Response PullsUpdateBranch(ctx, owner, repo, pullNumber).PullsUpdateBranchRequest(pullsUpdateBranchRequest).Execute()

Update a pull request branch



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    pullsUpdateBranchRequest := *openapiclient.NewPullsUpdateBranchRequest() // PullsUpdateBranchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsUpdateBranch(context.Background(), owner, repo, pullNumber).PullsUpdateBranchRequest(pullsUpdateBranchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsUpdateBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsUpdateBranch`: ActivityMarkRepoNotificationsAsRead202Response
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsUpdateBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsUpdateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsUpdateBranchRequest** | [**PullsUpdateBranchRequest**](PullsUpdateBranchRequest.md) |  | 

### Return type

[**ActivityMarkRepoNotificationsAsRead202Response**](ActivityMarkRepoNotificationsAsRead202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsUpdateReview

> PullRequestReview PullsUpdateReview(ctx, owner, repo, pullNumber, reviewId).PullsUpdateReviewRequest(pullsUpdateReviewRequest).Execute()

Update a review for a pull request



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
    pullNumber := int32(56) // int32 | The number that identifies the pull request.
    reviewId := int32(56) // int32 | The unique identifier of the review.
    pullsUpdateReviewRequest := *openapiclient.NewPullsUpdateReviewRequest("Body_example") // PullsUpdateReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsUpdateReview(context.Background(), owner, repo, pullNumber, reviewId).PullsUpdateReviewRequest(pullsUpdateReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsUpdateReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsUpdateReview`: PullRequestReview
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsUpdateReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**pullNumber** | **int32** | The number that identifies the pull request. | 
**reviewId** | **int32** | The unique identifier of the review. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsUpdateReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pullsUpdateReviewRequest** | [**PullsUpdateReviewRequest**](PullsUpdateReviewRequest.md) |  | 

### Return type

[**PullRequestReview**](PullRequestReview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullsUpdateReviewComment

> PullRequestReviewComment PullsUpdateReviewComment(ctx, owner, repo, commentId).PullsUpdateReviewCommentRequest(pullsUpdateReviewCommentRequest).Execute()

Update a review comment for a pull request



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
    commentId := int32(56) // int32 | The unique identifier of the comment.
    pullsUpdateReviewCommentRequest := *openapiclient.NewPullsUpdateReviewCommentRequest("Body_example") // PullsUpdateReviewCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullsApi.PullsUpdateReviewComment(context.Background(), owner, repo, commentId).PullsUpdateReviewCommentRequest(pullsUpdateReviewCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullsApi.PullsUpdateReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullsUpdateReviewComment`: PullRequestReviewComment
    fmt.Fprintf(os.Stdout, "Response from `PullsApi.PullsUpdateReviewComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**commentId** | **int32** | The unique identifier of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPullsUpdateReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pullsUpdateReviewCommentRequest** | [**PullsUpdateReviewCommentRequest**](PullsUpdateReviewCommentRequest.md) |  | 

### Return type

[**PullRequestReviewComment**](PullRequestReviewComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

