# PullRequestSimple

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
**Number** | **int32** |  | 
**State** | **string** |  | 
**Locked** | **bool** |  | 
**Title** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Body** | **NullableString** |  | 
**Labels** | [**[]PullRequestSimpleLabelsInner**](PullRequestSimpleLabelsInner.md) |  | 
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
**RequestedTeams** | Pointer to [**[]Team**](Team.md) |  | [optional] 
**Head** | [**PullRequestSimpleHead**](PullRequestSimpleHead.md) |  | 
**Base** | [**PullRequestSimpleHead**](PullRequestSimpleHead.md) |  | 
**Links** | [**PullRequestSimpleLinks**](PullRequestSimpleLinks.md) |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**AutoMerge** | [**NullableAutoMerge**](AutoMerge.md) |  | 
**Draft** | Pointer to **bool** | Indicates whether or not the pull request is a draft. | [optional] 

## Methods

### NewPullRequestSimple

`func NewPullRequestSimple(url string, id int32, nodeId string, htmlUrl string, diffUrl string, patchUrl string, issueUrl string, commitsUrl string, reviewCommentsUrl string, reviewCommentUrl string, commentsUrl string, statusesUrl string, number int32, state string, locked bool, title string, user NullableNullableSimpleUser, body NullableString, labels []PullRequestSimpleLabelsInner, milestone NullableNullableMilestone, createdAt time.Time, updatedAt time.Time, closedAt NullableTime, mergedAt NullableTime, mergeCommitSha NullableString, assignee NullableNullableSimpleUser, head PullRequestSimpleHead, base PullRequestSimpleHead, links PullRequestSimpleLinks, authorAssociation AuthorAssociation, autoMerge NullableAutoMerge, ) *PullRequestSimple`

NewPullRequestSimple instantiates a new PullRequestSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestSimpleWithDefaults

`func NewPullRequestSimpleWithDefaults() *PullRequestSimple`

NewPullRequestSimpleWithDefaults instantiates a new PullRequestSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PullRequestSimple) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestSimple) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestSimple) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *PullRequestSimple) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestSimple) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestSimple) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PullRequestSimple) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PullRequestSimple) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PullRequestSimple) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHtmlUrl

`func (o *PullRequestSimple) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *PullRequestSimple) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *PullRequestSimple) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDiffUrl

`func (o *PullRequestSimple) GetDiffUrl() string`

GetDiffUrl returns the DiffUrl field if non-nil, zero value otherwise.

### GetDiffUrlOk

`func (o *PullRequestSimple) GetDiffUrlOk() (*string, bool)`

GetDiffUrlOk returns a tuple with the DiffUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffUrl

`func (o *PullRequestSimple) SetDiffUrl(v string)`

SetDiffUrl sets DiffUrl field to given value.


### GetPatchUrl

`func (o *PullRequestSimple) GetPatchUrl() string`

GetPatchUrl returns the PatchUrl field if non-nil, zero value otherwise.

### GetPatchUrlOk

`func (o *PullRequestSimple) GetPatchUrlOk() (*string, bool)`

GetPatchUrlOk returns a tuple with the PatchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchUrl

`func (o *PullRequestSimple) SetPatchUrl(v string)`

SetPatchUrl sets PatchUrl field to given value.


### GetIssueUrl

`func (o *PullRequestSimple) GetIssueUrl() string`

GetIssueUrl returns the IssueUrl field if non-nil, zero value otherwise.

### GetIssueUrlOk

`func (o *PullRequestSimple) GetIssueUrlOk() (*string, bool)`

GetIssueUrlOk returns a tuple with the IssueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueUrl

`func (o *PullRequestSimple) SetIssueUrl(v string)`

SetIssueUrl sets IssueUrl field to given value.


### GetCommitsUrl

`func (o *PullRequestSimple) GetCommitsUrl() string`

GetCommitsUrl returns the CommitsUrl field if non-nil, zero value otherwise.

### GetCommitsUrlOk

`func (o *PullRequestSimple) GetCommitsUrlOk() (*string, bool)`

GetCommitsUrlOk returns a tuple with the CommitsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsUrl

`func (o *PullRequestSimple) SetCommitsUrl(v string)`

SetCommitsUrl sets CommitsUrl field to given value.


### GetReviewCommentsUrl

`func (o *PullRequestSimple) GetReviewCommentsUrl() string`

GetReviewCommentsUrl returns the ReviewCommentsUrl field if non-nil, zero value otherwise.

### GetReviewCommentsUrlOk

`func (o *PullRequestSimple) GetReviewCommentsUrlOk() (*string, bool)`

GetReviewCommentsUrlOk returns a tuple with the ReviewCommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCommentsUrl

`func (o *PullRequestSimple) SetReviewCommentsUrl(v string)`

SetReviewCommentsUrl sets ReviewCommentsUrl field to given value.


### GetReviewCommentUrl

`func (o *PullRequestSimple) GetReviewCommentUrl() string`

GetReviewCommentUrl returns the ReviewCommentUrl field if non-nil, zero value otherwise.

### GetReviewCommentUrlOk

`func (o *PullRequestSimple) GetReviewCommentUrlOk() (*string, bool)`

GetReviewCommentUrlOk returns a tuple with the ReviewCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCommentUrl

`func (o *PullRequestSimple) SetReviewCommentUrl(v string)`

SetReviewCommentUrl sets ReviewCommentUrl field to given value.


### GetCommentsUrl

`func (o *PullRequestSimple) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *PullRequestSimple) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *PullRequestSimple) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetStatusesUrl

