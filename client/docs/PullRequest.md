# PullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**HtmlUrl** | **string** |  | 
**DiffUrl** | **string** |  | 
**PatchUrl** | **string** |  | 
**IssueUrl** | **string** |  | 
**CommitsUrl** | **string** |  | 
**ReviewCommentsUrl** | **string** |  | 
**ReviewCommentUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**StatusesUrl** | **string** |  | 
**Number** | **int32** | Number uniquely identifying the pull request within its repository. | 
**State** | **string** | State of this Pull Request. Either &#x60;open&#x60; or &#x60;closed&#x60;. | 
**Locked** | **bool** |  | 
**Title** | **string** | The title of the pull request. | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Body** | **NullableString** |  | 
**Labels** | [**[]PullRequestLabelsInner**](PullRequestLabelsInner.md) |  | 
**Milestone** | [**NullableNullableMilestone**](NullableMilestone.md) |  | 
**ActiveLockReason** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ClosedAt** | **NullableTime** |  | 
**MergedAt** | **NullableTime** |  | 
**MergeCommitSha** | **NullableString** |  | 
**Assignee** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Assignees** | Pointer to [**[]SimpleUser**](SimpleUser.md) |  | [optional] 
**RequestedReviewers** | Pointer to [**[]SimpleUser**](SimpleUser.md) |  | [optional] 
**RequestedTeams** | Pointer to [**[]TeamSimple**](TeamSimple.md) |  | [optional] 
**Head** | [**PullRequestHead**](PullRequestHead.md) |  | 
**Base** | [**PullRequestBase**](PullRequestBase.md) |  | 
**Links** | [**PullRequestSimpleLinks**](PullRequestSimpleLinks.md) |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**AutoMerge** | [**NullableAutoMerge**](AutoMerge.md) |  | 
**Draft** | Pointer to **bool** | Indicates whether or not the pull request is a draft. | [optional] 
**Merged** | **bool** |  | 
**Mergeable** | **NullableBool** |  | 
**Rebaseable** | Pointer to **NullableBool** |  | [optional] 
**MergeableState** | **string** |  | 
**MergedBy** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Comments** | **int32** |  | 
**ReviewComments** | **int32** |  | 
**MaintainerCanModify** | **bool** | Indicates whether maintainers can modify the pull request. | 
**Commits** | **int32** |  | 
**Additions** | **int32** |  | 
**Deletions** | **int32** |  | 
**ChangedFiles** | **int32** |  | 

## Methods

### NewPullRequest

`func NewPullRequest(url string, id int32, nodeId string, htmlUrl string, diffUrl string, patchUrl string, issueUrl string, commitsUrl string, reviewCommentsUrl string, reviewCommentUrl string, commentsUrl string, statusesUrl string, number int32, state string, locked bool, title string, user NullableNullableSimpleUser, body NullableString, labels []PullRequestLabelsInner, milestone NullableNullableMilestone, createdAt time.Time, updatedAt time.Time, closedAt NullableTime, mergedAt NullableTime, mergeCommitSha NullableString, assignee NullableNullableSimpleUser, head PullRequestHead, base PullRequestBase, links PullRequestSimpleLinks, authorAssociation AuthorAssociation, autoMerge NullableAutoMerge, merged bool, mergeable NullableBool, mergeableState string, mergedBy NullableNullableSimpleUser, comments int32, reviewComments int32, maintainerCanModify bool, commits int32, additions int32, deletions int32, changedFiles int32, ) *PullRequest`

NewPullRequest instantiates a new PullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestWithDefaults

`func NewPullRequestWithDefaults() *PullRequest`

NewPullRequestWithDefaults instantiates a new PullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *PullRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PullRequest) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PullRequest) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PullRequest) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHtmlUrl

`func (o *PullRequest) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PullRequest) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PullRequest) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDiffUrl

`func (o *PullRequest) GetDiffUrl() string`

GetDiffUrl returns the DiffUrl field if non-nil, zero value otherwise.

### GetDiffUrlOk

`func (o *PullRequest) GetDiffUrlOk() (*string, bool)`

GetDiffUrlOk returns a tuple with the DiffUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffUrl

`func (o *PullRequest) SetDiffUrl(v string)`

SetDiffUrl sets DiffUrl field to given value.


### GetPatchUrl

`func (o *PullRequest) GetPatchUrl() string`

GetPatchUrl returns the PatchUrl field if non-nil, zero value otherwise.

### GetPatchUrlOk

`func (o *PullRequest) GetPatchUrlOk() (*string, bool)`

GetPatchUrlOk returns a tuple with the PatchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchUrl

`func (o *PullRequest) SetPatchUrl(v string)`

SetPatchUrl sets PatchUrl field to given value.


### GetIssueUrl

`func (o *PullRequest) GetIssueUrl() string`

GetIssueUrl returns the IssueUrl field if non-nil, zero value otherwise.

### GetIssueUrlOk

`func (o *PullRequest) GetIssueUrlOk() (*string, bool)`

GetIssueUrlOk returns a tuple with the IssueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueUrl

`func (o *PullRequest) SetIssueUrl(v string)`

SetIssueUrl sets IssueUrl field to given value.


### GetCommitsUrl

`func (o *PullRequest) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *PullRequest) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *PullRequest) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetReviewCommentsUrl

`func (o *PullRequest) GetReviewCommentsUrl() string`

GetReviewCommentsUrl returns the ReviewCommentsUrl field if non-nil, zero value otherwise.

### GetReviewCommentsUrlOk

`func (o *PullRequest) GetReviewCommentsUrlOk() (*string, bool)`

GetReviewCommentsUrlOk returns a tuple with the ReviewCommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCommentsUrl

`func (o *PullRequest) SetReviewCommentsUrl(v string)`

SetReviewCommentsUrl sets ReviewCommentsUrl field to given value.


### GetReviewCommentUrl

`func (o *PullRequest) GetReviewCommentUrl() string`

GetReviewCommentUrl returns the ReviewCommentUrl field if non-nil, zero value otherwise.

### GetReviewCommentUrlOk

`func (o *PullRequest) GetReviewCommentUrlOk() (*string, bool)`

GetReviewCommentUrlOk returns a tuple with the ReviewCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCommentUrl

`func (o *PullRequest) SetReviewCommentUrl(v string)`

SetReviewCommentUrl sets ReviewCommentUrl field to given value.


### GetCommentsUrl

`func (o *PullRequest) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *PullRequest) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *PullRequest) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetStatusesUrl

`func (o *PullRequest) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *PullRequest) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *PullRequest) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetNumber

`func (o *PullRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *PullRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequest) SetState(v string)`

SetState sets State field to given value.


### GetLocked

`func (o *PullRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *PullRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *PullRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetTitle

`func (o *PullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUser

`func (o *PullRequest) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequest) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequest) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *PullRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PullRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetBody

`func (o *PullRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullRequest) SetBody(v string)`

SetBody sets Body field to given value.


### SetBodyNil

`func (o *PullRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *PullRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetLabels

`func (o *PullRequest) GetLabels() []PullRequestLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PullRequest) GetLabelsOk() (*[]PullRequestLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PullRequest) SetLabels(v []PullRequestLabelsInner)`

SetLabels sets Labels field to given value.


### GetMilestone

`func (o *PullRequest) GetMilestone() NullableMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *PullRequest) GetMilestoneOk() (*NullableMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *PullRequest) SetMilestone(v NullableMilestone)`

SetMilestone sets Milestone field to given value.


### SetMilestoneNil

`func (o *PullRequest) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *PullRequest) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetActiveLockReason

`func (o *PullRequest) GetActiveLockReason() string`

GetActiveLockReason returns the ActiveLockReason field if non-nil, zero value otherwise.

### GetActiveLockReasonOk

`func (o *PullRequest) GetActiveLockReasonOk() (*string, bool)`

GetActiveLockReasonOk returns a tuple with the ActiveLockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLockReason

`func (o *PullRequest) SetActiveLockReason(v string)`

SetActiveLockReason sets ActiveLockReason field to given value.

### HasActiveLockReason

`func (o *PullRequest) HasActiveLockReason() bool`

HasActiveLockReason returns a boolean if a field has been set.

### SetActiveLockReasonNil

`func (o *PullRequest) SetActiveLockReasonNil(b bool)`

 SetActiveLockReasonNil sets the value for ActiveLockReason to be an explicit nil

### UnsetActiveLockReason
`func (o *PullRequest) UnsetActiveLockReason()`

UnsetActiveLockReason ensures that no value is present for ActiveLockReason, not even an explicit nil
### GetCreatedAt

`func (o *PullRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PullRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *PullRequest) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *PullRequest) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *PullRequest) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *PullRequest) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *PullRequest) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetMergedAt

`func (o *PullRequest) GetMergedAt() time.Time`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *PullRequest) GetMergedAtOk() (*time.Time, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *PullRequest) SetMergedAt(v time.Time)`

SetMergedAt sets MergedAt field to given value.


### SetMergedAtNil

`func (o *PullRequest) SetMergedAtNil(b bool)`

 SetMergedAtNil sets the value for MergedAt to be an explicit nil

### UnsetMergedAt
`func (o *PullRequest) UnsetMergedAt()`

UnsetMergedAt ensures that no value is present for MergedAt, not even an explicit nil
### GetMergeCommitSha

`func (o *PullRequest) GetMergeCommitSha() string`

GetMergeCommitSha returns the MergeCommitSha field if non-nil, zero value otherwise.

### GetMergeCommitShaOk

`func (o *PullRequest) GetMergeCommitShaOk() (*string, bool)`

GetMergeCommitShaOk returns a tuple with the MergeCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitSha

`func (o *PullRequest) SetMergeCommitSha(v string)`

SetMergeCommitSha sets MergeCommitSha field to given value.


### SetMergeCommitShaNil

`func (o *PullRequest) SetMergeCommitShaNil(b bool)`

 SetMergeCommitShaNil sets the value for MergeCommitSha to be an explicit nil

### UnsetMergeCommitSha
`func (o *PullRequest) UnsetMergeCommitSha()`

UnsetMergeCommitSha ensures that no value is present for MergeCommitSha, not even an explicit nil
### GetAssignee

`func (o *PullRequest) GetAssignee() NullableSimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *PullRequest) GetAssigneeOk() (*NullableSimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *PullRequest) SetAssignee(v NullableSimpleUser)`

SetAssignee sets Assignee field to given value.


### SetAssigneeNil

`func (o *PullRequest) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *PullRequest) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAssignees

`func (o *PullRequest) GetAssignees() []SimpleUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *PullRequest) GetAssigneesOk() (*[]SimpleUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *PullRequest) SetAssignees(v []SimpleUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *PullRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### SetAssigneesNil

`func (o *PullRequest) SetAssigneesNil(b bool)`

 SetAssigneesNil sets the value for Assignees to be an explicit nil

### UnsetAssignees
`func (o *PullRequest) UnsetAssignees()`

UnsetAssignees ensures that no value is present for Assignees, not even an explicit nil
### GetRequestedReviewers

`func (o *PullRequest) GetRequestedReviewers() []SimpleUser`

GetRequestedReviewers returns the RequestedReviewers field if non-nil, zero value otherwise.

### GetRequestedReviewersOk

`func (o *PullRequest) GetRequestedReviewersOk() (*[]SimpleUser, bool)`

GetRequestedReviewersOk returns a tuple with the RequestedReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewers

`func (o *PullRequest) SetRequestedReviewers(v []SimpleUser)`

SetRequestedReviewers sets RequestedReviewers field to given value.

### HasRequestedReviewers

`func (o *PullRequest) HasRequestedReviewers() bool`

HasRequestedReviewers returns a boolean if a field has been set.

### SetRequestedReviewersNil

`func (o *PullRequest) SetRequestedReviewersNil(b bool)`

 SetRequestedReviewersNil sets the value for RequestedReviewers to be an explicit nil

### UnsetRequestedReviewers
`func (o *PullRequest) UnsetRequestedReviewers()`

UnsetRequestedReviewers ensures that no value is present for RequestedReviewers, not even an explicit nil
### GetRequestedTeams

`func (o *PullRequest) GetRequestedTeams() []TeamSimple`

GetRequestedTeams returns the RequestedTeams field if non-nil, zero value otherwise.

### GetRequestedTeamsOk

`func (o *PullRequest) GetRequestedTeamsOk() (*[]TeamSimple, bool)`

GetRequestedTeamsOk returns a tuple with the RequestedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeams

`func (o *PullRequest) SetRequestedTeams(v []TeamSimple)`

SetRequestedTeams sets RequestedTeams field to given value.

### HasRequestedTeams

`func (o *PullRequest) HasRequestedTeams() bool`

HasRequestedTeams returns a boolean if a field has been set.

### SetRequestedTeamsNil

`func (o *PullRequest) SetRequestedTeamsNil(b bool)`

 SetRequestedTeamsNil sets the value for RequestedTeams to be an explicit nil

### UnsetRequestedTeams
`func (o *PullRequest) UnsetRequestedTeams()`

UnsetRequestedTeams ensures that no value is present for RequestedTeams, not even an explicit nil
### GetHead

`func (o *PullRequest) GetHead() PullRequestHead`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullRequest) GetHeadOk() (*PullRequestHead, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullRequest) SetHead(v PullRequestHead)`

SetHead sets Head field to given value.


### GetBase

`func (o *PullRequest) GetBase() PullRequestBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullRequest) GetBaseOk() (*PullRequestBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullRequest) SetBase(v PullRequestBase)`

SetBase sets Base field to given value.


### GetLinks

`func (o *PullRequest) GetLinks() PullRequestSimpleLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PullRequest) GetLinksOk() (*PullRequestSimpleLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PullRequest) SetLinks(v PullRequestSimpleLinks)`

SetLinks sets Links field to given value.


### GetAuthorAssociation

`func (o *PullRequest) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *PullRequest) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *PullRequest) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetAutoMerge

`func (o *PullRequest) GetAutoMerge() AutoMerge`

GetAutoMerge returns the AutoMerge field if non-nil, zero value otherwise.

### GetAutoMergeOk

`func (o *PullRequest) GetAutoMergeOk() (*AutoMerge, bool)`

GetAutoMergeOk returns a tuple with the AutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMerge

`func (o *PullRequest) SetAutoMerge(v AutoMerge)`

SetAutoMerge sets AutoMerge field to given value.


### SetAutoMergeNil

`func (o *PullRequest) SetAutoMergeNil(b bool)`

 SetAutoMergeNil sets the value for AutoMerge to be an explicit nil

### UnsetAutoMerge
`func (o *PullRequest) UnsetAutoMerge()`

UnsetAutoMerge ensures that no value is present for AutoMerge, not even an explicit nil
### GetDraft

`func (o *PullRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PullRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PullRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *PullRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetMerged

`func (o *PullRequest) GetMerged() bool`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *PullRequest) GetMergedOk() (*bool, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *PullRequest) SetMerged(v bool)`

SetMerged sets Merged field to given value.


### GetMergeable

`func (o *PullRequest) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *PullRequest) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *PullRequest) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.


### SetMergeableNil

`func (o *PullRequest) SetMergeableNil(b bool)`

 SetMergeableNil sets the value for Mergeable to be an explicit nil

### UnsetMergeable
`func (o *PullRequest) UnsetMergeable()`

UnsetMergeable ensures that no value is present for Mergeable, not even an explicit nil
### GetRebaseable

`func (o *PullRequest) GetRebaseable() bool`

GetRebaseable returns the Rebaseable field if non-nil, zero value otherwise.

### GetRebaseableOk

`func (o *PullRequest) GetRebaseableOk() (*bool, bool)`

GetRebaseableOk returns a tuple with the Rebaseable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebaseable

`func (o *PullRequest) SetRebaseable(v bool)`

