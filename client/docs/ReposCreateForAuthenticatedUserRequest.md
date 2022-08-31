# ReposCreateForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the repository. | 
**Description** | Pointer to **string** | A short description of the repository. | [optional] 
**Homepage** | Pointer to **string** | A URL with more information about the repository. | [optional] 
**Private** | Pointer to **bool** | Whether the repository is private. | [optional] [default to false]
**HasIssues** | Pointer to **bool** | Whether issues are enabled. | [optional] [default to true]
**HasProjects** | Pointer to **bool** | Whether projects are enabled. | [optional] [default to true]
**HasWiki** | Pointer to **bool** | Whether the wiki is enabled. | [optional] [default to true]
**TeamId** | Pointer to **int32** | The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization. | [optional] 
**AutoInit** | Pointer to **bool** | Whether the repository is initialized with a minimal README. | [optional] [default to false]
**GitignoreTemplate** | Pointer to **string** | The desired language or platform to apply to the .gitignore. | [optional] 
**LicenseTemplate** | Pointer to **string** | The license keyword of the open source license for this repository. | [optional] 
**AllowSquashMerge** | Pointer to **bool** | Whether to allow squash merges for pull requests. | [optional] [default to true]
**AllowMergeCommit** | Pointer to **bool** | Whether to allow merge commits for pull requests. | [optional] [default to true]
**AllowRebaseMerge** | Pointer to **bool** | Whether to allow rebase merges for pull requests. | [optional] [default to true]
**AllowAutoMerge** | Pointer to **bool** | Whether to allow Auto-merge to be used on pull requests. | [optional] [default to false]
**DeleteBranchOnMerge** | Pointer to **bool** | Whether to delete head branches when pull requests are merged | [optional] [default to false]
**SquashMergeCommitTitle** | Pointer to **string** | The default value for a squash merge commit title:  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;COMMIT_OR_PR_TITLE&#x60; - default to the commit&#39;s title (if only one commit) or the pull request&#39;s title (when more than one commit). | [optional] 
**SquashMergeCommitMessage** | Pointer to **string** | The default value for a squash merge commit message:  - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;COMMIT_MESSAGES&#x60; - default to the branch&#39;s commit messages. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**MergeCommitTitle** | Pointer to **string** | The default value for a merge commit title.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;MERGE_MESSAGE&#x60; - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name). | [optional] 
**MergeCommitMessage** | Pointer to **string** | The default value for a merge commit message.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**HasDownloads** | Pointer to **bool** | Whether downloads are enabled. | [optional] [default to true]
**IsTemplate** | Pointer to **bool** | Whether this repository acts as a template that can be used to generate new repositories. | [optional] [default to false]

## Methods

### NewReposCreateForAuthenticatedUserRequest

`func NewReposCreateForAuthenticatedUserRequest(name string, ) *ReposCreateForAuthenticatedUserRequest`

NewReposCreateForAuthenticatedUserRequest instantiates a new ReposCreateForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateForAuthenticatedUserRequestWithDefaults

`func NewReposCreateForAuthenticatedUserRequestWithDefaults() *ReposCreateForAuthenticatedUserRequest`

NewReposCreateForAuthenticatedUserRequestWithDefaults instantiates a new ReposCreateForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReposCreateForAuthenticatedUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateForAuthenticatedUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReposCreateForAuthenticatedUserRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposCreateForAuthenticatedUserRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposCreateForAuthenticatedUserRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomepage

`func (o *ReposCreateForAuthenticatedUserRequest) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *ReposCreateForAuthenticatedUserRequest) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *ReposCreateForAuthenticatedUserRequest) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetPrivate

`func (o *ReposCreateForAuthenticatedUserRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ReposCreateForAuthenticatedUserRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ReposCreateForAuthenticatedUserRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetHasIssues

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *ReposCreateForAuthenticatedUserRequest) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.

### HasHasIssues

`func (o *ReposCreateForAuthenticatedUserRequest) HasHasIssues() bool`

HasHasIssues returns a boolean if a field has been set.

### GetHasProjects

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *ReposCreateForAuthenticatedUserRequest) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.

### HasHasProjects

`func (o *ReposCreateForAuthenticatedUserRequest) HasHasProjects() bool`

HasHasProjects returns a boolean if a field has been set.

### GetHasWiki

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *ReposCreateForAuthenticatedUserRequest) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.

### HasHasWiki

`func (o *ReposCreateForAuthenticatedUserRequest) HasHasWiki() bool`

HasHasWiki returns a boolean if a field has been set.

### GetTeamId

`func (o *ReposCreateForAuthenticatedUserRequest) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ReposCreateForAuthenticatedUserRequest) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ReposCreateForAuthenticatedUserRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetAutoInit

`func (o *ReposCreateForAuthenticatedUserRequest) GetAutoInit() bool`

GetAutoInit returns the AutoInit field if non-nil, zero value otherwise.

### GetAutoInitOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetAutoInitOk() (*bool, bool)`

GetAutoInitOk returns a tuple with the AutoInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInit

`func (o *ReposCreateForAuthenticatedUserRequest) SetAutoInit(v bool)`

SetAutoInit sets AutoInit field to given value.

