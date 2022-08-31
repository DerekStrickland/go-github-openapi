# IssueEventForIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Actor** | [**SimpleUser**](SimpleUser.md) |  | 
**Event** | **string** |  | 
**CommitId** | **NullableString** |  | 
**CommitUrl** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**PerformedViaGithubApp** | [**Integration**](Integration.md) |  | 
**Label** | [**LabeledIssueEventLabel**](LabeledIssueEventLabel.md) |  | 
**Assignee** | [**SimpleUser**](SimpleUser.md) |  | 
**Assigner** | [**SimpleUser**](SimpleUser.md) |  | 
**Milestone** | [**MilestonedIssueEventMilestone**](MilestonedIssueEventMilestone.md) |  | 
**Rename** | [**RenamedIssueEventRename**](RenamedIssueEventRename.md) |  | 
**ReviewRequester** | [**SimpleUser**](SimpleUser.md) |  | 
**RequestedTeam** | Pointer to [**Team**](Team.md) |  | [optional] 
**RequestedReviewer** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**DismissedReview** | [**ReviewDismissedIssueEventDismissedReview**](ReviewDismissedIssueEventDismissedReview.md) |  | 
**LockReason** | **NullableString** |  | 
**ProjectCard** | Pointer to [**AddedToProjectIssueEventProjectCard**](AddedToProjectIssueEventProjectCard.md) |  | [optional] 

## Methods

### NewIssueEventForIssue

`func NewIssueEventForIssue(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp Integration, label LabeledIssueEventLabel, assignee SimpleUser, assigner SimpleUser, milestone MilestonedIssueEventMilestone, rename RenamedIssueEventRename, reviewRequester SimpleUser, dismissedReview ReviewDismissedIssueEventDismissedReview, lockReason NullableString, ) *IssueEventForIssue`

NewIssueEventForIssue instantiates a new IssueEventForIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEventForIssueWithDefaults

`func NewIssueEventForIssueWithDefaults() *IssueEventForIssue`