SetRebaseable sets Rebaseable field to given value.

### HasRebaseable

`func (o *PullRequest) HasRebaseable() bool`

HasRebaseable returns a boolean if a field has been set.

### SetRebaseableNil

`func (o *PullRequest) SetRebaseableNil(b bool)`

 SetRebaseableNil sets the value for Rebaseable to be an explicit nil

### UnsetRebaseable
`func (o *PullRequest) UnsetRebaseable()`

UnsetRebaseable ensures that no value is present for Rebaseable, not even an explicit nil
### GetMergeableState

`func (o *PullRequest) GetMergeableState() string`

GetMergeableState returns the MergeableState field if non-nil, zero value otherwise.

### GetMergeableStateOk

`func (o *PullRequest) GetMergeableStateOk() (*string, bool)`

GetMergeableStateOk returns a tuple with the MergeableState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeableState

`func (o *PullRequest) SetMergeableState(v string)`

SetMergeableState sets MergeableState field to given value.


### GetMergedBy

`func (o *PullRequest) GetMergedBy() NullableSimpleUser`

GetMergedBy returns the MergedBy field if non-nil, zero value otherwise.

### GetMergedByOk

`func (o *PullRequest) GetMergedByOk() (*NullableSimpleUser, bool)`

GetMergedByOk returns a tuple with the MergedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedBy

`func (o *PullRequest) SetMergedBy(v NullableSimpleUser)`

SetMergedBy sets MergedBy field to given value.


### SetMergedByNil

`func (o *PullRequest) SetMergedByNil(b bool)`

 SetMergedByNil sets the value for MergedBy to be an explicit nil

### UnsetMergedBy
`func (o *PullRequest) UnsetMergedBy()`

UnsetMergedBy ensures that no value is present for MergedBy, not even an explicit nil
### GetComments

`func (o *PullRequest) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PullRequest) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PullRequest) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetReviewComments

`func (o *PullRequest) GetReviewComments() int32`

GetReviewComments returns the ReviewComments field if non-nil, zero value otherwise.

### GetReviewCommentsOk

`func (o *PullRequest) GetReviewCommentsOk() (*int32, bool)`

GetReviewCommentsOk returns a tuple with the ReviewComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewComments

`func (o *PullRequest) SetReviewComments(v int32)`

SetReviewComments sets ReviewComments field to given value.


### GetMaintainerCanModify

`func (o *PullRequest) GetMaintainerCanModify() bool`

GetMaintainerCanModify returns the MaintainerCanModify field if non-nil, zero value otherwise.

### GetMaintainerCanModifyOk

`func (o *PullRequest) GetMaintainerCanModifyOk() (*bool, bool)`

GetMaintainerCanModifyOk returns a tuple with the MaintainerCanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerCanModify

`func (o *PullRequest) SetMaintainerCanModify(v bool)`

SetMaintainerCanModify sets MaintainerCanModify field to given value.


### GetCommits

`func (o *PullRequest) GetCommits() int32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *PullRequest) GetCommitsOk() (*int32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *PullRequest) SetCommits(v int32)`

SetCommits sets Commits field to given value.


### GetAdditions

`func (o *PullRequest) GetAdditions() int32`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *PullRequest) GetAdditionsOk() (*int32, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *PullRequest) SetAdditions(v int32)`

SetAdditions sets Additions field to given value.


### GetDeletions

`func (o *PullRequest) GetDeletions() int32`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *PullRequest) GetDeletionsOk() (*int32, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *PullRequest) SetDeletions(v int32)`

SetDeletions sets Deletions field to given value.


### GetChangedFiles

`func (o *PullRequest) GetChangedFiles() int32`

GetChangedFiles returns the ChangedFiles field if non-nil, zero value otherwise.

### GetChangedFilesOk

`func (o *PullRequest) GetChangedFilesOk() (*int32, bool)`

GetChangedFilesOk returns a tuple with the ChangedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedFiles

`func (o *PullRequest) SetChangedFiles(v int32)`

SetChangedFiles sets ChangedFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


