# PrivateUser

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
**PrivateGists** | **int32** |  | 
**TotalPrivateRepos** | **int32** |  | 
**OwnedPrivateRepos** | **int32** |  | 
**DiskUsage** | **int32** |  | 
**Collaborators** | **int32** |  | 
**TwoFactorAuthentication** | **bool** |  | 
**Plan** | Pointer to [**PublicUserPlan**](PublicUserPlan.md) |  | [optional] 
**SuspendedAt** | Pointer to **NullableTime** |  | [optional] 
**BusinessPlus** | Pointer to **bool** |  | [optional] 
**LdapDn** | Pointer to **string** |  | [optional] 

## Methods

### NewPrivateUser

`func NewPrivateUser(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, name NullableString, company NullableString, blog NullableString, location NullableString, email NullableString, hireable NullableBool, bio NullableString, publicRepos int32, publicGists int32, followers int32, following int32, createdAt time.Time, updatedAt time.Time, privateGists int32, totalPrivateRepos int32, ownedPrivateRepos int32, diskUsage int32, collaborators int32, twoFactorAuthentication bool, ) *PrivateUser`

NewPrivateUser instantiates a new PrivateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateUserWithDefaults

`func NewPrivateUserWithDefaults() *PrivateUser`

NewPrivateUserWithDefaults instantiates a new PrivateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *PrivateUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PrivateUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PrivateUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *PrivateUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateUser) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PrivateUser) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PrivateUser) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PrivateUser) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *PrivateUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *PrivateUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *PrivateUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *PrivateUser) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *PrivateUser) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *PrivateUser) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *PrivateUser) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *PrivateUser) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *PrivateUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PrivateUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PrivateUser) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *PrivateUser) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PrivateUser) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PrivateUser) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *PrivateUser) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *PrivateUser) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *PrivateUser) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetFollowingUrl

`func (o *PrivateUser) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *PrivateUser) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *PrivateUser) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *PrivateUser) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *PrivateUser) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *PrivateUser) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *PrivateUser) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *PrivateUser) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *PrivateUser) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetSubscriptionsUrl

`func (o *PrivateUser) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *PrivateUser) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *PrivateUser) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *PrivateUser) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *PrivateUser) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *PrivateUser) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *PrivateUser) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *PrivateUser) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *PrivateUser) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *PrivateUser) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *PrivateUser) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *PrivateUser) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetReceivedEventsUrl

`func (o *PrivateUser) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *PrivateUser) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *PrivateUser) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *PrivateUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateUser) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *PrivateUser) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *PrivateUser) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *PrivateUser) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetName

`func (o *PrivateUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateUser) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PrivateUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PrivateUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *PrivateUser) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PrivateUser) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PrivateUser) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *PrivateUser) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PrivateUser) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetBlog

`func (o *PrivateUser) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *PrivateUser) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *PrivateUser) SetBlog(v string)`

SetBlog sets Blog field to given value.


### SetBlogNil

`func (o *PrivateUser) SetBlogNil(b bool)`

 SetBlogNil sets the value for Blog to be an explicit nil

### UnsetBlog
`func (o *PrivateUser) UnsetBlog()`

UnsetBlog ensures that no value is present for Blog, not even an explicit nil
### GetLocation

`func (o *PrivateUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrivateUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrivateUser) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *PrivateUser) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PrivateUser) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEmail

`func (o *PrivateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrivateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrivateUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *PrivateUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PrivateUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetHireable

`func (o *PrivateUser) GetHireable() bool`

GetHireable returns the Hireable field if non-nil, zero value otherwise.

### GetHireableOk

`func (o *PrivateUser) GetHireableOk() (*bool, bool)`

GetHireableOk returns a tuple with the Hireable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireable

`func (o *PrivateUser) SetHireable(v bool)`

SetHireable sets Hireable field to given value.


### SetHireableNil

`func (o *PrivateUser) SetHireableNil(b bool)`

 SetHireableNil sets the value for Hireable to be an explicit nil

### UnsetHireable
`func (o *PrivateUser) UnsetHireable()`

UnsetHireable ensures that no value is present for Hireable, not even an explicit nil
### GetBio

`func (o *PrivateUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *PrivateUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *PrivateUser) SetBio(v string)`

SetBio sets Bio field to given value.


### SetBioNil

`func (o *PrivateUser) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *PrivateUser) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetTwitterUsername

`func (o *PrivateUser) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *PrivateUser) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *PrivateUser) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *PrivateUser) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *PrivateUser) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *PrivateUser) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetPublicRepos

`func (o *PrivateUser) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *PrivateUser) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *PrivateUser) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.


### GetPublicGists

`func (o *PrivateUser) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *PrivateUser) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *PrivateUser) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.


