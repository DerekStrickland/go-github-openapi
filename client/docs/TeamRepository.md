# TeamRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the repository | 
**NodeId** | **string** |  | 
**Name** | **string** | The name of the repository. | 
**FullName** | **string** |  | 
**License** | [**NullableNullableLicenseSimple**](NullableLicenseSimple.md) |  | 
**Forks** | **int32** |  | 
**Permissions** | Pointer to [**RepositoryPermissions**](RepositoryPermissions.md) |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Owner** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Private** | **bool** | Whether the repository is private or public. | [default to false]
**HtmlUrl** | **string** |  | 
**Description** | **NullableString** |  | 
**Fork** | **bool** |  | 
**Url** | **string** |  | 
**ArchiveUrl** | **string** |  | 
**AssigneesUrl** | **string** |  | 
**BlobsUrl** | **string** |  | 
**BranchesUrl** | **string** |  | 
**CollaboratorsUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**CommitsUrl** | **string** |  | 
**CompareUrl** | **string** |  | 
**ContentsUrl** | **string** |  | 
**ContributorsUrl** | **string** |  | 
**DeploymentsUrl** | **string** |  | 
**DownloadsUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**ForksUrl** | **string** |  | 
**GitCommitsUrl** | **string** |  | 
**GitRefsUrl** | **string** |  | 
**GitTagsUrl** | **string** |  | 
**GitUrl** | **string** |  | 
**IssueCommentUrl** | **string** |  | 
**IssueEventsUrl** | **string** |  | 
**IssuesUrl** | **string** |  | 
**KeysUrl** | **string** |  | 
**LabelsUrl** | **string** |  | 
**LanguagesUrl** | **string** |  | 
**MergesUrl** | **string** |  | 
**MilestonesUrl** | **string** |  | 
**NotificationsUrl** | **string** |  | 
**PullsUrl** | **string** |  | 
**ReleasesUrl** | **string** |  | 
**SshUrl** | **string** |  | 
**StargazersUrl** | **string** |  | 
**StatusesUrl** | **string** |  | 
**SubscribersUrl** | **string** |  | 
**SubscriptionUrl** | **string** |  | 
**TagsUrl** | **string** |  | 
**TeamsUrl** | **string** |  | 
**TreesUrl** | **string** |  | 
**CloneUrl** | **string** |  | 
**MirrorUrl** | **NullableString** |  | 
**HooksUrl** | **string** |  | 
**SvnUrl** | **string** |  | 
**Homepage** | **NullableString** |  | 
**Language** | **NullableString** |  | 
**ForksCount** | **int32** |  | 
**StargazersCount** | **int32** |  | 
**WatchersCount** | **int32** |  | 
**Size** | **int32** |  | 
**DefaultBranch** | **string** | The default branch of the repository. | 
**OpenIssuesCount** | **int32** |  | 
**IsTemplate** | Pointer to **bool** | Whether this repository acts as a template that can be used to generate new repositories. | [optional] [default to false]
**Topics** | Pointer to **[]string** |  | [optional] 
**HasIssues** | **bool** | Whether issues are enabled. | [default to true]
**HasProjects** | **bool** | Whether projects are enabled. | [default to true]
**HasWiki** | **bool** | Whether the wiki is enabled. | [default to true]
**HasPages** | **bool** |  | 
**HasDownloads** | **bool** | Whether downloads are enabled. | [default to true]
**Archived** | **bool** | Whether the repository is archived. | [default to false]
**Disabled** | **bool** | Returns whether or not this repository disabled. | 
**Visibility** | Pointer to **string** | The repository visibility: public, private, or internal. | [optional] [default to "public"]
**PushedAt** | **NullableTime** |  | 
**CreatedAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 
**AllowRebaseMerge** | Pointer to **bool** | Whether to allow rebase merges for pull requests. | [optional] [default to true]
**TemplateRepository** | Pointer to [**NullableNullableRepository**](NullableRepository.md) |  | [optional] 
**TempCloneToken** | Pointer to **string** |  | [optional] 
**AllowSquashMerge** | Pointer to **bool** | Whether to allow squash merges for pull requests. | [optional] [default to true]
**AllowAutoMerge** | Pointer to **bool** | Whether to allow Auto-merge to be used on pull requests. | [optional] [default to false]
**DeleteBranchOnMerge** | Pointer to **bool** | Whether to delete head branches when pull requests are merged | [optional] [default to false]
**AllowMergeCommit** | Pointer to **bool** | Whether to allow merge commits for pull requests. | [optional] [default to true]
**AllowForking** | Pointer to **bool** | Whether to allow forking this repo | [optional] [default to false]
**WebCommitSignoffRequired** | Pointer to **bool** | Whether to require contributors to sign off on web-based commits | [optional] [default to false]
**SubscribersCount** | Pointer to **int32** |  | [optional] 
**NetworkCount** | Pointer to **int32** |  | [optional] 
**OpenIssues** | **int32** |  | 
**Watchers** | **int32** |  | 
**MasterBranch** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamRepository

