# \CodespacesApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodespacesAddRepositoryForSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesAddRepositoryForSecretForAuthenticatedUser) | **Put** /user/codespaces/secrets/{secret_name}/repositories/{repository_id} | Add a selected repository to a user secret
[**CodespacesCodespaceMachinesForAuthenticatedUser**](CodespacesApi.md#CodespacesCodespaceMachinesForAuthenticatedUser) | **Get** /user/codespaces/{codespace_name}/machines | List machine types for a codespace
[**CodespacesCreateForAuthenticatedUser**](CodespacesApi.md#CodespacesCreateForAuthenticatedUser) | **Post** /user/codespaces | Create a codespace for the authenticated user
[**CodespacesCreateOrUpdateRepoSecret**](CodespacesApi.md#CodespacesCreateOrUpdateRepoSecret) | **Put** /repos/{owner}/{repo}/codespaces/secrets/{secret_name} | Create or update a repository secret
[**CodespacesCreateOrUpdateSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesCreateOrUpdateSecretForAuthenticatedUser) | **Put** /user/codespaces/secrets/{secret_name} | Create or update a secret for the authenticated user
[**CodespacesCreateWithPrForAuthenticatedUser**](CodespacesApi.md#CodespacesCreateWithPrForAuthenticatedUser) | **Post** /repos/{owner}/{repo}/pulls/{pull_number}/codespaces | Create a codespace from a pull request
[**CodespacesCreateWithRepoForAuthenticatedUser**](CodespacesApi.md#CodespacesCreateWithRepoForAuthenticatedUser) | **Post** /repos/{owner}/{repo}/codespaces | Create a codespace in a repository
[**CodespacesDeleteForAuthenticatedUser**](CodespacesApi.md#CodespacesDeleteForAuthenticatedUser) | **Delete** /user/codespaces/{codespace_name} | Delete a codespace for the authenticated user
[**CodespacesDeleteFromOrganization**](CodespacesApi.md#CodespacesDeleteFromOrganization) | **Delete** /orgs/{org}/members/{username}/codespaces/{codespace_name} | Delete a codespace from the organization
[**CodespacesDeleteRepoSecret**](CodespacesApi.md#CodespacesDeleteRepoSecret) | **Delete** /repos/{owner}/{repo}/codespaces/secrets/{secret_name} | Delete a repository secret
[**CodespacesDeleteSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesDeleteSecretForAuthenticatedUser) | **Delete** /user/codespaces/secrets/{secret_name} | Delete a secret for the authenticated user
[**CodespacesExportForAuthenticatedUser**](CodespacesApi.md#CodespacesExportForAuthenticatedUser) | **Post** /user/codespaces/{codespace_name}/exports | Export a codespace for the authenticated user
[**CodespacesGetExportDetailsForAuthenticatedUser**](CodespacesApi.md#CodespacesGetExportDetailsForAuthenticatedUser) | **Get** /user/codespaces/{codespace_name}/exports/{export_id} | Get details about a codespace export
[**CodespacesGetForAuthenticatedUser**](CodespacesApi.md#CodespacesGetForAuthenticatedUser) | **Get** /user/codespaces/{codespace_name} | Get a codespace for the authenticated user
[**CodespacesGetPublicKeyForAuthenticatedUser**](CodespacesApi.md#CodespacesGetPublicKeyForAuthenticatedUser) | **Get** /user/codespaces/secrets/public-key | Get public key for the authenticated user
[**CodespacesGetRepoPublicKey**](CodespacesApi.md#CodespacesGetRepoPublicKey) | **Get** /repos/{owner}/{repo}/codespaces/secrets/public-key | Get a repository public key
[**CodespacesGetRepoSecret**](CodespacesApi.md#CodespacesGetRepoSecret) | **Get** /repos/{owner}/{repo}/codespaces/secrets/{secret_name} | Get a repository secret
[**CodespacesGetSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesGetSecretForAuthenticatedUser) | **Get** /user/codespaces/secrets/{secret_name} | Get a secret for the authenticated user
[**CodespacesListDevcontainersInRepositoryForAuthenticatedUser**](CodespacesApi.md#CodespacesListDevcontainersInRepositoryForAuthenticatedUser) | **Get** /repos/{owner}/{repo}/codespaces/devcontainers | List devcontainer configurations in a repository for the authenticated user
[**CodespacesListForAuthenticatedUser**](CodespacesApi.md#CodespacesListForAuthenticatedUser) | **Get** /user/codespaces | List codespaces for the authenticated user
[**CodespacesListInOrganization**](CodespacesApi.md#CodespacesListInOrganization) | **Get** /orgs/{org}/codespaces | List codespaces for the organization
[**CodespacesListInRepositoryForAuthenticatedUser**](CodespacesApi.md#CodespacesListInRepositoryForAuthenticatedUser) | **Get** /repos/{owner}/{repo}/codespaces | List codespaces in a repository for the authenticated user
[**CodespacesListRepoSecrets**](CodespacesApi.md#CodespacesListRepoSecrets) | **Get** /repos/{owner}/{repo}/codespaces/secrets | List repository secrets
[**CodespacesListRepositoriesForSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesListRepositoriesForSecretForAuthenticatedUser) | **Get** /user/codespaces/secrets/{secret_name}/repositories | List selected repositories for a user secret
[**CodespacesListSecretsForAuthenticatedUser**](CodespacesApi.md#CodespacesListSecretsForAuthenticatedUser) | **Get** /user/codespaces/secrets | List secrets for the authenticated user
[**CodespacesPreFlightWithRepoForAuthenticatedUser**](CodespacesApi.md#CodespacesPreFlightWithRepoForAuthenticatedUser) | **Get** /repos/{owner}/{repo}/codespaces/new | Get default attributes for a codespace
[**CodespacesRemoveRepositoryForSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesRemoveRepositoryForSecretForAuthenticatedUser) | **Delete** /user/codespaces/secrets/{secret_name}/repositories/{repository_id} | Remove a selected repository from a user secret
[**CodespacesRepoMachinesForAuthenticatedUser**](CodespacesApi.md#CodespacesRepoMachinesForAuthenticatedUser) | **Get** /repos/{owner}/{repo}/codespaces/machines | List available machine types for a repository
[**CodespacesSetRepositoriesForSecretForAuthenticatedUser**](CodespacesApi.md#CodespacesSetRepositoriesForSecretForAuthenticatedUser) | **Put** /user/codespaces/secrets/{secret_name}/repositories | Set selected repositories for a user secret
[**CodespacesStartForAuthenticatedUser**](CodespacesApi.md#CodespacesStartForAuthenticatedUser) | **Post** /user/codespaces/{codespace_name}/start | Start a codespace for the authenticated user
[**CodespacesStopForAuthenticatedUser**](CodespacesApi.md#CodespacesStopForAuthenticatedUser) | **Post** /user/codespaces/{codespace_name}/stop | Stop a codespace for the authenticated user
[**CodespacesStopInOrganization**](CodespacesApi.md#CodespacesStopInOrganization) | **Post** /orgs/{org}/members/{username}/codespaces/{codespace_name}/stop | Stop a codespace for an organization user
[**CodespacesUpdateForAuthenticatedUser**](CodespacesApi.md#CodespacesUpdateForAuthenticatedUser) | **Patch** /user/codespaces/{codespace_name} | Update a codespace for the authenticated user



## CodespacesAddRepositoryForSecretForAuthenticatedUser

> CodespacesAddRepositoryForSecretForAuthenticatedUser(ctx, secretName, repositoryId).Execute()

Add a selected repository to a user secret



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
    secretName := "secretName_example" // string | The name of the secret.
    repositoryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesAddRepositoryForSecretForAuthenticatedUser(context.Background(), secretName, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesAddRepositoryForSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 
**repositoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesAddRepositoryForSecretForAuthenticatedUserRequest struct via the builder pattern


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


## CodespacesCodespaceMachinesForAuthenticatedUser

> CodespacesRepoMachinesForAuthenticatedUser200Response CodespacesCodespaceMachinesForAuthenticatedUser(ctx, codespaceName).Execute()

List machine types for a codespace



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesCodespaceMachinesForAuthenticatedUser(context.Background(), codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesCodespaceMachinesForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesCodespaceMachinesForAuthenticatedUser`: CodespacesRepoMachinesForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesCodespaceMachinesForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesCodespaceMachinesForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CodespacesRepoMachinesForAuthenticatedUser200Response**](CodespacesRepoMachinesForAuthenticatedUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesCreateForAuthenticatedUser

> Codespace CodespacesCreateForAuthenticatedUser(ctx).CodespacesCreateForAuthenticatedUserRequest(codespacesCreateForAuthenticatedUserRequest).Execute()

Create a codespace for the authenticated user



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
    codespacesCreateForAuthenticatedUserRequest := openapiclient.codespaces_create_for_authenticated_user_request{CodespacesCreateForAuthenticatedUserRequestOneOf: openapiclient.NewCodespacesCreateForAuthenticatedUserRequestOneOf(int32(123))} // CodespacesCreateForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesCreateForAuthenticatedUser(context.Background()).CodespacesCreateForAuthenticatedUserRequest(codespacesCreateForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesCreateForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesCreateForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesCreateForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesCreateForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codespacesCreateForAuthenticatedUserRequest** | [**CodespacesCreateForAuthenticatedUserRequest**](CodespacesCreateForAuthenticatedUserRequest.md) |  | 

### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesCreateOrUpdateRepoSecret

> map[string]interface{} CodespacesCreateOrUpdateRepoSecret(ctx, owner, repo, secretName).CodespacesCreateOrUpdateRepoSecretRequest(codespacesCreateOrUpdateRepoSecretRequest).Execute()

Create or update a repository secret



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
    secretName := "secretName_example" // string | The name of the secret.
    codespacesCreateOrUpdateRepoSecretRequest := *openapiclient.NewCodespacesCreateOrUpdateRepoSecretRequest() // CodespacesCreateOrUpdateRepoSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesCreateOrUpdateRepoSecret(context.Background(), owner, repo, secretName).CodespacesCreateOrUpdateRepoSecretRequest(codespacesCreateOrUpdateRepoSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesCreateOrUpdateRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesCreateOrUpdateRepoSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesCreateOrUpdateRepoSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesCreateOrUpdateRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **codespacesCreateOrUpdateRepoSecretRequest** | [**CodespacesCreateOrUpdateRepoSecretRequest**](CodespacesCreateOrUpdateRepoSecretRequest.md) |  | 

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


## CodespacesCreateOrUpdateSecretForAuthenticatedUser

> map[string]interface{} CodespacesCreateOrUpdateSecretForAuthenticatedUser(ctx, secretName).CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest(codespacesCreateOrUpdateSecretForAuthenticatedUserRequest).Execute()

Create or update a secret for the authenticated user



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
    secretName := "secretName_example" // string | The name of the secret.
    codespacesCreateOrUpdateSecretForAuthenticatedUserRequest := *openapiclient.NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest("KeyId_example") // CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesCreateOrUpdateSecretForAuthenticatedUser(context.Background(), secretName).CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest(codespacesCreateOrUpdateSecretForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesCreateOrUpdateSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesCreateOrUpdateSecretForAuthenticatedUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesCreateOrUpdateSecretForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codespacesCreateOrUpdateSecretForAuthenticatedUserRequest** | [**CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest**](CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest.md) |  | 

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


## CodespacesCreateWithPrForAuthenticatedUser

> Codespace CodespacesCreateWithPrForAuthenticatedUser(ctx, owner, repo, pullNumber).CodespacesCreateWithPrForAuthenticatedUserRequest(codespacesCreateWithPrForAuthenticatedUserRequest).Execute()

Create a codespace from a pull request



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
    codespacesCreateWithPrForAuthenticatedUserRequest := *openapiclient.NewCodespacesCreateWithPrForAuthenticatedUserRequest() // CodespacesCreateWithPrForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesCreateWithPrForAuthenticatedUser(context.Background(), owner, repo, pullNumber).CodespacesCreateWithPrForAuthenticatedUserRequest(codespacesCreateWithPrForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesCreateWithPrForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesCreateWithPrForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesCreateWithPrForAuthenticatedUser`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCodespacesCreateWithPrForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **codespacesCreateWithPrForAuthenticatedUserRequest** | [**CodespacesCreateWithPrForAuthenticatedUserRequest**](CodespacesCreateWithPrForAuthenticatedUserRequest.md) |  | 

### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesCreateWithRepoForAuthenticatedUser

> Codespace CodespacesCreateWithRepoForAuthenticatedUser(ctx, owner, repo).CodespacesCreateWithRepoForAuthenticatedUserRequest(codespacesCreateWithRepoForAuthenticatedUserRequest).Execute()

Create a codespace in a repository



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
    codespacesCreateWithRepoForAuthenticatedUserRequest := *openapiclient.NewCodespacesCreateWithRepoForAuthenticatedUserRequest() // CodespacesCreateWithRepoForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesCreateWithRepoForAuthenticatedUser(context.Background(), owner, repo).CodespacesCreateWithRepoForAuthenticatedUserRequest(codespacesCreateWithRepoForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesCreateWithRepoForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesCreateWithRepoForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesCreateWithRepoForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesCreateWithRepoForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **codespacesCreateWithRepoForAuthenticatedUserRequest** | [**CodespacesCreateWithRepoForAuthenticatedUserRequest**](CodespacesCreateWithRepoForAuthenticatedUserRequest.md) |  | 

### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesDeleteForAuthenticatedUser

> map[string]interface{} CodespacesDeleteForAuthenticatedUser(ctx, codespaceName).Execute()

Delete a codespace for the authenticated user



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesDeleteForAuthenticatedUser(context.Background(), codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesDeleteForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesDeleteForAuthenticatedUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesDeleteForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesDeleteForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesDeleteFromOrganization

> map[string]interface{} CodespacesDeleteFromOrganization(ctx, org, username, codespaceName).Execute()

Delete a codespace from the organization



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
    username := "username_example" // string | The handle for the GitHub user account.
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesDeleteFromOrganization(context.Background(), org, username, codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesDeleteFromOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesDeleteFromOrganization`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesDeleteFromOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesDeleteFromOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesDeleteRepoSecret

> CodespacesDeleteRepoSecret(ctx, owner, repo, secretName).Execute()

Delete a repository secret



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesDeleteRepoSecret(context.Background(), owner, repo, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesDeleteRepoSecret``: %v\n", err)
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
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesDeleteRepoSecretRequest struct via the builder pattern


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


## CodespacesDeleteSecretForAuthenticatedUser

> CodespacesDeleteSecretForAuthenticatedUser(ctx, secretName).Execute()

Delete a secret for the authenticated user



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesDeleteSecretForAuthenticatedUser(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesDeleteSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesDeleteSecretForAuthenticatedUserRequest struct via the builder pattern


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


## CodespacesExportForAuthenticatedUser

> CodespaceExportDetails CodespacesExportForAuthenticatedUser(ctx, codespaceName).Execute()

Export a codespace for the authenticated user



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesExportForAuthenticatedUser(context.Background(), codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesExportForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesExportForAuthenticatedUser`: CodespaceExportDetails
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesExportForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesExportForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CodespaceExportDetails**](CodespaceExportDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesGetExportDetailsForAuthenticatedUser

> CodespaceExportDetails CodespacesGetExportDetailsForAuthenticatedUser(ctx, codespaceName, exportId).Execute()

Get details about a codespace export



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.
    exportId := "exportId_example" // string | The ID of the export operation, or `latest`. Currently only `latest` is currently supported.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesGetExportDetailsForAuthenticatedUser(context.Background(), codespaceName, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesGetExportDetailsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesGetExportDetailsForAuthenticatedUser`: CodespaceExportDetails
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesGetExportDetailsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 
**exportId** | **string** | The ID of the export operation, or &#x60;latest&#x60;. Currently only &#x60;latest&#x60; is currently supported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesGetExportDetailsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CodespaceExportDetails**](CodespaceExportDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesGetForAuthenticatedUser

> Codespace CodespacesGetForAuthenticatedUser(ctx, codespaceName).Execute()

Get a codespace for the authenticated user



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesGetForAuthenticatedUser(context.Background(), codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesGetForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesGetForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesGetForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesGetForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesGetPublicKeyForAuthenticatedUser

> CodespacesUserPublicKey CodespacesGetPublicKeyForAuthenticatedUser(ctx).Execute()

Get public key for the authenticated user



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
    resp, r, err := apiClient.CodespacesApi.CodespacesGetPublicKeyForAuthenticatedUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesGetPublicKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesGetPublicKeyForAuthenticatedUser`: CodespacesUserPublicKey
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesGetPublicKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesGetPublicKeyForAuthenticatedUserRequest struct via the builder pattern


### Return type

[**CodespacesUserPublicKey**](CodespacesUserPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesGetRepoPublicKey

> CodespacesPublicKey CodespacesGetRepoPublicKey(ctx, owner, repo).Execute()

Get a repository public key



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
    resp, r, err := apiClient.CodespacesApi.CodespacesGetRepoPublicKey(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesGetRepoPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesGetRepoPublicKey`: CodespacesPublicKey
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesGetRepoPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesGetRepoPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CodespacesPublicKey**](CodespacesPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesGetRepoSecret

> RepoCodespacesSecret CodespacesGetRepoSecret(ctx, owner, repo, secretName).Execute()

Get a repository secret



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesGetRepoSecret(context.Background(), owner, repo, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesGetRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesGetRepoSecret`: RepoCodespacesSecret
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesGetRepoSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesGetRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RepoCodespacesSecret**](RepoCodespacesSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesGetSecretForAuthenticatedUser

> CodespacesSecret CodespacesGetSecretForAuthenticatedUser(ctx, secretName).Execute()

Get a secret for the authenticated user



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesGetSecretForAuthenticatedUser(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesGetSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesGetSecretForAuthenticatedUser`: CodespacesSecret
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesGetSecretForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesGetSecretForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CodespacesSecret**](CodespacesSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListDevcontainersInRepositoryForAuthenticatedUser

> CodespacesListDevcontainersInRepositoryForAuthenticatedUser200Response CodespacesListDevcontainersInRepositoryForAuthenticatedUser(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List devcontainer configurations in a repository for the authenticated user



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
    resp, r, err := apiClient.CodespacesApi.CodespacesListDevcontainersInRepositoryForAuthenticatedUser(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListDevcontainersInRepositoryForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListDevcontainersInRepositoryForAuthenticatedUser`: CodespacesListDevcontainersInRepositoryForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListDevcontainersInRepositoryForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListDevcontainersInRepositoryForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**CodespacesListDevcontainersInRepositoryForAuthenticatedUser200Response**](CodespacesListDevcontainersInRepositoryForAuthenticatedUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListForAuthenticatedUser

> CodespacesListInOrganization200Response CodespacesListForAuthenticatedUser(ctx).PerPage(perPage).Page(page).RepositoryId(repositoryId).Execute()

List codespaces for the authenticated user



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
    repositoryId := int32(56) // int32 | ID of the Repository to filter on (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesListForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).RepositoryId(repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListForAuthenticatedUser`: CodespacesListInOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **repositoryId** | **int32** | ID of the Repository to filter on | 

### Return type

[**CodespacesListInOrganization200Response**](CodespacesListInOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListInOrganization

> CodespacesListInOrganization200Response CodespacesListInOrganization(ctx, org).PerPage(perPage).Page(page).Execute()

List codespaces for the organization



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
    resp, r, err := apiClient.CodespacesApi.CodespacesListInOrganization(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListInOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListInOrganization`: CodespacesListInOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListInOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListInOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**CodespacesListInOrganization200Response**](CodespacesListInOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListInRepositoryForAuthenticatedUser

> CodespacesListInOrganization200Response CodespacesListInRepositoryForAuthenticatedUser(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List codespaces in a repository for the authenticated user



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
    resp, r, err := apiClient.CodespacesApi.CodespacesListInRepositoryForAuthenticatedUser(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListInRepositoryForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListInRepositoryForAuthenticatedUser`: CodespacesListInOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListInRepositoryForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListInRepositoryForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**CodespacesListInOrganization200Response**](CodespacesListInOrganization200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListRepoSecrets

> CodespacesListRepoSecrets200Response CodespacesListRepoSecrets(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

List repository secrets



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
    resp, r, err := apiClient.CodespacesApi.CodespacesListRepoSecrets(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListRepoSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListRepoSecrets`: CodespacesListRepoSecrets200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListRepoSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListRepoSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**CodespacesListRepoSecrets200Response**](CodespacesListRepoSecrets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListRepositoriesForSecretForAuthenticatedUser

> ActionsListSelectedReposForOrgSecret200Response CodespacesListRepositoriesForSecretForAuthenticatedUser(ctx, secretName).Execute()

List selected repositories for a user secret



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesListRepositoriesForSecretForAuthenticatedUser(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListRepositoriesForSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListRepositoriesForSecretForAuthenticatedUser`: ActionsListSelectedReposForOrgSecret200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListRepositoriesForSecretForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListRepositoriesForSecretForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionsListSelectedReposForOrgSecret200Response**](ActionsListSelectedReposForOrgSecret200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesListSecretsForAuthenticatedUser

> CodespacesListSecretsForAuthenticatedUser200Response CodespacesListSecretsForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List secrets for the authenticated user



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
    resp, r, err := apiClient.CodespacesApi.CodespacesListSecretsForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesListSecretsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesListSecretsForAuthenticatedUser`: CodespacesListSecretsForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesListSecretsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesListSecretsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**CodespacesListSecretsForAuthenticatedUser200Response**](CodespacesListSecretsForAuthenticatedUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesPreFlightWithRepoForAuthenticatedUser

> CodespacesPreFlightWithRepoForAuthenticatedUser200Response CodespacesPreFlightWithRepoForAuthenticatedUser(ctx, owner, repo).Ref(ref).ClientIp(clientIp).Execute()

Get default attributes for a codespace



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
    ref := "main" // string | The branch or commit to check for a default devcontainer path. If not specified, the default branch will be checked. (optional)
    clientIp := "1.2.3.4" // string | An alternative IP for default location auto-detection, such as when proxying a request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesPreFlightWithRepoForAuthenticatedUser(context.Background(), owner, repo).Ref(ref).ClientIp(clientIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesPreFlightWithRepoForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesPreFlightWithRepoForAuthenticatedUser`: CodespacesPreFlightWithRepoForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesPreFlightWithRepoForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesPreFlightWithRepoForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **string** | The branch or commit to check for a default devcontainer path. If not specified, the default branch will be checked. | 
 **clientIp** | **string** | An alternative IP for default location auto-detection, such as when proxying a request. | 

### Return type

[**CodespacesPreFlightWithRepoForAuthenticatedUser200Response**](CodespacesPreFlightWithRepoForAuthenticatedUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesRemoveRepositoryForSecretForAuthenticatedUser

> CodespacesRemoveRepositoryForSecretForAuthenticatedUser(ctx, secretName, repositoryId).Execute()

Remove a selected repository from a user secret



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
    secretName := "secretName_example" // string | The name of the secret.
    repositoryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesRemoveRepositoryForSecretForAuthenticatedUser(context.Background(), secretName, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesRemoveRepositoryForSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 
**repositoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesRemoveRepositoryForSecretForAuthenticatedUserRequest struct via the builder pattern


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


## CodespacesRepoMachinesForAuthenticatedUser

> CodespacesRepoMachinesForAuthenticatedUser200Response CodespacesRepoMachinesForAuthenticatedUser(ctx, owner, repo).Location(location).ClientIp(clientIp).Execute()

List available machine types for a repository



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
    location := "WestUs2" // string | The location to check for available machines. Assigned by IP if not provided. (optional)
    clientIp := "clientIp_example" // string | IP for location auto-detection when proxying a request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesRepoMachinesForAuthenticatedUser(context.Background(), owner, repo).Location(location).ClientIp(clientIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesRepoMachinesForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesRepoMachinesForAuthenticatedUser`: CodespacesRepoMachinesForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesRepoMachinesForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesRepoMachinesForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **location** | **string** | The location to check for available machines. Assigned by IP if not provided. | 
 **clientIp** | **string** | IP for location auto-detection when proxying a request | 

### Return type

[**CodespacesRepoMachinesForAuthenticatedUser200Response**](CodespacesRepoMachinesForAuthenticatedUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesSetRepositoriesForSecretForAuthenticatedUser

> CodespacesSetRepositoriesForSecretForAuthenticatedUser(ctx, secretName).CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest(codespacesSetRepositoriesForSecretForAuthenticatedUserRequest).Execute()

Set selected repositories for a user secret



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
    secretName := "secretName_example" // string | The name of the secret.
    codespacesSetRepositoriesForSecretForAuthenticatedUserRequest := *openapiclient.NewCodespacesSetRepositoriesForSecretForAuthenticatedUserRequest([]int32{int32(123)}) // CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesSetRepositoriesForSecretForAuthenticatedUser(context.Background(), secretName).CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest(codespacesSetRepositoriesForSecretForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesSetRepositoriesForSecretForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesSetRepositoriesForSecretForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codespacesSetRepositoriesForSecretForAuthenticatedUserRequest** | [**CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest**](CodespacesSetRepositoriesForSecretForAuthenticatedUserRequest.md) |  | 

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


## CodespacesStartForAuthenticatedUser

> Codespace CodespacesStartForAuthenticatedUser(ctx, codespaceName).Execute()

Start a codespace for the authenticated user



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesStartForAuthenticatedUser(context.Background(), codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesStartForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesStartForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesStartForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesStartForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesStopForAuthenticatedUser

> Codespace CodespacesStopForAuthenticatedUser(ctx, codespaceName).Execute()

Stop a codespace for the authenticated user



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesStopForAuthenticatedUser(context.Background(), codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesStopForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesStopForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesStopForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesStopForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesStopInOrganization

> Codespace CodespacesStopInOrganization(ctx, org, username, codespaceName).Execute()

Stop a codespace for an organization user



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
    username := "username_example" // string | The handle for the GitHub user account.
    codespaceName := "codespaceName_example" // string | The name of the codespace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesStopInOrganization(context.Background(), org, username, codespaceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesStopInOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesStopInOrganization`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesStopInOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesStopInOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodespacesUpdateForAuthenticatedUser

> Codespace CodespacesUpdateForAuthenticatedUser(ctx, codespaceName).CodespacesUpdateForAuthenticatedUserRequest(codespacesUpdateForAuthenticatedUserRequest).Execute()

Update a codespace for the authenticated user



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
    codespaceName := "codespaceName_example" // string | The name of the codespace.
    codespacesUpdateForAuthenticatedUserRequest := *openapiclient.NewCodespacesUpdateForAuthenticatedUserRequest() // CodespacesUpdateForAuthenticatedUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodespacesApi.CodespacesUpdateForAuthenticatedUser(context.Background(), codespaceName).CodespacesUpdateForAuthenticatedUserRequest(codespacesUpdateForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodespacesApi.CodespacesUpdateForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodespacesUpdateForAuthenticatedUser`: Codespace
    fmt.Fprintf(os.Stdout, "Response from `CodespacesApi.CodespacesUpdateForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codespaceName** | **string** | The name of the codespace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodespacesUpdateForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codespacesUpdateForAuthenticatedUserRequest** | [**CodespacesUpdateForAuthenticatedUserRequest**](CodespacesUpdateForAuthenticatedUserRequest.md) |  | 

### Return type

[**Codespace**](Codespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