`func (o *PullRequestSimple) GetStatusesUrl() string`

GetStatusesUrl returns the StatusesUrl field if non-nil, zero value otherwise.

### GetStatusesUrlOk

`func (o *PullRequestSimple) GetStatusesUrlOk() (*string, bool)`

GetStatusesUrlOk returns a tuple with the StatusesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesUrl

`func (o *PullRequestSimple) SetStatusesUrl(v string)`

SetStatusesUrl sets StatusesUrl field to given value.


### GetNumber

`func (o *PullRequestSimple) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullRequestSimple) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullRequestSimple) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *PullRequestSimple) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PullRequestSimple) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PullRequestSimple) SetState(v string)`

SetState sets State field to given value.


### GetLocked

`func (o *PullRequestSimple) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *PullRequestSimple) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *PullRequestSimple) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetTitle

`func (o *PullRequestSimple) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequestSimple) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequestSimple) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUser

`func (o *PullRequestSimple) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestSimple) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestSimple) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *PullRequestSimple) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PullRequestSimple) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetBody

`func (o *PullRequestSimple) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullRequestSimple) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullRequestSimple) SetBody(v string)`

SetBody sets Body field to given value.


### SetBodyNil

`func (o *PullRequestSimple) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *PullRequestSimple) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetLabels

`func (o *PullRequestSimple) GetLabels() []PullRequestSimpleLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PullRequestSimple) GetLabelsOk() (*[]PullRequestSimpleLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PullRequestSimple) SetLabels(v []PullRequestSimpleLabelsInner)`

SetLabels sets Labels field to given value.


### GetMilestone

`func (o *PullRequestSimple) GetMilestone() NullableMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *PullRequestSimple) GetMilestoneOk() (*NullableMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *PullRequestSimple) SetMilestone(v NullableMilestone)`

SetMilestone sets Milestone field to given value.


### SetMilestoneNil

