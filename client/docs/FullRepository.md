# FullRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** |  | 
**FullName** | **string** |  | 
**Owner** | [**SimpleUser**](SimpleUser.md) |  | 
**Private** | **bool** |  | 
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
**DefaultBranch** | **string** |  | 
**OpenIssuesCount** | **int32** |  | 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**HasIssues** | **bool** |  | 
**HasProjects** | **bool** |  | 
**HasWiki** | **bool** |  | 
**HasPages** | **bool** |  | 
**HasDownloads** | **bool** |  | 
**Archived** | **bool** |  | 
**Disabled** | **bool** | Returns whether or not this repository disabled. | 
**Visibility** | Pointer to **string** | The repository visibility: public, private, or internal. | [optional] 
**PushedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Permissions** | Pointer to [**FullRepositoryPermissions**](FullRepositoryPermissions.md) |  | [optional] 
**AllowRebaseMerge** | Pointer to **bool** |  | [optional] 
**TemplateRepository** | Pointer to [**NullableNullableRepository**](NullableRepository.md) |  | [optional] 
**TempCloneToken** | Pointer to **NullableString** |  | [optional] 
**AllowSquashMerge** | Pointer to **bool** |  | [optional] 
**AllowAutoMerge** | Pointer to **bool** |  | [optional] 
**DeleteBranchOnMerge** | Pointer to **bool** |  | [optional] 
**AllowMergeCommit** | Pointer to **bool** |  | [optional] 
**AllowUpdateBranch** | Pointer to **bool** |  | [optional] 
**UseSquashPrTitleAsDefault** | Pointer to **bool** |  | [optional] 
**SquashMergeCommitTitle** | Pointer to **string** | The default value for a squash merge commit title:  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;COMMIT_OR_PR_TITLE&#x60; - default to the commit&#39;s title (if only one commit) or the pull request&#39;s title (when more than one commit). | [optional] 
**SquashMergeCommitMessage** | Pointer to **string** | The default value for a squash merge commit message:  - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;COMMIT_MESSAGES&#x60; - default to the branch&#39;s commit messages. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**MergeCommitTitle** | Pointer to **string** | The default value for a merge commit title.    - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title.   - &#x60;MERGE_MESSAGE&#x60; - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name). | [optional] 
**MergeCommitMessage** | Pointer to **string** | The default value for a merge commit message.  - &#x60;PR_TITLE&#x60; - default to the pull request&#39;s title. - &#x60;PR_BODY&#x60; - default to the pull request&#39;s body. - &#x60;BLANK&#x60; - default to a blank commit message. | [optional] 
**AllowForking** | Pointer to **bool** |  | [optional] 
**WebCommitSignoffRequired** | Pointer to **bool** |  | [optional] 
**SubscribersCount** | **int32** |  | 
**NetworkCount** | **int32** |  | 
**License** | [**NullableNullableLicenseSimple**](NullableLicenseSimple.md) |  | 
**Organization** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**Parent** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Source** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Forks** | **int32** |  | 
**MasterBranch** | Pointer to **string** |  | [optional] 
**OpenIssues** | **int32** |  | 
**Watchers** | **int32** |  | 
**AnonymousAccessEnabled** | Pointer to **bool** | Whether anonymous git access is allowed. | [optional] [default to true]
**CodeOfConduct** | Pointer to [**CodeOfConductSimple**](CodeOfConductSimple.md) |  | [optional] 
**SecurityAndAnalysis** | Pointer to [**NullableSecurityAndAnalysis**](SecurityAndAnalysis.md) |  | [optional] 

## Methods

### NewFullRepository

`func NewFullRepository(id int32, nodeId string, name string, fullName string, owner SimpleUser, private bool, htmlUrl string, description NullableString, fork bool, url string, archiveUrl string, assigneesUrl string, blobsUrl string, branchesUrl string, collaboratorsUrl string, commentsUrl string, commitsUrl string, compareUrl string, contentsUrl string, contributorsUrl string, deploymentsUrl string, downloadsUrl string, eventsUrl string, forksUrl string, gitCommitsUrl string, gitRefsUrl string, gitTagsUrl string, gitUrl string, issueCommentUrl string, issueEventsUrl string, issuesUrl string, keysUrl string, labelsUrl string, languagesUrl string, mergesUrl string, milestonesUrl string, notificationsUrl string, pullsUrl string, releasesUrl string, sshUrl string, stargazersUrl string, statusesUrl string, subscribersUrl string, subscriptionUrl string, tagsUrl string, teamsUrl string, treesUrl string, cloneUrl string, mirrorUrl NullableString, hooksUrl string, svnUrl string, homepage NullableString, language NullableString, forksCount int32, stargazersCount int32, watchersCount int32, size int32, defaultBranch string, openIssuesCount int32, hasIssues bool, hasProjects bool, hasWiki bool, hasPages bool, hasDownloads bool, archived bool, disabled bool, pushedAt time.Time, createdAt time.Time, updatedAt time.Time, subscribersCount int32, networkCount int32, license NullableNullableLicenseSimple, forks int32, openIssues int32, watchers int32, ) *FullRepository`

