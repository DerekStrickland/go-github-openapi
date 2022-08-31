# OrganizationFull

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
**AdvancedSecurityEnabledForNewRepositories** | Pointer to **bool** | Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization.  This field is only visible to organization owners or members of a team with the security manager role. | [optional] 
**DependabotAlertsEnabledForNewRepositories** | Pointer to **bool** | Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred to this organization.  This field is only visible to organization owners or members of a team with the security manager role. | [optional] 
**DependabotSecurityUpdatesEnabledForNewRepositories** | Pointer to **bool** | Whether dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization.  This field is only visible to organization owners or members of a team with the security manager role. | [optional] 
**DependencyGraphEnabledForNewRepositories** | Pointer to **bool** | Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization.  This field is only visible to organization owners or members of a team with the security manager role. | [optional] 
**SecretScanningEnabledForNewRepositories** | Pointer to **bool** | Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization.  This field is only visible to organization owners or members of a team with the security manager role. | [optional] 
**SecretScanningPushProtectionEnabledForNewRepositories** | Pointer to **bool** | Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization.  This field is only visible to organization owners or members of a team with the security manager role. | [optional] 

## Methods

### NewOrganizationFull

`func NewOrganizationFull(login string, id int32, nodeId string, url string, reposUrl string, eventsUrl string, hooksUrl string, issuesUrl string, membersUrl string, publicMembersUrl string, avatarUrl string, description NullableString, hasOrganizationProjects bool, hasRepositoryProjects bool, publicRepos int32, publicGists int32, followers int32, following int32, htmlUrl string, createdAt time.Time, type_ string, updatedAt time.Time, ) *OrganizationFull`

NewOrganizationFull instantiates a new OrganizationFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFullWithDefaults

`func NewOrganizationFullWithDefaults() *OrganizationFull`

NewOrganizationFullWithDefaults instantiates a new OrganizationFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *OrganizationFull) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *OrganizationFull) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *OrganizationFull) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetId

`func (o *OrganizationFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationFull) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationFull) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *OrganizationFull) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *OrganizationFull) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *OrganizationFull) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *OrganizationFull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationFull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationFull) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReposUrl

`func (o *OrganizationFull) GetReposUrl() string`

GetReposUrl returns the ReposUrl field if non-nil, zero value otherwise.

### GetReposUrlOk

`func (o *OrganizationFull) GetReposUrlOk() (*string, bool)`

GetReposUrlOk returns a tuple with the ReposUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposUrl

`func (o *OrganizationFull) SetReposUrl(v string)`

SetReposUrl sets ReposUrl field to given value.


### GetEventsUrl

`func (o *OrganizationFull) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *OrganizationFull) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *OrganizationFull) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetHooksUrl

`func (o *OrganizationFull) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *OrganizationFull) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *OrganizationFull) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetIssuesUrl

`func (o *OrganizationFull) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *OrganizationFull) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *OrganizationFull) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetMembersUrl

`func (o *OrganizationFull) GetMembersUrl() string`

GetMembersUrl returns the MembersUrl field if non-nil, zero value otherwise.

### GetMembersUrlOk

`func (o *OrganizationFull) GetMembersUrlOk() (*string, bool)`

GetMembersUrlOk returns a tuple with the MembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersUrl

`func (o *OrganizationFull) SetMembersUrl(v string)`

SetMembersUrl sets MembersUrl field to given value.


### GetPublicMembersUrl

`func (o *OrganizationFull) GetPublicMembersUrl() string`

GetPublicMembersUrl returns the PublicMembersUrl field if non-nil, zero value otherwise.

### GetPublicMembersUrlOk

`func (o *OrganizationFull) GetPublicMembersUrlOk() (*string, bool)`

GetPublicMembersUrlOk returns a tuple with the PublicMembersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicMembersUrl

`func (o *OrganizationFull) SetPublicMembersUrl(v string)`

SetPublicMembersUrl sets PublicMembersUrl field to given value.


### GetAvatarUrl

`func (o *OrganizationFull) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *OrganizationFull) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *OrganizationFull) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetDescription

`func (o *OrganizationFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationFull) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *OrganizationFull) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrganizationFull) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *OrganizationFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompany

`func (o *OrganizationFull) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OrganizationFull) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OrganizationFull) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *OrganizationFull) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetBlog

`func (o *OrganizationFull) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *OrganizationFull) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *OrganizationFull) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *OrganizationFull) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetLocation

`func (o *OrganizationFull) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OrganizationFull) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OrganizationFull) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OrganizationFull) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationFull) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationFull) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationFull) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationFull) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *OrganizationFull) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *OrganizationFull) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *OrganizationFull) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *OrganizationFull) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *OrganizationFull) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *OrganizationFull) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetIsVerified

`func (o *OrganizationFull) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *OrganizationFull) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *OrganizationFull) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *OrganizationFull) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetHasOrganizationProjects

`func (o *OrganizationFull) GetHasOrganizationProjects() bool`

GetHasOrganizationProjects returns the HasOrganizationProjects field if non-nil, zero value otherwise.

### GetHasOrganizationProjectsOk

`func (o *OrganizationFull) GetHasOrganizationProjectsOk() (*bool, bool)`

GetHasOrganizationProjectsOk returns a tuple with the HasOrganizationProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOrganizationProjects

`func (o *OrganizationFull) SetHasOrganizationProjects(v bool)`

SetHasOrganizationProjects sets HasOrganizationProjects field to given value.


### GetHasRepositoryProjects

`func (o *OrganizationFull) GetHasRepositoryProjects() bool`

GetHasRepositoryProjects returns the HasRepositoryProjects field if non-nil, zero value otherwise.

### GetHasRepositoryProjectsOk

`func (o *OrganizationFull) GetHasRepositoryProjectsOk() (*bool, bool)`

GetHasRepositoryProjectsOk returns a tuple with the HasRepositoryProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRepositoryProjects

`func (o *OrganizationFull) SetHasRepositoryProjects(v bool)`

SetHasRepositoryProjects sets HasRepositoryProjects field to given value.


### GetPublicRepos

`func (o *OrganizationFull) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *OrganizationFull) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *OrganizationFull) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.


### GetPublicGists

`func (o *OrganizationFull) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *OrganizationFull) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *OrganizationFull) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.


### GetFollowers

`func (o *OrganizationFull) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *OrganizationFull) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *OrganizationFull) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.


### GetFollowing

`func (o *OrganizationFull) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *OrganizationFull) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *OrganizationFull) SetFollowing(v int32)`

SetFollowing sets Following field to given value.


### GetHtmlUrl

`func (o *OrganizationFull) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *OrganizationFull) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *OrganizationFull) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCreatedAt

`func (o *OrganizationFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *OrganizationFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationFull) SetType(v string)`

SetType sets Type field to given value.


### GetTotalPrivateRepos

`func (o *OrganizationFull) GetTotalPrivateRepos() int32`

GetTotalPrivateRepos returns the TotalPrivateRepos field if non-nil, zero value otherwise.

### GetTotalPrivateReposOk

`func (o *OrganizationFull) GetTotalPrivateReposOk() (*int32, bool)`

GetTotalPrivateReposOk returns a tuple with the TotalPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrivateRepos

`func (o *OrganizationFull) SetTotalPrivateRepos(v int32)`

SetTotalPrivateRepos sets TotalPrivateRepos field to given value.

### HasTotalPrivateRepos

`func (o *OrganizationFull) HasTotalPrivateRepos() bool`

HasTotalPrivateRepos returns a boolean if a field has been set.

### GetOwnedPrivateRepos

`func (o *OrganizationFull) GetOwnedPrivateRepos() int32`

GetOwnedPrivateRepos returns the OwnedPrivateRepos field if non-nil, zero value otherwise.

### GetOwnedPrivateReposOk

`func (o *OrganizationFull) GetOwnedPrivateReposOk() (*int32, bool)`

GetOwnedPrivateReposOk returns a tuple with the OwnedPrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedPrivateRepos

`func (o *OrganizationFull) SetOwnedPrivateRepos(v int32)`

SetOwnedPrivateRepos sets OwnedPrivateRepos field to given value.

### HasOwnedPrivateRepos

`func (o *OrganizationFull) HasOwnedPrivateRepos() bool`

HasOwnedPrivateRepos returns a boolean if a field has been set.

### GetPrivateGists

`func (o *OrganizationFull) GetPrivateGists() int32`

GetPrivateGists returns the PrivateGists field if non-nil, zero value otherwise.

