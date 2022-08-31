# \SecretScanningApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretScanningGetAlert**](SecretScanningApi.md#SecretScanningGetAlert) | **Get** /repos/{owner}/{repo}/secret-scanning/alerts/{alert_number} | Get a secret scanning alert
[**SecretScanningListAlertsForEnterprise**](SecretScanningApi.md#SecretScanningListAlertsForEnterprise) | **Get** /enterprises/{enterprise}/secret-scanning/alerts | List secret scanning alerts for an enterprise
[**SecretScanningListAlertsForOrg**](SecretScanningApi.md#SecretScanningListAlertsForOrg) | **Get** /orgs/{org}/secret-scanning/alerts | List secret scanning alerts for an organization
[**SecretScanningListAlertsForRepo**](SecretScanningApi.md#SecretScanningListAlertsForRepo) | **Get** /repos/{owner}/{repo}/secret-scanning/alerts | List secret scanning alerts for a repository
[**SecretScanningListLocationsForAlert**](SecretScanningApi.md#SecretScanningListLocationsForAlert) | **Get** /repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}/locations | List locations for a secret scanning alert
[**SecretScanningUpdateAlert**](SecretScanningApi.md#SecretScanningUpdateAlert) | **Patch** /repos/{owner}/{repo}/secret-scanning/alerts/{alert_number} | Update a secret scanning alert



## SecretScanningGetAlert

> SecretScanningAlert SecretScanningGetAlert(ctx, owner, repo, alertNumber).Execute()

Get a secret scanning alert



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
    alertNumber := int32(56) // int32 | The number that identifies an alert. You can find this at the end of the URL for a code scanning alert within GitHub, and in the `number` field in the response from the `GET /repos/{owner}/{repo}/code-scanning/alerts` operation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanningApi.SecretScanningGetAlert(context.Background(), owner, repo, alertNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanningApi.SecretScanningGetAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretScanningGetAlert`: SecretScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `SecretScanningApi.SecretScanningGetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**alertNumber** | **int32** | The number that identifies an alert. You can find this at the end of the URL for a code scanning alert within GitHub, and in the &#x60;number&#x60; field in the response from the &#x60;GET /repos/{owner}/{repo}/code-scanning/alerts&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretScanningGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SecretScanningAlert**](SecretScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretScanningListAlertsForEnterprise

> []OrganizationSecretScanningAlert SecretScanningListAlertsForEnterprise(ctx, enterprise).State(state).SecretType(secretType).Resolution(resolution).Sort(sort).Direction(direction).PerPage(perPage).Before(before).After(after).Execute()

List secret scanning alerts for an enterprise



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
    enterprise := "enterprise_example" // string | The slug version of the enterprise name. You can also substitute this value with the enterprise id.
    state := "state_example" // string | Set to `open` or `resolved` to only list secret scanning alerts in a specific state. (optional)
    secretType := "secretType_example" // string | A comma-separated list of secret types to return. By default all secret types are returned. See \"[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)\" for a complete list of secret types. (optional)
    resolution := "resolution_example" // string | A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are `false_positive`, `wont_fix`, `revoked`, `pattern_edited`, `pattern_deleted` or `used_in_tests`. (optional)
    sort := "sort_example" // string | The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanningApi.SecretScanningListAlertsForEnterprise(context.Background(), enterprise).State(state).SecretType(secretType).Resolution(resolution).Sort(sort).Direction(direction).PerPage(perPage).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanningApi.SecretScanningListAlertsForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretScanningListAlertsForEnterprise`: []OrganizationSecretScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `SecretScanningApi.SecretScanningListAlertsForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretScanningListAlertsForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Set to &#x60;open&#x60; or &#x60;resolved&#x60; to only list secret scanning alerts in a specific state. | 
 **secretType** | **string** | A comma-separated list of secret types to return. By default all secret types are returned. See \&quot;[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)\&quot; for a complete list of secret types. | 
 **resolution** | **string** | A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are &#x60;false_positive&#x60;, &#x60;wont_fix&#x60;, &#x60;revoked&#x60;, &#x60;pattern_edited&#x60;, &#x60;pattern_deleted&#x60; or &#x60;used_in_tests&#x60;. | 
 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the alert was created. &#x60;updated&#x60; means when the alert was updated or resolved. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. | 

### Return type

[**[]OrganizationSecretScanningAlert**](OrganizationSecretScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretScanningListAlertsForOrg

> []OrganizationSecretScanningAlert SecretScanningListAlertsForOrg(ctx, org).State(state).SecretType(secretType).Resolution(resolution).Sort(sort).Direction(direction).Page(page).PerPage(perPage).Before(before).After(after).Execute()

List secret scanning alerts for an organization



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
    state := "state_example" // string | Set to `open` or `resolved` to only list secret scanning alerts in a specific state. (optional)
    secretType := "secretType_example" // string | A comma-separated list of secret types to return. By default all secret types are returned. See \"[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)\" for a complete list of secret types. (optional)
    resolution := "resolution_example" // string | A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are `false_positive`, `wont_fix`, `revoked`, `pattern_edited`, `pattern_deleted` or `used_in_tests`. (optional)
    sort := "sort_example" // string | The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty \"before\" query string. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty \"after\" query string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanningApi.SecretScanningListAlertsForOrg(context.Background(), org).State(state).SecretType(secretType).Resolution(resolution).Sort(sort).Direction(direction).Page(page).PerPage(perPage).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanningApi.SecretScanningListAlertsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretScanningListAlertsForOrg`: []OrganizationSecretScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `SecretScanningApi.SecretScanningListAlertsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretScanningListAlertsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Set to &#x60;open&#x60; or &#x60;resolved&#x60; to only list secret scanning alerts in a specific state. | 
 **secretType** | **string** | A comma-separated list of secret types to return. By default all secret types are returned. See \&quot;[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)\&quot; for a complete list of secret types. | 
 **resolution** | **string** | A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are &#x60;false_positive&#x60;, &#x60;wont_fix&#x60;, &#x60;revoked&#x60;, &#x60;pattern_edited&#x60;, &#x60;pattern_deleted&#x60; or &#x60;used_in_tests&#x60;. | 
 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the alert was created. &#x60;updated&#x60; means when the alert was updated or resolved. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty \&quot;before\&quot; query string. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty \&quot;after\&quot; query string. | 

### Return type

[**[]OrganizationSecretScanningAlert**](OrganizationSecretScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretScanningListAlertsForRepo

> []SecretScanningAlert SecretScanningListAlertsForRepo(ctx, owner, repo).State(state).SecretType(secretType).Resolution(resolution).Sort(sort).Direction(direction).Page(page).PerPage(perPage).Before(before).After(after).Execute()

List secret scanning alerts for a repository



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
    state := "state_example" // string | Set to `open` or `resolved` to only list secret scanning alerts in a specific state. (optional)
    secretType := "secretType_example" // string | A comma-separated list of secret types to return. By default all secret types are returned. See \"[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)\" for a complete list of secret types. (optional)
    resolution := "resolution_example" // string | A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are `false_positive`, `wont_fix`, `revoked`, `pattern_edited`, `pattern_deleted` or `used_in_tests`. (optional)
    sort := "sort_example" // string | The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved. (optional) (default to "created")
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty \"before\" query string. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty \"after\" query string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanningApi.SecretScanningListAlertsForRepo(context.Background(), owner, repo).State(state).SecretType(secretType).Resolution(resolution).Sort(sort).Direction(direction).Page(page).PerPage(perPage).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanningApi.SecretScanningListAlertsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretScanningListAlertsForRepo`: []SecretScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `SecretScanningApi.SecretScanningListAlertsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretScanningListAlertsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | Set to &#x60;open&#x60; or &#x60;resolved&#x60; to only list secret scanning alerts in a specific state. | 
 **secretType** | **string** | A comma-separated list of secret types to return. By default all secret types are returned. See \&quot;[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)\&quot; for a complete list of secret types. | 
 **resolution** | **string** | A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are &#x60;false_positive&#x60;, &#x60;wont_fix&#x60;, &#x60;revoked&#x60;, &#x60;pattern_edited&#x60;, &#x60;pattern_deleted&#x60; or &#x60;used_in_tests&#x60;. | 
 **sort** | **string** | The property to sort the results by. &#x60;created&#x60; means when the alert was created. &#x60;updated&#x60; means when the alert was updated or resolved. | [default to &quot;created&quot;]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty \&quot;before\&quot; query string. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty \&quot;after\&quot; query string. | 

### Return type

[**[]SecretScanningAlert**](SecretScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretScanningListLocationsForAlert

> []SecretScanningLocation SecretScanningListLocationsForAlert(ctx, owner, repo, alertNumber).Page(page).PerPage(perPage).Execute()

List locations for a secret scanning alert



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
    alertNumber := int32(56) // int32 | The number that identifies an alert. You can find this at the end of the URL for a code scanning alert within GitHub, and in the `number` field in the response from the `GET /repos/{owner}/{repo}/code-scanning/alerts` operation.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanningApi.SecretScanningListLocationsForAlert(context.Background(), owner, repo, alertNumber).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanningApi.SecretScanningListLocationsForAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretScanningListLocationsForAlert`: []SecretScanningLocation
    fmt.Fprintf(os.Stdout, "Response from `SecretScanningApi.SecretScanningListLocationsForAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**alertNumber** | **int32** | The number that identifies an alert. You can find this at the end of the URL for a code scanning alert within GitHub, and in the &#x60;number&#x60; field in the response from the &#x60;GET /repos/{owner}/{repo}/code-scanning/alerts&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretScanningListLocationsForAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**[]SecretScanningLocation**](SecretScanningLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretScanningUpdateAlert

> SecretScanningAlert SecretScanningUpdateAlert(ctx, owner, repo, alertNumber).SecretScanningUpdateAlertRequest(secretScanningUpdateAlertRequest).Execute()

Update a secret scanning alert



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
    alertNumber := int32(56) // int32 | The number that identifies an alert. You can find this at the end of the URL for a code scanning alert within GitHub, and in the `number` field in the response from the `GET /repos/{owner}/{repo}/code-scanning/alerts` operation.
    secretScanningUpdateAlertRequest := *openapiclient.NewSecretScanningUpdateAlertRequest(openapiclient.secret-scanning-alert-state("open")) // SecretScanningUpdateAlertRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanningApi.SecretScanningUpdateAlert(context.Background(), owner, repo, alertNumber).SecretScanningUpdateAlertRequest(secretScanningUpdateAlertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanningApi.SecretScanningUpdateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretScanningUpdateAlert`: SecretScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `SecretScanningApi.SecretScanningUpdateAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**alertNumber** | **int32** | The number that identifies an alert. You can find this at the end of the URL for a code scanning alert within GitHub, and in the &#x60;number&#x60; field in the response from the &#x60;GET /repos/{owner}/{repo}/code-scanning/alerts&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretScanningUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **secretScanningUpdateAlertRequest** | [**SecretScanningUpdateAlertRequest**](SecretScanningUpdateAlertRequest.md) |  | 

### Return type

[**SecretScanningAlert**](SecretScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

