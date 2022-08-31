# UsersGetAuthenticated200Response

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

### NewUsersGetAuthenticated200Response

`func NewUsersGetAuthenticated200Response(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, name NullableString, company NullableString, blog NullableString, location NullableString, email NullableString, hireable NullableBool, bio NullableString, publicRepos int32, publicGists int32, followers int32, following int32, createdAt time.Time, updatedAt time.Time, privateGists int32, totalPrivateRepos int32, ownedPrivateRepos int32, diskUsage int32, collaborators int32, twoFactorAuthentication bool, ) *UsersGetAuthenticated200Response`

NewUsersGetAuthenticated200Response instantiates a new UsersGetAuthenticated200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersGetAuthenticated200ResponseWithDefaults

`func NewUsersGetAuthenticated200ResponseWithDefaults() *UsersGetAuthenticated200Response`

NewUsersGetAuthenticated200ResponseWithDefaults instantiates a new UsersGetAuthenticated200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *UsersGetAuthenticated200Response) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UsersGetAuthenticated200Response) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UsersGetAuthenticated200Response) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *UsersGetAuthenticated200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersGetAuthenticated200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersGetAuthenticated200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *UsersGetAuthenticated200Response) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *UsersGetAuthenticated200Response) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *UsersGetAuthenticated200Response) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *UsersGetAuthenticated200Response) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UsersGetAuthenticated200Response) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UsersGetAuthenticated200Response) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *UsersGetAuthenticated200Response) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *UsersGetAuthenticated200Response) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *UsersGetAuthenticated200Response) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *UsersGetAuthenticated200Response) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *UsersGetAuthenticated200Response) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *UsersGetAuthenticated200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UsersGetAuthenticated200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UsersGetAuthenticated200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *UsersGetAuthenticated200Response) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *UsersGetAuthenticated200Response) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *UsersGetAuthenticated200Response) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetFollowersUrl

`func (o *UsersGetAuthenticated200Response) GetFollowersUrl() string`

GetFollowersUrl returns the FollowersUrl field if non-nil, zero value otherwise.

### GetFollowersUrlOk

`func (o *UsersGetAuthenticated200Response) GetFollowersUrlOk() (*string, bool)`

GetFollowersUrlOk returns a tuple with the FollowersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersUrl

`func (o *UsersGetAuthenticated200Response) SetFollowersUrl(v string)`

SetFollowersUrl sets FollowersUrl field to given value.


### GetFollowingUrl

`func (o *UsersGetAuthenticated200Response) GetFollowingUrl() string`

GetFollowingUrl returns the FollowingUrl field if non-nil, zero value otherwise.

### GetFollowingUrlOk

`func (o *UsersGetAuthenticated200Response) GetFollowingUrlOk() (*string, bool)`

GetFollowingUrlOk returns a tuple with the FollowingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUrl

`func (o *UsersGetAuthenticated200Response) SetFollowingUrl(v string)`

SetFollowingUrl sets FollowingUrl field to given value.


### GetGistsUrl

`func (o *UsersGetAuthenticated200Response) GetGistsUrl() string`

GetGistsUrl returns the GistsUrl field if non-nil, zero value otherwise.

### GetGistsUrlOk

`func (o *UsersGetAuthenticated200Response) GetGistsUrlOk() (*string, bool)`

GetGistsUrlOk returns a tuple with the GistsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGistsUrl

`func (o *UsersGetAuthenticated200Response) SetGistsUrl(v string)`

SetGistsUrl sets GistsUrl field to given value.


### GetStarredUrl

`func (o *UsersGetAuthenticated200Response) GetStarredUrl() string`

GetStarredUrl returns the StarredUrl field if non-nil, zero value otherwise.

### GetStarredUrlOk

`func (o *UsersGetAuthenticated200Response) GetStarredUrlOk() (*string, bool)`

GetStarredUrlOk returns a tuple with the StarredUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredUrl

