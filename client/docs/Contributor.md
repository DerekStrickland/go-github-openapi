# Contributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**GravatarId** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**FollowersUrl** | Pointer to **string** |  | [optional] 
**FollowingUrl** | Pointer to **string** |  | [optional] 
**GistsUrl** | Pointer to **string** |  | [optional] 
**StarredUrl** | Pointer to **string** |  | [optional] 
**SubscriptionsUrl** | Pointer to **string** |  | [optional] 
**OrganizationsUrl** | Pointer to **string** |  | [optional] 
**ReposUrl** | Pointer to **string** |  | [optional] 
**EventsUrl** | Pointer to **string** |  | [optional] 
**ReceivedEventsUrl** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**SiteAdmin** | Pointer to **bool** |  | [optional] 
**Contributions** | **int32** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewContributor

`func NewContributor(type_ string, contributions int32, ) *Contributor`

NewContributor instantiates a new Contributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorWithDefaults

`func NewContributorWithDefaults() *Contributor`

NewContributorWithDefaults instantiates a new Contributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *Contributor) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Contributor) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Contributor) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *Contributor) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetId

`func (o *Contributor) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contributor) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contributor) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Contributor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeId

`func (o *Contributor) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Contributor) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Contributor) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *Contributor) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *Contributor) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Contributor) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Contributor) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Contributor) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetGravatarId

`func (o *Contributor) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *Contributor) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *Contributor) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.

### HasGravatarId

`func (o *Contributor) HasGravatarId() bool`

HasGravatarId returns a boolean if a field has been set.

### SetGravatarIdNil

`func (o *Contributor) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *Contributor) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *Contributor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Contributor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Contributor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Contributor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *Contributor) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Contributor) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Contributor) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *Contributor) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetFollowersUrl

`func (o *Contributor) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *Contributor) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *Contributor) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.

### HasFollowersUrl

`func (o *Contributor) HasFollowersUrl() bool`

HasFollowersUrl returns a boolean if a field has been set.

### GetFollowingUrl

`func (o *Contributor) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *Contributor) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *Contributor) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.

### HasFollowingUrl

`func (o *Contributor) HasFollowingUrl() bool`

HasFollowingUrl returns a boolean if a field has been set.

### GetGistsUrl

`func (o *Contributor) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *Contributor) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *Contributor) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.

### HasGistsUrl

`func (o *Contributor) HasGistsUrl() bool`

HasGistsUrl returns a boolean if a field has been set.

### GetStarredUrl

`func (o *Contributor) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *Contributor) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *Contributor) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.

### HasStarredUrl

`func (o *Contributor) HasStarredUrl() bool`

HasStarredUrl returns a boolean if a field has been set.

### GetSubscriptionsUrl

`func (o *Contributor) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *Contributor) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *Contributor) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.

### HasSubscriptionsUrl

`func (o *Contributor) HasSubscriptionsUrl() bool`

HasSubscriptionsUrl returns a boolean if a field has been set.

### GetOrganizationsUrl

`func (o *Contributor) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *Contributor) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *Contributor) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.

### HasOrganizationsUrl

`func (o *Contributor) HasOrganizationsUrl() bool`

HasOrganizationsUrl returns a boolean if a field has been set.

### GetReposUrl

`func (o *Contributor) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *Contributor) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *Contributor) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.

### HasReposUrl

`func (o *Contributor) HasReposUrl() bool`

HasReposUrl returns a boolean if a field has been set.

### GetEventsUrl

`func (o *Contributor) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *Contributor) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *Contributor) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.

### HasEventsUrl

`func (o *Contributor) HasEventsUrl() bool`

HasEventsUrl returns a boolean if a field has been set.

### GetReceivedEventsUrl

`func (o *Contributor) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *Contributor) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *Contributor) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.

### HasReceivedEventsUrl

`func (o *Contributor) HasReceivedEventsUrl() bool`

HasReceivedEventsUrl returns a boolean if a field has been set.

### GetType

`func (o *Contributor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contributor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contributor) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *Contributor) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *Contributor) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *Contributor) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.

### HasSiteAdmin

`func (o *Contributor) HasSiteAdmin() bool`

HasSiteAdmin returns a boolean if a field has been set.

### GetContributions

`func (o *Contributor) GetContributions() int32`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *Contributor) GetContributionsOk() (*int32, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *Contributor) SetContributions(v int32)`

SetContributions sets Contributions field to given value.


### GetEmail

`func (o *Contributor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contributor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contributor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contributor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Contributor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contributor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contributor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Contributor) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


