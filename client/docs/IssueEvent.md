# IssueEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Actor** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Event** | **string** |  | 
**CommitId** | **NullableString** |  | 
**CommitUrl** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**Issue** | Pointer to [**NullableNullableIssue**](NullableIssue.md) |  | [optional] 
**Label** | Pointer to [**IssueEventLabel**](IssueEventLabel.md) |  | [optional] 
**Assignee** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**Assigner** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**ReviewRequester** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**RequestedReviewer** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**RequestedTeam** | Pointer to [**Team**](Team.md) |  | [optional] 
**DismissedReview** | Pointer to [**IssueEventDismissedReview**](IssueEventDismissedReview.md) |  | [optional] 
**Milestone** | Pointer to [**IssueEventMilestone**](IssueEventMilestone.md) |  | [optional] 
**ProjectCard** | Pointer to [**IssueEventProjectCard**](IssueEventProjectCard.md) |  | [optional] 
**Rename** | Pointer to [**IssueEventRename**](IssueEventRename.md) |  | [optional] 
**AuthorAssociation** | Pointer to [**AuthorAssociation**](AuthorAssociation.md) |  | [optional] 
**LockReason** | Pointer to **NullableString** |  | [optional] 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 

## Methods

### NewIssueEvent

`func NewIssueEvent(id int32, nodeId string, url string, actor NullableNullableSimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt time.Time, ) *IssueEvent`

NewIssueEvent instantiates a new IssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEventWithDefaults

`func NewIssueEventWithDefaults() *IssueEvent`

NewIssueEventWithDefaults instantiates a new IssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *IssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *IssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *IssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *IssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *IssueEvent) GetActor() NullableSimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *IssueEvent) GetActorOk() (*NullableSimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *IssueEvent) SetActor(v NullableSimpleUser)`

SetActor sets Actor field to given value.


### SetActorNil

`func (o *IssueEvent) SetActorNil(b bool)`

 SetActorNil sets the value for Actor to be an explicit nil

### UnsetActor
`func (o *IssueEvent) UnsetActor()`

UnsetActor ensures that no value is present for Actor, not even an explicit nil
### GetEvent

`func (o *IssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *IssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *IssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *IssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *IssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *IssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *IssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *IssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *IssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *IssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *IssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *IssueEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIssue

`func (o *IssueEvent) GetIssue() NullableIssue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *IssueEvent) GetIssueOk() (*NullableIssue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *IssueEvent) SetIssue(v NullableIssue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *IssueEvent) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### SetIssueNil

`func (o *IssueEvent) SetIssueNil(b bool)`

 SetIssueNil sets the value for Issue to be an explicit nil

### UnsetIssue
`func (o *IssueEvent) UnsetIssue()`

UnsetIssue ensures that no value is present for Issue, not even an explicit nil
### GetLabel

`func (o *IssueEvent) GetLabel() IssueEventLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IssueEvent) GetLabelOk() (*IssueEventLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IssueEvent) SetLabel(v IssueEventLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IssueEvent) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAssignee

`func (o *IssueEvent) GetAssignee() NullableSimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssueEvent) GetAssigneeOk() (*NullableSimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssueEvent) SetAssignee(v NullableSimpleUser)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *IssueEvent) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *IssueEvent) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *IssueEvent) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAssigner

`func (o *IssueEvent) GetAssigner() NullableSimpleUser`

GetAssigner returns the Assigner field if non-nil, zero value otherwise.

### GetAssignerOk

`func (o *IssueEvent) GetAssignerOk() (*NullableSimpleUser, bool)`

GetAssignerOk returns a tuple with the Assigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigner

`func (o *IssueEvent) SetAssigner(v NullableSimpleUser)`

SetAssigner sets Assigner field to given value.

### HasAssigner

`func (o *IssueEvent) HasAssigner() bool`

HasAssigner returns a boolean if a field has been set.

### SetAssignerNil

`func (o *IssueEvent) SetAssignerNil(b bool)`

 SetAssignerNil sets the value for Assigner to be an explicit nil

### UnsetAssigner
`func (o *IssueEvent) UnsetAssigner()`

UnsetAssigner ensures that no value is present for Assigner, not even an explicit nil
### GetReviewRequester

`func (o *IssueEvent) GetReviewRequester() NullableSimpleUser`

GetReviewRequester returns the ReviewRequester field if non-nil, zero value otherwise.

### GetReviewRequesterOk

`func (o *IssueEvent) GetReviewRequesterOk() (*NullableSimpleUser, bool)`

GetReviewRequesterOk returns a tuple with the ReviewRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequester

`func (o *IssueEvent) SetReviewRequester(v NullableSimpleUser)`

SetReviewRequester sets ReviewRequester field to given value.

### HasReviewRequester

`func (o *IssueEvent) HasReviewRequester() bool`

HasReviewRequester returns a boolean if a field has been set.

### SetReviewRequesterNil

`func (o *IssueEvent) SetReviewRequesterNil(b bool)`

 SetReviewRequesterNil sets the value for ReviewRequester to be an explicit nil

### UnsetReviewRequester
`func (o *IssueEvent) UnsetReviewRequester()`

UnsetReviewRequester ensures that no value is present for ReviewRequester, not even an explicit nil
### GetRequestedReviewer

`func (o *IssueEvent) GetRequestedReviewer() NullableSimpleUser`

GetRequestedReviewer returns the RequestedReviewer field if non-nil, zero value otherwise.

### GetRequestedReviewerOk

`func (o *IssueEvent) GetRequestedReviewerOk() (*NullableSimpleUser, bool)`

GetRequestedReviewerOk returns a tuple with the RequestedReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewer

`func (o *IssueEvent) SetRequestedReviewer(v NullableSimpleUser)`

SetRequestedReviewer sets RequestedReviewer field to given value.

### HasRequestedReviewer

`func (o *IssueEvent) HasRequestedReviewer() bool`

HasRequestedReviewer returns a boolean if a field has been set.

### SetRequestedReviewerNil

`func (o *IssueEvent) SetRequestedReviewerNil(b bool)`

 SetRequestedReviewerNil sets the value for RequestedReviewer to be an explicit nil

### UnsetRequestedReviewer
`func (o *IssueEvent) UnsetRequestedReviewer()`

UnsetRequestedReviewer ensures that no value is present for RequestedReviewer, not even an explicit nil
### GetRequestedTeam

`func (o *IssueEvent) GetRequestedTeam() Team`

GetRequestedTeam returns the RequestedTeam field if non-nil, zero value otherwise.

### GetRequestedTeamOk

`func (o *IssueEvent) GetRequestedTeamOk() (*Team, bool)`

GetRequestedTeamOk returns a tuple with the RequestedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeam

`func (o *IssueEvent) SetRequestedTeam(v Team)`

SetRequestedTeam sets RequestedTeam field to given value.

### HasRequestedTeam

`func (o *IssueEvent) HasRequestedTeam() bool`

HasRequestedTeam returns a boolean if a field has been set.

### GetDismissedReview

`func (o *IssueEvent) GetDismissedReview() IssueEventDismissedReview`

GetDismissedReview returns the DismissedReview field if non-nil, zero value otherwise.

### GetDismissedReviewOk

`func (o *IssueEvent) GetDismissedReviewOk() (*IssueEventDismissedReview, bool)`

GetDismissedReviewOk returns a tuple with the DismissedReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedReview

`func (o *IssueEvent) SetDismissedReview(v IssueEventDismissedReview)`

SetDismissedReview sets DismissedReview field to given value.

### HasDismissedReview

`func (o *IssueEvent) HasDismissedReview() bool`

HasDismissedReview returns a boolean if a field has been set.

### GetMilestone

`func (o *IssueEvent) GetMilestone() IssueEventMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *IssueEvent) GetMilestoneOk() (*IssueEventMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *IssueEvent) SetMilestone(v IssueEventMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *IssueEvent) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetProjectCard

`func (o *IssueEvent) GetProjectCard() IssueEventProjectCard`

GetProjectCard returns the ProjectCard field if non-nil, zero value otherwise.

### GetProjectCardOk

`func (o *IssueEvent) GetProjectCardOk() (*IssueEventProjectCard, bool)`

GetProjectCardOk returns a tuple with the ProjectCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCard

`func (o *IssueEvent) SetProjectCard(v IssueEventProjectCard)`

SetProjectCard sets ProjectCard field to given value.

### HasProjectCard

`func (o *IssueEvent) HasProjectCard() bool`

HasProjectCard returns a boolean if a field has been set.

### GetRename

`func (o *IssueEvent) GetRename() IssueEventRename`

GetRename returns the Rename field if non-nil, zero value otherwise.

### GetRenameOk

`func (o *IssueEvent) GetRenameOk() (*IssueEventRename, bool)`

GetRenameOk returns a tuple with the Rename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRename

`func (o *IssueEvent) SetRename(v IssueEventRename)`

SetRename sets Rename field to given value.

### HasRename

`func (o *IssueEvent) HasRename() bool`

HasRename returns a boolean if a field has been set.

### GetAuthorAssociation

`func (o *IssueEvent) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *IssueEvent) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *IssueEvent) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.

### HasAuthorAssociation

`func (o *IssueEvent) HasAuthorAssociation() bool`

HasAuthorAssociation returns a boolean if a field has been set.

### GetLockReason

`func (o *IssueEvent) GetLockReason() string`

GetLockReason returns the LockReason field if non-nil, zero value otherwise.

### GetLockReasonOk

`func (o *IssueEvent) GetLockReasonOk() (*string, bool)`

GetLockReasonOk returns a tuple with the LockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockReason

`func (o *IssueEvent) SetLockReason(v string)`

SetLockReason sets LockReason field to given value.

### HasLockReason

`func (o *IssueEvent) HasLockReason() bool`

HasLockReason returns a boolean if a field has been set.

### SetLockReasonNil

`func (o *IssueEvent) SetLockReasonNil(b bool)`

 SetLockReasonNil sets the value for LockReason to be an explicit nil

### UnsetLockReason
`func (o *IssueEvent) UnsetLockReason()`

UnsetLockReason ensures that no value is present for LockReason, not even an explicit nil
### GetPerformedViaGithubApp

`func (o *IssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *IssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *IssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *IssueEvent) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *IssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *IssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


