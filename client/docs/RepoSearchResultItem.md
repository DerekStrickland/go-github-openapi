# RepoSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Name** | **string** |  | 
**FullName** | **string** |  | 
**Owner** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Private** | **bool** |  | 
**HtmlUrl** | **string** |  | 
**Description** | **NullableString** |  | 
**Fork** | **bool** |  | 
**Url** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**PushedAt** | **time.Time** |  | 
**Homepage** | **NullableString** |  | 
**Size** | **int32** |  | 
**StargazersCount** | **int32** |  | 
**WatchersCount** | **int32** |  | 
**Language** | **NullableString** |  | 
**ForksCount** | **int32** |  | 
**OpenIssuesCount** | **int32** |  | 
**MasterBranch** | Pointer to **string** |  | [optional] 
**DefaultBranch** | **string** |  | 
**Score** | **float32** |  | 
**ForksUrl** | **string** |  | 
**KeysUrl** | **string** |  | 
**CollaboratorsUrl** | **string** |  | 
**TeamsUrl** | **string** |  | 
**HooksUrl** | **string** |  | 
**IssueEventsUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**AssigneesUrl** | **string** |  | 
**BranchesUrl** | **string** |  | 
**TagsUrl** | **string** |  | 
**BlobsUrl** | **string** |  | 
**GitTagsUrl** | **string** |  | 
**GitRefsUrl** | **string** |  | 
**TreesUrl** | **string** |  | 
**StatusesUrl** | **string** |  | 
**LanguagesUrl** | **string** |  | 
**StargazersUrl** | **string** |  | 
**ContributorsUrl** | **string** |  | 
**SubscribersUrl** | **string** |  | 
**SubscriptionUrl** | **string** |  | 
**CommitsUrl** | **string** |  | 
**GitCommitsUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**IssueCommentUrl** | **string** |  | 
**ContentsUrl** | **string** |  | 
**CompareUrl** | **string** |  | 
**MergesUrl** | **string** |  | 
**ArchiveUrl** | **string** |  | 
**DownloadsUrl** | **string** |  | 
**IssuesUrl** | **string** |  | 
**PullsUrl** | **string** |  | 
**MilestonesUrl** | **string** |  | 
**NotificationsUrl** | **string** |  | 
**LabelsUrl** | **string** |  | 
**ReleasesUrl** | **string** |  | 
**DeploymentsUrl** | **string** |  | 
**GitUrl** | **string** |  | 
**SshUrl** | **string** |  | 
**CloneUrl** | **string** |  | 
**SvnUrl** | **string** |  | 
**Forks** | **int32** |  | 
**OpenIssues** | **int32** |  | 
**Watchers** | **int32** |  | 
**Topics** | Pointer to **[]string** |  | [optional] 
**MirrorUrl** | **NullableString** |  | 
**HasIssues** | **bool** |  | 
**HasProjects** | **bool** |  | 
**HasPages** | **bool** |  | 
**HasWiki** | **bool** |  | 
**HasDownloads** | **bool** |  | 
**Archived** | **bool** |  | 
**Disabled** | **bool** | Returns whether or not this repository disabled. | 
**Visibility** | Pointer to **string** | The repository visibility: public, private, or internal. | [optional] 
**License** | [**NullableNullableLicenseSimple**](NullableLicenseSimple.md) |  | 
**Permissions** | Pointer to [**FullRepositoryPermissions**](FullRepositoryPermissions.md) |  | [optional] 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 
**TempCloneToken** | Pointer to **string** |  | [optional] 
**AllowMergeCommit** | Pointer to **bool** |  | [optional] 
**AllowSquashMerge** | Pointer to **bool** |  | [optional] 
**AllowRebaseMerge** | Pointer to **bool** |  | [optional] 
**AllowAutoMerge** | Pointer to **bool** |  | [optional] 
**DeleteBranchOnMerge** | Pointer to **bool** |  | [optional] 
**AllowForking** | Pointer to **bool** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**WebCommitSignoffRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewRepoSearchResultItem

`func NewRepoSearchResultItem(id int32, nodeId string, name string, fullName string, owner NullableNullableSimpleUser, private bool, htmlUrl string, description NullableString, fork bool, url string, createdAt time.Time, updatedAt time.Time, pushedAt time.Time, homepage NullableString, size int32, stargazersCount int32, watchersCount int32, language NullableString, forksCount int32, openIssuesCount int32, defaultBranch string, score float32, forksUrl string, keysUrl string, collaboratorsUrl string, teamsUrl string, hooksUrl string, issueEventsUrl string, eventsUrl string, assigneesUrl string, branchesUrl string, tagsUrl string, blobsUrl string, gitTagsUrl string, gitRefsUrl string, treesUrl string, statusesUrl string, languagesUrl string, stargazersUrl string, contributorsUrl string, subscribersUrl string, subscriptionUrl string, commitsUrl string, gitCommitsUrl string, commentsUrl string, issueCommentUrl string, contentsUrl string, compareUrl string, mergesUrl string, archiveUrl string, downloadsUrl string, issuesUrl string, pullsUrl string, milestonesUrl string, notificationsUrl string, labelsUrl string, releasesUrl string, deploymentsUrl string, gitUrl string, sshUrl string, cloneUrl string, svnUrl string, forks int32, openIssues int32, watchers int32, mirrorUrl NullableString, hasIssues bool, hasProjects bool, hasPages bool, hasWiki bool, hasDownloads bool, archived bool, disabled bool, license NullableNullableLicenseSimple, ) *RepoSearchResultItem`

NewRepoSearchResultItem instantiates a new RepoSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepoSearchResultItemWithDefaults

`func NewRepoSearchResultItemWithDefaults() *RepoSearchResultItem`

NewRepoSearchResultItemWithDefaults instantiates a new RepoSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepoSearchResultItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepoSearchResultItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepoSearchResultItem) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *RepoSearchResultItem) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RepoSearchResultItem) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RepoSearchResultItem) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *RepoSearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepoSearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepoSearchResultItem) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *RepoSearchResultItem) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *RepoSearchResultItem) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *RepoSearchResultItem) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetOwner

`func (o *RepoSearchResultItem) GetOwner() NullableSimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RepoSearchResultItem) GetOwnerOk() (*NullableSimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RepoSearchResultItem) SetOwner(v NullableSimpleUser)`

SetOwner sets Owner field to given value.


### SetOwnerNil

`func (o *RepoSearchResultItem) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *RepoSearchResultItem) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPrivate

`func (o *RepoSearchResultItem) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *RepoSearchResultItem) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *RepoSearchResultItem) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetHtmlUrl

`func (o *RepoSearchResultItem) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *RepoSearchResultItem) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *RepoSearchResultItem) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDescription

`func (o *RepoSearchResultItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepoSearchResultItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepoSearchResultItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *RepoSearchResultItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RepoSearchResultItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFork

`func (o *RepoSearchResultItem) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *RepoSearchResultItem) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *RepoSearchResultItem) SetFork(v bool)`

SetFork sets Fork field to given value.


### GetUrl

`func (o *RepoSearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepoSearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepoSearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *RepoSearchResultItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepoSearchResultItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepoSearchResultItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RepoSearchResultItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepoSearchResultItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepoSearchResultItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPushedAt

`func (o *RepoSearchResultItem) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *RepoSearchResultItem) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *RepoSearchResultItem) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.


### GetHomepage

`func (o *RepoSearchResultItem) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *RepoSearchResultItem) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *RepoSearchResultItem) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.


### SetHomepageNil

`func (o *RepoSearchResultItem) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *RepoSearchResultItem) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil
### GetSize

`func (o *RepoSearchResultItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RepoSearchResultItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RepoSearchResultItem) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStargazersCount

`func (o *RepoSearchResultItem) GetStargazersCount() int32`

GetStargazersCount returns the StargazersCount field if non-nil, zero value otherwise.

### GetStargazersCountOk

`func (o *RepoSearchResultItem) GetStargazersCountOk() (*int32, bool)`

GetStargazersCountOk returns a tuple with the StargazersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersCount

`func (o *RepoSearchResultItem) SetStargazersCount(v int32)`

SetStargazersCount sets StargazersCount field to given value.


### GetWatchersCount

`func (o *RepoSearchResultItem) GetWatchersCount() int32`

GetWatchersCount returns the WatchersCount field if non-nil, zero value otherwise.

### GetWatchersCountOk

`func (o *RepoSearchResultItem) GetWatchersCountOk() (*int32, bool)`

GetWatchersCountOk returns a tuple with the WatchersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchersCount

`func (o *RepoSearchResultItem) SetWatchersCount(v int32)`

SetWatchersCount sets WatchersCount field to given value.


### GetLanguage

`func (o *RepoSearchResultItem) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RepoSearchResultItem) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RepoSearchResultItem) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### SetLanguageNil

