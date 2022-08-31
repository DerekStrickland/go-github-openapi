# PublicUser

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
**Name** | **NullableString** |  | 
**Company** | **NullableString** |  | 
**Blog** | **NullableString** |  | 
**Location** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**Hireable** | **NullableBool** |  | 
**Bio** | **NullableString** |  | 
**TwitterUsername** | Pointer to **NullableString** |  | [optional] 
**PublicRepos** | **int32** |  | 
**PublicGists** | **int32** |  | 
**Followers** | **int32** |  | 
**Following** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Plan** | Pointer to [**PublicUserPlan**](PublicUserPlan.md) |  | [optional] 
**SuspendedAt** | Pointer to **NullableTime** |  | [optional] 
**PrivateGists** | Pointer to **int32** |  | [optional] 
**TotalPrivateRepos** | Pointer to **int32** |  | [optional] 
**OwnedPrivateRepos** | Pointer to **int32** |  | [optional] 
**DiskUsage** | Pointer to **int32** |  | [optional] 
**Collaborators** | Pointer to **int32** |  | [optional] 

## Methods

### NewPublicUser

`func NewPublicUser(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, name NullableString, company NullableString, blog NullableString, location NullableString, email NullableString, hireable NullableBool, bio NullableString, publicRepos int32, publicGists int32, followers int32, following int32, createdAt time.Time, updatedAt time.Time, ) *PublicUser`

NewPublicUser instantiates a new PublicUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUserWithDefaults

`func NewPublicUserWithDefaults() *PublicUser`

NewPublicUserWithDefaults instantiates a new PublicUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *PublicUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PublicUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PublicUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *PublicUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicUser) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PublicUser) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PublicUser) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PublicUser) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *PublicUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PublicUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PublicUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *PublicUser) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *PublicUser) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *PublicUser) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *PublicUser) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *PublicUser) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *PublicUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicUser) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *PublicUser) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PublicUser) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PublicUser) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *PublicUser) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *PublicUser) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *PublicUser) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetFollowingUrl

`func (o *PublicUser) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *PublicUser) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *PublicUser) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *PublicUser) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *PublicUser) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *PublicUser) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *PublicUser) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *PublicUser) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *PublicUser) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetSubscriptionsUrl

`func (o *PublicUser) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *PublicUser) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *PublicUser) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *PublicUser) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *PublicUser) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *PublicUser) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *PublicUser) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *PublicUser) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *PublicUser) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *PublicUser) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *PublicUser) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *PublicUser) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetReceivedEventsUrl

`func (o *PublicUser) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *PublicUser) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *PublicUser) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *PublicUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicUser) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *PublicUser) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *PublicUser) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *PublicUser) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetName

`func (o *PublicUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicUser) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PublicUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PublicUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *PublicUser) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PublicUser) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PublicUser) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *PublicUser) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PublicUser) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetBlog

`func (o *PublicUser) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *PublicUser) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *PublicUser) SetBlog(v string)`

SetBlog sets Blog field to given value.


### SetBlogNil

`func (o *PublicUser) SetBlogNil(b bool)`

 SetBlogNil sets the value for Blog to be an explicit nil

### UnsetBlog
`func (o *PublicUser) UnsetBlog()`

UnsetBlog ensures that no value is present for Blog, not even an explicit nil
### GetLocation

`func (o *PublicUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PublicUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PublicUser) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *PublicUser) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PublicUser) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEmail

`func (o *PublicUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *PublicUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PublicUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetHireable

`func (o *PublicUser) GetHireable() bool`

GetHireable returns the Hireable field if non-nil, zero value otherwise.

### GetHireableOk

`func (o *PublicUser) GetHireableOk() (*bool, bool)`

GetHireableOk returns a tuple with the Hireable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireable

`func (o *PublicUser) SetHireable(v bool)`

SetHireable sets Hireable field to given value.


### SetHireableNil

`func (o *PublicUser) SetHireableNil(b bool)`

 SetHireableNil sets the value for Hireable to be an explicit nil

### UnsetHireable
`func (o *PublicUser) UnsetHireable()`

UnsetHireable ensures that no value is present for Hireable, not even an explicit nil
### GetBio

`func (o *PublicUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *PublicUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *PublicUser) SetBio(v string)`

SetBio sets Bio field to given value.


### SetBioNil

`func (o *PublicUser) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *PublicUser) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetTwitterUsername

`func (o *PublicUser) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *PublicUser) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *PublicUser) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *PublicUser) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *PublicUser) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *PublicUser) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetPublicRepos

