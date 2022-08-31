# NullableSimpleUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
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

## Methods

### NewNullableSimpleUser

`func NewNullableSimpleUser(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, ) *NullableSimpleUser`

NewNullableSimpleUser instantiates a new NullableSimpleUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableSimpleUserWithDefaults

`func NewNullableSimpleUserWithDefaults() *NullableSimpleUser`

NewNullableSimpleUserWithDefaults instantiates a new NullableSimpleUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NullableSimpleUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NullableSimpleUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NullableSimpleUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NullableSimpleUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NullableSimpleUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NullableSimpleUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *NullableSimpleUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NullableSimpleUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NullableSimpleUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NullableSimpleUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NullableSimpleUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NullableSimpleUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetLogin

`func (o *NullableSimpleUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *NullableSimpleUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *NullableSimpleUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *NullableSimpleUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableSimpleUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableSimpleUser) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *NullableSimpleUser) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NullableSimpleUser) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NullableSimpleUser) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *NullableSimpleUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *NullableSimpleUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *NullableSimpleUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *NullableSimpleUser) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *NullableSimpleUser) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *NullableSimpleUser) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *NullableSimpleUser) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *NullableSimpleUser) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *NullableSimpleUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NullableSimpleUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NullableSimpleUser) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *NullableSimpleUser) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *NullableSimpleUser) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *NullableSimpleUser) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *NullableSimpleUser) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *NullableSimpleUser) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *NullableSimpleUser) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetFollowingUrl

`func (o *NullableSimpleUser) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *NullableSimpleUser) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *NullableSimpleUser) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *NullableSimpleUser) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *NullableSimpleUser) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *NullableSimpleUser) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *NullableSimpleUser) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *NullableSimpleUser) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *NullableSimpleUser) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetSubscriptionsUrl

`func (o *NullableSimpleUser) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *NullableSimpleUser) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *NullableSimpleUser) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *NullableSimpleUser) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *NullableSimpleUser) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *NullableSimpleUser) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *NullableSimpleUser) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *NullableSimpleUser) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *NullableSimpleUser) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *NullableSimpleUser) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *NullableSimpleUser) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *NullableSimpleUser) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetReceivedEventsUrl

`func (o *NullableSimpleUser) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *NullableSimpleUser) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *NullableSimpleUser) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *NullableSimpleUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NullableSimpleUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NullableSimpleUser) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *NullableSimpleUser) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *NullableSimpleUser) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *NullableSimpleUser) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetStarredAt

`func (o *NullableSimpleUser) GetStarredAt() string`

GetStarredAt returns the StarredAt field if non-nil, zero value otherwise.

### GetStarredAtOk

`func (o *NullableSimpleUser) GetStarredAtOk() (*string, bool)`

GetStarredAtOk returns a tuple with the StarredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredAt

`func (o *NullableSimpleUser) SetStarredAt(v string)`

SetStarredAt sets StarredAt field to given value.

### HasStarredAt

`func (o *NullableSimpleUser) HasStarredAt() bool`

HasStarredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


