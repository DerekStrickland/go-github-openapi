# \TeamsApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsAddMemberLegacy**](TeamsApi.md#TeamsAddMemberLegacy) | **Put** /teams/{team_id}/members/{username} | Add team member (Legacy)
[**TeamsAddOrUpdateMembershipForUserInOrg**](TeamsApi.md#TeamsAddOrUpdateMembershipForUserInOrg) | **Put** /orgs/{org}/teams/{team_slug}/memberships/{username} | Add or update team membership for a user
[**TeamsAddOrUpdateMembershipForUserLegacy**](TeamsApi.md#TeamsAddOrUpdateMembershipForUserLegacy) | **Put** /teams/{team_id}/memberships/{username} | Add or update team membership for a user (Legacy)
[**TeamsAddOrUpdateProjectPermissionsInOrg**](TeamsApi.md#TeamsAddOrUpdateProjectPermissionsInOrg) | **Put** /orgs/{org}/teams/{team_slug}/projects/{project_id} | Add or update team project permissions
[**TeamsAddOrUpdateProjectPermissionsLegacy**](TeamsApi.md#TeamsAddOrUpdateProjectPermissionsLegacy) | **Put** /teams/{team_id}/projects/{project_id} | Add or update team project permissions (Legacy)
[**TeamsAddOrUpdateRepoPermissionsInOrg**](TeamsApi.md#TeamsAddOrUpdateRepoPermissionsInOrg) | **Put** /orgs/{org}/teams/{team_slug}/repos/{owner}/{repo} | Add or update team repository permissions
[**TeamsAddOrUpdateRepoPermissionsLegacy**](TeamsApi.md#TeamsAddOrUpdateRepoPermissionsLegacy) | **Put** /teams/{team_id}/repos/{owner}/{repo} | Add or update team repository permissions (Legacy)
[**TeamsCheckPermissionsForProjectInOrg**](TeamsApi.md#TeamsCheckPermissionsForProjectInOrg) | **Get** /orgs/{org}/teams/{team_slug}/projects/{project_id} | Check team permissions for a project
[**TeamsCheckPermissionsForProjectLegacy**](TeamsApi.md#TeamsCheckPermissionsForProjectLegacy) | **Get** /teams/{team_id}/projects/{project_id} | Check team permissions for a project (Legacy)
[**TeamsCheckPermissionsForRepoInOrg**](TeamsApi.md#TeamsCheckPermissionsForRepoInOrg) | **Get** /orgs/{org}/teams/{team_slug}/repos/{owner}/{repo} | Check team permissions for a repository
[**TeamsCheckPermissionsForRepoLegacy**](TeamsApi.md#TeamsCheckPermissionsForRepoLegacy) | **Get** /teams/{team_id}/repos/{owner}/{repo} | Check team permissions for a repository (Legacy)
[**TeamsCreate**](TeamsApi.md#TeamsCreate) | **Post** /orgs/{org}/teams | Create a team
[**TeamsCreateDiscussionCommentInOrg**](TeamsApi.md#TeamsCreateDiscussionCommentInOrg) | **Post** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments | Create a discussion comment
[**TeamsCreateDiscussionCommentLegacy**](TeamsApi.md#TeamsCreateDiscussionCommentLegacy) | **Post** /teams/{team_id}/discussions/{discussion_number}/comments | Create a discussion comment (Legacy)
[**TeamsCreateDiscussionInOrg**](TeamsApi.md#TeamsCreateDiscussionInOrg) | **Post** /orgs/{org}/teams/{team_slug}/discussions | Create a discussion
[**TeamsCreateDiscussionLegacy**](TeamsApi.md#TeamsCreateDiscussionLegacy) | **Post** /teams/{team_id}/discussions | Create a discussion (Legacy)
[**TeamsCreateOrUpdateIdpGroupConnectionsInOrg**](TeamsApi.md#TeamsCreateOrUpdateIdpGroupConnectionsInOrg) | **Patch** /orgs/{org}/teams/{team_slug}/team-sync/group-mappings | Create or update IdP group connections
[**TeamsCreateOrUpdateIdpGroupConnectionsLegacy**](TeamsApi.md#TeamsCreateOrUpdateIdpGroupConnectionsLegacy) | **Patch** /teams/{team_id}/team-sync/group-mappings | Create or update IdP group connections (Legacy)
[**TeamsDeleteDiscussionCommentInOrg**](TeamsApi.md#TeamsDeleteDiscussionCommentInOrg) | **Delete** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number} | Delete a discussion comment
[**TeamsDeleteDiscussionCommentLegacy**](TeamsApi.md#TeamsDeleteDiscussionCommentLegacy) | **Delete** /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number} | Delete a discussion comment (Legacy)
[**TeamsDeleteDiscussionInOrg**](TeamsApi.md#TeamsDeleteDiscussionInOrg) | **Delete** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number} | Delete a discussion
[**TeamsDeleteDiscussionLegacy**](TeamsApi.md#TeamsDeleteDiscussionLegacy) | **Delete** /teams/{team_id}/discussions/{discussion_number} | Delete a discussion (Legacy)
[**TeamsDeleteInOrg**](TeamsApi.md#TeamsDeleteInOrg) | **Delete** /orgs/{org}/teams/{team_slug} | Delete a team
[**TeamsDeleteLegacy**](TeamsApi.md#TeamsDeleteLegacy) | **Delete** /teams/{team_id} | Delete a team (Legacy)
[**TeamsExternalIdpGroupInfoForOrg**](TeamsApi.md#TeamsExternalIdpGroupInfoForOrg) | **Get** /orgs/{org}/external-group/{group_id} | Get an external group
[**TeamsGetByName**](TeamsApi.md#TeamsGetByName) | **Get** /orgs/{org}/teams/{team_slug} | Get a team by name
[**TeamsGetDiscussionCommentInOrg**](TeamsApi.md#TeamsGetDiscussionCommentInOrg) | **Get** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number} | Get a discussion comment
[**TeamsGetDiscussionCommentLegacy**](TeamsApi.md#TeamsGetDiscussionCommentLegacy) | **Get** /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number} | Get a discussion comment (Legacy)
[**TeamsGetDiscussionInOrg**](TeamsApi.md#TeamsGetDiscussionInOrg) | **Get** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number} | Get a discussion
[**TeamsGetDiscussionLegacy**](TeamsApi.md#TeamsGetDiscussionLegacy) | **Get** /teams/{team_id}/discussions/{discussion_number} | Get a discussion (Legacy)
[**TeamsGetLegacy**](TeamsApi.md#TeamsGetLegacy) | **Get** /teams/{team_id} | Get a team (Legacy)
[**TeamsGetMemberLegacy**](TeamsApi.md#TeamsGetMemberLegacy) | **Get** /teams/{team_id}/members/{username} | Get team member (Legacy)
[**TeamsGetMembershipForUserInOrg**](TeamsApi.md#TeamsGetMembershipForUserInOrg) | **Get** /orgs/{org}/teams/{team_slug}/memberships/{username} | Get team membership for a user
[**TeamsGetMembershipForUserLegacy**](TeamsApi.md#TeamsGetMembershipForUserLegacy) | **Get** /teams/{team_id}/memberships/{username} | Get team membership for a user (Legacy)
[**TeamsLinkExternalIdpGroupToTeamForOrg**](TeamsApi.md#TeamsLinkExternalIdpGroupToTeamForOrg) | **Patch** /orgs/{org}/teams/{team_slug}/external-groups | Update the connection between an external group and a team
[**TeamsList**](TeamsApi.md#TeamsList) | **Get** /orgs/{org}/teams | List teams
[**TeamsListChildInOrg**](TeamsApi.md#TeamsListChildInOrg) | **Get** /orgs/{org}/teams/{team_slug}/teams | List child teams
[**TeamsListChildLegacy**](TeamsApi.md#TeamsListChildLegacy) | **Get** /teams/{team_id}/teams | List child teams (Legacy)
[**TeamsListDiscussionCommentsInOrg**](TeamsApi.md#TeamsListDiscussionCommentsInOrg) | **Get** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments | List discussion comments
[**TeamsListDiscussionCommentsLegacy**](TeamsApi.md#TeamsListDiscussionCommentsLegacy) | **Get** /teams/{team_id}/discussions/{discussion_number}/comments | List discussion comments (Legacy)
[**TeamsListDiscussionsInOrg**](TeamsApi.md#TeamsListDiscussionsInOrg) | **Get** /orgs/{org}/teams/{team_slug}/discussions | List discussions
[**TeamsListDiscussionsLegacy**](TeamsApi.md#TeamsListDiscussionsLegacy) | **Get** /teams/{team_id}/discussions | List discussions (Legacy)
[**TeamsListExternalIdpGroupsForOrg**](TeamsApi.md#TeamsListExternalIdpGroupsForOrg) | **Get** /orgs/{org}/external-groups | List external groups in an organization
[**TeamsListForAuthenticatedUser**](TeamsApi.md#TeamsListForAuthenticatedUser) | **Get** /user/teams | List teams for the authenticated user
[**TeamsListIdpGroupsForLegacy**](TeamsApi.md#TeamsListIdpGroupsForLegacy) | **Get** /teams/{team_id}/team-sync/group-mappings | List IdP groups for a team (Legacy)
[**TeamsListIdpGroupsForOrg**](TeamsApi.md#TeamsListIdpGroupsForOrg) | **Get** /orgs/{org}/team-sync/groups | List IdP groups for an organization
[**TeamsListIdpGroupsInOrg**](TeamsApi.md#TeamsListIdpGroupsInOrg) | **Get** /orgs/{org}/teams/{team_slug}/team-sync/group-mappings | List IdP groups for a team
[**TeamsListLinkedExternalIdpGroupsToTeamForOrg**](TeamsApi.md#TeamsListLinkedExternalIdpGroupsToTeamForOrg) | **Get** /orgs/{org}/teams/{team_slug}/external-groups | List a connection between an external group and a team
[**TeamsListMembersInOrg**](TeamsApi.md#TeamsListMembersInOrg) | **Get** /orgs/{org}/teams/{team_slug}/members | List team members
[**TeamsListMembersLegacy**](TeamsApi.md#TeamsListMembersLegacy) | **Get** /teams/{team_id}/members | List team members (Legacy)
[**TeamsListPendingInvitationsInOrg**](TeamsApi.md#TeamsListPendingInvitationsInOrg) | **Get** /orgs/{org}/teams/{team_slug}/invitations | List pending team invitations
[**TeamsListPendingInvitationsLegacy**](TeamsApi.md#TeamsListPendingInvitationsLegacy) | **Get** /teams/{team_id}/invitations | List pending team invitations (Legacy)
[**TeamsListProjectsInOrg**](TeamsApi.md#TeamsListProjectsInOrg) | **Get** /orgs/{org}/teams/{team_slug}/projects | List team projects
[**TeamsListProjectsLegacy**](TeamsApi.md#TeamsListProjectsLegacy) | **Get** /teams/{team_id}/projects | List team projects (Legacy)
[**TeamsListReposInOrg**](TeamsApi.md#TeamsListReposInOrg) | **Get** /orgs/{org}/teams/{team_slug}/repos | List team repositories
[**TeamsListReposLegacy**](TeamsApi.md#TeamsListReposLegacy) | **Get** /teams/{team_id}/repos | List team repositories (Legacy)
[**TeamsRemoveMemberLegacy**](TeamsApi.md#TeamsRemoveMemberLegacy) | **Delete** /teams/{team_id}/members/{username} | Remove team member (Legacy)
[**TeamsRemoveMembershipForUserInOrg**](TeamsApi.md#TeamsRemoveMembershipForUserInOrg) | **Delete** /orgs/{org}/teams/{team_slug}/memberships/{username} | Remove team membership for a user
[**TeamsRemoveMembershipForUserLegacy**](TeamsApi.md#TeamsRemoveMembershipForUserLegacy) | **Delete** /teams/{team_id}/memberships/{username} | Remove team membership for a user (Legacy)
[**TeamsRemoveProjectInOrg**](TeamsApi.md#TeamsRemoveProjectInOrg) | **Delete** /orgs/{org}/teams/{team_slug}/projects/{project_id} | Remove a project from a team
[**TeamsRemoveProjectLegacy**](TeamsApi.md#TeamsRemoveProjectLegacy) | **Delete** /teams/{team_id}/projects/{project_id} | Remove a project from a team (Legacy)
[**TeamsRemoveRepoInOrg**](TeamsApi.md#TeamsRemoveRepoInOrg) | **Delete** /orgs/{org}/teams/{team_slug}/repos/{owner}/{repo} | Remove a repository from a team
[**TeamsRemoveRepoLegacy**](TeamsApi.md#TeamsRemoveRepoLegacy) | **Delete** /teams/{team_id}/repos/{owner}/{repo} | Remove a repository from a team (Legacy)
[**TeamsUnlinkExternalIdpGroupFromTeamForOrg**](TeamsApi.md#TeamsUnlinkExternalIdpGroupFromTeamForOrg) | **Delete** /orgs/{org}/teams/{team_slug}/external-groups | Remove the connection between an external group and a team
[**TeamsUpdateDiscussionCommentInOrg**](TeamsApi.md#TeamsUpdateDiscussionCommentInOrg) | **Patch** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number} | Update a discussion comment
[**TeamsUpdateDiscussionCommentLegacy**](TeamsApi.md#TeamsUpdateDiscussionCommentLegacy) | **Patch** /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number} | Update a discussion comment (Legacy)
[**TeamsUpdateDiscussionInOrg**](TeamsApi.md#TeamsUpdateDiscussionInOrg) | **Patch** /orgs/{org}/teams/{team_slug}/discussions/{discussion_number} | Update a discussion
[**TeamsUpdateDiscussionLegacy**](TeamsApi.md#TeamsUpdateDiscussionLegacy) | **Patch** /teams/{team_id}/discussions/{discussion_number} | Update a discussion (Legacy)
[**TeamsUpdateInOrg**](TeamsApi.md#TeamsUpdateInOrg) | **Patch** /orgs/{org}/teams/{team_slug} | Update a team
[**TeamsUpdateLegacy**](TeamsApi.md#TeamsUpdateLegacy) | **Patch** /teams/{team_id} | Update a team (Legacy)



## TeamsAddMemberLegacy

> TeamsAddMemberLegacy(ctx, teamId, username).Execute()

Add team member (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddMemberLegacy(context.Background(), teamId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddMemberLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddMemberLegacyRequest struct via the builder pattern


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


## TeamsAddOrUpdateMembershipForUserInOrg

> TeamMembership TeamsAddOrUpdateMembershipForUserInOrg(ctx, org, teamSlug, username).TeamsAddOrUpdateMembershipForUserInOrgRequest(teamsAddOrUpdateMembershipForUserInOrgRequest).Execute()

Add or update team membership for a user



### Example

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
    username := "username_example" // string | The handle for the GitHub user account.
    teamsAddOrUpdateMembershipForUserInOrgRequest := *openapiclient.NewTeamsAddOrUpdateMembershipForUserInOrgRequest() // TeamsAddOrUpdateMembershipForUserInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddOrUpdateMembershipForUserInOrg(context.Background(), org, teamSlug, username).TeamsAddOrUpdateMembershipForUserInOrgRequest(teamsAddOrUpdateMembershipForUserInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddOrUpdateMembershipForUserInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsAddOrUpdateMembershipForUserInOrg`: TeamMembership
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsAddOrUpdateMembershipForUserInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddOrUpdateMembershipForUserInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamsAddOrUpdateMembershipForUserInOrgRequest** | [**TeamsAddOrUpdateMembershipForUserInOrgRequest**](TeamsAddOrUpdateMembershipForUserInOrgRequest.md) |  | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsAddOrUpdateMembershipForUserLegacy

> TeamMembership TeamsAddOrUpdateMembershipForUserLegacy(ctx, teamId, username).TeamsAddOrUpdateMembershipForUserInOrgRequest(teamsAddOrUpdateMembershipForUserInOrgRequest).Execute()

Add or update team membership for a user (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    username := "username_example" // string | The handle for the GitHub user account.
    teamsAddOrUpdateMembershipForUserInOrgRequest := *openapiclient.NewTeamsAddOrUpdateMembershipForUserInOrgRequest() // TeamsAddOrUpdateMembershipForUserInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddOrUpdateMembershipForUserLegacy(context.Background(), teamId, username).TeamsAddOrUpdateMembershipForUserInOrgRequest(teamsAddOrUpdateMembershipForUserInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddOrUpdateMembershipForUserLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsAddOrUpdateMembershipForUserLegacy`: TeamMembership
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsAddOrUpdateMembershipForUserLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddOrUpdateMembershipForUserLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsAddOrUpdateMembershipForUserInOrgRequest** | [**TeamsAddOrUpdateMembershipForUserInOrgRequest**](TeamsAddOrUpdateMembershipForUserInOrgRequest.md) |  | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsAddOrUpdateProjectPermissionsInOrg

> TeamsAddOrUpdateProjectPermissionsInOrg(ctx, org, teamSlug, projectId).TeamsAddOrUpdateProjectPermissionsInOrgRequest(teamsAddOrUpdateProjectPermissionsInOrgRequest).Execute()

Add or update team project permissions



### Example

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
    projectId := int32(56) // int32 | The unique identifier of the project.
    teamsAddOrUpdateProjectPermissionsInOrgRequest := *openapiclient.NewTeamsAddOrUpdateProjectPermissionsInOrgRequest() // TeamsAddOrUpdateProjectPermissionsInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddOrUpdateProjectPermissionsInOrg(context.Background(), org, teamSlug, projectId).TeamsAddOrUpdateProjectPermissionsInOrgRequest(teamsAddOrUpdateProjectPermissionsInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddOrUpdateProjectPermissionsInOrg``: %v\n", err)
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
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddOrUpdateProjectPermissionsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamsAddOrUpdateProjectPermissionsInOrgRequest** | [**TeamsAddOrUpdateProjectPermissionsInOrgRequest**](TeamsAddOrUpdateProjectPermissionsInOrgRequest.md) |  | 

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


## TeamsAddOrUpdateProjectPermissionsLegacy

> TeamsAddOrUpdateProjectPermissionsLegacy(ctx, teamId, projectId).TeamsAddOrUpdateProjectPermissionsLegacyRequest(teamsAddOrUpdateProjectPermissionsLegacyRequest).Execute()

Add or update team project permissions (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    projectId := int32(56) // int32 | The unique identifier of the project.
    teamsAddOrUpdateProjectPermissionsLegacyRequest := *openapiclient.NewTeamsAddOrUpdateProjectPermissionsLegacyRequest() // TeamsAddOrUpdateProjectPermissionsLegacyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddOrUpdateProjectPermissionsLegacy(context.Background(), teamId, projectId).TeamsAddOrUpdateProjectPermissionsLegacyRequest(teamsAddOrUpdateProjectPermissionsLegacyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddOrUpdateProjectPermissionsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddOrUpdateProjectPermissionsLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsAddOrUpdateProjectPermissionsLegacyRequest** | [**TeamsAddOrUpdateProjectPermissionsLegacyRequest**](TeamsAddOrUpdateProjectPermissionsLegacyRequest.md) |  | 

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


## TeamsAddOrUpdateRepoPermissionsInOrg

> TeamsAddOrUpdateRepoPermissionsInOrg(ctx, org, teamSlug, owner, repo).TeamsAddOrUpdateRepoPermissionsInOrgRequest(teamsAddOrUpdateRepoPermissionsInOrgRequest).Execute()

Add or update team repository permissions



### Example

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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    teamsAddOrUpdateRepoPermissionsInOrgRequest := *openapiclient.NewTeamsAddOrUpdateRepoPermissionsInOrgRequest() // TeamsAddOrUpdateRepoPermissionsInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddOrUpdateRepoPermissionsInOrg(context.Background(), org, teamSlug, owner, repo).TeamsAddOrUpdateRepoPermissionsInOrgRequest(teamsAddOrUpdateRepoPermissionsInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddOrUpdateRepoPermissionsInOrg``: %v\n", err)
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
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddOrUpdateRepoPermissionsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **teamsAddOrUpdateRepoPermissionsInOrgRequest** | [**TeamsAddOrUpdateRepoPermissionsInOrgRequest**](TeamsAddOrUpdateRepoPermissionsInOrgRequest.md) |  | 

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


## TeamsAddOrUpdateRepoPermissionsLegacy

> TeamsAddOrUpdateRepoPermissionsLegacy(ctx, teamId, owner, repo).TeamsAddOrUpdateRepoPermissionsLegacyRequest(teamsAddOrUpdateRepoPermissionsLegacyRequest).Execute()

Add or update team repository permissions (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.
    teamsAddOrUpdateRepoPermissionsLegacyRequest := *openapiclient.NewTeamsAddOrUpdateRepoPermissionsLegacyRequest() // TeamsAddOrUpdateRepoPermissionsLegacyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsAddOrUpdateRepoPermissionsLegacy(context.Background(), teamId, owner, repo).TeamsAddOrUpdateRepoPermissionsLegacyRequest(teamsAddOrUpdateRepoPermissionsLegacyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsAddOrUpdateRepoPermissionsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsAddOrUpdateRepoPermissionsLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamsAddOrUpdateRepoPermissionsLegacyRequest** | [**TeamsAddOrUpdateRepoPermissionsLegacyRequest**](TeamsAddOrUpdateRepoPermissionsLegacyRequest.md) |  | 

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


## TeamsCheckPermissionsForProjectInOrg

> TeamProject TeamsCheckPermissionsForProjectInOrg(ctx, org, teamSlug, projectId).Execute()

Check team permissions for a project



### Example

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
    projectId := int32(56) // int32 | The unique identifier of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCheckPermissionsForProjectInOrg(context.Background(), org, teamSlug, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCheckPermissionsForProjectInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCheckPermissionsForProjectInOrg`: TeamProject
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCheckPermissionsForProjectInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCheckPermissionsForProjectInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TeamProject**](TeamProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCheckPermissionsForProjectLegacy

> TeamProject TeamsCheckPermissionsForProjectLegacy(ctx, teamId, projectId).Execute()

Check team permissions for a project (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    projectId := int32(56) // int32 | The unique identifier of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCheckPermissionsForProjectLegacy(context.Background(), teamId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCheckPermissionsForProjectLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCheckPermissionsForProjectLegacy`: TeamProject
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCheckPermissionsForProjectLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCheckPermissionsForProjectLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TeamProject**](TeamProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCheckPermissionsForRepoInOrg

> TeamRepository TeamsCheckPermissionsForRepoInOrg(ctx, org, teamSlug, owner, repo).Execute()

Check team permissions for a repository



### Example

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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCheckPermissionsForRepoInOrg(context.Background(), org, teamSlug, owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCheckPermissionsForRepoInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCheckPermissionsForRepoInOrg`: TeamRepository
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCheckPermissionsForRepoInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCheckPermissionsForRepoInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**TeamRepository**](TeamRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCheckPermissionsForRepoLegacy

> TeamRepository TeamsCheckPermissionsForRepoLegacy(ctx, teamId, owner, repo).Execute()

Check team permissions for a repository (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCheckPermissionsForRepoLegacy(context.Background(), teamId, owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCheckPermissionsForRepoLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCheckPermissionsForRepoLegacy`: TeamRepository
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCheckPermissionsForRepoLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCheckPermissionsForRepoLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TeamRepository**](TeamRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreate

> TeamFull TeamsCreate(ctx, org).TeamsCreateRequest(teamsCreateRequest).Execute()

Create a team



### Example

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
    teamsCreateRequest := *openapiclient.NewTeamsCreateRequest("Name_example") // TeamsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreate(context.Background(), org).TeamsCreateRequest(teamsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreate`: TeamFull
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamsCreateRequest** | [**TeamsCreateRequest**](TeamsCreateRequest.md) |  | 

### Return type

[**TeamFull**](TeamFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreateDiscussionCommentInOrg

> TeamDiscussionComment TeamsCreateDiscussionCommentInOrg(ctx, org, teamSlug, discussionNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()

Create a discussion comment



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    teamsCreateDiscussionCommentInOrgRequest := *openapiclient.NewTeamsCreateDiscussionCommentInOrgRequest("Body_example") // TeamsCreateDiscussionCommentInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreateDiscussionCommentInOrg(context.Background(), org, teamSlug, discussionNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreateDiscussionCommentInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateDiscussionCommentInOrg`: TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreateDiscussionCommentInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateDiscussionCommentInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamsCreateDiscussionCommentInOrgRequest** | [**TeamsCreateDiscussionCommentInOrgRequest**](TeamsCreateDiscussionCommentInOrgRequest.md) |  | 

### Return type

[**TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreateDiscussionCommentLegacy

> TeamDiscussionComment TeamsCreateDiscussionCommentLegacy(ctx, teamId, discussionNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()

Create a discussion comment (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    teamsCreateDiscussionCommentInOrgRequest := *openapiclient.NewTeamsCreateDiscussionCommentInOrgRequest("Body_example") // TeamsCreateDiscussionCommentInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreateDiscussionCommentLegacy(context.Background(), teamId, discussionNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreateDiscussionCommentLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateDiscussionCommentLegacy`: TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreateDiscussionCommentLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateDiscussionCommentLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsCreateDiscussionCommentInOrgRequest** | [**TeamsCreateDiscussionCommentInOrgRequest**](TeamsCreateDiscussionCommentInOrgRequest.md) |  | 

### Return type

[**TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreateDiscussionInOrg

> TeamDiscussion TeamsCreateDiscussionInOrg(ctx, org, teamSlug).TeamsCreateDiscussionInOrgRequest(teamsCreateDiscussionInOrgRequest).Execute()

Create a discussion



### Example

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
    teamsCreateDiscussionInOrgRequest := *openapiclient.NewTeamsCreateDiscussionInOrgRequest("Title_example", "Body_example") // TeamsCreateDiscussionInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreateDiscussionInOrg(context.Background(), org, teamSlug).TeamsCreateDiscussionInOrgRequest(teamsCreateDiscussionInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreateDiscussionInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateDiscussionInOrg`: TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreateDiscussionInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateDiscussionInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsCreateDiscussionInOrgRequest** | [**TeamsCreateDiscussionInOrgRequest**](TeamsCreateDiscussionInOrgRequest.md) |  | 

### Return type

[**TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreateDiscussionLegacy

> TeamDiscussion TeamsCreateDiscussionLegacy(ctx, teamId).TeamsCreateDiscussionInOrgRequest(teamsCreateDiscussionInOrgRequest).Execute()

Create a discussion (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    teamsCreateDiscussionInOrgRequest := *openapiclient.NewTeamsCreateDiscussionInOrgRequest("Title_example", "Body_example") // TeamsCreateDiscussionInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreateDiscussionLegacy(context.Background(), teamId).TeamsCreateDiscussionInOrgRequest(teamsCreateDiscussionInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreateDiscussionLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateDiscussionLegacy`: TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreateDiscussionLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateDiscussionLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamsCreateDiscussionInOrgRequest** | [**TeamsCreateDiscussionInOrgRequest**](TeamsCreateDiscussionInOrgRequest.md) |  | 

### Return type

[**TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreateOrUpdateIdpGroupConnectionsInOrg

> GroupMapping TeamsCreateOrUpdateIdpGroupConnectionsInOrg(ctx, org, teamSlug).TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest(teamsCreateOrUpdateIdpGroupConnectionsInOrgRequest).Execute()

Create or update IdP group connections



### Example

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
    teamsCreateOrUpdateIdpGroupConnectionsInOrgRequest := *openapiclient.NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest() // TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreateOrUpdateIdpGroupConnectionsInOrg(context.Background(), org, teamSlug).TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest(teamsCreateOrUpdateIdpGroupConnectionsInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreateOrUpdateIdpGroupConnectionsInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateOrUpdateIdpGroupConnectionsInOrg`: GroupMapping
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreateOrUpdateIdpGroupConnectionsInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsCreateOrUpdateIdpGroupConnectionsInOrgRequest** | [**TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest**](TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest.md) |  | 

### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreateOrUpdateIdpGroupConnectionsLegacy

> GroupMapping TeamsCreateOrUpdateIdpGroupConnectionsLegacy(ctx, teamId).TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest(teamsCreateOrUpdateIdpGroupConnectionsLegacyRequest).Execute()

Create or update IdP group connections (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    teamsCreateOrUpdateIdpGroupConnectionsLegacyRequest := *openapiclient.NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest([]openapiclient.TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner{*openapiclient.NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner("GroupId_example", "GroupName_example", "GroupDescription_example")}) // TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreateOrUpdateIdpGroupConnectionsLegacy(context.Background(), teamId).TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest(teamsCreateOrUpdateIdpGroupConnectionsLegacyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreateOrUpdateIdpGroupConnectionsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateOrUpdateIdpGroupConnectionsLegacy`: GroupMapping
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreateOrUpdateIdpGroupConnectionsLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamsCreateOrUpdateIdpGroupConnectionsLegacyRequest** | [**TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest**](TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest.md) |  | 

### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDeleteDiscussionCommentInOrg

> TeamsDeleteDiscussionCommentInOrg(ctx, org, teamSlug, discussionNumber, commentNumber).Execute()

Delete a discussion comment



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsDeleteDiscussionCommentInOrg(context.Background(), org, teamSlug, discussionNumber, commentNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDeleteDiscussionCommentInOrg``: %v\n", err)
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
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteDiscussionCommentInOrgRequest struct via the builder pattern


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


## TeamsDeleteDiscussionCommentLegacy

> TeamsDeleteDiscussionCommentLegacy(ctx, teamId, discussionNumber, commentNumber).Execute()

Delete a discussion comment (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsDeleteDiscussionCommentLegacy(context.Background(), teamId, discussionNumber, commentNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDeleteDiscussionCommentLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteDiscussionCommentLegacyRequest struct via the builder pattern


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


## TeamsDeleteDiscussionInOrg

> TeamsDeleteDiscussionInOrg(ctx, org, teamSlug, discussionNumber).Execute()

Delete a discussion



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsDeleteDiscussionInOrg(context.Background(), org, teamSlug, discussionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDeleteDiscussionInOrg``: %v\n", err)
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
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteDiscussionInOrgRequest struct via the builder pattern


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


## TeamsDeleteDiscussionLegacy

> TeamsDeleteDiscussionLegacy(ctx, teamId, discussionNumber).Execute()

Delete a discussion (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsDeleteDiscussionLegacy(context.Background(), teamId, discussionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDeleteDiscussionLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteDiscussionLegacyRequest struct via the builder pattern


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


## TeamsDeleteInOrg

> TeamsDeleteInOrg(ctx, org, teamSlug).Execute()

Delete a team



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsDeleteInOrg(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDeleteInOrg``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsDeleteInOrgRequest struct via the builder pattern


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


## TeamsDeleteLegacy

> TeamsDeleteLegacy(ctx, teamId).Execute()

Delete a team (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsDeleteLegacy(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDeleteLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteLegacyRequest struct via the builder pattern


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


## TeamsExternalIdpGroupInfoForOrg

> ExternalGroup TeamsExternalIdpGroupInfoForOrg(ctx, org, groupId).Execute()

Get an external group



### Example

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
    groupId := int32(56) // int32 | The unique identifier of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsExternalIdpGroupInfoForOrg(context.Background(), org, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsExternalIdpGroupInfoForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsExternalIdpGroupInfoForOrg`: ExternalGroup
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsExternalIdpGroupInfoForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**groupId** | **int32** | The unique identifier of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsExternalIdpGroupInfoForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExternalGroup**](ExternalGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetByName

> TeamFull TeamsGetByName(ctx, org, teamSlug).Execute()

Get a team by name



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsGetByName(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetByName`: TeamFull
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TeamFull**](TeamFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetDiscussionCommentInOrg

> TeamDiscussionComment TeamsGetDiscussionCommentInOrg(ctx, org, teamSlug, discussionNumber, commentNumber).Execute()

Get a discussion comment



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetDiscussionCommentInOrg(context.Background(), org, teamSlug, discussionNumber, commentNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetDiscussionCommentInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetDiscussionCommentInOrg`: TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetDiscussionCommentInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetDiscussionCommentInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetDiscussionCommentLegacy

> TeamDiscussionComment TeamsGetDiscussionCommentLegacy(ctx, teamId, discussionNumber, commentNumber).Execute()

Get a discussion comment (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetDiscussionCommentLegacy(context.Background(), teamId, discussionNumber, commentNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetDiscussionCommentLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetDiscussionCommentLegacy`: TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetDiscussionCommentLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetDiscussionCommentLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetDiscussionInOrg

> TeamDiscussion TeamsGetDiscussionInOrg(ctx, org, teamSlug, discussionNumber).Execute()

Get a discussion



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetDiscussionInOrg(context.Background(), org, teamSlug, discussionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetDiscussionInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetDiscussionInOrg`: TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetDiscussionInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetDiscussionInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetDiscussionLegacy

> TeamDiscussion TeamsGetDiscussionLegacy(ctx, teamId, discussionNumber).Execute()

Get a discussion (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetDiscussionLegacy(context.Background(), teamId, discussionNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetDiscussionLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetDiscussionLegacy`: TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetDiscussionLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetDiscussionLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetLegacy

> TeamFull TeamsGetLegacy(ctx, teamId).Execute()

Get a team (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetLegacy(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetLegacy`: TeamFull
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamFull**](TeamFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetMemberLegacy

> TeamsGetMemberLegacy(ctx, teamId, username).Execute()

Get team member (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetMemberLegacy(context.Background(), teamId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetMemberLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetMemberLegacyRequest struct via the builder pattern


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


## TeamsGetMembershipForUserInOrg

> TeamMembership TeamsGetMembershipForUserInOrg(ctx, org, teamSlug, username).Execute()

Get team membership for a user



### Example

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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetMembershipForUserInOrg(context.Background(), org, teamSlug, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetMembershipForUserInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetMembershipForUserInOrg`: TeamMembership
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetMembershipForUserInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetMembershipForUserInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetMembershipForUserLegacy

> TeamMembership TeamsGetMembershipForUserLegacy(ctx, teamId, username).Execute()

Get team membership for a user (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsGetMembershipForUserLegacy(context.Background(), teamId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGetMembershipForUserLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetMembershipForUserLegacy`: TeamMembership
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGetMembershipForUserLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetMembershipForUserLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsLinkExternalIdpGroupToTeamForOrg

> ExternalGroup TeamsLinkExternalIdpGroupToTeamForOrg(ctx, org, teamSlug).TeamsLinkExternalIdpGroupToTeamForOrgRequest(teamsLinkExternalIdpGroupToTeamForOrgRequest).Execute()

Update the connection between an external group and a team



### Example

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
    teamsLinkExternalIdpGroupToTeamForOrgRequest := *openapiclient.NewTeamsLinkExternalIdpGroupToTeamForOrgRequest(int32(1)) // TeamsLinkExternalIdpGroupToTeamForOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsLinkExternalIdpGroupToTeamForOrg(context.Background(), org, teamSlug).TeamsLinkExternalIdpGroupToTeamForOrgRequest(teamsLinkExternalIdpGroupToTeamForOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsLinkExternalIdpGroupToTeamForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsLinkExternalIdpGroupToTeamForOrg`: ExternalGroup
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsLinkExternalIdpGroupToTeamForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsLinkExternalIdpGroupToTeamForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsLinkExternalIdpGroupToTeamForOrgRequest** | [**TeamsLinkExternalIdpGroupToTeamForOrgRequest**](TeamsLinkExternalIdpGroupToTeamForOrgRequest.md) |  | 

### Return type

[**ExternalGroup**](ExternalGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsList

> []Team TeamsList(ctx, org).PerPage(perPage).Page(page).Execute()

List teams



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsList(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsList`: []Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListRequest struct via the builder pattern


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


## TeamsListChildInOrg

> []Team TeamsListChildInOrg(ctx, org, teamSlug).PerPage(perPage).Page(page).Execute()

List child teams



### Example

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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListChildInOrg(context.Background(), org, teamSlug).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListChildInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListChildInOrg`: []Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListChildInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListChildInOrgRequest struct via the builder pattern


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


## TeamsListChildLegacy

> []Team TeamsListChildLegacy(ctx, teamId).PerPage(perPage).Page(page).Execute()

List child teams (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListChildLegacy(context.Background(), teamId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListChildLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListChildLegacy`: []Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListChildLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListChildLegacyRequest struct via the builder pattern


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


## TeamsListDiscussionCommentsInOrg

> []TeamDiscussionComment TeamsListDiscussionCommentsInOrg(ctx, org, teamSlug, discussionNumber).Direction(direction).PerPage(perPage).Page(page).Execute()

List discussion comments



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListDiscussionCommentsInOrg(context.Background(), org, teamSlug, discussionNumber).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListDiscussionCommentsInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListDiscussionCommentsInOrg`: []TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListDiscussionCommentsInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListDiscussionCommentsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListDiscussionCommentsLegacy

> []TeamDiscussionComment TeamsListDiscussionCommentsLegacy(ctx, teamId, discussionNumber).Direction(direction).PerPage(perPage).Page(page).Execute()

List discussion comments (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListDiscussionCommentsLegacy(context.Background(), teamId, discussionNumber).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListDiscussionCommentsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListDiscussionCommentsLegacy`: []TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListDiscussionCommentsLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListDiscussionCommentsLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListDiscussionsInOrg

> []TeamDiscussion TeamsListDiscussionsInOrg(ctx, org, teamSlug).Direction(direction).PerPage(perPage).Page(page).Pinned(pinned).Execute()

List discussions



### Example

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
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)
    pinned := "pinned_example" // string | Pinned discussions only filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListDiscussionsInOrg(context.Background(), org, teamSlug).Direction(direction).PerPage(perPage).Page(page).Pinned(pinned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListDiscussionsInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListDiscussionsInOrg`: []TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListDiscussionsInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListDiscussionsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]
 **pinned** | **string** | Pinned discussions only filter | 

### Return type

[**[]TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListDiscussionsLegacy

> []TeamDiscussion TeamsListDiscussionsLegacy(ctx, teamId).Direction(direction).PerPage(perPage).Page(page).Execute()

List discussions (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    direction := "direction_example" // string | The direction to sort the results by. (optional) (default to "desc")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListDiscussionsLegacy(context.Background(), teamId).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListDiscussionsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListDiscussionsLegacy`: []TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListDiscussionsLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListDiscussionsLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **string** | The direction to sort the results by. | [default to &quot;desc&quot;]
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListExternalIdpGroupsForOrg

> ExternalGroups TeamsListExternalIdpGroupsForOrg(ctx, org).PerPage(perPage).Page(page).DisplayName(displayName).Execute()

List external groups in an organization



### Example

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
    displayName := "displayName_example" // string | Limits the list to groups containing the text in the group name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListExternalIdpGroupsForOrg(context.Background(), org).PerPage(perPage).Page(page).DisplayName(displayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListExternalIdpGroupsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListExternalIdpGroupsForOrg`: ExternalGroups
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListExternalIdpGroupsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListExternalIdpGroupsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page token | 
 **displayName** | **string** | Limits the list to groups containing the text in the group name | 

### Return type

[**ExternalGroups**](ExternalGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListForAuthenticatedUser

> []TeamFull TeamsListForAuthenticatedUser(ctx).PerPage(perPage).Page(page).Execute()

List teams for the authenticated user



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsListForAuthenticatedUser(context.Background()).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListForAuthenticatedUser`: []TeamFull
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TeamFull**](TeamFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListIdpGroupsForLegacy

> GroupMapping TeamsListIdpGroupsForLegacy(ctx, teamId).Execute()

List IdP groups for a team (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListIdpGroupsForLegacy(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListIdpGroupsForLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListIdpGroupsForLegacy`: GroupMapping
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListIdpGroupsForLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListIdpGroupsForLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListIdpGroupsForOrg

> GroupMapping TeamsListIdpGroupsForOrg(ctx, org).PerPage(perPage).Page(page).Execute()

List IdP groups for an organization



### Example

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
    page := "page_example" // string | Page token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListIdpGroupsForOrg(context.Background(), org).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListIdpGroupsForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListIdpGroupsForOrg`: GroupMapping
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListIdpGroupsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListIdpGroupsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **string** | Page token | 

### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListIdpGroupsInOrg

> GroupMapping TeamsListIdpGroupsInOrg(ctx, org, teamSlug).Execute()

List IdP groups for a team



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsListIdpGroupsInOrg(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListIdpGroupsInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListIdpGroupsInOrg`: GroupMapping
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListIdpGroupsInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListIdpGroupsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupMapping**](GroupMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListLinkedExternalIdpGroupsToTeamForOrg

> ExternalGroups TeamsListLinkedExternalIdpGroupsToTeamForOrg(ctx, org, teamSlug).Execute()

List a connection between an external group and a team



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsListLinkedExternalIdpGroupsToTeamForOrg(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListLinkedExternalIdpGroupsToTeamForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListLinkedExternalIdpGroupsToTeamForOrg`: ExternalGroups
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListLinkedExternalIdpGroupsToTeamForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListLinkedExternalIdpGroupsToTeamForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExternalGroups**](ExternalGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListMembersInOrg

> []SimpleUser TeamsListMembersInOrg(ctx, org, teamSlug).Role(role).PerPage(perPage).Page(page).Execute()

List team members



### Example

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
    role := "role_example" // string | Filters members returned by their role in the team. (optional) (default to "all")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListMembersInOrg(context.Background(), org, teamSlug).Role(role).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListMembersInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListMembersInOrg`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListMembersInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListMembersInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | **string** | Filters members returned by their role in the team. | [default to &quot;all&quot;]
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


## TeamsListMembersLegacy

> []SimpleUser TeamsListMembersLegacy(ctx, teamId).Role(role).PerPage(perPage).Page(page).Execute()

List team members (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    role := "role_example" // string | Filters members returned by their role in the team. (optional) (default to "all")
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListMembersLegacy(context.Background(), teamId).Role(role).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListMembersLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListMembersLegacy`: []SimpleUser
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListMembersLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListMembersLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** | Filters members returned by their role in the team. | [default to &quot;all&quot;]
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


## TeamsListPendingInvitationsInOrg

> []OrganizationInvitation TeamsListPendingInvitationsInOrg(ctx, org, teamSlug).PerPage(perPage).Page(page).Execute()

List pending team invitations



### Example

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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListPendingInvitationsInOrg(context.Background(), org, teamSlug).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListPendingInvitationsInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListPendingInvitationsInOrg`: []OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListPendingInvitationsInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListPendingInvitationsInOrgRequest struct via the builder pattern


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


## TeamsListPendingInvitationsLegacy

> []OrganizationInvitation TeamsListPendingInvitationsLegacy(ctx, teamId).PerPage(perPage).Page(page).Execute()

List pending team invitations (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListPendingInvitationsLegacy(context.Background(), teamId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListPendingInvitationsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListPendingInvitationsLegacy`: []OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListPendingInvitationsLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListPendingInvitationsLegacyRequest struct via the builder pattern


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


## TeamsListProjectsInOrg

> []TeamProject TeamsListProjectsInOrg(ctx, org, teamSlug).PerPage(perPage).Page(page).Execute()

List team projects



### Example

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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListProjectsInOrg(context.Background(), org, teamSlug).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListProjectsInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListProjectsInOrg`: []TeamProject
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListProjectsInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListProjectsInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TeamProject**](TeamProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListProjectsLegacy

> []TeamProject TeamsListProjectsLegacy(ctx, teamId).PerPage(perPage).Page(page).Execute()

List team projects (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListProjectsLegacy(context.Background(), teamId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListProjectsLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListProjectsLegacy`: []TeamProject
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListProjectsLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListProjectsLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of results per page (max 100). | [default to 30]
 **page** | **int32** | Page number of the results to fetch. | [default to 1]

### Return type

[**[]TeamProject**](TeamProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListReposInOrg

> []MinimalRepository TeamsListReposInOrg(ctx, org, teamSlug).PerPage(perPage).Page(page).Execute()

List team repositories



### Example

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
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListReposInOrg(context.Background(), org, teamSlug).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListReposInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListReposInOrg`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListReposInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListReposInOrgRequest struct via the builder pattern


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


## TeamsListReposLegacy

> []MinimalRepository TeamsListReposLegacy(ctx, teamId).PerPage(perPage).Page(page).Execute()

List team repositories (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    perPage := int32(56) // int32 | The number of results per page (max 100). (optional) (default to 30)
    page := int32(56) // int32 | Page number of the results to fetch. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsListReposLegacy(context.Background(), teamId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsListReposLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListReposLegacy`: []MinimalRepository
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsListReposLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListReposLegacyRequest struct via the builder pattern


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


## TeamsRemoveMemberLegacy

> TeamsRemoveMemberLegacy(ctx, teamId, username).Execute()

Remove team member (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveMemberLegacy(context.Background(), teamId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveMemberLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveMemberLegacyRequest struct via the builder pattern


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


## TeamsRemoveMembershipForUserInOrg

> TeamsRemoveMembershipForUserInOrg(ctx, org, teamSlug, username).Execute()

Remove team membership for a user



### Example

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
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveMembershipForUserInOrg(context.Background(), org, teamSlug, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveMembershipForUserInOrg``: %v\n", err)
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
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveMembershipForUserInOrgRequest struct via the builder pattern


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


## TeamsRemoveMembershipForUserLegacy

> TeamsRemoveMembershipForUserLegacy(ctx, teamId, username).Execute()

Remove team membership for a user (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    username := "username_example" // string | The handle for the GitHub user account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveMembershipForUserLegacy(context.Background(), teamId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveMembershipForUserLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**username** | **string** | The handle for the GitHub user account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveMembershipForUserLegacyRequest struct via the builder pattern


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


## TeamsRemoveProjectInOrg

> TeamsRemoveProjectInOrg(ctx, org, teamSlug, projectId).Execute()

Remove a project from a team



### Example

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
    projectId := int32(56) // int32 | The unique identifier of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveProjectInOrg(context.Background(), org, teamSlug, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveProjectInOrg``: %v\n", err)
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
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveProjectInOrgRequest struct via the builder pattern


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


## TeamsRemoveProjectLegacy

> TeamsRemoveProjectLegacy(ctx, teamId, projectId).Execute()

Remove a project from a team (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    projectId := int32(56) // int32 | The unique identifier of the project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveProjectLegacy(context.Background(), teamId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveProjectLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**projectId** | **int32** | The unique identifier of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveProjectLegacyRequest struct via the builder pattern


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


## TeamsRemoveRepoInOrg

> TeamsRemoveRepoInOrg(ctx, org, teamSlug, owner, repo).Execute()

Remove a repository from a team



### Example

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
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveRepoInOrg(context.Background(), org, teamSlug, owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveRepoInOrg``: %v\n", err)
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
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveRepoInOrgRequest struct via the builder pattern


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


## TeamsRemoveRepoLegacy

> TeamsRemoveRepoLegacy(ctx, teamId, owner, repo).Execute()

Remove a repository from a team (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    owner := "owner_example" // string | The account owner of the repository. The name is not case sensitive.
    repo := "repo_example" // string | The name of the repository. The name is not case sensitive.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsRemoveRepoLegacy(context.Background(), teamId, owner, repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsRemoveRepoLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**owner** | **string** | The account owner of the repository. The name is not case sensitive. | 
**repo** | **string** | The name of the repository. The name is not case sensitive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsRemoveRepoLegacyRequest struct via the builder pattern


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


## TeamsUnlinkExternalIdpGroupFromTeamForOrg

> TeamsUnlinkExternalIdpGroupFromTeamForOrg(ctx, org, teamSlug).Execute()

Remove the connection between an external group and a team



### Example

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
    resp, r, err := apiClient.TeamsApi.TeamsUnlinkExternalIdpGroupFromTeamForOrg(context.Background(), org, teamSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUnlinkExternalIdpGroupFromTeamForOrg``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsUnlinkExternalIdpGroupFromTeamForOrgRequest struct via the builder pattern


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


## TeamsUpdateDiscussionCommentInOrg

> TeamDiscussionComment TeamsUpdateDiscussionCommentInOrg(ctx, org, teamSlug, discussionNumber, commentNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()

Update a discussion comment



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    teamsCreateDiscussionCommentInOrgRequest := *openapiclient.NewTeamsCreateDiscussionCommentInOrgRequest("Body_example") // TeamsCreateDiscussionCommentInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdateDiscussionCommentInOrg(context.Background(), org, teamSlug, discussionNumber, commentNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdateDiscussionCommentInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdateDiscussionCommentInOrg`: TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdateDiscussionCommentInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateDiscussionCommentInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **teamsCreateDiscussionCommentInOrgRequest** | [**TeamsCreateDiscussionCommentInOrgRequest**](TeamsCreateDiscussionCommentInOrgRequest.md) |  | 

### Return type

[**TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateDiscussionCommentLegacy

> TeamDiscussionComment TeamsUpdateDiscussionCommentLegacy(ctx, teamId, discussionNumber, commentNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()

Update a discussion comment (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    commentNumber := int32(56) // int32 | The number that identifies the comment.
    teamsCreateDiscussionCommentInOrgRequest := *openapiclient.NewTeamsCreateDiscussionCommentInOrgRequest("Body_example") // TeamsCreateDiscussionCommentInOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdateDiscussionCommentLegacy(context.Background(), teamId, discussionNumber, commentNumber).TeamsCreateDiscussionCommentInOrgRequest(teamsCreateDiscussionCommentInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdateDiscussionCommentLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdateDiscussionCommentLegacy`: TeamDiscussionComment
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdateDiscussionCommentLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 
**commentNumber** | **int32** | The number that identifies the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateDiscussionCommentLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamsCreateDiscussionCommentInOrgRequest** | [**TeamsCreateDiscussionCommentInOrgRequest**](TeamsCreateDiscussionCommentInOrgRequest.md) |  | 

### Return type

[**TeamDiscussionComment**](TeamDiscussionComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateDiscussionInOrg

> TeamDiscussion TeamsUpdateDiscussionInOrg(ctx, org, teamSlug, discussionNumber).TeamsUpdateDiscussionInOrgRequest(teamsUpdateDiscussionInOrgRequest).Execute()

Update a discussion



### Example

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
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    teamsUpdateDiscussionInOrgRequest := *openapiclient.NewTeamsUpdateDiscussionInOrgRequest() // TeamsUpdateDiscussionInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdateDiscussionInOrg(context.Background(), org, teamSlug, discussionNumber).TeamsUpdateDiscussionInOrgRequest(teamsUpdateDiscussionInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdateDiscussionInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdateDiscussionInOrg`: TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdateDiscussionInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateDiscussionInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamsUpdateDiscussionInOrgRequest** | [**TeamsUpdateDiscussionInOrgRequest**](TeamsUpdateDiscussionInOrgRequest.md) |  | 

### Return type

[**TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateDiscussionLegacy

> TeamDiscussion TeamsUpdateDiscussionLegacy(ctx, teamId, discussionNumber).TeamsUpdateDiscussionInOrgRequest(teamsUpdateDiscussionInOrgRequest).Execute()

Update a discussion (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    discussionNumber := int32(56) // int32 | The number that identifies the discussion.
    teamsUpdateDiscussionInOrgRequest := *openapiclient.NewTeamsUpdateDiscussionInOrgRequest() // TeamsUpdateDiscussionInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdateDiscussionLegacy(context.Background(), teamId, discussionNumber).TeamsUpdateDiscussionInOrgRequest(teamsUpdateDiscussionInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdateDiscussionLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdateDiscussionLegacy`: TeamDiscussion
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdateDiscussionLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 
**discussionNumber** | **int32** | The number that identifies the discussion. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateDiscussionLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsUpdateDiscussionInOrgRequest** | [**TeamsUpdateDiscussionInOrgRequest**](TeamsUpdateDiscussionInOrgRequest.md) |  | 

### Return type

[**TeamDiscussion**](TeamDiscussion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateInOrg

> TeamFull TeamsUpdateInOrg(ctx, org, teamSlug).TeamsUpdateInOrgRequest(teamsUpdateInOrgRequest).Execute()

Update a team



### Example

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
    teamsUpdateInOrgRequest := *openapiclient.NewTeamsUpdateInOrgRequest() // TeamsUpdateInOrgRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdateInOrg(context.Background(), org, teamSlug).TeamsUpdateInOrgRequest(teamsUpdateInOrgRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdateInOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdateInOrg`: TeamFull
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdateInOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | The organization name. The name is not case sensitive. | 
**teamSlug** | **string** | The slug of the team name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateInOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsUpdateInOrgRequest** | [**TeamsUpdateInOrgRequest**](TeamsUpdateInOrgRequest.md) |  | 

### Return type

[**TeamFull**](TeamFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateLegacy

> TeamFull TeamsUpdateLegacy(ctx, teamId).TeamsUpdateLegacyRequest(teamsUpdateLegacyRequest).Execute()

Update a team (Legacy)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := int32(56) // int32 | The unique identifier of the team.
    teamsUpdateLegacyRequest := *openapiclient.NewTeamsUpdateLegacyRequest("Name_example") // TeamsUpdateLegacyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdateLegacy(context.Background(), teamId).TeamsUpdateLegacyRequest(teamsUpdateLegacyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdateLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdateLegacy`: TeamFull
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdateLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **int32** | The unique identifier of the team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamsUpdateLegacyRequest** | [**TeamsUpdateLegacyRequest**](TeamsUpdateLegacyRequest.md) |  | 

### Return type

[**TeamFull**](TeamFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

