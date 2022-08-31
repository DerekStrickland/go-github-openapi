# TeamOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**ReposUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**HooksUrl** | **string** |  | 
**IssuesUrl** | **string** |  | 
**MembersUrl** | **string** |  | 
**PublicMembersUrl** | **string** |  | 
**AvatarUrl** | **string** |  | 
**Description** | **NullableString** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Blog** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**TwitterUsername** | Pointer to **NullableString** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**HasOrganizationProjects** | **bool** |  | 
**HasRepositoryProjects** | **bool** |  | 
**PublicRepos** | **int32** |  | 
**PublicGists** | **int32** |  | 
**Followers** | **int32** |  | 
**Following** | **int32** |  | 
**HtmlUrl** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Type** | **string** |  | 
**TotalPrivateRepos** | Pointer to **int32** |  | [optional] 
**OwnedPrivateRepos** | Pointer to **int32** |  | [optional] 
**PrivateGists** | Pointer to **NullableInt32** |  | [optional] 
**DiskUsage** | Pointer to **NullableInt32** |  | [optional] 
**Collaborators** | Pointer to **NullableInt32** |  | [optional] 
**BillingEmail** | Pointer to **NullableString** |  | [optional] 
**Plan** | Pointer to [**OrganizationFullPlan**](OrganizationFullPlan.md) |  | [optional] 
**DefaultRepositoryPermission** | Pointer to **NullableString** |  | [optional] 
**MembersCanCreateRepositories** | Pointer to **NullableBool** |  | [optional] 
**TwoFactorRequirementEnabled** | Pointer to **NullableBool** |  | [optional] 
**MembersAllowedRepositoryCreationType** | Pointer to **string** |  | [optional] 
**MembersCanCreatePublicRepositories** | Pointer to **bool** |  | [optional] 
**MembersCanCreatePrivateRepositories** | Pointer to **bool** |  | [optional] 
**MembersCanCreateInternalRepositories** | Pointer to **bool** |  | [optional] 
**MembersCanCreatePages** | Pointer to **bool** |  | [optional] 
**MembersCanCreatePublicPages** | Pointer to **bool** |  | [optional] 
**MembersCanCreatePrivatePages** | Pointer to **bool** |  | [optional] 
**MembersCanForkPrivateRepositories** | Pointer to **NullableBool** |  | [optional] 
**WebCommitSignoffRequired** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewTeamOrganization

`func NewTeamOrganization(login string, id int32, nodeId string, url string, reposUrl string, eventsUrl string, hooksUrl string, issuesUrl string, membersUrl string, publicMembersUrl string, avatarUrl string, description NullableString, hasOrganizationProjects bool, hasRepositoryProjects bool, publicRepos int32, publicGists int32, followers int32, following int32, htmlUrl string, createdAt time.Time, type_ string, updatedAt time.Time, ) *TeamOrganization`

NewTeamOrganization instantiates a new TeamOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamOrganizationWithDefaults

`func NewTeamOrganizationWithDefaults() *TeamOrganization`

NewTeamOrganizationWithDefaults instantiates a new TeamOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *TeamOrganization) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *TeamOrganization) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *TeamOrganization) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *TeamOrganization) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamOrganization) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamOrganization) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TeamOrganization) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamOrganization) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamOrganization) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TeamOrganization) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamOrganization) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamOrganization) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReposUrl

`func (o *TeamOrganization) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *TeamOrganization) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *TeamOrganization) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *TeamOrganization) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *TeamOrganization) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *TeamOrganization) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetHooksUrl

`func (o *TeamOrganization) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *TeamOrganization) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *TeamOrganization) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetIssuesUrl

`func (o *TeamOrganization) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *TeamOrganization) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *TeamOrganization) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetMembersUrl

`func (o *TeamOrganization) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *TeamOrganization) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *TeamOrganization) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetPublicMembersUrl

`func (o *TeamOrganization) GetPublicMembersUrl() string`

