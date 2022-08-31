# RepositoryTemplateRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**RepositoryTemplateRepositoryOwner**](RepositoryTemplateRepositoryOwner.md) |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Fork** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ArchiveUrl** | Pointer to **string** |  | [optional] 
**AssigneesUrl** | Pointer to **string** |  | [optional] 
**BlobsUrl** | Pointer to **string** |  | [optional] 
**BranchesUrl** | Pointer to **string** |  | [optional] 
**CollaboratorsUrl** | Pointer to **string** |  | [optional] 
**CommentsUrl** | Pointer to **string** |  | [optional] 
**CommitsUrl** | Pointer to **string** |  | [optional] 
**CompareUrl** | Pointer to **string** |  | [optional] 
**ContentsUrl** | Pointer to **string** |  | [optional] 
**ContributorsUrl** | Pointer to **string** |  | [optional] 
**DeploymentsUrl** | Pointer to **string** |  | [optional] 
**DownloadsUrl** | Pointer to **string** |  | [optional] 
**EventsUrl** | Pointer to **string** |  | [optional] 
**ForksUrl** | Pointer to **string** |  | [optional] 
**GitCommitsUrl** | Pointer to **string** |  | [optional] 
**GitRefsUrl** | Pointer to **string** |  | [optional] 
**GitTagsUrl** | Pointer to **string** |  | [optional] 
**GitUrl** | Pointer to **string** |  | [optional] 
**IssueCommentUrl** | Pointer to **string** |  | [optional] 
**IssueEventsUrl** | Pointer to **string** |  | [optional] 
**IssuesUrl** | Pointer to **string** |  | [optional] 
**KeysUrl** | Pointer to **string** |  | [optional] 
**LabelsUrl** | Pointer to **string** |  | [optional] 
**LanguagesUrl** | Pointer to **string** |  | [optional] 
**MergesUrl** | Pointer to **string** |  | [optional] 
**MilestonesUrl** | Pointer to **string** |  | [optional] 
**NotificationsUrl** | Pointer to **string** |  | [optional] 
**PullsUrl** | Pointer to **string** |  | [optional] 
**ReleasesUrl** | Pointer to **string** |  | [optional] 
**SshUrl** | Pointer to **string** |  | [optional] 
**StargazersUrl** | Pointer to **string** |  | [optional] 
**StatusesUrl** | Pointer to **string** |  | [optional] 
**SubscribersUrl** | Pointer to **string** |  | [optional] 
**SubscriptionUrl** | Pointer to **string** |  | [optional] 
**TagsUrl** | Pointer to **string** |  | [optional] 
**TeamsUrl** | Pointer to **string** |  | [optional] 
**TreesUrl** | Pointer to **string** |  | [optional] 
**CloneUrl** | Pointer to **string** |  | [optional] 
**MirrorUrl** | Pointer to **string** |  | [optional] 
**HooksUrl** | Pointer to **string** |  | [optional] 
**SvnUrl** | Pointer to **string** |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**ForksCount** | Pointer to **int32** |  | [optional] 
**StargazersCount** | Pointer to **int32** |  | [optional] 
**WatchersCount** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**OpenIssuesCount** | Pointer to **int32** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**HasIssues** | Pointer to **bool** |  | [optional] 
**HasProjects** | Pointer to **bool** |  | [optional] 
**HasWiki** | Pointer to **bool** |  | [optional] 
**HasPages** | Pointer to **bool** |  | [optional] 
**HasDownloads** | Pointer to **bool** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**PushedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**RepositoryTemplateRepositoryPermissions**](RepositoryTemplateRepositoryPermissions.md) |  | [optional] 
**AllowRebaseMerge** | Pointer to **bool** |  | [optional] 
**TempCloneToken** | Pointer to **string** |  | [optional] 
**AllowSquashMerge** | Pointer to **bool** |  | [optional] 
**AllowAutoMerge** | Pointer to **bool** |  | [optional] 
**DeleteBranchOnMerge** | Pointer to **bool** |  | [optional] 
**AllowUpdateBranch** | Pointer to **bool** |  | [optional] 
**UseSquashPrTitleAsDefault** | Pointer to **bool** |  | [optional] 
**SquashMergeCommitTitle** | Pointer to **string** | The default value for a squash merge commit title:  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;COMMIT_OR_PR_TITLE&#x60; - default to the commit&#39;s title (if only one commit) or the pull request&#39;s title (when more than one commit). | [optional] 
**SquashMergeCommitMessage** | Pointer to **string** | The default value for a squash merge commit message:  - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;COMMIT_MESSAGES&#x60; - default to the branch&#39;s commit messages. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**MergeCommitTitle** | Pointer to **string** | The default value for a merge commit title.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;MERGE_MESSAGE&#x60; - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name). | [optional] 
**MergeCommitMessage** | Pointer to **string** | The default value for a merge commit message.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**AllowMergeCommit** | Pointer to **bool** |  | [optional] 
**SubscribersCount** | Pointer to **int32** |  | [optional] 
**NetworkCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewRepositoryTemplateRepository

`func NewRepositoryTemplateRepository() *RepositoryTemplateRepository`

NewRepositoryTemplateRepository instantiates a new RepositoryTemplateRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTemplateRepositoryWithDefaults

`func NewRepositoryTemplateRepositoryWithDefaults() *RepositoryTemplateRepository`

NewRepositoryTemplateRepositoryWithDefaults instantiates a new RepositoryTemplateRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepositoryTemplateRepository) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepositoryTemplateRepository) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepositoryTemplateRepository) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RepositoryTemplateRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeId

`func (o *RepositoryTemplateRepository) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RepositoryTemplateRepository) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RepositoryTemplateRepository) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *RepositoryTemplateRepository) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetName

`func (o *RepositoryTemplateRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryTemplateRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryTemplateRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryTemplateRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFullName

`func (o *RepositoryTemplateRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *RepositoryTemplateRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *RepositoryTemplateRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *RepositoryTemplateRepository) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetOwner

`func (o *RepositoryTemplateRepository) GetOwner() RepositoryTemplateRepositoryOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RepositoryTemplateRepository) GetOwnerOk() (*RepositoryTemplateRepositoryOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RepositoryTemplateRepository) SetOwner(v RepositoryTemplateRepositoryOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RepositoryTemplateRepository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivate

`func (o *RepositoryTemplateRepository) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *RepositoryTemplateRepository) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *RepositoryTemplateRepository) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *RepositoryTemplateRepository) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *RepositoryTemplateRepository) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *RepositoryTemplateRepository) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *RepositoryTemplateRepository) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *RepositoryTemplateRepository) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetDescription

`func (o *RepositoryTemplateRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositoryTemplateRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositoryTemplateRepository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepositoryTemplateRepository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFork

`func (o *RepositoryTemplateRepository) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *RepositoryTemplateRepository) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *RepositoryTemplateRepository) SetFork(v bool)`

SetFork sets Fork field to given value.

### HasFork

`func (o *RepositoryTemplateRepository) HasFork() bool`

HasFork returns a boolean if a field has been set.

### GetUrl

`func (o *RepositoryTemplateRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryTemplateRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryTemplateRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RepositoryTemplateRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetArchiveUrl

`func (o *RepositoryTemplateRepository) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *RepositoryTemplateRepository) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *RepositoryTemplateRepository) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.

### HasArchiveUrl

`func (o *RepositoryTemplateRepository) HasArchiveUrl() bool`

HasArchiveUrl returns a boolean if a field has been set.

### GetAssigneesUrl

`func (o *RepositoryTemplateRepository) GetAssigneesUrl() string`

GetAssigneesUrl returns the AssigneesUrl field if non-nil, zero value otherwise.

### GetAssigneesUrlOk

`func (o *RepositoryTemplateRepository) GetAssigneesUrlOk() (*string, bool)`

GetAssigneesUrlOk returns a tuple with the AssigneesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesUrl

`func (o *RepositoryTemplateRepository) SetAssigneesUrl(v string)`

SetAssigneesUrl sets AssigneesUrl field to given value.

### HasAssigneesUrl

`func (o *RepositoryTemplateRepository) HasAssigneesUrl() bool`

HasAssigneesUrl returns a boolean if a field has been set.

### GetBlobsUrl

`func (o *RepositoryTemplateRepository) GetBlobsUrl() string`

GetBlobsUrl returns the BlobsUrl field if non-nil, zero value otherwise.

### GetBlobsUrlOk

`func (o *RepositoryTemplateRepository) GetBlobsUrlOk() (*string, bool)`

GetBlobsUrlOk returns a tuple with the BlobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsUrl

`func (o *RepositoryTemplateRepository) SetBlobsUrl(v string)`

SetBlobsUrl sets BlobsUrl field to given value.

### HasBlobsUrl

`func (o *RepositoryTemplateRepository) HasBlobsUrl() bool`

HasBlobsUrl returns a boolean if a field has been set.

### GetBranchesUrl

`func (o *RepositoryTemplateRepository) GetBranchesUrl() string`

GetBranchesUrl returns the BranchesUrl field if non-nil, zero value otherwise.

### GetBranchesUrlOk

`func (o *RepositoryTemplateRepository) GetBranchesUrlOk() (*string, bool)`

GetBranchesUrlOk returns a tuple with the BranchesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesUrl

`func (o *RepositoryTemplateRepository) SetBranchesUrl(v string)`

SetBranchesUrl sets BranchesUrl field to given value.

### HasBranchesUrl

`func (o *RepositoryTemplateRepository) HasBranchesUrl() bool`

HasBranchesUrl returns a boolean if a field has been set.

### GetCollaboratorsUrl

`func (o *RepositoryTemplateRepository) GetCollaboratorsUrl() string`

GetCollaboratorsUrl returns the CollaboratorsUrl field if non-nil, zero value otherwise.

### GetCollaboratorsUrlOk

`func (o *RepositoryTemplateRepository) GetCollaboratorsUrlOk() (*string, bool)`

GetCollaboratorsUrlOk returns a tuple with the CollaboratorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorsUrl

`func (o *RepositoryTemplateRepository) SetCollaboratorsUrl(v string)`

SetCollaboratorsUrl sets CollaboratorsUrl field to given value.

### HasCollaboratorsUrl

`func (o *RepositoryTemplateRepository) HasCollaboratorsUrl() bool`

HasCollaboratorsUrl returns a boolean if a field has been set.

### GetCommentsUrl

`func (o *RepositoryTemplateRepository) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *RepositoryTemplateRepository) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *RepositoryTemplateRepository) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.

### HasCommentsUrl

`func (o *RepositoryTemplateRepository) HasCommentsUrl() bool`

HasCommentsUrl returns a boolean if a field has been set.

### GetCommitsUrl

`func (o *RepositoryTemplateRepository) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *RepositoryTemplateRepository) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *RepositoryTemplateRepository) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.