`func (o *RepoSearchResultItem) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *RepoSearchResultItem) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetForksCount

`func (o *RepoSearchResultItem) GetForksCount() int32`

GetForksCount returns the ForksCount field if non-nil, zero value otherwise.

### GetForksCountOk

`func (o *RepoSearchResultItem) GetForksCountOk() (*int32, bool)`

GetForksCountOk returns a tuple with the ForksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksCount

`func (o *RepoSearchResultItem) SetForksCount(v int32)`

SetForksCount sets ForksCount field to given value.


### GetOpenIssuesCount

`func (o *RepoSearchResultItem) GetOpenIssuesCount() int32`

GetOpenIssuesCount returns the OpenIssuesCount field if non-nil, zero value otherwise.

### GetOpenIssuesCountOk

`func (o *RepoSearchResultItem) GetOpenIssuesCountOk() (*int32, bool)`

GetOpenIssuesCountOk returns a tuple with the OpenIssuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssuesCount

`func (o *RepoSearchResultItem) SetOpenIssuesCount(v int32)`

SetOpenIssuesCount sets OpenIssuesCount field to given value.


### GetMasterBranch

`func (o *RepoSearchResultItem) GetMasterBranch() string`

GetMasterBranch returns the MasterBranch field if non-nil, zero value otherwise.

### GetMasterBranchOk

`func (o *RepoSearchResultItem) GetMasterBranchOk() (*string, bool)`

GetMasterBranchOk returns a tuple with the MasterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterBranch

`func (o *RepoSearchResultItem) SetMasterBranch(v string)`

SetMasterBranch sets MasterBranch field to given value.

### HasMasterBranch

`func (o *RepoSearchResultItem) HasMasterBranch() bool`

HasMasterBranch returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *RepoSearchResultItem) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *RepoSearchResultItem) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *RepoSearchResultItem) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetScore

`func (o *RepoSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RepoSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RepoSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetForksUrl

`func (o *RepoSearchResultItem) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *RepoSearchResultItem) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *RepoSearchResultItem) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.


### GetKeysUrl

`func (o *RepoSearchResultItem) GetKeysUrl() string`

GetKeysUrl returns the KeysUrl field if non-nil, zero value otherwise.

### GetKeysUrlOk

`func (o *RepoSearchResultItem) GetKeysUrlOk() (*string, bool)`

GetKeysUrlOk returns a tuple with the KeysUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysUrl

`func (o *RepoSearchResultItem) SetKeysUrl(v string)`

SetKeysUrl sets KeysUrl field to given value.


### GetCollaboratorsUrl

`func (o *RepoSearchResultItem) GetCollaboratorsUrl() string`

GetCollaboratorsUrl returns the CollaboratorsUrl field if non-nil, zero value otherwise.

### GetCollaboratorsUrlOk

`func (o *RepoSearchResultItem) GetCollaboratorsUrlOk() (*string, bool)`

GetCollaboratorsUrlOk returns a tuple with the CollaboratorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorsUrl

`func (o *RepoSearchResultItem) SetCollaboratorsUrl(v string)`

SetCollaboratorsUrl sets CollaboratorsUrl field to given value.


### GetTeamsUrl

`func (o *RepoSearchResultItem) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *RepoSearchResultItem) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *RepoSearchResultItem) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetHooksUrl

`func (o *RepoSearchResultItem) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *RepoSearchResultItem) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *RepoSearchResultItem) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.


### GetIssueEventsUrl

`func (o *RepoSearchResultItem) GetIssueEventsUrl() string`

GetIssueEventsUrl returns the IssueEventsUrl field if non-nil, zero value otherwise.

### GetIssueEventsUrlOk

`func (o *RepoSearchResultItem) GetIssueEventsUrlOk() (*string, bool)`

GetIssueEventsUrlOk returns a tuple with the IssueEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEventsUrl

`func (o *RepoSearchResultItem) SetIssueEventsUrl(v string)`

SetIssueEventsUrl sets IssueEventsUrl field to given value.


### GetEventsUrl

`func (o *RepoSearchResultItem) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *RepoSearchResultItem) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *RepoSearchResultItem) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetAssigneesUrl

`func (o *RepoSearchResultItem) GetAssigneesUrl() string`

GetAssigneesUrl returns the AssigneesUrl field if non-nil, zero value otherwise.

