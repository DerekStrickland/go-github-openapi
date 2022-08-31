# TimelineIssueEvents

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
**PerformedViaGithubApp** | [**NullableNullableIntegration**](NullableIntegration.md) |  | 
**Label** | [**LabeledIssueEventLabel**](LabeledIssueEventLabel.md) |  | 
**Milestone** | [**MilestonedIssueEventMilestone**](MilestonedIssueEventMilestone.md) |  | 
**Rename** | [**RenamedIssueEventRename**](RenamedIssueEventRename.md) |  | 
**ReviewRequester** | [**SimpleUser**](SimpleUser.md) |  | 
**RequestedTeam** | Pointer to [**Team**](Team.md) |  | [optional] 
**RequestedReviewer** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**DismissedReview** | [**ReviewDismissedIssueEventDismissedReview**](ReviewDismissedIssueEventDismissedReview.md) |  | 
**LockReason** | **NullableString** |  | 
**ProjectCard** | Pointer to [**AddedToProjectIssueEventProjectCard**](AddedToProjectIssueEventProjectCard.md) |  | [optional] 
**Body** | **NullableString** | The text of the review. | 
**BodyText** | Pointer to **string** |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**HtmlUrl** | **string** |  | 
**User** | [**SimpleUser**](SimpleUser.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**IssueUrl** | **string** |  | 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 
**Source** | [**TimelineCrossReferencedEventSource**](TimelineCrossReferencedEventSource.md) |  | 
**Sha** | **string** | SHA for the commit | 
**Author** | [**GitCommitAuthor**](GitCommitAuthor.md) |  | 
**Committer** | [**GitCommitAuthor**](GitCommitAuthor.md) |  | 
**Message** | **string** | Message describing the purpose of the commit | 
**Tree** | [**GitCommitTree**](GitCommitTree.md) |  | 
**Parents** | [**[]GitCommitParentsInner**](GitCommitParentsInner.md) |  | 
**Verification** | [**GitCommitVerification**](GitCommitVerification.md) |  | 
**State** | **string** |  | 
**PullRequestUrl** | **string** |  | 
**Links** | [**TimelineReviewedEventLinks**](TimelineReviewedEventLinks.md) |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**Comments** | Pointer to [**[]CommitComment**](CommitComment.md) |  | [optional] 
**Assignee** | [**SimpleUser**](SimpleUser.md) |  | 
**StateReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTimelineIssueEvents

`func NewTimelineIssueEvents(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, label LabeledIssueEventLabel, milestone MilestonedIssueEventMilestone, rename RenamedIssueEventRename, reviewRequester SimpleUser, dismissedReview ReviewDismissedIssueEventDismissedReview, lockReason NullableString, body NullableString, htmlUrl string, user SimpleUser, updatedAt time.Time, issueUrl string, authorAssociation AuthorAssociation, source TimelineCrossReferencedEventSource, sha string, author GitCommitAuthor, committer GitCommitAuthor, message string, tree GitCommitTree, parents []GitCommitParentsInner, verification GitCommitVerification, state string, pullRequestUrl string, links TimelineReviewedEventLinks, assignee SimpleUser, ) *TimelineIssueEvents`

NewTimelineIssueEvents instantiates a new TimelineIssueEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineIssueEventsWithDefaults

`func NewTimelineIssueEventsWithDefaults() *TimelineIssueEvents`

NewTimelineIssueEventsWithDefaults instantiates a new TimelineIssueEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimelineIssueEvents) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineIssueEvents) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineIssueEvents) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TimelineIssueEvents) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineIssueEvents) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineIssueEvents) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TimelineIssueEvents) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TimelineIssueEvents) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TimelineIssueEvents) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *TimelineIssueEvents) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TimelineIssueEvents) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TimelineIssueEvents) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *TimelineIssueEvents) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineIssueEvents) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineIssueEvents) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *TimelineIssueEvents) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *TimelineIssueEvents) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *TimelineIssueEvents) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *TimelineIssueEvents) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *TimelineIssueEvents) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *TimelineIssueEvents) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *TimelineIssueEvents) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *TimelineIssueEvents) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *TimelineIssueEvents) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *TimelineIssueEvents) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *TimelineIssueEvents) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineIssueEvents) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineIssueEvents) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *TimelineIssueEvents) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *TimelineIssueEvents) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *TimelineIssueEvents) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *TimelineIssueEvents) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *TimelineIssueEvents) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetLabel

`func (o *TimelineIssueEvents) GetLabel() LabeledIssueEventLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TimelineIssueEvents) GetLabelOk() (*LabeledIssueEventLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TimelineIssueEvents) SetLabel(v LabeledIssueEventLabel)`

SetLabel sets Label field to given value.


### GetMilestone

`func (o *TimelineIssueEvents) GetMilestone() MilestonedIssueEventMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *TimelineIssueEvents) GetMilestoneOk() (*MilestonedIssueEventMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *TimelineIssueEvents) SetMilestone(v MilestonedIssueEventMilestone)`

SetMilestone sets Milestone field to given value.


### GetRename

`func (o *TimelineIssueEvents) GetRename() RenamedIssueEventRename`