`func (o *UsersGetAuthenticated200Response) SetStarredUrl(v string)`

SetStarredUrl sets StarredUrl field to given value.


### GetSubscriptionsUrl

`func (o *UsersGetAuthenticated200Response) GetSubscriptionsUrl() string`

GetSubscriptionsUrl returns the SubscriptionsUrl field if non-nil, zero value otherwise.

### GetSubscriptionsUrlOk

`func (o *UsersGetAuthenticated200Response) GetSubscriptionsUrlOk() (*string, bool)`

GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsUrl

`func (o *UsersGetAuthenticated200Response) SetSubscriptionsUrl(v string)`

SetSubscriptionsUrl sets SubscriptionsUrl field to given value.


### GetOrganizationsUrl

`func (o *UsersGetAuthenticated200Response) GetOrganizationsUrl() string`

GetOrganizationsUrl returns the OrganizationsUrl field if non-nil, zero value otherwise.

### GetOrganizationsUrlOk

`func (o *UsersGetAuthenticated200Response) GetOrganizationsUrlOk() (*string, bool)`

GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsUrl

`func (o *UsersGetAuthenticated200Response) SetOrganizationsUrl(v string)`

SetOrganizationsUrl sets OrganizationsUrl field to given value.


### GetReposUrl

`func (o *UsersGetAuthenticated200Response) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *UsersGetAuthenticated200Response) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *UsersGetAuthenticated200Response) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *UsersGetAuthenticated200Response) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *UsersGetAuthenticated200Response) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *UsersGetAuthenticated200Response) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetReceivedEventsUrl

`func (o *UsersGetAuthenticated200Response) GetReceivedEventsUrl() string`

GetReceivedEventsUrl returns the ReceivedEventsUrl field if non-nil, zero value otherwise.

### GetReceivedEventsUrlOk

`func (o *UsersGetAuthenticated200Response) GetReceivedEventsUrlOk() (*string, bool)`

GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedEventsUrl

`func (o *UsersGetAuthenticated200Response) SetReceivedEventsUrl(v string)`

SetReceivedEventsUrl sets ReceivedEventsUrl field to given value.


### GetType

`func (o *UsersGetAuthenticated200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UsersGetAuthenticated200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UsersGetAuthenticated200Response) SetType(v string)`

SetType sets Type field to given value.


### GetSiteAdmin

`func (o *UsersGetAuthenticated200Response) GetSiteAdmin() bool`

GetSiteAdmin returns the SiteAdmin field if non-nil, zero value otherwise.

### GetSiteAdminOk

`func (o *UsersGetAuthenticated200Response) GetSiteAdminOk() (*bool, bool)`

GetSiteAdminOk returns a tuple with the SiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAdmin

`func (o *UsersGetAuthenticated200Response) SetSiteAdmin(v bool)`

SetSiteAdmin sets SiteAdmin field to given value.


### GetName

`func (o *UsersGetAuthenticated200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersGetAuthenticated200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersGetAuthenticated200Response) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UsersGetAuthenticated200Response) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UsersGetAuthenticated200Response) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCompany

`func (o *UsersGetAuthenticated200Response) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UsersGetAuthenticated200Response) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UsersGetAuthenticated200Response) SetCompany(v string)`

SetCompany sets Company field to given value.


### SetCompanyNil

`func (o *UsersGetAuthenticated200Response) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UsersGetAuthenticated200Response) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetBlog

`func (o *UsersGetAuthenticated200Response) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *UsersGetAuthenticated200Response) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *UsersGetAuthenticated200Response) SetBlog(v string)`

SetBlog sets Blog field to given value.


### SetBlogNil

`func (o *UsersGetAuthenticated200Response) SetBlogNil(b bool)`

 SetBlogNil sets the value for Blog to be an explicit nil

### UnsetBlog
`func (o *UsersGetAuthenticated200Response) UnsetBlog()`

UnsetBlog ensures that no value is present for Blog, not even an explicit nil
### GetLocation

`func (o *UsersGetAuthenticated200Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UsersGetAuthenticated200Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UsersGetAuthenticated200Response) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *UsersGetAuthenticated200Response) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *UsersGetAuthenticated200Response) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEmail

