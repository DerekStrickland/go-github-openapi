# TimelineLineCommentedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**[]PullRequestReviewComment**](PullRequestReviewComment.md) |  | [optional] 

## Methods

### NewTimelineLineCommentedEvent

`func NewTimelineLineCommentedEvent() *TimelineLineCommentedEvent`

NewTimelineLineCommentedEvent instantiates a new TimelineLineCommentedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineLineCommentedEventWithDefaults

`func NewTimelineLineCommentedEventWithDefaults() *TimelineLineCommentedEvent`

NewTimelineLineCommentedEventWithDefaults instantiates a new TimelineLineCommentedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TimelineLineCommentedEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineLineCommentedEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineLineCommentedEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TimelineLineCommentedEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetNodeId

`func (o *TimelineLineCommentedEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineLineCommentedEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineLineCommentedEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *TimelineLineCommentedEvent) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetComments

`func (o *TimelineLineCommentedEvent) GetComments() []PullRequestReviewComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TimelineLineCommentedEvent) GetCommentsOk() (*[]PullRequestReviewComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TimelineLineCommentedEvent) SetComments(v []PullRequestReviewComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TimelineLineCommentedEvent) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