### HasCommitsUrl

`func (o *RepositoryTemplateRepository) HasCommitsUrl() bool`

HasCommitsUrl returns a boolean if a field has been set.

### GetCompareUrl

`func (o *RepositoryTemplateRepository) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *RepositoryTemplateRepository) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *RepositoryTemplateRepository) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.

### HasCompareUrl

`func (o *RepositoryTemplateRepository) HasCompareUrl() bool`

HasCompareUrl returns a boolean if a field has been set.

### GetContentsUrl

`func (o *RepositoryTemplateRepository) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *RepositoryTemplateRepository) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *RepositoryTemplateRepository) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.

### HasContentsUrl

`func (o *RepositoryTemplateRepository) HasContentsUrl() bool`

HasContentsUrl returns a boolean if a field has been set.

### GetContributorsUrl

`func (o *RepositoryTemplateRepository) GetContributorsUrl() string`

GetContributorsUrl returns the ContributorsUrl field if non-nil, zero value otherwise.

### GetContributorsUrlOk

`func (o *RepositoryTemplateRepository) GetContributorsUrlOk() (*string, bool)`

GetContributorsUrlOk returns a tuple with the ContributorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsUrl

`func (o *RepositoryTemplateRepository) SetContributorsUrl(v string)`

SetContributorsUrl sets ContributorsUrl field to given value.

### HasContributorsUrl

`func (o *RepositoryTemplateRepository) HasContributorsUrl() bool`

HasContributorsUrl returns a boolean if a field has been set.

### GetDeploymentsUrl

`func (o *RepositoryTemplateRepository) GetDeploymentsUrl() string`

GetDeploymentsUrl returns the DeploymentsUrl field if non-nil, zero value otherwise.

### GetDeploymentsUrlOk

`func (o *RepositoryTemplateRepository) GetDeploymentsUrlOk() (*string, bool)`

GetDeploymentsUrlOk returns a tuple with the DeploymentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsUrl

`func (o *RepositoryTemplateRepository) SetDeploymentsUrl(v string)`

SetDeploymentsUrl sets DeploymentsUrl field to given value.

### HasDeploymentsUrl

`func (o *RepositoryTemplateRepository) HasDeploymentsUrl() bool`

HasDeploymentsUrl returns a boolean if a field has been set.

### GetDownloadsUrl

`func (o *RepositoryTemplateRepository) GetDownloadsUrl() string`

GetDownloadsUrl returns the DownloadsUrl field if non-nil, zero value otherwise.

### GetDownloadsUrlOk

`func (o *RepositoryTemplateRepository) GetDownloadsUrlOk() (*string, bool)`

GetDownloadsUrlOk returns a tuple with the DownloadsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsUrl

`func (o *RepositoryTemplateRepository) SetDownloadsUrl(v string)`

SetDownloadsUrl sets DownloadsUrl field to given value.

### HasDownloadsUrl

`func (o *RepositoryTemplateRepository) HasDownloadsUrl() bool`

HasDownloadsUrl returns a boolean if a field has been set.

### GetEventsUrl

`func (o *RepositoryTemplateRepository) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *RepositoryTemplateRepository) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *RepositoryTemplateRepository) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.

### HasEventsUrl

`func (o *RepositoryTemplateRepository) HasEventsUrl() bool`

HasEventsUrl returns a boolean if a field has been set.

### GetForksUrl

`func (o *RepositoryTemplateRepository) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *RepositoryTemplateRepository) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *RepositoryTemplateRepository) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.

### HasForksUrl

`func (o *RepositoryTemplateRepository) HasForksUrl() bool`

HasForksUrl returns a boolean if a field has been set.

### GetGitCommitsUrl

`func (o *RepositoryTemplateRepository) GetGitCommitsUrl() string`

GetGitCommitsUrl returns the GitCommitsUrl field if non-nil, zero value otherwise.

### GetGitCommitsUrlOk

`func (o *RepositoryTemplateRepository) GetGitCommitsUrlOk() (*string, bool)`

GetGitCommitsUrlOk returns a tuple with the GitCommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitsUrl

`func (o *RepositoryTemplateRepository) SetGitCommitsUrl(v string)`

SetGitCommitsUrl sets GitCommitsUrl field to given value.

### HasGitCommitsUrl

`func (o *RepositoryTemplateRepository) HasGitCommitsUrl() bool`

HasGitCommitsUrl returns a boolean if a field has been set.

### GetGitRefsUrl

`func (o *RepositoryTemplateRepository) GetGitRefsUrl() string`

GetGitRefsUrl returns the GitRefsUrl field if non-nil, zero value otherwise.

### GetGitRefsUrlOk

`func (o *RepositoryTemplateRepository) GetGitRefsUrlOk() (*string, bool)`

GetGitRefsUrlOk returns a tuple with the GitRefsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefsUrl

`func (o *RepositoryTemplateRepository) SetGitRefsUrl(v string)`

SetGitRefsUrl sets GitRefsUrl field to given value.

### HasGitRefsUrl

`func (o *RepositoryTemplateRepository) HasGitRefsUrl() bool`

HasGitRefsUrl returns a boolean if a field has been set.

### GetGitTagsUrl

`func (o *RepositoryTemplateRepository) GetGitTagsUrl() string`

GetGitTagsUrl returns the GitTagsUrl field if non-nil, zero value otherwise.

### GetGitTagsUrlOk

`func (o *RepositoryTemplateRepository) GetGitTagsUrlOk() (*string, bool)`

