# PendingDeploymentReviewersInnerReviewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Login** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**AvatarUrl** | **string** |  | 
**GravatarId** | **NullableString** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**FollowersUrl** | **string** |  | 
**FollowingUrl** | **string** |  | 
**GistsUrl** | **string** |  | 
**StarredUrl** | **string** |  | 
**SubscriptionsUrl** | **string** |  | 
**OrganizationsUrl** | **string** |  | 
**ReposUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**ReceivedEventsUrl** | **string** |  | 
**Type** | **string** |  | 
**SiteAdmin** | **bool** |  | 
**StarredAt** | Pointer to **string** |  | [optional] 
**Slug** | **string** |  | 
**Description** | **NullableString** |  | 
**Privacy** | Pointer to **string** |  | [optional] 
**Permission** | **string** |  | 
**Permissions** | Pointer to [**TeamPermissions**](TeamPermissions.md) |  | [optional] 
**MembersUrl** | **string** |  | 
**RepositoriesUrl** | **string** |  | 
**Parent** | [**NullableNullableTeamSimple**](NullableTeamSimple.md) |  | 

## Methods

### NewPendingDeploymentReviewersInnerReviewer

`func NewPendingDeploymentReviewersInnerReviewer(name string, login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, slug string, description NullableString, permission string, membersUrl string, repositoriesUrl string, parent NullableNullableTeamSimple, ) *PendingDeploymentReviewersInnerReviewer`

NewPendingDeploymentReviewersInnerReviewer instantiates a new PendingDeploymentReviewersInnerReviewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingDeploymentReviewersInnerReviewerWithDefaults

`func NewPendingDeploymentReviewersInnerReviewerWithDefaults() *PendingDeploymentReviewersInnerReviewer`

NewPendingDeploymentReviewersInnerReviewerWithDefaults instantiates a new PendingDeploymentReviewersInnerReviewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PendingDeploymentReviewersInnerReviewer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PendingDeploymentReviewersInnerReviewer) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *PendingDeploymentReviewersInnerReviewer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PendingDeploymentReviewersInnerReviewer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PendingDeploymentReviewersInnerReviewer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PendingDeploymentReviewersInnerReviewer) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PendingDeploymentReviewersInnerReviewer) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetLogin

`func (o *PendingDeploymentReviewersInnerReviewer) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PendingDeploymentReviewersInnerReviewer) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *PendingDeploymentReviewersInnerReviewer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PendingDeploymentReviewersInnerReviewer) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PendingDeploymentReviewersInnerReviewer) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PendingDeploymentReviewersInnerReviewer) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *PendingDeploymentReviewersInnerReviewer) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *PendingDeploymentReviewersInnerReviewer) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *PendingDeploymentReviewersInnerReviewer) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *PendingDeploymentReviewersInnerReviewer) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetFollowingUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetSubscriptionsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetReceivedEventsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *PendingDeploymentReviewersInnerReviewer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingDeploymentReviewersInnerReviewer) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *PendingDeploymentReviewersInnerReviewer) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *PendingDeploymentReviewersInnerReviewer) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetStarredAt

`func (o *PendingDeploymentReviewersInnerReviewer) GetStarredAt() string`

GetStarredAt returns the StarredAt field if non-nil, zero value otherwise.

### GetStarredAtOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetStarredAtOk() (*string, bool)`

GetStarredAtOk returns a tuple with the StarredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredAt

`func (o *PendingDeploymentReviewersInnerReviewer) SetStarredAt(v string)`

SetStarredAt sets StarredAt field to given value.

### HasStarredAt

`func (o *PendingDeploymentReviewersInnerReviewer) HasStarredAt() bool`

HasStarredAt returns a boolean if a field has been set.

### GetSlug

`func (o *PendingDeploymentReviewersInnerReviewer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PendingDeploymentReviewersInnerReviewer) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *PendingDeploymentReviewersInnerReviewer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PendingDeploymentReviewersInnerReviewer) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PendingDeploymentReviewersInnerReviewer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PendingDeploymentReviewersInnerReviewer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrivacy

`func (o *PendingDeploymentReviewersInnerReviewer) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *PendingDeploymentReviewersInnerReviewer) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *PendingDeploymentReviewersInnerReviewer) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPermission

`func (o *PendingDeploymentReviewersInnerReviewer) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PendingDeploymentReviewersInnerReviewer) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetPermissions

`func (o *PendingDeploymentReviewersInnerReviewer) GetPermissions() TeamPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetPermissionsOk() (*TeamPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PendingDeploymentReviewersInnerReviewer) SetPermissions(v TeamPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PendingDeploymentReviewersInnerReviewer) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetMembersUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetRepositoriesUrl

`func (o *PendingDeploymentReviewersInnerReviewer) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *PendingDeploymentReviewersInnerReviewer) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetParent

`func (o *PendingDeploymentReviewersInnerReviewer) GetParent() NullableTeamSimple`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PendingDeploymentReviewersInnerReviewer) GetParentOk() (*NullableTeamSimple, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PendingDeploymentReviewersInnerReviewer) SetParent(v NullableTeamSimple)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *PendingDeploymentReviewersInnerReviewer) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PendingDeploymentReviewersInnerReviewer) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