### GetPrivateGistsOk

`func (o *OrganizationFull) GetPrivateGistsOk() (*int32, bool)`

GetPrivateGistsOk returns a tuple with the PrivateGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateGists

`func (o *OrganizationFull) SetPrivateGists(v int32)`

SetPrivateGists sets PrivateGists field to given value.

### HasPrivateGists

`func (o *OrganizationFull) HasPrivateGists() bool`

HasPrivateGists returns a boolean if a field has been set.

### SetPrivateGistsNil

`func (o *OrganizationFull) SetPrivateGistsNil(b bool)`

 SetPrivateGistsNil sets the value for PrivateGists to be an explicit nil

### UnsetPrivateGists
`func (o *OrganizationFull) UnsetPrivateGists()`

UnsetPrivateGists ensures that no value is present for PrivateGists, not even an explicit nil
### GetDiskUsage

`func (o *OrganizationFull) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *OrganizationFull) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *OrganizationFull) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *OrganizationFull) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### SetDiskUsageNil

`func (o *OrganizationFull) SetDiskUsageNil(b bool)`

 SetDiskUsageNil sets the value for DiskUsage to be an explicit nil

### UnsetDiskUsage
`func (o *OrganizationFull) UnsetDiskUsage()`

UnsetDiskUsage ensures that no value is present for DiskUsage, not even an explicit nil
### GetCollaborators

`func (o *OrganizationFull) GetCollaborators() int32`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *OrganizationFull) GetCollaboratorsOk() (*int32, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *OrganizationFull) SetCollaborators(v int32)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *OrganizationFull) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### SetCollaboratorsNil

`func (o *OrganizationFull) SetCollaboratorsNil(b bool)`

 SetCollaboratorsNil sets the value for Collaborators to be an explicit nil

### UnsetCollaborators
`func (o *OrganizationFull) UnsetCollaborators()`

UnsetCollaborators ensures that no value is present for Collaborators, not even an explicit nil
### GetBillingEmail

`func (o *OrganizationFull) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrganizationFull) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrganizationFull) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrganizationFull) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *OrganizationFull) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *OrganizationFull) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetPlan

`func (o *OrganizationFull) GetPlan() OrganizationFullPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationFull) GetPlanOk() (*OrganizationFullPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationFull) SetPlan(v OrganizationFullPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrganizationFull) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDefaultRepositoryPermission

`func (o *OrganizationFull) GetDefaultRepositoryPermission() string`

GetDefaultRepositoryPermission returns the DefaultRepositoryPermission field if non-nil, zero value otherwise.

### GetDefaultRepositoryPermissionOk

`func (o *OrganizationFull) GetDefaultRepositoryPermissionOk() (*string, bool)`

GetDefaultRepositoryPermissionOk returns a tuple with the DefaultRepositoryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepositoryPermission

`func (o *OrganizationFull) SetDefaultRepositoryPermission(v string)`

SetDefaultRepositoryPermission sets DefaultRepositoryPermission field to given value.

### HasDefaultRepositoryPermission

`func (o *OrganizationFull) HasDefaultRepositoryPermission() bool`

HasDefaultRepositoryPermission returns a boolean if a field has been set.

### SetDefaultRepositoryPermissionNil

`func (o *OrganizationFull) SetDefaultRepositoryPermissionNil(b bool)`

 SetDefaultRepositoryPermissionNil sets the value for DefaultRepositoryPermission to be an explicit nil

### UnsetDefaultRepositoryPermission
`func (o *OrganizationFull) UnsetDefaultRepositoryPermission()`

UnsetDefaultRepositoryPermission ensures that no value is present for DefaultRepositoryPermission, not even an explicit nil
### GetMembersCanCreateRepositories

`func (o *OrganizationFull) GetMembersCanCreateRepositories() bool`

GetMembersCanCreateRepositories returns the MembersCanCreateRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreateRepositoriesOk

`func (o *OrganizationFull) GetMembersCanCreateRepositoriesOk() (*bool, bool)`

GetMembersCanCreateRepositoriesOk returns a tuple with the MembersCanCreateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreateRepositories

`func (o *OrganizationFull) SetMembersCanCreateRepositories(v bool)`

SetMembersCanCreateRepositories sets MembersCanCreateRepositories field to given value.

### HasMembersCanCreateRepositories

`func (o *OrganizationFull) HasMembersCanCreateRepositories() bool`

HasMembersCanCreateRepositories returns a boolean if a field has been set.

### SetMembersCanCreateRepositoriesNil

`func (o *OrganizationFull) SetMembersCanCreateRepositoriesNil(b bool)`

 SetMembersCanCreateRepositoriesNil sets the value for MembersCanCreateRepositories to be an explicit nil

### UnsetMembersCanCreateRepositories
`func (o *OrganizationFull) UnsetMembersCanCreateRepositories()`

UnsetMembersCanCreateRepositories ensures that no value is present for MembersCanCreateRepositories, not even an explicit nil
### GetTwoFactorRequirementEnabled

`func (o *OrganizationFull) GetTwoFactorRequirementEnabled() bool`

GetTwoFactorRequirementEnabled returns the TwoFactorRequirementEnabled field if non-nil, zero value otherwise.

### GetTwoFactorRequirementEnabledOk

`func (o *OrganizationFull) GetTwoFactorRequirementEnabledOk() (*bool, bool)`

GetTwoFactorRequirementEnabledOk returns a tuple with the TwoFactorRequirementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorRequirementEnabled

`func (o *OrganizationFull) SetTwoFactorRequirementEnabled(v bool)`

SetTwoFactorRequirementEnabled sets TwoFactorRequirementEnabled field to given value.

### HasTwoFactorRequirementEnabled

`func (o *OrganizationFull) HasTwoFactorRequirementEnabled() bool`

HasTwoFactorRequirementEnabled returns a boolean if a field has been set.

### SetTwoFactorRequirementEnabledNil

`func (o *OrganizationFull) SetTwoFactorRequirementEnabledNil(b bool)`

 SetTwoFactorRequirementEnabledNil sets the value for TwoFactorRequirementEnabled to be an explicit nil

### UnsetTwoFactorRequirementEnabled
`func (o *OrganizationFull) UnsetTwoFactorRequirementEnabled()`

UnsetTwoFactorRequirementEnabled ensures that no value is present for TwoFactorRequirementEnabled, not even an explicit nil
### GetMembersAllowedRepositoryCreationType

`func (o *OrganizationFull) GetMembersAllowedRepositoryCreationType() string`

GetMembersAllowedRepositoryCreationType returns the MembersAllowedRepositoryCreationType field if non-nil, zero value otherwise.

### GetMembersAllowedRepositoryCreationTypeOk

`func (o *OrganizationFull) GetMembersAllowedRepositoryCreationTypeOk() (*string, bool)`

GetMembersAllowedRepositoryCreationTypeOk returns a tuple with the MembersAllowedRepositoryCreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersAllowedRepositoryCreationType

`func (o *OrganizationFull) SetMembersAllowedRepositoryCreationType(v string)`

SetMembersAllowedRepositoryCreationType sets MembersAllowedRepositoryCreationType field to given value.

### HasMembersAllowedRepositoryCreationType

`func (o *OrganizationFull) HasMembersAllowedRepositoryCreationType() bool`

HasMembersAllowedRepositoryCreationType returns a boolean if a field has been set.

### GetMembersCanCreatePublicRepositories

`func (o *OrganizationFull) GetMembersCanCreatePublicRepositories() bool`

GetMembersCanCreatePublicRepositories returns the MembersCanCreatePublicRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreatePublicRepositoriesOk

`func (o *OrganizationFull) GetMembersCanCreatePublicRepositoriesOk() (*bool, bool)`

GetMembersCanCreatePublicRepositoriesOk returns a tuple with the MembersCanCreatePublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePublicRepositories

`func (o *OrganizationFull) SetMembersCanCreatePublicRepositories(v bool)`

SetMembersCanCreatePublicRepositories sets MembersCanCreatePublicRepositories field to given value.

### HasMembersCanCreatePublicRepositories

`func (o *OrganizationFull) HasMembersCanCreatePublicRepositories() bool`

HasMembersCanCreatePublicRepositories returns a boolean if a field has been set.

### GetMembersCanCreatePrivateRepositories

`func (o *OrganizationFull) GetMembersCanCreatePrivateRepositories() bool`

GetMembersCanCreatePrivateRepositories returns the MembersCanCreatePrivateRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreatePrivateRepositoriesOk

`func (o *OrganizationFull) GetMembersCanCreatePrivateRepositoriesOk() (*bool, bool)`

GetMembersCanCreatePrivateRepositoriesOk returns a tuple with the MembersCanCreatePrivateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePrivateRepositories

`func (o *OrganizationFull) SetMembersCanCreatePrivateRepositories(v bool)`

SetMembersCanCreatePrivateRepositories sets MembersCanCreatePrivateRepositories field to given value.

### HasMembersCanCreatePrivateRepositories

`func (o *OrganizationFull) HasMembersCanCreatePrivateRepositories() bool`

HasMembersCanCreatePrivateRepositories returns a boolean if a field has been set.

### GetMembersCanCreateInternalRepositories

`func (o *OrganizationFull) GetMembersCanCreateInternalRepositories() bool`

GetMembersCanCreateInternalRepositories returns the MembersCanCreateInternalRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreateInternalRepositoriesOk

`func (o *OrganizationFull) GetMembersCanCreateInternalRepositoriesOk() (*bool, bool)`

GetMembersCanCreateInternalRepositoriesOk returns a tuple with the MembersCanCreateInternalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreateInternalRepositories

`func (o *OrganizationFull) SetMembersCanCreateInternalRepositories(v bool)`

SetMembersCanCreateInternalRepositories sets MembersCanCreateInternalRepositories field to given value.

### HasMembersCanCreateInternalRepositories

`func (o *OrganizationFull) HasMembersCanCreateInternalRepositories() bool`

HasMembersCanCreateInternalRepositories returns a boolean if a field has been set.

### GetMembersCanCreatePages

`func (o *OrganizationFull) GetMembersCanCreatePages() bool`

GetMembersCanCreatePages returns the MembersCanCreatePages field if non-nil, zero value otherwise.

### GetMembersCanCreatePagesOk

`func (o *OrganizationFull) GetMembersCanCreatePagesOk() (*bool, bool)`

GetMembersCanCreatePagesOk returns a tuple with the MembersCanCreatePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePages

`func (o *OrganizationFull) SetMembersCanCreatePages(v bool)`

SetMembersCanCreatePages sets MembersCanCreatePages field to given value.

### HasMembersCanCreatePages

`func (o *OrganizationFull) HasMembersCanCreatePages() bool`

HasMembersCanCreatePages returns a boolean if a field has been set.

### GetMembersCanCreatePublicPages

`func (o *OrganizationFull) GetMembersCanCreatePublicPages() bool`

GetMembersCanCreatePublicPages returns the MembersCanCreatePublicPages field if non-nil, zero value otherwise.

### GetMembersCanCreatePublicPagesOk

`func (o *OrganizationFull) GetMembersCanCreatePublicPagesOk() (*bool, bool)`

GetMembersCanCreatePublicPagesOk returns a tuple with the MembersCanCreatePublicPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePublicPages

`func (o *OrganizationFull) SetMembersCanCreatePublicPages(v bool)`

SetMembersCanCreatePublicPages sets MembersCanCreatePublicPages field to given value.

### HasMembersCanCreatePublicPages

`func (o *OrganizationFull) HasMembersCanCreatePublicPages() bool`

HasMembersCanCreatePublicPages returns a boolean if a field has been set.

### GetMembersCanCreatePrivatePages

`func (o *OrganizationFull) GetMembersCanCreatePrivatePages() bool`

GetMembersCanCreatePrivatePages returns the MembersCanCreatePrivatePages field if non-nil, zero value otherwise.

### GetMembersCanCreatePrivatePagesOk

`func (o *OrganizationFull) GetMembersCanCreatePrivatePagesOk() (*bool, bool)`

GetMembersCanCreatePrivatePagesOk returns a tuple with the MembersCanCreatePrivatePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePrivatePages

`func (o *OrganizationFull) SetMembersCanCreatePrivatePages(v bool)`

SetMembersCanCreatePrivatePages sets MembersCanCreatePrivatePages field to given value.

### HasMembersCanCreatePrivatePages

`func (o *OrganizationFull) HasMembersCanCreatePrivatePages() bool`

HasMembersCanCreatePrivatePages returns a boolean if a field has been set.

### GetMembersCanForkPrivateRepositories

`func (o *OrganizationFull) GetMembersCanForkPrivateRepositories() bool`

GetMembersCanForkPrivateRepositories returns the MembersCanForkPrivateRepositories field if non-nil, zero value otherwise.

### GetMembersCanForkPrivateRepositoriesOk

`func (o *OrganizationFull) GetMembersCanForkPrivateRepositoriesOk() (*bool, bool)`

GetMembersCanForkPrivateRepositoriesOk returns a tuple with the MembersCanForkPrivateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanForkPrivateRepositories

`func (o *OrganizationFull) SetMembersCanForkPrivateRepositories(v bool)`

SetMembersCanForkPrivateRepositories sets MembersCanForkPrivateRepositories field to given value.

### HasMembersCanForkPrivateRepositories

`func (o *OrganizationFull) HasMembersCanForkPrivateRepositories() bool`

HasMembersCanForkPrivateRepositories returns a boolean if a field has been set.

### SetMembersCanForkPrivateRepositoriesNil

`func (o *OrganizationFull) SetMembersCanForkPrivateRepositoriesNil(b bool)`

 SetMembersCanForkPrivateRepositoriesNil sets the value for MembersCanForkPrivateRepositories to be an explicit nil

### UnsetMembersCanForkPrivateRepositories
`func (o *OrganizationFull) UnsetMembersCanForkPrivateRepositories()`

UnsetMembersCanForkPrivateRepositories ensures that no value is present for MembersCanForkPrivateRepositories, not even an explicit nil
### GetWebCommitSignoffRequired

`func (o *OrganizationFull) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *OrganizationFull) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *OrganizationFull) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *OrganizationFull) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAdvancedSecurityEnabledForNewRepositories