NewFullRepository instantiates a new FullRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullRepositoryWithDefaults

`func NewFullRepositoryWithDefaults() *FullRepository`

NewFullRepositoryWithDefaults instantiates a new FullRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullRepository) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullRepository) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullRepository) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *FullRepository) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FullRepository) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FullRepository) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *FullRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullRepository) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *FullRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *FullRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *FullRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetOwner

`func (o *FullRepository) GetOwner() SimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FullRepository) GetOwnerOk() (*SimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FullRepository) SetOwner(v SimpleUser)`

SetOwner sets Owner field to given value.


### GetPrivate

`func (o *FullRepository) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FullRepository) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FullRepository) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetHtmlUrl

`func (o *FullRepository) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *FullRepository) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *FullRepository) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDescription

`func (o *FullRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullRepository) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *FullRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FullRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFork

`func (o *FullRepository) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *FullRepository) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *FullRepository) SetFork(v bool)`

SetFork sets Fork field to given value.


### GetUrl

`func (o *FullRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FullRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FullRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetArchiveUrl

`func (o *FullRepository) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *FullRepository) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *FullRepository) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.


### GetAssigneesUrl

`func (o *FullRepository) GetAssigneesUrl() string`

GetAssigneesUrl returns the AssigneesUrl field if non-nil, zero value otherwise.

### GetAssigneesUrlOk

`func (o *FullRepository) GetAssigneesUrlOk() (*string, bool)`

GetAssigneesUrlOk returns a tuple with the AssigneesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesUrl

`func (o *FullRepository) SetAssigneesUrl(v string)`

SetAssigneesUrl sets AssigneesUrl field to given value.


### GetBlobsUrl

`func (o *FullRepository) GetBlobsUrl() string`

GetBlobsUrl returns the BlobsUrl field if non-nil, zero value otherwise.

### GetBlobsUrlOk

`func (o *FullRepository) GetBlobsUrlOk() (*string, bool)`

GetBlobsUrlOk returns a tuple with the BlobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsUrl

`func (o *FullRepository) SetBlobsUrl(v string)`

SetBlobsUrl sets BlobsUrl field to given value.


### GetBranchesUrl

`func (o *FullRepository) GetBranchesUrl() string`

GetBranchesUrl returns the BranchesUrl field if non-nil, zero value otherwise.

### GetBranchesUrlOk

`func (o *FullRepository) GetBranchesUrlOk() (*string, bool)`

GetBranchesUrlOk returns a tuple with the BranchesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesUrl

`func (o *FullRepository) SetBranchesUrl(v string)`

SetBranchesUrl sets BranchesUrl field to given value.


### GetCollaboratorsUrl

`func (o *FullRepository) GetCollaboratorsUrl() string`

GetCollaboratorsUrl returns the CollaboratorsUrl field if non-nil, zero value otherwise.

### GetCollaboratorsUrlOk

`func (o *FullRepository) GetCollaboratorsUrlOk() (*string, bool)`

GetCollaboratorsUrlOk returns a tuple with the CollaboratorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorsUrl

`func (o *FullRepository) SetCollaboratorsUrl(v string)`

SetCollaboratorsUrl sets CollaboratorsUrl field to given value.


### GetCommentsUrl

`func (o *FullRepository) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *FullRepository) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *FullRepository) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCommitsUrl

`func (o *FullRepository) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *FullRepository) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *FullRepository) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetCompareUrl

`func (o *FullRepository) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *FullRepository) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *FullRepository) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.


### GetContentsUrl

