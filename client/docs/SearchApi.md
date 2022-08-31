# \SearchApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCode**](SearchApi.md#SearchCode) | **Get** /search/code | Search code
[**SearchCommits**](SearchApi.md#SearchCommits) | **Get** /search/commits | Search commits
[**SearchIssuesAndPullRequests**](SearchApi.md#SearchIssuesAndPullRequests) | **Get** /search/issues | Search issues and pull requests
[**SearchLabels**](SearchApi.md#SearchLabels) | **Get** /search/labels | Search labels
[**SearchRepos**](SearchApi.md#SearchRepos) | **Get** /search/repositories | Search repositories
[**SearchTopics**](SearchApi.md#SearchTopics) | **Get** /search/topics | Search topics
[**SearchUsers**](SearchApi.md#SearchUsers) | **Get** /search/users | Search users



## SearchCode

> SearchCode200Response SearchCode(ctx).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()

Search code



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
    q := "q_example" // string | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \"[Searching code](https://docs.github.com/search-github/searching-on-github/searching-code)\" for a detailed list of qualifiers.
    sort := "sort_example" // string | Sorts the results of your query. Can only be `indexed`, which indicates how recently a file has been indexed by the GitHub search infrastructure. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) (optional)
    order := "order_example" // string | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchCode(context.Background()).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCode`: SearchCode200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \&quot;[Searching code](https://docs.github.com/search-github/searching-on-github/searching-code)\&quot; for a detailed list of qualifiers. | 
 **sort** | **string** | Sorts the results of your query. Can only be &#x60;indexed&#x60;, which indicates how recently a file has been indexed by the GitHub search infrastructure. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) | 
 **order** | **string** | Determines whether the first search result returned is the highest number of matches (&#x60;desc&#x60;) or lowest number of matches (&#x60;asc&#x60;). This parameter is ignored unless you provide &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchCode200Response**](SearchCode200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCommits

> SearchCommits200Response SearchCommits(ctx).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()

Search commits



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
    q := "q_example" // string | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \"[Searching commits](https://docs.github.com/search-github/searching-on-github/searching-commits)\" for a detailed list of qualifiers.
    sort := "sort_example" // string | Sorts the results of your query by `author-date` or `committer-date`. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) (optional)
    order := "order_example" // string | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchCommits(context.Background()).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCommits`: SearchCommits200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \&quot;[Searching commits](https://docs.github.com/search-github/searching-on-github/searching-commits)\&quot; for a detailed list of qualifiers. | 
 **sort** | **string** | Sorts the results of your query by &#x60;author-date&#x60; or &#x60;committer-date&#x60;. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) | 
 **order** | **string** | Determines whether the first search result returned is the highest number of matches (&#x60;desc&#x60;) or lowest number of matches (&#x60;asc&#x60;). This parameter is ignored unless you provide &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchCommits200Response**](SearchCommits200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIssuesAndPullRequests

> SearchIssuesAndPullRequests200Response SearchIssuesAndPullRequests(ctx).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()

Search issues and pull requests



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
    q := "q_example" // string | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \"[Searching issues and pull requests](https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests)\" for a detailed list of qualifiers.
    sort := "sort_example" // string | Sorts the results of your query by the number of `comments`, `reactions`, `reactions-+1`, `reactions--1`, `reactions-smile`, `reactions-thinking_face`, `reactions-heart`, `reactions-tada`, or `interactions`. You can also sort results by how recently the items were `created` or `updated`, Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) (optional)
    order := "order_example" // string | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchIssuesAndPullRequests(context.Background()).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchIssuesAndPullRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIssuesAndPullRequests`: SearchIssuesAndPullRequests200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchIssuesAndPullRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIssuesAndPullRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \&quot;[Searching issues and pull requests](https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests)\&quot; for a detailed list of qualifiers. | 
 **sort** | **string** | Sorts the results of your query by the number of &#x60;comments&#x60;, &#x60;reactions&#x60;, &#x60;reactions-+1&#x60;, &#x60;reactions--1&#x60;, &#x60;reactions-smile&#x60;, &#x60;reactions-thinking_face&#x60;, &#x60;reactions-heart&#x60;, &#x60;reactions-tada&#x60;, or &#x60;interactions&#x60;. You can also sort results by how recently the items were &#x60;created&#x60; or &#x60;updated&#x60;, Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) | 
 **order** | **string** | Determines whether the first search result returned is the highest number of matches (&#x60;desc&#x60;) or lowest number of matches (&#x60;asc&#x60;). This parameter is ignored unless you provide &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchIssuesAndPullRequests200Response**](SearchIssuesAndPullRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLabels

> SearchLabels200Response SearchLabels(ctx).RepositoryId(repositoryId).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()

Search labels



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
    repositoryId := int32(56) // int32 | The id of the repository.
    q := "q_example" // string | The search keywords. This endpoint does not accept qualifiers in the query. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query).
    sort := "sort_example" // string | Sorts the results of your query by when the label was `created` or `updated`. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) (optional)
    order := "order_example" // string | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchLabels(context.Background()).RepositoryId(repositoryId).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchLabels`: SearchLabels200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryId** | **int32** | The id of the repository. | 
 **q** | **string** | The search keywords. This endpoint does not accept qualifiers in the query. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). | 
 **sort** | **string** | Sorts the results of your query by when the label was &#x60;created&#x60; or &#x60;updated&#x60;. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) | 
 **order** | **string** | Determines whether the first search result returned is the highest number of matches (&#x60;desc&#x60;) or lowest number of matches (&#x60;asc&#x60;). This parameter is ignored unless you provide &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchLabels200Response**](SearchLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRepos

> SearchRepos200Response SearchRepos(ctx).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()

Search repositories



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
    q := "q_example" // string | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \"[Searching for repositories](https://docs.github.com/articles/searching-for-repositories/)\" for a detailed list of qualifiers.
    sort := "sort_example" // string | Sorts the results of your query by number of `stars`, `forks`, or `help-wanted-issues` or how recently the items were `updated`. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) (optional)
    order := "order_example" // string | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchRepos(context.Background()).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRepos`: SearchRepos200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \&quot;[Searching for repositories](https://docs.github.com/articles/searching-for-repositories/)\&quot; for a detailed list of qualifiers. | 
 **sort** | **string** | Sorts the results of your query by number of &#x60;stars&#x60;, &#x60;forks&#x60;, or &#x60;help-wanted-issues&#x60; or how recently the items were &#x60;updated&#x60;. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) | 
 **order** | **string** | Determines whether the first search result returned is the highest number of matches (&#x60;desc&#x60;) or lowest number of matches (&#x60;asc&#x60;). This parameter is ignored unless you provide &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchRepos200Response**](SearchRepos200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTopics

> SearchTopics200Response SearchTopics(ctx).Q(q).PerPage(perPage).Page(page).Execute()

Search topics



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
    q := "q_example" // string | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query).
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchTopics(context.Background()).Q(q).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTopics`: SearchTopics200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchTopics200Response**](SearchTopics200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> SearchUsers200Response SearchUsers(ctx).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()

Search users



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
    q := "q_example" // string | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \"[Searching users](https://docs.github.com/search-github/searching-on-github/searching-users)\" for a detailed list of qualifiers.
    sort := "sort_example" // string | Sorts the results of your query by number of `followers` or `repositories`, or when the person `joined` GitHub. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) (optional)
    order := "order_example" // string | Determines whether the first search result returned is the highest number of matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored unless you provide `sort`. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchUsers(context.Background()).Q(q).Sort(sort).Order(order).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsers`: SearchUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The query contains one or more search keywords and qualifiers. Qualifiers allow you to limit your search to specific areas of GitHub. The REST API supports the same qualifiers as the web interface for GitHub. To learn more about the format of the query, see [Constructing a search query](https://docs.github.com/rest/reference/search#constructing-a-search-query). See \&quot;[Searching users](https://docs.github.com/search-github/searching-on-github/searching-users)\&quot; for a detailed list of qualifiers. | 
 **sort** | **string** | Sorts the results of your query by number of &#x60;followers&#x60; or &#x60;repositories&#x60;, or when the person &#x60;joined&#x60; GitHub. Default: [best match](https://docs.github.com/rest/reference/search#ranking-search-results) | 
 **order** | **string** | Determines whether the first search result returned is the highest number of matches (&#x60;desc&#x60;) or lowest number of matches (&#x60;asc&#x60;). This parameter is ignored unless you provide &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**SearchUsers200Response**](SearchUsers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

