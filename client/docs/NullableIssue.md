# NullableIssue

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

### NewNullableIssue

`func NewNullableIssue(id int32, nodeId string, url string, repositoryUrl string, labelsUrl string, commentsUrl string, eventsUrl string, htmlUrl string, number int32, state string, title string, user NullableNullableSimpleUser, labels []IssueLabelsInner, assignee NullableNullableSimpleUser, milestone NullableNullableMilestone, locked bool, comments int32, closedAt NullableTime, createdAt time.Time, updatedAt time.Time, authorAssociation AuthorAssociation, ) *NullableIssue`

NewNullableIssue instantiates a new NullableIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableIssueWithDefaults

`func NewNullableIssueWithDefaults() *NullableIssue`

NewNullableIssueWithDefaults instantiates a new NullableIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NullableIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableIssue) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *NullableIssue) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NullableIssue) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NullableIssue) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *NullableIssue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NullableIssue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NullableIssue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRepositoryUrl

`func (o *NullableIssue) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *NullableIssue) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *NullableIssue) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetLabelsUrl

`func (o *NullableIssue) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *NullableIssue) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *NullableIssue) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetCommentsUrl

`func (o *NullableIssue) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *NullableIssue) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *NullableIssue) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetEventsUrl

`func (o *NullableIssue) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *NullableIssue) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *NullableIssue) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetHtmlUrl

`func (o *NullableIssue) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *NullableIssue) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *NullableIssue) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetNumber

`func (o *NullableIssue) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NullableIssue) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NullableIssue) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *NullableIssue) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NullableIssue) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NullableIssue) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *NullableIssue) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *NullableIssue) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *NullableIssue) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *NullableIssue) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *NullableIssue) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *NullableIssue) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
### GetTitle

`func (o *NullableIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NullableIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NullableIssue) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *NullableIssue) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NullableIssue) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NullableIssue) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NullableIssue) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *NullableIssue) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *NullableIssue) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetUser

`func (o *NullableIssue) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NullableIssue) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NullableIssue) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *NullableIssue) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *NullableIssue) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetLabels

`func (o *NullableIssue) GetLabels() []IssueLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NullableIssue) GetLabelsOk() (*[]IssueLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NullableIssue) SetLabels(v []IssueLabelsInner)`

SetLabels sets Labels field to given value.


### GetAssignee

`func (o *NullableIssue) GetAssignee() NullableSimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *NullableIssue) GetAssigneeOk() (*NullableSimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *NullableIssue) SetAssignee(v NullableSimpleUser)`

SetAssignee sets Assignee field to given value.


### SetAssigneeNil

`func (o *NullableIssue) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *NullableIssue) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetAssignees

`func (o *NullableIssue) GetAssignees() []SimpleUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *NullableIssue) GetAssigneesOk() (*[]SimpleUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *NullableIssue) SetAssignees(v []SimpleUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *NullableIssue) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### SetAssigneesNil

`func (o *NullableIssue) SetAssigneesNil(b bool)`

 SetAssigneesNil sets the value for Assignees to be an explicit nil

### UnsetAssignees
`func (o *NullableIssue) UnsetAssignees()`

UnsetAssignees ensures that no value is present for Assignees, not even an explicit nil
### GetMilestone

`func (o *NullableIssue) GetMilestone() NullableMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *NullableIssue) GetMilestoneOk() (*NullableMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *NullableIssue) SetMilestone(v NullableMilestone)`

SetMilestone sets Milestone field to given value.


### SetMilestoneNil

`func (o *NullableIssue) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *NullableIssue) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetLocked

`func (o *NullableIssue) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *NullableIssue) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *NullableIssue) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetActiveLockReason

`func (o *NullableIssue) GetActiveLockReason() string`

GetActiveLockReason returns the ActiveLockReason field if non-nil, zero value otherwise.

### GetActiveLockReasonOk

`func (o *NullableIssue) GetActiveLockReasonOk() (*string, bool)`

GetActiveLockReasonOk returns a tuple with the ActiveLockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLockReason

`func (o *NullableIssue) SetActiveLockReason(v string)`

SetActiveLockReason sets ActiveLockReason field to given value.

### HasActiveLockReason

`func (o *NullableIssue) HasActiveLockReason() bool`

HasActiveLockReason returns a boolean if a field has been set.

### SetActiveLockReasonNil

`func (o *NullableIssue) SetActiveLockReasonNil(b bool)`

 SetActiveLockReasonNil sets the value for ActiveLockReason to be an explicit nil

### UnsetActiveLockReason
`func (o *NullableIssue) UnsetActiveLockReason()`

UnsetActiveLockReason ensures that no value is present for ActiveLockReason, not even an explicit nil
### GetComments

`func (o *NullableIssue) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *NullableIssue) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *NullableIssue) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetPullRequest

`func (o *NullableIssue) GetPullRequest() IssuePullRequest`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *NullableIssue) GetPullRequestOk() (*IssuePullRequest, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *NullableIssue) SetPullRequest(v IssuePullRequest)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *NullableIssue) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetClosedAt

`func (o *NullableIssue) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *NullableIssue) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *NullableIssue) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *NullableIssue) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *NullableIssue) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetCreatedAt

`func (o *NullableIssue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NullableIssue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NullableIssue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NullableIssue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NullableIssue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NullableIssue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDraft

`func (o *NullableIssue) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *NullableIssue) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *NullableIssue) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *NullableIssue) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetClosedBy

`func (o *NullableIssue) GetClosedBy() NullableSimpleUser`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *NullableIssue) GetClosedByOk() (*NullableSimpleUser, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *NullableIssue) SetClosedBy(v NullableSimpleUser)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *NullableIssue) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### SetClosedByNil

`func (o *NullableIssue) SetClosedByNil(b bool)`

 SetClosedByNil sets the value for ClosedBy to be an explicit nil

### UnsetClosedBy
`func (o *NullableIssue) UnsetClosedBy()`

UnsetClosedBy ensures that no value is present for ClosedBy, not even an explicit nil
### GetBodyHtml

`func (o *NullableIssue) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *NullableIssue) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *NullableIssue) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *NullableIssue) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *NullableIssue) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *NullableIssue) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *NullableIssue) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *NullableIssue) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetTimelineUrl

`func (o *NullableIssue) GetTimelineUrl() string`

GetTimelineUrl returns the TimelineUrl field if non-nil, zero value otherwise.

### GetTimelineUrlOk

`func (o *NullableIssue) GetTimelineUrlOk() (*string, bool)`

GetTimelineUrlOk returns a tuple with the TimelineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineUrl

`func (o *NullableIssue) SetTimelineUrl(v string)`

SetTimelineUrl sets TimelineUrl field to given value.

### HasTimelineUrl

`func (o *NullableIssue) HasTimelineUrl() bool`

HasTimelineUrl returns a boolean if a field has been set.

### GetRepository

`func (o *NullableIssue) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NullableIssue) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NullableIssue) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *NullableIssue) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPerformedViaGithubApp

`func (o *NullableIssue) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *NullableIssue) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *NullableIssue) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *NullableIssue) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *NullableIssue) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *NullableIssue) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetAuthorAssociation

`func (o *NullableIssue) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *NullableIssue) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *NullableIssue) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetReactions

`func (o *NullableIssue) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *NullableIssue) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *NullableIssue) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *NullableIssue) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


