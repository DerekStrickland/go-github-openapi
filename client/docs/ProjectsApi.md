# \ProjectsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsAddCollaborator**](ProjectsApi.md#ProjectsAddCollaborator) | **Put** /projects/{project_id}/collaborators/{username} | Add project collaborator
[**ProjectsCreateCard**](ProjectsApi.md#ProjectsCreateCard) | **Post** /projects/columns/{column_id}/cards | Create a project card
[**ProjectsCreateColumn**](ProjectsApi.md#ProjectsCreateColumn) | **Post** /projects/{project_id}/columns | Create a project column
[**ProjectsCreateForAuthenticatedUser**](ProjectsApi.md#ProjectsCreateForAuthenticatedUser) | **Post** /user/projects | Create a user project
[**ProjectsCreateForOrg**](ProjectsApi.md#ProjectsCreateForOrg) | **Post** /orgs/{org}/projects | Create an organization project
[**ProjectsCreateForRepo**](ProjectsApi.md#ProjectsCreateForRepo) | **Post** /repos/{owner}/{repo}/projects | Create a repository project
[**ProjectsDelete**](ProjectsApi.md#ProjectsDelete) | **Delete** /projects/{project_id} | Delete a project
[**ProjectsDeleteCard**](ProjectsApi.md#ProjectsDeleteCard) | **Delete** /projects/columns/cards/{card_id} | Delete a project card
[**ProjectsDeleteColumn**](ProjectsApi.md#ProjectsDeleteColumn) | **Delete** /projects/columns/{column_id} | Delete a project column
[**ProjectsGet**](ProjectsApi.md#ProjectsGet) | **Get** /projects/{project_id} | Get a project
[**ProjectsGetCard**](ProjectsApi.md#ProjectsGetCard) | **Get** /projects/columns/cards/{card_id} | Get a project card
[**ProjectsGetColumn**](ProjectsApi.md#ProjectsGetColumn) | **Get** /projects/columns/{column_id} | Get a project column
[**ProjectsGetPermissionForUser**](ProjectsApi.md#ProjectsGetPermissionForUser) | **Get** /projects/{project_id}/collaborators/{username}/permission | Get project permission for a user
[**ProjectsListCards**](ProjectsApi.md#ProjectsListCards) | **Get** /projects/columns/{column_id}/cards | List project cards
[**ProjectsListCollaborators**](ProjectsApi.md#ProjectsListCollaborators) | **Get** /projects/{project_id}/collaborators | List project collaborators
[**ProjectsListColumns**](ProjectsApi.md#ProjectsListColumns) | **Get** /projects/{project_id}/columns | List project columns
[**ProjectsListForOrg**](ProjectsApi.md#ProjectsListForOrg) | **Get** /orgs/{org}/projects | List organization projects
[**ProjectsListForRepo**](ProjectsApi.md#ProjectsListForRepo) | **Get** /repos/{owner}/{repo}/projects | List repository projects
[**ProjectsListForUser**](ProjectsApi.md#ProjectsListForUser) | **Get** /users/{username}/projects | List user projects
[**ProjectsMoveCard**](ProjectsApi.md#ProjectsMoveCard) | **Post** /projects/columns/cards/{card_id}/moves | Move a project card
[**ProjectsMoveColumn**](ProjectsApi.md#ProjectsMoveColumn) | **Post** /projects/columns/{column_id}/moves | Move a project column
[**ProjectsRemoveCollaborator**](ProjectsApi.md#ProjectsRemoveCollaborator) | **Delete** /projects/{project_id}/collaborators/{username} | Remove user as a collaborator
[**ProjectsUpdate**](ProjectsApi.md#ProjectsUpdate) | **Patch** /projects/{project_id} | Update a project
[**ProjectsUpdateCard**](ProjectsApi.md#ProjectsUpdateCard) | **Patch** /projects/columns/cards/{card_id} | Update an existing project card
[**ProjectsUpdateColumn**](ProjectsApi.md#ProjectsUpdateColumn) | **Patch** /projects/columns/{column_id} | Update an existing project column



## ProjectsAddCollaborator

> ProjectsAddCollaborator(ctx, projectId, username).ProjectsAddCollaboratorRequest(projectsAddCollaboratorRequest).Execute()

Add project collaborator



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    username := "username_example" // string | The handle for the GitHub user account.
    projectsAddCollaboratorRequest := *openapiclient.NewProjectsAddCollaboratorRequest() // ProjectsAddCollaboratorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsAddCollaborator(context.Background(), projectId, username).ProjectsAddCollaboratorRequest(projectsAddCollaboratorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsAddCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsAddCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectsAddCollaboratorRequest** | [**ProjectsAddCollaboratorRequest**](ProjectsAddCollaboratorRequest.md) |  | 

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


## ProjectsCreateCard

> ProjectCard ProjectsCreateCard(ctx, columnId).ProjectsCreateCardRequest(projectsCreateCardRequest).Execute()

Create a project card



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
    columnId := int32(56) // int32 | The unique identifier of the column.
    projectsCreateCardRequest := openapiclient.projects_create_card_request{ProjectsCreateCardRequestOneOf: openapiclient.NewProjectsCreateCardRequestOneOf("Update all gems")} // ProjectsCreateCardRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreateCard(context.Background(), columnId).ProjectsCreateCardRequest(projectsCreateCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreateCard`: ProjectCard
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreateCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** | The unique identifier of the column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsCreateCardRequest** | [**ProjectsCreateCardRequest**](ProjectsCreateCardRequest.md) |  | 

### Return type

[**ProjectCard**](ProjectCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreateColumn

> ProjectColumn ProjectsCreateColumn(ctx, projectId).ProjectsUpdateColumnRequest(projectsUpdateColumnRequest).Execute()

Create a project column



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    projectsUpdateColumnRequest := *openapiclient.NewProjectsUpdateColumnRequest("Remaining tasks") // ProjectsUpdateColumnRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreateColumn(context.Background(), projectId).ProjectsUpdateColumnRequest(projectsUpdateColumnRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreateColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreateColumn`: ProjectColumn
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreateColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUpdateColumnRequest** | [**ProjectsUpdateColumnRequest**](ProjectsUpdateColumnRequest.md) |  | 

### Return type

[**ProjectColumn**](ProjectColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreateForAuthenticatedUser

> Project ProjectsCreateForAuthenticatedUser(ctx).ProjectsCreateForAuthenticatedUserRequest(projectsCreateForAuthenticatedUserRequest).Execute()

Create a user project



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
    projectsCreateForAuthenticatedUserRequest := *openapiclient.NewProjectsCreateForAuthenticatedUserRequest("Week One Sprint") // ProjectsCreateForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreateForAuthenticatedUser(context.Background()).ProjectsCreateForAuthenticatedUserRequest(projectsCreateForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreateForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreateForAuthenticatedUser`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreateForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsCreateForAuthenticatedUserRequest** | [**ProjectsCreateForAuthenticatedUserRequest**](ProjectsCreateForAuthenticatedUserRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreateForOrg

> Project ProjectsCreateForOrg(ctx, org).ProjectsCreateForOrgRequest(projectsCreateForOrgRequest).Execute()

Create an organization project



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
    projectsCreateForOrgRequest := *openapiclient.NewProjectsCreateForOrgRequest("Name_example") // ProjectsCreateForOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreateForOrg(context.Background(), org).ProjectsCreateForOrgRequest(projectsCreateForOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreateForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreateForOrg`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreateForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsCreateForOrgRequest** | [**ProjectsCreateForOrgRequest**](ProjectsCreateForOrgRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreateForRepo

> Project ProjectsCreateForRepo(ctx, owner, repo).ProjectsCreateForOrgRequest(projectsCreateForOrgRequest).Execute()

Create a repository project



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
    projectsCreateForOrgRequest := *openapiclient.NewProjectsCreateForOrgRequest("Name_example") // ProjectsCreateForOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreateForRepo(context.Background(), owner, repo).ProjectsCreateForOrgRequest(projectsCreateForOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreateForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreateForRepo`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreateForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectsCreateForOrgRequest** | [**ProjectsCreateForOrgRequest**](ProjectsCreateForOrgRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDelete

> ProjectsDelete(ctx, projectId).Execute()

Delete a project



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
    projectId := int32(56) // int32 | The unique identifier of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDelete(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


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


## ProjectsDeleteCard

> ProjectsDeleteCard(ctx, cardId).Execute()

Delete a project card



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
    cardId := int32(56) // int32 | The unique identifier of the card.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDeleteCard(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDeleteCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **int32** | The unique identifier of the card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteCardRequest struct via the builder pattern


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


## ProjectsDeleteColumn

> ProjectsDeleteColumn(ctx, columnId).Execute()

Delete a project column



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
    columnId := int32(56) // int32 | The unique identifier of the column.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDeleteColumn(context.Background(), columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDeleteColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** | The unique identifier of the column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteColumnRequest struct via the builder pattern


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


## ProjectsGet

> Project ProjectsGet(ctx, projectId).Execute()

Get a project



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
    projectId := int32(56) // int32 | The unique identifier of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsGet(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsGet`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGetCard

> ProjectCard ProjectsGetCard(ctx, cardId).Execute()

Get a project card



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
    cardId := int32(56) // int32 | The unique identifier of the card.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsGetCard(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsGetCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsGetCard`: ProjectCard
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsGetCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **int32** | The unique identifier of the card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectCard**](ProjectCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGetColumn

> ProjectColumn ProjectsGetColumn(ctx, columnId).Execute()

Get a project column



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
    columnId := int32(56) // int32 | The unique identifier of the column.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsGetColumn(context.Background(), columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsGetColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsGetColumn`: ProjectColumn
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsGetColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** | The unique identifier of the column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectColumn**](ProjectColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGetPermissionForUser

> ProjectCollaboratorPermission ProjectsGetPermissionForUser(ctx, projectId, username).Execute()

Get project permission for a user



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsGetPermissionForUser(context.Background(), projectId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsGetPermissionForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsGetPermissionForUser`: ProjectCollaboratorPermission
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsGetPermissionForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetPermissionForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectCollaboratorPermission**](ProjectCollaboratorPermission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListCards

> []ProjectCard ProjectsListCards(ctx, columnId).ArchivedState(archivedState).PerPage(perPage).Page(page).Execute()

List project cards



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
    columnId := int32(56) // int32 | The unique identifier of the column.
    archivedState := "archivedState_example" // string | Filters the project cards that are returned by the card's state. (optional) (default to "not_archived")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListCards(context.Background(), columnId).ArchivedState(archivedState).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListCards`: []ProjectCard
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** | The unique identifier of the column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archivedState** | **string** | Filters the project cards that are returned by the card&#39;s state. | [default to &quot;not_archived&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]ProjectCard**](ProjectCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListCollaborators

> []SimpleUser ProjectsListCollaborators(ctx, projectId).Affiliation(affiliation).PerPage(perPage).Page(page).Execute()

List project collaborators



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    affiliation := "affiliation_example" // string | Filters the collaborators by their affiliation. `outside` means outside collaborators of a project that are not a member of the project's organization. `direct` means collaborators with permissions to a project, regardless of organization membership status. `all` means all collaborators the authenticated user can see. (optional) (default to "all")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListCollaborators(context.Background(), projectId).Affiliation(affiliation).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListCollaborators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListCollaborators`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListCollaborators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListCollaboratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **affiliation** | **string** | Filters the collaborators by their affiliation. &#x60;outside&#x60; means outside collaborators of a project that are not a member of the project&#39;s organization. &#x60;direct&#x60; means collaborators with permissions to a project, regardless of organization membership status. &#x60;all&#x60; means all collaborators the authenticated user can see. | [default to &quot;all&quot;]
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


## ProjectsListColumns

> []ProjectColumn ProjectsListColumns(ctx, projectId).PerPage(perPage).Page(page).Execute()

List project columns



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListColumns(context.Background(), projectId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListColumns`: []ProjectColumn
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListColumnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]ProjectColumn**](ProjectColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListForOrg

> []Project ProjectsListForOrg(ctx, org).State(state).PerPage(perPage).Page(page).Execute()

List organization projects



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
    state := "state_example" // string | Indicates the state of the projects to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListForOrg(context.Background(), org).State(state).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListForOrg`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Indicates the state of the projects to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListForRepo

> []Project ProjectsListForRepo(ctx, owner, repo).State(state).PerPage(perPage).Page(page).Execute()

List repository projects



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
    state := "state_example" // string | Indicates the state of the projects to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListForRepo(context.Background(), owner, repo).State(state).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListForRepo`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | Indicates the state of the projects to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListForUser

> []Project ProjectsListForUser(ctx, username).State(state).PerPage(perPage).Page(page).Execute()

List user projects



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
    state := "state_example" // string | Indicates the state of the projects to return. Can be either `open`, `closed`, or `all`. (optional) (default to "open")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsListForUser(context.Background(), username).State(state).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsListForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsListForUser`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsListForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Indicates the state of the projects to return. Can be either &#x60;open&#x60;, &#x60;closed&#x60;, or &#x60;all&#x60;. | [default to &quot;open&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsMoveCard

> map[string]interface{} ProjectsMoveCard(ctx, cardId).ProjectsMoveCardRequest(projectsMoveCardRequest).Execute()

Move a project card



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
    cardId := int32(56) // int32 | The unique identifier of the card.
    projectsMoveCardRequest := *openapiclient.NewProjectsMoveCardRequest("bottom") // ProjectsMoveCardRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsMoveCard(context.Background(), cardId).ProjectsMoveCardRequest(projectsMoveCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsMoveCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsMoveCard`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsMoveCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **int32** | The unique identifier of the card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsMoveCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsMoveCardRequest** | [**ProjectsMoveCardRequest**](ProjectsMoveCardRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsMoveColumn

> map[string]interface{} ProjectsMoveColumn(ctx, columnId).ProjectsMoveColumnRequest(projectsMoveColumnRequest).Execute()

Move a project column



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
    columnId := int32(56) // int32 | The unique identifier of the column.
    projectsMoveColumnRequest := *openapiclient.NewProjectsMoveColumnRequest("last") // ProjectsMoveColumnRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsMoveColumn(context.Background(), columnId).ProjectsMoveColumnRequest(projectsMoveColumnRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsMoveColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsMoveColumn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsMoveColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** | The unique identifier of the column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsMoveColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsMoveColumnRequest** | [**ProjectsMoveColumnRequest**](ProjectsMoveColumnRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsRemoveCollaborator

> ProjectsRemoveCollaborator(ctx, projectId, username).Execute()

Remove user as a collaborator



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsRemoveCollaborator(context.Background(), projectId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsRemoveCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsRemoveCollaboratorRequest struct via the builder pattern


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


## ProjectsUpdate

> Project ProjectsUpdate(ctx, projectId).ProjectsUpdateRequest(projectsUpdateRequest).Execute()

Update a project



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
    projectId := int32(56) // int32 | The unique identifier of the project.
    projectsUpdateRequest := *openapiclient.NewProjectsUpdateRequest() // ProjectsUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsUpdate(context.Background(), projectId).ProjectsUpdateRequest(projectsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsUpdate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUpdateRequest** | [**ProjectsUpdateRequest**](ProjectsUpdateRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdateCard

> ProjectCard ProjectsUpdateCard(ctx, cardId).ProjectsUpdateCardRequest(projectsUpdateCardRequest).Execute()

Update an existing project card



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
    cardId := int32(56) // int32 | The unique identifier of the card.
    projectsUpdateCardRequest := *openapiclient.NewProjectsUpdateCardRequest() // ProjectsUpdateCardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsUpdateCard(context.Background(), cardId).ProjectsUpdateCardRequest(projectsUpdateCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsUpdateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsUpdateCard`: ProjectCard
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsUpdateCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **int32** | The unique identifier of the card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUpdateCardRequest** | [**ProjectsUpdateCardRequest**](ProjectsUpdateCardRequest.md) |  | 

### Return type

[**ProjectCard**](ProjectCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdateColumn

> ProjectColumn ProjectsUpdateColumn(ctx, columnId).ProjectsUpdateColumnRequest(projectsUpdateColumnRequest).Execute()

Update an existing project column



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
    columnId := int32(56) // int32 | The unique identifier of the column.
    projectsUpdateColumnRequest := *openapiclient.NewProjectsUpdateColumnRequest("Remaining tasks") // ProjectsUpdateColumnRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsUpdateColumn(context.Background(), columnId).ProjectsUpdateColumnRequest(projectsUpdateColumnRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsUpdateColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsUpdateColumn`: ProjectColumn
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsUpdateColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **int32** | The unique identifier of the column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUpdateColumnRequest** | [**ProjectsUpdateColumnRequest**](ProjectsUpdateColumnRequest.md) |  | 

### Return type

[**ProjectColumn**](ProjectColumn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

