# SimpleRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | A unique identifier of the repository. | 
**NodeId** | **string** | The GraphQL identifier of the repository. | 
**Name** | **string** | The name of the repository. | 
**FullName** | **string** | The full, globally unique, name of the repository. | 
**Owner** | [**SimpleUser**](SimpleUser.md) |  | 
**Private** | **bool** | Whether the repository is private. | 
**HtmlUrl** | **string** | The URL to view the repository on GitHub.com. | 
**Description** | **NullableString** | The repository description. | 
**Fork** | **bool** | Whether the repository is a fork. | 
**Url** | **string** | The URL to get more information about the repository from the GitHub API. | 
**ArchiveUrl** | **string** | A template for the API URL to download the repository as an archive. | 
**AssigneesUrl** | **string** | A template for the API URL to list the available assignees for issues in the repository. | 
**BlobsUrl** | **string** | A template for the API URL to create or retrieve a raw Git blob in the repository. | 
**BranchesUrl** | **string** | A template for the API URL to get information about branches in the repository. | 
**CollaboratorsUrl** | **string** | A template for the API URL to get information about collaborators of the repository. | 
**CommentsUrl** | **string** | A template for the API URL to get information about comments on the repository. | 
**CommitsUrl** | **string** | A template for the API URL to get information about commits on the repository. | 
**CompareUrl** | **string** | A template for the API URL to compare two commits or refs. | 
**ContentsUrl** | **string** | A template for the API URL to get the contents of the repository. | 
**ContributorsUrl** | **string** | A template for the API URL to list the contributors to the repository. | 
**DeploymentsUrl** | **string** | The API URL to list the deployments of the repository. | 
**DownloadsUrl** | **string** | The API URL to list the downloads on the repository. | 
**EventsUrl** | **string** | The API URL to list the events of the repository. | 
**ForksUrl** | **string** | The API URL to list the forks of the repository. | 
**GitCommitsUrl** | **string** | A template for the API URL to get information about Git commits of the repository. | 
**GitRefsUrl** | **string** | A template for the API URL to get information about Git refs of the repository. | 
**GitTagsUrl** | **string** | A template for the API URL to get information about Git tags of the repository. | 
**IssueCommentUrl** | **string** | A template for the API URL to get information about issue comments on the repository. | 
**IssueEventsUrl** | **string** | A template for the API URL to get information about issue events on the repository. | 
**IssuesUrl** | **string** | A template for the API URL to get information about issues on the repository. | 
**KeysUrl** | **string** | A template for the API URL to get information about deploy keys on the repository. | 
**LabelsUrl** | **string** | A template for the API URL to get information about labels of the repository. | 
**LanguagesUrl** | **string** | The API URL to get information about the languages of the repository. | 
**MergesUrl** | **string** | The API URL to merge branches in the repository. | 
**MilestonesUrl** | **string** | A template for the API URL to get information about milestones of the repository. | 
**NotificationsUrl** | **string** | A template for the API URL to get information about notifications on the repository. | 
**PullsUrl** | **string** | A template for the API URL to get information about pull requests on the repository. | 
**ReleasesUrl** | **string** | A template for the API URL to get information about releases on the repository. | 
**StargazersUrl** | **string** | The API URL to list the stargazers on the repository. | 
**StatusesUrl** | **string** | A template for the API URL to get information about statuses of a commit. | 
**SubscribersUrl** | **string** | The API URL to list the subscribers on the repository. | 
**SubscriptionUrl** | **string** | The API URL to subscribe to notifications for this repository. | 
**TagsUrl** | **string** | The API URL to get information about tags on the repository. | 
**TeamsUrl** | **string** | The API URL to list the teams on the repository. | 
**TreesUrl** | **string** | A template for the API URL to create or retrieve a raw Git tree of the repository. | 
**HooksUrl** | **string** | The API URL to list the hooks on the repository. | 

## Methods

### NewSimpleRepository

