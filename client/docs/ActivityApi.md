# \ActivityApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivityCheckRepoIsStarredByAuthenticatedUser**](ActivityApi.md#ActivityCheckRepoIsStarredByAuthenticatedUser) | **Get** /user/starred/{owner}/{repo} | Check if a repository is starred by the authenticated user
[**ActivityDeleteRepoSubscription**](ActivityApi.md#ActivityDeleteRepoSubscription) | **Delete** /repos/{owner}/{repo}/subscription | Delete a repository subscription
[**ActivityDeleteThreadSubscription**](ActivityApi.md#ActivityDeleteThreadSubscription) | **Delete** /notifications/threads/{thread_id}/subscription | Delete a thread subscription
[**ActivityGetFeeds**](ActivityApi.md#ActivityGetFeeds) | **Get** /feeds | Get feeds
[**ActivityGetRepoSubscription**](ActivityApi.md#ActivityGetRepoSubscription) | **Get** /repos/{owner}/{repo}/subscription | Get a repository subscription
[**ActivityGetThread**](ActivityApi.md#ActivityGetThread) | **Get** /notifications/threads/{thread_id} | Get a thread
[**ActivityGetThreadSubscriptionForAuthenticatedUser**](ActivityApi.md#ActivityGetThreadSubscriptionForAuthenticatedUser) | **Get** /notifications/threads/{thread_id}/subscription | Get a thread subscription for the authenticated user
[**ActivityListEventsForAuthenticatedUser**](ActivityApi.md#ActivityListEventsForAuthenticatedUser) | **Get** /users/{username}/events | List events for the authenticated user
[**ActivityListNotificationsForAuthenticatedUser**](ActivityApi.md#ActivityListNotificationsForAuthenticatedUser) | **Get** /notifications | List notifications for the authenticated user
[**ActivityListOrgEventsForAuthenticatedUser**](ActivityApi.md#ActivityListOrgEventsForAuthenticatedUser) | **Get** /users/{username}/events/orgs/{org} | List organization events for the authenticated user
[**ActivityListPublicEvents**](ActivityApi.md#ActivityListPublicEvents) | **Get** /events | List public events
[**ActivityListPublicEventsForRepoNetwork**](ActivityApi.md#ActivityListPublicEventsForRepoNetwork) | **Get** /networks/{owner}/{repo}/events | List public events for a network of repositories
[**ActivityListPublicEventsForUser**](ActivityApi.md#ActivityListPublicEventsForUser) | **Get** /users/{username}/events/public | List public events for a user
[**ActivityListPublicOrgEvents**](ActivityApi.md#ActivityListPublicOrgEvents) | **Get** /orgs/{org}/events | List public organization events
[**ActivityListReceivedEventsForUser**](ActivityApi.md#ActivityListReceivedEventsForUser) | **Get** /users/{username}/received_events | List events received by the authenticated user
[**ActivityListReceivedPublicEventsForUser**](ActivityApi.md#ActivityListReceivedPublicEventsForUser) | **Get** /users/{username}/received_events/public | List public events received by a user
[**ActivityListRepoEvents**](ActivityApi.md#ActivityListRepoEvents) | **Get** /repos/{owner}/{repo}/events | List repository events
[**ActivityListRepoNotificationsForAuthenticatedUser**](ActivityApi.md#ActivityListRepoNotificationsForAuthenticatedUser) | **Get** /repos/{owner}/{repo}/notifications | List repository notifications for the authenticated user
[**ActivityListReposStarredByAuthenticatedUser**](ActivityApi.md#ActivityListReposStarredByAuthenticatedUser) | **Get** /user/starred | List repositories starred by the authenticated user
[**ActivityListReposStarredByUser**](ActivityApi.md#ActivityListReposStarredByUser) | **Get** /users/{username}/starred | List repositories starred by a user
[**ActivityListReposWatchedByUser**](ActivityApi.md#ActivityListReposWatchedByUser) | **Get** /users/{username}/subscriptions | List repositories watched by a user
[**ActivityListStargazersForRepo**](ActivityApi.md#ActivityListStargazersForRepo) | **Get** /repos/{owner}/{repo}/stargazers | List stargazers
[**ActivityListWatchedReposForAuthenticatedUser**](ActivityApi.md#ActivityListWatchedReposForAuthenticatedUser) | **Get** /user/subscriptions | List repositories watched by the authenticated user
[**ActivityListWatchersForRepo**](ActivityApi.md#ActivityListWatchersForRepo) | **Get** /repos/{owner}/{repo}/subscribers | List watchers
[**ActivityMarkNotificationsAsRead**](ActivityApi.md#ActivityMarkNotificationsAsRead) | **Put** /notifications | Mark notifications as read
[**ActivityMarkRepoNotificationsAsRead**](ActivityApi.md#ActivityMarkRepoNotificationsAsRead) | **Put** /repos/{owner}/{repo}/notifications | Mark repository notifications as read
[**ActivityMarkThreadAsRead**](ActivityApi.md#ActivityMarkThreadAsRead) | **Patch** /notifications/threads/{thread_id} | Mark a thread as read
[**ActivitySetRepoSubscription**](ActivityApi.md#ActivitySetRepoSubscription) | **Put** /repos/{owner}/{repo}/subscription | Set a repository subscription
[**ActivitySetThreadSubscription**](ActivityApi.md#ActivitySetThreadSubscription) | **Put** /notifications/threads/{thread_id}/subscription | Set a thread subscription
[**ActivityStarRepoForAuthenticatedUser**](ActivityApi.md#ActivityStarRepoForAuthenticatedUser) | **Put** /user/starred/{owner}/{repo} | Star a repository for the authenticated user
[**ActivityUnstarRepoForAuthenticatedUser**](ActivityApi.md#ActivityUnstarRepoForAuthenticatedUser) | **Delete** /user/starred/{owner}/{repo} | Unstar a repository for the authenticated user



## ActivityCheckRepoIsStarredByAuthenticatedUser

> ActivityCheckRepoIsStarredByAuthenticatedUser(ctx, owner, repo).Execute()

Check if a repository is starred by the authenticated user



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
    resp, r, err := apiClient.ActivityApi.ActivityCheckRepoIsStarredByAuthenticatedUser(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityCheckRepoIsStarredByAuthenticatedUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivityCheckRepoIsStarredByAuthenticatedUserRequest struct via the builder pattern


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


## ActivityDeleteRepoSubscription

> ActivityDeleteRepoSubscription(ctx, owner, repo).Execute()

Delete a repository subscription



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
    resp, r, err := apiClient.ActivityApi.ActivityDeleteRepoSubscription(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityDeleteRepoSubscription``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivityDeleteRepoSubscriptionRequest struct via the builder pattern


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


## ActivityDeleteThreadSubscription

> ActivityDeleteThreadSubscription(ctx, threadId).Execute()

Delete a thread subscription



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
    threadId := int32(56) // int32 | The unique identifier of the pull request thread.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityDeleteThreadSubscription(context.Background(), threadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityDeleteThreadSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int32** | The unique identifier of the pull request thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityDeleteThreadSubscriptionRequest struct via the builder pattern


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


## ActivityGetFeeds

> Feed ActivityGetFeeds(ctx).Execute()

Get feeds



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
    resp, r, err := apiClient.ActivityApi.ActivityGetFeeds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityGetFeeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityGetFeeds`: Feed
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityGetFeeds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivityGetFeedsRequest struct via the builder pattern


### Return type

[**Feed**](Feed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityGetRepoSubscription

> RepositorySubscription ActivityGetRepoSubscription(ctx, owner, repo).Execute()

Get a repository subscription



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
    resp, r, err := apiClient.ActivityApi.ActivityGetRepoSubscription(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityGetRepoSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityGetRepoSubscription`: RepositorySubscription
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityGetRepoSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityGetRepoSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositorySubscription**](RepositorySubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityGetThread

> Thread ActivityGetThread(ctx, threadId).Execute()

Get a thread



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
    threadId := int32(56) // int32 | The unique identifier of the pull request thread.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityGetThread(context.Background(), threadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityGetThread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityGetThread`: Thread
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityGetThread`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int32** | The unique identifier of the pull request thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityGetThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityGetThreadSubscriptionForAuthenticatedUser

> ThreadSubscription ActivityGetThreadSubscriptionForAuthenticatedUser(ctx, threadId).Execute()

Get a thread subscription for the authenticated user



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
    threadId := int32(56) // int32 | The unique identifier of the pull request thread.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityGetThreadSubscriptionForAuthenticatedUser(context.Background(), threadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityGetThreadSubscriptionForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityGetThreadSubscriptionForAuthenticatedUser`: ThreadSubscription
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityGetThreadSubscriptionForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int32** | The unique identifier of the pull request thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityGetThreadSubscriptionForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThreadSubscription**](ThreadSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListEventsForAuthenticatedUser

> []Event ActivityListEventsForAuthenticatedUser(ctx, username).PerPage(perPage).Page(page).Execute()

List events for the authenticated user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListEventsForAuthenticatedUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListEventsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListEventsForAuthenticatedUser`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListEventsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListEventsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListNotificationsForAuthenticatedUser

> []Thread ActivityListNotificationsForAuthenticatedUser(ctx).All(all).Participating(participating).Since(since).Before(before).Page(page).PerPage(perPage).Execute()

List notifications for the authenticated user



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
    all := true // bool | If `true`, show notifications marked as read. (optional) (default to false)
    participating := true // bool | If `true`, only shows notifications in which the user is directly participating or mentioned. (optional) (default to false)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    before := time.Now() // time.Time | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 50). (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListNotificationsForAuthenticatedUser(context.Background()).All(all).Participating(participating).Since(since).Before(before).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListNotificationsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListNotificationsForAuthenticatedUser`: []Thread
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListNotificationsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityListNotificationsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** | If &#x60;true&#x60;, show notifications marked as read. | [default to false]
 **participating** | **bool** | If &#x60;true&#x60;, only shows notifications in which the user is directly participating or mentioned. | [default to false]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **before** | **time.Time** | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 50). | [default to 50]

### Return type

[**[]Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListOrgEventsForAuthenticatedUser

> []Event ActivityListOrgEventsForAuthenticatedUser(ctx, username, org).PerPage(perPage).Page(page).Execute()

List organization events for the authenticated user



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
    org := "org_example" // string | The organization name. The name is not case sensitive.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListOrgEventsForAuthenticatedUser(context.Background(), username, org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListOrgEventsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListOrgEventsForAuthenticatedUser`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListOrgEventsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListOrgEventsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListPublicEvents

> []Event ActivityListPublicEvents(ctx).PerPage(perPage).Page(page).Execute()

List public events



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListPublicEvents(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListPublicEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListPublicEvents`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListPublicEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityListPublicEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListPublicEventsForRepoNetwork

> []Event ActivityListPublicEventsForRepoNetwork(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List public events for a network of repositories



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
    resp, r, err := apiClient.ActivityApi.ActivityListPublicEventsForRepoNetwork(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListPublicEventsForRepoNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListPublicEventsForRepoNetwork`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListPublicEventsForRepoNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListPublicEventsForRepoNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListPublicEventsForUser

> []Event ActivityListPublicEventsForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List public events for a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListPublicEventsForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListPublicEventsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListPublicEventsForUser`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListPublicEventsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListPublicEventsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListPublicOrgEvents

> []Event ActivityListPublicOrgEvents(ctx, org).PerPage(perPage).Page(page).Execute()

List public organization events



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
    resp, r, err := apiClient.ActivityApi.ActivityListPublicOrgEvents(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListPublicOrgEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListPublicOrgEvents`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListPublicOrgEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListPublicOrgEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListReceivedEventsForUser

> []Event ActivityListReceivedEventsForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List events received by the authenticated user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListReceivedEventsForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListReceivedEventsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListReceivedEventsForUser`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListReceivedEventsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListReceivedEventsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListReceivedPublicEventsForUser

> []Event ActivityListReceivedPublicEventsForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List public events received by a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListReceivedPublicEventsForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListReceivedPublicEventsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListReceivedPublicEventsForUser`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListReceivedPublicEventsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListReceivedPublicEventsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListRepoEvents

> []Event ActivityListRepoEvents(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository events



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
    resp, r, err := apiClient.ActivityApi.ActivityListRepoEvents(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListRepoEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListRepoEvents`: []Event
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListRepoEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListRepoEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListRepoNotificationsForAuthenticatedUser

> []Thread ActivityListRepoNotificationsForAuthenticatedUser(ctx, owner, repo).All(all).Participating(participating).Since(since).Before(before).PerPage(perPage).Page(page).Execute()

List repository notifications for the authenticated user



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
    all := true // bool | If `true`, show notifications marked as read. (optional) (default to false)
    participating := true // bool | If `true`, only shows notifications in which the user is directly participating or mentioned. (optional) (default to false)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    before := time.Now() // time.Time | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListRepoNotificationsForAuthenticatedUser(context.Background(), owner, repo).All(all).Participating(participating).Since(since).Before(before).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListRepoNotificationsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListRepoNotificationsForAuthenticatedUser`: []Thread
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListRepoNotificationsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListRepoNotificationsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **all** | **bool** | If &#x60;true&#x60;, show notifications marked as read. | [default to false]
 **participating** | **bool** | If &#x60;true&#x60;, only shows notifications in which the user is directly participating or mentioned. | [default to false]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **before** | **time.Time** | Only show notifications updated before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Thread**](Thread.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListReposStarredByAuthenticatedUser

> []Repository ActivityListReposStarredByAuthenticatedUser(ctx).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List repositories starred by the authenticated user



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
    sort := "sort_example" // string | The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListReposStarredByAuthenticatedUser(context.Background()).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListReposStarredByAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListReposStarredByAuthenticatedUser`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListReposStarredByAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityListReposStarredByAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the repository was starred. &#x60;updated&#x60; means when the repository was last pushed to. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.github.v3.star+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListReposStarredByUser

> ActivityListReposStarredByUser200Response ActivityListReposStarredByUser(ctx, username).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List repositories starred by a user



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
    sort := "sort_example" // string | The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListReposStarredByUser(context.Background(), username).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListReposStarredByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListReposStarredByUser`: ActivityListReposStarredByUser200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListReposStarredByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListReposStarredByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the repository was starred. &#x60;updated&#x60; means when the repository was last pushed to. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActivityListReposStarredByUser200Response**](ActivityListReposStarredByUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListReposWatchedByUser

> []MinimalRepository ActivityListReposWatchedByUser(ctx, username).PerPage(perPage).Page(page).Execute()

List repositories watched by a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListReposWatchedByUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListReposWatchedByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListReposWatchedByUser`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListReposWatchedByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListReposWatchedByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListStargazersForRepo

> ActivityListStargazersForRepo200Response ActivityListStargazersForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List stargazers



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
    resp, r, err := apiClient.ActivityApi.ActivityListStargazersForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListStargazersForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListStargazersForRepo`: ActivityListStargazersForRepo200Response
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListStargazersForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListStargazersForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**ActivityListStargazersForRepo200Response**](ActivityListStargazersForRepo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListWatchedReposForAuthenticatedUser

> []MinimalRepository ActivityListWatchedReposForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List repositories watched by the authenticated user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityListWatchedReposForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListWatchedReposForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListWatchedReposForAuthenticatedUser`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListWatchedReposForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityListWatchedReposForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MinimalRepository**](MinimalRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityListWatchersForRepo

> []SimpleUser ActivityListWatchersForRepo(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List watchers



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
    resp, r, err := apiClient.ActivityApi.ActivityListWatchersForRepo(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityListWatchersForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityListWatchersForRepo`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityListWatchersForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityListWatchersForRepoRequest struct via the builder pattern


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


## ActivityMarkNotificationsAsRead

> ActivityMarkNotificationsAsRead202Response ActivityMarkNotificationsAsRead(ctx).ActivityMarkNotificationsAsReadRequest(activityMarkNotificationsAsReadRequest).Execute()

Mark notifications as read



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
    activityMarkNotificationsAsReadRequest := *openapiclient.NewActivityMarkNotificationsAsReadRequest() // ActivityMarkNotificationsAsReadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityMarkNotificationsAsRead(context.Background()).ActivityMarkNotificationsAsReadRequest(activityMarkNotificationsAsReadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityMarkNotificationsAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityMarkNotificationsAsRead`: ActivityMarkNotificationsAsRead202Response
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityMarkNotificationsAsRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityMarkNotificationsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityMarkNotificationsAsReadRequest** | [**ActivityMarkNotificationsAsReadRequest**](ActivityMarkNotificationsAsReadRequest.md) |  | 

### Return type

[**ActivityMarkNotificationsAsRead202Response**](ActivityMarkNotificationsAsRead202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityMarkRepoNotificationsAsRead

> ActivityMarkRepoNotificationsAsRead202Response ActivityMarkRepoNotificationsAsRead(ctx, owner, repo).ActivityMarkRepoNotificationsAsReadRequest(activityMarkRepoNotificationsAsReadRequest).Execute()

Mark repository notifications as read



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
    activityMarkRepoNotificationsAsReadRequest := *openapiclient.NewActivityMarkRepoNotificationsAsReadRequest() // ActivityMarkRepoNotificationsAsReadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityMarkRepoNotificationsAsRead(context.Background(), owner, repo).ActivityMarkRepoNotificationsAsReadRequest(activityMarkRepoNotificationsAsReadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityMarkRepoNotificationsAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityMarkRepoNotificationsAsRead`: ActivityMarkRepoNotificationsAsRead202Response
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivityMarkRepoNotificationsAsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityMarkRepoNotificationsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activityMarkRepoNotificationsAsReadRequest** | [**ActivityMarkRepoNotificationsAsReadRequest**](ActivityMarkRepoNotificationsAsReadRequest.md) |  | 

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


## ActivityMarkThreadAsRead

> ActivityMarkThreadAsRead(ctx, threadId).Execute()

Mark a thread as read



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
    threadId := int32(56) // int32 | The unique identifier of the pull request thread.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivityMarkThreadAsRead(context.Background(), threadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityMarkThreadAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int32** | The unique identifier of the pull request thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityMarkThreadAsReadRequest struct via the builder pattern


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


## ActivitySetRepoSubscription

> RepositorySubscription ActivitySetRepoSubscription(ctx, owner, repo).ActivitySetRepoSubscriptionRequest(activitySetRepoSubscriptionRequest).Execute()

Set a repository subscription



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
    activitySetRepoSubscriptionRequest := *openapiclient.NewActivitySetRepoSubscriptionRequest() // ActivitySetRepoSubscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivitySetRepoSubscription(context.Background(), owner, repo).ActivitySetRepoSubscriptionRequest(activitySetRepoSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivitySetRepoSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivitySetRepoSubscription`: RepositorySubscription
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivitySetRepoSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivitySetRepoSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activitySetRepoSubscriptionRequest** | [**ActivitySetRepoSubscriptionRequest**](ActivitySetRepoSubscriptionRequest.md) |  | 

### Return type

[**RepositorySubscription**](RepositorySubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivitySetThreadSubscription

> ThreadSubscription ActivitySetThreadSubscription(ctx, threadId).ActivitySetThreadSubscriptionRequest(activitySetThreadSubscriptionRequest).Execute()

Set a thread subscription



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
    threadId := int32(56) // int32 | The unique identifier of the pull request thread.
    activitySetThreadSubscriptionRequest := *openapiclient.NewActivitySetThreadSubscriptionRequest() // ActivitySetThreadSubscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ActivitySetThreadSubscription(context.Background(), threadId).ActivitySetThreadSubscriptionRequest(activitySetThreadSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivitySetThreadSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivitySetThreadSubscription`: ThreadSubscription
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ActivitySetThreadSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **int32** | The unique identifier of the pull request thread. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivitySetThreadSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activitySetThreadSubscriptionRequest** | [**ActivitySetThreadSubscriptionRequest**](ActivitySetThreadSubscriptionRequest.md) |  | 

### Return type

[**ThreadSubscription**](ThreadSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivityStarRepoForAuthenticatedUser

> ActivityStarRepoForAuthenticatedUser(ctx, owner, repo).Execute()

Star a repository for the authenticated user



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
    resp, r, err := apiClient.ActivityApi.ActivityStarRepoForAuthenticatedUser(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityStarRepoForAuthenticatedUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivityStarRepoForAuthenticatedUserRequest struct via the builder pattern


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


## ActivityUnstarRepoForAuthenticatedUser

> ActivityUnstarRepoForAuthenticatedUser(ctx, owner, repo).Execute()

Unstar a repository for the authenticated user



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
    resp, r, err := apiClient.ActivityApi.ActivityUnstarRepoForAuthenticatedUser(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ActivityUnstarRepoForAuthenticatedUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivityUnstarRepoForAuthenticatedUserRequest struct via the builder pattern


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