`func (o *PublicUser) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *PublicUser) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *PublicUser) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.


### GetPublicGists

`func (o *PublicUser) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *PublicUser) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *PublicUser) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.


### GetFollowers

`func (o *PublicUser) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PublicUser) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PublicUser) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.


### GetFollowing

`func (o *PublicUser) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *PublicUser) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *PublicUser) SetFollowing(v int32)`

SetFollowing sets Following field to given value.


### GetCreatedAt

`func (o *PublicUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PublicUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPlan

`func (o *PublicUser) GetPlan() PublicUserPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PublicUser) GetPlanOk() (*PublicUserPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PublicUser) SetPlan(v PublicUserPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PublicUser) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *PublicUser) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *PublicUser) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *PublicUser) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *PublicUser) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### SetSuspendedAtNil

`func (o *PublicUser) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *PublicUser) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetPrivateGists

`func (o *PublicUser) GetPrivateGists() int32`

GetPrivateGists returns the PrivateGists field if non-nil, zero value otherwise.

### GetPrivateGistsOk

`func (o *PublicUser) GetPrivateGistsOk() (*int32, bool)`

GetPrivateGistsOk returns a tuple with the PrivateGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateGists

`func (o *PublicUser) SetPrivateGists(v int32)`

SetPrivateGists sets PrivateGists field to given value.

### HasPrivateGists

`func (o *PublicUser) HasPrivateGists() bool`

HasPrivateGists returns a boolean if a field has been set.

### GetTotalPrivateRepos

`func (o *PublicUser) GetTotalPrivateRepos() int32`

GetTotalPrivateRepos returns the TotalPrivateRepos field if non-nil, zero value otherwise.

### GetTotalPrivateReposOk

`func (o *PublicUser) GetTotalPrivateReposOk() (*int32, bool)`

GetTotalPrivateReposOk returns a tuple with the TotalPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateRepos

`func (o *PublicUser) SetTotalPrivateRepos(v int32)`

SetTotalPrivateRepos sets TotalPrivateRepos field to given value.

### HasTotalPrivateRepos

`func (o *PublicUser) HasTotalPrivateRepos() bool`

HasTotalPrivateRepos returns a boolean if a field has been set.

### GetOwnedPrivateRepos

`func (o *PublicUser) GetOwnedPrivateRepos() int32`

GetOwnedPrivateRepos returns the OwnedPrivateRepos field if non-nil, zero value otherwise.

### GetOwnedPrivateReposOk

`func (o *PublicUser) GetOwnedPrivateReposOk() (*int32, bool)`

GetOwnedPrivateReposOk returns a tuple with the OwnedPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedPrivateRepos

`func (o *PublicUser) SetOwnedPrivateRepos(v int32)`

SetOwnedPrivateRepos sets OwnedPrivateRepos field to given value.

### HasOwnedPrivateRepos

`func (o *PublicUser) HasOwnedPrivateRepos() bool`

HasOwnedPrivateRepos returns a boolean if a field has been set.

### GetDiskUsage

`func (o *PublicUser) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *PublicUser) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *PublicUser) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *PublicUser) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### GetCollaborators

`func (o *PublicUser) GetCollaborators() int32`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *PublicUser) GetCollaboratorsOk() (*int32, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *PublicUser) SetCollaborators(v int32)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *PublicUser) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