GetGitTagsUrlOk returns a tuple with the GitTagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagsUrl

`func (o *RepositoryTemplateRepository) SetGitTagsUrl(v string)`

SetGitTagsUrl sets GitTagsUrl field to given value.

### HasGitTagsUrl

`func (o *RepositoryTemplateRepository) HasGitTagsUrl() bool`

HasGitTagsUrl returns a boolean if a field has been set.

### GetGitUrl

`func (o *RepositoryTemplateRepository) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *RepositoryTemplateRepository) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *RepositoryTemplateRepository) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *RepositoryTemplateRepository) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetIssueCommentUrl

`func (o *RepositoryTemplateRepository) GetIssueCommentUrl() string`

GetIssueCommentUrl returns the IssueCommentUrl field if non-nil, zero value otherwise.

### GetIssueCommentUrlOk

`func (o *RepositoryTemplateRepository) GetIssueCommentUrlOk() (*string, bool)`

GetIssueCommentUrlOk returns a tuple with the IssueCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCommentUrl

`func (o *RepositoryTemplateRepository) SetIssueCommentUrl(v string)`

SetIssueCommentUrl sets IssueCommentUrl field to given value.

### HasIssueCommentUrl

`func (o *RepositoryTemplateRepository) HasIssueCommentUrl() bool`

HasIssueCommentUrl returns a boolean if a field has been set.

### GetIssueEventsUrl

`func (o *RepositoryTemplateRepository) GetIssueEventsUrl() string`

GetIssueEventsUrl returns the IssueEventsUrl field if non-nil, zero value otherwise.

### GetIssueEventsUrlOk

`func (o *RepositoryTemplateRepository) GetIssueEventsUrlOk() (*string, bool)`

GetIssueEventsUrlOk returns a tuple with the IssueEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEventsUrl

`func (o *RepositoryTemplateRepository) SetIssueEventsUrl(v string)`

SetIssueEventsUrl sets IssueEventsUrl field to given value.

### HasIssueEventsUrl

`func (o *RepositoryTemplateRepository) HasIssueEventsUrl() bool`

HasIssueEventsUrl returns a boolean if a field has been set.

### GetIssuesUrl

`func (o *RepositoryTemplateRepository) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *RepositoryTemplateRepository) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *RepositoryTemplateRepository) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.

### HasIssuesUrl

`func (o *RepositoryTemplateRepository) HasIssuesUrl() bool`

HasIssuesUrl returns a boolean if a field has been set.

### GetKeysUrl

`func (o *RepositoryTemplateRepository) GetKeysUrl() string`

GetKeysUrl returns the KeysUrl field if non-nil, zero value otherwise.

### GetKeysUrlOk

`func (o *RepositoryTemplateRepository) GetKeysUrlOk() (*string, bool)`

GetKeysUrlOk returns a tuple with the KeysUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysUrl

`func (o *RepositoryTemplateRepository) SetKeysUrl(v string)`

SetKeysUrl sets KeysUrl field to given value.

### HasKeysUrl

`func (o *RepositoryTemplateRepository) HasKeysUrl() bool`

HasKeysUrl returns a boolean if a field has been set.

### GetLabelsUrl

`func (o *RepositoryTemplateRepository) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *RepositoryTemplateRepository) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *RepositoryTemplateRepository) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.

### HasLabelsUrl

`func (o *RepositoryTemplateRepository) HasLabelsUrl() bool`

HasLabelsUrl returns a boolean if a field has been set.

### GetLanguagesUrl

`func (o *RepositoryTemplateRepository) GetLanguagesUrl() string`

GetLanguagesUrl returns the LanguagesUrl field if non-nil, zero value otherwise.

### GetLanguagesUrlOk

`func (o *RepositoryTemplateRepository) GetLanguagesUrlOk() (*string, bool)`

GetLanguagesUrlOk returns a tuple with the LanguagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesUrl

`func (o *RepositoryTemplateRepository) SetLanguagesUrl(v string)`

SetLanguagesUrl sets LanguagesUrl field to given value.

### HasLanguagesUrl

`func (o *RepositoryTemplateRepository) HasLanguagesUrl() bool`

HasLanguagesUrl returns a boolean if a field has been set.

### GetMergesUrl

`func (o *RepositoryTemplateRepository) GetMergesUrl() string`

GetMergesUrl returns the MergesUrl field if non-nil, zero value otherwise.

### GetMergesUrlOk

`func (o *RepositoryTemplateRepository) GetMergesUrlOk() (*string, bool)`

GetMergesUrlOk returns a tuple with the MergesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesUrl

`func (o *RepositoryTemplateRepository) SetMergesUrl(v string)`

SetMergesUrl sets MergesUrl field to given value.

### HasMergesUrl

`func (o *RepositoryTemplateRepository) HasMergesUrl() bool`

HasMergesUrl returns a boolean if a field has been set.

### GetMilestonesUrl

`func (o *RepositoryTemplateRepository) GetMilestonesUrl() string`

GetMilestonesUrl returns the MilestonesUrl field if non-nil, zero value otherwise.

### GetMilestonesUrlOk

`func (o *RepositoryTemplateRepository) GetMilestonesUrlOk() (*string, bool)`

GetMilestonesUrlOk returns a tuple with the MilestonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestonesUrl

`func (o *RepositoryTemplateRepository) SetMilestonesUrl(v string)`

SetMilestonesUrl sets MilestonesUrl field to given value.

### HasMilestonesUrl

`func (o *RepositoryTemplateRepository) HasMilestonesUrl() bool`

HasMilestonesUrl returns a boolean if a field has been set.

### GetNotificationsUrl

`func (o *RepositoryTemplateRepository) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *RepositoryTemplateRepository) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *RepositoryTemplateRepository) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.

