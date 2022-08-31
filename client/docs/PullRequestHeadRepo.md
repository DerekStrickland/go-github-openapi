# PullRequestHeadRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Description** | **NullableString** |  | 
**DownloadsUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**Fork** | **bool** |  | 
**ForksUrl** | **string** |  | 
**FullName** | **string** |  | 
**GitCommitsUrl** | **string** |  | 
**GitRefsUrl** | **string** |  | 
**GitTagsUrl** | **string** |  | 
**HooksUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**IssueCommentUrl** | **string** |  | 
**IssueEventsUrl** | **string** |  | 
**IssuesUrl** | **string** |  | 
**KeysUrl** | **string** |  | 
**LabelsUrl** | **string** |  | 
**LanguagesUrl** | **string** |  | 
**MergesUrl** | **string** |  | 
**MilestonesUrl** | **string** |  | 
**Name** | **string** |  | 
**NotificationsUrl** | **string** |  | 
**Owner** | [**PullRequestHeadRepoOwner**](PullRequestHeadRepoOwner.md) |  | 
**Private** | **bool** |  | 
**PullsUrl** | **string** |  | 
**ReleasesUrl** | **string** |  | 
**StargazersUrl** | **string** |  | 
**StatusesUrl** | **string** |  | 
**SubscribersUrl** | **string** |  | 
**SubscriptionUrl** | **string** |  | 
**TagsUrl** | **string** |  | 
**TeamsUrl** | **string** |  | 
**TreesUrl** | **string** |  | 
**Url** | **string** |  | 
**CloneUrl** | **string** |  | 
**DefaultBranch** | **string** |  | 
**Forks** | **int32** |  | 
**ForksCount** | **int32** |  | 
**GitUrl** | **string** |  | 
**HasDownloads** | **bool** |  | 
**HasIssues** | **bool** |  | 
**HasProjects** | **bool** |  | 
**HasWiki** | **bool** |  | 
**HasPages** | **bool** |  | 
**Homepage** | **NullableString** |  | 
**Language** | **NullableString** |  | 
**MasterBranch** | Pointer to **string** |  | [optional] 
**Archived** | **bool** |  | 
**Disabled** | **bool** |  | 
**Visibility** | Pointer to **string** | The repository visibility: public, private, or internal. | [optional] 
**MirrorUrl** | **NullableString** |  | 
**OpenIssues** | **int32** |  | 
**OpenIssuesCount** | **int32** |  | 
**Permissions** | Pointer to [**FullRepositoryPermissions**](FullRepositoryPermissions.md) |  | [optional] 
**TempCloneToken** | Pointer to **string** |  | [optional] 
**AllowMergeCommit** | Pointer to **bool** |  | [optional] 
**AllowSquashMerge** | Pointer to **bool** |  | [optional] 
**AllowRebaseMerge** | Pointer to **bool** |  | [optional] 
**License** | [**NullablePullRequestHeadRepoLicense**](PullRequestHeadRepoLicense.md) |  | 
**PushedAt** | **time.Time** |  | 
**Size** | **int32** |  | 
**SshUrl** | **string** |  | 
**StargazersCount** | **int32** |  | 
**SvnUrl** | **string** |  | 
**Topics** | Pointer to **[]string** |  | [optional] 
**Watchers** | **int32** |  | 
**WatchersCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AllowForking** | Pointer to **bool** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**WebCommitSignoffRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewPullRequestHeadRepo

`func NewPullRequestHeadRepo(archiveUrl string, assigneesUrl string, blobsUrl string, branchesUrl string, collaboratorsUrl string, commentsUrl string, commitsUrl string, compareUrl string, contentsUrl string, contributorsUrl string, deploymentsUrl string, description NullableString, downloadsUrl string, eventsUrl string, fork bool, forksUrl string, fullName string, gitCommitsUrl string, gitRefsUrl string, gitTagsUrl string, hooksUrl string, htmlUrl string, id int32, nodeId string, issueCommentUrl string, issueEventsUrl string, issuesUrl string, keysUrl string, labelsUrl string, languagesUrl string, mergesUrl string, milestonesUrl string, name string, notificationsUrl string, owner PullRequestHeadRepoOwner, private bool, pullsUrl string, releasesUrl string, stargazersUrl string, statusesUrl string, subscribersUrl string, subscriptionUrl string, tagsUrl string, teamsUrl string, treesUrl string, url string, cloneUrl string, defaultBranch string, forks int32, forksCount int32, gitUrl string, hasDownloads bool, hasIssues bool, hasProjects bool, hasWiki bool, hasPages bool, homepage NullableString, language NullableString, archived bool, disabled bool, mirrorUrl NullableString, openIssues int32, openIssuesCount int32, license NullablePullRequestHeadRepoLicense, pushedAt time.Time, size int32, sshUrl string, stargazersCount int32, svnUrl string, watchers int32, watchersCount int32, createdAt time.Time, updatedAt time.Time, ) *PullRequestHeadRepo`

NewPullRequestHeadRepo instantiates a new PullRequestHeadRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestHeadRepoWithDefaults

`func NewPullRequestHeadRepoWithDefaults() *PullRequestHeadRepo`

NewPullRequestHeadRepoWithDefaults instantiates a new PullRequestHeadRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveUrl

`func (o *PullRequestHeadRepo) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *PullRequestHeadRepo) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *PullRequestHeadRepo) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.


### GetAssigneesUrl

`func (o *PullRequestHeadRepo) GetAssigneesUrl() string`

GetAssigneesUrl returns the AssigneesUrl field if non-nil, zero value otherwise.

### GetAssigneesUrlOk

`func (o *PullRequestHeadRepo) GetAssigneesUrlOk() (*string, bool)`

GetAssigneesUrlOk returns a tuple with the AssigneesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesUrl

`func (o *PullRequestHeadRepo) SetAssigneesUrl(v string)`

SetAssigneesUrl sets AssigneesUrl field to given value.


### GetBlobsUrl

`func (o *PullRequestHeadRepo) GetBlobsUrl() string`

GetBlobsUrl returns the BlobsUrl field if non-nil, zero value otherwise.

### GetBlobsUrlOk

`func (o *PullRequestHeadRepo) GetBlobsUrlOk() (*string, bool)`

GetBlobsUrlOk returns a tuple with the BlobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsUrl

`func (o *PullRequestHeadRepo) SetBlobsUrl(v string)`

SetBlobsUrl sets BlobsUrl field to given value.


### GetBranchesUrl

`func (o *PullRequestHeadRepo) GetBranchesUrl() string`

GetBranchesUrl returns the BranchesUrl field if non-nil, zero value otherwise.

### GetBranchesUrlOk

`func (o *PullRequestHeadRepo) GetBranchesUrlOk() (*string, bool)`

GetBranchesUrlOk returns a tuple with the BranchesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesUrl

`func (o *PullRequestHeadRepo) SetBranchesUrl(v string)`

SetBranchesUrl sets BranchesUrl field to given value.


### GetCollaboratorsUrl

`func (o *PullRequestHeadRepo) GetCollaboratorsUrl() string`

GetCollaboratorsUrl returns the CollaboratorsUrl field if non-nil, zero value otherwise.

### GetCollaboratorsUrlOk

`func (o *PullRequestHeadRepo) GetCollaboratorsUrlOk() (*string, bool)`

GetCollaboratorsUrlOk returns a tuple with the CollaboratorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorsUrl

`func (o *PullRequestHeadRepo) SetCollaboratorsUrl(v string)`

SetCollaboratorsUrl sets CollaboratorsUrl field to given value.


### GetCommentsUrl

`func (o *PullRequestHeadRepo) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *PullRequestHeadRepo) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *PullRequestHeadRepo) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCommitsUrl

`func (o *PullRequestHeadRepo) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *PullRequestHeadRepo) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *PullRequestHeadRepo) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetCompareUrl

`func (o *PullRequestHeadRepo) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *PullRequestHeadRepo) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *PullRequestHeadRepo) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.


### GetContentsUrl

`func (o *PullRequestHeadRepo) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *PullRequestHeadRepo) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *PullRequestHeadRepo) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetContributorsUrl

`func (o *PullRequestHeadRepo) GetContributorsUrl() string`

GetContributorsUrl returns the ContributorsUrl field if non-nil, zero value otherwise.

### GetContributorsUrlOk

`func (o *PullRequestHeadRepo) GetContributorsUrlOk() (*string, bool)`

GetContributorsUrlOk returns a tuple with the ContributorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsUrl

`func (o *PullRequestHeadRepo) SetContributorsUrl(v string)`

SetContributorsUrl sets ContributorsUrl field to given value.


### GetDeploymentsUrl

`func (o *PullRequestHeadRepo) GetDeploymentsUrl() string`

GetDeploymentsUrl returns the DeploymentsUrl field if non-nil, zero value otherwise.

### GetDeploymentsUrlOk

`func (o *PullRequestHeadRepo) GetDeploymentsUrlOk() (*string, bool)`

GetDeploymentsUrlOk returns a tuple with the DeploymentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsUrl

`func (o *PullRequestHeadRepo) SetDeploymentsUrl(v string)`

SetDeploymentsUrl sets DeploymentsUrl field to given value.


### GetDescription

`func (o *PullRequestHeadRepo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PullRequestHeadRepo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PullRequestHeadRepo) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PullRequestHeadRepo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PullRequestHeadRepo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDownloadsUrl

`func (o *PullRequestHeadRepo) GetDownloadsUrl() string`

GetDownloadsUrl returns the DownloadsUrl field if non-nil, zero value otherwise.

### GetDownloadsUrlOk

`func (o *PullRequestHeadRepo) GetDownloadsUrlOk() (*string, bool)`

GetDownloadsUrlOk returns a tuple with the DownloadsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsUrl

`func (o *PullRequestHeadRepo) SetDownloadsUrl(v string)`

SetDownloadsUrl sets DownloadsUrl field to given value.


### GetEventsUrl

`func (o *PullRequestHeadRepo) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *PullRequestHeadRepo) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *PullRequestHeadRepo) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetFork

`func (o *PullRequestHeadRepo) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *PullRequestHeadRepo) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *PullRequestHeadRepo) SetFork(v bool)`

SetFork sets Fork field to given value.


### GetForksUrl

`func (o *PullRequestHeadRepo) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *PullRequestHeadRepo) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *PullRequestHeadRepo) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.


### GetFullName

`func (o *PullRequestHeadRepo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PullRequestHeadRepo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PullRequestHeadRepo) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetGitCommitsUrl

`func (o *PullRequestHeadRepo) GetGitCommitsUrl() string`

GetGitCommitsUrl returns the GitCommitsUrl field if non-nil, zero value otherwise.

### GetGitCommitsUrlOk

`func (o *PullRequestHeadRepo) GetGitCommitsUrlOk() (*string, bool)`

GetGitCommitsUrlOk returns a tuple with the GitCommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitsUrl

`func (o *PullRequestHeadRepo) SetGitCommitsUrl(v string)`

SetGitCommitsUrl sets GitCommitsUrl field to given value.


### GetGitRefsUrl

`func (o *PullRequestHeadRepo) GetGitRefsUrl() string`

GetGitRefsUrl returns the GitRefsUrl field if non-nil, zero value otherwise.

### GetGitRefsUrlOk

`func (o *PullRequestHeadRepo) GetGitRefsUrlOk() (*string, bool)`

GetGitRefsUrlOk returns a tuple with the GitRefsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefsUrl

`func (o *PullRequestHeadRepo) SetGitRefsUrl(v string)`

SetGitRefsUrl sets GitRefsUrl field to given value.


### GetGitTagsUrl

`func (o *PullRequestHeadRepo) GetGitTagsUrl() string`

GetGitTagsUrl returns the GitTagsUrl field if non-nil, zero value otherwise.

### GetGitTagsUrlOk

`func (o *PullRequestHeadRepo) GetGitTagsUrlOk() (*string, bool)`

GetGitTagsUrlOk returns a tuple with the GitTagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagsUrl

`func (o *PullRequestHeadRepo) SetGitTagsUrl(v string)`

SetGitTagsUrl sets GitTagsUrl field to given value.


### GetHooksUrl

`func (o *PullRequestHeadRepo) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *PullRequestHeadRepo) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *PullRequestHeadRepo) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetHtmlUrl

`func (o *PullRequestHeadRepo) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PullRequestHeadRepo) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PullRequestHeadRepo) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetId

`func (o *PullRequestHeadRepo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestHeadRepo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestHeadRepo) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PullRequestHeadRepo) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PullRequestHeadRepo) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PullRequestHeadRepo) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetIssueCommentUrl

`func (o *PullRequestHeadRepo) GetIssueCommentUrl() string`

GetIssueCommentUrl returns the IssueCommentUrl field if non-nil, zero value otherwise.

### GetIssueCommentUrlOk

`func (o *PullRequestHeadRepo) GetIssueCommentUrlOk() (*string, bool)`

GetIssueCommentUrlOk returns a tuple with the IssueCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCommentUrl

`func (o *PullRequestHeadRepo) SetIssueCommentUrl(v string)`

SetIssueCommentUrl sets IssueCommentUrl field to given value.


### GetIssueEventsUrl

`func (o *PullRequestHeadRepo) GetIssueEventsUrl() string`

GetIssueEventsUrl returns the IssueEventsUrl field if non-nil, zero value otherwise.

### GetIssueEventsUrlOk

`func (o *PullRequestHeadRepo) GetIssueEventsUrlOk() (*string, bool)`

GetIssueEventsUrlOk returns a tuple with the IssueEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEventsUrl

`func (o *PullRequestHeadRepo) SetIssueEventsUrl(v string)`

SetIssueEventsUrl sets IssueEventsUrl field to given value.


### GetIssuesUrl

`func (o *PullRequestHeadRepo) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *PullRequestHeadRepo) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *PullRequestHeadRepo) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetKeysUrl

`func (o *PullRequestHeadRepo) GetKeysUrl() string`

GetKeysUrl returns the KeysUrl field if non-nil, zero value otherwise.

### GetKeysUrlOk

`func (o *PullRequestHeadRepo) GetKeysUrlOk() (*string, bool)`

GetKeysUrlOk returns a tuple with the KeysUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysUrl

`func (o *PullRequestHeadRepo) SetKeysUrl(v string)`

SetKeysUrl sets KeysUrl field to given value.


### GetLabelsUrl

`func (o *PullRequestHeadRepo) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *PullRequestHeadRepo) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *PullRequestHeadRepo) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetLanguagesUrl

`func (o *PullRequestHeadRepo) GetLanguagesUrl() string`

GetLanguagesUrl returns the LanguagesUrl field if non-nil, zero value otherwise.

### GetLanguagesUrlOk

`func (o *PullRequestHeadRepo) GetLanguagesUrlOk() (*string, bool)`

GetLanguagesUrlOk returns a tuple with the LanguagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesUrl

`func (o *PullRequestHeadRepo) SetLanguagesUrl(v string)`

SetLanguagesUrl sets LanguagesUrl field to given value.


### GetMergesUrl

`func (o *PullRequestHeadRepo) GetMergesUrl() string`

GetMergesUrl returns the MergesUrl field if non-nil, zero value otherwise.

### GetMergesUrlOk

