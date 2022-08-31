# \OrgsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgsAddSecurityManagerTeam**](OrgsApi.md#OrgsAddSecurityManagerTeam) | **Put** /orgs/{org}/security-managers/teams/{team_slug} | Add a security manager team
[**OrgsBlockUser**](OrgsApi.md#OrgsBlockUser) | **Put** /orgs/{org}/blocks/{username} | Block a user from an organization
[**OrgsCancelInvitation**](OrgsApi.md#OrgsCancelInvitation) | **Delete** /orgs/{org}/invitations/{invitation_id} | Cancel an organization invitation
[**OrgsCheckBlockedUser**](OrgsApi.md#OrgsCheckBlockedUser) | **Get** /orgs/{org}/blocks/{username} | Check if a user is blocked by an organization
[**OrgsCheckMembershipForUser**](OrgsApi.md#OrgsCheckMembershipForUser) | **Get** /orgs/{org}/members/{username} | Check organization membership for a user
[**OrgsCheckPublicMembershipForUser**](OrgsApi.md#OrgsCheckPublicMembershipForUser) | **Get** /orgs/{org}/public_members/{username} | Check public organization membership for a user
[**OrgsConvertMemberToOutsideCollaborator**](OrgsApi.md#OrgsConvertMemberToOutsideCollaborator) | **Put** /orgs/{org}/outside_collaborators/{username} | Convert an organization member to outside collaborator
[**OrgsCreateInvitation**](OrgsApi.md#OrgsCreateInvitation) | **Post** /orgs/{org}/invitations | Create an organization invitation
[**OrgsCreateWebhook**](OrgsApi.md#OrgsCreateWebhook) | **Post** /orgs/{org}/hooks | Create an organization webhook
[**OrgsDeleteWebhook**](OrgsApi.md#OrgsDeleteWebhook) | **Delete** /orgs/{org}/hooks/{hook_id} | Delete an organization webhook
[**OrgsEnableOrDisableSecurityProductOnAllOrgRepos**](OrgsApi.md#OrgsEnableOrDisableSecurityProductOnAllOrgRepos) | **Post** /orgs/{org}/{security_product}/{enablement} | Enable or disable a security feature for an organization
[**OrgsGet**](OrgsApi.md#OrgsGet) | **Get** /orgs/{org} | Get an organization
[**OrgsGetAuditLog**](OrgsApi.md#OrgsGetAuditLog) | **Get** /orgs/{org}/audit-log | Get the audit log for an organization
[**OrgsGetMembershipForAuthenticatedUser**](OrgsApi.md#OrgsGetMembershipForAuthenticatedUser) | **Get** /user/memberships/orgs/{org} | Get an organization membership for the authenticated user
[**OrgsGetMembershipForUser**](OrgsApi.md#OrgsGetMembershipForUser) | **Get** /orgs/{org}/memberships/{username} | Get organization membership for a user
[**OrgsGetWebhook**](OrgsApi.md#OrgsGetWebhook) | **Get** /orgs/{org}/hooks/{hook_id} | Get an organization webhook
[**OrgsGetWebhookConfigForOrg**](OrgsApi.md#OrgsGetWebhookConfigForOrg) | **Get** /orgs/{org}/hooks/{hook_id}/config | Get a webhook configuration for an organization
[**OrgsGetWebhookDelivery**](OrgsApi.md#OrgsGetWebhookDelivery) | **Get** /orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id} | Get a webhook delivery for an organization webhook
[**OrgsList**](OrgsApi.md#OrgsList) | **Get** /organizations | List organizations
[**OrgsListAppInstallations**](OrgsApi.md#OrgsListAppInstallations) | **Get** /orgs/{org}/installations | List app installations for an organization
[**OrgsListBlockedUsers**](OrgsApi.md#OrgsListBlockedUsers) | **Get** /orgs/{org}/blocks | List users blocked by an organization
[**OrgsListCustomRoles**](OrgsApi.md#OrgsListCustomRoles) | **Get** /organizations/{organization_id}/custom_roles | List custom repository roles in an organization
[**OrgsListFailedInvitations**](OrgsApi.md#OrgsListFailedInvitations) | **Get** /orgs/{org}/failed_invitations | List failed organization invitations
[**OrgsListForAuthenticatedUser**](OrgsApi.md#OrgsListForAuthenticatedUser) | **Get** /user/orgs | List organizations for the authenticated user
[**OrgsListForUser**](OrgsApi.md#OrgsListForUser) | **Get** /users/{username}/orgs | List organizations for a user
[**OrgsListInvitationTeams**](OrgsApi.md#OrgsListInvitationTeams) | **Get** /orgs/{org}/invitations/{invitation_id}/teams | List organization invitation teams
[**OrgsListMembers**](OrgsApi.md#OrgsListMembers) | **Get** /orgs/{org}/members | List organization members
[**OrgsListMembershipsForAuthenticatedUser**](OrgsApi.md#OrgsListMembershipsForAuthenticatedUser) | **Get** /user/memberships/orgs | List organization memberships for the authenticated user
[**OrgsListOutsideCollaborators**](OrgsApi.md#OrgsListOutsideCollaborators) | **Get** /orgs/{org}/outside_collaborators | List outside collaborators for an organization
[**OrgsListPendingInvitations**](OrgsApi.md#OrgsListPendingInvitations) | **Get** /orgs/{org}/invitations | List pending organization invitations
[**OrgsListPublicMembers**](OrgsApi.md#OrgsListPublicMembers) | **Get** /orgs/{org}/public_members | List public organization members
[**OrgsListSamlSsoAuthorizations**](OrgsApi.md#OrgsListSamlSsoAuthorizations) | **Get** /orgs/{org}/credential-authorizations | List SAML SSO authorizations for an organization
[**OrgsListSecurityManagerTeams**](OrgsApi.md#OrgsListSecurityManagerTeams) | **Get** /orgs/{org}/security-managers | List security manager teams
[**OrgsListWebhookDeliveries**](OrgsApi.md#OrgsListWebhookDeliveries) | **Get** /orgs/{org}/hooks/{hook_id}/deliveries | List deliveries for an organization webhook
[**OrgsListWebhooks**](OrgsApi.md#OrgsListWebhooks) | **Get** /orgs/{org}/hooks | List organization webhooks
[**OrgsPingWebhook**](OrgsApi.md#OrgsPingWebhook) | **Post** /orgs/{org}/hooks/{hook_id}/pings | Ping an organization webhook
[**OrgsRedeliverWebhookDelivery**](OrgsApi.md#OrgsRedeliverWebhookDelivery) | **Post** /orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}/attempts | Redeliver a delivery for an organization webhook
[**OrgsRemoveMember**](OrgsApi.md#OrgsRemoveMember) | **Delete** /orgs/{org}/members/{username} | Remove an organization member
[**OrgsRemoveMembershipForUser**](OrgsApi.md#OrgsRemoveMembershipForUser) | **Delete** /orgs/{org}/memberships/{username} | Remove organization membership for a user
[**OrgsRemoveOutsideCollaborator**](OrgsApi.md#OrgsRemoveOutsideCollaborator) | **Delete** /orgs/{org}/outside_collaborators/{username} | Remove outside collaborator from an organization
[**OrgsRemovePublicMembershipForAuthenticatedUser**](OrgsApi.md#OrgsRemovePublicMembershipForAuthenticatedUser) | **Delete** /orgs/{org}/public_members/{username} | Remove public organization membership for the authenticated user
[**OrgsRemoveSamlSsoAuthorization**](OrgsApi.md#OrgsRemoveSamlSsoAuthorization) | **Delete** /orgs/{org}/credential-authorizations/{credential_id} | Remove a SAML SSO authorization for an organization
[**OrgsRemoveSecurityManagerTeam**](OrgsApi.md#OrgsRemoveSecurityManagerTeam) | **Delete** /orgs/{org}/security-managers/teams/{team_slug} | Remove a security manager team
[**OrgsSetMembershipForUser**](OrgsApi.md#OrgsSetMembershipForUser) | **Put** /orgs/{org}/memberships/{username} | Set organization membership for a user
[**OrgsSetPublicMembershipForAuthenticatedUser**](OrgsApi.md#OrgsSetPublicMembershipForAuthenticatedUser) | **Put** /orgs/{org}/public_members/{username} | Set public organization membership for the authenticated user
[**OrgsUnblockUser**](OrgsApi.md#OrgsUnblockUser) | **Delete** /orgs/{org}/blocks/{username} | Unblock a user from an organization
[**OrgsUpdate**](OrgsApi.md#OrgsUpdate) | **Patch** /orgs/{org} | Update an organization
[**OrgsUpdateMembershipForAuthenticatedUser**](OrgsApi.md#OrgsUpdateMembershipForAuthenticatedUser) | **Patch** /user/memberships/orgs/{org} | Update an organization membership for the authenticated user
[**OrgsUpdateWebhook**](OrgsApi.md#OrgsUpdateWebhook) | **Patch** /orgs/{org}/hooks/{hook_id} | Update an organization webhook
[**OrgsUpdateWebhookConfigForOrg**](OrgsApi.md#OrgsUpdateWebhookConfigForOrg) | **Patch** /orgs/{org}/hooks/{hook_id}/config | Update a webhook configuration for an organization



## OrgsAddSecurityManagerTeam

> OrgsAddSecurityManagerTeam(ctx, org, teamSlug).Execute()

Add a security manager team



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsAddSecurityManagerTeam(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsAddSecurityManagerTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsAddSecurityManagerTeamRequest struct via the builder pattern


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


## OrgsBlockUser

> OrgsBlockUser(ctx, org, username).Execute()

Block a user from an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsBlockUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsBlockUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsBlockUserRequest struct via the builder pattern


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


## OrgsCancelInvitation

> OrgsCancelInvitation(ctx, org, invitationId).Execute()

Cancel an organization invitation



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
    invitationId := int32(56) // int32 | The unique identifier of the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsCancelInvitation(context.Background(), org, invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsCancelInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**invitationId** | **int32** | The unique identifier of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsCancelInvitationRequest struct via the builder pattern


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


## OrgsCheckBlockedUser

> OrgsCheckBlockedUser(ctx, org, username).Execute()

Check if a user is blocked by an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsCheckBlockedUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsCheckBlockedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsCheckBlockedUserRequest struct via the builder pattern


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


## OrgsCheckMembershipForUser

> OrgsCheckMembershipForUser(ctx, org, username).Execute()

Check organization membership for a user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsCheckMembershipForUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsCheckMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsCheckMembershipForUserRequest struct via the builder pattern


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


## OrgsCheckPublicMembershipForUser

> OrgsCheckPublicMembershipForUser(ctx, org, username).Execute()

Check public organization membership for a user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsCheckPublicMembershipForUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsCheckPublicMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsCheckPublicMembershipForUserRequest struct via the builder pattern


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


## OrgsConvertMemberToOutsideCollaborator

> map[string]interface{} OrgsConvertMemberToOutsideCollaborator(ctx, org, username).OrgsConvertMemberToOutsideCollaboratorRequest(orgsConvertMemberToOutsideCollaboratorRequest).Execute()

Convert an organization member to outside collaborator



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
    orgsConvertMemberToOutsideCollaboratorRequest := *openapiclient.NewOrgsConvertMemberToOutsideCollaboratorRequest() // OrgsConvertMemberToOutsideCollaboratorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsConvertMemberToOutsideCollaborator(context.Background(), org, username).OrgsConvertMemberToOutsideCollaboratorRequest(orgsConvertMemberToOutsideCollaboratorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsConvertMemberToOutsideCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsConvertMemberToOutsideCollaborator`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsConvertMemberToOutsideCollaborator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsConvertMemberToOutsideCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgsConvertMemberToOutsideCollaboratorRequest** | [**OrgsConvertMemberToOutsideCollaboratorRequest**](OrgsConvertMemberToOutsideCollaboratorRequest.md) |  | 

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


## OrgsCreateInvitation

> OrganizationInvitation OrgsCreateInvitation(ctx, org).OrgsCreateInvitationRequest(orgsCreateInvitationRequest).Execute()

Create an organization invitation



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
    orgsCreateInvitationRequest := *openapiclient.NewOrgsCreateInvitationRequest() // OrgsCreateInvitationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsCreateInvitation(context.Background(), org).OrgsCreateInvitationRequest(orgsCreateInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsCreateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsCreateInvitation`: OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsCreateInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgsCreateInvitationRequest** | [**OrgsCreateInvitationRequest**](OrgsCreateInvitationRequest.md) |  | 

### Return type

[**OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsCreateWebhook

> OrgHook OrgsCreateWebhook(ctx, org).OrgsCreateWebhookRequest(orgsCreateWebhookRequest).Execute()

Create an organization webhook



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
    orgsCreateWebhookRequest := *openapiclient.NewOrgsCreateWebhookRequest("Name_example", *openapiclient.NewOrgsCreateWebhookRequestConfig("https://example.com/webhook")) // OrgsCreateWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsCreateWebhook(context.Background(), org).OrgsCreateWebhookRequest(orgsCreateWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsCreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsCreateWebhook`: OrgHook
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsCreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgsCreateWebhookRequest** | [**OrgsCreateWebhookRequest**](OrgsCreateWebhookRequest.md) |  | 

### Return type

[**OrgHook**](OrgHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsDeleteWebhook

> OrgsDeleteWebhook(ctx, org, hookId).Execute()

Delete an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsDeleteWebhook(context.Background(), org, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsDeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsDeleteWebhookRequest struct via the builder pattern


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


## OrgsEnableOrDisableSecurityProductOnAllOrgRepos

> OrgsEnableOrDisableSecurityProductOnAllOrgRepos(ctx, org, securityProduct, enablement).Execute()

Enable or disable a security feature for an organization



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
    securityProduct := "securityProduct_example" // string | The security feature to enable or disable.
    enablement := "enablement_example" // string | The action to take.  `enable_all` means to enable the specified security feature for all repositories in the organization. `disable_all` means to disable the specified security feature for all repositories in the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsEnableOrDisableSecurityProductOnAllOrgRepos(context.Background(), org, securityProduct, enablement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsEnableOrDisableSecurityProductOnAllOrgRepos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**securityProduct** | **string** | The security feature to enable or disable. | 
**enablement** | **string** | The action to take.  &#x60;enable_all&#x60; means to enable the specified security feature for all repositories in the organization. &#x60;disable_all&#x60; means to disable the specified security feature for all repositories in the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsEnableOrDisableSecurityProductOnAllOrgReposRequest struct via the builder pattern


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


## OrgsGet

> OrganizationFull OrgsGet(ctx, org).Execute()

Get an organization



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
    resp, r, err := apiClient.OrgsApi.OrgsGet(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGet`: OrganizationFull
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationFull**](OrganizationFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsGetAuditLog

> []AuditLogEvent OrgsGetAuditLog(ctx, org).Phrase(phrase).Include(include).After(after).Before(before).Order(order).PerPage(perPage).Execute()

Get the audit log for an organization



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
    phrase := "phrase_example" // string | A search phrase. For more information, see [Searching the audit log](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization#searching-the-audit-log). (optional)
    include := "include_example" // string | The event types to include:  - `web` - returns web (non-Git) events. - `git` - returns Git events. - `all` - returns both web and Git events.  The default is `web`. (optional)
    after := "after_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. (optional)
    before := "before_example" // string | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. (optional)
    order := "order_example" // string | The order of audit log events. To list newest events first, specify `desc`. To list oldest events first, specify `asc`.  The default is `desc`. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsGetAuditLog(context.Background(), org).Phrase(phrase).Include(include).After(after).Before(before).Order(order).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGetAuditLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGetAuditLog`: []AuditLogEvent
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGetAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phrase** | **string** | A search phrase. For more information, see [Searching the audit log](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/reviewing-the-audit-log-for-your-organization#searching-the-audit-log). | 
 **include** | **string** | The event types to include:  - &#x60;web&#x60; - returns web (non-Git) events. - &#x60;git&#x60; - returns Git events. - &#x60;all&#x60; - returns both web and Git events.  The default is &#x60;web&#x60;. | 
 **after** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events after this cursor. | 
 **before** | **string** | A cursor, as given in the [Link header](https://docs.github.com/rest/overview/resources-in-the-rest-api#link-header). If specified, the query only searches for events before this cursor. | 
 **order** | **string** | The order of audit log events. To list newest events first, specify &#x60;desc&#x60;. To list oldest events first, specify &#x60;asc&#x60;.  The default is &#x60;desc&#x60;. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**[]AuditLogEvent**](AuditLogEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsGetMembershipForAuthenticatedUser

> OrgMembership OrgsGetMembershipForAuthenticatedUser(ctx, org).Execute()

Get an organization membership for the authenticated user



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
    resp, r, err := apiClient.OrgsApi.OrgsGetMembershipForAuthenticatedUser(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGetMembershipForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGetMembershipForAuthenticatedUser`: OrgMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGetMembershipForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetMembershipForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsGetMembershipForUser

> OrgMembership OrgsGetMembershipForUser(ctx, org, username).Execute()

Get organization membership for a user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsGetMembershipForUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGetMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGetMembershipForUser`: OrgMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGetMembershipForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetMembershipForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsGetWebhook

> OrgHook OrgsGetWebhook(ctx, org, hookId).Execute()

Get an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsGetWebhook(context.Background(), org, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGetWebhook`: OrgHook
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgHook**](OrgHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsGetWebhookConfigForOrg

> WebhookConfig OrgsGetWebhookConfigForOrg(ctx, org, hookId).Execute()

Get a webhook configuration for an organization



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsGetWebhookConfigForOrg(context.Background(), org, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGetWebhookConfigForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGetWebhookConfigForOrg`: WebhookConfig
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGetWebhookConfigForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetWebhookConfigForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## OrgsGetWebhookDelivery

> HookDelivery OrgsGetWebhookDelivery(ctx, org, hookId, deliveryId).Execute()

Get a webhook delivery for an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    deliveryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsGetWebhookDelivery(context.Background(), org, hookId, deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsGetWebhookDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsGetWebhookDelivery`: HookDelivery
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsGetWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsGetWebhookDeliveryRequest struct via the builder pattern


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


## OrgsList

> []OrganizationSimple OrgsList(ctx).Since(since).PerPage(perPage).Execute()

List organizations



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
    since := int32(56) // int32 | An organization ID. Only return organizations with an ID greater than this ID. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsList(context.Background()).Since(since).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsList`: []OrganizationSimple
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int32** | An organization ID. Only return organizations with an ID greater than this ID. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]

### Return type

[**[]OrganizationSimple**](OrganizationSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListAppInstallations

> OrgsListAppInstallations200Response OrgsListAppInstallations(ctx, org).PerPage(perPage).Page(page).Execute()

List app installations for an organization



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
    resp, r, err := apiClient.OrgsApi.OrgsListAppInstallations(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListAppInstallations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListAppInstallations`: OrgsListAppInstallations200Response
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListAppInstallations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListAppInstallationsRequest struct via the builder pattern


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


## OrgsListBlockedUsers

> []SimpleUser OrgsListBlockedUsers(ctx, org).Execute()

List users blocked by an organization



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
    resp, r, err := apiClient.OrgsApi.OrgsListBlockedUsers(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListBlockedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListBlockedUsers`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListBlockedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListBlockedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## OrgsListCustomRoles

> OrgsListCustomRoles200Response OrgsListCustomRoles(ctx, organizationId).Execute()

List custom repository roles in an organization



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
    organizationId := "organizationId_example" // string | The unique identifier of the organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListCustomRoles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListCustomRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListCustomRoles`: OrgsListCustomRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListCustomRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique identifier of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListCustomRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgsListCustomRoles200Response**](OrgsListCustomRoles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListFailedInvitations

> []OrganizationInvitation OrgsListFailedInvitations(ctx, org).PerPage(perPage).Page(page).Execute()

List failed organization invitations



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
    resp, r, err := apiClient.OrgsApi.OrgsListFailedInvitations(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListFailedInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListFailedInvitations`: []OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListFailedInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListFailedInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListForAuthenticatedUser

> []OrganizationSimple OrgsListForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List organizations for the authenticated user



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
    resp, r, err := apiClient.OrgsApi.OrgsListForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListForAuthenticatedUser`: []OrganizationSimple
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]OrganizationSimple**](OrganizationSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListForUser

> []OrganizationSimple OrgsListForUser(ctx, username).PerPage(perPage).Page(page).Execute()

List organizations for a user



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
    resp, r, err := apiClient.OrgsApi.OrgsListForUser(context.Background(), username).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListForUser`: []OrganizationSimple
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]OrganizationSimple**](OrganizationSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListInvitationTeams

> []Team OrgsListInvitationTeams(ctx, org, invitationId).PerPage(perPage).Page(page).Execute()

List organization invitation teams



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
    invitationId := int32(56) // int32 | The unique identifier of the invitation.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListInvitationTeams(context.Background(), org, invitationId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListInvitationTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListInvitationTeams`: []Team
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListInvitationTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**invitationId** | **int32** | The unique identifier of the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListInvitationTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListMembers

> []SimpleUser OrgsListMembers(ctx, org).Filter(filter).Role(role).PerPage(perPage).Page(page).Execute()

List organization members



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
    filter := "filter_example" // string | Filter members returned in the list. `2fa_disabled` means that only members without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned. This options is only available for organization owners. (optional) (default to "all")
    role := "role_example" // string | Filter members returned by their role. (optional) (default to "all")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListMembers(context.Background(), org).Filter(filter).Role(role).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListMembers`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter members returned in the list. &#x60;2fa_disabled&#x60; means that only members without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned. This options is only available for organization owners. | [default to &quot;all&quot;]
 **role** | **string** | Filter members returned by their role. | [default to &quot;all&quot;]
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


## OrgsListMembershipsForAuthenticatedUser

> []OrgMembership OrgsListMembershipsForAuthenticatedUser(ctx).State(state).PerPage(perPage).Page(page).Execute()

List organization memberships for the authenticated user



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
    state := "state_example" // string | Indicates the state of the memberships to return. Can be either `active` or `pending`. If not specified, the API returns both active and pending memberships. (optional)
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListMembershipsForAuthenticatedUser(context.Background()).State(state).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListMembershipsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListMembershipsForAuthenticatedUser`: []OrgMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListMembershipsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListMembershipsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Indicates the state of the memberships to return. Can be either &#x60;active&#x60; or &#x60;pending&#x60;. If not specified, the API returns both active and pending memberships. | 
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]OrgMembership**](OrgMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListOutsideCollaborators

> []SimpleUser OrgsListOutsideCollaborators(ctx, org).Filter(filter).PerPage(perPage).Page(page).Execute()

List outside collaborators for an organization



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
    filter := "filter_example" // string | Filter the list of outside collaborators. `2fa_disabled` means that only outside collaborators without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned. (optional) (default to "all")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListOutsideCollaborators(context.Background(), org).Filter(filter).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListOutsideCollaborators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListOutsideCollaborators`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListOutsideCollaborators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListOutsideCollaboratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter the list of outside collaborators. &#x60;2fa_disabled&#x60; means that only outside collaborators without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned. | [default to &quot;all&quot;]
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


## OrgsListPendingInvitations

> []OrganizationInvitation OrgsListPendingInvitations(ctx, org).PerPage(perPage).Page(page).Execute()

List pending organization invitations



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
    resp, r, err := apiClient.OrgsApi.OrgsListPendingInvitations(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListPendingInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListPendingInvitations`: []OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListPendingInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListPendingInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListPublicMembers

> []SimpleUser OrgsListPublicMembers(ctx, org).PerPage(perPage).Page(page).Execute()

List public organization members



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
    resp, r, err := apiClient.OrgsApi.OrgsListPublicMembers(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListPublicMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListPublicMembers`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListPublicMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListPublicMembersRequest struct via the builder pattern


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


## OrgsListSamlSsoAuthorizations

> []CredentialAuthorization OrgsListSamlSsoAuthorizations(ctx, org).PerPage(perPage).Page(page).Login(login).Execute()

List SAML SSO authorizations for an organization



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
    page := int32(56) // int32 | Page token (optional)
    login := "login_example" // string | Limits the list of credentials authorizations for an organization to a specific login (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListSamlSsoAuthorizations(context.Background(), org).PerPage(perPage).Page(page).Login(login).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListSamlSsoAuthorizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListSamlSsoAuthorizations`: []CredentialAuthorization
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListSamlSsoAuthorizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListSamlSsoAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page token | 
 **login** | **string** | Limits the list of credentials authorizations for an organization to a specific login | 

### Return type

[**[]CredentialAuthorization**](CredentialAuthorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListSecurityManagerTeams

> []TeamSimple OrgsListSecurityManagerTeams(ctx, org).Execute()

List security manager teams



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
    resp, r, err := apiClient.OrgsApi.OrgsListSecurityManagerTeams(context.Background(), org).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListSecurityManagerTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListSecurityManagerTeams`: []TeamSimple
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListSecurityManagerTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListSecurityManagerTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamSimple**](TeamSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsListWebhookDeliveries

> []HookDeliveryItem OrgsListWebhookDeliveries(ctx, org, hookId).PerPage(perPage).Cursor(cursor).Execute()

List deliveries for an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    cursor := "cursor_example" // string | Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the `link` header for the next and previous page cursors. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsListWebhookDeliveries(context.Background(), org, hookId).PerPage(perPage).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListWebhookDeliveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListWebhookDeliveries`: []HookDeliveryItem
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListWebhookDeliveriesRequest struct via the builder pattern


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


## OrgsListWebhooks

> []OrgHook OrgsListWebhooks(ctx, org).PerPage(perPage).Page(page).Execute()

List organization webhooks



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
    resp, r, err := apiClient.OrgsApi.OrgsListWebhooks(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsListWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsListWebhooks`: []OrgHook
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsListWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]OrgHook**](OrgHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsPingWebhook

> OrgsPingWebhook(ctx, org, hookId).Execute()

Ping an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsPingWebhook(context.Background(), org, hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsPingWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsPingWebhookRequest struct via the builder pattern


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


## OrgsRedeliverWebhookDelivery

> map[string]interface{} OrgsRedeliverWebhookDelivery(ctx, org, hookId, deliveryId).Execute()

Redeliver a delivery for an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    deliveryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRedeliverWebhookDelivery(context.Background(), org, hookId, deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRedeliverWebhookDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsRedeliverWebhookDelivery`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsRedeliverWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRedeliverWebhookDeliveryRequest struct via the builder pattern


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


## OrgsRemoveMember

> OrgsRemoveMember(ctx, org, username).Execute()

Remove an organization member



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRemoveMember(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRemoveMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRemoveMemberRequest struct via the builder pattern


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


## OrgsRemoveMembershipForUser

> OrgsRemoveMembershipForUser(ctx, org, username).Execute()

Remove organization membership for a user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRemoveMembershipForUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRemoveMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRemoveMembershipForUserRequest struct via the builder pattern


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


## OrgsRemoveOutsideCollaborator

> OrgsRemoveOutsideCollaborator(ctx, org, username).Execute()

Remove outside collaborator from an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRemoveOutsideCollaborator(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRemoveOutsideCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRemoveOutsideCollaboratorRequest struct via the builder pattern


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


## OrgsRemovePublicMembershipForAuthenticatedUser

> OrgsRemovePublicMembershipForAuthenticatedUser(ctx, org, username).Execute()

Remove public organization membership for the authenticated user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRemovePublicMembershipForAuthenticatedUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRemovePublicMembershipForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRemovePublicMembershipForAuthenticatedUserRequest struct via the builder pattern


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


## OrgsRemoveSamlSsoAuthorization

> OrgsRemoveSamlSsoAuthorization(ctx, org, credentialId).Execute()

Remove a SAML SSO authorization for an organization



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
    credentialId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRemoveSamlSsoAuthorization(context.Background(), org, credentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRemoveSamlSsoAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**credentialId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRemoveSamlSsoAuthorizationRequest struct via the builder pattern


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


## OrgsRemoveSecurityManagerTeam

> OrgsRemoveSecurityManagerTeam(ctx, org, teamSlug).Execute()

Remove a security manager team



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
    teamSlug := "teamSlug_example" // string | The slug of the team name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsRemoveSecurityManagerTeam(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsRemoveSecurityManagerTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsRemoveSecurityManagerTeamRequest struct via the builder pattern


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


## OrgsSetMembershipForUser

> OrgMembership OrgsSetMembershipForUser(ctx, org, username).OrgsSetMembershipForUserRequest(orgsSetMembershipForUserRequest).Execute()

Set organization membership for a user



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
    orgsSetMembershipForUserRequest := *openapiclient.NewOrgsSetMembershipForUserRequest() // OrgsSetMembershipForUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsSetMembershipForUser(context.Background(), org, username).OrgsSetMembershipForUserRequest(orgsSetMembershipForUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsSetMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsSetMembershipForUser`: OrgMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsSetMembershipForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsSetMembershipForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgsSetMembershipForUserRequest** | [**OrgsSetMembershipForUserRequest**](OrgsSetMembershipForUserRequest.md) |  | 

### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsSetPublicMembershipForAuthenticatedUser

> OrgsSetPublicMembershipForAuthenticatedUser(ctx, org, username).Execute()

Set public organization membership for the authenticated user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsSetPublicMembershipForAuthenticatedUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsSetPublicMembershipForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsSetPublicMembershipForAuthenticatedUserRequest struct via the builder pattern


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


## OrgsUnblockUser

> OrgsUnblockUser(ctx, org, username).Execute()

Unblock a user from an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsUnblockUser(context.Background(), org, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsUnblockUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsUnblockUserRequest struct via the builder pattern


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


## OrgsUpdate

> OrganizationFull OrgsUpdate(ctx, org).OrgsUpdateRequest(orgsUpdateRequest).Execute()

Update an organization



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
    orgsUpdateRequest := *openapiclient.NewOrgsUpdateRequest() // OrgsUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsUpdate(context.Background(), org).OrgsUpdateRequest(orgsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsUpdate`: OrganizationFull
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgsUpdateRequest** | [**OrgsUpdateRequest**](OrgsUpdateRequest.md) |  | 

### Return type

[**OrganizationFull**](OrganizationFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsUpdateMembershipForAuthenticatedUser

> OrgMembership OrgsUpdateMembershipForAuthenticatedUser(ctx, org).OrgsUpdateMembershipForAuthenticatedUserRequest(orgsUpdateMembershipForAuthenticatedUserRequest).Execute()

Update an organization membership for the authenticated user



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
    orgsUpdateMembershipForAuthenticatedUserRequest := *openapiclient.NewOrgsUpdateMembershipForAuthenticatedUserRequest("State_example") // OrgsUpdateMembershipForAuthenticatedUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsUpdateMembershipForAuthenticatedUser(context.Background(), org).OrgsUpdateMembershipForAuthenticatedUserRequest(orgsUpdateMembershipForAuthenticatedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsUpdateMembershipForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsUpdateMembershipForAuthenticatedUser`: OrgMembership
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsUpdateMembershipForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsUpdateMembershipForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgsUpdateMembershipForAuthenticatedUserRequest** | [**OrgsUpdateMembershipForAuthenticatedUserRequest**](OrgsUpdateMembershipForAuthenticatedUserRequest.md) |  | 

### Return type

[**OrgMembership**](OrgMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsUpdateWebhook

> OrgHook OrgsUpdateWebhook(ctx, org, hookId).OrgsUpdateWebhookRequest(orgsUpdateWebhookRequest).Execute()

Update an organization webhook



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    orgsUpdateWebhookRequest := *openapiclient.NewOrgsUpdateWebhookRequest() // OrgsUpdateWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsUpdateWebhook(context.Background(), org, hookId).OrgsUpdateWebhookRequest(orgsUpdateWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsUpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsUpdateWebhook`: OrgHook
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgsUpdateWebhookRequest** | [**OrgsUpdateWebhookRequest**](OrgsUpdateWebhookRequest.md) |  | 

### Return type

[**OrgHook**](OrgHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsUpdateWebhookConfigForOrg

> WebhookConfig OrgsUpdateWebhookConfigForOrg(ctx, org, hookId).AppsUpdateWebhookConfigForAppRequest(appsUpdateWebhookConfigForAppRequest).Execute()

Update a webhook configuration for an organization



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
    hookId := int32(56) // int32 | The unique identifier of the hook.
    appsUpdateWebhookConfigForAppRequest := *openapiclient.NewAppsUpdateWebhookConfigForAppRequest() // AppsUpdateWebhookConfigForAppRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.OrgsUpdateWebhookConfigForOrg(context.Background(), org, hookId).AppsUpdateWebhookConfigForAppRequest(appsUpdateWebhookConfigForAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgsUpdateWebhookConfigForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgsUpdateWebhookConfigForOrg`: WebhookConfig
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgsUpdateWebhookConfigForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**hookId** | **int32** | The unique identifier of the hook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgsUpdateWebhookConfigForOrgRequest struct via the builder pattern


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