`func (o *UsersGetAuthenticated200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersGetAuthenticated200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersGetAuthenticated200Response) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *UsersGetAuthenticated200Response) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UsersGetAuthenticated200Response) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetHireable

`func (o *UsersGetAuthenticated200Response) GetHireable() bool`

GetHireable returns the Hireable field if non-nil, zero value otherwise.

### GetHireableOk

`func (o *UsersGetAuthenticated200Response) GetHireableOk() (*bool, bool)`

GetHireableOk returns a tuple with the Hireable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireable

`func (o *UsersGetAuthenticated200Response) SetHireable(v bool)`

SetHireable sets Hireable field to given value.


### SetHireableNil

`func (o *UsersGetAuthenticated200Response) SetHireableNil(b bool)`

 SetHireableNil sets the value for Hireable to be an explicit nil

### UnsetHireable
`func (o *UsersGetAuthenticated200Response) UnsetHireable()`

UnsetHireable ensures that no value is present for Hireable, not even an explicit nil
### GetBio

`func (o *UsersGetAuthenticated200Response) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UsersGetAuthenticated200Response) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UsersGetAuthenticated200Response) SetBio(v string)`

SetBio sets Bio field to given value.


### SetBioNil

`func (o *UsersGetAuthenticated200Response) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *UsersGetAuthenticated200Response) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetTwitterUsername

`func (o *UsersGetAuthenticated200Response) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *UsersGetAuthenticated200Response) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *UsersGetAuthenticated200Response) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *UsersGetAuthenticated200Response) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *UsersGetAuthenticated200Response) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *UsersGetAuthenticated200Response) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetPublicRepos

`func (o *UsersGetAuthenticated200Response) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *UsersGetAuthenticated200Response) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *UsersGetAuthenticated200Response) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.


### GetPublicGists

`func (o *UsersGetAuthenticated200Response) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *UsersGetAuthenticated200Response) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *UsersGetAuthenticated200Response) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.


### GetFollowers

`func (o *UsersGetAuthenticated200Response) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *UsersGetAuthenticated200Response) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *UsersGetAuthenticated200Response) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.


### GetFollowing

`func (o *UsersGetAuthenticated200Response) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *UsersGetAuthenticated200Response) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *UsersGetAuthenticated200Response) SetFollowing(v int32)`

SetFollowing sets Following field to given value.


### GetCreatedAt

`func (o *UsersGetAuthenticated200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersGetAuthenticated200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersGetAuthenticated200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UsersGetAuthenticated200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UsersGetAuthenticated200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UsersGetAuthenticated200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPrivateGists

`func (o *UsersGetAuthenticated200Response) GetPrivateGists() int32`

GetPrivateGists returns the PrivateGists field if non-nil, zero value otherwise.

### GetPrivateGistsOk

`func (o *UsersGetAuthenticated200Response) GetPrivateGistsOk() (*int32, bool)`

GetPrivateGistsOk returns a tuple with the PrivateGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateGists

`func (o *UsersGetAuthenticated200Response) SetPrivateGists(v int32)`

SetPrivateGists sets PrivateGists field to given value.


### GetTotalPrivateRepos

`func (o *UsersGetAuthenticated200Response) GetTotalPrivateRepos() int32`

GetTotalPrivateRepos returns the TotalPrivateRepos field if non-nil, zero value otherwise.

### GetTotalPrivateReposOk

`func (o *UsersGetAuthenticated200Response) GetTotalPrivateReposOk() (*int32, bool)`

GetTotalPrivateReposOk returns a tuple with the TotalPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateRepos

`func (o *UsersGetAuthenticated200Response) SetTotalPrivateRepos(v int32)`

SetTotalPrivateRepos sets TotalPrivateRepos field to given value.


### GetOwnedPrivateRepos

`func (o *UsersGetAuthenticated200Response) GetOwnedPrivateRepos() int32`

GetOwnedPrivateRepos returns the OwnedPrivateRepos field if non-nil, zero value otherwise.

### GetOwnedPrivateReposOk

`func (o *UsersGetAuthenticated200Response) GetOwnedPrivateReposOk() (*int32, bool)`

GetOwnedPrivateReposOk returns a tuple with the OwnedPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedPrivateRepos

`func (o *UsersGetAuthenticated200Response) SetOwnedPrivateRepos(v int32)`

SetOwnedPrivateRepos sets OwnedPrivateRepos field to given value.


### GetDiskUsage

`func (o *UsersGetAuthenticated200Response) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *UsersGetAuthenticated200Response) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *UsersGetAuthenticated200Response) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.