`func (o *PullRequestHeadRepo) GetMergesUrlOk() (*string, bool)`

GetMergesUrlOk returns a tuple with the MergesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesUrl

`func (o *PullRequestHeadRepo) SetMergesUrl(v string)`

SetMergesUrl sets MergesUrl field to given value.


### GetMilestonesUrl

`func (o *PullRequestHeadRepo) GetMilestonesUrl() string`

GetMilestonesUrl returns the MilestonesUrl field if non-nil, zero value otherwise.

### GetMilestonesUrlOk

`func (o *PullRequestHeadRepo) GetMilestonesUrlOk() (*string, bool)`

GetMilestonesUrlOk returns a tuple with the MilestonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestonesUrl

`func (o *PullRequestHeadRepo) SetMilestonesUrl(v string)`

SetMilestonesUrl sets MilestonesUrl field to given value.


### GetName

`func (o *PullRequestHeadRepo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PullRequestHeadRepo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PullRequestHeadRepo) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationsUrl

`func (o *PullRequestHeadRepo) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *PullRequestHeadRepo) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *PullRequestHeadRepo) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.


### GetOwner

`func (o *PullRequestHeadRepo) GetOwner() PullRequestHeadRepoOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PullRequestHeadRepo) GetOwnerOk() (*PullRequestHeadRepoOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PullRequestHeadRepo) SetOwner(v PullRequestHeadRepoOwner)`

SetOwner sets Owner field to given value.


### GetPrivate

`func (o *PullRequestHeadRepo) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *PullRequestHeadRepo) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *PullRequestHeadRepo) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetPullsUrl

`func (o *PullRequestHeadRepo) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *PullRequestHeadRepo) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *PullRequestHeadRepo) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.


### GetReleasesUrl

`func (o *PullRequestHeadRepo) GetReleasesUrl() string`

GetReleasesUrl returns the ReleasesUrl field if non-nil, zero value otherwise.

### GetReleasesUrlOk

`func (o *PullRequestHeadRepo) GetReleasesUrlOk() (*string, bool)`

GetReleasesUrlOk returns a tuple with the ReleasesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesUrl

`func (o *PullRequestHeadRepo) SetReleasesUrl(v string)`

SetReleasesUrl sets ReleasesUrl field to given value.


### GetStargazersUrl

`func (o *PullRequestHeadRepo) GetStargazersUrl() string`

GetStargazersUrl returns the StargazersUrl field if non-nil, zero value otherwise.

### GetStargazersUrlOk

`func (o *PullRequestHeadRepo) GetStargazersUrlOk() (*string, bool)`

GetStargazersUrlOk returns a tuple with the StargazersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersUrl

`func (o *PullRequestHeadRepo) SetStargazersUrl(v string)`

SetStargazersUrl sets StargazersUrl field to given value.


### GetStatusesUrl

`func (o *PullRequestHeadRepo) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *PullRequestHeadRepo) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *PullRequestHeadRepo) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetSubscribersUrl

`func (o *PullRequestHeadRepo) GetSubscribersUrl() string`

GetSubscribersUrl returns the SubscribersUrl field if non-nil, zero value otherwise.

### GetSubscribersUrlOk

`func (o *PullRequestHeadRepo) GetSubscribersUrlOk() (*string, bool)`

GetSubscribersUrlOk returns a tuple with the SubscribersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersUrl

`func (o *PullRequestHeadRepo) SetSubscribersUrl(v string)`

SetSubscribersUrl sets SubscribersUrl field to given value.


### GetSubscriptionUrl

`func (o *PullRequestHeadRepo) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *PullRequestHeadRepo) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *PullRequestHeadRepo) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetTagsUrl

`func (o *PullRequestHeadRepo) GetTagsUrl() string`

GetTagsUrl returns the TagsUrl field if non-nil, zero value otherwise.

### GetTagsUrlOk

