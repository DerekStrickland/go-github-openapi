# PullsDismissReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The message for the pull request review dismissal | 
**Event** | Pointer to **string** |  | [optional] 

## Methods

### NewPullsDismissReviewRequest

`func NewPullsDismissReviewRequest(message string, ) *PullsDismissReviewRequest`

NewPullsDismissReviewRequest instantiates a new PullsDismissReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsDismissReviewRequestWithDefaults

`func NewPullsDismissReviewRequestWithDefaults() *PullsDismissReviewRequest`

NewPullsDismissReviewRequestWithDefaults instantiates a new PullsDismissReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PullsDismissReviewRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PullsDismissReviewRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PullsDismissReviewRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEvent

`func (o *PullsDismissReviewRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PullsDismissReviewRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PullsDismissReviewRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PullsDismissReviewRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


