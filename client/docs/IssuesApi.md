# \IssuesApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssuesAddAssignees**](IssuesApi.md#IssuesAddAssignees) | **Post** /repos/{owner}/{repo}/issues/{issue_number}/assignees | Add assignees to an issue
[**IssuesAddLabels**](IssuesApi.md#IssuesAddLabels) | **Post** /repos/{owner}/{repo}/issues/{issue_number}/labels | Add labels to an issue
[**IssuesCheckUserCanBeAssigned**](IssuesApi.md#IssuesCheckUserCanBeAssigned) | **Get** /repos/{owner}/{repo}/assignees/{assignee} | Check if a user can be assigned
[**IssuesCreate**](IssuesApi.md#IssuesCreate) | **Post** /repos/{owner}/{repo}/issues | Create an issue
[**IssuesCreateComment**](IssuesApi.md#IssuesCreateComment) | **Post** /repos/{owner}/{repo}/issues/{issue_number}/comments | Create an issue comment
[**IssuesCreateLabel**](IssuesApi.md#IssuesCreateLabel) | **Post** /repos/{owner}/{repo}/labels | Create a label
[**IssuesCreateMilestone**](IssuesApi.md#IssuesCreateMilestone) | **Post** /repos/{owner}/{repo}/milestones | Create a milestone
[**IssuesDeleteComment**](IssuesApi.md#IssuesDeleteComment) | **Delete** /repos/{owner}/{repo}/issues/comments/{comment_id} | Delete an issue comment
[**IssuesDeleteLabel**](IssuesApi.md#IssuesDeleteLabel) | **Delete** /repos/{owner}/{repo}/labels/{name} | Delete a label
[**IssuesDeleteMilestone**](IssuesApi.md#IssuesDeleteMilestone) | **Delete** /repos/{owner}/{repo}/milestones/{milestone_number} | Delete a milestone
[**IssuesGet**](IssuesApi.md#IssuesGet) | **Get** /repos/{owner}/{repo}/issues/{issue_number} | Get an issue
[**IssuesGetComment**](IssuesApi.md#IssuesGetComment) | **Get** /repos/{owner}/{repo}/issues/comments/{comment_id} | Get an issue comment
[**IssuesGetEvent**](IssuesApi.md#IssuesGetEvent) | **Get** /repos/{owner}/{repo}/issues/events/{event_id} | Get an issue event
[**IssuesGetLabel**](IssuesApi.md#IssuesGetLabel) | **Get** /repos/{owner}/{repo}/labels/{name} | Get a label
[**IssuesGetMilestone**](IssuesApi.md#IssuesGetMilestone) | **Get** /repos/{owner}/{repo}/milestones/{milestone_number} | Get a milestone
[**IssuesList**](IssuesApi.md#IssuesList) | **Get** /issues | List issues assigned to the authenticated user
[**IssuesListAssignees**](IssuesApi.md#IssuesListAssignees) | **Get** /repos/{owner}/{repo}/assignees | List assignees
[**IssuesListComments**](IssuesApi.md#IssuesListComments) | **Get** /repos/{owner}/{repo}/issues/{issue_number}/comments | List issue comments
[**IssuesListCommentsForRepo**](IssuesApi.md#IssuesListCommentsForRepo) | **Get** /repos/{owner}/{repo}/issues/comments | List issue comments for a repository
[**IssuesListEvents**](IssuesApi.md#IssuesListEvents) | **Get** /repos/{owner}/{repo}/issues/{issue_number}/events | List issue events
[**IssuesListEventsForRepo**](IssuesApi.md#IssuesListEventsForRepo) | **Get** /repos/{owner}/{repo}/issues/events | List issue events for a repository
[**IssuesListEventsForTimeline**](IssuesApi.md#IssuesListEventsForTimeline) | **Get** /repos/{owner}/{repo}/issues/{issue_number}/timeline | List timeline events for an issue
[**IssuesListForAuthenticatedUser**](IssuesApi.md#IssuesListForAuthenticatedUser) | **Get** /user/issues | List user account issues assigned to the authenticated user
[**IssuesListForOrg**](IssuesApi.md#IssuesListForOrg) | **Get** /orgs/{org}/issues | List organization issues assigned to the authenticated user
[**IssuesListForRepo**](IssuesApi.md#IssuesListForRepo) | **Get** /repos/{owner}/{repo}/issues | List repository issues
[**IssuesListLabelsForMilestone**](IssuesApi.md#IssuesListLabelsForMilestone) | **Get** /repos/{owner}/{repo}/milestones/{milestone_number}/labels | List labels for issues in a milestone
[**IssuesListLabelsForRepo**](IssuesApi.md#IssuesListLabelsForRepo) | **Get** /repos/{owner}/{repo}/labels | List labels for a repository
[**IssuesListLabelsOnIssue**](IssuesApi.md#IssuesListLabelsOnIssue) | **Get** /repos/{owner}/{repo}/issues/{issue_number}/labels | List labels for an issue
[**IssuesListMilestones**](IssuesApi.md#IssuesListMilestones) | **Get** /repos/{owner}/{repo}/milestones | List milestones
[**IssuesLock**](IssuesApi.md#IssuesLock) | **Put** /repos/{owner}/{repo}/issues/{issue_number}/lock | Lock an issue
[**IssuesRemoveAllLabels**](IssuesApi.md#IssuesRemoveAllLabels) | **Delete** /repos/{owner}/{repo}/issues/{issue_number}/labels | Remove all labels from an issue
[**IssuesRemoveAssignees**](IssuesApi.md#IssuesRemoveAssignees) | **Delete** /repos/{owner}/{repo}/issues/{issue_number}/assignees | Remove assignees from an issue
[**IssuesRemoveLabel**](IssuesApi.md#IssuesRemoveLabel) | **Delete** /repos/{owner}/{repo}/issues/{issue_number}/labels/{name} | Remove a label from an issue
[**IssuesSetLabels**](IssuesApi.md#IssuesSetLabels) | **Put** /repos/{owner}/{repo}/issues/{issue_number}/labels | Set labels for an issue
[**IssuesUnlock**](IssuesApi.md#IssuesUnlock) | **Delete** /repos/{owner}/{repo}/issues/{issue_number}/lock | Unlock an issue
[**IssuesUpdate**](IssuesApi.md#IssuesUpdate) | **Patch** /repos/{owner}/{repo}/issues/{issue_number} | Update an issue
[**IssuesUpdateComment**](IssuesApi.md#IssuesUpdateComment) | **Patch** /repos/{owner}/{repo}/issues/comments/{comment_id} | Update an issue comment
[**IssuesUpdateLabel**](IssuesApi.md#IssuesUpdateLabel) | **Patch** /repos/{owner}/{repo}/labels/{name} | Update a label
[**IssuesUpdateMilestone**](IssuesApi.md#IssuesUpdateMilestone) | **Patch** /repos/{owner}/{repo}/milestones/{milestone_number} | Update a milestone



## IssuesAddAssignees

> Issue IssuesAddAssignees(ctx, owner, repo, issueNumber).IssuesAddAssigneesRequest(issuesAddAssigneesRequest).Execute()

Add assignees to an issue



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
    issuesAddAssigneesRequest := *openapiclient.NewIssuesAddAssigneesRequest() // IssuesAddAssigneesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesAddAssignees(context.Background(), owner, repo, issueNumber).IssuesAddAssigneesRequest(issuesAddAssigneesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesAddAssignees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesAddAssignees`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesAddAssignees`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesAddAssigneesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesAddAssigneesRequest** | [**IssuesAddAssigneesRequest**](IssuesAddAssigneesRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesAddLabels

> []Label IssuesAddLabels(ctx, owner, repo, issueNumber).IssuesAddLabelsRequest(issuesAddLabelsRequest).Execute()

Add labels to an issue



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
    issuesAddLabelsRequest := openapiclient.issues_add_labels_request{IssuesAddLabelsRequestOneOf: openapiclient.NewIssuesAddLabelsRequestOneOf()} // IssuesAddLabelsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesAddLabels(context.Background(), owner, repo, issueNumber).IssuesAddLabelsRequest(issuesAddLabelsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesAddLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesAddLabels`: []Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesAddLabels`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesAddLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesAddLabelsRequest** | [**IssuesAddLabelsRequest**](IssuesAddLabelsRequest.md) |  | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesCheckUserCanBeAssigned

> IssuesCheckUserCanBeAssigned(ctx, owner, repo, assignee).Execute()

Check if a user can be assigned



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
    assignee := "assignee_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesCheckUserCanBeAssigned(context.Background(), owner, repo, assignee).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesCheckUserCanBeAssigned``: %v\n", err)
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
**assignee** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesCheckUserCanBeAssignedRequest struct via the builder pattern


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


## IssuesCreate

> Issue IssuesCreate(ctx, owner, repo).IssuesCreateRequest(issuesCreateRequest).Execute()

Create an issue



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
    issuesCreateRequest := *openapiclient.NewIssuesCreateRequest(openapiclient.issues_create_request_title{Int32: new(int32)}) // IssuesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesCreate(context.Background(), owner, repo).IssuesCreateRequest(issuesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesCreate`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issuesCreateRequest** | [**IssuesCreateRequest**](IssuesCreateRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesCreateComment

> IssueComment IssuesCreateComment(ctx, owner, repo, issueNumber).IssuesUpdateCommentRequest(issuesUpdateCommentRequest).Execute()

Create an issue comment



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
    issuesUpdateCommentRequest := *openapiclient.NewIssuesUpdateCommentRequest("Body_example") // IssuesUpdateCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesCreateComment(context.Background(), owner, repo, issueNumber).IssuesUpdateCommentRequest(issuesUpdateCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesCreateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesCreateComment`: IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesCreateComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesUpdateCommentRequest** | [**IssuesUpdateCommentRequest**](IssuesUpdateCommentRequest.md) |  | 

### Return type

[**IssueComment**](IssueComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesCreateLabel

> Label IssuesCreateLabel(ctx, owner, repo).IssuesCreateLabelRequest(issuesCreateLabelRequest).Execute()

Create a label



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
    issuesCreateLabelRequest := *openapiclient.NewIssuesCreateLabelRequest("Name_example") // IssuesCreateLabelRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesCreateLabel(context.Background(), owner, repo).IssuesCreateLabelRequest(issuesCreateLabelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesCreateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesCreateLabel`: Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesCreateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issuesCreateLabelRequest** | [**IssuesCreateLabelRequest**](IssuesCreateLabelRequest.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesCreateMilestone

> Milestone IssuesCreateMilestone(ctx, owner, repo).IssuesCreateMilestoneRequest(issuesCreateMilestoneRequest).Execute()

Create a milestone



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
    issuesCreateMilestoneRequest := *openapiclient.NewIssuesCreateMilestoneRequest("Title_example") // IssuesCreateMilestoneRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesCreateMilestone(context.Background(), owner, repo).IssuesCreateMilestoneRequest(issuesCreateMilestoneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesCreateMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesCreateMilestone`: Milestone
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesCreateMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesCreateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issuesCreateMilestoneRequest** | [**IssuesCreateMilestoneRequest**](IssuesCreateMilestoneRequest.md) |  | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesDeleteComment

> IssuesDeleteComment(ctx, owner, repo, commentId).Execute()

Delete an issue comment



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
    resp, r, err := apiClient.IssuesApi.IssuesDeleteComment(context.Background(), owner, repo, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesDeleteComment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIssuesDeleteCommentRequest struct via the builder pattern


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


## IssuesDeleteLabel

> IssuesDeleteLabel(ctx, owner, repo, name).Execute()

Delete a label



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesDeleteLabel(context.Background(), owner, repo, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesDeleteLabel``: %v\n", err)
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
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesDeleteLabelRequest struct via the builder pattern


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


## IssuesDeleteMilestone

> IssuesDeleteMilestone(ctx, owner, repo, milestoneNumber).Execute()

Delete a milestone



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
    milestoneNumber := int32(56) // int32 | The number that identifies the milestone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesDeleteMilestone(context.Background(), owner, repo, milestoneNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesDeleteMilestone``: %v\n", err)
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
**milestoneNumber** | **int32** | The number that identifies the milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesDeleteMilestoneRequest struct via the builder pattern


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


## IssuesGet

> Issue IssuesGet(ctx, owner, repo, issueNumber).Execute()

Get an issue



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesGet(context.Background(), owner, repo, issueNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesGet`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesGetComment

> IssueComment IssuesGetComment(ctx, owner, repo, commentId).Execute()

Get an issue comment



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
    resp, r, err := apiClient.IssuesApi.IssuesGetComment(context.Background(), owner, repo, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesGetComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesGetComment`: IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesGetComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IssueComment**](IssueComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesGetEvent

> IssueEvent IssuesGetEvent(ctx, owner, repo, eventId).Execute()

Get an issue event



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
    eventId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesGetEvent(context.Background(), owner, repo, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesGetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesGetEvent`: IssueEvent
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesGetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**eventId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IssueEvent**](IssueEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesGetLabel

> Label IssuesGetLabel(ctx, owner, repo, name).Execute()

Get a label



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesGetLabel(context.Background(), owner, repo, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesGetLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesGetLabel`: Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesGetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesGetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesGetMilestone

> Milestone IssuesGetMilestone(ctx, owner, repo, milestoneNumber).Execute()

Get a milestone



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
    milestoneNumber := int32(56) // int32 | The number that identifies the milestone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesGetMilestone(context.Background(), owner, repo, milestoneNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesGetMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesGetMilestone`: Milestone
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesGetMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**milestoneNumber** | **int32** | The number that identifies the milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesGetMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesList

> []Issue IssuesList(ctx).Filter(filter).State(state).Labels(labels).Sort(sort).Direction(direction).Since(since).Collab(collab).Orgs(orgs).Owned(owned).Pulls(pulls).PerPage(perPage).Page(page).Execute()

List issues assigned to the authenticated user



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
    filter := "filter_example" // string | Indicates which sorts of issues to return. `assigned` means issues assigned to you. `created` means issues created by you. `mentioned` means issues mentioning you. `subscribed` means issues you're subscribed to updates for. `all` or `repos` means all issues you can see, regardless of participation or creation. (optional) (default to "assigned")
    state := "state_example" // string | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    labels := "labels_example" // string | A list of comma separated label names. Example: `bug,ui,@high` (optional)
    sort := "sort_example" // string | What to sort results by. Can be either `created`, `updated`, `comments`. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    collab := true // bool |  (optional)
    orgs := true // bool |  (optional)
    owned := true // bool |  (optional)
    pulls := true // bool |  (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesList(context.Background()).Filter(filter).State(state).Labels(labels).Sort(sort).Direction(direction).Since(since).Collab(collab).Orgs(orgs).Owned(owned).Pulls(pulls).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesList`: []Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Indicates which sorts of issues to return. &#x60;assigned&#x60; means issues assigned to you. &#x60;created&#x60; means issues created by you. &#x60;mentioned&#x60; means issues mentioning you. &#x60;subscribed&#x60; means issues you&#39;re subscribed to updates for. &#x60;all&#x60; or &#x60;repos&#x60; means all issues you can see, regardless of participation or creation. | [default to &quot;assigned&quot;]
 **state** | **string** | Indicates the state of the issues to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **labels** | **string** | A list of comma separated label names. Example: &#x60;bug,ui,@high&#x60; | 
 **sort** | **string** | What to sort results by. Can be either &#x60;created&#x60;, &#x60;updated&#x60;, &#x60;comments&#x60;. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **collab** | **bool** |  | 
 **orgs** | **bool** |  | 
 **owned** | **bool** |  | 
 **pulls** | **bool** |  | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListAssignees

> []SimpleUser IssuesListAssignees(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List assignees



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListAssignees(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListAssignees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListAssignees`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListAssignees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListAssigneesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]SimpleUser**](SimpleUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListComments

> []IssueComment IssuesListComments(ctx, owner, repo, issueNumber).Since(since).PerPage(perPage).Page(page).Execute()

List issue comments



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
    issueNumber := int32(56) // int32 | The number that identifies the issue.
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListComments(context.Background(), owner, repo, issueNumber).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListComments`: []IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListComments`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesListCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]IssueComment**](IssueComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListCommentsForRepo

> []IssueComment IssuesListCommentsForRepo(ctx, owner, repo).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()

List issue comments for a repository



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
    sort := "sort_example" // string | The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to. (optional) (default to "created")
    direction := "direction_example" // string | Either `asc` or `desc`. Ignored without the `sort` parameter. (optional)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListCommentsForRepo(context.Background(), owner, repo).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListCommentsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListCommentsForRepo`: []IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListCommentsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListCommentsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the repository was starred. &#x60;updated&#x60; means when the repository was last pushed to. | [default to &quot;created&quot;]
 **direction** | **string** | Either &#x60;asc&#x60; or &#x60;desc&#x60;. Ignored without the &#x60;sort&#x60; parameter. | 
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]IssueComment**](IssueComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListEvents

> []IssueEventForIssue IssuesListEvents(ctx, owner, repo, issueNumber).PerPage(perPage).Page(page).Execute()

List issue events



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListEvents(context.Background(), owner, repo, issueNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListEvents`: []IssueEventForIssue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListEvents`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]IssueEventForIssue**](IssueEventForIssue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListEventsForRepo

> []IssueEvent IssuesListEventsForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List issue events for a repository



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListEventsForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListEventsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListEventsForRepo`: []IssueEvent
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListEventsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListEventsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]IssueEvent**](IssueEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListEventsForTimeline

> []TimelineIssueEvents IssuesListEventsForTimeline(ctx, owner, repo, issueNumber).PerPage(perPage).Page(page).Execute()

List timeline events for an issue



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListEventsForTimeline(context.Background(), owner, repo, issueNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListEventsForTimeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListEventsForTimeline`: []TimelineIssueEvents
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListEventsForTimeline`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesListEventsForTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TimelineIssueEvents**](TimelineIssueEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListForAuthenticatedUser

> []Issue IssuesListForAuthenticatedUser(ctx).Filter(filter).State(state).Labels(labels).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()

List user account issues assigned to the authenticated user



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
    filter := "filter_example" // string | Indicates which sorts of issues to return. `assigned` means issues assigned to you. `created` means issues created by you. `mentioned` means issues mentioning you. `subscribed` means issues you're subscribed to updates for. `all` or `repos` means all issues you can see, regardless of participation or creation. (optional) (default to "assigned")
    state := "state_example" // string | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    labels := "labels_example" // string | A list of comma separated label names. Example: `bug,ui,@high` (optional)
    sort := "sort_example" // string | What to sort results by. Can be either `created`, `updated`, `comments`. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListForAuthenticatedUser(context.Background()).Filter(filter).State(state).Labels(labels).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListForAuthenticatedUser`: []Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Indicates which sorts of issues to return. &#x60;assigned&#x60; means issues assigned to you. &#x60;created&#x60; means issues created by you. &#x60;mentioned&#x60; means issues mentioning you. &#x60;subscribed&#x60; means issues you&#39;re subscribed to updates for. &#x60;all&#x60; or &#x60;repos&#x60; means all issues you can see, regardless of participation or creation. | [default to &quot;assigned&quot;]
 **state** | **string** | Indicates the state of the issues to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **labels** | **string** | A list of comma separated label names. Example: &#x60;bug,ui,@high&#x60; | 
 **sort** | **string** | What to sort results by. Can be either &#x60;created&#x60;, &#x60;updated&#x60;, &#x60;comments&#x60;. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListForOrg

> []Issue IssuesListForOrg(ctx, org).Filter(filter).State(state).Labels(labels).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()

List organization issues assigned to the authenticated user



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    filter := "filter_example" // string | Indicates which sorts of issues to return. `assigned` means issues assigned to you. `created` means issues created by you. `mentioned` means issues mentioning you. `subscribed` means issues you're subscribed to updates for. `all` or `repos` means all issues you can see, regardless of participation or creation. (optional) (default to "assigned")
    state := "state_example" // string | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    labels := "labels_example" // string | A list of comma separated label names. Example: `bug,ui,@high` (optional)
    sort := "sort_example" // string | What to sort results by. Can be either `created`, `updated`, `comments`. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListForOrg(context.Background(), org).Filter(filter).State(state).Labels(labels).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListForOrg`: []Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Indicates which sorts of issues to return. &#x60;assigned&#x60; means issues assigned to you. &#x60;created&#x60; means issues created by you. &#x60;mentioned&#x60; means issues mentioning you. &#x60;subscribed&#x60; means issues you&#39;re subscribed to updates for. &#x60;all&#x60; or &#x60;repos&#x60; means all issues you can see, regardless of participation or creation. | [default to &quot;assigned&quot;]
 **state** | **string** | Indicates the state of the issues to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **labels** | **string** | A list of comma separated label names. Example: &#x60;bug,ui,@high&#x60; | 
 **sort** | **string** | What to sort results by. Can be either &#x60;created&#x60;, &#x60;updated&#x60;, &#x60;comments&#x60;. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListForRepo

> []Issue IssuesListForRepo(ctx, owner, repo).Milestone(milestone).State(state).Assignee(assignee).Creator(creator).Mentioned(mentioned).Labels(labels).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()

List repository issues



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
    milestone := "milestone_example" // string | If an `integer` is passed, it should refer to a milestone by its `number` field. If the string `*` is passed, issues with any milestone are accepted. If the string `none` is passed, issues without milestones are returned. (optional)
    state := "state_example" // string | Indicates the state of the issues to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    assignee := "assignee_example" // string | Can be the name of a user. Pass in `none` for issues with no assigned user, and `*` for issues assigned to any user. (optional)
    creator := "creator_example" // string | The user that created the issue. (optional)
    mentioned := "mentioned_example" // string | A user that's mentioned in the issue. (optional)
    labels := "labels_example" // string | A list of comma separated label names. Example: `bug,ui,@high` (optional)
    sort := "sort_example" // string | What to sort results by. Can be either `created`, `updated`, `comments`. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListForRepo(context.Background(), owner, repo).Milestone(milestone).State(state).Assignee(assignee).Creator(creator).Mentioned(mentioned).Labels(labels).Sort(sort).Direction(direction).Since(since).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListForRepo`: []Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **milestone** | **string** | If an &#x60;integer&#x60; is passed, it should refer to a milestone by its &#x60;number&#x60; field. If the string &#x60;*&#x60; is passed, issues with any milestone are accepted. If the string &#x60;none&#x60; is passed, issues without milestones are returned. | 
 **state** | **string** | Indicates the state of the issues to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **assignee** | **string** | Can be the name of a user. Pass in &#x60;none&#x60; for issues with no assigned user, and &#x60;*&#x60; for issues assigned to any user. | 
 **creator** | **string** | The user that created the issue. | 
 **mentioned** | **string** | A user that&#39;s mentioned in the issue. | 
 **labels** | **string** | A list of comma separated label names. Example: &#x60;bug,ui,@high&#x60; | 
 **sort** | **string** | What to sort results by. Can be either &#x60;created&#x60;, &#x60;updated&#x60;, &#x60;comments&#x60;. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListLabelsForMilestone

> []Label IssuesListLabelsForMilestone(ctx, owner, repo, milestoneNumber).PerPage(perPage).Page(page).Execute()

List labels for issues in a milestone



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
    milestoneNumber := int32(56) // int32 | The number that identifies the milestone.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListLabelsForMilestone(context.Background(), owner, repo, milestoneNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListLabelsForMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListLabelsForMilestone`: []Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListLabelsForMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**milestoneNumber** | **int32** | The number that identifies the milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListLabelsForMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListLabelsForRepo

> []Label IssuesListLabelsForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List labels for a repository



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListLabelsForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListLabelsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListLabelsForRepo`: []Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListLabelsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListLabelsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListLabelsOnIssue

> []Label IssuesListLabelsOnIssue(ctx, owner, repo, issueNumber).PerPage(perPage).Page(page).Execute()

List labels for an issue



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListLabelsOnIssue(context.Background(), owner, repo, issueNumber).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListLabelsOnIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListLabelsOnIssue`: []Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListLabelsOnIssue`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesListLabelsOnIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesListMilestones

> []Milestone IssuesListMilestones(ctx, owner, repo).State(state).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List milestones



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
    state := "state_example" // string | The state of the milestone. Either `open`, `closed`, or `all`. (optional) (default to "open")
    sort := "sort_example" // string | What to sort results by. Either `due_on` or `completeness`. (optional) (default to "due_on")
    direction := "direction_example" // string | The direction of the sort. Either `asc` or `desc`. (optional) (default to "asc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesListMilestones(context.Background(), owner, repo).State(state).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesListMilestones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesListMilestones`: []Milestone
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesListMilestones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesListMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | The state of the milestone. Either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **sort** | **string** | What to sort results by. Either &#x60;due_on&#x60; or &#x60;completeness&#x60;. | [default to &quot;due_on&quot;]
 **direction** | **string** | The direction of the sort. Either &#x60;asc&#x60; or &#x60;desc&#x60;. | [default to &quot;asc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesLock

> IssuesLock(ctx, owner, repo, issueNumber).IssuesLockRequest(issuesLockRequest).Execute()

Lock an issue



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
    issuesLockRequest := *openapiclient.NewIssuesLockRequest() // IssuesLockRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesLock(context.Background(), owner, repo, issueNumber).IssuesLockRequest(issuesLockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesLock``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesLockRequest** | [**IssuesLockRequest**](IssuesLockRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesRemoveAllLabels

> IssuesRemoveAllLabels(ctx, owner, repo, issueNumber).Execute()

Remove all labels from an issue



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesRemoveAllLabels(context.Background(), owner, repo, issueNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesRemoveAllLabels``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesRemoveAllLabelsRequest struct via the builder pattern


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


## IssuesRemoveAssignees

> Issue IssuesRemoveAssignees(ctx, owner, repo, issueNumber).IssuesRemoveAssigneesRequest(issuesRemoveAssigneesRequest).Execute()

Remove assignees from an issue



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
    issuesRemoveAssigneesRequest := *openapiclient.NewIssuesRemoveAssigneesRequest() // IssuesRemoveAssigneesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesRemoveAssignees(context.Background(), owner, repo, issueNumber).IssuesRemoveAssigneesRequest(issuesRemoveAssigneesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesRemoveAssignees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesRemoveAssignees`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesRemoveAssignees`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesRemoveAssigneesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesRemoveAssigneesRequest** | [**IssuesRemoveAssigneesRequest**](IssuesRemoveAssigneesRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesRemoveLabel

> []Label IssuesRemoveLabel(ctx, owner, repo, issueNumber, name).Execute()

Remove a label from an issue



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesRemoveLabel(context.Background(), owner, repo, issueNumber, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesRemoveLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesRemoveLabel`: []Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesRemoveLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**issueNumber** | **int32** | The number that identifies the issue. | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesRemoveLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesSetLabels

> []Label IssuesSetLabels(ctx, owner, repo, issueNumber).IssuesSetLabelsRequest(issuesSetLabelsRequest).Execute()

Set labels for an issue



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
    issuesSetLabelsRequest := openapiclient.issues_set_labels_request{IssuesSetLabelsRequestOneOf: openapiclient.NewIssuesSetLabelsRequestOneOf()} // IssuesSetLabelsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesSetLabels(context.Background(), owner, repo, issueNumber).IssuesSetLabelsRequest(issuesSetLabelsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesSetLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesSetLabels`: []Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesSetLabels`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesSetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesSetLabelsRequest** | [**IssuesSetLabelsRequest**](IssuesSetLabelsRequest.md) |  | 

### Return type

[**[]Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesUnlock

> IssuesUnlock(ctx, owner, repo, issueNumber).Execute()

Unlock an issue



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesUnlock(context.Background(), owner, repo, issueNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesUnlock``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesUnlockRequest struct via the builder pattern


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


## IssuesUpdate

> Issue IssuesUpdate(ctx, owner, repo, issueNumber).IssuesUpdateRequest(issuesUpdateRequest).Execute()

Update an issue



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
    issuesUpdateRequest := *openapiclient.NewIssuesUpdateRequest() // IssuesUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesUpdate(context.Background(), owner, repo, issueNumber).IssuesUpdateRequest(issuesUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesUpdate`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesUpdateRequest** | [**IssuesUpdateRequest**](IssuesUpdateRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesUpdateComment

> IssueComment IssuesUpdateComment(ctx, owner, repo, commentId).IssuesUpdateCommentRequest(issuesUpdateCommentRequest).Execute()

Update an issue comment



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
    issuesUpdateCommentRequest := *openapiclient.NewIssuesUpdateCommentRequest("Body_example") // IssuesUpdateCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesUpdateComment(context.Background(), owner, repo, commentId).IssuesUpdateCommentRequest(issuesUpdateCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesUpdateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesUpdateComment`: IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesUpdateComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiIssuesUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesUpdateCommentRequest** | [**IssuesUpdateCommentRequest**](IssuesUpdateCommentRequest.md) |  | 

### Return type

[**IssueComment**](IssueComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesUpdateLabel

> Label IssuesUpdateLabel(ctx, owner, repo, name).IssuesUpdateLabelRequest(issuesUpdateLabelRequest).Execute()

Update a label



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
    name := "name_example" // string | 
    issuesUpdateLabelRequest := *openapiclient.NewIssuesUpdateLabelRequest() // IssuesUpdateLabelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesUpdateLabel(context.Background(), owner, repo, name).IssuesUpdateLabelRequest(issuesUpdateLabelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesUpdateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesUpdateLabel`: Label
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesUpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesUpdateLabelRequest** | [**IssuesUpdateLabelRequest**](IssuesUpdateLabelRequest.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuesUpdateMilestone

> Milestone IssuesUpdateMilestone(ctx, owner, repo, milestoneNumber).IssuesUpdateMilestoneRequest(issuesUpdateMilestoneRequest).Execute()

Update a milestone



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
    milestoneNumber := int32(56) // int32 | The number that identifies the milestone.
    issuesUpdateMilestoneRequest := *openapiclient.NewIssuesUpdateMilestoneRequest() // IssuesUpdateMilestoneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuesApi.IssuesUpdateMilestone(context.Background(), owner, repo, milestoneNumber).IssuesUpdateMilestoneRequest(issuesUpdateMilestoneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuesApi.IssuesUpdateMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuesUpdateMilestone`: Milestone
    fmt.Fprintf(os.Stdout, "Response from `IssuesApi.IssuesUpdateMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**milestoneNumber** | **int32** | The number that identifies the milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssuesUpdateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **issuesUpdateMilestoneRequest** | [**IssuesUpdateMilestoneRequest**](IssuesUpdateMilestoneRequest.md) |  | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