`func (o *PullRequestHeadRepo) GetTagsUrlOk() (*string, bool)`

GetTagsUrlOk returns a tuple with the TagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsUrl

`func (o *PullRequestHeadRepo) SetTagsUrl(v string)`

SetTagsUrl sets TagsUrl field to given value.


### GetTeamsUrl

`func (o *PullRequestHeadRepo) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *PullRequestHeadRepo) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *PullRequestHeadRepo) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetTreesUrl

`func (o *PullRequestHeadRepo) GetTreesUrl() string`

GetTreesUrl returns the TreesUrl field if non-nil, zero value otherwise.

### GetTreesUrlOk

`func (o *PullRequestHeadRepo) GetTreesUrlOk() (*string, bool)`

GetTreesUrlOk returns a tuple with the TreesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreesUrl

`func (o *PullRequestHeadRepo) SetTreesUrl(v string)`

SetTreesUrl sets TreesUrl field to given value.


### GetUrl

`func (o *PullRequestHeadRepo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestHeadRepo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestHeadRepo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCloneUrl

`func (o *PullRequestHeadRepo) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *PullRequestHeadRepo) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *PullRequestHeadRepo) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetDefaultBranch

`func (o *PullRequestHeadRepo) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *PullRequestHeadRepo) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *PullRequestHeadRepo) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetForks

`func (o *PullRequestHeadRepo) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *PullRequestHeadRepo) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *PullRequestHeadRepo) SetForks(v int32)`

SetForks sets Forks field to given value.


### GetForksCount

`func (o *PullRequestHeadRepo) GetForksCount() int32`

GetForksCount returns the ForksCount field if non-nil, zero value otherwise.

### GetForksCountOk

`func (o *PullRequestHeadRepo) GetForksCountOk() (*int32, bool)`

GetForksCountOk returns a tuple with the ForksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksCount

`func (o *PullRequestHeadRepo) SetForksCount(v int32)`

SetForksCount sets ForksCount field to given value.


### GetGitUrl

`func (o *PullRequestHeadRepo) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *PullRequestHeadRepo) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *PullRequestHeadRepo) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### GetHasDownloads

`func (o *PullRequestHeadRepo) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *PullRequestHeadRepo) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *PullRequestHeadRepo) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.


### GetHasIssues

`func (o *PullRequestHeadRepo) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *PullRequestHeadRepo) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *PullRequestHeadRepo) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.


### GetHasProjects

`func (o *PullRequestHeadRepo) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *PullRequestHeadRepo) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *PullRequestHeadRepo) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.


### GetHasWiki

`func (o *PullRequestHeadRepo) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *PullRequestHeadRepo) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *PullRequestHeadRepo) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.


### GetHasPages

`func (o *PullRequestHeadRepo) GetHasPages() bool`

GetHasPages returns the HasPages field if non-nil, zero value otherwise.

### GetHasPagesOk

`func (o *PullRequestHeadRepo) GetHasPagesOk() (*bool, bool)`

GetHasPagesOk returns a tuple with the HasPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPages

`func (o *PullRequestHeadRepo) SetHasPages(v bool)`

SetHasPages sets HasPages field to given value.


### GetHomepage

`func (o *PullRequestHeadRepo) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *PullRequestHeadRepo) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *PullRequestHeadRepo) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.


### SetHomepageNil

`func (o *PullRequestHeadRepo) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *PullRequestHeadRepo) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil
### GetLanguage

`func (o *PullRequestHeadRepo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PullRequestHeadRepo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PullRequestHeadRepo) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### SetLanguageNil

