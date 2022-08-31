# \AppsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAddRepoToInstallationForAuthenticatedUser**](AppsApi.md#AppsAddRepoToInstallationForAuthenticatedUser) | **Put** /user/installations/{installation_id}/repositories/{repository_id} | Add a repository to an app installation
[**AppsCheckToken**](AppsApi.md#AppsCheckToken) | **Post** /applications/{client_id}/token | Check a token
[**AppsCreateFromManifest**](AppsApi.md#AppsCreateFromManifest) | **Post** /app-manifests/{code}/conversions | Create a GitHub App from a manifest
[**AppsCreateInstallationAccessToken**](AppsApi.md#AppsCreateInstallationAccessToken) | **Post** /app/installations/{installation_id}/access_tokens | Create an installation access token for an app
[**AppsDeleteAuthorization**](AppsApi.md#AppsDeleteAuthorization) | **Delete** /applications/{client_id}/grant | Delete an app authorization
[**AppsDeleteInstallation**](AppsApi.md#AppsDeleteInstallation) | **Delete** /app/installations/{installation_id} | Delete an installation for the authenticated app
[**AppsDeleteToken**](AppsApi.md#AppsDeleteToken) | **Delete** /applications/{client_id}/token | Delete an app token
[**AppsGetAuthenticated**](AppsApi.md#AppsGetAuthenticated) | **Get** /app | Get the authenticated app
[**AppsGetBySlug**](AppsApi.md#AppsGetBySlug) | **Get** /apps/{app_slug} | Get an app
[**AppsGetInstallation**](AppsApi.md#AppsGetInstallation) | **Get** /app/installations/{installation_id} | Get an installation for the authenticated app
[**AppsGetOrgInstallation**](AppsApi.md#AppsGetOrgInstallation) | **Get** /orgs/{org}/installation | Get an organization installation for the authenticated app
[**AppsGetRepoInstallation**](AppsApi.md#AppsGetRepoInstallation) | **Get** /repos/{owner}/{repo}/installation | Get a repository installation for the authenticated app
[**AppsGetSubscriptionPlanForAccount**](AppsApi.md#AppsGetSubscriptionPlanForAccount) | **Get** /marketplace_listing/accounts/{account_id} | Get a subscription plan for an account
[**AppsGetSubscriptionPlanForAccountStubbed**](AppsApi.md#AppsGetSubscriptionPlanForAccountStubbed) | **Get** /marketplace_listing/stubbed/accounts/{account_id} | Get a subscription plan for an account (stubbed)
[**AppsGetUserInstallation**](AppsApi.md#AppsGetUserInstallation) | **Get** /users/{username}/installation | Get a user installation for the authenticated app
[**AppsGetWebhookConfigForApp**](AppsApi.md#AppsGetWebhookConfigForApp) | **Get** /app/hook/config | Get a webhook configuration for an app
[**AppsGetWebhookDelivery**](AppsApi.md#AppsGetWebhookDelivery) | **Get** /app/hook/deliveries/{delivery_id} | Get a delivery for an app webhook
[**AppsListAccountsForPlan**](AppsApi.md#AppsListAccountsForPlan) | **Get** /marketplace_listing/plans/{plan_id}/accounts | List accounts for a plan
[**AppsListAccountsForPlanStubbed**](AppsApi.md#AppsListAccountsForPlanStubbed) | **Get** /marketplace_listing/stubbed/plans/{plan_id}/accounts | List accounts for a plan (stubbed)
[**AppsListInstallationReposForAuthenticatedUser**](AppsApi.md#AppsListInstallationReposForAuthenticatedUser) | **Get** /user/installations/{installation_id}/repositories | List repositories accessible to the user access token
[**AppsListInstallations**](AppsApi.md#AppsListInstallations) | **Get** /app/installations | List installations for the authenticated app
[**AppsListInstallationsForAuthenticatedUser**](AppsApi.md#AppsListInstallationsForAuthenticatedUser) | **Get** /user/installations | List app installations accessible to the user access token
[**AppsListPlans**](AppsApi.md#AppsListPlans) | **Get** /marketplace_listing/plans | List plans
[**AppsListPlansStubbed**](AppsApi.md#AppsListPlansStubbed) | **Get** /marketplace_listing/stubbed/plans | List plans (stubbed)
[**AppsListReposAccessibleToInstallation**](AppsApi.md#AppsListReposAccessibleToInstallation) | **Get** /installation/repositories | List repositories accessible to the app installation
[**AppsListSubscriptionsForAuthenticatedUser**](AppsApi.md#AppsListSubscriptionsForAuthenticatedUser) | **Get** /user/marketplace_purchases | List subscriptions for the authenticated user
[**AppsListSubscriptionsForAuthenticatedUserStubbed**](AppsApi.md#AppsListSubscriptionsForAuthenticatedUserStubbed) | **Get** /user/marketplace_purchases/stubbed | List subscriptions for the authenticated user (stubbed)
[**AppsListWebhookDeliveries**](AppsApi.md#AppsListWebhookDeliveries) | **Get** /app/hook/deliveries | List deliveries for an app webhook
[**AppsRedeliverWebhookDelivery**](AppsApi.md#AppsRedeliverWebhookDelivery) | **Post** /app/hook/deliveries/{delivery_id}/attempts | Redeliver a delivery for an app webhook
[**AppsRemoveRepoFromInstallationForAuthenticatedUser**](AppsApi.md#AppsRemoveRepoFromInstallationForAuthenticatedUser) | **Delete** /user/installations/{installation_id}/repositories/{repository_id} | Remove a repository from an app installation
[**AppsResetToken**](AppsApi.md#AppsResetToken) | **Patch** /applications/{client_id}/token | Reset a token
[**AppsRevokeInstallationAccessToken**](AppsApi.md#AppsRevokeInstallationAccessToken) | **Delete** /installation/token | Revoke an installation access token
[**AppsScopeToken**](AppsApi.md#AppsScopeToken) | **Post** /applications/{client_id}/token/scoped | Create a scoped access token
[**AppsSuspendInstallation**](AppsApi.md#AppsSuspendInstallation) | **Put** /app/installations/{installation_id}/suspended | Suspend an app installation
[**AppsUnsuspendInstallation**](AppsApi.md#AppsUnsuspendInstallation) | **Delete** /app/installations/{installation_id}/suspended | Unsuspend an app installation
[**AppsUpdateWebhookConfigForApp**](AppsApi.md#AppsUpdateWebhookConfigForApp) | **Patch** /app/hook/config | Update a webhook configuration for an app



## AppsAddRepoToInstallationForAuthenticatedUser

> AppsAddRepoToInstallationForAuthenticatedUser(ctx, installationId, repositoryId).Execute()

Add a repository to an app installation



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
    installationId := int32(1) // int32 | The unique identifier of the installation.
    repositoryId := int32(56) // int32 | The unique identifier of the repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsAddRepoToInstallationForAuthenticatedUser(context.Background(), installationId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsAddRepoToInstallationForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 
**repositoryId** | **int32** | The unique identifier of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAddRepoToInstallationForAuthenticatedUserRequest struct via the builder pattern


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


## AppsCheckToken

> Authorization AppsCheckToken(ctx, clientId).AppsCheckTokenRequest(appsCheckTokenRequest).Execute()

Check a token



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
    clientId := "Iv1.8a61f9b3a7aba766" // string | The client ID of the GitHub app.
    appsCheckTokenRequest := *openapiclient.NewAppsCheckTokenRequest("AccessToken_example") // AppsCheckTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsCheckToken(context.Background(), clientId).AppsCheckTokenRequest(appsCheckTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsCheckToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsCheckToken`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsCheckToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID of the GitHub app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCheckTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsCheckTokenRequest** | [**AppsCheckTokenRequest**](AppsCheckTokenRequest.md) |  | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsCreateFromManifest

> AppsCreateFromManifest201Response AppsCreateFromManifest(ctx, code).Execute()

Create a GitHub App from a manifest



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
    code := "code_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsCreateFromManifest(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsCreateFromManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsCreateFromManifest`: AppsCreateFromManifest201Response
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsCreateFromManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCreateFromManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppsCreateFromManifest201Response**](AppsCreateFromManifest201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsCreateInstallationAccessToken

> InstallationToken AppsCreateInstallationAccessToken(ctx, installationId).AppsCreateInstallationAccessTokenRequest(appsCreateInstallationAccessTokenRequest).Execute()

Create an installation access token for an app



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
    installationId := int32(1) // int32 | The unique identifier of the installation.
    appsCreateInstallationAccessTokenRequest := *openapiclient.NewAppsCreateInstallationAccessTokenRequest() // AppsCreateInstallationAccessTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsCreateInstallationAccessToken(context.Background(), installationId).AppsCreateInstallationAccessTokenRequest(appsCreateInstallationAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsCreateInstallationAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsCreateInstallationAccessToken`: InstallationToken
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsCreateInstallationAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCreateInstallationAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsCreateInstallationAccessTokenRequest** | [**AppsCreateInstallationAccessTokenRequest**](AppsCreateInstallationAccessTokenRequest.md) |  | 

### Return type

[**InstallationToken**](InstallationToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsDeleteAuthorization

> AppsDeleteAuthorization(ctx, clientId).AppsDeleteAuthorizationRequest(appsDeleteAuthorizationRequest).Execute()

Delete an app authorization



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
    clientId := "Iv1.8a61f9b3a7aba766" // string | The client ID of the GitHub app.
    appsDeleteAuthorizationRequest := *openapiclient.NewAppsDeleteAuthorizationRequest("AccessToken_example") // AppsDeleteAuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsDeleteAuthorization(context.Background(), clientId).AppsDeleteAuthorizationRequest(appsDeleteAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsDeleteAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID of the GitHub app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDeleteAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsDeleteAuthorizationRequest** | [**AppsDeleteAuthorizationRequest**](AppsDeleteAuthorizationRequest.md) |  | 

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


## AppsDeleteInstallation

> AppsDeleteInstallation(ctx, installationId).Execute()

Delete an installation for the authenticated app



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
    installationId := int32(1) // int32 | The unique identifier of the installation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsDeleteInstallation(context.Background(), installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsDeleteInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDeleteInstallationRequest struct via the builder pattern


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


## AppsDeleteToken

> AppsDeleteToken(ctx, clientId).AppsDeleteAuthorizationRequest(appsDeleteAuthorizationRequest).Execute()

Delete an app token



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
    clientId := "Iv1.8a61f9b3a7aba766" // string | The client ID of the GitHub app.
    appsDeleteAuthorizationRequest := *openapiclient.NewAppsDeleteAuthorizationRequest("AccessToken_example") // AppsDeleteAuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsDeleteToken(context.Background(), clientId).AppsDeleteAuthorizationRequest(appsDeleteAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsDeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID of the GitHub app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsDeleteAuthorizationRequest** | [**AppsDeleteAuthorizationRequest**](AppsDeleteAuthorizationRequest.md) |  | 

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


## AppsGetAuthenticated

> Integration AppsGetAuthenticated(ctx).Execute()

Get the authenticated app



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
    resp, r, err := apiClient.AppsApi.AppsGetAuthenticated(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetAuthenticated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetAuthenticated`: Integration
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetAuthenticated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetAuthenticatedRequest struct via the builder pattern


### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetBySlug

> Integration AppsGetBySlug(ctx, appSlug).Execute()

Get an app



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
    appSlug := "appSlug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetBySlug(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetBySlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetBySlug`: Integration
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetInstallation

> Installation AppsGetInstallation(ctx, installationId).Execute()

Get an installation for the authenticated app



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
    installationId := int32(1) // int32 | The unique identifier of the installation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetInstallation(context.Background(), installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetInstallation`: Installation
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetInstallation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Installation**](Installation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetOrgInstallation

> Installation AppsGetOrgInstallation(ctx, org).Execute()

Get an organization installation for the authenticated app



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
    resp, r, err := apiClient.AppsApi.AppsGetOrgInstallation(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetOrgInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetOrgInstallation`: Installation
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetOrgInstallation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetOrgInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Installation**](Installation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetRepoInstallation

> Installation AppsGetRepoInstallation(ctx, owner, repo).Execute()

Get a repository installation for the authenticated app



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
    resp, r, err := apiClient.AppsApi.AppsGetRepoInstallation(context.Background(), owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetRepoInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetRepoInstallation`: Installation
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetRepoInstallation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetRepoInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Installation**](Installation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetSubscriptionPlanForAccount

> MarketplacePurchase AppsGetSubscriptionPlanForAccount(ctx, accountId).Execute()

Get a subscription plan for an account



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
    accountId := int32(56) // int32 | account_id parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetSubscriptionPlanForAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetSubscriptionPlanForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetSubscriptionPlanForAccount`: MarketplacePurchase
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetSubscriptionPlanForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetSubscriptionPlanForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplacePurchase**](MarketplacePurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetSubscriptionPlanForAccountStubbed

> MarketplacePurchase AppsGetSubscriptionPlanForAccountStubbed(ctx, accountId).Execute()

Get a subscription plan for an account (stubbed)



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
    accountId := int32(56) // int32 | account_id parameter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetSubscriptionPlanForAccountStubbed(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetSubscriptionPlanForAccountStubbed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetSubscriptionPlanForAccountStubbed`: MarketplacePurchase
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetSubscriptionPlanForAccountStubbed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** | account_id parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetSubscriptionPlanForAccountStubbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplacePurchase**](MarketplacePurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetUserInstallation

> Installation AppsGetUserInstallation(ctx, username).Execute()

Get a user installation for the authenticated app



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetUserInstallation(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetUserInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetUserInstallation`: Installation
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetUserInstallation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetUserInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Installation**](Installation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetWebhookConfigForApp

> WebhookConfig AppsGetWebhookConfigForApp(ctx).Execute()

Get a webhook configuration for an app



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
    resp, r, err := apiClient.AppsApi.AppsGetWebhookConfigForApp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetWebhookConfigForApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetWebhookConfigForApp`: WebhookConfig
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetWebhookConfigForApp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetWebhookConfigForAppRequest struct via the builder pattern


### Return type

[**WebhookConfig**](WebhookConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGetWebhookDelivery

> HookDelivery AppsGetWebhookDelivery(ctx, deliveryId).Execute()

Get a delivery for an app webhook



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
    deliveryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsGetWebhookDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsGetWebhookDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsGetWebhookDelivery`: HookDelivery
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsGetWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HookDelivery**](HookDelivery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListAccountsForPlan

> []MarketplacePurchase AppsListAccountsForPlan(ctx, planId).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List accounts for a plan



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
    planId := int32(56) // int32 | The unique identifier of the plan.
    sort := "sort_example" // string | The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to. (optional) (default to "created")
    direction := "direction_example" // string | To return the oldest accounts first, set to `asc`. Ignored without the `sort` parameter. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsListAccountsForPlan(context.Background(), planId).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListAccountsForPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListAccountsForPlan`: []MarketplacePurchase
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListAccountsForPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int32** | The unique identifier of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsListAccountsForPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the repository was starred. &#x60;updated&#x60; means when the repository was last pushed to. | [default to &quot;created&quot;]
 **direction** | **string** | To return the oldest accounts first, set to &#x60;asc&#x60;. Ignored without the &#x60;sort&#x60; parameter. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MarketplacePurchase**](MarketplacePurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListAccountsForPlanStubbed

> []MarketplacePurchase AppsListAccountsForPlanStubbed(ctx, planId).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()

List accounts for a plan (stubbed)



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
    planId := int32(56) // int32 | The unique identifier of the plan.
    sort := "sort_example" // string | The property to sort the results by. `created` means when the repository was starred. `updated` means when the repository was last pushed to. (optional) (default to "created")
    direction := "direction_example" // string | To return the oldest accounts first, set to `asc`. Ignored without the `sort` parameter. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsListAccountsForPlanStubbed(context.Background(), planId).Sort(sort).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListAccountsForPlanStubbed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListAccountsForPlanStubbed`: []MarketplacePurchase
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListAccountsForPlanStubbed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int32** | The unique identifier of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsListAccountsForPlanStubbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the repository was starred. &#x60;updated&#x60; means when the repository was last pushed to. | [default to &quot;created&quot;]
 **direction** | **string** | To return the oldest accounts first, set to &#x60;asc&#x60;. Ignored without the &#x60;sort&#x60; parameter. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MarketplacePurchase**](MarketplacePurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListInstallationReposForAuthenticatedUser

> AppsListInstallationReposForAuthenticatedUser200Response AppsListInstallationReposForAuthenticatedUser(ctx, installationId).PerPage(perPage).Page(page).Execute()

List repositories accessible to the user access token



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
    installationId := int32(1) // int32 | The unique identifier of the installation.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsListInstallationReposForAuthenticatedUser(context.Background(), installationId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListInstallationReposForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListInstallationReposForAuthenticatedUser`: AppsListInstallationReposForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListInstallationReposForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsListInstallationReposForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**AppsListInstallationReposForAuthenticatedUser200Response**](AppsListInstallationReposForAuthenticatedUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListInstallations

> []Installation AppsListInstallations(ctx).PerPage(perPage).Page(page).Since(since).Outdated(outdated).Execute()

List installations for the authenticated app



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    since := time.Now() // time.Time | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. (optional)
    outdated := "outdated_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsListInstallations(context.Background()).PerPage(perPage).Page(page).Since(since).Outdated(outdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListInstallations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListInstallations`: []Installation
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListInstallations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListInstallationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **since** | **time.Time** | Only show notifications updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | 
 **outdated** | **string** |  | 

### Return type

[**[]Installation**](Installation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListInstallationsForAuthenticatedUser

> OrgsListAppInstallations200Response AppsListInstallationsForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List app installations accessible to the user access token



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
    resp, r, err := apiClient.AppsApi.AppsListInstallationsForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListInstallationsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListInstallationsForAuthenticatedUser`: OrgsListAppInstallations200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListInstallationsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListInstallationsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**OrgsListAppInstallations200Response**](OrgsListAppInstallations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListPlans

> []MarketplaceListingPlan AppsListPlans(ctx).PerPage(perPage).Page(page).Execute()

List plans



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
    resp, r, err := apiClient.AppsApi.AppsListPlans(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListPlans`: []MarketplaceListingPlan
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MarketplaceListingPlan**](MarketplaceListingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListPlansStubbed

> []MarketplaceListingPlan AppsListPlansStubbed(ctx).PerPage(perPage).Page(page).Execute()

List plans (stubbed)



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
    resp, r, err := apiClient.AppsApi.AppsListPlansStubbed(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListPlansStubbed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListPlansStubbed`: []MarketplaceListingPlan
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListPlansStubbed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListPlansStubbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]MarketplaceListingPlan**](MarketplaceListingPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListReposAccessibleToInstallation

> AppsListReposAccessibleToInstallation200Response AppsListReposAccessibleToInstallation(ctx).PerPage(perPage).Page(page).Execute()

List repositories accessible to the app installation



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
    resp, r, err := apiClient.AppsApi.AppsListReposAccessibleToInstallation(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListReposAccessibleToInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListReposAccessibleToInstallation`: AppsListReposAccessibleToInstallation200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListReposAccessibleToInstallation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListReposAccessibleToInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**AppsListReposAccessibleToInstallation200Response**](AppsListReposAccessibleToInstallation200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListSubscriptionsForAuthenticatedUser

> []UserMarketplacePurchase AppsListSubscriptionsForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List subscriptions for the authenticated user



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
    resp, r, err := apiClient.AppsApi.AppsListSubscriptionsForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListSubscriptionsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListSubscriptionsForAuthenticatedUser`: []UserMarketplacePurchase
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListSubscriptionsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListSubscriptionsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]UserMarketplacePurchase**](UserMarketplacePurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListSubscriptionsForAuthenticatedUserStubbed

> []UserMarketplacePurchase AppsListSubscriptionsForAuthenticatedUserStubbed(ctx).PerPage(perPage).Page(page).Execute()

List subscriptions for the authenticated user (stubbed)



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
    resp, r, err := apiClient.AppsApi.AppsListSubscriptionsForAuthenticatedUserStubbed(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListSubscriptionsForAuthenticatedUserStubbed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListSubscriptionsForAuthenticatedUserStubbed`: []UserMarketplacePurchase
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListSubscriptionsForAuthenticatedUserStubbed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListSubscriptionsForAuthenticatedUserStubbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]UserMarketplacePurchase**](UserMarketplacePurchase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsListWebhookDeliveries

> []HookDeliveryItem AppsListWebhookDeliveries(ctx).PerPage(perPage).Cursor(cursor).Execute()

List deliveries for an app webhook



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
    cursor := "cursor_example" // string | Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the `link` header for the next and previous page cursors. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsListWebhookDeliveries(context.Background()).PerPage(perPage).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsListWebhookDeliveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsListWebhookDeliveries`: []HookDeliveryItem
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsListWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsListWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **cursor** | **string** | Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the &#x60;link&#x60; header for the next and previous page cursors. | 

### Return type

[**[]HookDeliveryItem**](HookDeliveryItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsRedeliverWebhookDelivery

> map[string]interface{} AppsRedeliverWebhookDelivery(ctx, deliveryId).Execute()

Redeliver a delivery for an app webhook



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
    deliveryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsRedeliverWebhookDelivery(context.Background(), deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsRedeliverWebhookDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsRedeliverWebhookDelivery`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsRedeliverWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsRedeliverWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsRemoveRepoFromInstallationForAuthenticatedUser

> AppsRemoveRepoFromInstallationForAuthenticatedUser(ctx, installationId, repositoryId).Execute()

Remove a repository from an app installation



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
    installationId := int32(1) // int32 | The unique identifier of the installation.
    repositoryId := int32(56) // int32 | The unique identifier of the repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsRemoveRepoFromInstallationForAuthenticatedUser(context.Background(), installationId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsRemoveRepoFromInstallationForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 
**repositoryId** | **int32** | The unique identifier of the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsRemoveRepoFromInstallationForAuthenticatedUserRequest struct via the builder pattern


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


## AppsResetToken

> Authorization AppsResetToken(ctx, clientId).AppsCheckTokenRequest(appsCheckTokenRequest).Execute()

Reset a token



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
    clientId := "Iv1.8a61f9b3a7aba766" // string | The client ID of the GitHub app.
    appsCheckTokenRequest := *openapiclient.NewAppsCheckTokenRequest("AccessToken_example") // AppsCheckTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsResetToken(context.Background(), clientId).AppsCheckTokenRequest(appsCheckTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsResetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsResetToken`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsResetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID of the GitHub app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsResetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsCheckTokenRequest** | [**AppsCheckTokenRequest**](AppsCheckTokenRequest.md) |  | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsRevokeInstallationAccessToken

> AppsRevokeInstallationAccessToken(ctx).Execute()

Revoke an installation access token



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
    resp, r, err := apiClient.AppsApi.AppsRevokeInstallationAccessToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsRevokeInstallationAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsRevokeInstallationAccessTokenRequest struct via the builder pattern


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


## AppsScopeToken

> Authorization AppsScopeToken(ctx, clientId).AppsScopeTokenRequest(appsScopeTokenRequest).Execute()

Create a scoped access token



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
    clientId := "Iv1.8a61f9b3a7aba766" // string | The client ID of the GitHub app.
    appsScopeTokenRequest := *openapiclient.NewAppsScopeTokenRequest("e72e16c7e42f292c6912e7710c838347ae178b4a") // AppsScopeTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsScopeToken(context.Background(), clientId).AppsScopeTokenRequest(appsScopeTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsScopeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsScopeToken`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsScopeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID of the GitHub app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsScopeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsScopeTokenRequest** | [**AppsScopeTokenRequest**](AppsScopeTokenRequest.md) |  | 

### Return type

[**Authorization**](Authorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSuspendInstallation

> AppsSuspendInstallation(ctx, installationId).Execute()

Suspend an app installation



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
    installationId := int32(1) // int32 | The unique identifier of the installation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsSuspendInstallation(context.Background(), installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsSuspendInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSuspendInstallationRequest struct via the builder pattern


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


## AppsUnsuspendInstallation

> AppsUnsuspendInstallation(ctx, installationId).Execute()

Unsuspend an app installation



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
    installationId := int32(1) // int32 | The unique identifier of the installation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsUnsuspendInstallation(context.Background(), installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsUnsuspendInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installationId** | **int32** | The unique identifier of the installation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsUnsuspendInstallationRequest struct via the builder pattern


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


## AppsUpdateWebhookConfigForApp

> WebhookConfig AppsUpdateWebhookConfigForApp(ctx).AppsUpdateWebhookConfigForAppRequest(appsUpdateWebhookConfigForAppRequest).Execute()

Update a webhook configuration for an app



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
    appsUpdateWebhookConfigForAppRequest := *openapiclient.NewAppsUpdateWebhookConfigForAppRequest() // AppsUpdateWebhookConfigForAppRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.AppsUpdateWebhookConfigForApp(context.Background()).AppsUpdateWebhookConfigForAppRequest(appsUpdateWebhookConfigForAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.AppsUpdateWebhookConfigForApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsUpdateWebhookConfigForApp`: WebhookConfig
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.AppsUpdateWebhookConfigForApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsUpdateWebhookConfigForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appsUpdateWebhookConfigForAppRequest** | [**AppsUpdateWebhookConfigForAppRequest**](AppsUpdateWebhookConfigForAppRequest.md) |  | 

### Return type

[**WebhookConfig**](WebhookConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

