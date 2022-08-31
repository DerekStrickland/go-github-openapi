# \DependencyGraphApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DependencyGraphCreateRepositorySnapshot**](DependencyGraphApi.md#DependencyGraphCreateRepositorySnapshot) | **Post** /repos/{owner}/{repo}/dependency-graph/snapshots | Create a snapshot of dependencies for a repository
[**DependencyGraphDiffRange**](DependencyGraphApi.md#DependencyGraphDiffRange) | **Get** /repos/{owner}/{repo}/dependency-graph/compare/{basehead} | Get a diff of the dependencies between commits



## DependencyGraphCreateRepositorySnapshot

> DependencyGraphCreateRepositorySnapshot201Response DependencyGraphCreateRepositorySnapshot(ctx, owner, repo).Snapshot(snapshot).Execute()

Create a snapshot of dependencies for a repository



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
    snapshot := *openapiclient.NewSnapshot(int32(123), *openapiclient.NewSnapshotJob("5622a2b0-63f6-4732-8c34-a1ab27e102a11", "yourworkflowname_yourjobname"), "ddc951f4b1293222421f2c8df679786153acf689", "refs/heads/main", *openapiclient.NewSnapshotDetector("docker buildtime detector", "1.0.0", "http://example.com/docker-buildtimer-detector"), time.Now()) // Snapshot | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependencyGraphApi.DependencyGraphCreateRepositorySnapshot(context.Background(), owner, repo).Snapshot(snapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependencyGraphApi.DependencyGraphCreateRepositorySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependencyGraphCreateRepositorySnapshot`: DependencyGraphCreateRepositorySnapshot201Response
    fmt.Fprintf(os.Stdout, "Response from `DependencyGraphApi.DependencyGraphCreateRepositorySnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependencyGraphCreateRepositorySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snapshot** | [**Snapshot**](Snapshot.md) |  | 

### Return type

[**DependencyGraphCreateRepositorySnapshot201Response**](DependencyGraphCreateRepositorySnapshot201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependencyGraphDiffRange

> []DependencyGraphDiffInner DependencyGraphDiffRange(ctx, owner, repo, basehead).Name(name).Execute()

Get a diff of the dependencies between commits



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
    basehead := "basehead_example" // string | The base and head Git revisions to compare. The Git revisions will be resolved to commit SHAs. Named revisions will be resolved to their corresponding HEAD commits, and an appropriate merge base will be determined. This parameter expects the format `{base}...{head}`.
    name := "name_example" // string | The full path, relative to the repository root, of the dependency manifest file. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependencyGraphApi.DependencyGraphDiffRange(context.Background(), owner, repo, basehead).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependencyGraphApi.DependencyGraphDiffRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependencyGraphDiffRange`: []DependencyGraphDiffInner
    fmt.Fprintf(os.Stdout, "Response from `DependencyGraphApi.DependencyGraphDiffRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**basehead** | **string** | The base and head Git revisions to compare. The Git revisions will be resolved to commit SHAs. Named revisions will be resolved to their corresponding HEAD commits, and an appropriate merge base will be determined. This parameter expects the format &#x60;{base}...{head}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependencyGraphDiffRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **string** | The full path, relative to the repository root, of the dependency manifest file. | 

### Return type

[**[]DependencyGraphDiffInner**](DependencyGraphDiffInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

