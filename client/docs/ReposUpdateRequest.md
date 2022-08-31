# ReposUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the repository. | [optional] 
**Description** | Pointer to **string** | A short description of the repository. | [optional] 
**Homepage** | Pointer to **string** | A URL with more information about the repository. | [optional] 
**Private** | Pointer to **bool** | Either &#x60;true&#x60; to make the repository private or &#x60;false&#x60; to make it public. Default: &#x60;false&#x60;.   **Note**: You will get a &#x60;422&#x60; error if the organization restricts [changing repository visibility](https://docs.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private. | [optional] [default to false]
**Visibility** | Pointer to **string** | Can be &#x60;public&#x60; or &#x60;private&#x60;. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, &#x60;visibility&#x60; can also be &#x60;internal&#x60;.\&quot; | [optional] 
**SecurityAndAnalysis** | Pointer to [**NullableReposUpdateRequestSecurityAndAnalysis**](ReposUpdateRequestSecurityAndAnalysis.md) |  | [optional] 
**HasIssues** | Pointer to **bool** | Either &#x60;true&#x60; to enable issues for this repository or &#x60;false&#x60; to disable them. | [optional] [default to true]
**HasProjects** | Pointer to **bool** | Either &#x60;true&#x60; to enable projects for this repository or &#x60;false&#x60; to disable them. **Note:** If you&#39;re creating a repository in an organization that has disabled repository projects, the default is &#x60;false&#x60;, and if you pass &#x60;true&#x60;, the API returns an error. | [optional] [default to true]
**HasWiki** | Pointer to **bool** | Either &#x60;true&#x60; to enable the wiki for this repository or &#x60;false&#x60; to disable it. | [optional] [default to true]
**IsTemplate** | Pointer to **bool** | Either &#x60;true&#x60; to make this repo available as a template repository or &#x60;false&#x60; to prevent it. | [optional] [default to false]
**DefaultBranch** | Pointer to **string** | Updates the default branch for this repository. | [optional] 
**AllowSquashMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow squash-merging pull requests, or &#x60;false&#x60; to prevent squash-merging. | [optional] [default to true]
**AllowMergeCommit** | Pointer to **bool** | Either &#x60;true&#x60; to allow merging pull requests with a merge commit, or &#x60;false&#x60; to prevent merging pull requests with merge commits. | [optional] [default to true]
**AllowRebaseMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow rebase-merging pull requests, or &#x60;false&#x60; to prevent rebase-merging. | [optional] [default to true]
**AllowAutoMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow auto-merge on pull requests, or &#x60;false&#x60; to disallow auto-merge. | [optional] [default to false]
**DeleteBranchOnMerge** | Pointer to **bool** | Either &#x60;true&#x60; to allow automatically deleting head branches when pull requests are merged, or &#x60;false&#x60; to prevent automatic deletion. | [optional] [default to false]
**AllowUpdateBranch** | Pointer to **bool** | Either &#x60;true&#x60; to always allow a pull request head branch that is behind its base branch to be updated even if it is not required to be up to date before merging, or false otherwise. | [optional] [default to false]
**UseSquashPrTitleAsDefault** | Pointer to **bool** | Either &#x60;true&#x60; to allow squash-merge commits to use pull request title, or &#x60;false&#x60; to use commit message. | [optional] [default to false]
**SquashMergeCommitTitle** | Pointer to **string** | The default value for a squash merge commit title:  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;COMMIT_OR_PR_TITLE&#x60; - default to the commit&#39;s title (if only one commit) or the pull request&#39;s title (when more than one commit). | [optional] 
**SquashMergeCommitMessage** | Pointer to **string** | The default value for a squash merge commit message:  - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;COMMIT_MESSAGES&#x60; - default to the branch&#39;s commit messages. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**MergeCommitTitle** | Pointer to **string** | The default value for a merge commit title.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;MERGE_MESSAGE&#x60; - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name). | [optional] 
**MergeCommitMessage** | Pointer to **string** | The default value for a merge commit message.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**Archived** | Pointer to **bool** | &#x60;true&#x60; to archive this repository. **Note**: You cannot unarchive repositories through the API. | [optional] [default to false]
**AllowForking** | Pointer to **bool** | Either &#x60;true&#x60; to allow private forks, or &#x60;false&#x60; to prevent private forks. | [optional] [default to false]
**WebCommitSignoffRequired** | Pointer to **bool** | Either &#x60;true&#x60; to require contributors to sign off on web-based commits, or &#x60;false&#x60; to not require contributors to sign off on web-based commits. | [optional] [default to false]

## Methods

### NewReposUpdateRequest

`func NewReposUpdateRequest() *ReposUpdateRequest`

NewReposUpdateRequest instantiates a new ReposUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateRequestWithDefaults

`func NewReposUpdateRequestWithDefaults() *ReposUpdateRequest`