`func NewSimpleRepository(id int32, nodeId string, name string, fullName string, owner SimpleUser, private bool, htmlUrl string, description NullableString, fork bool, url string, archiveUrl string, assigneesUrl string, blobsUrl string, branchesUrl string, collaboratorsUrl string, commentsUrl string, commitsUrl string, compareUrl string, contentsUrl string, contributorsUrl string, deploymentsUrl string, downloadsUrl string, eventsUrl string, forksUrl string, gitCommitsUrl string, gitRefsUrl string, gitTagsUrl string, issueCommentUrl string, issueEventsUrl string, issuesUrl string, keysUrl string, labelsUrl string, languagesUrl string, mergesUrl string, milestonesUrl string, notificationsUrl string, pullsUrl string, releasesUrl string, stargazersUrl string, statusesUrl string, subscribersUrl string, subscriptionUrl string, tagsUrl string, teamsUrl string, treesUrl string, hooksUrl string, ) *SimpleRepository`

NewSimpleRepository instantiates a new SimpleRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleRepositoryWithDefaults

`func NewSimpleRepositoryWithDefaults() *SimpleRepository`

NewSimpleRepositoryWithDefaults instantiates a new SimpleRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleRepository) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleRepository) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleRepository) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *SimpleRepository) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SimpleRepository) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SimpleRepository) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetName

`func (o *SimpleRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleRepository) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *SimpleRepository) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SimpleRepository) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SimpleRepository) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetOwner

`func (o *SimpleRepository) GetOwner() SimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SimpleRepository) GetOwnerOk() (*SimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SimpleRepository) SetOwner(v SimpleUser)`

SetOwner sets Owner field to given value.


### GetPrivate

`func (o *SimpleRepository) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SimpleRepository) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SimpleRepository) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetHtmlUrl

`func (o *SimpleRepository) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *SimpleRepository) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *SimpleRepository) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDescription

`func (o *SimpleRepository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleRepository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleRepository) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SimpleRepository) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SimpleRepository) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFork

`func (o *SimpleRepository) GetFork() bool`

GetFork returns the Fork field if non-nil, zero value otherwise.

### GetForkOk

`func (o *SimpleRepository) GetForkOk() (*bool, bool)`

GetForkOk returns a tuple with the Fork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFork

`func (o *SimpleRepository) SetFork(v bool)`

SetFork sets Fork field to given value.


### GetUrl

`func (o *SimpleRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimpleRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimpleRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetArchiveUrl

`func (o *SimpleRepository) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *SimpleRepository) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *SimpleRepository) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.


### GetAssigneesUrl

`func (o *SimpleRepository) GetAssigneesUrl() string`

GetAssigneesUrl returns the AssigneesUrl field if non-nil, zero value otherwise.

### GetAssigneesUrlOk

`func (o *SimpleRepository) GetAssigneesUrlOk() (*string, bool)`

GetAssigneesUrlOk returns a tuple with the AssigneesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesUrl

`func (o *SimpleRepository) SetAssigneesUrl(v string)`

SetAssigneesUrl sets AssigneesUrl field to given value.


### GetBlobsUrl

`func (o *SimpleRepository) GetBlobsUrl() string`

GetBlobsUrl returns the BlobsUrl field if non-nil, zero value otherwise.

### GetBlobsUrlOk

`func (o *SimpleRepository) GetBlobsUrlOk() (*string, bool)`

GetBlobsUrlOk returns a tuple with the BlobsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsUrl

`func (o *SimpleRepository) SetBlobsUrl(v string)`

SetBlobsUrl sets BlobsUrl field to given value.


### GetBranchesUrl

`func (o *SimpleRepository) GetBranchesUrl() string`

GetBranchesUrl returns the BranchesUrl field if non-nil, zero value otherwise.

### GetBranchesUrlOk

`func (o *SimpleRepository) GetBranchesUrlOk() (*string, bool)`

GetBranchesUrlOk returns a tuple with the BranchesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchesUrl

`func (o *SimpleRepository) SetBranchesUrl(v string)`

SetBranchesUrl sets BranchesUrl field to given value.


### GetCollaboratorsUrl

