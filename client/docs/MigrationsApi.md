# \MigrationsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrationsCancelImport**](MigrationsApi.md#MigrationsCancelImport) | **Delete** /repos/{owner}/{repo}/import | Cancel an import
[**MigrationsDeleteArchiveForAuthenticatedUser**](MigrationsApi.md#MigrationsDeleteArchiveForAuthenticatedUser) | **Delete** /user/migrations/{migration_id}/archive | Delete a user migration archive
[**MigrationsDeleteArchiveForOrg**](MigrationsApi.md#MigrationsDeleteArchiveForOrg) | **Delete** /orgs/{org}/migrations/{migration_id}/archive | Delete an organization migration archive
[**MigrationsDownloadArchiveForOrg**](MigrationsApi.md#MigrationsDownloadArchiveForOrg) | **Get** /orgs/{org}/migrations/{migration_id}/archive | Download an organization migration archive
[**MigrationsGetArchiveForAuthenticatedUser**](MigrationsApi.md#MigrationsGetArchiveForAuthenticatedUser) | **Get** /user/migrations/{migration_id}/archive | Download a user migration archive
[**MigrationsGetCommitAuthors**](MigrationsApi.md#MigrationsGetCommitAuthors) | **Get** /repos/{owner}/{repo}/import/authors | Get commit authors
[**MigrationsGetImportStatus**](MigrationsApi.md#MigrationsGetImportStatus) | **Get** /repos/{owner}/{repo}/import | Get an import status
[**MigrationsGetLargeFiles**](MigrationsApi.md#MigrationsGetLargeFiles) | **Get** /repos/{owner}/{repo}/import/large_files | Get large files
[**MigrationsGetStatusForAuthenticatedUser**](MigrationsApi.md#MigrationsGetStatusForAuthenticatedUser) | **Get** /user/migrations/{migration_id} | Get a user migration status
[**MigrationsGetStatusForOrg**](MigrationsApi.md#MigrationsGetStatusForOrg) | **Get** /orgs/{org}/migrations/{migration_id} | Get an organization migration status
[**MigrationsListForAuthenticatedUser**](MigrationsApi.md#MigrationsListForAuthenticatedUser) | **Get** /user/migrations | List user migrations
[**MigrationsListForOrg**](MigrationsApi.md#MigrationsListForOrg) | **Get** /orgs/{org}/migrations | List organization migrations
[**MigrationsListReposForAuthenticatedUser**](MigrationsApi.md#MigrationsListReposForAuthenticatedUser) | **Get** /user/migrations/{migration_id}/repositories | List repositories for a user migration
[**MigrationsListReposForOrg**](MigrationsApi.md#MigrationsListReposForOrg) | **Get** /orgs/{org}/migrations/{migration_id}/repositories | List repositories in an organization migration
[**MigrationsMapCommitAuthor**](MigrationsApi.md#MigrationsMapCommitAuthor) | **Patch** /repos/{owner}/{repo}/import/authors/{author_id} | Map a commit author
[**MigrationsSetLfsPreference**](MigrationsApi.md#MigrationsSetLfsPreference) | **Patch** /repos/{owner}/{repo}/import/lfs | Update Git LFS preference
[**MigrationsStartForAuthenticatedUser**](MigrationsApi.md#MigrationsStartForAuthenticatedUser) | **Post** /user/migrations | Start a user migration
[**MigrationsStartForOrg**](MigrationsApi.md#MigrationsStartForOrg) | **Post** /orgs/{org}/migrations | Start an organization migration
[**MigrationsStartImport**](MigrationsApi.md#MigrationsStartImport) | **Put** /repos/{owner}/{repo}/import | Start an import
[**MigrationsUnlockRepoForAuthenticatedUser**](MigrationsApi.md#MigrationsUnlockRepoForAuthenticatedUser) | **Delete** /user/migrations/{migration_id}/repos/{repo_name}/lock | Unlock a user repository
[**MigrationsUnlockRepoForOrg**](MigrationsApi.md#MigrationsUnlockRepoForOrg) | **Delete** /orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock | Unlock an organization repository
[**MigrationsUpdateImport**](MigrationsApi.md#MigrationsUpdateImport) | **Patch** /repos/{owner}/{repo}/import | Update an import



## MigrationsCancelImport

> MigrationsCancelImport(ctx, owner, repo).Execute()

Cancel an import



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
    resp, r, err := apiClient.MigrationsApi.MigrationsCancelImport(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsCancelImport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMigrationsCancelImportRequest struct via the builder pattern


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


## MigrationsDeleteArchiveForAuthenticatedUser

> MigrationsDeleteArchiveForAuthenticatedUser(ctx, migrationId).Execute()

Delete a user migration archive



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsDeleteArchiveForAuthenticatedUser(context.Background(), migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsDeleteArchiveForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsDeleteArchiveForAuthenticatedUserRequest struct via the builder pattern


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


## MigrationsDeleteArchiveForOrg

> MigrationsDeleteArchiveForOrg(ctx, org, migrationId).Execute()

Delete an organization migration archive



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsDeleteArchiveForOrg(context.Background(), org, migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsDeleteArchiveForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsDeleteArchiveForOrgRequest struct via the builder pattern


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


## MigrationsDownloadArchiveForOrg

> MigrationsDownloadArchiveForOrg(ctx, org, migrationId).Execute()

Download an organization migration archive



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsDownloadArchiveForOrg(context.Background(), org, migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsDownloadArchiveForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsDownloadArchiveForOrgRequest struct via the builder pattern


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


## MigrationsGetArchiveForAuthenticatedUser

> MigrationsGetArchiveForAuthenticatedUser(ctx, migrationId).Execute()

Download a user migration archive



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsGetArchiveForAuthenticatedUser(context.Background(), migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsGetArchiveForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsGetArchiveForAuthenticatedUserRequest struct via the builder pattern


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


## MigrationsGetCommitAuthors

> []PorterAuthor MigrationsGetCommitAuthors(ctx, owner, repo).Since(since).Execute()

Get commit authors



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
    since := int32(56) // int32 | A user ID. Only return users with an ID greater than this ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsGetCommitAuthors(context.Background(), owner, repo).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsGetCommitAuthors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsGetCommitAuthors`: []PorterAuthor
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsGetCommitAuthors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsGetCommitAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **since** | **int32** | A user ID. Only return users with an ID greater than this ID. | 

### Return type

[**[]PorterAuthor**](PorterAuthor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsGetImportStatus

> ModelImport MigrationsGetImportStatus(ctx, owner, repo).Execute()

Get an import status



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
    resp, r, err := apiClient.MigrationsApi.MigrationsGetImportStatus(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsGetImportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsGetImportStatus`: ModelImport
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsGetImportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsGetImportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelImport**](ModelImport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsGetLargeFiles

> []PorterLargeFile MigrationsGetLargeFiles(ctx, owner, repo).Execute()

Get large files



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
    resp, r, err := apiClient.MigrationsApi.MigrationsGetLargeFiles(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsGetLargeFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsGetLargeFiles`: []PorterLargeFile
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsGetLargeFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsGetLargeFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PorterLargeFile**](PorterLargeFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsGetStatusForAuthenticatedUser

> Migration MigrationsGetStatusForAuthenticatedUser(ctx, migrationId).Exclude(exclude).Execute()

Get a user migration status



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.
    exclude := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsGetStatusForAuthenticatedUser(context.Background(), migrationId).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsGetStatusForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsGetStatusForAuthenticatedUser`: Migration
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsGetStatusForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsGetStatusForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclude** | **[]string** |  | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsGetStatusForOrg

> Migration MigrationsGetStatusForOrg(ctx, org, migrationId).Exclude(exclude).Execute()

Get an organization migration status



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.
    exclude := []string{"repositories"} // []string | Exclude attributes from the API response to improve performance (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsGetStatusForOrg(context.Background(), org, migrationId).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsGetStatusForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsGetStatusForOrg`: Migration
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsGetStatusForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsGetStatusForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclude** | **[]string** | Exclude attributes from the API response to improve performance | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsListForAuthenticatedUser

> []Migration MigrationsListForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List user migrations



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
    resp, r, err := apiClient.MigrationsApi.MigrationsListForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsListForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsListForAuthenticatedUser`: []Migration
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsListForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsListForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsListForOrg

> []Migration MigrationsListForOrg(ctx, org).PerPage(perPage).Page(page).Exclude(exclude).Execute()

List organization migrations



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
    exclude := []string{"repositories"} // []string | Exclude attributes from the API response to improve performance (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsListForOrg(context.Background(), org).PerPage(perPage).Page(page).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsListForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsListForOrg`: []Migration
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsListForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsListForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **exclude** | **[]string** | Exclude attributes from the API response to improve performance | 

### Return type

[**[]Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsListReposForAuthenticatedUser

> []MinimalRepository MigrationsListReposForAuthenticatedUser(ctx, migrationId).PerPage(perPage).Page(page).Execute()

List repositories for a user migration



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsListReposForAuthenticatedUser(context.Background(), migrationId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsListReposForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsListReposForAuthenticatedUser`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsListReposForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsListReposForAuthenticatedUserRequest struct via the builder pattern


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


## MigrationsListReposForOrg

> []MinimalRepository MigrationsListReposForOrg(ctx, org, migrationId).PerPage(perPage).Page(page).Execute()

List repositories in an organization migration



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsListReposForOrg(context.Background(), org, migrationId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsListReposForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsListReposForOrg`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsListReposForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**migrationId** | **int32** | The unique identifier of the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsListReposForOrgRequest struct via the builder pattern


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


## MigrationsMapCommitAuthor

> PorterAuthor MigrationsMapCommitAuthor(ctx, owner, repo, authorId).MigrationsMapCommitAuthorRequest(migrationsMapCommitAuthorRequest).Execute()

Map a commit author



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
    authorId := int32(56) // int32 | 
    migrationsMapCommitAuthorRequest := *openapiclient.NewMigrationsMapCommitAuthorRequest() // MigrationsMapCommitAuthorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsMapCommitAuthor(context.Background(), owner, repo, authorId).MigrationsMapCommitAuthorRequest(migrationsMapCommitAuthorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsMapCommitAuthor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsMapCommitAuthor`: PorterAuthor
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsMapCommitAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**authorId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsMapCommitAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **migrationsMapCommitAuthorRequest** | [**MigrationsMapCommitAuthorRequest**](MigrationsMapCommitAuthorRequest.md) |  | 

### Return type

[**PorterAuthor**](PorterAuthor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsSetLfsPreference

> ModelImport MigrationsSetLfsPreference(ctx, owner, repo).MigrationsSetLfsPreferenceRequest(migrationsSetLfsPreferenceRequest).Execute()

Update Git LFS preference



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
    migrationsSetLfsPreferenceRequest := *openapiclient.NewMigrationsSetLfsPreferenceRequest("UseLfs_example") // MigrationsSetLfsPreferenceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsSetLfsPreference(context.Background(), owner, repo).MigrationsSetLfsPreferenceRequest(migrationsSetLfsPreferenceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsSetLfsPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsSetLfsPreference`: ModelImport
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsSetLfsPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsSetLfsPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **migrationsSetLfsPreferenceRequest** | [**MigrationsSetLfsPreferenceRequest**](MigrationsSetLfsPreferenceRequest.md) |  | 

### Return type

[**ModelImport**](ModelImport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsStartForAuthenticatedUser

> Migration MigrationsStartForAuthenticatedUser(ctx).MigrationsStartForAuthenticatedUserRequest(migrationsStartForAuthenticatedUserRequest).Execute()

Start a user migration



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
    migrationsStartForAuthenticatedUserRequest := *openapiclient.NewMigrationsStartForAuthenticatedUserRequest([]string{"acme/widgets"}) // MigrationsStartForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsStartForAuthenticatedUser(context.Background()).MigrationsStartForAuthenticatedUserRequest(migrationsStartForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsStartForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsStartForAuthenticatedUser`: Migration
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsStartForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsStartForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationsStartForAuthenticatedUserRequest** | [**MigrationsStartForAuthenticatedUserRequest**](MigrationsStartForAuthenticatedUserRequest.md) |  | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsStartForOrg

> Migration MigrationsStartForOrg(ctx, org).MigrationsStartForOrgRequest(migrationsStartForOrgRequest).Execute()

Start an organization migration



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
    migrationsStartForOrgRequest := *openapiclient.NewMigrationsStartForOrgRequest([]string{"Repositories_example"}) // MigrationsStartForOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsStartForOrg(context.Background(), org).MigrationsStartForOrgRequest(migrationsStartForOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsStartForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsStartForOrg`: Migration
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsStartForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsStartForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **migrationsStartForOrgRequest** | [**MigrationsStartForOrgRequest**](MigrationsStartForOrgRequest.md) |  | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsStartImport

> ModelImport MigrationsStartImport(ctx, owner, repo).MigrationsStartImportRequest(migrationsStartImportRequest).Execute()

Start an import



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
    migrationsStartImportRequest := *openapiclient.NewMigrationsStartImportRequest("VcsUrl_example") // MigrationsStartImportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsStartImport(context.Background(), owner, repo).MigrationsStartImportRequest(migrationsStartImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsStartImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsStartImport`: ModelImport
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsStartImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsStartImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **migrationsStartImportRequest** | [**MigrationsStartImportRequest**](MigrationsStartImportRequest.md) |  | 

### Return type

[**ModelImport**](ModelImport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationsUnlockRepoForAuthenticatedUser

> MigrationsUnlockRepoForAuthenticatedUser(ctx, migrationId, repoName).Execute()

Unlock a user repository



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.
    repoName := "repoName_example" // string | repo_name parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsUnlockRepoForAuthenticatedUser(context.Background(), migrationId, repoName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsUnlockRepoForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**migrationId** | **int32** | The unique identifier of the migration. | 
**repoName** | **string** | repo_name parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsUnlockRepoForAuthenticatedUserRequest struct via the builder pattern


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


## MigrationsUnlockRepoForOrg

> MigrationsUnlockRepoForOrg(ctx, org, migrationId, repoName).Execute()

Unlock an organization repository



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
    migrationId := int32(56) // int32 | The unique identifier of the migration.
    repoName := "repoName_example" // string | repo_name parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsUnlockRepoForOrg(context.Background(), org, migrationId, repoName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsUnlockRepoForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**migrationId** | **int32** | The unique identifier of the migration. | 
**repoName** | **string** | repo_name parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsUnlockRepoForOrgRequest struct via the builder pattern


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


## MigrationsUpdateImport

> ModelImport MigrationsUpdateImport(ctx, owner, repo).MigrationsUpdateImportRequest(migrationsUpdateImportRequest).Execute()

Update an import



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
    migrationsUpdateImportRequest := *openapiclient.NewMigrationsUpdateImportRequest() // MigrationsUpdateImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.MigrationsUpdateImport(context.Background(), owner, repo).MigrationsUpdateImportRequest(migrationsUpdateImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.MigrationsUpdateImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationsUpdateImport`: ModelImport
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.MigrationsUpdateImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationsUpdateImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **migrationsUpdateImportRequest** | [**MigrationsUpdateImportRequest**](MigrationsUpdateImportRequest.md) |  | 

### Return type

[**ModelImport**](ModelImport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

