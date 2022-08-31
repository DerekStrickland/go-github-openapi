# \CodeScanningApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodeScanningDeleteAnalysis**](CodeScanningApi.md#CodeScanningDeleteAnalysis) | **Delete** /repos/{owner}/{repo}/code-scanning/analyses/{analysis_id} | Delete a code scanning analysis from a repository
[**CodeScanningGetAlert**](CodeScanningApi.md#CodeScanningGetAlert) | **Get** /repos/{owner}/{repo}/code-scanning/alerts/{alert_number} | Get a code scanning alert
[**CodeScanningGetAnalysis**](CodeScanningApi.md#CodeScanningGetAnalysis) | **Get** /repos/{owner}/{repo}/code-scanning/analyses/{analysis_id} | Get a code scanning analysis for a repository
[**CodeScanningGetSarif**](CodeScanningApi.md#CodeScanningGetSarif) | **Get** /repos/{owner}/{repo}/code-scanning/sarifs/{sarif_id} | Get information about a SARIF upload
[**CodeScanningListAlertInstances**](CodeScanningApi.md#CodeScanningListAlertInstances) | **Get** /repos/{owner}/{repo}/code-scanning/alerts/{alert_number}/instances | List instances of a code scanning alert
[**CodeScanningListAlertsForEnterprise**](CodeScanningApi.md#CodeScanningListAlertsForEnterprise) | **Get** /enterprises/{enterprise}/code-scanning/alerts | List code scanning alerts for an enterprise
[**CodeScanningListAlertsForOrg**](CodeScanningApi.md#CodeScanningListAlertsForOrg) | **Get** /orgs/{org}/code-scanning/alerts | List code scanning alerts for an organization
[**CodeScanningListAlertsForRepo**](CodeScanningApi.md#CodeScanningListAlertsForRepo) | **Get** /repos/{owner}/{repo}/code-scanning/alerts | List code scanning alerts for a repository
[**CodeScanningListRecentAnalyses**](CodeScanningApi.md#CodeScanningListRecentAnalyses) | **Get** /repos/{owner}/{repo}/code-scanning/analyses | List code scanning analyses for a repository
[**CodeScanningUpdateAlert**](CodeScanningApi.md#CodeScanningUpdateAlert) | **Patch** /repos/{owner}/{repo}/code-scanning/alerts/{alert_number} | Update a code scanning alert
[**CodeScanningUploadSarif**](CodeScanningApi.md#CodeScanningUploadSarif) | **Post** /repos/{owner}/{repo}/code-scanning/sarifs | Upload an analysis as SARIF data



## CodeScanningDeleteAnalysis

> CodeScanningAnalysisDeletion CodeScanningDeleteAnalysis(ctx, owner, repo, analysisId).ConfirmDelete(confirmDelete).Execute()

Delete a code scanning analysis from a repository



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
    analysisId := int32(56) // int32 | The ID of the analysis, as returned from the `GET /repos/{owner}/{repo}/code-scanning/analyses` operation.
    confirmDelete := "confirmDelete_example" // string | Allow deletion if the specified analysis is the last in a set. If you attempt to delete the final analysis in a set without setting this parameter to `true`, you'll get a 400 response with the message: `Analysis is last of its type and deletion may result in the loss of historical alert data. Please specify confirm_delete.` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningDeleteAnalysis(context.Background(), owner, repo, analysisId).ConfirmDelete(confirmDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningDeleteAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningDeleteAnalysis`: CodeScanningAnalysisDeletion
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningDeleteAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**analysisId** | **int32** | The ID of the analysis, as returned from the &#x60;GET /repos/{owner}/{repo}/code-scanning/analyses&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningDeleteAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **confirmDelete** | **string** | Allow deletion if the specified analysis is the last in a set. If you attempt to delete the final analysis in a set without setting this parameter to &#x60;true&#x60;, you&#39;ll get a 400 response with the message: &#x60;Analysis is last of its type and deletion may result in the loss of historical alert data. Please specify confirm_delete.&#x60; | 

### Return type

[**CodeScanningAnalysisDeletion**](CodeScanningAnalysisDeletion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningGetAlert

> CodeScanningAlert CodeScanningGetAlert(ctx, owner, repo, alertNumber).Execute()

Get a code scanning alert



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
    resp, r, err := apiClient.CodeScanningApi.CodeScanningGetAlert(context.Background(), owner, repo, alertNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningGetAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningGetAlert`: CodeScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningGetAlert`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCodeScanningGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CodeScanningAlert**](CodeScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningGetAnalysis

> CodeScanningAnalysis CodeScanningGetAnalysis(ctx, owner, repo, analysisId).Execute()

Get a code scanning analysis for a repository



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
    analysisId := int32(56) // int32 | The ID of the analysis, as returned from the `GET /repos/{owner}/{repo}/code-scanning/analyses` operation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningGetAnalysis(context.Background(), owner, repo, analysisId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningGetAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningGetAnalysis`: CodeScanningAnalysis
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningGetAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**analysisId** | **int32** | The ID of the analysis, as returned from the &#x60;GET /repos/{owner}/{repo}/code-scanning/analyses&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningGetAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CodeScanningAnalysis**](CodeScanningAnalysis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json+sarif

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningGetSarif

> CodeScanningSarifsStatus CodeScanningGetSarif(ctx, owner, repo, sarifId).Execute()

Get information about a SARIF upload



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
    sarifId := "sarifId_example" // string | The SARIF ID obtained after uploading.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningGetSarif(context.Background(), owner, repo, sarifId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningGetSarif``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningGetSarif`: CodeScanningSarifsStatus
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningGetSarif`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 
**sarifId** | **string** | The SARIF ID obtained after uploading. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningGetSarifRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CodeScanningSarifsStatus**](CodeScanningSarifsStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningListAlertInstances

> []CodeScanningAlertInstance CodeScanningListAlertInstances(ctx, owner, repo, alertNumber).Page(page).PerPage(perPage).Ref(ref).Execute()

List instances of a code scanning alert



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
    ref := "ref_example" // string | The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningListAlertInstances(context.Background(), owner, repo, alertNumber).Page(page).PerPage(perPage).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningListAlertInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningListAlertInstances`: []CodeScanningAlertInstance
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningListAlertInstances`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCodeScanningListAlertInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **ref** | **string** | The Git reference for the results you want to list. The &#x60;ref&#x60; for a branch can be formatted either as &#x60;refs/heads/&lt;branch name&gt;&#x60; or simply &#x60;&lt;branch name&gt;&#x60;. To reference a pull request use &#x60;refs/pull/&lt;number&gt;/merge&#x60;. | 

### Return type

[**[]CodeScanningAlertInstance**](CodeScanningAlertInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningListAlertsForEnterprise

> []CodeScanningOrganizationAlertItems CodeScanningListAlertsForEnterprise(ctx, enterprise).ToolName(toolName).ToolGuid(toolGuid).Before(before).After(after).Page(page).PerPage(perPage).Direction(direction).State(state).Sort(sort).Execute()

List code scanning alerts for an enterprise



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
    toolName := "toolName_example" // string | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both. (optional)
    toolGuid := "toolGuid_example" // string | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both. (optional)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. (optional)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    state := openapiclient.code-scanning-alert-state("open") // CodeScanningAlertState | If specified, only code scanning alerts with this state will be returned. (optional)
    sort := "sort_example" // string | The property by which to sort the results. (optional) (default to "created")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningListAlertsForEnterprise(context.Background(), enterprise).ToolName(toolName).ToolGuid(toolGuid).Before(before).After(after).Page(page).PerPage(perPage).Direction(direction).State(state).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningListAlertsForEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningListAlertsForEnterprise`: []CodeScanningOrganizationAlertItems
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningListAlertsForEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterprise** | **string** | The slug version of the enterprise name. You can also substitute this value with the enterprise id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningListAlertsForEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolName** | **string** | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either &#x60;tool_name&#x60; or &#x60;tool_guid&#x60;, but not both. | 
 **toolGuid** | **string** | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either &#x60;tool_guid&#x60; or &#x60;tool_name&#x60;, but not both. | 
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. | 
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **state** | [**CodeScanningAlertState**](CodeScanningAlertState.md) | If specified, only code scanning alerts with this state will be returned. | 
 **sort** | **string** | The property by which to sort the results. | [default to &quot;created&quot;]

### Return type

[**[]CodeScanningOrganizationAlertItems**](CodeScanningOrganizationAlertItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningListAlertsForOrg

> []CodeScanningOrganizationAlertItems CodeScanningListAlertsForOrg(ctx, org).ToolName(toolName).ToolGuid(toolGuid).Before(before).After(after).Page(page).PerPage(perPage).Direction(direction).State(state).Sort(sort).Execute()

List code scanning alerts for an organization



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
    toolName := "toolName_example" // string | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both. (optional)
    toolGuid := "toolGuid_example" // string | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both. (optional)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. (optional)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    state := openapiclient.code-scanning-alert-state("open") // CodeScanningAlertState | If specified, only code scanning alerts with this state will be returned. (optional)
    sort := "sort_example" // string | The property by which to sort the results. (optional) (default to "created")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningListAlertsForOrg(context.Background(), org).ToolName(toolName).ToolGuid(toolGuid).Before(before).After(after).Page(page).PerPage(perPage).Direction(direction).State(state).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningListAlertsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningListAlertsForOrg`: []CodeScanningOrganizationAlertItems
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningListAlertsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningListAlertsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolName** | **string** | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either &#x60;tool_name&#x60; or &#x60;tool_guid&#x60;, but not both. | 
 **toolGuid** | **string** | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either &#x60;tool_guid&#x60; or &#x60;tool_name&#x60;, but not both. | 
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. | 
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **state** | [**CodeScanningAlertState**](CodeScanningAlertState.md) | If specified, only code scanning alerts with this state will be returned. | 
 **sort** | **string** | The property by which to sort the results. | [default to &quot;created&quot;]

### Return type

[**[]CodeScanningOrganizationAlertItems**](CodeScanningOrganizationAlertItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningListAlertsForRepo

> []CodeScanningAlertItems CodeScanningListAlertsForRepo(ctx, owner, repo).ToolName(toolName).ToolGuid(toolGuid).Page(page).PerPage(perPage).Ref(ref).Direction(direction).Sort(sort).State(state).Execute()

List code scanning alerts for a repository



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
    toolName := "toolName_example" // string | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both. (optional)
    toolGuid := "toolGuid_example" // string | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both. (optional)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    ref := "ref_example" // string | The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`. (optional)
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    sort := "sort_example" // string | The property by which to sort the results. (optional) (default to "created")
    state := openapiclient.code-scanning-alert-state("open") // CodeScanningAlertState | If specified, only code scanning alerts with this state will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningListAlertsForRepo(context.Background(), owner, repo).ToolName(toolName).ToolGuid(toolGuid).Page(page).PerPage(perPage).Ref(ref).Direction(direction).Sort(sort).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningListAlertsForRepo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningListAlertsForRepo`: []CodeScanningAlertItems
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningListAlertsForRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningListAlertsForRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **toolName** | **string** | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either &#x60;tool_name&#x60; or &#x60;tool_guid&#x60;, but not both. | 
 **toolGuid** | **string** | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either &#x60;tool_guid&#x60; or &#x60;tool_name&#x60;, but not both. | 
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **ref** | **string** | The Git reference for the results you want to list. The &#x60;ref&#x60; for a branch can be formatted either as &#x60;refs/heads/&lt;branch name&gt;&#x60; or simply &#x60;&lt;branch name&gt;&#x60;. To reference a pull request use &#x60;refs/pull/&lt;number&gt;/merge&#x60;. | 
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **sort** | **string** | The property by which to sort the results. | [default to &quot;created&quot;]
 **state** | [**CodeScanningAlertState**](CodeScanningAlertState.md) | If specified, only code scanning alerts with this state will be returned. | 

### Return type

[**[]CodeScanningAlertItems**](CodeScanningAlertItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningListRecentAnalyses

> []CodeScanningAnalysis CodeScanningListRecentAnalyses(ctx, owner, repo).ToolName(toolName).ToolGuid(toolGuid).Page(page).PerPage(perPage).Ref(ref).SarifId(sarifId).Direction(direction).Sort(sort).Execute()

List code scanning analyses for a repository



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
    toolName := "toolName_example" // string | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both. (optional)
    toolGuid := "toolGuid_example" // string | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both. (optional)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    ref := "ref_example" // string | The Git reference for the analyses you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`. (optional)
    sarifId := "sarifId_example" // string | Filter analyses belonging to the same SARIF upload. (optional)
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    sort := "sort_example" // string | The property by which to sort the results. (optional) (default to "created")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningListRecentAnalyses(context.Background(), owner, repo).ToolName(toolName).ToolGuid(toolGuid).Page(page).PerPage(perPage).Ref(ref).SarifId(sarifId).Direction(direction).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningListRecentAnalyses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningListRecentAnalyses`: []CodeScanningAnalysis
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningListRecentAnalyses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningListRecentAnalysesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **toolName** | **string** | The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either &#x60;tool_name&#x60; or &#x60;tool_guid&#x60;, but not both. | 
 **toolGuid** | **string** | The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either &#x60;tool_guid&#x60; or &#x60;tool_name&#x60;, but not both. | 
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **ref** | **string** | The Git reference for the analyses you want to list. The &#x60;ref&#x60; for a branch can be formatted either as &#x60;refs/heads/&lt;branch name&gt;&#x60; or simply &#x60;&lt;branch name&gt;&#x60;. To reference a pull request use &#x60;refs/pull/&lt;number&gt;/merge&#x60;. | 
 **sarifId** | **string** | Filter analyses belonging to the same SARIF upload. | 
 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **sort** | **string** | The property by which to sort the results. | [default to &quot;created&quot;]

### Return type

[**[]CodeScanningAnalysis**](CodeScanningAnalysis.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningUpdateAlert

> CodeScanningAlert CodeScanningUpdateAlert(ctx, owner, repo, alertNumber).CodeScanningUpdateAlertRequest(codeScanningUpdateAlertRequest).Execute()

Update a code scanning alert



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
    codeScanningUpdateAlertRequest := *openapiclient.NewCodeScanningUpdateAlertRequest(openapiclient.code-scanning-alert-set-state("open")) // CodeScanningUpdateAlertRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningUpdateAlert(context.Background(), owner, repo, alertNumber).CodeScanningUpdateAlertRequest(codeScanningUpdateAlertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningUpdateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningUpdateAlert`: CodeScanningAlert
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningUpdateAlert`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCodeScanningUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **codeScanningUpdateAlertRequest** | [**CodeScanningUpdateAlertRequest**](CodeScanningUpdateAlertRequest.md) |  | 

### Return type

[**CodeScanningAlert**](CodeScanningAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeScanningUploadSarif

> CodeScanningSarifsReceipt CodeScanningUploadSarif(ctx, owner, repo).CodeScanningUploadSarifRequest(codeScanningUploadSarifRequest).Execute()

Upload an analysis as SARIF data



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
    codeScanningUploadSarifRequest := *openapiclient.NewCodeScanningUploadSarifRequest("CommitSha_example", "Ref_example", "Sarif_example") // CodeScanningUploadSarifRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeScanningApi.CodeScanningUploadSarif(context.Background(), owner, repo).CodeScanningUploadSarifRequest(codeScanningUploadSarifRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeScanningApi.CodeScanningUploadSarif``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CodeScanningUploadSarif`: CodeScanningSarifsReceipt
    fmt.Fprintf(os.Stdout, "Response from `CodeScanningApi.CodeScanningUploadSarif`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCodeScanningUploadSarifRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **codeScanningUploadSarifRequest** | [**CodeScanningUploadSarifRequest**](CodeScanningUploadSarifRequest.md) |  | 

### Return type

[**CodeScanningSarifsReceipt**](CodeScanningSarifsReceipt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