GetPublicMembersUrl returns the PublicMembersUrl field if non-nil, zero value otherwise.

### GetPublicMembersUrlOk

`func (o *TeamOrganization) GetPublicMembersUrlOk() (*string, bool)`

GetPublicMembersUrlOk returns a tuple with the PublicMembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMembersUrl

`func (o *TeamOrganization) SetPublicMembersUrl(v string)`

SetPublicMembersUrl sets PublicMembersUrl field to given value.


### GetAvatarUrl

`func (o *TeamOrganization) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TeamOrganization) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TeamOrganization) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetDescription

`func (o *TeamOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamOrganization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamOrganization) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TeamOrganization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamOrganization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *TeamOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompany

`func (o *TeamOrganization) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TeamOrganization) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TeamOrganization) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TeamOrganization) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetBlog

`func (o *TeamOrganization) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *TeamOrganization) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *TeamOrganization) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *TeamOrganization) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetLocation

`func (o *TeamOrganization) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TeamOrganization) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TeamOrganization) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TeamOrganization) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEmail

`func (o *TeamOrganization) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TeamOrganization) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TeamOrganization) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TeamOrganization) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *TeamOrganization) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *TeamOrganization) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *TeamOrganization) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *TeamOrganization) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *TeamOrganization) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *TeamOrganization) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetIsVerified

`func (o *TeamOrganization) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *TeamOrganization) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *TeamOrganization) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *TeamOrganization) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetHasOrganizationProjects

`func (o *TeamOrganization) GetHasOrganizationProjects() bool`

GetHasOrganizationProjects returns the HasOrganizationProjects field if non-nil, zero value otherwise.

### GetHasOrganizationProjectsOk

`func (o *TeamOrganization) GetHasOrganizationProjectsOk() (*bool, bool)`

GetHasOrganizationProjectsOk returns a tuple with the HasOrganizationProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOrganizationProjects

`func (o *TeamOrganization) SetHasOrganizationProjects(v bool)`

SetHasOrganizationProjects sets HasOrganizationProjects field to given value.


### GetHasRepositoryProjects

`func (o *TeamOrganization) GetHasRepositoryProjects() bool`

GetHasRepositoryProjects returns the HasRepositoryProjects field if non-nil, zero value otherwise.

### GetHasRepositoryProjectsOk

`func (o *TeamOrganization) GetHasRepositoryProjectsOk() (*bool, bool)`

GetHasRepositoryProjectsOk returns a tuple with the HasRepositoryProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRepositoryProjects

`func (o *TeamOrganization) SetHasRepositoryProjects(v bool)`

SetHasRepositoryProjects sets HasRepositoryProjects field to given value.


### GetPublicRepos

`func (o *TeamOrganization) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *TeamOrganization) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *TeamOrganization) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.


### GetPublicGists

`func (o *TeamOrganization) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *TeamOrganization) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *TeamOrganization) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.


### GetFollowers

`func (o *TeamOrganization) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *TeamOrganization) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *TeamOrganization) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.


### GetFollowing

`func (o *TeamOrganization) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *TeamOrganization) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *TeamOrganization) SetFollowing(v int32)`

SetFollowing sets Following field to given value.


### GetHtmlUrl