`func (o *FullRepository) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *FullRepository) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *FullRepository) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetContributorsUrl

`func (o *FullRepository) GetContributorsUrl() string`

GetContributorsUrl returns the ContributorsUrl field if non-nil, zero value otherwise.

### GetContributorsUrlOk

`func (o *FullRepository) GetContributorsUrlOk() (*string, bool)`

GetContributorsUrlOk returns a tuple with the ContributorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsUrl

`func (o *FullRepository) SetContributorsUrl(v string)`

SetContributorsUrl sets ContributorsUrl field to given value.


### GetDeploymentsUrl

`func (o *FullRepository) GetDeploymentsUrl() string`

GetDeploymentsUrl returns the DeploymentsUrl field if non-nil, zero value otherwise.

### GetDeploymentsUrlOk

`func (o *FullRepository) GetDeploymentsUrlOk() (*string, bool)`

GetDeploymentsUrlOk returns a tuple with the DeploymentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsUrl

`func (o *FullRepository) SetDeploymentsUrl(v string)`

SetDeploymentsUrl sets DeploymentsUrl field to given value.


### GetDownloadsUrl

`func (o *FullRepository) GetDownloadsUrl() string`

GetDownloadsUrl returns the DownloadsUrl field if non-nil, zero value otherwise.

### GetDownloadsUrlOk

`func (o *FullRepository) GetDownloadsUrlOk() (*string, bool)`

GetDownloadsUrlOk returns a tuple with the DownloadsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsUrl

`func (o *FullRepository) SetDownloadsUrl(v string)`

SetDownloadsUrl sets DownloadsUrl field to given value.


### GetEventsUrl

`func (o *FullRepository) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *FullRepository) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *FullRepository) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetForksUrl

`func (o *FullRepository) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *FullRepository) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *FullRepository) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.


### GetGitCommitsUrl

`func (o *FullRepository) GetGitCommitsUrl() string`

GetGitCommitsUrl returns the GitCommitsUrl field if non-nil, zero value otherwise.

### GetGitCommitsUrlOk

`func (o *FullRepository) GetGitCommitsUrlOk() (*string, bool)`

GetGitCommitsUrlOk returns a tuple with the GitCommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitsUrl

`func (o *FullRepository) SetGitCommitsUrl(v string)`

SetGitCommitsUrl sets GitCommitsUrl field to given value.


### GetGitRefsUrl

`func (o *FullRepository) GetGitRefsUrl() string`

GetGitRefsUrl returns the GitRefsUrl field if non-nil, zero value otherwise.

### GetGitRefsUrlOk

`func (o *FullRepository) GetGitRefsUrlOk() (*string, bool)`

GetGitRefsUrlOk returns a tuple with the GitRefsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefsUrl

`func (o *FullRepository) SetGitRefsUrl(v string)`

SetGitRefsUrl sets GitRefsUrl field to given value.


### GetGitTagsUrl

`func (o *FullRepository) GetGitTagsUrl() string`

GetGitTagsUrl returns the GitTagsUrl field if non-nil, zero value otherwise.

### GetGitTagsUrlOk

`func (o *FullRepository) GetGitTagsUrlOk() (*string, bool)`

GetGitTagsUrlOk returns a tuple with the GitTagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagsUrl

`func (o *FullRepository) SetGitTagsUrl(v string)`

SetGitTagsUrl sets GitTagsUrl field to given value.


### GetGitUrl

`func (o *FullRepository) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *FullRepository) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *FullRepository) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### GetIssueCommentUrl

`func (o *FullRepository) GetIssueCommentUrl() string`

GetIssueCommentUrl returns the IssueCommentUrl field if non-nil, zero value otherwise.

### GetIssueCommentUrlOk

`func (o *FullRepository) GetIssueCommentUrlOk() (*string, bool)`

GetIssueCommentUrlOk returns a tuple with the IssueCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCommentUrl

`func (o *FullRepository) SetIssueCommentUrl(v string)`

SetIssueCommentUrl sets IssueCommentUrl field to given value.


### GetIssueEventsUrl

`func (o *FullRepository) GetIssueEventsUrl() string`

GetIssueEventsUrl returns the IssueEventsUrl field if non-nil, zero value otherwise.

### GetIssueEventsUrlOk

`func (o *FullRepository) GetIssueEventsUrlOk() (*string, bool)`

GetIssueEventsUrlOk returns a tuple with the IssueEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEventsUrl

`func (o *FullRepository) SetIssueEventsUrl(v string)`

SetIssueEventsUrl sets IssueEventsUrl field to given value.


### GetIssuesUrl

`func (o *FullRepository) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *FullRepository) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *FullRepository) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetKeysUrl

`func (o *FullRepository) GetKeysUrl() string`

GetKeysUrl returns the KeysUrl field if non-nil, zero value otherwise.

### GetKeysUrlOk

`func (o *FullRepository) GetKeysUrlOk() (*string, bool)`

GetKeysUrlOk returns a tuple with the KeysUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysUrl

`func (o *FullRepository) SetKeysUrl(v string)`

SetKeysUrl sets KeysUrl field to given value.


### GetLabelsUrl

`func (o *FullRepository) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *FullRepository) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *FullRepository) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetLanguagesUrl

`func (o *FullRepository) GetLanguagesUrl() string`

GetLanguagesUrl returns the LanguagesUrl field if non-nil, zero value otherwise.

### GetLanguagesUrlOk

