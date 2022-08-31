# IssueSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**RepositoryUrl** | **string** |  | 
**LabelsUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**EventsUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Number** | **int32** |  | 
**Title** | **string** |  | 
**Locked** | **bool** |  | 
**ActiveLockReason** | Pointer to **NullableString** |  | [optional] 
**Assignees** | Pointer to [**[]SimpleUser**](SimpleUser.md) |  | [optional] 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Labels** | [**[]IssueSearchResultItemLabelsInner**](IssueSearchResultItemLabelsInner.md) |  | 
**State** | **string** |  | 
**StateReason** | Pointer to **NullableString** |  | [optional] 
**Assignee** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Milestone** | [**NullableNullableMilestone**](NullableMilestone.md) |  | 
**Comments** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ClosedAt** | **NullableTime** |  | 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 
**PullRequest** | Pointer to [**IssuePullRequest**](IssuePullRequest.md) |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Score** | **float32** |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**Draft** | Pointer to **bool** |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**TimelineUrl** | Pointer to **string** |  | [optional] 
**PerformedViaGithubApp** | Pointer to [**NullableNullableIntegration**](NullableIntegration.md) |  | [optional] 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewIssueSearchResultItem

`func NewIssueSearchResultItem(url string, repositoryUrl string, labelsUrl string, commentsUrl string, eventsUrl string, htmlUrl string, id int32, nodeId string, number int32, title string, locked bool, user NullableNullableSimpleUser, labels []IssueSearchResultItemLabelsInner, state string, assignee NullableNullableSimpleUser, milestone NullableNullableMilestone, comments int32, createdAt time.Time, updatedAt time.Time, closedAt NullableTime, score float32, authorAssociation AuthorAssociation, ) *IssueSearchResultItem`

NewIssueSearchResultItem instantiates a new IssueSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueSearchResultItemWithDefaults

`func NewIssueSearchResultItemWithDefaults() *IssueSearchResultItem`

NewIssueSearchResultItemWithDefaults instantiates a new IssueSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IssueSearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssueSearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssueSearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRepositoryUrl

`func (o *IssueSearchResultItem) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *IssueSearchResultItem) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *IssueSearchResultItem) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetLabelsUrl

`func (o *IssueSearchResultItem) GetLabelsUrl() string`

GetLabelsUrl returns the LabelsUrl field if non-nil, zero value otherwise.

### GetLabelsUrlOk

`func (o *IssueSearchResultItem) GetLabelsUrlOk() (*string, bool)`

GetLabelsUrlOk returns a tuple with the LabelsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsUrl

`func (o *IssueSearchResultItem) SetLabelsUrl(v string)`

SetLabelsUrl sets LabelsUrl field to given value.


### GetCommentsUrl

`func (o *IssueSearchResultItem) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *IssueSearchResultItem) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *IssueSearchResultItem) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetEventsUrl

`func (o *IssueSearchResultItem) GetEventsUrl() string`

GetEventsUrl returns the EventsUrl field if non-nil, zero value otherwise.

### GetEventsUrlOk

`func (o *IssueSearchResultItem) GetEventsUrlOk() (*string, bool)`

GetEventsUrlOk returns a tuple with the EventsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUrl

`func (o *IssueSearchResultItem) SetEventsUrl(v string)`

SetEventsUrl sets EventsUrl field to given value.


### GetHtmlUrl

`func (o *IssueSearchResultItem) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *IssueSearchResultItem) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *IssueSearchResultItem) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetId

`func (o *IssueSearchResultItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueSearchResultItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueSearchResultItem) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *IssueSearchResultItem) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *IssueSearchResultItem) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *IssueSearchResultItem) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNumber

`func (o *IssueSearchResultItem) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IssueSearchResultItem) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IssueSearchResultItem) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetTitle

`func (o *IssueSearchResultItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueSearchResultItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueSearchResultItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLocked

`func (o *IssueSearchResultItem) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *IssueSearchResultItem) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *IssueSearchResultItem) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetActiveLockReason

`func (o *IssueSearchResultItem) GetActiveLockReason() string`

GetActiveLockReason returns the ActiveLockReason field if non-nil, zero value otherwise.

### GetActiveLockReasonOk

`func (o *IssueSearchResultItem) GetActiveLockReasonOk() (*string, bool)`

GetActiveLockReasonOk returns a tuple with the ActiveLockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLockReason

`func (o *IssueSearchResultItem) SetActiveLockReason(v string)`

SetActiveLockReason sets ActiveLockReason field to given value.

### HasActiveLockReason

`func (o *IssueSearchResultItem) HasActiveLockReason() bool`

HasActiveLockReason returns a boolean if a field has been set.

### SetActiveLockReasonNil

`func (o *IssueSearchResultItem) SetActiveLockReasonNil(b bool)`

 SetActiveLockReasonNil sets the value for ActiveLockReason to be an explicit nil

### UnsetActiveLockReason
`func (o *IssueSearchResultItem) UnsetActiveLockReason()`

UnsetActiveLockReason ensures that no value is present for ActiveLockReason, not even an explicit nil
### GetAssignees

`func (o *IssueSearchResultItem) GetAssignees() []SimpleUser`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *IssueSearchResultItem) GetAssigneesOk() (*[]SimpleUser, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *IssueSearchResultItem) SetAssignees(v []SimpleUser)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *IssueSearchResultItem) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### SetAssigneesNil

`func (o *IssueSearchResultItem) SetAssigneesNil(b bool)`

 SetAssigneesNil sets the value for Assignees to be an explicit nil

### UnsetAssignees
`func (o *IssueSearchResultItem) UnsetAssignees()`

UnsetAssignees ensures that no value is present for Assignees, not even an explicit nil
### GetUser

`func (o *IssueSearchResultItem) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IssueSearchResultItem) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IssueSearchResultItem) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *IssueSearchResultItem) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IssueSearchResultItem) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetLabels

`func (o *IssueSearchResultItem) GetLabels() []IssueSearchResultItemLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssueSearchResultItem) GetLabelsOk() (*[]IssueSearchResultItemLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssueSearchResultItem) SetLabels(v []IssueSearchResultItemLabelsInner)`

SetLabels sets Labels field to given value.


### GetState

`func (o *IssueSearchResultItem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IssueSearchResultItem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IssueSearchResultItem) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *IssueSearchResultItem) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *IssueSearchResultItem) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *IssueSearchResultItem) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *IssueSearchResultItem) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *IssueSearchResultItem) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *IssueSearchResultItem) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
### GetAssignee

`func (o *IssueSearchResultItem) GetAssignee() NullableSimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssueSearchResultItem) GetAssigneeOk() (*NullableSimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssueSearchResultItem) SetAssignee(v NullableSimpleUser)`

SetAssignee sets Assignee field to given value.


### SetAssigneeNil

`func (o *IssueSearchResultItem) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *IssueSearchResultItem) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetMilestone

`func (o *IssueSearchResultItem) GetMilestone() NullableMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *IssueSearchResultItem) GetMilestoneOk() (*NullableMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *IssueSearchResultItem) SetMilestone(v NullableMilestone)`

SetMilestone sets Milestone field to given value.


### SetMilestoneNil

`func (o *IssueSearchResultItem) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *IssueSearchResultItem) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetComments

`func (o *IssueSearchResultItem) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IssueSearchResultItem) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IssueSearchResultItem) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetCreatedAt

`func (o *IssueSearchResultItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueSearchResultItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueSearchResultItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IssueSearchResultItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueSearchResultItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueSearchResultItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetClosedAt

`func (o *IssueSearchResultItem) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *IssueSearchResultItem) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *IssueSearchResultItem) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *IssueSearchResultItem) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *IssueSearchResultItem) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetTextMatches

`func (o *IssueSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *IssueSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *IssueSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *IssueSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.

### GetPullRequest

`func (o *IssueSearchResultItem) GetPullRequest() IssuePullRequest`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *IssueSearchResultItem) GetPullRequestOk() (*IssuePullRequest, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *IssueSearchResultItem) SetPullRequest(v IssuePullRequest)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *IssueSearchResultItem) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetBody

`func (o *IssueSearchResultItem) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IssueSearchResultItem) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IssueSearchResultItem) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IssueSearchResultItem) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetScore

`func (o *IssueSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IssueSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IssueSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetAuthorAssociation

`func (o *IssueSearchResultItem) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *IssueSearchResultItem) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *IssueSearchResultItem) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetDraft

`func (o *IssueSearchResultItem) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *IssueSearchResultItem) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *IssueSearchResultItem) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *IssueSearchResultItem) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetRepository

`func (o *IssueSearchResultItem) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *IssueSearchResultItem) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *IssueSearchResultItem) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *IssueSearchResultItem) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBodyHtml

`func (o *IssueSearchResultItem) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *IssueSearchResultItem) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *IssueSearchResultItem) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *IssueSearchResultItem) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *IssueSearchResultItem) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *IssueSearchResultItem) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *IssueSearchResultItem) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *IssueSearchResultItem) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetTimelineUrl

`func (o *IssueSearchResultItem) GetTimelineUrl() string`

GetTimelineUrl returns the TimelineUrl field if non-nil, zero value otherwise.

### GetTimelineUrlOk

`func (o *IssueSearchResultItem) GetTimelineUrlOk() (*string, bool)`

GetTimelineUrlOk returns a tuple with the TimelineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineUrl

`func (o *IssueSearchResultItem) SetTimelineUrl(v string)`

SetTimelineUrl sets TimelineUrl field to given value.

### HasTimelineUrl

`func (o *IssueSearchResultItem) HasTimelineUrl() bool`

HasTimelineUrl returns a boolean if a field has been set.

### GetPerformedViaGithubApp

`func (o *IssueSearchResultItem) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *IssueSearchResultItem) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *IssueSearchResultItem) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.

### HasPerformedViaGithubApp

`func (o *IssueSearchResultItem) HasPerformedViaGithubApp() bool`

HasPerformedViaGithubApp returns a boolean if a field has been set.

### SetPerformedViaGithubAppNil

`func (o *IssueSearchResultItem) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *IssueSearchResultItem) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetReactions

`func (o *IssueSearchResultItem) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *IssueSearchResultItem) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *IssueSearchResultItem) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *IssueSearchResultItem) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