`func (o *TeamOrganization) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamOrganization) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamOrganization) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCreatedAt

`func (o *TeamOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *TeamOrganization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamOrganization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamOrganization) SetType(v string)`

SetType sets Type field to given value.


### GetTotalPrivateRepos

`func (o *TeamOrganization) GetTotalPrivateRepos() int32`

GetTotalPrivateRepos returns the TotalPrivateRepos field if non-nil, zero value otherwise.

### GetTotalPrivateReposOk

`func (o *TeamOrganization) GetTotalPrivateReposOk() (*int32, bool)`

GetTotalPrivateReposOk returns a tuple with the TotalPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateRepos

`func (o *TeamOrganization) SetTotalPrivateRepos(v int32)`

SetTotalPrivateRepos sets TotalPrivateRepos field to given value.

### HasTotalPrivateRepos

`func (o *TeamOrganization) HasTotalPrivateRepos() bool`

HasTotalPrivateRepos returns a boolean if a field has been set.

### GetOwnedPrivateRepos

`func (o *TeamOrganization) GetOwnedPrivateRepos() int32`

GetOwnedPrivateRepos returns the OwnedPrivateRepos field if non-nil, zero value otherwise.

### GetOwnedPrivateReposOk

`func (o *TeamOrganization) GetOwnedPrivateReposOk() (*int32, bool)`

GetOwnedPrivateReposOk returns a tuple with the OwnedPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedPrivateRepos

`func (o *TeamOrganization) SetOwnedPrivateRepos(v int32)`

SetOwnedPrivateRepos sets OwnedPrivateRepos field to given value.

### HasOwnedPrivateRepos

`func (o *TeamOrganization) HasOwnedPrivateRepos() bool`

HasOwnedPrivateRepos returns a boolean if a field has been set.

### GetPrivateGists

`func (o *TeamOrganization) GetPrivateGists() int32`

GetPrivateGists returns the PrivateGists field if non-nil, zero value otherwise.

### GetPrivateGistsOk

`func (o *TeamOrganization) GetPrivateGistsOk() (*int32, bool)`

GetPrivateGistsOk returns a tuple with the PrivateGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateGists

`func (o *TeamOrganization) SetPrivateGists(v int32)`

SetPrivateGists sets PrivateGists field to given value.

### HasPrivateGists

`func (o *TeamOrganization) HasPrivateGists() bool`

HasPrivateGists returns a boolean if a field has been set.

### SetPrivateGistsNil

`func (o *TeamOrganization) SetPrivateGistsNil(b bool)`

 SetPrivateGistsNil sets the value for PrivateGists to be an explicit nil

### UnsetPrivateGists
`func (o *TeamOrganization) UnsetPrivateGists()`

UnsetPrivateGists ensures that no value is present for PrivateGists, not even an explicit nil
### GetDiskUsage

`func (o *TeamOrganization) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *TeamOrganization) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *TeamOrganization) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *TeamOrganization) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### SetDiskUsageNil

`func (o *TeamOrganization) SetDiskUsageNil(b bool)`

 SetDiskUsageNil sets the value for DiskUsage to be an explicit nil

### UnsetDiskUsage
`func (o *TeamOrganization) UnsetDiskUsage()`

UnsetDiskUsage ensures that no value is present for DiskUsage, not even an explicit nil
### GetCollaborators

`func (o *TeamOrganization) GetCollaborators() int32`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *TeamOrganization) GetCollaboratorsOk() (*int32, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *TeamOrganization) SetCollaborators(v int32)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *TeamOrganization) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### SetCollaboratorsNil

`func (o *TeamOrganization) SetCollaboratorsNil(b bool)`

 SetCollaboratorsNil sets the value for Collaborators to be an explicit nil

### UnsetCollaborators
`func (o *TeamOrganization) UnsetCollaborators()`

UnsetCollaborators ensures that no value is present for Collaborators, not even an explicit nil
### GetBillingEmail

`func (o *TeamOrganization) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *TeamOrganization) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *TeamOrganization) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *TeamOrganization) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *TeamOrganization) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *TeamOrganization) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetPlan

`func (o *TeamOrganization) GetPlan() OrganizationFullPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *TeamOrganization) GetPlanOk() (*OrganizationFullPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *TeamOrganization) SetPlan(v OrganizationFullPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *TeamOrganization) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDefaultRepositoryPermission

`func (o *TeamOrganization) GetDefaultRepositoryPermission() string`

GetDefaultRepositoryPermission returns the DefaultRepositoryPermission field if non-nil, zero value otherwise.

### GetDefaultRepositoryPermissionOk

`func (o *TeamOrganization) GetDefaultRepositoryPermissionOk() (*string, bool)`