`func (o *FullRepository) GetLanguagesUrlOk() (*string, bool)`

GetLanguagesUrlOk returns a tuple with the LanguagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesUrl

`func (o *FullRepository) SetLanguagesUrl(v string)`

SetLanguagesUrl sets LanguagesUrl field to given value.


### GetMergesUrl

`func (o *FullRepository) GetMergesUrl() string`

GetMergesUrl returns the MergesUrl field if non-nil, zero value otherwise.

### GetMergesUrlOk

`func (o *FullRepository) GetMergesUrlOk() (*string, bool)`

GetMergesUrlOk returns a tuple with the MergesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesUrl

`func (o *FullRepository) SetMergesUrl(v string)`

SetMergesUrl sets MergesUrl field to given value.


### GetMilestonesUrl

`func (o *FullRepository) GetMilestonesUrl() string`

GetMilestonesUrl returns the MilestonesUrl field if non-nil, zero value otherwise.

### GetMilestonesUrlOk

`func (o *FullRepository) GetMilestonesUrlOk() (*string, bool)`

GetMilestonesUrlOk returns a tuple with the MilestonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestonesUrl

`func (o *FullRepository) SetMilestonesUrl(v string)`

SetMilestonesUrl sets MilestonesUrl field to given value.


### GetNotificationsUrl

`func (o *FullRepository) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *FullRepository) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *FullRepository) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.


### GetPullsUrl

`func (o *FullRepository) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *FullRepository) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *FullRepository) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.


### GetReleasesUrl

`func (o *FullRepository) GetReleasesUrl() string`

GetReleasesUrl returns the ReleasesUrl field if non-nil, zero value otherwise.

### GetReleasesUrlOk

`func (o *FullRepository) GetReleasesUrlOk() (*string, bool)`

GetReleasesUrlOk returns a tuple with the ReleasesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesUrl

`func (o *FullRepository) SetReleasesUrl(v string)`

SetReleasesUrl sets ReleasesUrl field to given value.


### GetSshUrl

`func (o *FullRepository) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *FullRepository) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *FullRepository) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.


### GetStargazersUrl

`func (o *FullRepository) GetStargazersUrl() string`

GetStargazersUrl returns the StargazersUrl field if non-nil, zero value otherwise.

### GetStargazersUrlOk

`func (o *FullRepository) GetStargazersUrlOk() (*string, bool)`

GetStargazersUrlOk returns a tuple with the StargazersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersUrl

`func (o *FullRepository) SetStargazersUrl(v string)`

SetStargazersUrl sets StargazersUrl field to given value.


### GetStatusesUrl

`func (o *FullRepository) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *FullRepository) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *FullRepository) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetSubscribersUrl

`func (o *FullRepository) GetSubscribersUrl() string`

GetSubscribersUrl returns the SubscribersUrl field if non-nil, zero value otherwise.

### GetSubscribersUrlOk

`func (o *FullRepository) GetSubscribersUrlOk() (*string, bool)`

GetSubscribersUrlOk returns a tuple with the SubscribersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersUrl

`func (o *FullRepository) SetSubscribersUrl(v string)`

SetSubscribersUrl sets SubscribersUrl field to given value.


### GetSubscriptionUrl

`func (o *FullRepository) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *FullRepository) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *FullRepository) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetTagsUrl

`func (o *FullRepository) GetTagsUrl() string`

GetTagsUrl returns the TagsUrl field if non-nil, zero value otherwise.

### GetTagsUrlOk

`func (o *FullRepository) GetTagsUrlOk() (*string, bool)`

GetTagsUrlOk returns a tuple with the TagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsUrl

`func (o *FullRepository) SetTagsUrl(v string)`

SetTagsUrl sets TagsUrl field to given value.


### GetTeamsUrl

`func (o *FullRepository) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *FullRepository) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *FullRepository) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetTreesUrl

`func (o *FullRepository) GetTreesUrl() string`

GetTreesUrl returns the TreesUrl field if non-nil, zero value otherwise.

### GetTreesUrlOk

`func (o *FullRepository) GetTreesUrlOk() (*string, bool)`

GetTreesUrlOk returns a tuple with the TreesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreesUrl

`func (o *FullRepository) SetTreesUrl(v string)`

SetTreesUrl sets TreesUrl field to given value.


### GetCloneUrl