### GetFollowers

`func (o *PrivateUser) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PrivateUser) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PrivateUser) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.


### GetFollowing

`func (o *PrivateUser) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *PrivateUser) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *PrivateUser) SetFollowing(v int32)`

SetFollowing sets Following field to given value.


### GetCreatedAt

`func (o *PrivateUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PrivateUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrivateUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrivateUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPrivateGists

`func (o *PrivateUser) GetPrivateGists() int32`

GetPrivateGists returns the PrivateGists field if non-nil, zero value otherwise.

### GetPrivateGistsOk

`func (o *PrivateUser) GetPrivateGistsOk() (*int32, bool)`

GetPrivateGistsOk returns a tuple with the PrivateGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateGists

`func (o *PrivateUser) SetPrivateGists(v int32)`

SetPrivateGists sets PrivateGists field to given value.


### GetTotalPrivateRepos

`func (o *PrivateUser) GetTotalPrivateRepos() int32`

GetTotalPrivateRepos returns the TotalPrivateRepos field if non-nil, zero value otherwise.

### GetTotalPrivateReposOk

`func (o *PrivateUser) GetTotalPrivateReposOk() (*int32, bool)`

GetTotalPrivateReposOk returns a tuple with the TotalPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateRepos

`func (o *PrivateUser) SetTotalPrivateRepos(v int32)`

SetTotalPrivateRepos sets TotalPrivateRepos field to given value.


### GetOwnedPrivateRepos

`func (o *PrivateUser) GetOwnedPrivateRepos() int32`

GetOwnedPrivateRepos returns the OwnedPrivateRepos field if non-nil, zero value otherwise.

### GetOwnedPrivateReposOk

`func (o *PrivateUser) GetOwnedPrivateReposOk() (*int32, bool)`

GetOwnedPrivateReposOk returns a tuple with the OwnedPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedPrivateRepos

`func (o *PrivateUser) SetOwnedPrivateRepos(v int32)`

SetOwnedPrivateRepos sets OwnedPrivateRepos field to given value.


### GetDiskUsage

`func (o *PrivateUser) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *PrivateUser) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *PrivateUser) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.


### GetCollaborators

`func (o *PrivateUser) GetCollaborators() int32`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *PrivateUser) GetCollaboratorsOk() (*int32, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *PrivateUser) SetCollaborators(v int32)`

SetCollaborators sets Collaborators field to given value.


### GetTwoFactorAuthentication

`func (o *PrivateUser) GetTwoFactorAuthentication() bool`

GetTwoFactorAuthentication returns the TwoFactorAuthentication field if non-nil, zero value otherwise.

### GetTwoFactorAuthenticationOk

`func (o *PrivateUser) GetTwoFactorAuthenticationOk() (*bool, bool)`

GetTwoFactorAuthenticationOk returns a tuple with the TwoFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthentication

`func (o *PrivateUser) SetTwoFactorAuthentication(v bool)`

SetTwoFactorAuthentication sets TwoFactorAuthentication field to given value.


### GetPlan

`func (o *PrivateUser) GetPlan() PublicUserPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PrivateUser) GetPlanOk() (*PublicUserPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PrivateUser) SetPlan(v PublicUserPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PrivateUser) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *PrivateUser) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *PrivateUser) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *PrivateUser) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *PrivateUser) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### SetSuspendedAtNil

`func (o *PrivateUser) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *PrivateUser) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetBusinessPlus

`func (o *PrivateUser) GetBusinessPlus() bool`

GetBusinessPlus returns the BusinessPlus field if non-nil, zero value otherwise.

### GetBusinessPlusOk

`func (o *PrivateUser) GetBusinessPlusOk() (*bool, bool)`

GetBusinessPlusOk returns a tuple with the BusinessPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPlus

`func (o *PrivateUser) SetBusinessPlus(v bool)`

SetBusinessPlus sets BusinessPlus field to given value.

### HasBusinessPlus

`func (o *PrivateUser) HasBusinessPlus() bool`

HasBusinessPlus returns a boolean if a field has been set.

### GetLdapDn

`func (o *PrivateUser) GetLdapDn() string`

GetLdapDn returns the LdapDn field if non-nil, zero value otherwise.

### GetLdapDnOk

`func (o *PrivateUser) GetLdapDnOk() (*string, bool)`

GetLdapDnOk returns a tuple with the LdapDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDn

`func (o *PrivateUser) SetLdapDn(v string)`

SetLdapDn sets LdapDn field to given value.

### HasLdapDn

`func (o *PrivateUser) HasLdapDn() bool`

HasLdapDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