NewReposUpdateRequestWithDefaults instantiates a new ReposUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReposUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ReposUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomepage

`func (o *ReposUpdateRequest) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *ReposUpdateRequest) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *ReposUpdateRequest) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *ReposUpdateRequest) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetPrivate

`func (o *ReposUpdateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ReposUpdateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ReposUpdateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ReposUpdateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetVisibility

`func (o *ReposUpdateRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ReposUpdateRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ReposUpdateRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ReposUpdateRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSecurityAndAnalysis

`func (o *ReposUpdateRequest) GetSecurityAndAnalysis() ReposUpdateRequestSecurityAndAnalysis`

GetSecurityAndAnalysis returns the SecurityAndAnalysis field if non-nil, zero value otherwise.

### GetSecurityAndAnalysisOk

`func (o *ReposUpdateRequest) GetSecurityAndAnalysisOk() (*ReposUpdateRequestSecurityAndAnalysis, bool)`

GetSecurityAndAnalysisOk returns a tuple with the SecurityAndAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAndAnalysis

`func (o *ReposUpdateRequest) SetSecurityAndAnalysis(v ReposUpdateRequestSecurityAndAnalysis)`

SetSecurityAndAnalysis sets SecurityAndAnalysis field to given value.

### HasSecurityAndAnalysis

`func (o *ReposUpdateRequest) HasSecurityAndAnalysis() bool`

HasSecurityAndAnalysis returns a boolean if a field has been set.

### SetSecurityAndAnalysisNil

`func (o *ReposUpdateRequest) SetSecurityAndAnalysisNil(b bool)`

 SetSecurityAndAnalysisNil sets the value for SecurityAndAnalysis to be an explicit nil

### UnsetSecurityAndAnalysis
`func (o *ReposUpdateRequest) UnsetSecurityAndAnalysis()`

UnsetSecurityAndAnalysis ensures that no value is present for SecurityAndAnalysis, not even an explicit nil
### GetHasIssues

`func (o *ReposUpdateRequest) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *ReposUpdateRequest) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *ReposUpdateRequest) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.

### HasHasIssues

`func (o *ReposUpdateRequest) HasHasIssues() bool`

HasHasIssues returns a boolean if a field has been set.

### GetHasProjects

`func (o *ReposUpdateRequest) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *ReposUpdateRequest) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *ReposUpdateRequest) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.

### HasHasProjects

`func (o *ReposUpdateRequest) HasHasProjects() bool`

HasHasProjects returns a boolean if a field has been set.

### GetHasWiki

`func (o *ReposUpdateRequest) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *ReposUpdateRequest) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *ReposUpdateRequest) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.

### HasHasWiki

`func (o *ReposUpdateRequest) HasHasWiki() bool`

HasHasWiki returns a boolean if a field has been set.

### GetIsTemplate