`func (o *PullRequestHeadRepo) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *PullRequestHeadRepo) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetMasterBranch

`func (o *PullRequestHeadRepo) GetMasterBranch() string`

GetMasterBranch returns the MasterBranch field if non-nil, zero value otherwise.

### GetMasterBranchOk

`func (o *PullRequestHeadRepo) GetMasterBranchOk() (*string, bool)`

GetMasterBranchOk returns a tuple with the MasterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterBranch

`func (o *PullRequestHeadRepo) SetMasterBranch(v string)`

SetMasterBranch sets MasterBranch field to given value.

### HasMasterBranch

`func (o *PullRequestHeadRepo) HasMasterBranch() bool`

HasMasterBranch returns a boolean if a field has been set.

### GetArchived

`func (o *PullRequestHeadRepo) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PullRequestHeadRepo) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PullRequestHeadRepo) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetDisabled

`func (o *PullRequestHeadRepo) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PullRequestHeadRepo) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PullRequestHeadRepo) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetVisibility

`func (o *PullRequestHeadRepo) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PullRequestHeadRepo) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PullRequestHeadRepo) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PullRequestHeadRepo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMirrorUrl

`func (o *PullRequestHeadRepo) GetMirrorUrl() string`

GetMirrorUrl returns the MirrorUrl field if non-nil, zero value otherwise.

### GetMirrorUrlOk

`func (o *PullRequestHeadRepo) GetMirrorUrlOk() (*string, bool)`

GetMirrorUrlOk returns a tuple with the MirrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorUrl

`func (o *PullRequestHeadRepo) SetMirrorUrl(v string)`

SetMirrorUrl sets MirrorUrl field to given value.


### SetMirrorUrlNil

`func (o *PullRequestHeadRepo) SetMirrorUrlNil(b bool)`

 SetMirrorUrlNil sets the value for MirrorUrl to be an explicit nil

### UnsetMirrorUrl
`func (o *PullRequestHeadRepo) UnsetMirrorUrl()`

UnsetMirrorUrl ensures that no value is present for MirrorUrl, not even an explicit nil
### GetOpenIssues

`func (o *PullRequestHeadRepo) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *PullRequestHeadRepo) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *PullRequestHeadRepo) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetOpenIssuesCount

`func (o *PullRequestHeadRepo) GetOpenIssuesCount() int32`

GetOpenIssuesCount returns the OpenIssuesCount field if non-nil, zero value otherwise.

### GetOpenIssuesCountOk

`func (o *PullRequestHeadRepo) GetOpenIssuesCountOk() (*int32, bool)`

GetOpenIssuesCountOk returns a tuple with the OpenIssuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssuesCount

`func (o *PullRequestHeadRepo) SetOpenIssuesCount(v int32)`

SetOpenIssuesCount sets OpenIssuesCount field to given value.


### GetPermissions

`func (o *PullRequestHeadRepo) GetPermissions() FullRepositoryPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PullRequestHeadRepo) GetPermissionsOk() (*FullRepositoryPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PullRequestHeadRepo) SetPermissions(v FullRepositoryPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PullRequestHeadRepo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTempCloneToken

`func (o *PullRequestHeadRepo) GetTempCloneToken() string`

GetTempCloneToken returns the TempCloneToken field if non-nil, zero value otherwise.

### GetTempCloneTokenOk

`func (o *PullRequestHeadRepo) GetTempCloneTokenOk() (*string, bool)`

GetTempCloneTokenOk returns a tuple with the TempCloneToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempCloneToken

`func (o *PullRequestHeadRepo) SetTempCloneToken(v string)`

SetTempCloneToken sets TempCloneToken field to given value.

### HasTempCloneToken

`func (o *PullRequestHeadRepo) HasTempCloneToken() bool`

HasTempCloneToken returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *PullRequestHeadRepo) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *PullRequestHeadRepo) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *PullRequestHeadRepo) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *PullRequestHeadRepo) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *PullRequestHeadRepo) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *PullRequestHeadRepo) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *PullRequestHeadRepo) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *PullRequestHeadRepo) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *PullRequestHeadRepo) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *PullRequestHeadRepo) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *PullRequestHeadRepo) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *PullRequestHeadRepo) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetLicense

`func (o *PullRequestHeadRepo) GetLicense() PullRequestHeadRepoLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PullRequestHeadRepo) GetLicenseOk() (*PullRequestHeadRepoLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PullRequestHeadRepo) SetLicense(v PullRequestHeadRepoLicense)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *PullRequestHeadRepo) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *PullRequestHeadRepo) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetPushedAt

`func (o *PullRequestHeadRepo) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *PullRequestHeadRepo) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *PullRequestHeadRepo) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.


### GetSize

`func (o *PullRequestHeadRepo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PullRequestHeadRepo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PullRequestHeadRepo) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSshUrl

`func (o *PullRequestHeadRepo) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *PullRequestHeadRepo) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *PullRequestHeadRepo) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.


### GetStargazersCount

`func (o *PullRequestHeadRepo) GetStargazersCount() int32`

GetStargazersCount returns the StargazersCount field if non-nil, zero value otherwise.

### GetStargazersCountOk

`func (o *PullRequestHeadRepo) GetStargazersCountOk() (*int32, bool)`

GetStargazersCountOk returns a tuple with the StargazersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersCount

`func (o *PullRequestHeadRepo) SetStargazersCount(v int32)`

SetStargazersCount sets StargazersCount field to given value.


### GetSvnUrl

`func (o *PullRequestHeadRepo) GetSvnUrl() string`

GetSvnUrl returns the SvnUrl field if non-nil, zero value otherwise.

### GetSvnUrlOk

`func (o *PullRequestHeadRepo) GetSvnUrlOk() (*string, bool)`

GetSvnUrlOk returns a tuple with the SvnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnUrl

`func (o *PullRequestHeadRepo) SetSvnUrl(v string)`

SetSvnUrl sets SvnUrl field to given value.


### GetTopics

`func (o *PullRequestHeadRepo) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *PullRequestHeadRepo) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *PullRequestHeadRepo) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *PullRequestHeadRepo) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetWatchers

`func (o *PullRequestHeadRepo) GetWatchers() int32`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *PullRequestHeadRepo) GetWatchersOk() (*int32, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *PullRequestHeadRepo) SetWatchers(v int32)`

SetWatchers sets Watchers field to given value.


### GetWatchersCount

`func (o *PullRequestHeadRepo) GetWatchersCount() int32`

GetWatchersCount returns the WatchersCount field if non-nil, zero value otherwise.

### GetWatchersCountOk

`func (o *PullRequestHeadRepo) GetWatchersCountOk() (*int32, bool)`

GetWatchersCountOk returns a tuple with the WatchersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchersCount

`func (o *PullRequestHeadRepo) SetWatchersCount(v int32)`

SetWatchersCount sets WatchersCount field to given value.


### GetCreatedAt

`func (o *PullRequestHeadRepo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequestHeadRepo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequestHeadRepo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PullRequestHeadRepo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequestHeadRepo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequestHeadRepo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAllowForking

`func (o *PullRequestHeadRepo) GetAllowForking() bool`

GetAllowForking returns the AllowForking field if non-nil, zero value otherwise.

### GetAllowForkingOk

`func (o *PullRequestHeadRepo) GetAllowForkingOk() (*bool, bool)`

GetAllowForkingOk returns a tuple with the AllowForking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForking

`func (o *PullRequestHeadRepo) SetAllowForking(v bool)`

SetAllowForking sets AllowForking field to given value.

### HasAllowForking

`func (o *PullRequestHeadRepo) HasAllowForking() bool`

HasAllowForking returns a boolean if a field has been set.

### GetIsTemplate

`func (o *PullRequestHeadRepo) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *PullRequestHeadRepo) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *PullRequestHeadRepo) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *PullRequestHeadRepo) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetWebCommitSignoffRequired

`func (o *PullRequestHeadRepo) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *PullRequestHeadRepo) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *PullRequestHeadRepo) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *PullRequestHeadRepo) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


