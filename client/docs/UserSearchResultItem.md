# UserSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**AvatarUrl** | **string** |  | 
**GravatarId** | **NullableString** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**FollowersUrl** | **string** |  | 
**SubscriptionsUrl** | **string** |  | 
**OrganizationsUrl** | **string** |  | 
**ReposUrl** | **string** |  | 
**ReceivedEventsUrl** | **string** |  | 
**Type** | **string** |  | 
**Score** | **float32** |  | 
**FollowingUrl** | **string** |  | 
**GistsUrl** | **string** |  | 
**StarredUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**PublicRepos** | Pointer to **int32** |  | [optional] 
**PublicGists** | Pointer to **int32** |  | [optional] 
**Followers** | Pointer to **int32** |  | [optional] 
**Following** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Bio** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**SiteAdmin** | **bool** |  | 
**Hireable** | Pointer to **NullableBool** |  | [optional] 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 
**Blog** | Pointer to **NullableString** |  | [optional] 
**Company** | Pointer to **NullableString** |  | [optional] 
**SuspendedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewUserSearchResultItem

`func NewUserSearchResultItem(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, receivedEventsUrl string, type_ string, score float32, followingUrl string, gistsUrl string, starredUrl string, eventsUrl string, siteAdmin bool, ) *UserSearchResultItem`

NewUserSearchResultItem instantiates a new UserSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSearchResultItemWithDefaults

`func NewUserSearchResultItemWithDefaults() *UserSearchResultItem`

NewUserSearchResultItemWithDefaults instantiates a new UserSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *UserSearchResultItem) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserSearchResultItem) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserSearchResultItem) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *UserSearchResultItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSearchResultItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSearchResultItem) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *UserSearchResultItem) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *UserSearchResultItem) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *UserSearchResultItem) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *UserSearchResultItem) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserSearchResultItem) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserSearchResultItem) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *UserSearchResultItem) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *UserSearchResultItem) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *UserSearchResultItem) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *UserSearchResultItem) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *UserSearchResultItem) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *UserSearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *UserSearchResultItem) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *UserSearchResultItem) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *UserSearchResultItem) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *UserSearchResultItem) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *UserSearchResultItem) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *UserSearchResultItem) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetSubscriptionsUrl

`func (o *UserSearchResultItem) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *UserSearchResultItem) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *UserSearchResultItem) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *UserSearchResultItem) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *UserSearchResultItem) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *UserSearchResultItem) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *UserSearchResultItem) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *UserSearchResultItem) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *UserSearchResultItem) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetReceivedEventsUrl

`func (o *UserSearchResultItem) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *UserSearchResultItem) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *UserSearchResultItem) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *UserSearchResultItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSearchResultItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSearchResultItem) SetType(v string)`

SetType sets Type field to given value.


### GetScore

`func (o *UserSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UserSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UserSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetFollowingUrl

`func (o *UserSearchResultItem) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *UserSearchResultItem) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *UserSearchResultItem) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *UserSearchResultItem) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *UserSearchResultItem) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *UserSearchResultItem) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *UserSearchResultItem) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *UserSearchResultItem) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *UserSearchResultItem) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetEventsUrl

`func (o *UserSearchResultItem) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *UserSearchResultItem) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *UserSearchResultItem) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetPublicRepos

`func (o *UserSearchResultItem) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *UserSearchResultItem) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *UserSearchResultItem) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.

### HasPublicRepos

`func (o *UserSearchResultItem) HasPublicRepos() bool`

HasPublicRepos returns a boolean if a field has been set.

### GetPublicGists

`func (o *UserSearchResultItem) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *UserSearchResultItem) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *UserSearchResultItem) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.

### HasPublicGists

`func (o *UserSearchResultItem) HasPublicGists() bool`

HasPublicGists returns a boolean if a field has been set.

### GetFollowers

`func (o *UserSearchResultItem) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *UserSearchResultItem) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *UserSearchResultItem) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *UserSearchResultItem) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetFollowing

`func (o *UserSearchResultItem) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *UserSearchResultItem) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *UserSearchResultItem) SetFollowing(v int32)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *UserSearchResultItem) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserSearchResultItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserSearchResultItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserSearchResultItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserSearchResultItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserSearchResultItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserSearchResultItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserSearchResultItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserSearchResultItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *UserSearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSearchResultItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSearchResultItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserSearchResultItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserSearchResultItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBio

`func (o *UserSearchResultItem) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserSearchResultItem) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserSearchResultItem) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UserSearchResultItem) HasBio() bool`

HasBio returns a boolean if a field has been set.

### SetBioNil

`func (o *UserSearchResultItem) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *UserSearchResultItem) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetEmail

`func (o *UserSearchResultItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSearchResultItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSearchResultItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSearchResultItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserSearchResultItem) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserSearchResultItem) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetLocation

`func (o *UserSearchResultItem) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UserSearchResultItem) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UserSearchResultItem) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UserSearchResultItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *UserSearchResultItem) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *UserSearchResultItem) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSiteAdmin

`func (o *UserSearchResultItem) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *UserSearchResultItem) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *UserSearchResultItem) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetHireable

`func (o *UserSearchResultItem) GetHireable() bool`

GetHireable returns the Hireable field if non-nil, zero value otherwise.

### GetHireableOk

`func (o *UserSearchResultItem) GetHireableOk() (*bool, bool)`

GetHireableOk returns a tuple with the Hireable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireable

`func (o *UserSearchResultItem) SetHireable(v bool)`

SetHireable sets Hireable field to given value.

### HasHireable

`func (o *UserSearchResultItem) HasHireable() bool`

HasHireable returns a boolean if a field has been set.

### SetHireableNil

`func (o *UserSearchResultItem) SetHireableNil(b bool)`

 SetHireableNil sets the value for Hireable to be an explicit nil

### UnsetHireable
`func (o *UserSearchResultItem) UnsetHireable()`

UnsetHireable ensures that no value is present for Hireable, not even an explicit nil
### GetTextMatches

`func (o *UserSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *UserSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *UserSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *UserSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.

### GetBlog

`func (o *UserSearchResultItem) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *UserSearchResultItem) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *UserSearchResultItem) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *UserSearchResultItem) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### SetBlogNil

`func (o *UserSearchResultItem) SetBlogNil(b bool)`

 SetBlogNil sets the value for Blog to be an explicit nil

### UnsetBlog
`func (o *UserSearchResultItem) UnsetBlog()`

UnsetBlog ensures that no value is present for Blog, not even an explicit nil
### GetCompany

`func (o *UserSearchResultItem) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserSearchResultItem) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserSearchResultItem) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserSearchResultItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UserSearchResultItem) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UserSearchResultItem) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetSuspendedAt

`func (o *UserSearchResultItem) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *UserSearchResultItem) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *UserSearchResultItem) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *UserSearchResultItem) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### SetSuspendedAtNil

`func (o *UserSearchResultItem) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *UserSearchResultItem) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