`func (o *FullRepository) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *FullRepository) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *FullRepository) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetMirrorUrl

`func (o *FullRepository) GetMirrorUrl() string`

GetMirrorUrl returns the MirrorUrl field if non-nil, zero value otherwise.

### GetMirrorUrlOk

`func (o *FullRepository) GetMirrorUrlOk() (*string, bool)`

GetMirrorUrlOk returns a tuple with the MirrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorUrl

`func (o *FullRepository) SetMirrorUrl(v string)`

SetMirrorUrl sets MirrorUrl field to given value.


### SetMirrorUrlNil

`func (o *FullRepository) SetMirrorUrlNil(b bool)`

 SetMirrorUrlNil sets the value for MirrorUrl to be an explicit nil

### UnsetMirrorUrl
`func (o *FullRepository) UnsetMirrorUrl()`

UnsetMirrorUrl ensures that no value is present for MirrorUrl, not even an explicit nil
### GetHooksUrl

`func (o *FullRepository) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *FullRepository) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *FullRepository) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetSvnUrl

`func (o *FullRepository) GetSvnUrl() string`

GetSvnUrl returns the SvnUrl field if non-nil, zero value otherwise.

### GetSvnUrlOk

`func (o *FullRepository) GetSvnUrlOk() (*string, bool)`

GetSvnUrlOk returns a tuple with the SvnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnUrl

`func (o *FullRepository) SetSvnUrl(v string)`

SetSvnUrl sets SvnUrl field to given value.


### GetHomepage

`func (o *FullRepository) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *FullRepository) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *FullRepository) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.


### SetHomepageNil

`func (o *FullRepository) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *FullRepository) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil
### GetLanguage

`func (o *FullRepository) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FullRepository) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FullRepository) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### SetLanguageNil

`func (o *FullRepository) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *FullRepository) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetForksCount

`func (o *FullRepository) GetForksCount() int32`

GetForksCount returns the ForksCount field if non-nil, zero value otherwise.

### GetForksCountOk

`func (o *FullRepository) GetForksCountOk() (*int32, bool)`

GetForksCountOk returns a tuple with the ForksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksCount

`func (o *FullRepository) SetForksCount(v int32)`

SetForksCount sets ForksCount field to given value.


### GetStargazersCount

`func (o *FullRepository) GetStargazersCount() int32`

GetStargazersCount returns the StargazersCount field if non-nil, zero value otherwise.

### GetStargazersCountOk

`func (o *FullRepository) GetStargazersCountOk() (*int32, bool)`

GetStargazersCountOk returns a tuple with the StargazersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersCount

`func (o *FullRepository) SetStargazersCount(v int32)`

SetStargazersCount sets StargazersCount field to given value.


### GetWatchersCount

`func (o *FullRepository) GetWatchersCount() int32`

GetWatchersCount returns the WatchersCount field if non-nil, zero value otherwise.

### GetWatchersCountOk

`func (o *FullRepository) GetWatchersCountOk() (*int32, bool)`

GetWatchersCountOk returns a tuple with the WatchersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchersCount

`func (o *FullRepository) SetWatchersCount(v int32)`

SetWatchersCount sets WatchersCount field to given value.


### GetSize

`func (o *FullRepository) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FullRepository) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FullRepository) SetSize(v int32)`

SetSize sets Size field to given value.


### GetDefaultBranch

`func (o *FullRepository) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *FullRepository) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *FullRepository) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetOpenIssuesCount

`func (o *FullRepository) GetOpenIssuesCount() int32`

GetOpenIssuesCount returns the OpenIssuesCount field if non-nil, zero value otherwise.

### GetOpenIssuesCountOk

`func (o *FullRepository) GetOpenIssuesCountOk() (*int32, bool)`

GetOpenIssuesCountOk returns a tuple with the OpenIssuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssuesCount

`func (o *FullRepository) SetOpenIssuesCount(v int32)`

SetOpenIssuesCount sets OpenIssuesCount field to given value.


### GetIsTemplate

`func (o *FullRepository) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *FullRepository) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *FullRepository) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *FullRepository) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetTopics

`func (o *FullRepository) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *FullRepository) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *FullRepository) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *FullRepository) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetHasIssues

`func (o *FullRepository) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *FullRepository) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *FullRepository) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.


### GetHasProjects

`func (o *FullRepository) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *FullRepository) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *FullRepository) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.


### GetHasWiki

`func (o *FullRepository) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *FullRepository) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *FullRepository) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.


### GetHasPages

`func (o *FullRepository) GetHasPages() bool`

GetHasPages returns the HasPages field if non-nil, zero value otherwise.

### GetHasPagesOk

`func (o *FullRepository) GetHasPagesOk() (*bool, bool)`

GetHasPagesOk returns a tuple with the HasPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPages

`func (o *FullRepository) SetHasPages(v bool)`

SetHasPages sets HasPages field to given value.


### GetHasDownloads

`func (o *FullRepository) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *FullRepository) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *FullRepository) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.


### GetArchived

`func (o *FullRepository) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FullRepository) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FullRepository) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetDisabled

`func (o *FullRepository) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FullRepository) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FullRepository) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetVisibility

`func (o *FullRepository) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FullRepository) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FullRepository) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FullRepository) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetPushedAt

`func (o *FullRepository) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *FullRepository) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *FullRepository) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.


### GetCreatedAt

`func (o *FullRepository) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullRepository) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullRepository) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FullRepository) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FullRepository) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FullRepository) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPermissions

`func (o *FullRepository) GetPermissions() FullRepositoryPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FullRepository) GetPermissionsOk() (*FullRepositoryPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FullRepository) SetPermissions(v FullRepositoryPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FullRepository) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *FullRepository) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *FullRepository) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *FullRepository) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *FullRepository) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetTemplateRepository

`func (o *FullRepository) GetTemplateRepository() NullableRepository`

GetTemplateRepository returns the TemplateRepository field if non-nil, zero value otherwise.

### GetTemplateRepositoryOk

`func (o *FullRepository) GetTemplateRepositoryOk() (*NullableRepository, bool)`

GetTemplateRepositoryOk returns a tuple with the TemplateRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRepository

`func (o *FullRepository) SetTemplateRepository(v NullableRepository)`

SetTemplateRepository sets TemplateRepository field to given value.

### HasTemplateRepository

`func (o *FullRepository) HasTemplateRepository() bool`

HasTemplateRepository returns a boolean if a field has been set.

### SetTemplateRepositoryNil

`func (o *FullRepository) SetTemplateRepositoryNil(b bool)`

 SetTemplateRepositoryNil sets the value for TemplateRepository to be an explicit nil

### UnsetTemplateRepository
`func (o *FullRepository) UnsetTemplateRepository()`

UnsetTemplateRepository ensures that no value is present for TemplateRepository, not even an explicit nil
### GetTempCloneToken

`func (o *FullRepository) GetTempCloneToken() string`

GetTempCloneToken returns the TempCloneToken field if non-nil, zero value otherwise.

### GetTempCloneTokenOk

`func (o *FullRepository) GetTempCloneTokenOk() (*string, bool)`

GetTempCloneTokenOk returns a tuple with the TempCloneToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempCloneToken

`func (o *FullRepository) SetTempCloneToken(v string)`

SetTempCloneToken sets TempCloneToken field to given value.

### HasTempCloneToken

`func (o *FullRepository) HasTempCloneToken() bool`

HasTempCloneToken returns a boolean if a field has been set.

### SetTempCloneTokenNil

`func (o *FullRepository) SetTempCloneTokenNil(b bool)`

 SetTempCloneTokenNil sets the value for TempCloneToken to be an explicit nil

### UnsetTempCloneToken
`func (o *FullRepository) UnsetTempCloneToken()`

UnsetTempCloneToken ensures that no value is present for TempCloneToken, not even an explicit nil
### GetAllowSquashMerge

`func (o *FullRepository) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *FullRepository) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *FullRepository) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *FullRepository) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *FullRepository) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *FullRepository) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *FullRepository) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *FullRepository) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *FullRepository) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *FullRepository) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *FullRepository) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *FullRepository) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *FullRepository) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *FullRepository) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *FullRepository) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *FullRepository) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowUpdateBranch

`func (o *FullRepository) GetAllowUpdateBranch() bool`

GetAllowUpdateBranch returns the AllowUpdateBranch field if non-nil, zero value otherwise.

### GetAllowUpdateBranchOk

`func (o *FullRepository) GetAllowUpdateBranchOk() (*bool, bool)`

GetAllowUpdateBranchOk returns a tuple with the AllowUpdateBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdateBranch

`func (o *FullRepository) SetAllowUpdateBranch(v bool)`

SetAllowUpdateBranch sets AllowUpdateBranch field to given value.

### HasAllowUpdateBranch

`func (o *FullRepository) HasAllowUpdateBranch() bool`

HasAllowUpdateBranch returns a boolean if a field has been set.

### GetUseSquashPrTitleAsDefault

`func (o *FullRepository) GetUseSquashPrTitleAsDefault() bool`

GetUseSquashPrTitleAsDefault returns the UseSquashPrTitleAsDefault field if non-nil, zero value otherwise.

### GetUseSquashPrTitleAsDefaultOk

`func (o *FullRepository) GetUseSquashPrTitleAsDefaultOk() (*bool, bool)`

GetUseSquashPrTitleAsDefaultOk returns a tuple with the UseSquashPrTitleAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSquashPrTitleAsDefault

`func (o *FullRepository) SetUseSquashPrTitleAsDefault(v bool)`

SetUseSquashPrTitleAsDefault sets UseSquashPrTitleAsDefault field to given value.

### HasUseSquashPrTitleAsDefault