### GetAssigneesUrlOk

`func (o *RepoSearchResultItem) GetAssigneesUrlOk() (*string, bool)`

GetAssigneesUrlOk returns a tuple with the AssigneesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesUrl

`func (o *RepoSearchResultItem) SetAssigneesUrl(v string)`

SetAssigneesUrl sets AssigneesUrl field to given value.


### GetBranchesUrl

`func (o *RepoSearchResultItem) GetBranchesUrl() string`

GetBranchesUrl returns the BranchesUrl field if non-nil, zero value otherwise.

### GetBranchesUrlOk

`func (o *RepoSearchResultItem) GetBranchesUrlOk() (*string, bool)`

GetBranchesUrlOk returns a tuple with the BranchesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesUrl

`func (o *RepoSearchResultItem) SetBranchesUrl(v string)`

SetBranchesUrl sets BranchesUrl field to given value.


### GetTagsUrl

`func (o *RepoSearchResultItem) GetTagsUrl() string`

GetTagsUrl returns the TagsUrl field if non-nil, zero value otherwise.

### GetTagsUrlOk

`func (o *RepoSearchResultItem) GetTagsUrlOk() (*string, bool)`

GetTagsUrlOk returns a tuple with the TagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsUrl

`func (o *RepoSearchResultItem) SetTagsUrl(v string)`

SetTagsUrl sets TagsUrl field to given value.


### GetBlobsUrl

`func (o *RepoSearchResultItem) GetBlobsUrl() string`

GetBlobsUrl returns the BlobsUrl field if non-nil, zero value otherwise.

### GetBlobsUrlOk

`func (o *RepoSearchResultItem) GetBlobsUrlOk() (*string, bool)`

GetBlobsUrlOk returns a tuple with the BlobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsUrl

`func (o *RepoSearchResultItem) SetBlobsUrl(v string)`

SetBlobsUrl sets BlobsUrl field to given value.


### GetGitTagsUrl

`func (o *RepoSearchResultItem) GetGitTagsUrl() string`

GetGitTagsUrl returns the GitTagsUrl field if non-nil, zero value otherwise.

### GetGitTagsUrlOk

`func (o *RepoSearchResultItem) GetGitTagsUrlOk() (*string, bool)`

GetGitTagsUrlOk returns a tuple with the GitTagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagsUrl

`func (o *RepoSearchResultItem) SetGitTagsUrl(v string)`

SetGitTagsUrl sets GitTagsUrl field to given value.


### GetGitRefsUrl

`func (o *RepoSearchResultItem) GetGitRefsUrl() string`

GetGitRefsUrl returns the GitRefsUrl field if non-nil, zero value otherwise.

### GetGitRefsUrlOk

`func (o *RepoSearchResultItem) GetGitRefsUrlOk() (*string, bool)`

GetGitRefsUrlOk returns a tuple with the GitRefsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefsUrl

`func (o *RepoSearchResultItem) SetGitRefsUrl(v string)`

SetGitRefsUrl sets GitRefsUrl field to given value.


### GetTreesUrl

`func (o *RepoSearchResultItem) GetTreesUrl() string`

GetTreesUrl returns the TreesUrl field if non-nil, zero value otherwise.

### GetTreesUrlOk

`func (o *RepoSearchResultItem) GetTreesUrlOk() (*string, bool)`

GetTreesUrlOk returns a tuple with the TreesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreesUrl

`func (o *RepoSearchResultItem) SetTreesUrl(v string)`

SetTreesUrl sets TreesUrl field to given value.


### GetStatusesUrl