`func (o *SimpleRepository) GetCollaboratorsUrl() string`

GetCollaboratorsUrl returns the CollaboratorsUrl field if non-nil, zero value otherwise.

### GetCollaboratorsUrlOk

`func (o *SimpleRepository) GetCollaboratorsUrlOk() (*string, bool)`

GetCollaboratorsUrlOk returns a tuple with the CollaboratorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboratorsUrl

`func (o *SimpleRepository) SetCollaboratorsUrl(v string)`

SetCollaboratorsUrl sets CollaboratorsUrl field to given value.


### GetCommentsUrl

`func (o *SimpleRepository) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *SimpleRepository) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *SimpleRepository) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCommitsUrl

`func (o *SimpleRepository) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *SimpleRepository) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *SimpleRepository) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetCompareUrl

`func (o *SimpleRepository) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *SimpleRepository) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *SimpleRepository) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.


### GetContentsUrl

`func (o *SimpleRepository) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *SimpleRepository) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *SimpleRepository) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.


### GetContributorsUrl

`func (o *SimpleRepository) GetContributorsUrl() string`

GetContributorsUrl returns the ContributorsUrl field if non-nil, zero value otherwise.

### GetContributorsUrlOk

`func (o *SimpleRepository) GetContributorsUrlOk() (*string, bool)`

GetContributorsUrlOk returns a tuple with the ContributorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorsUrl

`func (o *SimpleRepository) SetContributorsUrl(v string)`

SetContributorsUrl sets ContributorsUrl field to given value.


### GetDeploymentsUrl

`func (o *SimpleRepository) GetDeploymentsUrl() string`

GetDeploymentsUrl returns the DeploymentsUrl field if non-nil, zero value otherwise.

### GetDeploymentsUrlOk

`func (o *SimpleRepository) GetDeploymentsUrlOk() (*string, bool)`

GetDeploymentsUrlOk returns a tuple with the DeploymentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsUrl

`func (o *SimpleRepository) SetDeploymentsUrl(v string)`

SetDeploymentsUrl sets DeploymentsUrl field to given value.


### GetDownloadsUrl

`func (o *SimpleRepository) GetDownloadsUrl() string`

GetDownloadsUrl returns the DownloadsUrl field if non-nil, zero value otherwise.

### GetDownloadsUrlOk

`func (o *SimpleRepository) GetDownloadsUrlOk() (*string, bool)`

GetDownloadsUrlOk returns a tuple with the DownloadsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsUrl

`func (o *SimpleRepository) SetDownloadsUrl(v string)`

SetDownloadsUrl sets DownloadsUrl field to given value.


### GetEventsUrl