`func (o *OrganizationFull) GetAdvancedSecurityEnabledForNewRepositories() bool`

GetAdvancedSecurityEnabledForNewRepositories returns the AdvancedSecurityEnabledForNewRepositories field if non-nil, zero value otherwise.

### GetAdvancedSecurityEnabledForNewRepositoriesOk

`func (o *OrganizationFull) GetAdvancedSecurityEnabledForNewRepositoriesOk() (*bool, bool)`

GetAdvancedSecurityEnabledForNewRepositoriesOk returns a tuple with the AdvancedSecurityEnabledForNewRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSecurityEnabledForNewRepositories

`func (o *OrganizationFull) SetAdvancedSecurityEnabledForNewRepositories(v bool)`

SetAdvancedSecurityEnabledForNewRepositories sets AdvancedSecurityEnabledForNewRepositories field to given value.

### HasAdvancedSecurityEnabledForNewRepositories

`func (o *OrganizationFull) HasAdvancedSecurityEnabledForNewRepositories() bool`

HasAdvancedSecurityEnabledForNewRepositories returns a boolean if a field has been set.

### GetDependabotAlertsEnabledForNewRepositories

`func (o *OrganizationFull) GetDependabotAlertsEnabledForNewRepositories() bool`

GetDependabotAlertsEnabledForNewRepositories returns the DependabotAlertsEnabledForNewRepositories field if non-nil, zero value otherwise.