`func NewTeamRepository(id int32, nodeId string, name string, fullName string, license NullableNullableLicenseSimple, forks int32, owner NullableNullableSimpleUser, private bool, htmlUrl string, description NullableString, fork bool, url string, archiveUrl string, assigneesUrl string, blobsUrl string, branchesUrl string, collaboratorsUrl string, commentsUrl string, commitsUrl string, compareUrl string, contentsUrl string, contributorsUrl string, deploymentsUrl string, downloadsUrl string, eventsUrl string, forksUrl string, gitCommitsUrl string, gitRefsUrl string, gitTagsUrl string, gitUrl string, issueCommentUrl string, issueEventsUrl string, issuesUrl string, keysUrl string, labelsUrl string, languagesUrl string, mergesUrl string, milestonesUrl string, notificationsUrl string, pullsUrl string, releasesUrl string, sshUrl string, stargazersUrl string, statusesUrl string, subscribersUrl string, subscriptionUrl string, tagsUrl string, teamsUrl string, treesUrl string, cloneUrl string, mirrorUrl NullableString, hooksUrl string, svnUrl string, homepage NullableString, language NullableString, forksCount int32, stargazersCount int32, watchersCount int32, size int32, defaultBranch string, openIssuesCount int32, hasIssues bool, hasProjects bool, hasWiki bool, hasPages bool, hasDownloads bool, archived bool, disabled bool, pushedAt NullableTime, createdAt NullableTime, updatedAt NullableTime, openIssues int32, watchers int32, ) *TeamRepository`

NewTeamRepository instantiates a new TeamRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRepositoryWithDefaults

`func NewTeamRepositoryWithDefaults() *TeamRepository`

NewTeamRepositoryWithDefaults instantiates a new TeamRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamRepository) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamRepository) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamRepository) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TeamRepository) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TeamRepository) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TeamRepository) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *TeamRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamRepository) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *TeamRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *TeamRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *TeamRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetLicense

`func (o *TeamRepository) GetLicense() NullableLicenseSimple`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *TeamRepository) GetLicenseOk() (*NullableLicenseSimple, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *TeamRepository) SetLicense(v NullableLicenseSimple)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *TeamRepository) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *TeamRepository) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetForks

`func (o *TeamRepository) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *TeamRepository) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *TeamRepository) SetForks(v int32)`

SetForks sets Forks field to given value.


### GetPermissions

`func (o *TeamRepository) GetPermissions() RepositoryPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TeamRepository) GetPermissionsOk() (*RepositoryPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TeamRepository) SetPermissions(v RepositoryPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TeamRepository) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRoleName

`func (o *TeamRepository) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *TeamRepository) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *TeamRepository) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *TeamRepository) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetOwner

`func (o *TeamRepository) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TeamRepository) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TeamRepository) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *TeamRepository) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *TeamRepository) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPrivate

`func (o *TeamRepository) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *TeamRepository) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *TeamRepository) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetHtmlUrl

`func (o *TeamRepository) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TeamRepository) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TeamRepository) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDescription

`func (o *TeamRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamRepository) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *TeamRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TeamRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFork

`func (o *TeamRepository) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *TeamRepository) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *TeamRepository) SetFork(v bool)`

SetFork sets Fork field to given value.


### GetUrl

`func (o *TeamRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetArchiveUrl

`func (o *TeamRepository) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *TeamRepository) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *TeamRepository) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.


### GetAssigneesUrl

`func (o *TeamRepository) GetAssigneesUrl() string`

GetAssigneesUrl returns the AssigneesUrl field if non-nil, zero value otherwise.

### GetAssigneesUrlOk

`func (o *TeamRepository) GetAssigneesUrlOk() (*string, bool)`

GetAssigneesUrlOk returns a tuple with the AssigneesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesUrl

`func (o *TeamRepository) SetAssigneesUrl(v string)`

SetAssigneesUrl sets AssigneesUrl field to given value.


### GetBlobsUrl

`func (o *TeamRepository) GetBlobsUrl() string`

GetBlobsUrl returns the BlobsUrl field if non-nil, zero value otherwise.

### GetBlobsUrlOk

`func (o *TeamRepository) GetBlobsUrlOk() (*string, bool)`

GetBlobsUrlOk returns a tuple with the BlobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsUrl

`func (o *TeamRepository) SetBlobsUrl(v string)`

SetBlobsUrl sets BlobsUrl field to given value.


### GetBranchesUrl

`func (o *TeamRepository) GetBranchesUrl() string`

GetBranchesUrl returns the BranchesUrl field if non-nil, zero value otherwise.

### GetBranchesUrlOk

`func (o *TeamRepository) GetBranchesUrlOk() (*string, bool)`

GetBranchesUrlOk returns a tuple with the BranchesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesUrl

`func (o *TeamRepository) SetBranchesUrl(v string)`

SetBranchesUrl sets BranchesUrl field to given value.


### GetCollaboratorsUrl

`func (o *TeamRepository) GetCollaboratorsUrl() string`

GetCollaboratorsUrl returns the CollaboratorsUrl field if non-nil, zero value otherwise.

### GetCollaboratorsUrlOk

`func (o *TeamRepository) GetCollaboratorsUrlOk() (*string, bool)`

GetCollaboratorsUrlOk returns a tuple with the CollaboratorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorsUrl

`func (o *TeamRepository) SetCollaboratorsUrl(v string)`

SetCollaboratorsUrl sets CollaboratorsUrl field to given value.


### GetCommentsUrl

`func (o *TeamRepository) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *TeamRepository) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *TeamRepository) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCommitsUrl

`func (o *TeamRepository) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *TeamRepository) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *TeamRepository) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetCompareUrl

`func (o *TeamRepository) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *TeamRepository) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *TeamRepository) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.


### GetContentsUrl

`func (o *TeamRepository) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *TeamRepository) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *TeamRepository) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetContributorsUrl

`func (o *TeamRepository) GetContributorsUrl() string`

GetContributorsUrl returns the ContributorsUrl field if non-nil, zero value otherwise.

### GetContributorsUrlOk

`func (o *TeamRepository) GetContributorsUrlOk() (*string, bool)`

GetContributorsUrlOk returns a tuple with the ContributorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsUrl

`func (o *TeamRepository) SetContributorsUrl(v string)`

SetContributorsUrl sets ContributorsUrl field to given value.


### GetDeploymentsUrl

`func (o *TeamRepository) GetDeploymentsUrl() string`

GetDeploymentsUrl returns the DeploymentsUrl field if non-nil, zero value otherwise.

### GetDeploymentsUrlOk

`func (o *TeamRepository) GetDeploymentsUrlOk() (*string, bool)`

GetDeploymentsUrlOk returns a tuple with the DeploymentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsUrl

`func (o *TeamRepository) SetDeploymentsUrl(v string)`

SetDeploymentsUrl sets DeploymentsUrl field to given value.


### GetDownloadsUrl

`func (o *TeamRepository) GetDownloadsUrl() string`

GetDownloadsUrl returns the DownloadsUrl field if non-nil, zero value otherwise.

### GetDownloadsUrlOk

`func (o *TeamRepository) GetDownloadsUrlOk() (*string, bool)`

GetDownloadsUrlOk returns a tuple with the DownloadsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsUrl

`func (o *TeamRepository) SetDownloadsUrl(v string)`

SetDownloadsUrl sets DownloadsUrl field to given value.


### GetEventsUrl

`func (o *TeamRepository) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *TeamRepository) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *TeamRepository) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetForksUrl

`func (o *TeamRepository) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *TeamRepository) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *TeamRepository) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.


### GetGitCommitsUrl

`func (o *TeamRepository) GetGitCommitsUrl() string`

GetGitCommitsUrl returns the GitCommitsUrl field if non-nil, zero value otherwise.

### GetGitCommitsUrlOk

`func (o *TeamRepository) GetGitCommitsUrlOk() (*string, bool)`

GetGitCommitsUrlOk returns a tuple with the GitCommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitsUrl

`func (o *TeamRepository) SetGitCommitsUrl(v string)`

SetGitCommitsUrl sets GitCommitsUrl field to given value.


### GetGitRefsUrl

`func (o *TeamRepository) GetGitRefsUrl() string`

GetGitRefsUrl returns the GitRefsUrl field if non-nil, zero value otherwise.

### GetGitRefsUrlOk

`func (o *TeamRepository) GetGitRefsUrlOk() (*string, bool)`

GetGitRefsUrlOk returns a tuple with the GitRefsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefsUrl

`func (o *TeamRepository) SetGitRefsUrl(v string)`

SetGitRefsUrl sets GitRefsUrl field to given value.


### GetGitTagsUrl

`func (o *TeamRepository) GetGitTagsUrl() string`

GetGitTagsUrl returns the GitTagsUrl field if non-nil, zero value otherwise.

### GetGitTagsUrlOk

`func (o *TeamRepository) GetGitTagsUrlOk() (*string, bool)`

GetGitTagsUrlOk returns a tuple with the GitTagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagsUrl

`func (o *TeamRepository) SetGitTagsUrl(v string)`

SetGitTagsUrl sets GitTagsUrl field to given value.


### GetGitUrl

`func (o *TeamRepository) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *TeamRepository) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *TeamRepository) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### GetIssueCommentUrl

`func (o *TeamRepository) GetIssueCommentUrl() string`

GetIssueCommentUrl returns the IssueCommentUrl field if non-nil, zero value otherwise.

### GetIssueCommentUrlOk

`func (o *TeamRepository) GetIssueCommentUrlOk() (*string, bool)`

GetIssueCommentUrlOk returns a tuple with the IssueCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCommentUrl

`func (o *TeamRepository) SetIssueCommentUrl(v string)`

SetIssueCommentUrl sets IssueCommentUrl field to given value.


### GetIssueEventsUrl

`func (o *TeamRepository) GetIssueEventsUrl() string`

GetIssueEventsUrl returns the IssueEventsUrl field if non-nil, zero value otherwise.

### GetIssueEventsUrlOk

`func (o *TeamRepository) GetIssueEventsUrlOk() (*string, bool)`

GetIssueEventsUrlOk returns a tuple with the IssueEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEventsUrl

`func (o *TeamRepository) SetIssueEventsUrl(v string)`

SetIssueEventsUrl sets IssueEventsUrl field to given value.


### GetIssuesUrl

`func (o *TeamRepository) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *TeamRepository) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *TeamRepository) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetKeysUrl

`func (o *TeamRepository) GetKeysUrl() string`

GetKeysUrl returns the KeysUrl field if non-nil, zero value otherwise.

