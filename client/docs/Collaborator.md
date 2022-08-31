# Collaborator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Id** | **int32** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
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
**Permissions** | Pointer to [**CollaboratorPermissions**](CollaboratorPermissions.md) |  | [optional] 
**RoleName** | **string** |  | 

## Methods

### NewCollaborator

`func NewCollaborator(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, roleName string, ) *Collaborator`

NewCollaborator instantiates a new Collaborator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollaboratorWithDefaults

`func NewCollaboratorWithDefaults() *Collaborator`

NewCollaboratorWithDefaults instantiates a new Collaborator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *Collaborator) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Collaborator) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Collaborator) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *Collaborator) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collaborator) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collaborator) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *Collaborator) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Collaborator) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Collaborator) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Collaborator) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Collaborator) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Collaborator) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *Collaborator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collaborator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collaborator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Collaborator) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Collaborator) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Collaborator) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeId

`func (o *Collaborator) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Collaborator) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Collaborator) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *Collaborator) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Collaborator) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Collaborator) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *Collaborator) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *Collaborator) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *Collaborator) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *Collaborator) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *Collaborator) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *Collaborator) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Collaborator) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Collaborator) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Collaborator) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Collaborator) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Collaborator) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *Collaborator) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *Collaborator) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *Collaborator) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetFollowingUrl

`func (o *Collaborator) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *Collaborator) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *Collaborator) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *Collaborator) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *Collaborator) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *Collaborator) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *Collaborator) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *Collaborator) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *Collaborator) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetSubscriptionsUrl

`func (o *Collaborator) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *Collaborator) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *Collaborator) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *Collaborator) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *Collaborator) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *Collaborator) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *Collaborator) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *Collaborator) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *Collaborator) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *Collaborator) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *Collaborator) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *Collaborator) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetReceivedEventsUrl

`func (o *Collaborator) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *Collaborator) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *Collaborator) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *Collaborator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Collaborator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Collaborator) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *Collaborator) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *Collaborator) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *Collaborator) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetPermissions

`func (o *Collaborator) GetPermissions() CollaboratorPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Collaborator) GetPermissionsOk() (*CollaboratorPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Collaborator) SetPermissions(v CollaboratorPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Collaborator) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRoleName

`func (o *Collaborator) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Collaborator) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Collaborator) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


