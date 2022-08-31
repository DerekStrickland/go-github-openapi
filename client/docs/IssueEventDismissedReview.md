# IssueEventDismissedReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**ReviewId** | **int32** |  | 
**DismissalMessage** | **NullableString** |  | 
**DismissalCommitId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIssueEventDismissedReview

`func NewIssueEventDismissedReview(state string, reviewId int32, dismissalMessage NullableString, ) *IssueEventDismissedReview`

NewIssueEventDismissedReview instantiates a new IssueEventDismissedReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEventDismissedReviewWithDefaults

`func NewIssueEventDismissedReviewWithDefaults() *IssueEventDismissedReview`

NewIssueEventDismissedReviewWithDefaults instantiates a new IssueEventDismissedReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *IssueEventDismissedReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IssueEventDismissedReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IssueEventDismissedReview) SetState(v string)`

SetState sets State field to given value.


### GetReviewId

`func (o *IssueEventDismissedReview) GetReviewId() int32`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *IssueEventDismissedReview) GetReviewIdOk() (*int32, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *IssueEventDismissedReview) SetReviewId(v int32)`

SetReviewId sets ReviewId field to given value.


### GetDismissalMessage

`func (o *IssueEventDismissedReview) GetDismissalMessage() string`

GetDismissalMessage returns the DismissalMessage field if non-nil, zero value otherwise.

### GetDismissalMessageOk

`func (o *IssueEventDismissedReview) GetDismissalMessageOk() (*string, bool)`

GetDismissalMessageOk returns a tuple with the DismissalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalMessage

`func (o *IssueEventDismissedReview) SetDismissalMessage(v string)`

SetDismissalMessage sets DismissalMessage field to given value.


### SetDismissalMessageNil

`func (o *IssueEventDismissedReview) SetDismissalMessageNil(b bool)`

 SetDismissalMessageNil sets the value for DismissalMessage to be an explicit nil

### UnsetDismissalMessage
`func (o *IssueEventDismissedReview) UnsetDismissalMessage()`

UnsetDismissalMessage ensures that no value is present for DismissalMessage, not even an explicit nil
### GetDismissalCommitId

`func (o *IssueEventDismissedReview) GetDismissalCommitId() string`

GetDismissalCommitId returns the DismissalCommitId field if non-nil, zero value otherwise.

### GetDismissalCommitIdOk

`func (o *IssueEventDismissedReview) GetDismissalCommitIdOk() (*string, bool)`

GetDismissalCommitIdOk returns a tuple with the DismissalCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalCommitId

`func (o *IssueEventDismissedReview) SetDismissalCommitId(v string)`

SetDismissalCommitId sets DismissalCommitId field to given value.

### HasDismissalCommitId

`func (o *IssueEventDismissedReview) HasDismissalCommitId() bool`

HasDismissalCommitId returns a boolean if a field has been set.

### SetDismissalCommitIdNil

`func (o *IssueEventDismissedReview) SetDismissalCommitIdNil(b bool)`

 SetDismissalCommitIdNil sets the value for DismissalCommitId to be an explicit nil

### UnsetDismissalCommitId
`func (o *IssueEventDismissedReview) UnsetDismissalCommitId()`

UnsetDismissalCommitId ensures that no value is present for DismissalCommitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