GetRename returns the Rename field if non-nil, zero value otherwise.

### GetRenameOk

`func (o *TimelineIssueEvents) GetRenameOk() (*RenamedIssueEventRename, bool)`

GetRenameOk returns a tuple with the Rename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRename

`func (o *TimelineIssueEvents) SetRename(v RenamedIssueEventRename)`

SetRename sets Rename field to given value.


### GetReviewRequester

`func (o *TimelineIssueEvents) GetReviewRequester() SimpleUser`

GetReviewRequester returns the ReviewRequester field if non-nil, zero value otherwise.

### GetReviewRequesterOk

`func (o *TimelineIssueEvents) GetReviewRequesterOk() (*SimpleUser, bool)`

GetReviewRequesterOk returns a tuple with the ReviewRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequester

`func (o *TimelineIssueEvents) SetReviewRequester(v SimpleUser)`

SetReviewRequester sets ReviewRequester field to given value.


### GetRequestedTeam

`func (o *TimelineIssueEvents) GetRequestedTeam() Team`

GetRequestedTeam returns the RequestedTeam field if non-nil, zero value otherwise.

### GetRequestedTeamOk

`func (o *TimelineIssueEvents) GetRequestedTeamOk() (*Team, bool)`

GetRequestedTeamOk returns a tuple with the RequestedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeam

`func (o *TimelineIssueEvents) SetRequestedTeam(v Team)`

SetRequestedTeam sets RequestedTeam field to given value.

### HasRequestedTeam

`func (o *TimelineIssueEvents) HasRequestedTeam() bool`

HasRequestedTeam returns a boolean if a field has been set.

### GetRequestedReviewer

`func (o *TimelineIssueEvents) GetRequestedReviewer() SimpleUser`

GetRequestedReviewer returns the RequestedReviewer field if non-nil, zero value otherwise.

### GetRequestedReviewerOk

`func (o *TimelineIssueEvents) GetRequestedReviewerOk() (*SimpleUser, bool)`

GetRequestedReviewerOk returns a tuple with the RequestedReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewer

`func (o *TimelineIssueEvents) SetRequestedReviewer(v SimpleUser)`

SetRequestedReviewer sets RequestedReviewer field to given value.

### HasRequestedReviewer

`func (o *TimelineIssueEvents) HasRequestedReviewer() bool`

HasRequestedReviewer returns a boolean if a field has been set.

### GetDismissedReview

`func (o *TimelineIssueEvents) GetDismissedReview() ReviewDismissedIssueEventDismissedReview`

GetDismissedReview returns the DismissedReview field if non-nil, zero value otherwise.

### GetDismissedReviewOk

`func (o *TimelineIssueEvents) GetDismissedReviewOk() (*ReviewDismissedIssueEventDismissedReview, bool)`

GetDismissedReviewOk returns a tuple with the DismissedReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedReview

`func (o *TimelineIssueEvents) SetDismissedReview(v ReviewDismissedIssueEventDismissedReview)`

SetDismissedReview sets DismissedReview field to given value.


### GetLockReason

`func (o *TimelineIssueEvents) GetLockReason() string`

GetLockReason returns the LockReason field if non-nil, zero value otherwise.

### GetLockReasonOk

`func (o *TimelineIssueEvents) GetLockReasonOk() (*string, bool)`

GetLockReasonOk returns a tuple with the LockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockReason

`func (o *TimelineIssueEvents) SetLockReason(v string)`

SetLockReason sets LockReason field to given value.


### SetLockReasonNil

`func (o *TimelineIssueEvents) SetLockReasonNil(b bool)`

 SetLockReasonNil sets the value for LockReason to be an explicit nil

### UnsetLockReason
`func (o *TimelineIssueEvents) UnsetLockReason()`

UnsetLockReason ensures that no value is present for LockReason, not even an explicit nil
### GetProjectCard

`func (o *TimelineIssueEvents) GetProjectCard() AddedToProjectIssueEventProjectCard`

GetProjectCard returns the ProjectCard field if non-nil, zero value otherwise.

### GetProjectCardOk

`func (o *TimelineIssueEvents) GetProjectCardOk() (*AddedToProjectIssueEventProjectCard, bool)`

GetProjectCardOk returns a tuple with the ProjectCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCard

`func (o *TimelineIssueEvents) SetProjectCard(v AddedToProjectIssueEventProjectCard)`

SetProjectCard sets ProjectCard field to given value.

### HasProjectCard

`func (o *TimelineIssueEvents) HasProjectCard() bool`

HasProjectCard returns a boolean if a field has been set.

### GetBody

`func (o *TimelineIssueEvents) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TimelineIssueEvents) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TimelineIssueEvents) SetBody(v string)`

SetBody sets Body field to given value.


### SetBodyNil

`func (o *TimelineIssueEvents) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *TimelineIssueEvents) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyText

`func (o *TimelineIssueEvents) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *TimelineIssueEvents) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *TimelineIssueEvents) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *TimelineIssueEvents) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBodyHtml

`func (o *TimelineIssueEvents) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *TimelineIssueEvents) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *TimelineIssueEvents) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *TimelineIssueEvents) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *TimelineIssueEvents) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TimelineIssueEvents) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TimelineIssueEvents) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUser

`func (o *TimelineIssueEvents) GetUser() SimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TimelineIssueEvents) GetUserOk() (*SimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TimelineIssueEvents) SetUser(v SimpleUser)`

SetUser sets User field to given value.


### GetUpdatedAt

`func (o *TimelineIssueEvents) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimelineIssueEvents) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimelineIssueEvents) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIssueUrl

`func (o *TimelineIssueEvents) GetIssueUrl() string`

GetIssueUrl returns the IssueUrl field if non-nil, zero value otherwise.

### GetIssueUrlOk

`func (o *TimelineIssueEvents) GetIssueUrlOk() (*string, bool)`

GetIssueUrlOk returns a tuple with the IssueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueUrl

`func (o *TimelineIssueEvents) SetIssueUrl(v string)`

SetIssueUrl sets IssueUrl field to given value.


### GetAuthorAssociation

`func (o *TimelineIssueEvents) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *TimelineIssueEvents) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *TimelineIssueEvents) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.


### GetReactions

`func (o *TimelineIssueEvents) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *TimelineIssueEvents) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *TimelineIssueEvents) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *TimelineIssueEvents) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetSource

`func (o *TimelineIssueEvents) GetSource() TimelineCrossReferencedEventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TimelineIssueEvents) GetSourceOk() (*TimelineCrossReferencedEventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TimelineIssueEvents) SetSource(v TimelineCrossReferencedEventSource)`

SetSource sets Source field to given value.


### GetSha

`func (o *TimelineIssueEvents) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *TimelineIssueEvents) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *TimelineIssueEvents) SetSha(v string)`

SetSha sets Sha field to given value.


### GetAuthor

`func (o *TimelineIssueEvents) GetAuthor() GitCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TimelineIssueEvents) GetAuthorOk() (*GitCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TimelineIssueEvents) SetAuthor(v GitCommitAuthor)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *TimelineIssueEvents) GetCommitter() GitCommitAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *TimelineIssueEvents) GetCommitterOk() (*GitCommitAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *TimelineIssueEvents) SetCommitter(v GitCommitAuthor)`

SetCommitter sets Committer field to given value.


### GetMessage

`func (o *TimelineIssueEvents) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TimelineIssueEvents) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TimelineIssueEvents) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTree

`func (o *TimelineIssueEvents) GetTree() GitCommitTree`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *TimelineIssueEvents) GetTreeOk() (*GitCommitTree, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *TimelineIssueEvents) SetTree(v GitCommitTree)`

SetTree sets Tree field to given value.


### GetParents

`func (o *TimelineIssueEvents) GetParents() []GitCommitParentsInner`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *TimelineIssueEvents) GetParentsOk() (*[]GitCommitParentsInner, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *TimelineIssueEvents) SetParents(v []GitCommitParentsInner)`

SetParents sets Parents field to given value.


### GetVerification

`func (o *TimelineIssueEvents) GetVerification() GitCommitVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *TimelineIssueEvents) GetVerificationOk() (*GitCommitVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *TimelineIssueEvents) SetVerification(v GitCommitVerification)`

SetVerification sets Verification field to given value.


### GetState

`func (o *TimelineIssueEvents) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TimelineIssueEvents) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TimelineIssueEvents) SetState(v string)`

SetState sets State field to given value.


### GetPullRequestUrl

`func (o *TimelineIssueEvents) GetPullRequestUrl() string`

GetPullRequestUrl returns the PullRequestUrl field if non-nil, zero value otherwise.

### GetPullRequestUrlOk

`func (o *TimelineIssueEvents) GetPullRequestUrlOk() (*string, bool)`

GetPullRequestUrlOk returns a tuple with the PullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUrl

`func (o *TimelineIssueEvents) SetPullRequestUrl(v string)`

SetPullRequestUrl sets PullRequestUrl field to given value.


### GetLinks

`func (o *TimelineIssueEvents) GetLinks() TimelineReviewedEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TimelineIssueEvents) GetLinksOk() (*TimelineReviewedEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TimelineIssueEvents) SetLinks(v TimelineReviewedEventLinks)`

SetLinks sets Links field to given value.


### GetSubmittedAt

`func (o *TimelineIssueEvents) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *TimelineIssueEvents) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *TimelineIssueEvents) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *TimelineIssueEvents) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetComments

`func (o *TimelineIssueEvents) GetComments() []CommitComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TimelineIssueEvents) GetCommentsOk() (*[]CommitComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TimelineIssueEvents) SetComments(v []CommitComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TimelineIssueEvents) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAssignee

`func (o *TimelineIssueEvents) GetAssignee() SimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TimelineIssueEvents) GetAssigneeOk() (*SimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TimelineIssueEvents) SetAssignee(v SimpleUser)`

SetAssignee sets Assignee field to given value.


### GetStateReason

`func (o *TimelineIssueEvents) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *TimelineIssueEvents) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *TimelineIssueEvents) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *TimelineIssueEvents) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *TimelineIssueEvents) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *TimelineIssueEvents) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