### GetDependabotAlertsEnabledForNewRepositoriesOk

`func (o *OrganizationFull) GetDependabotAlertsEnabledForNewRepositoriesOk() (*bool, bool)`

GetDependabotAlertsEnabledForNewRepositoriesOk returns a tuple with the DependabotAlertsEnabledForNewRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependabotAlertsEnabledForNewRepositories

`func (o *OrganizationFull) SetDependabotAlertsEnabledForNewRepositories(v bool)`

SetDependabotAlertsEnabledForNewRepositories sets DependabotAlertsEnabledForNewRepositories field to given value.

### HasDependabotAlertsEnabledForNewRepositories

`func (o *OrganizationFull) HasDependabotAlertsEnabledForNewRepositories() bool`

HasDependabotAlertsEnabledForNewRepositories returns a boolean if a field has been set.

### GetDependabotSecurityUpdatesEnabledForNewRepositories

`func (o *OrganizationFull) GetDependabotSecurityUpdatesEnabledForNewRepositories() bool`

GetDependabotSecurityUpdatesEnabledForNewRepositories returns the DependabotSecurityUpdatesEnabledForNewRepositories field if non-nil, zero value otherwise.

### GetDependabotSecurityUpdatesEnabledForNewRepositoriesOk

`func (o *OrganizationFull) GetDependabotSecurityUpdatesEnabledForNewRepositoriesOk() (*bool, bool)`

