# \PackagesApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PackagesDeletePackageForAuthenticatedUser**](PackagesApi.md#PackagesDeletePackageForAuthenticatedUser) | **Delete** /user/packages/{package_type}/{package_name} | Delete a package for the authenticated user
[**PackagesDeletePackageForOrg**](PackagesApi.md#PackagesDeletePackageForOrg) | **Delete** /orgs/{org}/packages/{package_type}/{package_name} | Delete a package for an organization
[**PackagesDeletePackageForUser**](PackagesApi.md#PackagesDeletePackageForUser) | **Delete** /users/{username}/packages/{package_type}/{package_name} | Delete a package for a user
[**PackagesDeletePackageVersionForAuthenticatedUser**](PackagesApi.md#PackagesDeletePackageVersionForAuthenticatedUser) | **Delete** /user/packages/{package_type}/{package_name}/versions/{package_version_id} | Delete a package version for the authenticated user
[**PackagesDeletePackageVersionForOrg**](PackagesApi.md#PackagesDeletePackageVersionForOrg) | **Delete** /orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id} | Delete package version for an organization
[**PackagesDeletePackageVersionForUser**](PackagesApi.md#PackagesDeletePackageVersionForUser) | **Delete** /users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id} | Delete package version for a user
[**PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser**](PackagesApi.md#PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser) | **Get** /user/packages/{package_type}/{package_name}/versions | Get all package versions for a package owned by the authenticated user
[**PackagesGetAllPackageVersionsForPackageOwnedByOrg**](PackagesApi.md#PackagesGetAllPackageVersionsForPackageOwnedByOrg) | **Get** /orgs/{org}/packages/{package_type}/{package_name}/versions | Get all package versions for a package owned by an organization
[**PackagesGetAllPackageVersionsForPackageOwnedByUser**](PackagesApi.md#PackagesGetAllPackageVersionsForPackageOwnedByUser) | **Get** /users/{username}/packages/{package_type}/{package_name}/versions | Get all package versions for a package owned by a user
[**PackagesGetPackageForAuthenticatedUser**](PackagesApi.md#PackagesGetPackageForAuthenticatedUser) | **Get** /user/packages/{package_type}/{package_name} | Get a package for the authenticated user
[**PackagesGetPackageForOrganization**](PackagesApi.md#PackagesGetPackageForOrganization) | **Get** /orgs/{org}/packages/{package_type}/{package_name} | Get a package for an organization
[**PackagesGetPackageForUser**](PackagesApi.md#PackagesGetPackageForUser) | **Get** /users/{username}/packages/{package_type}/{package_name} | Get a package for a user
[**PackagesGetPackageVersionForAuthenticatedUser**](PackagesApi.md#PackagesGetPackageVersionForAuthenticatedUser) | **Get** /user/packages/{package_type}/{package_name}/versions/{package_version_id} | Get a package version for the authenticated user
[**PackagesGetPackageVersionForOrganization**](PackagesApi.md#PackagesGetPackageVersionForOrganization) | **Get** /orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id} | Get a package version for an organization
[**PackagesGetPackageVersionForUser**](PackagesApi.md#PackagesGetPackageVersionForUser) | **Get** /users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id} | Get a package version for a user
[**PackagesListPackagesForAuthenticatedUser**](PackagesApi.md#PackagesListPackagesForAuthenticatedUser) | **Get** /user/packages | List packages for the authenticated user&#39;s namespace
[**PackagesListPackagesForOrganization**](PackagesApi.md#PackagesListPackagesForOrganization) | **Get** /orgs/{org}/packages | List packages for an organization
[**PackagesListPackagesForUser**](PackagesApi.md#PackagesListPackagesForUser) | **Get** /users/{username}/packages | List packages for a user
[**PackagesRestorePackageForAuthenticatedUser**](PackagesApi.md#PackagesRestorePackageForAuthenticatedUser) | **Post** /user/packages/{package_type}/{package_name}/restore | Restore a package for the authenticated user
[**PackagesRestorePackageForOrg**](PackagesApi.md#PackagesRestorePackageForOrg) | **Post** /orgs/{org}/packages/{package_type}/{package_name}/restore | Restore a package for an organization
[**PackagesRestorePackageForUser**](PackagesApi.md#PackagesRestorePackageForUser) | **Post** /users/{username}/packages/{package_type}/{package_name}/restore | Restore a package for a user
[**PackagesRestorePackageVersionForAuthenticatedUser**](PackagesApi.md#PackagesRestorePackageVersionForAuthenticatedUser) | **Post** /user/packages/{package_type}/{package_name}/versions/{package_version_id}/restore | Restore a package version for the authenticated user
[**PackagesRestorePackageVersionForOrg**](PackagesApi.md#PackagesRestorePackageVersionForOrg) | **Post** /orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore | Restore package version for an organization
[**PackagesRestorePackageVersionForUser**](PackagesApi.md#PackagesRestorePackageVersionForUser) | **Post** /users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore | Restore package version for a user



## PackagesDeletePackageForAuthenticatedUser

> PackagesDeletePackageForAuthenticatedUser(ctx, packageType, packageName).Execute()

Delete a package for the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDeletePackageForAuthenticatedUser(context.Background(), packageType, packageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDeletePackageForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesDeletePackageForAuthenticatedUserRequest struct via the builder pattern


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


## PackagesDeletePackageForOrg

> PackagesDeletePackageForOrg(ctx, packageType, packageName, org).Execute()

Delete a package for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDeletePackageForOrg(context.Background(), packageType, packageName, org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDeletePackageForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesDeletePackageForOrgRequest struct via the builder pattern


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


## PackagesDeletePackageForUser

> PackagesDeletePackageForUser(ctx, packageType, packageName, username).Execute()

Delete a package for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDeletePackageForUser(context.Background(), packageType, packageName, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDeletePackageForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesDeletePackageForUserRequest struct via the builder pattern


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


## PackagesDeletePackageVersionForAuthenticatedUser

> PackagesDeletePackageVersionForAuthenticatedUser(ctx, packageType, packageName, packageVersionId).Execute()

Delete a package version for the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDeletePackageVersionForAuthenticatedUser(context.Background(), packageType, packageName, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDeletePackageVersionForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesDeletePackageVersionForAuthenticatedUserRequest struct via the builder pattern


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


## PackagesDeletePackageVersionForOrg

> PackagesDeletePackageVersionForOrg(ctx, packageType, packageName, org, packageVersionId).Execute()

Delete package version for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDeletePackageVersionForOrg(context.Background(), packageType, packageName, org, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDeletePackageVersionForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesDeletePackageVersionForOrgRequest struct via the builder pattern


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


## PackagesDeletePackageVersionForUser

> PackagesDeletePackageVersionForUser(ctx, packageType, packageName, username, packageVersionId).Execute()

Delete package version for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    username := "username_example" // string | The handle for the GitHub user account.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesDeletePackageVersionForUser(context.Background(), packageType, packageName, username, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesDeletePackageVersionForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**username** | **string** | The handle for the GitHub user account. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesDeletePackageVersionForUserRequest struct via the builder pattern


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


## PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser

> []PackageVersion PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser(ctx, packageType, packageName).Page(page).PerPage(perPage).State(state).Execute()

Get all package versions for a package owned by the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    state := "state_example" // string | The state of the package, either active or deleted. (optional) (default to "active")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser(context.Background(), packageType, packageName).Page(page).PerPage(perPage).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser`: []PackageVersion
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetAllPackageVersionsForPackageOwnedByAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **state** | **string** | The state of the package, either active or deleted. | [default to &quot;active&quot;]

### Return type

[**[]PackageVersion**](PackageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetAllPackageVersionsForPackageOwnedByOrg

> []PackageVersion PackagesGetAllPackageVersionsForPackageOwnedByOrg(ctx, packageType, packageName, org).Page(page).PerPage(perPage).State(state).Execute()

Get all package versions for a package owned by an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    state := "state_example" // string | The state of the package, either active or deleted. (optional) (default to "active")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByOrg(context.Background(), packageType, packageName, org).Page(page).PerPage(perPage).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetAllPackageVersionsForPackageOwnedByOrg`: []PackageVersion
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetAllPackageVersionsForPackageOwnedByOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **state** | **string** | The state of the package, either active or deleted. | [default to &quot;active&quot;]

### Return type

[**[]PackageVersion**](PackageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetAllPackageVersionsForPackageOwnedByUser

> []PackageVersion PackagesGetAllPackageVersionsForPackageOwnedByUser(ctx, packageType, packageName, username).Execute()

Get all package versions for a package owned by a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByUser(context.Background(), packageType, packageName, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetAllPackageVersionsForPackageOwnedByUser`: []PackageVersion
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetAllPackageVersionsForPackageOwnedByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetAllPackageVersionsForPackageOwnedByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]PackageVersion**](PackageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetPackageForAuthenticatedUser

> ModelPackage PackagesGetPackageForAuthenticatedUser(ctx, packageType, packageName).Execute()

Get a package for the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetPackageForAuthenticatedUser(context.Background(), packageType, packageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetPackageForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetPackageForAuthenticatedUser`: ModelPackage
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetPackageForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetPackageForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelPackage**](ModelPackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetPackageForOrganization

> ModelPackage PackagesGetPackageForOrganization(ctx, packageType, packageName, org).Execute()

Get a package for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetPackageForOrganization(context.Background(), packageType, packageName, org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetPackageForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetPackageForOrganization`: ModelPackage
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetPackageForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetPackageForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ModelPackage**](ModelPackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetPackageForUser

> ModelPackage PackagesGetPackageForUser(ctx, packageType, packageName, username).Execute()

Get a package for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetPackageForUser(context.Background(), packageType, packageName, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetPackageForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetPackageForUser`: ModelPackage
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetPackageForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetPackageForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ModelPackage**](ModelPackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetPackageVersionForAuthenticatedUser

> PackageVersion PackagesGetPackageVersionForAuthenticatedUser(ctx, packageType, packageName, packageVersionId).Execute()

Get a package version for the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetPackageVersionForAuthenticatedUser(context.Background(), packageType, packageName, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetPackageVersionForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetPackageVersionForAuthenticatedUser`: PackageVersion
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetPackageVersionForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetPackageVersionForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PackageVersion**](PackageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetPackageVersionForOrganization

> PackageVersion PackagesGetPackageVersionForOrganization(ctx, packageType, packageName, org, packageVersionId).Execute()

Get a package version for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetPackageVersionForOrganization(context.Background(), packageType, packageName, org, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetPackageVersionForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetPackageVersionForOrganization`: PackageVersion
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetPackageVersionForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetPackageVersionForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PackageVersion**](PackageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesGetPackageVersionForUser

> PackageVersion PackagesGetPackageVersionForUser(ctx, packageType, packageName, packageVersionId, username).Execute()

Get a package version for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesGetPackageVersionForUser(context.Background(), packageType, packageName, packageVersionId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesGetPackageVersionForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesGetPackageVersionForUser`: PackageVersion
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesGetPackageVersionForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesGetPackageVersionForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PackageVersion**](PackageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesListPackagesForAuthenticatedUser

> []ModelPackage PackagesListPackagesForAuthenticatedUser(ctx).PackageType(packageType).Visibility(visibility).Execute()

List packages for the authenticated user's namespace



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    visibility := "visibility_example" // string | The selected visibility of the packages. Only `container` package_types currently support `internal` visibility properly. For other ecosystems `internal` is synonymous with `private`. This parameter is optional and only filters an existing result set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesListPackagesForAuthenticatedUser(context.Background()).PackageType(packageType).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesListPackagesForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesListPackagesForAuthenticatedUser`: []ModelPackage
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesListPackagesForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPackagesListPackagesForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
 **visibility** | **string** | The selected visibility of the packages. Only &#x60;container&#x60; package_types currently support &#x60;internal&#x60; visibility properly. For other ecosystems &#x60;internal&#x60; is synonymous with &#x60;private&#x60;. This parameter is optional and only filters an existing result set. | 

### Return type

[**[]ModelPackage**](ModelPackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesListPackagesForOrganization

> []ModelPackage PackagesListPackagesForOrganization(ctx, org).PackageType(packageType).Visibility(visibility).Execute()

List packages for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    org := "org_example" // string | The organization name. The name is not case sensitive.
    visibility := "visibility_example" // string | The selected visibility of the packages. Only `container` package_types currently support `internal` visibility properly. For other ecosystems `internal` is synonymous with `private`. This parameter is optional and only filters an existing result set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesListPackagesForOrganization(context.Background(), org).PackageType(packageType).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesListPackagesForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesListPackagesForOrganization`: []ModelPackage
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesListPackagesForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesListPackagesForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 

 **visibility** | **string** | The selected visibility of the packages. Only &#x60;container&#x60; package_types currently support &#x60;internal&#x60; visibility properly. For other ecosystems &#x60;internal&#x60; is synonymous with &#x60;private&#x60;. This parameter is optional and only filters an existing result set. | 

### Return type

[**[]ModelPackage**](ModelPackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesListPackagesForUser

> []ModelPackage PackagesListPackagesForUser(ctx, username).PackageType(packageType).Visibility(visibility).Execute()

List packages for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    username := "username_example" // string | The handle for the GitHub user account.
    visibility := "visibility_example" // string | The selected visibility of the packages. Only `container` package_types currently support `internal` visibility properly. For other ecosystems `internal` is synonymous with `private`. This parameter is optional and only filters an existing result set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesListPackagesForUser(context.Background(), username).PackageType(packageType).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesListPackagesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PackagesListPackagesForUser`: []ModelPackage
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.PackagesListPackagesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesListPackagesForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 

 **visibility** | **string** | The selected visibility of the packages. Only &#x60;container&#x60; package_types currently support &#x60;internal&#x60; visibility properly. For other ecosystems &#x60;internal&#x60; is synonymous with &#x60;private&#x60;. This parameter is optional and only filters an existing result set. | 

### Return type

[**[]ModelPackage**](ModelPackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackagesRestorePackageForAuthenticatedUser

> PackagesRestorePackageForAuthenticatedUser(ctx, packageType, packageName).Token(token).Execute()

Restore a package for the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    token := "token_example" // string | package token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRestorePackageForAuthenticatedUser(context.Background(), packageType, packageName).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRestorePackageForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesRestorePackageForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | package token | 

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


## PackagesRestorePackageForOrg

> PackagesRestorePackageForOrg(ctx, packageType, packageName, org).Token(token).Execute()

Restore a package for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.
    token := "token_example" // string | package token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRestorePackageForOrg(context.Background(), packageType, packageName, org).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRestorePackageForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesRestorePackageForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **token** | **string** | package token | 

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


## PackagesRestorePackageForUser

> PackagesRestorePackageForUser(ctx, packageType, packageName, username).Token(token).Execute()

Restore a package for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    username := "username_example" // string | The handle for the GitHub user account.
    token := "token_example" // string | package token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRestorePackageForUser(context.Background(), packageType, packageName, username).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRestorePackageForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesRestorePackageForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **token** | **string** | package token | 

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


## PackagesRestorePackageVersionForAuthenticatedUser

> PackagesRestorePackageVersionForAuthenticatedUser(ctx, packageType, packageName, packageVersionId).Execute()

Restore a package version for the authenticated user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRestorePackageVersionForAuthenticatedUser(context.Background(), packageType, packageName, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRestorePackageVersionForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesRestorePackageVersionForAuthenticatedUserRequest struct via the builder pattern


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


## PackagesRestorePackageVersionForOrg

> PackagesRestorePackageVersionForOrg(ctx, packageType, packageName, org, packageVersionId).Execute()

Restore package version for an organization



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    org := "org_example" // string | The organization name. The name is not case sensitive.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRestorePackageVersionForOrg(context.Background(), packageType, packageName, org, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRestorePackageVersionForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**org** | **string** | The organization name. The name is not case sensitive. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesRestorePackageVersionForOrgRequest struct via the builder pattern


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


## PackagesRestorePackageVersionForUser

> PackagesRestorePackageVersionForUser(ctx, packageType, packageName, username, packageVersionId).Execute()

Restore package version for a user



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
    packageType := "packageType_example" // string | The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.
    packageName := "packageName_example" // string | The name of the package.
    username := "username_example" // string | The handle for the GitHub user account.
    packageVersionId := int32(56) // int32 | Unique identifier of the package version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PackagesApi.PackagesRestorePackageVersionForUser(context.Background(), packageType, packageName, username, packageVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.PackagesRestorePackageVersionForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageType** | **string** | The type of supported package. Packages in GitHub&#39;s Gradle registry have the type &#x60;maven&#x60;. Docker images pushed to GitHub&#39;s Container registry (&#x60;ghcr.io&#x60;) have the type &#x60;container&#x60;. You can use the type &#x60;docker&#x60; to find images that were pushed to GitHub&#39;s Docker registry (&#x60;docker.pkg.github.com&#x60;), even if these have now been migrated to the Container registry. | 
**packageName** | **string** | The name of the package. | 
**username** | **string** | The handle for the GitHub user account. | 
**packageVersionId** | **int32** | Unique identifier of the package version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPackagesRestorePackageVersionForUserRequest struct via the builder pattern


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