### HasNotificationsUrl

`func (o *RepositoryTemplateRepository) HasNotificationsUrl() bool`

HasNotificationsUrl returns a boolean if a field has been set.

### GetPullsUrl

`func (o *RepositoryTemplateRepository) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *RepositoryTemplateRepository) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *RepositoryTemplateRepository) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.

### HasPullsUrl

`func (o *RepositoryTemplateRepository) HasPullsUrl() bool`

HasPullsUrl returns a boolean if a field has been set.

### GetReleasesUrl

`func (o *RepositoryTemplateRepository) GetReleasesUrl() string`

GetReleasesUrl returns the ReleasesUrl field if non-nil, zero value otherwise.

### GetReleasesUrlOk

`func (o *RepositoryTemplateRepository) GetReleasesUrlOk() (*string, bool)`

GetReleasesUrlOk returns a tuple with the ReleasesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesUrl

`func (o *RepositoryTemplateRepository) SetReleasesUrl(v string)`

SetReleasesUrl sets ReleasesUrl field to given value.

### HasReleasesUrl

`func (o *RepositoryTemplateRepository) HasReleasesUrl() bool`

HasReleasesUrl returns a boolean if a field has been set.

### GetSshUrl

`func (o *RepositoryTemplateRepository) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *RepositoryTemplateRepository) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *RepositoryTemplateRepository) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.

### HasSshUrl

`func (o *RepositoryTemplateRepository) HasSshUrl() bool`

HasSshUrl returns a boolean if a field has been set.

### GetStargazersUrl

`func (o *RepositoryTemplateRepository) GetStargazersUrl() string`

GetStargazersUrl returns the StargazersUrl field if non-nil, zero value otherwise.

### GetStargazersUrlOk

`func (o *RepositoryTemplateRepository) GetStargazersUrlOk() (*string, bool)`

GetStargazersUrlOk returns a tuple with the StargazersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersUrl

`func (o *RepositoryTemplateRepository) SetStargazersUrl(v string)`

SetStargazersUrl sets StargazersUrl field to given value.

### HasStargazersUrl

`func (o *RepositoryTemplateRepository) HasStargazersUrl() bool`

HasStargazersUrl returns a boolean if a field has been set.

### GetStatusesUrl

`func (o *RepositoryTemplateRepository) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *RepositoryTemplateRepository) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *RepositoryTemplateRepository) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.

### HasStatusesUrl

`func (o *RepositoryTemplateRepository) HasStatusesUrl() bool`

HasStatusesUrl returns a boolean if a field has been set.

### GetSubscribersUrl

`func (o *RepositoryTemplateRepository) GetSubscribersUrl() string`

GetSubscribersUrl returns the SubscribersUrl field if non-nil, zero value otherwise.

### GetSubscribersUrlOk

`func (o *RepositoryTemplateRepository) GetSubscribersUrlOk() (*string, bool)`

GetSubscribersUrlOk returns a tuple with the SubscribersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersUrl

`func (o *RepositoryTemplateRepository) SetSubscribersUrl(v string)`

SetSubscribersUrl sets SubscribersUrl field to given value.

### HasSubscribersUrl

`func (o *RepositoryTemplateRepository) HasSubscribersUrl() bool`

HasSubscribersUrl returns a boolean if a field has been set.

### GetSubscriptionUrl

`func (o *RepositoryTemplateRepository) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *RepositoryTemplateRepository) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *RepositoryTemplateRepository) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.

### HasSubscriptionUrl

`func (o *RepositoryTemplateRepository) HasSubscriptionUrl() bool`

HasSubscriptionUrl returns a boolean if a field has been set.

### GetTagsUrl

`func (o *RepositoryTemplateRepository) GetTagsUrl() string`

GetTagsUrl returns the TagsUrl field if non-nil, zero value otherwise.

### GetTagsUrlOk

`func (o *RepositoryTemplateRepository) GetTagsUrlOk() (*string, bool)`

GetTagsUrlOk returns a tuple with the TagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsUrl

`func (o *RepositoryTemplateRepository) SetTagsUrl(v string)`

SetTagsUrl sets TagsUrl field to given value.

### HasTagsUrl

`func (o *RepositoryTemplateRepository) HasTagsUrl() bool`

HasTagsUrl returns a boolean if a field has been set.

### GetTeamsUrl

`func (o *RepositoryTemplateRepository) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *RepositoryTemplateRepository) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *RepositoryTemplateRepository) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.

### HasTeamsUrl

`func (o *RepositoryTemplateRepository) HasTeamsUrl() bool`

HasTeamsUrl returns a boolean if a field has been set.

### GetTreesUrl

`func (o *RepositoryTemplateRepository) GetTreesUrl() string`

GetTreesUrl returns the TreesUrl field if non-nil, zero value otherwise.

### GetTreesUrlOk

`func (o *RepositoryTemplateRepository) GetTreesUrlOk() (*string, bool)`

GetTreesUrlOk returns a tuple with the TreesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreesUrl

`func (o *RepositoryTemplateRepository) SetTreesUrl(v string)`

SetTreesUrl sets TreesUrl field to given value.

### HasTreesUrl

`func (o *RepositoryTemplateRepository) HasTreesUrl() bool`

HasTreesUrl returns a boolean if a field has been set.

### GetCloneUrl

`func (o *RepositoryTemplateRepository) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *RepositoryTemplateRepository) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *RepositoryTemplateRepository) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *RepositoryTemplateRepository) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetMirrorUrl

`func (o *RepositoryTemplateRepository) GetMirrorUrl() string`

GetMirrorUrl returns the MirrorUrl field if non-nil, zero value otherwise.

### GetMirrorUrlOk

`func (o *RepositoryTemplateRepository) GetMirrorUrlOk() (*string, bool)`

GetMirrorUrlOk returns a tuple with the MirrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorUrl

`func (o *RepositoryTemplateRepository) SetMirrorUrl(v string)`

SetMirrorUrl sets MirrorUrl field to given value.

### HasMirrorUrl

`func (o *RepositoryTemplateRepository) HasMirrorUrl() bool`

HasMirrorUrl returns a boolean if a field has been set.

### GetHooksUrl

`func (o *RepositoryTemplateRepository) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *RepositoryTemplateRepository) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *RepositoryTemplateRepository) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.

### HasHooksUrl

`func (o *RepositoryTemplateRepository) HasHooksUrl() bool`

HasHooksUrl returns a boolean if a field has been set.

### GetSvnUrl

`func (o *RepositoryTemplateRepository) GetSvnUrl() string`

GetSvnUrl returns the SvnUrl field if non-nil, zero value otherwise.

### GetSvnUrlOk

`func (o *RepositoryTemplateRepository) GetSvnUrlOk() (*string, bool)`

GetSvnUrlOk returns a tuple with the SvnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnUrl

`func (o *RepositoryTemplateRepository) SetSvnUrl(v string)`

SetSvnUrl sets SvnUrl field to given value.

### HasSvnUrl

`func (o *RepositoryTemplateRepository) HasSvnUrl() bool`

HasSvnUrl returns a boolean if a field has been set.

### GetHomepage

`func (o *RepositoryTemplateRepository) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *RepositoryTemplateRepository) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *RepositoryTemplateRepository) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *RepositoryTemplateRepository) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetLanguage

`func (o *RepositoryTemplateRepository) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RepositoryTemplateRepository) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RepositoryTemplateRepository) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *RepositoryTemplateRepository) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetForksCount

`func (o *RepositoryTemplateRepository) GetForksCount() int32`

GetForksCount returns the ForksCount field if non-nil, zero value otherwise.

### GetForksCountOk

`func (o *RepositoryTemplateRepository) GetForksCountOk() (*int32, bool)`

GetForksCountOk returns a tuple with the ForksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksCount

`func (o *RepositoryTemplateRepository) SetForksCount(v int32)`

SetForksCount sets ForksCount field to given value.

### HasForksCount

`func (o *RepositoryTemplateRepository) HasForksCount() bool`

HasForksCount returns a boolean if a field has been set.

### GetStargazersCount

`func (o *RepositoryTemplateRepository) GetStargazersCount() int32`

GetStargazersCount returns the StargazersCount field if non-nil, zero value otherwise.

### GetStargazersCountOk

`func (o *RepositoryTemplateRepository) GetStargazersCountOk() (*int32, bool)`

GetStargazersCountOk returns a tuple with the StargazersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersCount

`func (o *RepositoryTemplateRepository) SetStargazersCount(v int32)`

SetStargazersCount sets StargazersCount field to given value.

### HasStargazersCount

`func (o *RepositoryTemplateRepository) HasStargazersCount() bool`

HasStargazersCount returns a boolean if a field has been set.

### GetWatchersCount

`func (o *RepositoryTemplateRepository) GetWatchersCount() int32`

GetWatchersCount returns the WatchersCount field if non-nil, zero value otherwise.

### GetWatchersCountOk

`func (o *RepositoryTemplateRepository) GetWatchersCountOk() (*int32, bool)`

GetWatchersCountOk returns a tuple with the WatchersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchersCount

`func (o *RepositoryTemplateRepository) SetWatchersCount(v int32)`

SetWatchersCount sets WatchersCount field to given value.

### HasWatchersCount

`func (o *RepositoryTemplateRepository) HasWatchersCount() bool`

HasWatchersCount returns a boolean if a field has been set.

### GetSize

`func (o *RepositoryTemplateRepository) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RepositoryTemplateRepository) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RepositoryTemplateRepository) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RepositoryTemplateRepository) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *RepositoryTemplateRepository) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *RepositoryTemplateRepository) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *RepositoryTemplateRepository) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *RepositoryTemplateRepository) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetOpenIssuesCount

`func (o *RepositoryTemplateRepository) GetOpenIssuesCount() int32`

GetOpenIssuesCount returns the OpenIssuesCount field if non-nil, zero value otherwise.

### GetOpenIssuesCountOk

`func (o *RepositoryTemplateRepository) GetOpenIssuesCountOk() (*int32, bool)`

GetOpenIssuesCountOk returns a tuple with the OpenIssuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssuesCount

`func (o *RepositoryTemplateRepository) SetOpenIssuesCount(v int32)`

SetOpenIssuesCount sets OpenIssuesCount field to given value.

### HasOpenIssuesCount

`func (o *RepositoryTemplateRepository) HasOpenIssuesCount() bool`

HasOpenIssuesCount returns a boolean if a field has been set.

### GetIsTemplate

