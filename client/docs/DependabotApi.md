# \DependabotApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DependabotAddSelectedRepoToOrgSecret**](DependabotApi.md#DependabotAddSelectedRepoToOrgSecret) | **Put** /orgs/{org}/dependabot/secrets/{secret_name}/repositories/{repository_id} | Add selected repository to an organization secret
[**DependabotCreateOrUpdateOrgSecret**](DependabotApi.md#DependabotCreateOrUpdateOrgSecret) | **Put** /orgs/{org}/dependabot/secrets/{secret_name} | Create or update an organization secret
[**DependabotCreateOrUpdateRepoSecret**](DependabotApi.md#DependabotCreateOrUpdateRepoSecret) | **Put** /repos/{owner}/{repo}/dependabot/secrets/{secret_name} | Create or update a repository secret
[**DependabotDeleteOrgSecret**](DependabotApi.md#DependabotDeleteOrgSecret) | **Delete** /orgs/{org}/dependabot/secrets/{secret_name} | Delete an organization secret
[**DependabotDeleteRepoSecret**](DependabotApi.md#DependabotDeleteRepoSecret) | **Delete** /repos/{owner}/{repo}/dependabot/secrets/{secret_name} | Delete a repository secret
[**DependabotGetOrgPublicKey**](DependabotApi.md#DependabotGetOrgPublicKey) | **Get** /orgs/{org}/dependabot/secrets/public-key | Get an organization public key
[**DependabotGetOrgSecret**](DependabotApi.md#DependabotGetOrgSecret) | **Get** /orgs/{org}/dependabot/secrets/{secret_name} | Get an organization secret
[**DependabotGetRepoPublicKey**](DependabotApi.md#DependabotGetRepoPublicKey) | **Get** /repos/{owner}/{repo}/dependabot/secrets/public-key | Get a repository public key
[**DependabotGetRepoSecret**](DependabotApi.md#DependabotGetRepoSecret) | **Get** /repos/{owner}/{repo}/dependabot/secrets/{secret_name} | Get a repository secret
[**DependabotListOrgSecrets**](DependabotApi.md#DependabotListOrgSecrets) | **Get** /orgs/{org}/dependabot/secrets | List organization secrets
[**DependabotListRepoSecrets**](DependabotApi.md#DependabotListRepoSecrets) | **Get** /repos/{owner}/{repo}/dependabot/secrets | List repository secrets
[**DependabotListSelectedReposForOrgSecret**](DependabotApi.md#DependabotListSelectedReposForOrgSecret) | **Get** /orgs/{org}/dependabot/secrets/{secret_name}/repositories | List selected repositories for an organization secret
[**DependabotRemoveSelectedRepoFromOrgSecret**](DependabotApi.md#DependabotRemoveSelectedRepoFromOrgSecret) | **Delete** /orgs/{org}/dependabot/secrets/{secret_name}/repositories/{repository_id} | Remove selected repository from an organization secret
[**DependabotSetSelectedReposForOrgSecret**](DependabotApi.md#DependabotSetSelectedReposForOrgSecret) | **Put** /orgs/{org}/dependabot/secrets/{secret_name}/repositories | Set selected repositories for an organization secret



## DependabotAddSelectedRepoToOrgSecret

> DependabotAddSelectedRepoToOrgSecret(ctx, org, secretName, repositoryId).Execute()

Add selected repository to an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.
    repositoryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotAddSelectedRepoToOrgSecret(context.Background(), org, secretName, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotAddSelectedRepoToOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 
**repositoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotAddSelectedRepoToOrgSecretRequest struct via the builder pattern


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


## DependabotCreateOrUpdateOrgSecret

> map[string]interface{} DependabotCreateOrUpdateOrgSecret(ctx, org, secretName).DependabotCreateOrUpdateOrgSecretRequest(dependabotCreateOrUpdateOrgSecretRequest).Execute()

Create or update an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.
    dependabotCreateOrUpdateOrgSecretRequest := *openapiclient.NewDependabotCreateOrUpdateOrgSecretRequest("Visibility_example") // DependabotCreateOrUpdateOrgSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotCreateOrUpdateOrgSecret(context.Background(), org, secretName).DependabotCreateOrUpdateOrgSecretRequest(dependabotCreateOrUpdateOrgSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotCreateOrUpdateOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotCreateOrUpdateOrgSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotCreateOrUpdateOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotCreateOrUpdateOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dependabotCreateOrUpdateOrgSecretRequest** | [**DependabotCreateOrUpdateOrgSecretRequest**](DependabotCreateOrUpdateOrgSecretRequest.md) |  | 

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


## DependabotCreateOrUpdateRepoSecret

> map[string]interface{} DependabotCreateOrUpdateRepoSecret(ctx, owner, repo, secretName).DependabotCreateOrUpdateRepoSecretRequest(dependabotCreateOrUpdateRepoSecretRequest).Execute()

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
    dependabotCreateOrUpdateRepoSecretRequest := *openapiclient.NewDependabotCreateOrUpdateRepoSecretRequest() // DependabotCreateOrUpdateRepoSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotCreateOrUpdateRepoSecret(context.Background(), owner, repo, secretName).DependabotCreateOrUpdateRepoSecretRequest(dependabotCreateOrUpdateRepoSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotCreateOrUpdateRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotCreateOrUpdateRepoSecret`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotCreateOrUpdateRepoSecret`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDependabotCreateOrUpdateRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dependabotCreateOrUpdateRepoSecretRequest** | [**DependabotCreateOrUpdateRepoSecretRequest**](DependabotCreateOrUpdateRepoSecretRequest.md) |  | 

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


## DependabotDeleteOrgSecret

> DependabotDeleteOrgSecret(ctx, org, secretName).Execute()

Delete an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotDeleteOrgSecret(context.Background(), org, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotDeleteOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotDeleteOrgSecretRequest struct via the builder pattern


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


## DependabotDeleteRepoSecret

> DependabotDeleteRepoSecret(ctx, owner, repo, secretName).Execute()

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
    resp, r, err := apiClient.DependabotApi.DependabotDeleteRepoSecret(context.Background(), owner, repo, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotDeleteRepoSecret``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDependabotDeleteRepoSecretRequest struct via the builder pattern


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


## DependabotGetOrgPublicKey

> DependabotPublicKey DependabotGetOrgPublicKey(ctx, org).Execute()

Get an organization public key



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotGetOrgPublicKey(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotGetOrgPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotGetOrgPublicKey`: DependabotPublicKey
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotGetOrgPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotGetOrgPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DependabotPublicKey**](DependabotPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependabotGetOrgSecret

> OrganizationDependabotSecret DependabotGetOrgSecret(ctx, org, secretName).Execute()

Get an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotGetOrgSecret(context.Background(), org, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotGetOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotGetOrgSecret`: OrganizationDependabotSecret
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotGetOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotGetOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDependabotSecret**](OrganizationDependabotSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependabotGetRepoPublicKey

> DependabotPublicKey DependabotGetRepoPublicKey(ctx, owner, repo).Execute()

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
    resp, r, err := apiClient.DependabotApi.DependabotGetRepoPublicKey(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotGetRepoPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotGetRepoPublicKey`: DependabotPublicKey
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotGetRepoPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotGetRepoPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DependabotPublicKey**](DependabotPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependabotGetRepoSecret

> DependabotSecret DependabotGetRepoSecret(ctx, owner, repo, secretName).Execute()

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
    resp, r, err := apiClient.DependabotApi.DependabotGetRepoSecret(context.Background(), owner, repo, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotGetRepoSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotGetRepoSecret`: DependabotSecret
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotGetRepoSecret`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDependabotGetRepoSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DependabotSecret**](DependabotSecret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependabotListOrgSecrets

> DependabotListOrgSecrets200Response DependabotListOrgSecrets(ctx, org).PerPage(perPage).Page(page).Execute()

List organization secrets



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
    resp, r, err := apiClient.DependabotApi.DependabotListOrgSecrets(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotListOrgSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotListOrgSecrets`: DependabotListOrgSecrets200Response
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotListOrgSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotListOrgSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**DependabotListOrgSecrets200Response**](DependabotListOrgSecrets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependabotListRepoSecrets

> DependabotListRepoSecrets200Response DependabotListRepoSecrets(ctx, owner, repo).PerPage(perPage).Page(page).Execute()

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
    resp, r, err := apiClient.DependabotApi.DependabotListRepoSecrets(context.Background(), owner, repo).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotListRepoSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotListRepoSecrets`: DependabotListRepoSecrets200Response
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotListRepoSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotListRepoSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**DependabotListRepoSecrets200Response**](DependabotListRepoSecrets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DependabotListSelectedReposForOrgSecret

> ActionsListSelectedReposForOrgSecret200Response DependabotListSelectedReposForOrgSecret(ctx, org, secretName).Page(page).PerPage(perPage).Execute()

List selected repositories for an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotListSelectedReposForOrgSecret(context.Background(), org, secretName).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotListSelectedReposForOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DependabotListSelectedReposForOrgSecret`: ActionsListSelectedReposForOrgSecret200Response
    fmt.Fprintf(os.Stdout, "Response from `DependabotApi.DependabotListSelectedReposForOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotListSelectedReposForOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

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


## DependabotRemoveSelectedRepoFromOrgSecret

> DependabotRemoveSelectedRepoFromOrgSecret(ctx, org, secretName, repositoryId).Execute()

Remove selected repository from an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.
    repositoryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotRemoveSelectedRepoFromOrgSecret(context.Background(), org, secretName, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotRemoveSelectedRepoFromOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 
**repositoryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotRemoveSelectedRepoFromOrgSecretRequest struct via the builder pattern


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


## DependabotSetSelectedReposForOrgSecret

> DependabotSetSelectedReposForOrgSecret(ctx, org, secretName).DependabotSetSelectedReposForOrgSecretRequest(dependabotSetSelectedReposForOrgSecretRequest).Execute()

Set selected repositories for an organization secret



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
    secretName := "secretName_example" // string | The name of the secret.
    dependabotSetSelectedReposForOrgSecretRequest := *openapiclient.NewDependabotSetSelectedReposForOrgSecretRequest([]int32{int32(123)}) // DependabotSetSelectedReposForOrgSecretRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DependabotApi.DependabotSetSelectedReposForOrgSecret(context.Background(), org, secretName).DependabotSetSelectedReposForOrgSecretRequest(dependabotSetSelectedReposForOrgSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependabotApi.DependabotSetSelectedReposForOrgSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**secretName** | **string** | The name of the secret. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDependabotSetSelectedReposForOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dependabotSetSelectedReposForOrgSecretRequest** | [**DependabotSetSelectedReposForOrgSecretRequest**](DependabotSetSelectedReposForOrgSecretRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