### GetKeysUrlOk

`func (o *TeamRepository) GetKeysUrlOk() (*string, bool)`

GetKeysUrlOk returns a tuple with the KeysUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysUrl

`func (o *TeamRepository) SetKeysUrl(v string)`

SetKeysUrl sets KeysUrl field to given value.


### GetLabelsUrl

`func (o *TeamRepository) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *TeamRepository) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *TeamRepository) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetLanguagesUrl

`func (o *TeamRepository) GetLanguagesUrl() string`

GetLanguagesUrl returns the LanguagesUrl field if non-nil, zero value otherwise.

### GetLanguagesUrlOk

`func (o *TeamRepository) GetLanguagesUrlOk() (*string, bool)`

GetLanguagesUrlOk returns a tuple with the LanguagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesUrl

`func (o *TeamRepository) SetLanguagesUrl(v string)`

SetLanguagesUrl sets LanguagesUrl field to given value.


### GetMergesUrl

`func (o *TeamRepository) GetMergesUrl() string`

GetMergesUrl returns the MergesUrl field if non-nil, zero value otherwise.

### GetMergesUrlOk

`func (o *TeamRepository) GetMergesUrlOk() (*string, bool)`

GetMergesUrlOk returns a tuple with the MergesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesUrl

`func (o *TeamRepository) SetMergesUrl(v string)`

SetMergesUrl sets MergesUrl field to given value.


### GetMilestonesUrl

`func (o *TeamRepository) GetMilestonesUrl() string`

GetMilestonesUrl returns the MilestonesUrl field if non-nil, zero value otherwise.

### GetMilestonesUrlOk

`func (o *TeamRepository) GetMilestonesUrlOk() (*string, bool)`

GetMilestonesUrlOk returns a tuple with the MilestonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestonesUrl

`func (o *TeamRepository) SetMilestonesUrl(v string)`

SetMilestonesUrl sets MilestonesUrl field to given value.


### GetNotificationsUrl

`func (o *TeamRepository) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *TeamRepository) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *TeamRepository) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.


### GetPullsUrl

`func (o *TeamRepository) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *TeamRepository) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *TeamRepository) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.


### GetReleasesUrl

`func (o *TeamRepository) GetReleasesUrl() string`

GetReleasesUrl returns the ReleasesUrl field if non-nil, zero value otherwise.

### GetReleasesUrlOk

`func (o *TeamRepository) GetReleasesUrlOk() (*string, bool)`

GetReleasesUrlOk returns a tuple with the ReleasesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesUrl

`func (o *TeamRepository) SetReleasesUrl(v string)`

SetReleasesUrl sets ReleasesUrl field to given value.


### GetSshUrl

`func (o *TeamRepository) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *TeamRepository) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *TeamRepository) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.


### GetStargazersUrl

`func (o *TeamRepository) GetStargazersUrl() string`

GetStargazersUrl returns the StargazersUrl field if non-nil, zero value otherwise.

### GetStargazersUrlOk

`func (o *TeamRepository) GetStargazersUrlOk() (*string, bool)`

GetStargazersUrlOk returns a tuple with the StargazersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersUrl

`func (o *TeamRepository) SetStargazersUrl(v string)`

SetStargazersUrl sets StargazersUrl field to given value.


### GetStatusesUrl

`func (o *TeamRepository) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *TeamRepository) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *TeamRepository) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetSubscribersUrl

`func (o *TeamRepository) GetSubscribersUrl() string`

GetSubscribersUrl returns the SubscribersUrl field if non-nil, zero value otherwise.

### GetSubscribersUrlOk

`func (o *TeamRepository) GetSubscribersUrlOk() (*string, bool)`

GetSubscribersUrlOk returns a tuple with the SubscribersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersUrl

`func (o *TeamRepository) SetSubscribersUrl(v string)`

SetSubscribersUrl sets SubscribersUrl field to given value.


### GetSubscriptionUrl

`func (o *TeamRepository) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *TeamRepository) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *TeamRepository) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetTagsUrl

`func (o *TeamRepository) GetTagsUrl() string`

GetTagsUrl returns the TagsUrl field if non-nil, zero value otherwise.

### GetTagsUrlOk

`func (o *TeamRepository) GetTagsUrlOk() (*string, bool)`

GetTagsUrlOk returns a tuple with the TagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsUrl

`func (o *TeamRepository) SetTagsUrl(v string)`

SetTagsUrl sets TagsUrl field to given value.


### GetTeamsUrl

`func (o *TeamRepository) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *TeamRepository) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *TeamRepository) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetTreesUrl

`func (o *TeamRepository) GetTreesUrl() string`

GetTreesUrl returns the TreesUrl field if non-nil, zero value otherwise.

### GetTreesUrlOk

`func (o *TeamRepository) GetTreesUrlOk() (*string, bool)`

GetTreesUrlOk returns a tuple with the TreesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreesUrl

`func (o *TeamRepository) SetTreesUrl(v string)`

SetTreesUrl sets TreesUrl field to given value.


### GetCloneUrl

`func (o *TeamRepository) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *TeamRepository) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *TeamRepository) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetMirrorUrl

`func (o *TeamRepository) GetMirrorUrl() string`

GetMirrorUrl returns the MirrorUrl field if non-nil, zero value otherwise.

### GetMirrorUrlOk

`func (o *TeamRepository) GetMirrorUrlOk() (*string, bool)`

GetMirrorUrlOk returns a tuple with the MirrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorUrl

`func (o *TeamRepository) SetMirrorUrl(v string)`

SetMirrorUrl sets MirrorUrl field to given value.


### SetMirrorUrlNil

`func (o *TeamRepository) SetMirrorUrlNil(b bool)`

 SetMirrorUrlNil sets the value for MirrorUrl to be an explicit nil

### UnsetMirrorUrl
`func (o *TeamRepository) UnsetMirrorUrl()`

UnsetMirrorUrl ensures that no value is present for MirrorUrl, not even an explicit nil
### GetHooksUrl

`func (o *TeamRepository) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *TeamRepository) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *TeamRepository) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetSvnUrl

`func (o *TeamRepository) GetSvnUrl() string`

GetSvnUrl returns the SvnUrl field if non-nil, zero value otherwise.

### GetSvnUrlOk

`func (o *TeamRepository) GetSvnUrlOk() (*string, bool)`

GetSvnUrlOk returns a tuple with the SvnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnUrl

`func (o *TeamRepository) SetSvnUrl(v string)`

SetSvnUrl sets SvnUrl field to given value.


### GetHomepage

`func (o *TeamRepository) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *TeamRepository) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *TeamRepository) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.


### SetHomepageNil

`func (o *TeamRepository) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *TeamRepository) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil
### GetLanguage

`func (o *TeamRepository) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TeamRepository) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TeamRepository) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### SetLanguageNil

`func (o *TeamRepository) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *TeamRepository) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetForksCount

`func (o *TeamRepository) GetForksCount() int32`

GetForksCount returns the ForksCount field if non-nil, zero value otherwise.

### GetForksCountOk

`func (o *TeamRepository) GetForksCountOk() (*int32, bool)`

GetForksCountOk returns a tuple with the ForksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksCount

`func (o *TeamRepository) SetForksCount(v int32)`

SetForksCount sets ForksCount field to given value.


### GetStargazersCount

`func (o *TeamRepository) GetStargazersCount() int32`

GetStargazersCount returns the StargazersCount field if non-nil, zero value otherwise.

### GetStargazersCountOk

`func (o *TeamRepository) GetStargazersCountOk() (*int32, bool)`

GetStargazersCountOk returns a tuple with the StargazersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersCount

`func (o *TeamRepository) SetStargazersCount(v int32)`

SetStargazersCount sets StargazersCount field to given value.


### GetWatchersCount

`func (o *TeamRepository) GetWatchersCount() int32`

GetWatchersCount returns the WatchersCount field if non-nil, zero value otherwise.

### GetWatchersCountOk

`func (o *TeamRepository) GetWatchersCountOk() (*int32, bool)`

GetWatchersCountOk returns a tuple with the WatchersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchersCount

`func (o *TeamRepository) SetWatchersCount(v int32)`

SetWatchersCount sets WatchersCount field to given value.


### GetSize

`func (o *TeamRepository) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TeamRepository) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TeamRepository) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDefaultBranch

`func (o *TeamRepository) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *TeamRepository) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *TeamRepository) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetOpenIssuesCount

`func (o *TeamRepository) GetOpenIssuesCount() int32`

GetOpenIssuesCount returns the OpenIssuesCount field if non-nil, zero value otherwise.

### GetOpenIssuesCountOk

`func (o *TeamRepository) GetOpenIssuesCountOk() (*int32, bool)`

GetOpenIssuesCountOk returns a tuple with the OpenIssuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssuesCount

`func (o *TeamRepository) SetOpenIssuesCount(v int32)`

SetOpenIssuesCount sets OpenIssuesCount field to given value.


### GetIsTemplate

`func (o *TeamRepository) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *TeamRepository) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *TeamRepository) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *TeamRepository) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetTopics

`func (o *TeamRepository) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *TeamRepository) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *TeamRepository) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *TeamRepository) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetHasIssues

`func (o *TeamRepository) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *TeamRepository) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *TeamRepository) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.


### GetHasProjects

`func (o *TeamRepository) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *TeamRepository) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *TeamRepository) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.


### GetHasWiki

`func (o *TeamRepository) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *TeamRepository) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *TeamRepository) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.


### GetHasPages

`func (o *TeamRepository) GetHasPages() bool`

GetHasPages returns the HasPages field if non-nil, zero value otherwise.

### GetHasPagesOk

`func (o *TeamRepository) GetHasPagesOk() (*bool, bool)`

GetHasPagesOk returns a tuple with the HasPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPages

`func (o *TeamRepository) SetHasPages(v bool)`

SetHasPages sets HasPages field to given value.


### GetHasDownloads

`func (o *TeamRepository) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *TeamRepository) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *TeamRepository) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.


### GetArchived

`func (o *TeamRepository) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TeamRepository) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TeamRepository) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetDisabled

`func (o *TeamRepository) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TeamRepository) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TeamRepository) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetVisibility

`func (o *TeamRepository) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TeamRepository) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TeamRepository) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TeamRepository) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetPushedAt

`func (o *TeamRepository) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *TeamRepository) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *TeamRepository) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.


### SetPushedAtNil

`func (o *TeamRepository) SetPushedAtNil(b bool)`

 SetPushedAtNil sets the value for PushedAt to be an explicit nil

### UnsetPushedAt
`func (o *TeamRepository) UnsetPushedAt()`

UnsetPushedAt ensures that no value is present for PushedAt, not even an explicit nil
### GetCreatedAt

`func (o *TeamRepository) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamRepository) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamRepository) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *TeamRepository) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TeamRepository) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TeamRepository) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamRepository) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamRepository) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *TeamRepository) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TeamRepository) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAllowRebaseMerge

`func (o *TeamRepository) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *TeamRepository) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *TeamRepository) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *TeamRepository) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetTemplateRepository

`func (o *TeamRepository) GetTemplateRepository() NullableRepository`

GetTemplateRepository returns the TemplateRepository field if non-nil, zero value otherwise.

### GetTemplateRepositoryOk

`func (o *TeamRepository) GetTemplateRepositoryOk() (*NullableRepository, bool)`

GetTemplateRepositoryOk returns a tuple with the TemplateRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRepository

`func (o *TeamRepository) SetTemplateRepository(v NullableRepository)`

SetTemplateRepository sets TemplateRepository field to given value.

### HasTemplateRepository

`func (o *TeamRepository) HasTemplateRepository() bool`

HasTemplateRepository returns a boolean if a field has been set.

### SetTemplateRepositoryNil

`func (o *TeamRepository) SetTemplateRepositoryNil(b bool)`

 SetTemplateRepositoryNil sets the value for TemplateRepository to be an explicit nil

### UnsetTemplateRepository
`func (o *TeamRepository) UnsetTemplateRepository()`

UnsetTemplateRepository ensures that no value is present for TemplateRepository, not even an explicit nil
### GetTempCloneToken

`func (o *TeamRepository) GetTempCloneToken() string`

GetTempCloneToken returns the TempCloneToken field if non-nil, zero value otherwise.

### GetTempCloneTokenOk

`func (o *TeamRepository) GetTempCloneTokenOk() (*string, bool)`

GetTempCloneTokenOk returns a tuple with the TempCloneToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempCloneToken

`func (o *TeamRepository) SetTempCloneToken(v string)`

SetTempCloneToken sets TempCloneToken field to given value.

### HasTempCloneToken

`func (o *TeamRepository) HasTempCloneToken() bool`

HasTempCloneToken returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *TeamRepository) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *TeamRepository) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *TeamRepository) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *TeamRepository) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *TeamRepository) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *TeamRepository) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *TeamRepository) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *TeamRepository) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *TeamRepository) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *TeamRepository) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *TeamRepository) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *TeamRepository) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *TeamRepository) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *TeamRepository) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *TeamRepository) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *TeamRepository) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowForking

`func (o *TeamRepository) GetAllowForking() bool`

GetAllowForking returns the AllowForking field if non-nil, zero value otherwise.

### GetAllowForkingOk

`func (o *TeamRepository) GetAllowForkingOk() (*bool, bool)`

GetAllowForkingOk returns a tuple with the AllowForking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForking

`func (o *TeamRepository) SetAllowForking(v bool)`

SetAllowForking sets AllowForking field to given value.

### HasAllowForking

`func (o *TeamRepository) HasAllowForking() bool`

HasAllowForking returns a boolean if a field has been set.

### GetWebCommitSignoffRequired

`func (o *TeamRepository) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *TeamRepository) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *TeamRepository) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *TeamRepository) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.

### GetSubscribersCount

`func (o *TeamRepository) GetSubscribersCount() int32`

GetSubscribersCount returns the SubscribersCount field if non-nil, zero value otherwise.

### GetSubscribersCountOk

`func (o *TeamRepository) GetSubscribersCountOk() (*int32, bool)`

GetSubscribersCountOk returns a tuple with the SubscribersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersCount

`func (o *TeamRepository) SetSubscribersCount(v int32)`

SetSubscribersCount sets SubscribersCount field to given value.

### HasSubscribersCount

`func (o *TeamRepository) HasSubscribersCount() bool`

HasSubscribersCount returns a boolean if a field has been set.

### GetNetworkCount

`func (o *TeamRepository) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *TeamRepository) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *TeamRepository) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *TeamRepository) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetOpenIssues

`func (o *TeamRepository) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *TeamRepository) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *TeamRepository) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetWatchers

`func (o *TeamRepository) GetWatchers() int32`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *TeamRepository) GetWatchersOk() (*int32, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *TeamRepository) SetWatchers(v int32)`

SetWatchers sets Watchers field to given value.


### GetMasterBranch

`func (o *TeamRepository) GetMasterBranch() string`

GetMasterBranch returns the MasterBranch field if non-nil, zero value otherwise.

### GetMasterBranchOk

`func (o *TeamRepository) GetMasterBranchOk() (*string, bool)`

GetMasterBranchOk returns a tuple with the MasterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterBranch

`func (o *TeamRepository) SetMasterBranch(v string)`

SetMasterBranch sets MasterBranch field to given value.

### HasMasterBranch

`func (o *TeamRepository) HasMasterBranch() bool`

HasMasterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