`func (o *PullRequestSimple) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *PullRequestSimple) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetActiveLockReason

`func (o *PullRequestSimple) GetActiveLockReason() string`

GetActiveLockReason returns the ActiveLockReason field if non-nil, zero value otherwise.

### GetActiveLockReasonOk

`func (o *PullRequestSimple) GetActiveLockReasonOk() (*string, bool)`

GetActiveLockReasonOk returns a tuple with the ActiveLockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLockReason

`func (o *PullRequestSimple) SetActiveLockReason(v string)`

SetActiveLockReason sets ActiveLockReason field to given value.

### HasActiveLockReason

`func (o *PullRequestSimple) HasActiveLockReason() bool`

HasActiveLockReason returns a boolean if a field has been set.

### SetActiveLockReasonNil

`func (o *PullRequestSimple) SetActiveLockReasonNil(b bool)`

 SetActiveLockReasonNil sets the value for ActiveLockReason to be an explicit nil

### UnsetActiveLockReason
`func (o *PullRequestSimple) UnsetActiveLockReason()`

UnsetActiveLockReason ensures that no value is present for ActiveLockReason, not even an explicit nil
### GetCreatedAt

`func (o *PullRequestSimple) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequestSimple) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequestSimple) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PullRequestSimple) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequestSimple) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequestSimple) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *PullRequestSimple) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *PullRequestSimple) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *PullRequestSimple) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *PullRequestSimple) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *PullRequestSimple) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetMergedAt

`func (o *PullRequestSimple) GetMergedAt() time.Time`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *PullRequestSimple) GetMergedAtOk() (*time.Time, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *PullRequestSimple) SetMergedAt(v time.Time)`

SetMergedAt sets MergedAt field to given value.


### SetMergedAtNil

`func (o *PullRequestSimple) SetMergedAtNil(b bool)`

 SetMergedAtNil sets the value for MergedAt to be an explicit nil

### UnsetMergedAt
`func (o *PullRequestSimple) UnsetMergedAt()`

UnsetMergedAt ensures that no value is present for MergedAt, not even an explicit nil
### GetMergeCommitSha

`func (o *PullRequestSimple) GetMergeCommitSha() string`

GetMergeCommitSha returns the MergeCommitSha field if non-nil, zero value otherwise.

### GetMergeCommitShaOk

`func (o *PullRequestSimple) GetMergeCommitShaOk() (*string, bool)`

GetMergeCommitShaOk returns a tuple with the MergeCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitSha

`func (o *PullRequestSimple) SetMergeCommitSha(v string)`

SetMergeCommitSha sets MergeCommitSha field to given value.


### SetMergeCommitShaNil

`func (o *PullRequestSimple) SetMergeCommitShaNil(b bool)`

 SetMergeCommitShaNil sets the value for MergeCommitSha to be an explicit nil

### UnsetMergeCommitSha
`func (o *PullRequestSimple) UnsetMergeCommitSha()`

UnsetMergeCommitSha ensures that no value is present for MergeCommitSha, not even an explicit nil
### GetAssignee

`func (o *PullRequestSimple) GetAssignee() NullableSimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *PullRequestSimple) GetAssigneeOk() (*NullableSimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *PullRequestSimple) SetAssignee(v NullableSimpleUser)`

SetAssignee sets Assignee field to given value.


### SetAssigneeNil

`func (o *PullRequestSimple) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *PullRequestSimple) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAssignees

`func (o *PullRequestSimple) GetAssignees() []SimpleUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *PullRequestSimple) GetAssigneesOk() (*[]SimpleUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *PullRequestSimple) SetAssignees(v []SimpleUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *PullRequestSimple) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### SetAssigneesNil

`func (o *PullRequestSimple) SetAssigneesNil(b bool)`

 SetAssigneesNil sets the value for Assignees to be an explicit nil

### UnsetAssignees
`func (o *PullRequestSimple) UnsetAssignees()`

UnsetAssignees ensures that no value is present for Assignees, not even an explicit nil
### GetRequestedReviewers

`func (o *PullRequestSimple) GetRequestedReviewers() []SimpleUser`

GetRequestedReviewers returns the RequestedReviewers field if non-nil, zero value otherwise.

### GetRequestedReviewersOk

`func (o *PullRequestSimple) GetRequestedReviewersOk() (*[]SimpleUser, bool)`

GetRequestedReviewersOk returns a tuple with the RequestedReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewers

`func (o *PullRequestSimple) SetRequestedReviewers(v []SimpleUser)`

SetRequestedReviewers sets RequestedReviewers field to given value.

### HasRequestedReviewers

`func (o *PullRequestSimple) HasRequestedReviewers() bool`

HasRequestedReviewers returns a boolean if a field has been set.

### SetRequestedReviewersNil

`func (o *PullRequestSimple) SetRequestedReviewersNil(b bool)`

 SetRequestedReviewersNil sets the value for RequestedReviewers to be an explicit nil

### UnsetRequestedReviewers
`func (o *PullRequestSimple) UnsetRequestedReviewers()`

UnsetRequestedReviewers ensures that no value is present for RequestedReviewers, not even an explicit nil
### GetRequestedTeams

`func (o *PullRequestSimple) GetRequestedTeams() []Team`

GetRequestedTeams returns the RequestedTeams field if non-nil, zero value otherwise.

### GetRequestedTeamsOk

`func (o *PullRequestSimple) GetRequestedTeamsOk() (*[]Team, bool)`

GetRequestedTeamsOk returns a tuple with the RequestedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeams

`func (o *PullRequestSimple) SetRequestedTeams(v []Team)`

SetRequestedTeams sets RequestedTeams field to given value.

### HasRequestedTeams

`func (o *PullRequestSimple) HasRequestedTeams() bool`

HasRequestedTeams returns a boolean if a field has been set.

### SetRequestedTeamsNil

`func (o *PullRequestSimple) SetRequestedTeamsNil(b bool)`

 SetRequestedTeamsNil sets the value for RequestedTeams to be an explicit nil

### UnsetRequestedTeams
`func (o *PullRequestSimple) UnsetRequestedTeams()`

UnsetRequestedTeams ensures that no value is present for RequestedTeams, not even an explicit nil
### GetHead

`func (o *PullRequestSimple) GetHead() PullRequestSimpleHead`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullRequestSimple) GetHeadOk() (*PullRequestSimpleHead, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullRequestSimple) SetHead(v PullRequestSimpleHead)`

SetHead sets Head field to given value.


### GetBase

`func (o *PullRequestSimple) GetBase() PullRequestSimpleHead`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullRequestSimple) GetBaseOk() (*PullRequestSimpleHead, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullRequestSimple) SetBase(v PullRequestSimpleHead)`

SetBase sets Base field to given value.


### GetLinks

`func (o *PullRequestSimple) GetLinks() PullRequestSimpleLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PullRequestSimple) GetLinksOk() (*PullRequestSimpleLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PullRequestSimple) SetLinks(v PullRequestSimpleLinks)`

SetLinks sets Links field to given value.


### GetAuthorAssociation

`func (o *PullRequestSimple) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *PullRequestSimple) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *PullRequestSimple) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetAutoMerge

`func (o *PullRequestSimple) GetAutoMerge() AutoMerge`

GetAutoMerge returns the AutoMerge field if non-nil, zero value otherwise.

### GetAutoMergeOk

`func (o *PullRequestSimple) GetAutoMergeOk() (*AutoMerge, bool)`

GetAutoMergeOk returns a tuple with the AutoMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMerge

`func (o *PullRequestSimple) SetAutoMerge(v AutoMerge)`

SetAutoMerge sets AutoMerge field to given value.


### SetAutoMergeNil

`func (o *PullRequestSimple) SetAutoMergeNil(b bool)`

 SetAutoMergeNil sets the value for AutoMerge to be an explicit nil

### UnsetAutoMerge
`func (o *PullRequestSimple) UnsetAutoMerge()`

UnsetAutoMerge ensures that no value is present for AutoMerge, not even an explicit nil
### GetDraft

`func (o *PullRequestSimple) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PullRequestSimple) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PullRequestSimple) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *PullRequestSimple) HasDraft() bool`

HasDraft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