`func (o *FullRepository) HasUseSquashPrTitleAsDefault() bool`

HasUseSquashPrTitleAsDefault returns a boolean if a field has been set.

### GetSquashMergeCommitTitle

`func (o *FullRepository) GetSquashMergeCommitTitle() string`

GetSquashMergeCommitTitle returns the SquashMergeCommitTitle field if non-nil, zero value otherwise.

### GetSquashMergeCommitTitleOk

`func (o *FullRepository) GetSquashMergeCommitTitleOk() (*string, bool)`

GetSquashMergeCommitTitleOk returns a tuple with the SquashMergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitTitle

`func (o *FullRepository) SetSquashMergeCommitTitle(v string)`

SetSquashMergeCommitTitle sets SquashMergeCommitTitle field to given value.

### HasSquashMergeCommitTitle

`func (o *FullRepository) HasSquashMergeCommitTitle() bool`

HasSquashMergeCommitTitle returns a boolean if a field has been set.

### GetSquashMergeCommitMessage

`func (o *FullRepository) GetSquashMergeCommitMessage() string`

GetSquashMergeCommitMessage returns the SquashMergeCommitMessage field if non-nil, zero value otherwise.

### GetSquashMergeCommitMessageOk

`func (o *FullRepository) GetSquashMergeCommitMessageOk() (*string, bool)`

GetSquashMergeCommitMessageOk returns a tuple with the SquashMergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashMergeCommitMessage

`func (o *FullRepository) SetSquashMergeCommitMessage(v string)`

SetSquashMergeCommitMessage sets SquashMergeCommitMessage field to given value.

### HasSquashMergeCommitMessage

`func (o *FullRepository) HasSquashMergeCommitMessage() bool`

HasSquashMergeCommitMessage returns a boolean if a field has been set.

### GetMergeCommitTitle

`func (o *FullRepository) GetMergeCommitTitle() string`

GetMergeCommitTitle returns the MergeCommitTitle field if non-nil, zero value otherwise.

### GetMergeCommitTitleOk

`func (o *FullRepository) GetMergeCommitTitleOk() (*string, bool)`

GetMergeCommitTitleOk returns a tuple with the MergeCommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitTitle

`func (o *FullRepository) SetMergeCommitTitle(v string)`

SetMergeCommitTitle sets MergeCommitTitle field to given value.

### HasMergeCommitTitle

`func (o *FullRepository) HasMergeCommitTitle() bool`

HasMergeCommitTitle returns a boolean if a field has been set.

### GetMergeCommitMessage

`func (o *FullRepository) GetMergeCommitMessage() string`

GetMergeCommitMessage returns the MergeCommitMessage field if non-nil, zero value otherwise.

### GetMergeCommitMessageOk

`func (o *FullRepository) GetMergeCommitMessageOk() (*string, bool)`

GetMergeCommitMessageOk returns a tuple with the MergeCommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitMessage

`func (o *FullRepository) SetMergeCommitMessage(v string)`

SetMergeCommitMessage sets MergeCommitMessage field to given value.

### HasMergeCommitMessage

`func (o *FullRepository) HasMergeCommitMessage() bool`

HasMergeCommitMessage returns a boolean if a field has been set.

### GetAllowForking

`func (o *FullRepository) GetAllowForking() bool`

GetAllowForking returns the AllowForking field if non-nil, zero value otherwise.

### GetAllowForkingOk

`func (o *FullRepository) GetAllowForkingOk() (*bool, bool)`

GetAllowForkingOk returns a tuple with the AllowForking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForking

`func (o *FullRepository) SetAllowForking(v bool)`

SetAllowForking sets AllowForking field to given value.

### HasAllowForking

`func (o *FullRepository) HasAllowForking() bool`

HasAllowForking returns a boolean if a field has been set.

### GetWebCommitSignoffRequired

`func (o *FullRepository) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *FullRepository) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *FullRepository) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *FullRepository) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.

### GetSubscribersCount

`func (o *FullRepository) GetSubscribersCount() int32`

GetSubscribersCount returns the SubscribersCount field if non-nil, zero value otherwise.

### GetSubscribersCountOk

`func (o *FullRepository) GetSubscribersCountOk() (*int32, bool)`

GetSubscribersCountOk returns a tuple with the SubscribersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersCount

`func (o *FullRepository) SetSubscribersCount(v int32)`

SetSubscribersCount sets SubscribersCount field to given value.


### GetNetworkCount

`func (o *FullRepository) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *FullRepository) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *FullRepository) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.


### GetLicense

`func (o *FullRepository) GetLicense() NullableLicenseSimple`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *FullRepository) GetLicenseOk() (*NullableLicenseSimple, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *FullRepository) SetLicense(v NullableLicenseSimple)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *FullRepository) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *FullRepository) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetOrganization

`func (o *FullRepository) GetOrganization() NullableSimpleUser`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FullRepository) GetOrganizationOk() (*NullableSimpleUser, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FullRepository) SetOrganization(v NullableSimpleUser)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FullRepository) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FullRepository) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FullRepository) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetParent

`func (o *FullRepository) GetParent() Repository`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FullRepository) GetParentOk() (*Repository, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FullRepository) SetParent(v Repository)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FullRepository) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSource

`func (o *FullRepository) GetSource() Repository`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FullRepository) GetSourceOk() (*Repository, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FullRepository) SetSource(v Repository)`

SetSource sets Source field to given value.

### HasSource

`func (o *FullRepository) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetForks

`func (o *FullRepository) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *FullRepository) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *FullRepository) SetForks(v int32)`

SetForks sets Forks field to given value.


### GetMasterBranch

`func (o *FullRepository) GetMasterBranch() string`

GetMasterBranch returns the MasterBranch field if non-nil, zero value otherwise.

### GetMasterBranchOk

`func (o *FullRepository) GetMasterBranchOk() (*string, bool)`

GetMasterBranchOk returns a tuple with the MasterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterBranch

`func (o *FullRepository) SetMasterBranch(v string)`

SetMasterBranch sets MasterBranch field to given value.

### HasMasterBranch

`func (o *FullRepository) HasMasterBranch() bool`

HasMasterBranch returns a boolean if a field has been set.

### GetOpenIssues

`func (o *FullRepository) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *FullRepository) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *FullRepository) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetWatchers

`func (o *FullRepository) GetWatchers() int32`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *FullRepository) GetWatchersOk() (*int32, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *FullRepository) SetWatchers(v int32)`

SetWatchers sets Watchers field to given value.


### GetAnonymousAccessEnabled

`func (o *FullRepository) GetAnonymousAccessEnabled() bool`

GetAnonymousAccessEnabled returns the AnonymousAccessEnabled field if non-nil, zero value otherwise.

### GetAnonymousAccessEnabledOk

`func (o *FullRepository) GetAnonymousAccessEnabledOk() (*bool, bool)`

GetAnonymousAccessEnabledOk returns a tuple with the AnonymousAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousAccessEnabled

`func (o *FullRepository) SetAnonymousAccessEnabled(v bool)`

SetAnonymousAccessEnabled sets AnonymousAccessEnabled field to given value.

### HasAnonymousAccessEnabled

`func (o *FullRepository) HasAnonymousAccessEnabled() bool`

HasAnonymousAccessEnabled returns a boolean if a field has been set.

### GetCodeOfConduct

`func (o *FullRepository) GetCodeOfConduct() CodeOfConductSimple`

GetCodeOfConduct returns the CodeOfConduct field if non-nil, zero value otherwise.

### GetCodeOfConductOk

`func (o *FullRepository) GetCodeOfConductOk() (*CodeOfConductSimple, bool)`

GetCodeOfConductOk returns a tuple with the CodeOfConduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOfConduct

`func (o *FullRepository) SetCodeOfConduct(v CodeOfConductSimple)`

SetCodeOfConduct sets CodeOfConduct field to given value.

### HasCodeOfConduct

`func (o *FullRepository) HasCodeOfConduct() bool`

HasCodeOfConduct returns a boolean if a field has been set.

### GetSecurityAndAnalysis

`func (o *FullRepository) GetSecurityAndAnalysis() SecurityAndAnalysis`

GetSecurityAndAnalysis returns the SecurityAndAnalysis field if non-nil, zero value otherwise.

### GetSecurityAndAnalysisOk

`func (o *FullRepository) GetSecurityAndAnalysisOk() (*SecurityAndAnalysis, bool)`

GetSecurityAndAnalysisOk returns a tuple with the SecurityAndAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAndAnalysis

`func (o *FullRepository) SetSecurityAndAnalysis(v SecurityAndAnalysis)`

SetSecurityAndAnalysis sets SecurityAndAnalysis field to given value.

### HasSecurityAndAnalysis

`func (o *FullRepository) HasSecurityAndAnalysis() bool`

HasSecurityAndAnalysis returns a boolean if a field has been set.

### SetSecurityAndAnalysisNil

`func (o *FullRepository) SetSecurityAndAnalysisNil(b bool)`

 SetSecurityAndAnalysisNil sets the value for SecurityAndAnalysis to be an explicit nil

### UnsetSecurityAndAnalysis
`func (o *FullRepository) UnsetSecurityAndAnalysis()`

UnsetSecurityAndAnalysis ensures that no value is present for SecurityAndAnalysis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


