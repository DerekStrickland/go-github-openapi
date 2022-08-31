# \ReactionsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactionsCreateForCommitComment**](ReactionsApi.md#ReactionsCreateForCommitComment) | **Post** /repos/{owner}/{repo}/comments/{comment_id}/reactions | Create reaction for a commit comment
[**ReactionsCreateForIssue**](ReactionsApi.md#ReactionsCreateForIssue) | **Post** /repos/{owner}/{repo}/issues/{issue_number}/reactions | Create reaction for an issue
[**ReactionsCreateForIssueComment**](ReactionsApi.md#ReactionsCreateForIssueComment) | **Post** /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions | Create reaction for an issue comment
[**ReactionsCreateForPullRequestReviewComment**](ReactionsApi.md#ReactionsCreateForPullRequestReviewComment) | **Post** /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions | Create reaction for a pull request review comment
[**ReactionsCreateForRelease**](ReactionsApi.md#ReactionsCreateForRelease) | **Post** /repos/{owner}/{repo}/releases/{release_id}/reactions | Create reaction for a release
[**ReactionsCreateForTeamDiscussionCommentInOrg**](ReactionsApi.md#ReactionsCreateForTeamDiscussionCommentInOrg) | **Post** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions | Create reaction for a team discussion comment
[**ReactionsCreateForTeamDiscussionCommentLegacy**](ReactionsApi.md#ReactionsCreateForTeamDiscussionCommentLegacy) | **Post** /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions | Create reaction for a team discussion comment (Legacy)
[**ReactionsCreateForTeamDiscussionInOrg**](ReactionsApi.md#ReactionsCreateForTeamDiscussionInOrg) | **Post** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions | Create reaction for a team discussion
[**ReactionsCreateForTeamDiscussionLegacy**](ReactionsApi.md#ReactionsCreateForTeamDiscussionLegacy) | **Post** /teams/{team_id}/discussions/{discussion_number}/reactions | Create reaction for a team discussion (Legacy)
[**ReactionsDeleteForCommitComment**](ReactionsApi.md#ReactionsDeleteForCommitComment) | **Delete** /repos/{owner}/{repo}/comments/{comment_id}/reactions/{reaction_id} | Delete a commit comment reaction
[**ReactionsDeleteForIssue**](ReactionsApi.md#ReactionsDeleteForIssue) | **Delete** /repos/{owner}/{repo}/issues/{issue_number}/reactions/{reaction_id} | Delete an issue reaction
[**ReactionsDeleteForIssueComment**](ReactionsApi.md#ReactionsDeleteForIssueComment) | **Delete** /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions/{reaction_id} | Delete an issue comment reaction
[**ReactionsDeleteForPullRequestComment**](ReactionsApi.md#ReactionsDeleteForPullRequestComment) | **Delete** /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions/{reaction_id} | Delete a pull request comment reaction
[**ReactionsDeleteForRelease**](ReactionsApi.md#ReactionsDeleteForRelease) | **Delete** /repos/{owner}/{repo}/releases/{release_id}/reactions/{reaction_id} | Delete a release reaction
[**ReactionsDeleteForTeamDiscussion**](ReactionsApi.md#ReactionsDeleteForTeamDiscussion) | **Delete** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions/{reaction_id} | Delete team discussion reaction
[**ReactionsDeleteForTeamDiscussionComment**](ReactionsApi.md#ReactionsDeleteForTeamDiscussionComment) | **Delete** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions/{reaction_id} | Delete team discussion comment reaction
[**ReactionsListForCommitComment**](ReactionsApi.md#ReactionsListForCommitComment) | **Get** /repos/{owner}/{repo}/comments/{comment_id}/reactions | List reactions for a commit comment
[**ReactionsListForIssue**](ReactionsApi.md#ReactionsListForIssue) | **Get** /repos/{owner}/{repo}/issues/{issue_number}/reactions | List reactions for an issue
[**ReactionsListForIssueComment**](ReactionsApi.md#ReactionsListForIssueComment) | **Get** /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions | List reactions for an issue comment
[**ReactionsListForPullRequestReviewComment**](ReactionsApi.md#ReactionsListForPullRequestReviewComment) | **Get** /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions | List reactions for a pull request review comment
[**ReactionsListForRelease**](ReactionsApi.md#ReactionsListForRelease) | **Get** /repos/{owner}/{repo}/releases/{release_id}/reactions | List reactions for a release
[**ReactionsListForTeamDiscussionCommentInOrg**](ReactionsApi.md#ReactionsListForTeamDiscussionCommentInOrg) | **Get** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions | List reactions for a team discussion comment
[**ReactionsListForTeamDiscussionCommentLegacy**](ReactionsApi.md#ReactionsListForTeamDiscussionCommentLegacy) | **Get** /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions | List reactions for a team discussion comment (Legacy)
[**ReactionsListForTeamDiscussionInOrg**](ReactionsApi.md#ReactionsListForTeamDiscussionInOrg) | **Get** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions | List reactions for a team discussion
[**ReactionsListForTeamDiscussionLegacy**](ReactionsApi.md#ReactionsListForTeamDiscussionLegacy) | **Get** /teams/{team_id}/discussions/{discussion_number}/reactions | List reactions for a team discussion (Legacy)



## ReactionsCreateForCommitComment

> Reaction ReactionsCreateForCommitComment(ctx, owner, repo, commentId).ReactionsCreateForCommitCommentRequest(reactionsCreateForCommitCommentRequest).Execute()

Create reaction for a commit comment



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
    reactionsCreateForCommitCommentRequest := *openapiclient.NewReactionsCreateForCommitCommentRequest("Content_example") // ReactionsCreateForCommitCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForCommitComment(context.Background(), owner, repo, commentId).ReactionsCreateForCommitCommentRequest(reactionsCreateForCommitCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForCommitComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForCommitComment`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForCommitComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReactionsCreateForCommitCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForCommitCommentRequest** | [**ReactionsCreateForCommitCommentRequest**](ReactionsCreateForCommitCommentRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForIssue

> Reaction ReactionsCreateForIssue(ctx, owner, repo, issueNumber).ReactionsCreateForIssueRequest(reactionsCreateForIssueRequest).Execute()

Create reaction for an issue



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
    issueNumber := int32(56) // int32 | The number that identifies the issue.
    reactionsCreateForIssueRequest := *openapiclient.NewReactionsCreateForIssueRequest("Content_example") // ReactionsCreateForIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForIssue(context.Background(), owner, repo, issueNumber).ReactionsCreateForIssueRequest(reactionsCreateForIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForIssue`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**issueNumber** | **int32** | The number that identifies the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsCreateForIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForIssueRequest** | [**ReactionsCreateForIssueRequest**](ReactionsCreateForIssueRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForIssueComment

> Reaction ReactionsCreateForIssueComment(ctx, owner, repo, commentId).ReactionsCreateForIssueCommentRequest(reactionsCreateForIssueCommentRequest).Execute()

Create reaction for an issue comment



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
    reactionsCreateForIssueCommentRequest := *openapiclient.NewReactionsCreateForIssueCommentRequest("Content_example") // ReactionsCreateForIssueCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForIssueComment(context.Background(), owner, repo, commentId).ReactionsCreateForIssueCommentRequest(reactionsCreateForIssueCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForIssueComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForIssueComment`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForIssueComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReactionsCreateForIssueCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForIssueCommentRequest** | [**ReactionsCreateForIssueCommentRequest**](ReactionsCreateForIssueCommentRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForPullRequestReviewComment

> Reaction ReactionsCreateForPullRequestReviewComment(ctx, owner, repo, commentId).ReactionsCreateForPullRequestReviewCommentRequest(reactionsCreateForPullRequestReviewCommentRequest).Execute()

Create reaction for a pull request review comment



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
    reactionsCreateForPullRequestReviewCommentRequest := *openapiclient.NewReactionsCreateForPullRequestReviewCommentRequest("Content_example") // ReactionsCreateForPullRequestReviewCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForPullRequestReviewComment(context.Background(), owner, repo, commentId).ReactionsCreateForPullRequestReviewCommentRequest(reactionsCreateForPullRequestReviewCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForPullRequestReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForPullRequestReviewComment`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForPullRequestReviewComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReactionsCreateForPullRequestReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForPullRequestReviewCommentRequest** | [**ReactionsCreateForPullRequestReviewCommentRequest**](ReactionsCreateForPullRequestReviewCommentRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForRelease

> Reaction ReactionsCreateForRelease(ctx, owner, repo, releaseId).ReactionsCreateForReleaseRequest(reactionsCreateForReleaseRequest).Execute()

Create reaction for a release



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
    releaseId := int32(56) // int32 | The unique identifier of the release.
    reactionsCreateForReleaseRequest := *openapiclient.NewReactionsCreateForReleaseRequest("Content_example") // ReactionsCreateForReleaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForRelease(context.Background(), owner, repo, releaseId).ReactionsCreateForReleaseRequest(reactionsCreateForReleaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForRelease`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsCreateForReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForReleaseRequest** | [**ReactionsCreateForReleaseRequest**](ReactionsCreateForReleaseRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForTeamDiscussionCommentInOrg

> Reaction ReactionsCreateForTeamDiscussionCommentInOrg(ctx, org, teamSlug, discussionNumber, commentNumber).ReactionsCreateForTeamDiscussionCommentInOrgRequest(reactionsCreateForTeamDiscussionCommentInOrgRequest).Execute()

Create reaction for a team discussion comment



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    reactionsCreateForTeamDiscussionCommentInOrgRequest := *openapiclient.NewReactionsCreateForTeamDiscussionCommentInOrgRequest("Content_example") // ReactionsCreateForTeamDiscussionCommentInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForTeamDiscussionCommentInOrg(context.Background(), org, teamSlug, discussionNumber, commentNumber).ReactionsCreateForTeamDiscussionCommentInOrgRequest(reactionsCreateForTeamDiscussionCommentInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForTeamDiscussionCommentInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForTeamDiscussionCommentInOrg`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForTeamDiscussionCommentInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsCreateForTeamDiscussionCommentInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **reactionsCreateForTeamDiscussionCommentInOrgRequest** | [**ReactionsCreateForTeamDiscussionCommentInOrgRequest**](ReactionsCreateForTeamDiscussionCommentInOrgRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForTeamDiscussionCommentLegacy

> Reaction ReactionsCreateForTeamDiscussionCommentLegacy(ctx, teamId, discussionNumber, commentNumber).ReactionsCreateForTeamDiscussionCommentInOrgRequest(reactionsCreateForTeamDiscussionCommentInOrgRequest).Execute()

Create reaction for a team discussion comment (Legacy)



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
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    reactionsCreateForTeamDiscussionCommentInOrgRequest := *openapiclient.NewReactionsCreateForTeamDiscussionCommentInOrgRequest("Content_example") // ReactionsCreateForTeamDiscussionCommentInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForTeamDiscussionCommentLegacy(context.Background(), teamId, discussionNumber, commentNumber).ReactionsCreateForTeamDiscussionCommentInOrgRequest(reactionsCreateForTeamDiscussionCommentInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForTeamDiscussionCommentLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForTeamDiscussionCommentLegacy`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForTeamDiscussionCommentLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsCreateForTeamDiscussionCommentLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForTeamDiscussionCommentInOrgRequest** | [**ReactionsCreateForTeamDiscussionCommentInOrgRequest**](ReactionsCreateForTeamDiscussionCommentInOrgRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForTeamDiscussionInOrg

> Reaction ReactionsCreateForTeamDiscussionInOrg(ctx, org, teamSlug, discussionNumber).ReactionsCreateForTeamDiscussionInOrgRequest(reactionsCreateForTeamDiscussionInOrgRequest).Execute()

Create reaction for a team discussion



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    reactionsCreateForTeamDiscussionInOrgRequest := *openapiclient.NewReactionsCreateForTeamDiscussionInOrgRequest("Content_example") // ReactionsCreateForTeamDiscussionInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForTeamDiscussionInOrg(context.Background(), org, teamSlug, discussionNumber).ReactionsCreateForTeamDiscussionInOrgRequest(reactionsCreateForTeamDiscussionInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForTeamDiscussionInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForTeamDiscussionInOrg`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForTeamDiscussionInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsCreateForTeamDiscussionInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reactionsCreateForTeamDiscussionInOrgRequest** | [**ReactionsCreateForTeamDiscussionInOrgRequest**](ReactionsCreateForTeamDiscussionInOrgRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsCreateForTeamDiscussionLegacy

> Reaction ReactionsCreateForTeamDiscussionLegacy(ctx, teamId, discussionNumber).ReactionsCreateForTeamDiscussionInOrgRequest(reactionsCreateForTeamDiscussionInOrgRequest).Execute()

Create reaction for a team discussion (Legacy)



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
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    reactionsCreateForTeamDiscussionInOrgRequest := *openapiclient.NewReactionsCreateForTeamDiscussionInOrgRequest("Content_example") // ReactionsCreateForTeamDiscussionInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsCreateForTeamDiscussionLegacy(context.Background(), teamId, discussionNumber).ReactionsCreateForTeamDiscussionInOrgRequest(reactionsCreateForTeamDiscussionInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsCreateForTeamDiscussionLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsCreateForTeamDiscussionLegacy`: Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsCreateForTeamDiscussionLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsCreateForTeamDiscussionLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reactionsCreateForTeamDiscussionInOrgRequest** | [**ReactionsCreateForTeamDiscussionInOrgRequest**](ReactionsCreateForTeamDiscussionInOrgRequest.md) |  | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsDeleteForCommitComment

> ReactionsDeleteForCommitComment(ctx, owner, repo, commentId, reactionId).Execute()

Delete a commit comment reaction



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
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForCommitComment(context.Background(), owner, repo, commentId, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForCommitComment``: %v\n", err)
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
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForCommitCommentRequest struct via the builder pattern


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


## ReactionsDeleteForIssue

> ReactionsDeleteForIssue(ctx, owner, repo, issueNumber, reactionId).Execute()

Delete an issue reaction



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
    issueNumber := int32(56) // int32 | The number that identifies the issue.
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForIssue(context.Background(), owner, repo, issueNumber, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForIssue``: %v\n", err)
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
**issueNumber** | **int32** | The number that identifies the issue. | 
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForIssueRequest struct via the builder pattern


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


## ReactionsDeleteForIssueComment

> ReactionsDeleteForIssueComment(ctx, owner, repo, commentId, reactionId).Execute()

Delete an issue comment reaction



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
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForIssueComment(context.Background(), owner, repo, commentId, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForIssueComment``: %v\n", err)
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
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForIssueCommentRequest struct via the builder pattern


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


## ReactionsDeleteForPullRequestComment

> ReactionsDeleteForPullRequestComment(ctx, owner, repo, commentId, reactionId).Execute()

Delete a pull request comment reaction



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
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForPullRequestComment(context.Background(), owner, repo, commentId, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForPullRequestComment``: %v\n", err)
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
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForPullRequestCommentRequest struct via the builder pattern


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


## ReactionsDeleteForRelease

> ReactionsDeleteForRelease(ctx, owner, repo, releaseId, reactionId).Execute()

Delete a release reaction



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
    releaseId := int32(56) // int32 | The unique identifier of the release.
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForRelease(context.Background(), owner, repo, releaseId, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForRelease``: %v\n", err)
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
**releaseId** | **int32** | The unique identifier of the release. | 
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForReleaseRequest struct via the builder pattern


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


## ReactionsDeleteForTeamDiscussion

> ReactionsDeleteForTeamDiscussion(ctx, org, teamSlug, discussionNumber, reactionId).Execute()

Delete team discussion reaction



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForTeamDiscussion(context.Background(), org, teamSlug, discussionNumber, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForTeamDiscussion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForTeamDiscussionRequest struct via the builder pattern


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


## ReactionsDeleteForTeamDiscussionComment

> ReactionsDeleteForTeamDiscussionComment(ctx, org, teamSlug, discussionNumber, commentNumber, reactionId).Execute()

Delete team discussion comment reaction



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    reactionId := int32(56) // int32 | The unique identifier of the reaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsDeleteForTeamDiscussionComment(context.Background(), org, teamSlug, discussionNumber, commentNumber, reactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsDeleteForTeamDiscussionComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 
**reactionId** | **int32** | The unique identifier of the reaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsDeleteForTeamDiscussionCommentRequest struct via the builder pattern


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


## ReactionsListForCommitComment

> []Reaction ReactionsListForCommitComment(ctx, owner, repo, commentId).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a commit comment



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
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a commit comment. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForCommitComment(context.Background(), owner, repo, commentId).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForCommitComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForCommitComment`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForCommitComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReactionsListForCommitCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a commit comment. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForIssue

> []Reaction ReactionsListForIssue(ctx, owner, repo, issueNumber).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for an issue



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
    issueNumber := int32(56) // int32 | The number that identifies the issue.
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to an issue. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForIssue(context.Background(), owner, repo, issueNumber).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForIssue`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**issueNumber** | **int32** | The number that identifies the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsListForIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to an issue. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForIssueComment

> []Reaction ReactionsListForIssueComment(ctx, owner, repo, commentId).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for an issue comment



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
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to an issue comment. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForIssueComment(context.Background(), owner, repo, commentId).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForIssueComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForIssueComment`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForIssueComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReactionsListForIssueCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to an issue comment. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForPullRequestReviewComment

> []Reaction ReactionsListForPullRequestReviewComment(ctx, owner, repo, commentId).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a pull request review comment



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
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a pull request review comment. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForPullRequestReviewComment(context.Background(), owner, repo, commentId).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForPullRequestReviewComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForPullRequestReviewComment`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForPullRequestReviewComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReactionsListForPullRequestReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a pull request review comment. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForRelease

> []Reaction ReactionsListForRelease(ctx, owner, repo, releaseId).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a release



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
    releaseId := int32(56) // int32 | The unique identifier of the release.
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a release. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForRelease(context.Background(), owner, repo, releaseId).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForRelease`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**releaseId** | **int32** | The unique identifier of the release. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsListForReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a release. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForTeamDiscussionCommentInOrg

> []Reaction ReactionsListForTeamDiscussionCommentInOrg(ctx, org, teamSlug, discussionNumber, commentNumber).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a team discussion comment



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion comment. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForTeamDiscussionCommentInOrg(context.Background(), org, teamSlug, discussionNumber, commentNumber).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForTeamDiscussionCommentInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForTeamDiscussionCommentInOrg`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForTeamDiscussionCommentInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsListForTeamDiscussionCommentInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion comment. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForTeamDiscussionCommentLegacy

> []Reaction ReactionsListForTeamDiscussionCommentLegacy(ctx, teamId, discussionNumber, commentNumber).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a team discussion comment (Legacy)



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
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion comment. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForTeamDiscussionCommentLegacy(context.Background(), teamId, discussionNumber, commentNumber).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForTeamDiscussionCommentLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForTeamDiscussionCommentLegacy`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForTeamDiscussionCommentLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsListForTeamDiscussionCommentLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion comment. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForTeamDiscussionInOrg

> []Reaction ReactionsListForTeamDiscussionInOrg(ctx, org, teamSlug, discussionNumber).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a team discussion



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForTeamDiscussionInOrg(context.Background(), org, teamSlug, discussionNumber).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForTeamDiscussionInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForTeamDiscussionInOrg`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForTeamDiscussionInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsListForTeamDiscussionInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactionsListForTeamDiscussionLegacy

> []Reaction ReactionsListForTeamDiscussionLegacy(ctx, teamId, discussionNumber).Content(content).PerPage(perPage).Page(page).Execute()

List reactions for a team discussion (Legacy)



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
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    content := "content_example" // string | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReactionsApi.ReactionsListForTeamDiscussionLegacy(context.Background(), teamId, discussionNumber).Content(content).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReactionsApi.ReactionsListForTeamDiscussionLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReactionsListForTeamDiscussionLegacy`: []Reaction
    fmt.Fprintf(os.Stdout, "Response from `ReactionsApi.ReactionsListForTeamDiscussionLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsListForTeamDiscussionLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** | Returns a single [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types). Omit this parameter to list all reactions to a team discussion. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

