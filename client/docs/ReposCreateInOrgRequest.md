# ReposCreateInOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the repository. | 
**Description** | Pointer to **string** | A short description of the repository. | [optional] 
**Homepage** | Pointer to **string** | A URL with more information about the repository. | [optional] 
**Private** | Pointer to **bool** | Whether the repository is private. | [optional] [default to false]
**Visibility** | Pointer to **string** | Can be &#x60;public&#x60; or &#x60;private&#x60;. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, &#x60;visibility&#x60; can also be &#x60;internal&#x60;. Note: For GitHub Enterprise Server and GitHub AE, this endpoint will only list repositories available to all users on the enterprise. For more information, see \&quot;[Creating an internal repository](https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/about-repository-visibility#about-internal-repositories)\&quot; in the GitHub Help documentation. | [optional] 
**HasIssues** | Pointer to **bool** | Either &#x60;true&#x60; to enable issues for this repository or &#x60;false&#x60; to disable them. | [optional] [default to true]
**HasProjects** | Pointer to **bool** | Either &#x60;true&#x60; to enable projects for this repository or &#x60;false&#x60; to disable them. **Note:** If you&#39;re creating a repository in an organization that has disabled repository projects, the default is &#x60;false&#x60;, and if you pass &#x60;true&#x60;, the API returns an error. | [optional] [default to true]
**HasWiki** | Pointer to **bool** | Either &#x60;true&#x60; to enable the wiki for this repository or &#x60;false&#x60; to disable it. | [optional] [default to true]
**IsTemplate** | Pointer to **bool** | Either &#x60;true&#x60; to make this repo available as a template repository or &#x60;false&#x60; to prevent it. | [optional] [default to false]
**TeamId** | Pointer to **int32** | The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization. | [optional] 
**AutoInit** | Pointer to **bool** | Pass &#x60;true&#x60; to create an initial commit with empty README. | [optional] [default to false]
**GitignoreTemplate** | Pointer to **string** | Desired language or platform [.gitignore template](https://github.com/github/gitignore) to apply. Use the name of the template without the extension. For example, \&quot;Haskell\&quot;. | [optional] 
**LicenseTemplate** | Pointer to **string** | Choose an [open source license template](https://choosealicense.com/) that best suits your needs, and then use the [license keyword](https://docs.github.com/articles/licensing-a-repository/#searching-github-by-license-type) as the &#x60;license_template&#x60; string. For example, \&quot;mit\&quot; or \&quot;mpl-2.0\&quot;. | [optional] 
**AllowSquashMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow squash-merging pull requests, or &#x60;false&#x60; to prevent squash-merging. | [optional] [default to true]
**AllowMergeCommit** | Pointer to **bool** | Either &#x60;true&#x60; to allow merging pull requests with a merge commit, or &#x60;false&#x60; to prevent merging pull requests with merge commits. | [optional] [default to true]
**AllowRebaseMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow rebase-merging pull requests, or &#x60;false&#x60; to prevent rebase-merging. | [optional] [default to true]
**AllowAutoMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow auto-merge on pull requests, or &#x60;false&#x60; to disallow auto-merge. | [optional] [default to false]
**DeleteBranchOnMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow automatically deleting head branches when pull requests are merged, or &#x60;false&#x60; to prevent automatic deletion. | [optional] [default to false]
**UseSquashPrTitleAsDefault** | Pointer to **bool** | Either &#x60;true&#x60; to allow squash-merge commits to use pull request title, or &#x60;false&#x60; to use commit message. | [optional] [default to false]
**SquashMergeCommitTitle** | Pointer to **string** | The default value for a squash merge commit title:  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;COMMIT_OR_PR_TITLE&#x60; - default to the commit&#39;s title (if only one commit) or the pull request&#39;s title (when more than one commit). | [optional] 
**SquashMergeCommitMessage** | Pointer to **string** | The default value for a squash merge commit message:  - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;COMMIT_MESSAGES&#x60; - default to the branch&#39;s commit messages. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**MergeCommitTitle** | Pointer to **string** | The default value for a merge commit title.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;MERGE_MESSAGE&#x60; - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name). | [optional] 
**MergeCommitMessage** | Pointer to **string** | The default value for a merge commit message.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 

## Methods

### NewReposCreateInOrgRequest

`func NewReposCreateInOrgRequest(name string, ) *ReposCreateInOrgRequest`

NewReposCreateInOrgRequest instantiates a new ReposCreateInOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateInOrgRequestWithDefaults

`func NewReposCreateInOrgRequestWithDefaults() *ReposCreateInOrgRequest`

NewReposCreateInOrgRequestWithDefaults instantiates a new ReposCreateInOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReposCreateInOrgRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateInOrgRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateInOrgRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReposCreateInOrgRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposCreateInOrgRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposCreateInOrgRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposCreateInOrgRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomepage

`func (o *ReposCreateInOrgRequest) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *ReposCreateInOrgRequest) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *ReposCreateInOrgRequest) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *ReposCreateInOrgRequest) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetPrivate

`func (o *ReposCreateInOrgRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ReposCreateInOrgRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ReposCreateInOrgRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ReposCreateInOrgRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetVisibility

`func (o *ReposCreateInOrgRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ReposCreateInOrgRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ReposCreateInOrgRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ReposCreateInOrgRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetHasIssues

`func (o *ReposCreateInOrgRequest) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *ReposCreateInOrgRequest) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *ReposCreateInOrgRequest) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.

### HasHasIssues

`func (o *ReposCreateInOrgRequest) HasHasIssues() bool`

HasHasIssues returns a boolean if a field has been set.

### GetHasProjects

`func (o *ReposCreateInOrgRequest) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *ReposCreateInOrgRequest) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *ReposCreateInOrgRequest) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.

### HasHasProjects

`func (o *ReposCreateInOrgRequest) HasHasProjects() bool`

HasHasProjects returns a boolean if a field has been set.

### GetHasWiki

`func (o *ReposCreateInOrgRequest) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *ReposCreateInOrgRequest) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *ReposCreateInOrgRequest) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.