GetDependabotSecurityUpdatesEnabledForNewRepositoriesOk returns a tuple with the DependabotSecurityUpdatesEnabledForNewRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependabotSecurityUpdatesEnabledForNewRepositories

`func (o *OrganizationFull) SetDependabotSecurityUpdatesEnabledForNewRepositories(v bool)`

SetDependabotSecurityUpdatesEnabledForNewRepositories sets DependabotSecurityUpdatesEnabledForNewRepositories field to given value.

### HasDependabotSecurityUpdatesEnabledForNewRepositories

`func (o *OrganizationFull) HasDependabotSecurityUpdatesEnabledForNewRepositories() bool`

HasDependabotSecurityUpdatesEnabledForNewRepositories returns a boolean if a field has been set.

### GetDependencyGraphEnabledForNewRepositories

`func (o *OrganizationFull) GetDependencyGraphEnabledForNewRepositories() bool`

GetDependencyGraphEnabledForNewRepositories returns the DependencyGraphEnabledForNewRepositories field if non-nil, zero value otherwise.

### GetDependencyGraphEnabledForNewRepositoriesOk

`func (o *OrganizationFull) GetDependencyGraphEnabledForNewRepositoriesOk() (*bool, bool)`

GetDependencyGraphEnabledForNewRepositoriesOk returns a tuple with the DependencyGraphEnabledForNewRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyGraphEnabledForNewRepositories

`func (o *OrganizationFull) SetDependencyGraphEnabledForNewRepositories(v bool)`

SetDependencyGraphEnabledForNewRepositories sets DependencyGraphEnabledForNewRepositories field to given value.

### HasDependencyGraphEnabledForNewRepositories

`func (o *OrganizationFull) HasDependencyGraphEnabledForNewRepositories() bool`

HasDependencyGraphEnabledForNewRepositories returns a boolean if a field has been set.

### GetSecretScanningEnabledForNewRepositories

`func (o *OrganizationFull) GetSecretScanningEnabledForNewRepositories() bool`

GetSecretScanningEnabledForNewRepositories returns the SecretScanningEnabledForNewRepositories field if non-nil, zero value otherwise.

### GetSecretScanningEnabledForNewRepositoriesOk

`func (o *OrganizationFull) GetSecretScanningEnabledForNewRepositoriesOk() (*bool, bool)`

GetSecretScanningEnabledForNewRepositoriesOk returns a tuple with the SecretScanningEnabledForNewRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanningEnabledForNewRepositories

`func (o *OrganizationFull) SetSecretScanningEnabledForNewRepositories(v bool)`

SetSecretScanningEnabledForNewRepositories sets SecretScanningEnabledForNewRepositories field to given value.

### HasSecretScanningEnabledForNewRepositories

`func (o *OrganizationFull) HasSecretScanningEnabledForNewRepositories() bool`

HasSecretScanningEnabledForNewRepositories returns a boolean if a field has been set.

### GetSecretScanningPushProtectionEnabledForNewRepositories

`func (o *OrganizationFull) GetSecretScanningPushProtectionEnabledForNewRepositories() bool`

GetSecretScanningPushProtectionEnabledForNewRepositories returns the SecretScanningPushProtectionEnabledForNewRepositories field if non-nil, zero value otherwise.

### GetSecretScanningPushProtectionEnabledForNewRepositoriesOk

`func (o *OrganizationFull) GetSecretScanningPushProtectionEnabledForNewRepositoriesOk() (*bool, bool)`

GetSecretScanningPushProtectionEnabledForNewRepositoriesOk returns a tuple with the SecretScanningPushProtectionEnabledForNewRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanningPushProtectionEnabledForNewRepositories

`func (o *OrganizationFull) SetSecretScanningPushProtectionEnabledForNewRepositories(v bool)`

SetSecretScanningPushProtectionEnabledForNewRepositories sets SecretScanningPushProtectionEnabledForNewRepositories field to given value.

### HasSecretScanningPushProtectionEnabledForNewRepositories

`func (o *OrganizationFull) HasSecretScanningPushProtectionEnabledForNewRepositories() bool`

HasSecretScanningPushProtectionEnabledForNewRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