`func (o *RepoSearchResultItem) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *RepoSearchResultItem) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *RepoSearchResultItem) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetLanguagesUrl

`func (o *RepoSearchResultItem) GetLanguagesUrl() string`

GetLanguagesUrl returns the LanguagesUrl field if non-nil, zero value otherwise.

### GetLanguagesUrlOk

`func (o *RepoSearchResultItem) GetLanguagesUrlOk() (*string, bool)`

GetLanguagesUrlOk returns a tuple with the LanguagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesUrl

`func (o *RepoSearchResultItem) SetLanguagesUrl(v string)`

SetLanguagesUrl sets LanguagesUrl field to given value.


### GetStargazersUrl

`func (o *RepoSearchResultItem) GetStargazersUrl() string`

GetStargazersUrl returns the StargazersUrl field if non-nil, zero value otherwise.

### GetStargazersUrlOk

`func (o *RepoSearchResultItem) GetStargazersUrlOk() (*string, bool)`

GetStargazersUrlOk returns a tuple with the StargazersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersUrl

`func (o *RepoSearchResultItem) SetStargazersUrl(v string)`

SetStargazersUrl sets StargazersUrl field to given value.


### GetContributorsUrl

`func (o *RepoSearchResultItem) GetContributorsUrl() string`

GetContributorsUrl returns the ContributorsUrl field if non-nil, zero value otherwise.

### GetContributorsUrlOk

`func (o *RepoSearchResultItem) GetContributorsUrlOk() (*string, bool)`

GetContributorsUrlOk returns a tuple with the ContributorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsUrl

`func (o *RepoSearchResultItem) SetContributorsUrl(v string)`

SetContributorsUrl sets ContributorsUrl field to given value.


### GetSubscribersUrl

`func (o *RepoSearchResultItem) GetSubscribersUrl() string`

GetSubscribersUrl returns the SubscribersUrl field if non-nil, zero value otherwise.

### GetSubscribersUrlOk

`func (o *RepoSearchResultItem) GetSubscribersUrlOk() (*string, bool)`

GetSubscribersUrlOk returns a tuple with the SubscribersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersUrl

`func (o *RepoSearchResultItem) SetSubscribersUrl(v string)`

SetSubscribersUrl sets SubscribersUrl field to given value.


### GetSubscriptionUrl

`func (o *RepoSearchResultItem) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *RepoSearchResultItem) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *RepoSearchResultItem) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetCommitsUrl

`func (o *RepoSearchResultItem) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *RepoSearchResultItem) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *RepoSearchResultItem) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetGitCommitsUrl

`func (o *RepoSearchResultItem) GetGitCommitsUrl() string`

GetGitCommitsUrl returns the GitCommitsUrl field if non-nil, zero value otherwise.

### GetGitCommitsUrlOk

`func (o *RepoSearchResultItem) GetGitCommitsUrlOk() (*string, bool)`

GetGitCommitsUrlOk returns a tuple with the GitCommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitsUrl

`func (o *RepoSearchResultItem) SetGitCommitsUrl(v string)`

SetGitCommitsUrl sets GitCommitsUrl field to given value.


### GetCommentsUrl

`func (o *RepoSearchResultItem) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *RepoSearchResultItem) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *RepoSearchResultItem) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetIssueCommentUrl

`func (o *RepoSearchResultItem) GetIssueCommentUrl() string`

GetIssueCommentUrl returns the IssueCommentUrl field if non-nil, zero value otherwise.

### GetIssueCommentUrlOk

`func (o *RepoSearchResultItem) GetIssueCommentUrlOk() (*string, bool)`

GetIssueCommentUrlOk returns a tuple with the IssueCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCommentUrl

`func (o *RepoSearchResultItem) SetIssueCommentUrl(v string)`

SetIssueCommentUrl sets IssueCommentUrl field to given value.


### GetContentsUrl

`func (o *RepoSearchResultItem) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *RepoSearchResultItem) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *RepoSearchResultItem) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetCompareUrl

`func (o *RepoSearchResultItem) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *RepoSearchResultItem) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *RepoSearchResultItem) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.


### GetMergesUrl

`func (o *RepoSearchResultItem) GetMergesUrl() string`

GetMergesUrl returns the MergesUrl field if non-nil, zero value otherwise.

### GetMergesUrlOk

`func (o *RepoSearchResultItem) GetMergesUrlOk() (*string, bool)`

GetMergesUrlOk returns a tuple with the MergesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesUrl

`func (o *RepoSearchResultItem) SetMergesUrl(v string)`

SetMergesUrl sets MergesUrl field to given value.


### GetArchiveUrl

`func (o *RepoSearchResultItem) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *RepoSearchResultItem) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *RepoSearchResultItem) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.


### GetDownloadsUrl

`func (o *RepoSearchResultItem) GetDownloadsUrl() string`

GetDownloadsUrl returns the DownloadsUrl field if non-nil, zero value otherwise.

### GetDownloadsUrlOk

`func (o *RepoSearchResultItem) GetDownloadsUrlOk() (*string, bool)`

GetDownloadsUrlOk returns a tuple with the DownloadsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsUrl

`func (o *RepoSearchResultItem) SetDownloadsUrl(v string)`

SetDownloadsUrl sets DownloadsUrl field to given value.


### GetIssuesUrl

`func (o *RepoSearchResultItem) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *RepoSearchResultItem) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *RepoSearchResultItem) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetPullsUrl

`func (o *RepoSearchResultItem) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *RepoSearchResultItem) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *RepoSearchResultItem) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.


### GetMilestonesUrl

`func (o *RepoSearchResultItem) GetMilestonesUrl() string`

GetMilestonesUrl returns the MilestonesUrl field if non-nil, zero value otherwise.

### GetMilestonesUrlOk

`func (o *RepoSearchResultItem) GetMilestonesUrlOk() (*string, bool)`

GetMilestonesUrlOk returns a tuple with the MilestonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestonesUrl

`func (o *RepoSearchResultItem) SetMilestonesUrl(v string)`

SetMilestonesUrl sets MilestonesUrl field to given value.


### GetNotificationsUrl

`func (o *RepoSearchResultItem) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *RepoSearchResultItem) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *RepoSearchResultItem) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.


### GetLabelsUrl

`func (o *RepoSearchResultItem) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *RepoSearchResultItem) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *RepoSearchResultItem) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetReleasesUrl

`func (o *RepoSearchResultItem) GetReleasesUrl() string`

GetReleasesUrl returns the ReleasesUrl field if non-nil, zero value otherwise.

### GetReleasesUrlOk

`func (o *RepoSearchResultItem) GetReleasesUrlOk() (*string, bool)`

GetReleasesUrlOk returns a tuple with the ReleasesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesUrl

`func (o *RepoSearchResultItem) SetReleasesUrl(v string)`

SetReleasesUrl sets ReleasesUrl field to given value.


### GetDeploymentsUrl

`func (o *RepoSearchResultItem) GetDeploymentsUrl() string`

GetDeploymentsUrl returns the DeploymentsUrl field if non-nil, zero value otherwise.

### GetDeploymentsUrlOk

`func (o *RepoSearchResultItem) GetDeploymentsUrlOk() (*string, bool)`

GetDeploymentsUrlOk returns a tuple with the DeploymentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsUrl

`func (o *RepoSearchResultItem) SetDeploymentsUrl(v string)`

SetDeploymentsUrl sets DeploymentsUrl field to given value.


### GetGitUrl

`func (o *RepoSearchResultItem) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *RepoSearchResultItem) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *RepoSearchResultItem) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### GetSshUrl

`func (o *RepoSearchResultItem) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *RepoSearchResultItem) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *RepoSearchResultItem) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.


### GetCloneUrl

`func (o *RepoSearchResultItem) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *RepoSearchResultItem) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *RepoSearchResultItem) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetSvnUrl

`func (o *RepoSearchResultItem) GetSvnUrl() string`

GetSvnUrl returns the SvnUrl field if non-nil, zero value otherwise.

### GetSvnUrlOk

`func (o *RepoSearchResultItem) GetSvnUrlOk() (*string, bool)`

GetSvnUrlOk returns a tuple with the SvnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnUrl

`func (o *RepoSearchResultItem) SetSvnUrl(v string)`

SetSvnUrl sets SvnUrl field to given value.


### GetForks

`func (o *RepoSearchResultItem) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *RepoSearchResultItem) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *RepoSearchResultItem) SetForks(v int32)`

SetForks sets Forks field to given value.


### GetOpenIssues

`func (o *RepoSearchResultItem) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *RepoSearchResultItem) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *RepoSearchResultItem) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetWatchers

`func (o *RepoSearchResultItem) GetWatchers() int32`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *RepoSearchResultItem) GetWatchersOk() (*int32, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *RepoSearchResultItem) SetWatchers(v int32)`

SetWatchers sets Watchers field to given value.


### GetTopics

`func (o *RepoSearchResultItem) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *RepoSearchResultItem) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *RepoSearchResultItem) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *RepoSearchResultItem) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetMirrorUrl

`func (o *RepoSearchResultItem) GetMirrorUrl() string`

GetMirrorUrl returns the MirrorUrl field if non-nil, zero value otherwise.

### GetMirrorUrlOk

`func (o *RepoSearchResultItem) GetMirrorUrlOk() (*string, bool)`

GetMirrorUrlOk returns a tuple with the MirrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorUrl

`func (o *RepoSearchResultItem) SetMirrorUrl(v string)`

SetMirrorUrl sets MirrorUrl field to given value.


### SetMirrorUrlNil

`func (o *RepoSearchResultItem) SetMirrorUrlNil(b bool)`

 SetMirrorUrlNil sets the value for MirrorUrl to be an explicit nil

### UnsetMirrorUrl
`func (o *RepoSearchResultItem) UnsetMirrorUrl()`

UnsetMirrorUrl ensures that no value is present for MirrorUrl, not even an explicit nil
### GetHasIssues

`func (o *RepoSearchResultItem) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *RepoSearchResultItem) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *RepoSearchResultItem) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.