### HasAutoInit

`func (o *ReposCreateForAuthenticatedUserRequest) HasAutoInit() bool`

HasAutoInit returns a boolean if a field has been set.

### GetGitignoreTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) GetGitignoreTemplate() string`

GetGitignoreTemplate returns the GitignoreTemplate field if non-nil, zero value otherwise.

### GetGitignoreTemplateOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetGitignoreTemplateOk() (*string, bool)`

GetGitignoreTemplateOk returns a tuple with the GitignoreTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitignoreTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) SetGitignoreTemplate(v string)`

SetGitignoreTemplate sets GitignoreTemplate field to given value.

### HasGitignoreTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) HasGitignoreTemplate() bool`

HasGitignoreTemplate returns a boolean if a field has been set.

### GetLicenseTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) GetLicenseTemplate() string`

GetLicenseTemplate returns the LicenseTemplate field if non-nil, zero value otherwise.

### GetLicenseTemplateOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetLicenseTemplateOk() (*string, bool)`

GetLicenseTemplateOk returns a tuple with the LicenseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) SetLicenseTemplate(v string)`

SetLicenseTemplate sets LicenseTemplate field to given value.

### HasLicenseTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) HasLicenseTemplate() bool`

HasLicenseTemplate returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *ReposCreateForAuthenticatedUserRequest) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *ReposCreateForAuthenticatedUserRequest) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *ReposCreateForAuthenticatedUserRequest) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *ReposCreateForAuthenticatedUserRequest) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *ReposCreateForAuthenticatedUserRequest) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *ReposCreateForAuthenticatedUserRequest) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *ReposCreateForAuthenticatedUserRequest) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *ReposCreateForAuthenticatedUserRequest) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *ReposCreateForAuthenticatedUserRequest) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *ReposCreateForAuthenticatedUserRequest) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *ReposCreateForAuthenticatedUserRequest) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetSquashMergeCommitTitle

`func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitTitle() string`

GetSquashMergeCommitTitle returns the SquashMergeCommitTitle field if non-nil, zero value otherwise.

### GetSquashMergeCommitTitleOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitTitleOk() (*string, bool)`

GetSquashMergeCommitTitleOk returns a tuple with the SquashMergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitTitle

`func (o *ReposCreateForAuthenticatedUserRequest) SetSquashMergeCommitTitle(v string)`

SetSquashMergeCommitTitle sets SquashMergeCommitTitle field to given value.

### HasSquashMergeCommitTitle

`func (o *ReposCreateForAuthenticatedUserRequest) HasSquashMergeCommitTitle() bool`

HasSquashMergeCommitTitle returns a boolean if a field has been set.

### GetSquashMergeCommitMessage

`func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitMessage() string`

GetSquashMergeCommitMessage returns the SquashMergeCommitMessage field if non-nil, zero value otherwise.

### GetSquashMergeCommitMessageOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetSquashMergeCommitMessageOk() (*string, bool)`

GetSquashMergeCommitMessageOk returns a tuple with the SquashMergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitMessage

`func (o *ReposCreateForAuthenticatedUserRequest) SetSquashMergeCommitMessage(v string)`

SetSquashMergeCommitMessage sets SquashMergeCommitMessage field to given value.

### HasSquashMergeCommitMessage

`func (o *ReposCreateForAuthenticatedUserRequest) HasSquashMergeCommitMessage() bool`

HasSquashMergeCommitMessage returns a boolean if a field has been set.

### GetMergeCommitTitle

`func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitTitle() string`

GetMergeCommitTitle returns the MergeCommitTitle field if non-nil, zero value otherwise.

### GetMergeCommitTitleOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitTitleOk() (*string, bool)`

GetMergeCommitTitleOk returns a tuple with the MergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitTitle

`func (o *ReposCreateForAuthenticatedUserRequest) SetMergeCommitTitle(v string)`

SetMergeCommitTitle sets MergeCommitTitle field to given value.

### HasMergeCommitTitle

`func (o *ReposCreateForAuthenticatedUserRequest) HasMergeCommitTitle() bool`

HasMergeCommitTitle returns a boolean if a field has been set.

### GetMergeCommitMessage

`func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitMessage() string`

GetMergeCommitMessage returns the MergeCommitMessage field if non-nil, zero value otherwise.

### GetMergeCommitMessageOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetMergeCommitMessageOk() (*string, bool)`

GetMergeCommitMessageOk returns a tuple with the MergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitMessage

`func (o *ReposCreateForAuthenticatedUserRequest) SetMergeCommitMessage(v string)`

SetMergeCommitMessage sets MergeCommitMessage field to given value.

### HasMergeCommitMessage

`func (o *ReposCreateForAuthenticatedUserRequest) HasMergeCommitMessage() bool`

HasMergeCommitMessage returns a boolean if a field has been set.

### GetHasDownloads

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *ReposCreateForAuthenticatedUserRequest) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.

### HasHasDownloads

`func (o *ReposCreateForAuthenticatedUserRequest) HasHasDownloads() bool`

HasHasDownloads returns a boolean if a field has been set.

### GetIsTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *ReposCreateForAuthenticatedUserRequest) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *ReposCreateForAuthenticatedUserRequest) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


