# TimelineReviewedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Id** | **int32** | Unique identifier of the review | 
**NodeId** | **string** |  | 
**User** | [**SimpleUser**](SimpleUser.md) |  | 
**Body** | **NullableString** | The text of the review. | 
**State** | **string** |  | 
**HtmlUrl** | **string** |  | 
**PullRequestUrl** | **string** |  | 
**Links** | [**TimelineReviewedEventLinks**](TimelineReviewedEventLinks.md) |  | 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**CommitId** | **string** | A commit SHA for the review. | 
**BodyHtml** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**AuthorAssociation** | [**AuthorAssociation**](AuthorAssociation.md) |  | 

## Methods

### NewTimelineReviewedEvent

`func NewTimelineReviewedEvent(event string, id int32, nodeId string, user SimpleUser, body NullableString, state string, htmlUrl string, pullRequestUrl string, links TimelineReviewedEventLinks, commitId string, authorAssociation AuthorAssociation, ) *TimelineReviewedEvent`

NewTimelineReviewedEvent instantiates a new TimelineReviewedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineReviewedEventWithDefaults

`func NewTimelineReviewedEventWithDefaults() *TimelineReviewedEvent`

NewTimelineReviewedEventWithDefaults instantiates a new TimelineReviewedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TimelineReviewedEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineReviewedEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineReviewedEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *TimelineReviewedEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineReviewedEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineReviewedEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TimelineReviewedEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineReviewedEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineReviewedEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUser

`func (o *TimelineReviewedEvent) GetUser() SimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TimelineReviewedEvent) GetUserOk() (*SimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TimelineReviewedEvent) SetUser(v SimpleUser)`

SetUser sets User field to given value.


### GetBody

`func (o *TimelineReviewedEvent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TimelineReviewedEvent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TimelineReviewedEvent) SetBody(v string)`

SetBody sets Body field to given value.


### SetBodyNil

`func (o *TimelineReviewedEvent) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *TimelineReviewedEvent) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetState

`func (o *TimelineReviewedEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TimelineReviewedEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TimelineReviewedEvent) SetState(v string)`

SetState sets State field to given value.


### GetHtmlUrl

`func (o *TimelineReviewedEvent) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TimelineReviewedEvent) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TimelineReviewedEvent) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetPullRequestUrl

`func (o *TimelineReviewedEvent) GetPullRequestUrl() string`

GetPullRequestUrl returns the PullRequestUrl field if non-nil, zero value otherwise.

### GetPullRequestUrlOk

`func (o *TimelineReviewedEvent) GetPullRequestUrlOk() (*string, bool)`

GetPullRequestUrlOk returns a tuple with the PullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUrl

`func (o *TimelineReviewedEvent) SetPullRequestUrl(v string)`

SetPullRequestUrl sets PullRequestUrl field to given value.


### GetLinks

`func (o *TimelineReviewedEvent) GetLinks() TimelineReviewedEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TimelineReviewedEvent) GetLinksOk() (*TimelineReviewedEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TimelineReviewedEvent) SetLinks(v TimelineReviewedEventLinks)`

SetLinks sets Links field to given value.


### GetSubmittedAt

`func (o *TimelineReviewedEvent) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *TimelineReviewedEvent) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *TimelineReviewedEvent) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *TimelineReviewedEvent) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetCommitId

`func (o *TimelineReviewedEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *TimelineReviewedEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *TimelineReviewedEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetBodyHtml

`func (o *TimelineReviewedEvent) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *TimelineReviewedEvent) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *TimelineReviewedEvent) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *TimelineReviewedEvent) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *TimelineReviewedEvent) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *TimelineReviewedEvent) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *TimelineReviewedEvent) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *TimelineReviewedEvent) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetAuthorAssociation

`func (o *TimelineReviewedEvent) GetAuthorAssociation() AuthorAssociation`

GetAuthorAssociation returns the AuthorAssociation field if non-nil, zero value otherwise.

### GetAuthorAssociationOk

`func (o *TimelineReviewedEvent) GetAuthorAssociationOk() (*AuthorAssociation, bool)`

GetAuthorAssociationOk returns a tuple with the AuthorAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAssociation

`func (o *TimelineReviewedEvent) SetAuthorAssociation(v AuthorAssociation)`

SetAuthorAssociation sets AuthorAssociation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