`func (o *RepositoryTemplateRepository) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *RepositoryTemplateRepository) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *RepositoryTemplateRepository) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *RepositoryTemplateRepository) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetTopics

`func (o *RepositoryTemplateRepository) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *RepositoryTemplateRepository) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *RepositoryTemplateRepository) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *RepositoryTemplateRepository) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetHasIssues

`func (o *RepositoryTemplateRepository) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *RepositoryTemplateRepository) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *RepositoryTemplateRepository) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.

### HasHasIssues

`func (o *RepositoryTemplateRepository) HasHasIssues() bool`

HasHasIssues returns a boolean if a field has been set.

### GetHasProjects

`func (o *RepositoryTemplateRepository) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *RepositoryTemplateRepository) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *RepositoryTemplateRepository) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.

### HasHasProjects

`func (o *RepositoryTemplateRepository) HasHasProjects() bool`

HasHasProjects returns a boolean if a field has been set.

### GetHasWiki

`func (o *RepositoryTemplateRepository) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *RepositoryTemplateRepository) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *RepositoryTemplateRepository) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.

### HasHasWiki

`func (o *RepositoryTemplateRepository) HasHasWiki() bool`

HasHasWiki returns a boolean if a field has been set.

### GetHasPages

`func (o *RepositoryTemplateRepository) GetHasPages() bool`

GetHasPages returns the HasPages field if non-nil, zero value otherwise.

### GetHasPagesOk

`func (o *RepositoryTemplateRepository) GetHasPagesOk() (*bool, bool)`

GetHasPagesOk returns a tuple with the HasPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPages

`func (o *RepositoryTemplateRepository) SetHasPages(v bool)`

SetHasPages sets HasPages field to given value.

### HasHasPages

`func (o *RepositoryTemplateRepository) HasHasPages() bool`

HasHasPages returns a boolean if a field has been set.

### GetHasDownloads

`func (o *RepositoryTemplateRepository) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *RepositoryTemplateRepository) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *RepositoryTemplateRepository) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.

### HasHasDownloads

`func (o *RepositoryTemplateRepository) HasHasDownloads() bool`

HasHasDownloads returns a boolean if a field has been set.

### GetArchived

`func (o *RepositoryTemplateRepository) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *RepositoryTemplateRepository) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *RepositoryTemplateRepository) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *RepositoryTemplateRepository) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDisabled

`func (o *RepositoryTemplateRepository) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RepositoryTemplateRepository) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RepositoryTemplateRepository) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RepositoryTemplateRepository) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetVisibility

`func (o *RepositoryTemplateRepository) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RepositoryTemplateRepository) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RepositoryTemplateRepository) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *RepositoryTemplateRepository) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetPushedAt

`func (o *RepositoryTemplateRepository) GetPushedAt() string`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *RepositoryTemplateRepository) GetPushedAtOk() (*string, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *RepositoryTemplateRepository) SetPushedAt(v string)`

SetPushedAt sets PushedAt field to given value.

### HasPushedAt

`func (o *RepositoryTemplateRepository) HasPushedAt() bool`

HasPushedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepositoryTemplateRepository) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositoryTemplateRepository) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositoryTemplateRepository) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepositoryTemplateRepository) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RepositoryTemplateRepository) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepositoryTemplateRepository) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepositoryTemplateRepository) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepositoryTemplateRepository) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPermissions

`func (o *RepositoryTemplateRepository) GetPermissions() RepositoryTemplateRepositoryPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RepositoryTemplateRepository) GetPermissionsOk() (*RepositoryTemplateRepositoryPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RepositoryTemplateRepository) SetPermissions(v RepositoryTemplateRepositoryPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RepositoryTemplateRepository) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *RepositoryTemplateRepository) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *RepositoryTemplateRepository) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *RepositoryTemplateRepository) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *RepositoryTemplateRepository) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetTempCloneToken

`func (o *RepositoryTemplateRepository) GetTempCloneToken() string`

GetTempCloneToken returns the TempCloneToken field if non-nil, zero value otherwise.

### GetTempCloneTokenOk

`func (o *RepositoryTemplateRepository) GetTempCloneTokenOk() (*string, bool)`

GetTempCloneTokenOk returns a tuple with the TempCloneToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempCloneToken

`func (o *RepositoryTemplateRepository) SetTempCloneToken(v string)`

SetTempCloneToken sets TempCloneToken field to given value.

### HasTempCloneToken

`func (o *RepositoryTemplateRepository) HasTempCloneToken() bool`

HasTempCloneToken returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *RepositoryTemplateRepository) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *RepositoryTemplateRepository) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *RepositoryTemplateRepository) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *RepositoryTemplateRepository) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *RepositoryTemplateRepository) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *RepositoryTemplateRepository) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *RepositoryTemplateRepository) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *RepositoryTemplateRepository) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *RepositoryTemplateRepository) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *RepositoryTemplateRepository) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *RepositoryTemplateRepository) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *RepositoryTemplateRepository) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetAllowUpdateBranch

`func (o *RepositoryTemplateRepository) GetAllowUpdateBranch() bool`

GetAllowUpdateBranch returns the AllowUpdateBranch field if non-nil, zero value otherwise.

### GetAllowUpdateBranchOk

`func (o *RepositoryTemplateRepository) GetAllowUpdateBranchOk() (*bool, bool)`

GetAllowUpdateBranchOk returns a tuple with the AllowUpdateBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdateBranch

`func (o *RepositoryTemplateRepository) SetAllowUpdateBranch(v bool)`

SetAllowUpdateBranch sets AllowUpdateBranch field to given value.

### HasAllowUpdateBranch

`func (o *RepositoryTemplateRepository) HasAllowUpdateBranch() bool`

HasAllowUpdateBranch returns a boolean if a field has been set.

### GetUseSquashPrTitleAsDefault

`func (o *RepositoryTemplateRepository) GetUseSquashPrTitleAsDefault() bool`

GetUseSquashPrTitleAsDefault returns the UseSquashPrTitleAsDefault field if non-nil, zero value otherwise.

### GetUseSquashPrTitleAsDefaultOk

`func (o *RepositoryTemplateRepository) GetUseSquashPrTitleAsDefaultOk() (*bool, bool)`

GetUseSquashPrTitleAsDefaultOk returns a tuple with the UseSquashPrTitleAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSquashPrTitleAsDefault

`func (o *RepositoryTemplateRepository) SetUseSquashPrTitleAsDefault(v bool)`

SetUseSquashPrTitleAsDefault sets UseSquashPrTitleAsDefault field to given value.

### HasUseSquashPrTitleAsDefault

`func (o *RepositoryTemplateRepository) HasUseSquashPrTitleAsDefault() bool`

HasUseSquashPrTitleAsDefault returns a boolean if a field has been set.

### GetSquashMergeCommitTitle

`func (o *RepositoryTemplateRepository) GetSquashMergeCommitTitle() string`

GetSquashMergeCommitTitle returns the SquashMergeCommitTitle field if non-nil, zero value otherwise.

### GetSquashMergeCommitTitleOk

`func (o *RepositoryTemplateRepository) GetSquashMergeCommitTitleOk() (*string, bool)`

GetSquashMergeCommitTitleOk returns a tuple with the SquashMergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitTitle

`func (o *RepositoryTemplateRepository) SetSquashMergeCommitTitle(v string)`

SetSquashMergeCommitTitle sets SquashMergeCommitTitle field to given value.

### HasSquashMergeCommitTitle

`func (o *RepositoryTemplateRepository) HasSquashMergeCommitTitle() bool`

HasSquashMergeCommitTitle returns a boolean if a field has been set.

### GetSquashMergeCommitMessage

`func (o *RepositoryTemplateRepository) GetSquashMergeCommitMessage() string`

GetSquashMergeCommitMessage returns the SquashMergeCommitMessage field if non-nil, zero value otherwise.

### GetSquashMergeCommitMessageOk

`func (o *RepositoryTemplateRepository) GetSquashMergeCommitMessageOk() (*string, bool)`

GetSquashMergeCommitMessageOk returns a tuple with the SquashMergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitMessage

`func (o *RepositoryTemplateRepository) SetSquashMergeCommitMessage(v string)`

SetSquashMergeCommitMessage sets SquashMergeCommitMessage field to given value.

### HasSquashMergeCommitMessage

`func (o *RepositoryTemplateRepository) HasSquashMergeCommitMessage() bool`

HasSquashMergeCommitMessage returns a boolean if a field has been set.

### GetMergeCommitTitle

`func (o *RepositoryTemplateRepository) GetMergeCommitTitle() string`

GetMergeCommitTitle returns the MergeCommitTitle field if non-nil, zero value otherwise.

### GetMergeCommitTitleOk

`func (o *RepositoryTemplateRepository) GetMergeCommitTitleOk() (*string, bool)`

GetMergeCommitTitleOk returns a tuple with the MergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitTitle

`func (o *RepositoryTemplateRepository) SetMergeCommitTitle(v string)`

SetMergeCommitTitle sets MergeCommitTitle field to given value.

### HasMergeCommitTitle

`func (o *RepositoryTemplateRepository) HasMergeCommitTitle() bool`

HasMergeCommitTitle returns a boolean if a field has been set.

### GetMergeCommitMessage

`func (o *RepositoryTemplateRepository) GetMergeCommitMessage() string`

GetMergeCommitMessage returns the MergeCommitMessage field if non-nil, zero value otherwise.

### GetMergeCommitMessageOk

`func (o *RepositoryTemplateRepository) GetMergeCommitMessageOk() (*string, bool)`

GetMergeCommitMessageOk returns a tuple with the MergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitMessage

`func (o *RepositoryTemplateRepository) SetMergeCommitMessage(v string)`

SetMergeCommitMessage sets MergeCommitMessage field to given value.

### HasMergeCommitMessage

`func (o *RepositoryTemplateRepository) HasMergeCommitMessage() bool`

HasMergeCommitMessage returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *RepositoryTemplateRepository) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *RepositoryTemplateRepository) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *RepositoryTemplateRepository) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *RepositoryTemplateRepository) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetSubscribersCount

`func (o *RepositoryTemplateRepository) GetSubscribersCount() int32`

GetSubscribersCount returns the SubscribersCount field if non-nil, zero value otherwise.

### GetSubscribersCountOk

`func (o *RepositoryTemplateRepository) GetSubscribersCountOk() (*int32, bool)`

GetSubscribersCountOk returns a tuple with the SubscribersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersCount

`func (o *RepositoryTemplateRepository) SetSubscribersCount(v int32)`

SetSubscribersCount sets SubscribersCount field to given value.

### HasSubscribersCount

`func (o *RepositoryTemplateRepository) HasSubscribersCount() bool`

HasSubscribersCount returns a boolean if a field has been set.

### GetNetworkCount

`func (o *RepositoryTemplateRepository) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *RepositoryTemplateRepository) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *RepositoryTemplateRepository) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *RepositoryTemplateRepository) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


