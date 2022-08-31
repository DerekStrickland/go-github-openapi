# \UsersApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersAddEmailForAuthenticatedUser**](UsersApi.md#UsersAddEmailForAuthenticatedUser) | **Post** /user/emails | Add an email address for the authenticated user
[**UsersBlock**](UsersApi.md#UsersBlock) | **Put** /user/blocks/{username} | Block a user
[**UsersCheckBlocked**](UsersApi.md#UsersCheckBlocked) | **Get** /user/blocks/{username} | Check if a user is blocked by the authenticated user
[**UsersCheckFollowingForUser**](UsersApi.md#UsersCheckFollowingForUser) | **Get** /users/{username}/following/{target_user} | Check if a user follows another user
[**UsersCheckPersonIsFollowedByAuthenticated**](UsersApi.md#UsersCheckPersonIsFollowedByAuthenticated) | **Get** /user/following/{username} | Check if a person is followed by the authenticated user
[**UsersCreateGpgKeyForAuthenticatedUser**](UsersApi.md#UsersCreateGpgKeyForAuthenticatedUser) | **Post** /user/gpg_keys | Create a GPG key for the authenticated user
[**UsersCreatePublicSshKeyForAuthenticatedUser**](UsersApi.md#UsersCreatePublicSshKeyForAuthenticatedUser) | **Post** /user/keys | Create a public SSH key for the authenticated user
[**UsersCreateSshSigningKeyForAuthenticatedUser**](UsersApi.md#UsersCreateSshSigningKeyForAuthenticatedUser) | **Post** /user/ssh_signing_keys | Create a SSH signing key for the authenticated user
[**UsersDeleteEmailForAuthenticatedUser**](UsersApi.md#UsersDeleteEmailForAuthenticatedUser) | **Delete** /user/emails | Delete an email address for the authenticated user
[**UsersDeleteGpgKeyForAuthenticatedUser**](UsersApi.md#UsersDeleteGpgKeyForAuthenticatedUser) | **Delete** /user/gpg_keys/{gpg_key_id} | Delete a GPG key for the authenticated user
[**UsersDeletePublicSshKeyForAuthenticatedUser**](UsersApi.md#UsersDeletePublicSshKeyForAuthenticatedUser) | **Delete** /user/keys/{key_id} | Delete a public SSH key for the authenticated user
[**UsersDeleteSshSigningKeyForAuthenticatedUser**](UsersApi.md#UsersDeleteSshSigningKeyForAuthenticatedUser) | **Delete** /user/ssh_signing_keys/{ssh_signing_key_id} | Delete an SSH signing key for the authenticated user
[**UsersFollow**](UsersApi.md#UsersFollow) | **Put** /user/following/{username} | Follow a user
[**UsersGetAuthenticated**](UsersApi.md#UsersGetAuthenticated) | **Get** /user | Get the authenticated user
[**UsersGetByUsername**](UsersApi.md#UsersGetByUsername) | **Get** /users/{username} | Get a user
[**UsersGetContextForUser**](UsersApi.md#UsersGetContextForUser) | **Get** /users/{username}/hovercard | Get contextual information for a user
[**UsersGetGpgKeyForAuthenticatedUser**](UsersApi.md#UsersGetGpgKeyForAuthenticatedUser) | **Get** /user/gpg_keys/{gpg_key_id} | Get a GPG key for the authenticated user
[**UsersGetPublicSshKeyForAuthenticatedUser**](UsersApi.md#UsersGetPublicSshKeyForAuthenticatedUser) | **Get** /user/keys/{key_id} | Get a public SSH key for the authenticated user
[**UsersGetSshSigningKeyForAuthenticatedUser**](UsersApi.md#UsersGetSshSigningKeyForAuthenticatedUser) | **Get** /user/ssh_signing_keys/{ssh_signing_key_id} | Get an SSH signing key for the authenticated user
[**UsersList**](UsersApi.md#UsersList) | **Get** /users | List users
[**UsersListBlockedByAuthenticatedUser**](UsersApi.md#UsersListBlockedByAuthenticatedUser) | **Get** /user/blocks | List users blocked by the authenticated user
[**UsersListEmailsForAuthenticatedUser**](UsersApi.md#UsersListEmailsForAuthenticatedUser) | **Get** /user/emails | List email addresses for the authenticated user
[**UsersListFollowedByAuthenticatedUser**](UsersApi.md#UsersListFollowedByAuthenticatedUser) | **Get** /user/following | List the people the authenticated user follows
[**UsersListFollowersForAuthenticatedUser**](UsersApi.md#UsersListFollowersForAuthenticatedUser) | **Get** /user/followers | List followers of the authenticated user
[**UsersListFollowersForUser**](UsersApi.md#UsersListFollowersForUser) | **Get** /users/{username}/followers | List followers of a user
[**UsersListFollowingForUser**](UsersApi.md#UsersListFollowingForUser) | **Get** /users/{username}/following | List the people a user follows
[**UsersListGpgKeysForAuthenticatedUser**](UsersApi.md#UsersListGpgKeysForAuthenticatedUser) | **Get** /user/gpg_keys | List GPG keys for the authenticated user
[**UsersListGpgKeysForUser**](UsersApi.md#UsersListGpgKeysForUser) | **Get** /users/{username}/gpg_keys | List GPG keys for a user
[**UsersListPublicEmailsForAuthenticatedUser**](UsersApi.md#UsersListPublicEmailsForAuthenticatedUser) | **Get** /user/public_emails | List public email addresses for the authenticated user
[**UsersListPublicKeysForUser**](UsersApi.md#UsersListPublicKeysForUser) | **Get** /users/{username}/keys | List public keys for a user
[**UsersListPublicSshKeysForAuthenticatedUser**](UsersApi.md#UsersListPublicSshKeysForAuthenticatedUser) | **Get** /user/keys | List public SSH keys for the authenticated user
[**UsersListSshSigningKeysForAuthenticatedUser**](UsersApi.md#UsersListSshSigningKeysForAuthenticatedUser) | **Get** /user/ssh_signing_keys | List SSH signing keys for the authenticated user
[**UsersListSshSigningKeysForUser**](UsersApi.md#UsersListSshSigningKeysForUser) | **Get** /users/{username}/ssh_signing_keys | List SSH signing keys for a user
[**UsersSetPrimaryEmailVisibilityForAuthenticatedUser**](UsersApi.md#UsersSetPrimaryEmailVisibilityForAuthenticatedUser) | **Patch** /user/email/visibility | Set primary email visibility for the authenticated user
[**UsersUnblock**](UsersApi.md#UsersUnblock) | **Delete** /user/blocks/{username} | Unblock a user
[**UsersUnfollow**](UsersApi.md#UsersUnfollow) | **Delete** /user/following/{username} | Unfollow a user
[**UsersUpdateAuthenticated**](UsersApi.md#UsersUpdateAuthenticated) | **Patch** /user | Update the authenticated user



## UsersAddEmailForAuthenticatedUser

> []Email UsersAddEmailForAuthenticatedUser(ctx).UsersAddEmailForAuthenticatedUserRequest(usersAddEmailForAuthenticatedUserRequest).Execute()

Add an email address for the authenticated user



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
    usersAddEmailForAuthenticatedUserRequest := openapiclient.users_add_email_for_authenticated_user_request{UsersAddEmailForAuthenticatedUserRequestOneOf: openapiclient.NewUsersAddEmailForAuthenticatedUserRequestOneOf([]string{"username@example.com"})} // UsersAddEmailForAuthenticatedUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersAddEmailForAuthenticatedUser(context.Background()).UsersAddEmailForAuthenticatedUserRequest(usersAddEmailForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersAddEmailForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAddEmailForAuthenticatedUser`: []Email
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersAddEmailForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersAddEmailForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersAddEmailForAuthenticatedUserRequest** | [**UsersAddEmailForAuthenticatedUserRequest**](UsersAddEmailForAuthenticatedUserRequest.md) |  | 

### Return type

[**[]Email**](Email.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersBlock

> UsersBlock(ctx, username).Execute()

Block a user



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
    resp, r, err := apiClient.UsersApi.UsersBlock(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersBlockRequest struct via the builder pattern


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


## UsersCheckBlocked

> UsersCheckBlocked(ctx, username).Execute()

Check if a user is blocked by the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersCheckBlocked(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCheckBlocked``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCheckBlockedRequest struct via the builder pattern


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


## UsersCheckFollowingForUser

> UsersCheckFollowingForUser(ctx, username, targetUser).Execute()

Check if a user follows another user



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
    targetUser := "targetUser_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersCheckFollowingForUser(context.Background(), username, targetUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCheckFollowingForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 
**targetUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCheckFollowingForUserRequest struct via the builder pattern


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


## UsersCheckPersonIsFollowedByAuthenticated

> UsersCheckPersonIsFollowedByAuthenticated(ctx, username).Execute()

Check if a person is followed by the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersCheckPersonIsFollowedByAuthenticated(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCheckPersonIsFollowedByAuthenticated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCheckPersonIsFollowedByAuthenticatedRequest struct via the builder pattern


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


## UsersCreateGpgKeyForAuthenticatedUser

> GpgKey UsersCreateGpgKeyForAuthenticatedUser(ctx).UsersCreateGpgKeyForAuthenticatedUserRequest(usersCreateGpgKeyForAuthenticatedUserRequest).Execute()

Create a GPG key for the authenticated user



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
    usersCreateGpgKeyForAuthenticatedUserRequest := *openapiclient.NewUsersCreateGpgKeyForAuthenticatedUserRequest("ArmoredPublicKey_example") // UsersCreateGpgKeyForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersCreateGpgKeyForAuthenticatedUser(context.Background()).UsersCreateGpgKeyForAuthenticatedUserRequest(usersCreateGpgKeyForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCreateGpgKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateGpgKeyForAuthenticatedUser`: GpgKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersCreateGpgKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateGpgKeyForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersCreateGpgKeyForAuthenticatedUserRequest** | [**UsersCreateGpgKeyForAuthenticatedUserRequest**](UsersCreateGpgKeyForAuthenticatedUserRequest.md) |  | 

### Return type

[**GpgKey**](GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCreatePublicSshKeyForAuthenticatedUser

> Key UsersCreatePublicSshKeyForAuthenticatedUser(ctx).UsersCreatePublicSshKeyForAuthenticatedUserRequest(usersCreatePublicSshKeyForAuthenticatedUserRequest).Execute()

Create a public SSH key for the authenticated user



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
    usersCreatePublicSshKeyForAuthenticatedUserRequest := *openapiclient.NewUsersCreatePublicSshKeyForAuthenticatedUserRequest("Key_example") // UsersCreatePublicSshKeyForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersCreatePublicSshKeyForAuthenticatedUser(context.Background()).UsersCreatePublicSshKeyForAuthenticatedUserRequest(usersCreatePublicSshKeyForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCreatePublicSshKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreatePublicSshKeyForAuthenticatedUser`: Key
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersCreatePublicSshKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreatePublicSshKeyForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersCreatePublicSshKeyForAuthenticatedUserRequest** | [**UsersCreatePublicSshKeyForAuthenticatedUserRequest**](UsersCreatePublicSshKeyForAuthenticatedUserRequest.md) |  | 

### Return type

[**Key**](Key.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCreateSshSigningKeyForAuthenticatedUser

> SshSigningKey UsersCreateSshSigningKeyForAuthenticatedUser(ctx).UsersCreateSshSigningKeyForAuthenticatedUserRequest(usersCreateSshSigningKeyForAuthenticatedUserRequest).Execute()

Create a SSH signing key for the authenticated user



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
    usersCreateSshSigningKeyForAuthenticatedUserRequest := *openapiclient.NewUsersCreateSshSigningKeyForAuthenticatedUserRequest("Key_example") // UsersCreateSshSigningKeyForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersCreateSshSigningKeyForAuthenticatedUser(context.Background()).UsersCreateSshSigningKeyForAuthenticatedUserRequest(usersCreateSshSigningKeyForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCreateSshSigningKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateSshSigningKeyForAuthenticatedUser`: SshSigningKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersCreateSshSigningKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateSshSigningKeyForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersCreateSshSigningKeyForAuthenticatedUserRequest** | [**UsersCreateSshSigningKeyForAuthenticatedUserRequest**](UsersCreateSshSigningKeyForAuthenticatedUserRequest.md) |  | 

### Return type

[**SshSigningKey**](SshSigningKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteEmailForAuthenticatedUser

> UsersDeleteEmailForAuthenticatedUser(ctx).UsersDeleteEmailForAuthenticatedUserRequest(usersDeleteEmailForAuthenticatedUserRequest).Execute()

Delete an email address for the authenticated user



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
    usersDeleteEmailForAuthenticatedUserRequest := openapiclient.users_delete_email_for_authenticated_user_request{UsersDeleteEmailForAuthenticatedUserRequestOneOf: openapiclient.NewUsersDeleteEmailForAuthenticatedUserRequestOneOf([]string{"username@example.com"})} // UsersDeleteEmailForAuthenticatedUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersDeleteEmailForAuthenticatedUser(context.Background()).UsersDeleteEmailForAuthenticatedUserRequest(usersDeleteEmailForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDeleteEmailForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteEmailForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersDeleteEmailForAuthenticatedUserRequest** | [**UsersDeleteEmailForAuthenticatedUserRequest**](UsersDeleteEmailForAuthenticatedUserRequest.md) |  | 

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


## UsersDeleteGpgKeyForAuthenticatedUser

> UsersDeleteGpgKeyForAuthenticatedUser(ctx, gpgKeyId).Execute()

Delete a GPG key for the authenticated user



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
    gpgKeyId := int32(56) // int32 | The unique identifier of the GPG key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersDeleteGpgKeyForAuthenticatedUser(context.Background(), gpgKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDeleteGpgKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpgKeyId** | **int32** | The unique identifier of the GPG key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteGpgKeyForAuthenticatedUserRequest struct via the builder pattern


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


## UsersDeletePublicSshKeyForAuthenticatedUser

> UsersDeletePublicSshKeyForAuthenticatedUser(ctx, keyId).Execute()

Delete a public SSH key for the authenticated user



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
    keyId := int32(56) // int32 | The unique identifier of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersDeletePublicSshKeyForAuthenticatedUser(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDeletePublicSshKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int32** | The unique identifier of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeletePublicSshKeyForAuthenticatedUserRequest struct via the builder pattern


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


## UsersDeleteSshSigningKeyForAuthenticatedUser

> UsersDeleteSshSigningKeyForAuthenticatedUser(ctx, sshSigningKeyId).Execute()

Delete an SSH signing key for the authenticated user



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
    sshSigningKeyId := int32(56) // int32 | The unique identifier of the SSH signing key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersDeleteSshSigningKeyForAuthenticatedUser(context.Background(), sshSigningKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDeleteSshSigningKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshSigningKeyId** | **int32** | The unique identifier of the SSH signing key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteSshSigningKeyForAuthenticatedUserRequest struct via the builder pattern


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


## UsersFollow

> UsersFollow(ctx, username).Execute()

Follow a user



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
    resp, r, err := apiClient.UsersApi.UsersFollow(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersFollow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersFollowRequest struct via the builder pattern


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


## UsersGetAuthenticated

> UsersGetAuthenticated200Response UsersGetAuthenticated(ctx).Execute()

Get the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersGetAuthenticated(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetAuthenticated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetAuthenticated`: UsersGetAuthenticated200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetAuthenticated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetAuthenticatedRequest struct via the builder pattern


### Return type

[**UsersGetAuthenticated200Response**](UsersGetAuthenticated200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetByUsername

> UsersGetAuthenticated200Response UsersGetByUsername(ctx, username).Execute()

Get a user



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
    resp, r, err := apiClient.UsersApi.UsersGetByUsername(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetByUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetByUsername`: UsersGetAuthenticated200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetByUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetByUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsersGetAuthenticated200Response**](UsersGetAuthenticated200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetContextForUser

> Hovercard UsersGetContextForUser(ctx, username).SubjectType(subjectType).SubjectId(subjectId).Execute()

Get contextual information for a user



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
    subjectType := "subjectType_example" // string | Identifies which additional information you'd like to receive about the person's hovercard. Can be `organization`, `repository`, `issue`, `pull_request`. **Required** when using `subject_id`. (optional)
    subjectId := "subjectId_example" // string | Uses the ID for the `subject_type` you specified. **Required** when using `subject_type`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGetContextForUser(context.Background(), username).SubjectType(subjectType).SubjectId(subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetContextForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetContextForUser`: Hovercard
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetContextForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetContextForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectType** | **string** | Identifies which additional information you&#39;d like to receive about the person&#39;s hovercard. Can be &#x60;organization&#x60;, &#x60;repository&#x60;, &#x60;issue&#x60;, &#x60;pull_request&#x60;. **Required** when using &#x60;subject_id&#x60;. | 
 **subjectId** | **string** | Uses the ID for the &#x60;subject_type&#x60; you specified. **Required** when using &#x60;subject_type&#x60;. | 

### Return type

[**Hovercard**](Hovercard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetGpgKeyForAuthenticatedUser

> GpgKey UsersGetGpgKeyForAuthenticatedUser(ctx, gpgKeyId).Execute()

Get a GPG key for the authenticated user



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
    gpgKeyId := int32(56) // int32 | The unique identifier of the GPG key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGetGpgKeyForAuthenticatedUser(context.Background(), gpgKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetGpgKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetGpgKeyForAuthenticatedUser`: GpgKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetGpgKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpgKeyId** | **int32** | The unique identifier of the GPG key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetGpgKeyForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GpgKey**](GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetPublicSshKeyForAuthenticatedUser

> Key UsersGetPublicSshKeyForAuthenticatedUser(ctx, keyId).Execute()

Get a public SSH key for the authenticated user



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
    keyId := int32(56) // int32 | The unique identifier of the key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGetPublicSshKeyForAuthenticatedUser(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetPublicSshKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPublicSshKeyForAuthenticatedUser`: Key
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetPublicSshKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int32** | The unique identifier of the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPublicSshKeyForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Key**](Key.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetSshSigningKeyForAuthenticatedUser

> SshSigningKey UsersGetSshSigningKeyForAuthenticatedUser(ctx, sshSigningKeyId).Execute()

Get an SSH signing key for the authenticated user



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
    sshSigningKeyId := int32(56) // int32 | The unique identifier of the SSH signing key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGetSshSigningKeyForAuthenticatedUser(context.Background(), sshSigningKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetSshSigningKeyForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetSshSigningKeyForAuthenticatedUser`: SshSigningKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetSshSigningKeyForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshSigningKeyId** | **int32** | The unique identifier of the SSH signing key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetSshSigningKeyForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SshSigningKey**](SshSigningKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersList

> []SimpleUser UsersList(ctx).Since(since).PerPage(perPage).Execute()

List users



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
    since := int32(56) // int32 | A user ID. Only return users with an ID greater than this ID. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersList(context.Background()).Since(since).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersList`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int32** | A user ID. Only return users with an ID greater than this ID. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

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


## UsersListBlockedByAuthenticatedUser

> []SimpleUser UsersListBlockedByAuthenticatedUser(ctx).Execute()

List users blocked by the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListBlockedByAuthenticatedUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListBlockedByAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListBlockedByAuthenticatedUser`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListBlockedByAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListBlockedByAuthenticatedUserRequest struct via the builder pattern


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


## UsersListEmailsForAuthenticatedUser

> []Email UsersListEmailsForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List email addresses for the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListEmailsForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListEmailsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListEmailsForAuthenticatedUser`: []Email
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListEmailsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListEmailsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Email**](Email.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListFollowedByAuthenticatedUser

> []SimpleUser UsersListFollowedByAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List the people the authenticated user follows



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
    resp, r, err := apiClient.UsersApi.UsersListFollowedByAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListFollowedByAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListFollowedByAuthenticatedUser`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListFollowedByAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListFollowedByAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## UsersListFollowersForAuthenticatedUser

> []SimpleUser UsersListFollowersForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List followers of the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListFollowersForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListFollowersForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListFollowersForAuthenticatedUser`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListFollowersForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListFollowersForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## UsersListFollowersForUser

> []SimpleUser UsersListFollowersForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List followers of a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersListFollowersForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListFollowersForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListFollowersForUser`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListFollowersForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListFollowersForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## UsersListFollowingForUser

> []SimpleUser UsersListFollowingForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List the people a user follows



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersListFollowingForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListFollowingForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListFollowingForUser`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListFollowingForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListFollowingForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## UsersListGpgKeysForAuthenticatedUser

> []GpgKey UsersListGpgKeysForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List GPG keys for the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListGpgKeysForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListGpgKeysForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListGpgKeysForAuthenticatedUser`: []GpgKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListGpgKeysForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListGpgKeysForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]GpgKey**](GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListGpgKeysForUser

> []GpgKey UsersListGpgKeysForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List GPG keys for a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersListGpgKeysForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListGpgKeysForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListGpgKeysForUser`: []GpgKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListGpgKeysForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListGpgKeysForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]GpgKey**](GpgKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListPublicEmailsForAuthenticatedUser

> []Email UsersListPublicEmailsForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List public email addresses for the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListPublicEmailsForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListPublicEmailsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListPublicEmailsForAuthenticatedUser`: []Email
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListPublicEmailsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListPublicEmailsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Email**](Email.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListPublicKeysForUser

> []KeySimple UsersListPublicKeysForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List public keys for a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersListPublicKeysForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListPublicKeysForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListPublicKeysForUser`: []KeySimple
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListPublicKeysForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListPublicKeysForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]KeySimple**](KeySimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListPublicSshKeysForAuthenticatedUser

> []Key UsersListPublicSshKeysForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List public SSH keys for the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListPublicSshKeysForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListPublicSshKeysForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListPublicSshKeysForAuthenticatedUser`: []Key
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListPublicSshKeysForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListPublicSshKeysForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Key**](Key.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListSshSigningKeysForAuthenticatedUser

> []SshSigningKey UsersListSshSigningKeysForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List SSH signing keys for the authenticated user



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
    resp, r, err := apiClient.UsersApi.UsersListSshSigningKeysForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListSshSigningKeysForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListSshSigningKeysForAuthenticatedUser`: []SshSigningKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListSshSigningKeysForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListSshSigningKeysForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]SshSigningKey**](SshSigningKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListSshSigningKeysForUser

> []SshSigningKey UsersListSshSigningKeysForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List SSH signing keys for a user



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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersListSshSigningKeysForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersListSshSigningKeysForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListSshSigningKeysForUser`: []SshSigningKey
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersListSshSigningKeysForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListSshSigningKeysForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]SshSigningKey**](SshSigningKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSetPrimaryEmailVisibilityForAuthenticatedUser

> []Email UsersSetPrimaryEmailVisibilityForAuthenticatedUser(ctx).UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest(usersSetPrimaryEmailVisibilityForAuthenticatedUserRequest).Execute()

Set primary email visibility for the authenticated user



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
    usersSetPrimaryEmailVisibilityForAuthenticatedUserRequest := *openapiclient.NewUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest("Visibility_example") // UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersSetPrimaryEmailVisibilityForAuthenticatedUser(context.Background()).UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest(usersSetPrimaryEmailVisibilityForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersSetPrimaryEmailVisibilityForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSetPrimaryEmailVisibilityForAuthenticatedUser`: []Email
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersSetPrimaryEmailVisibilityForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersSetPrimaryEmailVisibilityForAuthenticatedUserRequest** | [**UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest**](UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest.md) |  | 

### Return type

[**[]Email**](Email.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUnblock

> UsersUnblock(ctx, username).Execute()

Unblock a user



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
    resp, r, err := apiClient.UsersApi.UsersUnblock(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUnblock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUnblockRequest struct via the builder pattern


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


## UsersUnfollow

> UsersUnfollow(ctx, username).Execute()

Unfollow a user



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
    resp, r, err := apiClient.UsersApi.UsersUnfollow(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUnfollow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUnfollowRequest struct via the builder pattern


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


## UsersUpdateAuthenticated

> PrivateUser UsersUpdateAuthenticated(ctx).UsersUpdateAuthenticatedRequest(usersUpdateAuthenticatedRequest).Execute()

Update the authenticated user



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
    usersUpdateAuthenticatedRequest := *openapiclient.NewUsersUpdateAuthenticatedRequest() // UsersUpdateAuthenticatedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUpdateAuthenticated(context.Background()).UsersUpdateAuthenticatedRequest(usersUpdateAuthenticatedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUpdateAuthenticated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUpdateAuthenticated`: PrivateUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUpdateAuthenticated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateAuthenticatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersUpdateAuthenticatedRequest** | [**UsersUpdateAuthenticatedRequest**](UsersUpdateAuthenticatedRequest.md) |  | 

### Return type

[**PrivateUser**](PrivateUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