### GetCollaborators

`func (o *UsersGetAuthenticated200Response) GetCollaborators() int32`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *UsersGetAuthenticated200Response) GetCollaboratorsOk() (*int32, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *UsersGetAuthenticated200Response) SetCollaborators(v int32)`

SetCollaborators sets Collaborators field to given value.


### GetTwoFactorAuthentication

`func (o *UsersGetAuthenticated200Response) GetTwoFactorAuthentication() bool`

GetTwoFactorAuthentication returns the TwoFactorAuthentication field if non-nil, zero value otherwise.

### GetTwoFactorAuthenticationOk

`func (o *UsersGetAuthenticated200Response) GetTwoFactorAuthenticationOk() (*bool, bool)`

GetTwoFactorAuthenticationOk returns a tuple with the TwoFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthentication

`func (o *UsersGetAuthenticated200Response) SetTwoFactorAuthentication(v bool)`

SetTwoFactorAuthentication sets TwoFactorAuthentication field to given value.


### GetPlan

`func (o *UsersGetAuthenticated200Response) GetPlan() PublicUserPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UsersGetAuthenticated200Response) GetPlanOk() (*PublicUserPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UsersGetAuthenticated200Response) SetPlan(v PublicUserPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UsersGetAuthenticated200Response) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *UsersGetAuthenticated200Response) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *UsersGetAuthenticated200Response) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *UsersGetAuthenticated200Response) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *UsersGetAuthenticated200Response) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### SetSuspendedAtNil

`func (o *UsersGetAuthenticated200Response) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *UsersGetAuthenticated200Response) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetBusinessPlus

`func (o *UsersGetAuthenticated200Response) GetBusinessPlus() bool`

GetBusinessPlus returns the BusinessPlus field if non-nil, zero value otherwise.

### GetBusinessPlusOk

`func (o *UsersGetAuthenticated200Response) GetBusinessPlusOk() (*bool, bool)`

GetBusinessPlusOk returns a tuple with the BusinessPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPlus

`func (o *UsersGetAuthenticated200Response) SetBusinessPlus(v bool)`

SetBusinessPlus sets BusinessPlus field to given value.

### HasBusinessPlus

`func (o *UsersGetAuthenticated200Response) HasBusinessPlus() bool`

HasBusinessPlus returns a boolean if a field has been set.

### GetLdapDn

`func (o *UsersGetAuthenticated200Response) GetLdapDn() string`

GetLdapDn returns the LdapDn field if non-nil, zero value otherwise.

### GetLdapDnOk

`func (o *UsersGetAuthenticated200Response) GetLdapDnOk() (*string, bool)`

GetLdapDnOk returns a tuple with the LdapDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDn

`func (o *UsersGetAuthenticated200Response) SetLdapDn(v string)`

SetLdapDn sets LdapDn field to given value.

### HasLdapDn

`func (o *UsersGetAuthenticated200Response) HasLdapDn() bool`

HasLdapDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