`func (o *SimpleRepository) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *SimpleRepository) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *SimpleRepository) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetForksUrl

`func (o *SimpleRepository) GetForksUrl() string`

GetForksUrl returns the ForksUrl field if non-nil, zero value otherwise.

### GetForksUrlOk

`func (o *SimpleRepository) GetForksUrlOk() (*string, bool)`

GetForksUrlOk returns a tuple with the ForksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForksUrl

`func (o *SimpleRepository) SetForksUrl(v string)`

SetForksUrl sets ForksUrl field to given value.


### GetGitCommitsUrl

`func (o *SimpleRepository) GetGitCommitsUrl() string`

GetGitCommitsUrl returns the GitCommitsUrl field if non-nil, zero value otherwise.

### GetGitCommitsUrlOk

`func (o *SimpleRepository) GetGitCommitsUrlOk() (*string, bool)`

GetGitCommitsUrlOk returns a tuple with the GitCommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitsUrl

`func (o *SimpleRepository) SetGitCommitsUrl(v string)`

SetGitCommitsUrl sets GitCommitsUrl field to given value.


### GetGitRefsUrl

`func (o *SimpleRepository) GetGitRefsUrl() string`

GetGitRefsUrl returns the GitRefsUrl field if non-nil, zero value otherwise.

### GetGitRefsUrlOk

`func (o *SimpleRepository) GetGitRefsUrlOk() (*string, bool)`

GetGitRefsUrlOk returns a tuple with the GitRefsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRefsUrl

`func (o *SimpleRepository) SetGitRefsUrl(v string)`

SetGitRefsUrl sets GitRefsUrl field to given value.


### GetGitTagsUrl

`func (o *SimpleRepository) GetGitTagsUrl() string`

GetGitTagsUrl returns the GitTagsUrl field if non-nil, zero value otherwise.

### GetGitTagsUrlOk

`func (o *SimpleRepository) GetGitTagsUrlOk() (*string, bool)`

GetGitTagsUrlOk returns a tuple with the GitTagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagsUrl

`func (o *SimpleRepository) SetGitTagsUrl(v string)`

SetGitTagsUrl sets GitTagsUrl field to given value.


### GetIssueCommentUrl

`func (o *SimpleRepository) GetIssueCommentUrl() string`

GetIssueCommentUrl returns the IssueCommentUrl field if non-nil, zero value otherwise.

### GetIssueCommentUrlOk

`func (o *SimpleRepository) GetIssueCommentUrlOk() (*string, bool)`

GetIssueCommentUrlOk returns a tuple with the IssueCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCommentUrl

`func (o *SimpleRepository) SetIssueCommentUrl(v string)`

SetIssueCommentUrl sets IssueCommentUrl field to given value.


### GetIssueEventsUrl

`func (o *SimpleRepository) GetIssueEventsUrl() string`

GetIssueEventsUrl returns the IssueEventsUrl field if non-nil, zero value otherwise.

### GetIssueEventsUrlOk

`func (o *SimpleRepository) GetIssueEventsUrlOk() (*string, bool)`

GetIssueEventsUrlOk returns a tuple with the IssueEventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueEventsUrl

`func (o *SimpleRepository) SetIssueEventsUrl(v string)`

SetIssueEventsUrl sets IssueEventsUrl field to given value.


### GetIssuesUrl

`func (o *SimpleRepository) GetIssuesUrl() string`

GetIssuesUrl returns the IssuesUrl field if non-nil, zero value otherwise.

### GetIssuesUrlOk

`func (o *SimpleRepository) GetIssuesUrlOk() (*string, bool)`

GetIssuesUrlOk returns a tuple with the IssuesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesUrl

`func (o *SimpleRepository) SetIssuesUrl(v string)`

SetIssuesUrl sets IssuesUrl field to given value.


### GetKeysUrl

`func (o *SimpleRepository) GetKeysUrl() string`

GetKeysUrl returns the KeysUrl field if non-nil, zero value otherwise.

### GetKeysUrlOk

`func (o *SimpleRepository) GetKeysUrlOk() (*string, bool)`

GetKeysUrlOk returns a tuple with the KeysUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysUrl

`func (o *SimpleRepository) SetKeysUrl(v string)`

SetKeysUrl sets KeysUrl field to given value.


### GetLabelsUrl

`func (o *SimpleRepository) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *SimpleRepository) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *SimpleRepository) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetLanguagesUrl

`func (o *SimpleRepository) GetLanguagesUrl() string`

GetLanguagesUrl returns the LanguagesUrl field if non-nil, zero value otherwise.

### GetLanguagesUrlOk

`func (o *SimpleRepository) GetLanguagesUrlOk() (*string, bool)`

GetLanguagesUrlOk returns a tuple with the LanguagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesUrl

`func (o *SimpleRepository) SetLanguagesUrl(v string)`

SetLanguagesUrl sets LanguagesUrl field to given value.


### GetMergesUrl

`func (o *SimpleRepository) GetMergesUrl() string`

GetMergesUrl returns the MergesUrl field if non-nil, zero value otherwise.

### GetMergesUrlOk

`func (o *SimpleRepository) GetMergesUrlOk() (*string, bool)`

GetMergesUrlOk returns a tuple with the MergesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergesUrl

`func (o *SimpleRepository) SetMergesUrl(v string)`

SetMergesUrl sets MergesUrl field to given value.


### GetMilestonesUrl