`func (o *ReposUpdateRequest) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *ReposUpdateRequest) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *ReposUpdateRequest) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *ReposUpdateRequest) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ReposUpdateRequest) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ReposUpdateRequest) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ReposUpdateRequest) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ReposUpdateRequest) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *ReposUpdateRequest) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *ReposUpdateRequest) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *ReposUpdateRequest) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *ReposUpdateRequest) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *ReposUpdateRequest) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *ReposUpdateRequest) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *ReposUpdateRequest) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *ReposUpdateRequest) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *ReposUpdateRequest) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *ReposUpdateRequest) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *ReposUpdateRequest) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *ReposUpdateRequest) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *ReposUpdateRequest) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *ReposUpdateRequest) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *ReposUpdateRequest) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *ReposUpdateRequest) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *ReposUpdateRequest) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *ReposUpdateRequest) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *ReposUpdateRequest) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *ReposUpdateRequest) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetAllowUpdateBranch

`func (o *ReposUpdateRequest) GetAllowUpdateBranch() bool`

GetAllowUpdateBranch returns the AllowUpdateBranch field if non-nil, zero value otherwise.

### GetAllowUpdateBranchOk

`func (o *ReposUpdateRequest) GetAllowUpdateBranchOk() (*bool, bool)`

GetAllowUpdateBranchOk returns a tuple with the AllowUpdateBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdateBranch

`func (o *ReposUpdateRequest) SetAllowUpdateBranch(v bool)`

SetAllowUpdateBranch sets AllowUpdateBranch field to given value.

### HasAllowUpdateBranch

`func (o *ReposUpdateRequest) HasAllowUpdateBranch() bool`

HasAllowUpdateBranch returns a boolean if a field has been set.

### GetUseSquashPrTitleAsDefault

`func (o *ReposUpdateRequest) GetUseSquashPrTitleAsDefault() bool`

GetUseSquashPrTitleAsDefault returns the UseSquashPrTitleAsDefault field if non-nil, zero value otherwise.

### GetUseSquashPrTitleAsDefaultOk

`func (o *ReposUpdateRequest) GetUseSquashPrTitleAsDefaultOk() (*bool, bool)`

GetUseSquashPrTitleAsDefaultOk returns a tuple with the UseSquashPrTitleAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSquashPrTitleAsDefault

`func (o *ReposUpdateRequest) SetUseSquashPrTitleAsDefault(v bool)`

SetUseSquashPrTitleAsDefault sets UseSquashPrTitleAsDefault field to given value.

### HasUseSquashPrTitleAsDefault

`func (o *ReposUpdateRequest) HasUseSquashPrTitleAsDefault() bool`

HasUseSquashPrTitleAsDefault returns a boolean if a field has been set.

### GetSquashMergeCommitTitle

`func (o *ReposUpdateRequest) GetSquashMergeCommitTitle() string`

GetSquashMergeCommitTitle returns the SquashMergeCommitTitle field if non-nil, zero value otherwise.

### GetSquashMergeCommitTitleOk

`func (o *ReposUpdateRequest) GetSquashMergeCommitTitleOk() (*string, bool)`

GetSquashMergeCommitTitleOk returns a tuple with the SquashMergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitTitle

`func (o *ReposUpdateRequest) SetSquashMergeCommitTitle(v string)`

SetSquashMergeCommitTitle sets SquashMergeCommitTitle field to given value.

### HasSquashMergeCommitTitle

`func (o *ReposUpdateRequest) HasSquashMergeCommitTitle() bool`

HasSquashMergeCommitTitle returns a boolean if a field has been set.

### GetSquashMergeCommitMessage

`func (o *ReposUpdateRequest) GetSquashMergeCommitMessage() string`

GetSquashMergeCommitMessage returns the SquashMergeCommitMessage field if non-nil, zero value otherwise.

### GetSquashMergeCommitMessageOk

`func (o *ReposUpdateRequest) GetSquashMergeCommitMessageOk() (*string, bool)`

GetSquashMergeCommitMessageOk returns a tuple with the SquashMergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitMessage

`func (o *ReposUpdateRequest) SetSquashMergeCommitMessage(v string)`

SetSquashMergeCommitMessage sets SquashMergeCommitMessage field to given value.

### HasSquashMergeCommitMessage

`func (o *ReposUpdateRequest) HasSquashMergeCommitMessage() bool`

HasSquashMergeCommitMessage returns a boolean if a field has been set.

### GetMergeCommitTitle

`func (o *ReposUpdateRequest) GetMergeCommitTitle() string`

GetMergeCommitTitle returns the MergeCommitTitle field if non-nil, zero value otherwise.

### GetMergeCommitTitleOk

`func (o *ReposUpdateRequest) GetMergeCommitTitleOk() (*string, bool)`

GetMergeCommitTitleOk returns a tuple with the MergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitTitle

`func (o *ReposUpdateRequest) SetMergeCommitTitle(v string)`

SetMergeCommitTitle sets MergeCommitTitle field to given value.

### HasMergeCommitTitle

`func (o *ReposUpdateRequest) HasMergeCommitTitle() bool`

HasMergeCommitTitle returns a boolean if a field has been set.

### GetMergeCommitMessage

`func (o *ReposUpdateRequest) GetMergeCommitMessage() string`

GetMergeCommitMessage returns the MergeCommitMessage field if non-nil, zero value otherwise.

### GetMergeCommitMessageOk

`func (o *ReposUpdateRequest) GetMergeCommitMessageOk() (*string, bool)`

GetMergeCommitMessageOk returns a tuple with the MergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitMessage

`func (o *ReposUpdateRequest) SetMergeCommitMessage(v string)`

SetMergeCommitMessage sets MergeCommitMessage field to given value.

### HasMergeCommitMessage

`func (o *ReposUpdateRequest) HasMergeCommitMessage() bool`

HasMergeCommitMessage returns a boolean if a field has been set.

### GetArchived

`func (o *ReposUpdateRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ReposUpdateRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ReposUpdateRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ReposUpdateRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetAllowForking

`func (o *ReposUpdateRequest) GetAllowForking() bool`

GetAllowForking returns the AllowForking field if non-nil, zero value otherwise.

### GetAllowForkingOk

`func (o *ReposUpdateRequest) GetAllowForkingOk() (*bool, bool)`

GetAllowForkingOk returns a tuple with the AllowForking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForking

`func (o *ReposUpdateRequest) SetAllowForking(v bool)`

SetAllowForking sets AllowForking field to given value.

### HasAllowForking

`func (o *ReposUpdateRequest) HasAllowForking() bool`

HasAllowForking returns a boolean if a field has been set.

### GetWebCommitSignoffRequired

`func (o *ReposUpdateRequest) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *ReposUpdateRequest) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *ReposUpdateRequest) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *ReposUpdateRequest) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


