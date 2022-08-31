# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** | URL for the issue | 
**RepositoryUrl** | **string** |  | 
**LabelsUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Number** | **int32** | Number uniquely identifying the issue within its repository | 
**State** | **string** | State of the issue; either &#39;open&#39; or &#39;closed&#39; | 
**StateReason** | Pointer to **NullableString** | The reason for the current state | [optional] 
**Title** | **string** | Title of the issue | 
**Body** | Pointer to **NullableString** | Contents of the issue | [optional] 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Labels** | [**[]IssueLabelsInner**](IssueLabelsInner.md) | Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository | 
**Assignee** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Assignees** | Pointer to [**[]SimpleUser**](SimpleUser.md) |  | [optional] 
**Milestone** | [**NullableNullableMilestone**](NullableMilestone.md) |  | 
**Locked** | **bool** |  | 
**ActiveLockReason** | Pointer to **NullableString** |  | [optional] 
**Comments** | **int32** |  | 
**PullRequest** | Pointer to [**IssuePullRequest**](IssuePullRequest.md) |  | [optional] 
**ClosedAt** | **NullableTime** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Draft** | Pointer to **bool** |  | [optional] 
**ClosedBy** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**TimelineUrl** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewIssue

`func NewIssue(id int32, nodeId string, url string, repositoryUrl string, labelsUrl string, commentsUrl string, eventsUrl string, htmlUrl string, number int32, state string, title string, user NullableNullableSimpleUser, labels []IssueLabelsInner, assignee NullableNullableSimpleUser, milestone NullableNullableMilestone, locked bool, comments int32, closedAt NullableTime, createdAt time.Time, updatedAt time.Time, authorAssociation AuthorAssociation, ) *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Issue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Issue) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Issue) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Issue) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *Issue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Issue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Issue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRepositoryUrl

`func (o *Issue) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *Issue) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *Issue) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetLabelsUrl

`func (o *Issue) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *Issue) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *Issue) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetCommentsUrl

`func (o *Issue) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *Issue) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *Issue) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetEventsUrl

`func (o *Issue) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *Issue) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *Issue) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetHtmlUrl

`func (o *Issue) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Issue) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Issue) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetNumber

`func (o *Issue) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Issue) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Issue) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *Issue) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Issue) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Issue) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *Issue) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Issue) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Issue) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Issue) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *Issue) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *Issue) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
### GetTitle

`func (o *Issue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Issue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Issue) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *Issue) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Issue) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Issue) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Issue) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Issue) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Issue) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetUser

`func (o *Issue) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Issue) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Issue) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *Issue) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Issue) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetLabels

`func (o *Issue) GetLabels() []IssueLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Issue) GetLabelsOk() (*[]IssueLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Issue) SetLabels(v []IssueLabelsInner)`

SetLabels sets Labels field to given value.


### GetAssignee

`func (o *Issue) GetAssignee() NullableSimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Issue) GetAssigneeOk() (*NullableSimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Issue) SetAssignee(v NullableSimpleUser)`

SetAssignee sets Assignee field to given value.


### SetAssigneeNil

`func (o *Issue) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *Issue) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAssignees

`func (o *Issue) GetAssignees() []SimpleUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *Issue) GetAssigneesOk() (*[]SimpleUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *Issue) SetAssignees(v []SimpleUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *Issue) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### SetAssigneesNil

`func (o *Issue) SetAssigneesNil(b bool)`

 SetAssigneesNil sets the value for Assignees to be an explicit nil

### UnsetAssignees
`func (o *Issue) UnsetAssignees()`

UnsetAssignees ensures that no value is present for Assignees, not even an explicit nil
### GetMilestone

`func (o *Issue) GetMilestone() NullableMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *Issue) GetMilestoneOk() (*NullableMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *Issue) SetMilestone(v NullableMilestone)`

SetMilestone sets Milestone field to given value.


### SetMilestoneNil

`func (o *Issue) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *Issue) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetLocked

`func (o *Issue) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Issue) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Issue) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetActiveLockReason

`func (o *Issue) GetActiveLockReason() string`

GetActiveLockReason returns the ActiveLockReason field if non-nil, zero value otherwise.

### GetActiveLockReasonOk

`func (o *Issue) GetActiveLockReasonOk() (*string, bool)`

GetActiveLockReasonOk returns a tuple with the ActiveLockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLockReason

`func (o *Issue) SetActiveLockReason(v string)`

SetActiveLockReason sets ActiveLockReason field to given value.

### HasActiveLockReason

`func (o *Issue) HasActiveLockReason() bool`

HasActiveLockReason returns a boolean if a field has been set.

### SetActiveLockReasonNil

`func (o *Issue) SetActiveLockReasonNil(b bool)`

 SetActiveLockReasonNil sets the value for ActiveLockReason to be an explicit nil

### UnsetActiveLockReason
`func (o *Issue) UnsetActiveLockReason()`

UnsetActiveLockReason ensures that no value is present for ActiveLockReason, not even an explicit nil
### GetComments

`func (o *Issue) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Issue) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Issue) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetPullRequest

`func (o *Issue) GetPullRequest() IssuePullRequest`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *Issue) GetPullRequestOk() (*IssuePullRequest, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *Issue) SetPullRequest(v IssuePullRequest)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *Issue) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetClosedAt

`func (o *Issue) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *Issue) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *Issue) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *Issue) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *Issue) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetCreatedAt

`func (o *Issue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Issue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Issue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Issue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Issue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Issue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDraft

`func (o *Issue) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *Issue) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *Issue) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *Issue) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetClosedBy

`func (o *Issue) GetClosedBy() NullableSimpleUser`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Issue) GetClosedByOk() (*NullableSimpleUser, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Issue) SetClosedBy(v NullableSimpleUser)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *Issue) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### SetClosedByNil

`func (o *Issue) SetClosedByNil(b bool)`

 SetClosedByNil sets the value for ClosedBy to be an explicit nil

### UnsetClosedBy
`func (o *Issue) UnsetClosedBy()`

UnsetClosedBy ensures that no value is present for ClosedBy, not even an explicit nil
### GetBodyHtml

`func (o *Issue) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *Issue) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *Issue) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *Issue) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *Issue) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *Issue) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *Issue) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *Issue) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetTimelineUrl

`func (o *Issue) GetTimelineUrl() string`

GetTimelineUrl returns the TimelineUrl field if non-nil, zero value otherwise.

### GetTimelineUrlOk

`func (o *Issue) GetTimelineUrlOk() (*string, bool)`

GetTimelineUrlOk returns a tuple with the TimelineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineUrl

`func (o *Issue) SetTimelineUrl(v string)`

SetTimelineUrl sets TimelineUrl field to given value.

### HasTimelineUrl

`func (o *Issue) HasTimelineUrl() bool`

HasTimelineUrl returns a boolean if a field has been set.

### GetRepository

`func (o *Issue) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Issue) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Issue) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Issue) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPerformedViaGithubApp

`func (o *Issue) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *Issue) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *Issue) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *Issue) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *Issue) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *Issue) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetAuthorAssociation

`func (o *Issue) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *Issue) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *Issue) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetReactions

`func (o *Issue) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Issue) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Issue) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *Issue) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


