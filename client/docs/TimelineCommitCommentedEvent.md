# TimelineCommitCommentedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**CommitId** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**[]CommitComment**](CommitComment.md) |  | [optional] 

## Methods

### NewTimelineCommitCommentedEvent

`func NewTimelineCommitCommentedEvent() *TimelineCommitCommentedEvent`

NewTimelineCommitCommentedEvent instantiates a new TimelineCommitCommentedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineCommitCommentedEventWithDefaults

`func NewTimelineCommitCommentedEventWithDefaults() *TimelineCommitCommentedEvent`

NewTimelineCommitCommentedEventWithDefaults instantiates a new TimelineCommitCommentedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TimelineCommitCommentedEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineCommitCommentedEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineCommitCommentedEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TimelineCommitCommentedEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetNodeId

`func (o *TimelineCommitCommentedEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineCommitCommentedEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineCommitCommentedEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *TimelineCommitCommentedEvent) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetCommitId

`func (o *TimelineCommitCommentedEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *TimelineCommitCommentedEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *TimelineCommitCommentedEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *TimelineCommitCommentedEvent) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetComments

`func (o *TimelineCommitCommentedEvent) GetComments() []CommitComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TimelineCommitCommentedEvent) GetCommentsOk() (*[]CommitComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TimelineCommitCommentedEvent) SetComments(v []CommitComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TimelineCommitCommentedEvent) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