NewIssueEventForIssueWithDefaults instantiates a new IssueEventForIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueEventForIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueEventForIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueEventForIssue) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *IssueEventForIssue) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *IssueEventForIssue) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *IssueEventForIssue) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *IssueEventForIssue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssueEventForIssue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssueEventForIssue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *IssueEventForIssue) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *IssueEventForIssue) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *IssueEventForIssue) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *IssueEventForIssue) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IssueEventForIssue) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IssueEventForIssue) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *IssueEventForIssue) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *IssueEventForIssue) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *IssueEventForIssue) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *IssueEventForIssue) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *IssueEventForIssue) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *IssueEventForIssue) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *IssueEventForIssue) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *IssueEventForIssue) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *IssueEventForIssue) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *IssueEventForIssue) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *IssueEventForIssue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueEventForIssue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueEventForIssue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *IssueEventForIssue) GetPerformedViaGithubApp() Integration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *IssueEventForIssue) GetPerformedViaGithubAppOk() (*Integration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *IssueEventForIssue) SetPerformedViaGithubApp(v Integration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### GetLabel

`func (o *IssueEventForIssue) GetLabel() LabeledIssueEventLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IssueEventForIssue) GetLabelOk() (*LabeledIssueEventLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IssueEventForIssue) SetLabel(v LabeledIssueEventLabel)`

SetLabel sets Label field to given value.


### GetAssignee

`func (o *IssueEventForIssue) GetAssignee() SimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssueEventForIssue) GetAssigneeOk() (*SimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssueEventForIssue) SetAssignee(v SimpleUser)`

SetAssignee sets Assignee field to given value.


### GetAssigner

`func (o *IssueEventForIssue) GetAssigner() SimpleUser`

GetAssigner returns the Assigner field if non-nil, zero value otherwise.

### GetAssignerOk

`func (o *IssueEventForIssue) GetAssignerOk() (*SimpleUser, bool)`

GetAssignerOk returns a tuple with the Assigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigner

`func (o *IssueEventForIssue) SetAssigner(v SimpleUser)`

SetAssigner sets Assigner field to given value.


### GetMilestone

`func (o *IssueEventForIssue) GetMilestone() MilestonedIssueEventMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *IssueEventForIssue) GetMilestoneOk() (*MilestonedIssueEventMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *IssueEventForIssue) SetMilestone(v MilestonedIssueEventMilestone)`

SetMilestone sets Milestone field to given value.


### GetRename

`func (o *IssueEventForIssue) GetRename() RenamedIssueEventRename`

GetRename returns the Rename field if non-nil, zero value otherwise.

### GetRenameOk

`func (o *IssueEventForIssue) GetRenameOk() (*RenamedIssueEventRename, bool)`

GetRenameOk returns a tuple with the Rename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRename

`func (o *IssueEventForIssue) SetRename(v RenamedIssueEventRename)`

SetRename sets Rename field to given value.


### GetReviewRequester

`func (o *IssueEventForIssue) GetReviewRequester() SimpleUser`

GetReviewRequester returns the ReviewRequester field if non-nil, zero value otherwise.

### GetReviewRequesterOk

`func (o *IssueEventForIssue) GetReviewRequesterOk() (*SimpleUser, bool)`

GetReviewRequesterOk returns a tuple with the ReviewRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequester

`func (o *IssueEventForIssue) SetReviewRequester(v SimpleUser)`

SetReviewRequester sets ReviewRequester field to given value.


### GetRequestedTeam

`func (o *IssueEventForIssue) GetRequestedTeam() Team`

GetRequestedTeam returns the RequestedTeam field if non-nil, zero value otherwise.

### GetRequestedTeamOk

`func (o *IssueEventForIssue) GetRequestedTeamOk() (*Team, bool)`

GetRequestedTeamOk returns a tuple with the RequestedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeam

`func (o *IssueEventForIssue) SetRequestedTeam(v Team)`

SetRequestedTeam sets RequestedTeam field to given value.

### HasRequestedTeam

`func (o *IssueEventForIssue) HasRequestedTeam() bool`

HasRequestedTeam returns a boolean if a field has been set.

### GetRequestedReviewer

`func (o *IssueEventForIssue) GetRequestedReviewer() SimpleUser`

GetRequestedReviewer returns the RequestedReviewer field if non-nil, zero value otherwise.

### GetRequestedReviewerOk

`func (o *IssueEventForIssue) GetRequestedReviewerOk() (*SimpleUser, bool)`

GetRequestedReviewerOk returns a tuple with the RequestedReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewer

`func (o *IssueEventForIssue) SetRequestedReviewer(v SimpleUser)`

SetRequestedReviewer sets RequestedReviewer field to given value.

### HasRequestedReviewer

`func (o *IssueEventForIssue) HasRequestedReviewer() bool`

HasRequestedReviewer returns a boolean if a field has been set.

### GetDismissedReview

`func (o *IssueEventForIssue) GetDismissedReview() ReviewDismissedIssueEventDismissedReview`

GetDismissedReview returns the DismissedReview field if non-nil, zero value otherwise.

### GetDismissedReviewOk

`func (o *IssueEventForIssue) GetDismissedReviewOk() (*ReviewDismissedIssueEventDismissedReview, bool)`

GetDismissedReviewOk returns a tuple with the DismissedReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedReview

`func (o *IssueEventForIssue) SetDismissedReview(v ReviewDismissedIssueEventDismissedReview)`

SetDismissedReview sets DismissedReview field to given value.


### GetLockReason

`func (o *IssueEventForIssue) GetLockReason() string`

GetLockReason returns the LockReason field if non-nil, zero value otherwise.

### GetLockReasonOk

`func (o *IssueEventForIssue) GetLockReasonOk() (*string, bool)`

GetLockReasonOk returns a tuple with the LockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockReason

`func (o *IssueEventForIssue) SetLockReason(v string)`

SetLockReason sets LockReason field to given value.


### SetLockReasonNil

`func (o *IssueEventForIssue) SetLockReasonNil(b bool)`

 SetLockReasonNil sets the value for LockReason to be an explicit nil

### UnsetLockReason
`func (o *IssueEventForIssue) UnsetLockReason()`

UnsetLockReason ensures that no value is present for LockReason, not even an explicit nil
### GetProjectCard

`func (o *IssueEventForIssue) GetProjectCard() AddedToProjectIssueEventProjectCard`

GetProjectCard returns the ProjectCard field if non-nil, zero value otherwise.

### GetProjectCardOk

`func (o *IssueEventForIssue) GetProjectCardOk() (*AddedToProjectIssueEventProjectCard, bool)`

GetProjectCardOk returns a tuple with the ProjectCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCard

`func (o *IssueEventForIssue) SetProjectCard(v AddedToProjectIssueEventProjectCard)`

SetProjectCard sets ProjectCard field to given value.

### HasProjectCard

`func (o *IssueEventForIssue) HasProjectCard() bool`

HasProjectCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