GetDefaultRepositoryPermissionOk returns a tuple with the DefaultRepositoryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepositoryPermission

`func (o *TeamOrganization) SetDefaultRepositoryPermission(v string)`

SetDefaultRepositoryPermission sets DefaultRepositoryPermission field to given value.

### HasDefaultRepositoryPermission

`func (o *TeamOrganization) HasDefaultRepositoryPermission() bool`

HasDefaultRepositoryPermission returns a boolean if a field has been set.

### SetDefaultRepositoryPermissionNil

`func (o *TeamOrganization) SetDefaultRepositoryPermissionNil(b bool)`

 SetDefaultRepositoryPermissionNil sets the value for DefaultRepositoryPermission to be an explicit nil

### UnsetDefaultRepositoryPermission
`func (o *TeamOrganization) UnsetDefaultRepositoryPermission()`

UnsetDefaultRepositoryPermission ensures that no value is present for DefaultRepositoryPermission, not even an explicit nil
### GetMembersCanCreateRepositories

`func (o *TeamOrganization) GetMembersCanCreateRepositories() bool`

GetMembersCanCreateRepositories returns the MembersCanCreateRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreateRepositoriesOk

`func (o *TeamOrganization) GetMembersCanCreateRepositoriesOk() (*bool, bool)`

GetMembersCanCreateRepositoriesOk returns a tuple with the MembersCanCreateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreateRepositories

`func (o *TeamOrganization) SetMembersCanCreateRepositories(v bool)`

SetMembersCanCreateRepositories sets MembersCanCreateRepositories field to given value.

### HasMembersCanCreateRepositories

`func (o *TeamOrganization) HasMembersCanCreateRepositories() bool`

HasMembersCanCreateRepositories returns a boolean if a field has been set.

### SetMembersCanCreateRepositoriesNil

`func (o *TeamOrganization) SetMembersCanCreateRepositoriesNil(b bool)`

 SetMembersCanCreateRepositoriesNil sets the value for MembersCanCreateRepositories to be an explicit nil

### UnsetMembersCanCreateRepositories
`func (o *TeamOrganization) UnsetMembersCanCreateRepositories()`

UnsetMembersCanCreateRepositories ensures that no value is present for MembersCanCreateRepositories, not even an explicit nil
### GetTwoFactorRequirementEnabled

`func (o *TeamOrganization) GetTwoFactorRequirementEnabled() bool`

GetTwoFactorRequirementEnabled returns the TwoFactorRequirementEnabled field if non-nil, zero value otherwise.

### GetTwoFactorRequirementEnabledOk

`func (o *TeamOrganization) GetTwoFactorRequirementEnabledOk() (*bool, bool)`

GetTwoFactorRequirementEnabledOk returns a tuple with the TwoFactorRequirementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorRequirementEnabled

`func (o *TeamOrganization) SetTwoFactorRequirementEnabled(v bool)`

SetTwoFactorRequirementEnabled sets TwoFactorRequirementEnabled field to given value.

### HasTwoFactorRequirementEnabled

`func (o *TeamOrganization) HasTwoFactorRequirementEnabled() bool`

HasTwoFactorRequirementEnabled returns a boolean if a field has been set.

### SetTwoFactorRequirementEnabledNil

`func (o *TeamOrganization) SetTwoFactorRequirementEnabledNil(b bool)`

 SetTwoFactorRequirementEnabledNil sets the value for TwoFactorRequirementEnabled to be an explicit nil

### UnsetTwoFactorRequirementEnabled
`func (o *TeamOrganization) UnsetTwoFactorRequirementEnabled()`

UnsetTwoFactorRequirementEnabled ensures that no value is present for TwoFactorRequirementEnabled, not even an explicit nil
### GetMembersAllowedRepositoryCreationType

`func (o *TeamOrganization) GetMembersAllowedRepositoryCreationType() string`

GetMembersAllowedRepositoryCreationType returns the MembersAllowedRepositoryCreationType field if non-nil, zero value otherwise.