### GetHasProjects

`func (o *RepoSearchResultItem) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *RepoSearchResultItem) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *RepoSearchResultItem) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.


### GetHasPages

`func (o *RepoSearchResultItem) GetHasPages() bool`

GetHasPages returns the HasPages field if non-nil, zero value otherwise.

### GetHasPagesOk

`func (o *RepoSearchResultItem) GetHasPagesOk() (*bool, bool)`

GetHasPagesOk returns a tuple with the HasPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPages

`func (o *RepoSearchResultItem) SetHasPages(v bool)`

SetHasPages sets HasPages field to given value.


### GetHasWiki

`func (o *RepoSearchResultItem) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *RepoSearchResultItem) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *RepoSearchResultItem) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.


### GetHasDownloads

`func (o *RepoSearchResultItem) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *RepoSearchResultItem) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *RepoSearchResultItem) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.


### GetArchived

`func (o *RepoSearchResultItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *RepoSearchResultItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *RepoSearchResultItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetDisabled

`func (o *RepoSearchResultItem) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RepoSearchResultItem) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RepoSearchResultItem) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetVisibility

`func (o *RepoSearchResultItem) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RepoSearchResultItem) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RepoSearchResultItem) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *RepoSearchResultItem) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLicense

`func (o *RepoSearchResultItem) GetLicense() NullableLicenseSimple`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *RepoSearchResultItem) GetLicenseOk() (*NullableLicenseSimple, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *RepoSearchResultItem) SetLicense(v NullableLicenseSimple)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *RepoSearchResultItem) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *RepoSearchResultItem) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetPermissions

`func (o *RepoSearchResultItem) GetPermissions() FullRepositoryPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RepoSearchResultItem) GetPermissionsOk() (*FullRepositoryPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RepoSearchResultItem) SetPermissions(v FullRepositoryPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RepoSearchResultItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTextMatches

`func (o *RepoSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *RepoSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *RepoSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *RepoSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.

### GetTempCloneToken

`func (o *RepoSearchResultItem) GetTempCloneToken() string`

GetTempCloneToken returns the TempCloneToken field if non-nil, zero value otherwise.

### GetTempCloneTokenOk

`func (o *RepoSearchResultItem) GetTempCloneTokenOk() (*string, bool)`

GetTempCloneTokenOk returns a tuple with the TempCloneToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempCloneToken

`func (o *RepoSearchResultItem) SetTempCloneToken(v string)`

SetTempCloneToken sets TempCloneToken field to given value.

### HasTempCloneToken

`func (o *RepoSearchResultItem) HasTempCloneToken() bool`

HasTempCloneToken returns a boolean if a field has been set.

### GetAllowMergeCommit

`func (o *RepoSearchResultItem) GetAllowMergeCommit() bool`

GetAllowMergeCommit returns the AllowMergeCommit field if non-nil, zero value otherwise.

### GetAllowMergeCommitOk

`func (o *RepoSearchResultItem) GetAllowMergeCommitOk() (*bool, bool)`

GetAllowMergeCommitOk returns a tuple with the AllowMergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMergeCommit

`func (o *RepoSearchResultItem) SetAllowMergeCommit(v bool)`

SetAllowMergeCommit sets AllowMergeCommit field to given value.

### HasAllowMergeCommit

`func (o *RepoSearchResultItem) HasAllowMergeCommit() bool`

HasAllowMergeCommit returns a boolean if a field has been set.

### GetAllowSquashMerge

`func (o *RepoSearchResultItem) GetAllowSquashMerge() bool`

GetAllowSquashMerge returns the AllowSquashMerge field if non-nil, zero value otherwise.

### GetAllowSquashMergeOk

`func (o *RepoSearchResultItem) GetAllowSquashMergeOk() (*bool, bool)`

GetAllowSquashMergeOk returns a tuple with the AllowSquashMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSquashMerge

`func (o *RepoSearchResultItem) SetAllowSquashMerge(v bool)`

SetAllowSquashMerge sets AllowSquashMerge field to given value.

### HasAllowSquashMerge

`func (o *RepoSearchResultItem) HasAllowSquashMerge() bool`

HasAllowSquashMerge returns a boolean if a field has been set.

### GetAllowRebaseMerge

`func (o *RepoSearchResultItem) GetAllowRebaseMerge() bool`

GetAllowRebaseMerge returns the AllowRebaseMerge field if non-nil, zero value otherwise.

### GetAllowRebaseMergeOk

`func (o *RepoSearchResultItem) GetAllowRebaseMergeOk() (*bool, bool)`

GetAllowRebaseMergeOk returns a tuple with the AllowRebaseMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRebaseMerge

`func (o *RepoSearchResultItem) SetAllowRebaseMerge(v bool)`

SetAllowRebaseMerge sets AllowRebaseMerge field to given value.

### HasAllowRebaseMerge

`func (o *RepoSearchResultItem) HasAllowRebaseMerge() bool`

HasAllowRebaseMerge returns a boolean if a field has been set.

### GetAllowAutoMerge

`func (o *RepoSearchResultItem) GetAllowAutoMerge() bool`

GetAllowAutoMerge returns the AllowAutoMerge field if non-nil, zero value otherwise.

### GetAllowAutoMergeOk

`func (o *RepoSearchResultItem) GetAllowAutoMergeOk() (*bool, bool)`

GetAllowAutoMergeOk returns a tuple with the AllowAutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoMerge

`func (o *RepoSearchResultItem) SetAllowAutoMerge(v bool)`

SetAllowAutoMerge sets AllowAutoMerge field to given value.

### HasAllowAutoMerge

`func (o *RepoSearchResultItem) HasAllowAutoMerge() bool`

HasAllowAutoMerge returns a boolean if a field has been set.

### GetDeleteBranchOnMerge

`func (o *RepoSearchResultItem) GetDeleteBranchOnMerge() bool`

GetDeleteBranchOnMerge returns the DeleteBranchOnMerge field if non-nil, zero value otherwise.

### GetDeleteBranchOnMergeOk

`func (o *RepoSearchResultItem) GetDeleteBranchOnMergeOk() (*bool, bool)`

GetDeleteBranchOnMergeOk returns a tuple with the DeleteBranchOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBranchOnMerge

`func (o *RepoSearchResultItem) SetDeleteBranchOnMerge(v bool)`

SetDeleteBranchOnMerge sets DeleteBranchOnMerge field to given value.

### HasDeleteBranchOnMerge

`func (o *RepoSearchResultItem) HasDeleteBranchOnMerge() bool`

HasDeleteBranchOnMerge returns a boolean if a field has been set.

### GetAllowForking

`func (o *RepoSearchResultItem) GetAllowForking() bool`

GetAllowForking returns the AllowForking field if non-nil, zero value otherwise.

### GetAllowForkingOk

`func (o *RepoSearchResultItem) GetAllowForkingOk() (*bool, bool)`

GetAllowForkingOk returns a tuple with the AllowForking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForking

`func (o *RepoSearchResultItem) SetAllowForking(v bool)`

SetAllowForking sets AllowForking field to given value.

### HasAllowForking

`func (o *RepoSearchResultItem) HasAllowForking() bool`

HasAllowForking returns a boolean if a field has been set.

### GetIsTemplate

`func (o *RepoSearchResultItem) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *RepoSearchResultItem) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *RepoSearchResultItem) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *RepoSearchResultItem) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetWebCommitSignoffRequired

`func (o *RepoSearchResultItem) GetWebCommitSignoffRequired() bool`

GetWebCommitSignoffRequired returns the WebCommitSignoffRequired field if non-nil, zero value otherwise.

### GetWebCommitSignoffRequiredOk

`func (o *RepoSearchResultItem) GetWebCommitSignoffRequiredOk() (*bool, bool)`

GetWebCommitSignoffRequiredOk returns a tuple with the WebCommitSignoffRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebCommitSignoffRequired

`func (o *RepoSearchResultItem) SetWebCommitSignoffRequired(v bool)`

SetWebCommitSignoffRequired sets WebCommitSignoffRequired field to given value.

### HasWebCommitSignoffRequired

`func (o *RepoSearchResultItem) HasWebCommitSignoffRequired() bool`

HasWebCommitSignoffRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