### HasHasWiki

`func (o *ReposCreateInOrgRequest) HasHasWiki() bool`

HasHasWiki returns a boolean if a field has been set.

### GetIsTemplate

`func (o *ReposCreateInOrgRequest) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *ReposCreateInOrgRequest) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *ReposCreateInOrgRequest) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *ReposCreateInOrgRequest) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetTeamId

`func (o *ReposCreateInOrgRequest) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ReposCreateInOrgRequest) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ReposCreateInOrgRequest) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ReposCreateInOrgRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetAutoInit

`func (o *ReposCreateInOrgRequest) GetAutoInit() bool`

GetAutoInit returns the AutoInit field if non-nil, zero value otherwise.

### GetAutoInitOk

`func (o *ReposCreateInOrgRequest) GetAutoInitOk() (*bool, bool)`

GetAutoInitOk returns a tuple with the AutoInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInit

`func (o *ReposCreateInOrgRequest) SetAutoInit(v bool)`

SetAutoInit sets AutoInit field to given value.

### HasAutoInit

`func (o *ReposCreateInOrgRequest) HasAutoInit() bool`

HasAutoInit returns a boolean if a field has been set.

### GetGitignoreTemplate

`func (o *ReposCreateInOrgRequest) GetGitignoreTemplate() string`

GetGitignoreTemplate returns the GitignoreTemplate field if non-nil, zero value otherwise.

### GetGitignoreTemplateOk

`func (o *ReposCreateInOrgRequest) GetGitignoreTemplateOk() (*string, bool)`

GetGitignoreTemplateOk returns a tuple with the GitignoreTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitignoreTemplate

`func (o *ReposCreateInOrgRequest) SetGitignoreTemplate(v string)`

SetGitignoreTemplate sets GitignoreTemplate field to given value.

### HasGitignoreTemplate

`func (o *ReposCreateInOrgRequest) HasGitignoreTemplate() bool`

HasGitignoreTemplate returns a boolean if a field has been set.

### GetLicenseTemplate

`func (o *ReposCreateInOrgRequest) GetLicenseTemplate() string`

GetLicenseTemplate returns the LicenseTemplate field if non-nil, zero value otherwise.

### GetLicenseTemplateOk

`func (o *ReposCreateInOrgRequest) GetLicenseTemplateOk() (*string, bool)`

GetLicenseTemplateOk returns a tuple with the LicenseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTemplate

`func (o *ReposCreateInOrgRequest) SetLicenseTemplate(v string)`

SetLicenseTemplate sets LicenseTemplate field to given value.

### HasLicenseTemplate

`func (o *ReposCreateInOrgRequest) HasLicenseTemplate() bool`

HasLicenseTemplate returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *ReposCreateInOrgRequest) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *ReposCreateInOrgRequest) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *ReposCreateInOrgRequest) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *ReposCreateInOrgRequest) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *ReposCreateInOrgRequest) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *ReposCreateInOrgRequest) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *ReposCreateInOrgRequest) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *ReposCreateInOrgRequest) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *ReposCreateInOrgRequest) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *ReposCreateInOrgRequest) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *ReposCreateInOrgRequest) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *ReposCreateInOrgRequest) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *ReposCreateInOrgRequest) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *ReposCreateInOrgRequest) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *ReposCreateInOrgRequest) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *ReposCreateInOrgRequest) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *ReposCreateInOrgRequest) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *ReposCreateInOrgRequest) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *ReposCreateInOrgRequest) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *ReposCreateInOrgRequest) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetUseSquashPrTitleAsDefault

`func (o *ReposCreateInOrgRequest) GetUseSquashPrTitleAsDefault() bool`

GetUseSquashPrTitleAsDefault returns the UseSquashPrTitleAsDefault field if non-nil, zero value otherwise.

### GetUseSquashPrTitleAsDefaultOk

`func (o *ReposCreateInOrgRequest) GetUseSquashPrTitleAsDefaultOk() (*bool, bool)`

GetUseSquashPrTitleAsDefaultOk returns a tuple with the UseSquashPrTitleAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSquashPrTitleAsDefault

`func (o *ReposCreateInOrgRequest) SetUseSquashPrTitleAsDefault(v bool)`

SetUseSquashPrTitleAsDefault sets UseSquashPrTitleAsDefault field to given value.

### HasUseSquashPrTitleAsDefault

`func (o *ReposCreateInOrgRequest) HasUseSquashPrTitleAsDefault() bool`

HasUseSquashPrTitleAsDefault returns a boolean if a field has been set.

### GetSquashMergeCommitTitle

`func (o *ReposCreateInOrgRequest) GetSquashMergeCommitTitle() string`

GetSquashMergeCommitTitle returns the SquashMergeCommitTitle field if non-nil, zero value otherwise.

### GetSquashMergeCommitTitleOk

`func (o *ReposCreateInOrgRequest) GetSquashMergeCommitTitleOk() (*string, bool)`

GetSquashMergeCommitTitleOk returns a tuple with the SquashMergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitTitle

`func (o *ReposCreateInOrgRequest) SetSquashMergeCommitTitle(v string)`

SetSquashMergeCommitTitle sets SquashMergeCommitTitle field to given value.

### HasSquashMergeCommitTitle

`func (o *ReposCreateInOrgRequest) HasSquashMergeCommitTitle() bool`

HasSquashMergeCommitTitle returns a boolean if a field has been set.

### GetSquashMergeCommitMessage

`func (o *ReposCreateInOrgRequest) GetSquashMergeCommitMessage() string`

GetSquashMergeCommitMessage returns the SquashMergeCommitMessage field if non-nil, zero value otherwise.

### GetSquashMergeCommitMessageOk

`func (o *ReposCreateInOrgRequest) GetSquashMergeCommitMessageOk() (*string, bool)`

GetSquashMergeCommitMessageOk returns a tuple with the SquashMergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitMessage

`func (o *ReposCreateInOrgRequest) SetSquashMergeCommitMessage(v string)`

SetSquashMergeCommitMessage sets SquashMergeCommitMessage field to given value.

### HasSquashMergeCommitMessage

`func (o *ReposCreateInOrgRequest) HasSquashMergeCommitMessage() bool`

HasSquashMergeCommitMessage returns a boolean if a field has been set.

### GetMergeCommitTitle

`func (o *ReposCreateInOrgRequest) GetMergeCommitTitle() string`

GetMergeCommitTitle returns the MergeCommitTitle field if non-nil, zero value otherwise.

### GetMergeCommitTitleOk

`func (o *ReposCreateInOrgRequest) GetMergeCommitTitleOk() (*string, bool)`

GetMergeCommitTitleOk returns a tuple with the MergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitTitle

`func (o *ReposCreateInOrgRequest) SetMergeCommitTitle(v string)`

SetMergeCommitTitle sets MergeCommitTitle field to given value.

### HasMergeCommitTitle

`func (o *ReposCreateInOrgRequest) HasMergeCommitTitle() bool`

HasMergeCommitTitle returns a boolean if a field has been set.

### GetMergeCommitMessage

`func (o *ReposCreateInOrgRequest) GetMergeCommitMessage() string`

GetMergeCommitMessage returns the MergeCommitMessage field if non-nil, zero value otherwise.

### GetMergeCommitMessageOk

`func (o *ReposCreateInOrgRequest) GetMergeCommitMessageOk() (*string, bool)`

GetMergeCommitMessageOk returns a tuple with the MergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitMessage

`func (o *ReposCreateInOrgRequest) SetMergeCommitMessage(v string)`

SetMergeCommitMessage sets MergeCommitMessage field to given value.

### HasMergeCommitMessage

`func (o *ReposCreateInOrgRequest) HasMergeCommitMessage() bool`

HasMergeCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