`func (o *SimpleRepository) GetMilestonesUrl() string`

GetMilestonesUrl returns the MilestonesUrl field if non-nil, zero value otherwise.

### GetMilestonesUrlOk

`func (o *SimpleRepository) GetMilestonesUrlOk() (*string, bool)`

GetMilestonesUrlOk returns a tuple with the MilestonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestonesUrl

`func (o *SimpleRepository) SetMilestonesUrl(v string)`

SetMilestonesUrl sets MilestonesUrl field to given value.


### GetNotificationsUrl

`func (o *SimpleRepository) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *SimpleRepository) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *SimpleRepository) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.


### GetPullsUrl

`func (o *SimpleRepository) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *SimpleRepository) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *SimpleRepository) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.


### GetReleasesUrl

`func (o *SimpleRepository) GetReleasesUrl() string`

GetReleasesUrl returns the ReleasesUrl field if non-nil, zero value otherwise.

### GetReleasesUrlOk

`func (o *SimpleRepository) GetReleasesUrlOk() (*string, bool)`

GetReleasesUrlOk returns a tuple with the ReleasesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasesUrl

`func (o *SimpleRepository) SetReleasesUrl(v string)`

SetReleasesUrl sets ReleasesUrl field to given value.


### GetStargazersUrl

`func (o *SimpleRepository) GetStargazersUrl() string`

GetStargazersUrl returns the StargazersUrl field if non-nil, zero value otherwise.

### GetStargazersUrlOk

`func (o *SimpleRepository) GetStargazersUrlOk() (*string, bool)`

GetStargazersUrlOk returns a tuple with the StargazersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStargazersUrl

`func (o *SimpleRepository) SetStargazersUrl(v string)`

SetStargazersUrl sets StargazersUrl field to given value.


### GetStatusesUrl

`func (o *SimpleRepository) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *SimpleRepository) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *SimpleRepository) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetSubscribersUrl

`func (o *SimpleRepository) GetSubscribersUrl() string`

GetSubscribersUrl returns the SubscribersUrl field if non-nil, zero value otherwise.

### GetSubscribersUrlOk

`func (o *SimpleRepository) GetSubscribersUrlOk() (*string, bool)`

GetSubscribersUrlOk returns a tuple with the SubscribersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersUrl

`func (o *SimpleRepository) SetSubscribersUrl(v string)`

SetSubscribersUrl sets SubscribersUrl field to given value.


### GetSubscriptionUrl

`func (o *SimpleRepository) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *SimpleRepository) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *SimpleRepository) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetTagsUrl

`func (o *SimpleRepository) GetTagsUrl() string`

GetTagsUrl returns the TagsUrl field if non-nil, zero value otherwise.

### GetTagsUrlOk

`func (o *SimpleRepository) GetTagsUrlOk() (*string, bool)`

GetTagsUrlOk returns a tuple with the TagsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsUrl

`func (o *SimpleRepository) SetTagsUrl(v string)`

SetTagsUrl sets TagsUrl field to given value.


### GetTeamsUrl

`func (o *SimpleRepository) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *SimpleRepository) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *SimpleRepository) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetTreesUrl

`func (o *SimpleRepository) GetTreesUrl() string`

GetTreesUrl returns the TreesUrl field if non-nil, zero value otherwise.

### GetTreesUrlOk

`func (o *SimpleRepository) GetTreesUrlOk() (*string, bool)`

GetTreesUrlOk returns a tuple with the TreesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreesUrl

`func (o *SimpleRepository) SetTreesUrl(v string)`

SetTreesUrl sets TreesUrl field to given value.


### GetHooksUrl

`func (o *SimpleRepository) GetHooksUrl() string`

GetHooksUrl returns the HooksUrl field if non-nil, zero value otherwise.

### GetHooksUrlOk

`func (o *SimpleRepository) GetHooksUrlOk() (*string, bool)`

GetHooksUrlOk returns a tuple with the HooksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksUrl

`func (o *SimpleRepository) SetHooksUrl(v string)`

SetHooksUrl sets HooksUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