### GetMembersAllowedRepositoryCreationTypeOk

`func (o *TeamOrganization) GetMembersAllowedRepositoryCreationTypeOk() (*string, bool)`

GetMembersAllowedRepositoryCreationTypeOk returns a tuple with the MembersAllowedRepositoryCreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersAllowedRepositoryCreationType

`func (o *TeamOrganization) SetMembersAllowedRepositoryCreationType(v string)`

SetMembersAllowedRepositoryCreationType sets MembersAllowedRepositoryCreationType field to given value.

### HasMembersAllowedRepositoryCreationType

`func (o *TeamOrganization) HasMembersAllowedRepositoryCreationType() bool`

HasMembersAllowedRepositoryCreationType returns a boolean if a field has been set.

### GetMembersCanCreatePublicRepositories

`func (o *TeamOrganization) GetMembersCanCreatePublicRepositories() bool`

GetMembersCanCreatePublicRepositories returns the MembersCanCreatePublicRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreatePublicRepositoriesOk

`func (o *TeamOrganization) GetMembersCanCreatePublicRepositoriesOk() (*bool, bool)`

GetMembersCanCreatePublicRepositoriesOk returns a tuple with the MembersCanCreatePublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePublicRepositories

`func (o *TeamOrganization) SetMembersCanCreatePublicRepositories(v bool)`

SetMembersCanCreatePublicRepositories sets MembersCanCreatePublicRepositories field to given value.

### HasMembersCanCreatePublicRepositories

`func (o *TeamOrganization) HasMembersCanCreatePublicRepositories() bool`

HasMembersCanCreatePublicRepositories returns a boolean if a field has been set.

### GetMembersCanCreatePrivateRepositories

`func (o *TeamOrganization) GetMembersCanCreatePrivateRepositories() bool`

GetMembersCanCreatePrivateRepositories returns the MembersCanCreatePrivateRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreatePrivateRepositoriesOk

`func (o *TeamOrganization) GetMembersCanCreatePrivateRepositoriesOk() (*bool, bool)`

GetMembersCanCreatePrivateRepositoriesOk returns a tuple with the MembersCanCreatePrivateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePrivateRepositories

`func (o *TeamOrganization) SetMembersCanCreatePrivateRepositories(v bool)`

SetMembersCanCreatePrivateRepositories sets MembersCanCreatePrivateRepositories field to given value.

### HasMembersCanCreatePrivateRepositories

`func (o *TeamOrganization) HasMembersCanCreatePrivateRepositories() bool`

HasMembersCanCreatePrivateRepositories returns a boolean if a field has been set.

### GetMembersCanCreateInternalRepositories

`func (o *TeamOrganization) GetMembersCanCreateInternalRepositories() bool`

GetMembersCanCreateInternalRepositories returns the MembersCanCreateInternalRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreateInternalRepositoriesOk

`func (o *TeamOrganization) GetMembersCanCreateInternalRepositoriesOk() (*bool, bool)`

GetMembersCanCreateInternalRepositoriesOk returns a tuple with the MembersCanCreateInternalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreateInternalRepositories

`func (o *TeamOrganization) SetMembersCanCreateInternalRepositories(v bool)`

SetMembersCanCreateInternalRepositories sets MembersCanCreateInternalRepositories field to given value.

### HasMembersCanCreateInternalRepositories

`func (o *TeamOrganization) HasMembersCanCreateInternalRepositories() bool`

HasMembersCanCreateInternalRepositories returns a boolean if a field has been set.

### GetMembersCanCreatePages

`func (o *TeamOrganization) GetMembersCanCreatePages() bool`

GetMembersCanCreatePages returns the MembersCanCreatePages field if non-nil, zero value otherwise.

### GetMembersCanCreatePagesOk

`func (o *TeamOrganization) GetMembersCanCreatePagesOk() (*bool, bool)`

GetMembersCanCreatePagesOk returns a tuple with the MembersCanCreatePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePages

`func (o *TeamOrganization) SetMembersCanCreatePages(v bool)`

SetMembersCanCreatePages sets MembersCanCreatePages field to given value.

### HasMembersCanCreatePages

`func (o *TeamOrganization) HasMembersCanCreatePages() bool`

HasMembersCanCreatePages returns a boolean if a field has been set.

### GetMembersCanCreatePublicPages

`func (o *TeamOrganization) GetMembersCanCreatePublicPages() bool`

GetMembersCanCreatePublicPages returns the MembersCanCreatePublicPages field if non-nil, zero value otherwise.

### GetMembersCanCreatePublicPagesOk

`func (o *TeamOrganization) GetMembersCanCreatePublicPagesOk() (*bool, bool)`

GetMembersCanCreatePublicPagesOk returns a tuple with the MembersCanCreatePublicPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePublicPages

`func (o *TeamOrganization) SetMembersCanCreatePublicPages(v bool)`

SetMembersCanCreatePublicPages sets MembersCanCreatePublicPages field to given value.

### HasMembersCanCreatePublicPages

`func (o *TeamOrganization) HasMembersCanCreatePublicPages() bool`

HasMembersCanCreatePublicPages returns a boolean if a field has been set.

### GetMembersCanCreatePrivatePages

`func (o *TeamOrganization) GetMembersCanCreatePrivatePages() bool`

GetMembersCanCreatePrivatePages returns the MembersCanCreatePrivatePages field if non-nil, zero value otherwise.

### GetMembersCanCreatePrivatePagesOk

`func (o *TeamOrganization) GetMembersCanCreatePrivatePagesOk() (*bool, bool)`

GetMembersCanCreatePrivatePagesOk returns a tuple with the MembersCanCreatePrivatePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePrivatePages

`func (o *TeamOrganization) SetMembersCanCreatePrivatePages(v bool)`

SetMembersCanCreatePrivatePages sets MembersCanCreatePrivatePages field to given value.

### HasMembersCanCreatePrivatePages

`func (o *TeamOrganization) HasMembersCanCreatePrivatePages() bool`

HasMembersCanCreatePrivatePages returns a boolean if a field has been set.

### GetMembersCanForkPrivateRepositories

`func (o *TeamOrganization) GetMembersCanForkPrivateRepositories() bool`

GetMembersCanForkPrivateRepositories returns the MembersCanForkPrivateRepositories field if non-nil, zero value otherwise.

### GetMembersCanForkPrivateRepositoriesOk

`func (o *TeamOrganization) GetMembersCanForkPrivateRepositoriesOk() (*bool, bool)`

GetMembersCanForkPrivateRepositoriesOk returns a tuple with the MembersCanForkPrivateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanForkPrivateRepositories

`func (o *TeamOrganization) SetMembersCanForkPrivateRepositories(v bool)`

SetMembersCanForkPrivateRepositories sets MembersCanForkPrivateRepositories field to given value.

### HasMembersCanForkPrivateRepositories

`func (o *TeamOrganization) HasMembersCanForkPrivateRepositories() bool`

HasMembersCanForkPrivateRepositories returns a boolean if a field has been set.

### SetMembersCanForkPrivateRepositoriesNil

`func (o *TeamOrganization) SetMembersCanForkPrivateRepositoriesNil(b bool)`

 SetMembersCanForkPrivateRepositoriesNil sets the value for MembersCanForkPrivateRepositories to be an explicit nil

### UnsetMembersCanForkPrivateRepositories
`func (o *TeamOrganization) UnsetMembersCanForkPrivateRepositories()`

UnsetMembersCanForkPrivateRepositories ensures that no value is present for MembersCanForkPrivateRepositories, not even an explicit nil
### GetWebCommitSignoffRequired

`func (o *TeamOrganization) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *TeamOrganization) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *TeamOrganization) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *TeamOrganization) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TeamOrganization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamOrganization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamOrganization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


