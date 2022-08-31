# ReviewDismissedIssueEventDismissedReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**ReviewId** | **int32** |  | 
**DismissalMessage** | **NullableString** |  | 
**DismissalCommitId** | Pointer to **string** |  | [optional] 

## Methods

### NewReviewDismissedIssueEventDismissedReview

`func NewReviewDismissedIssueEventDismissedReview(state string, reviewId int32, dismissalMessage NullableString, ) *ReviewDismissedIssueEventDismissedReview`

NewReviewDismissedIssueEventDismissedReview instantiates a new ReviewDismissedIssueEventDismissedReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewDismissedIssueEventDismissedReviewWithDefaults

`func NewReviewDismissedIssueEventDismissedReviewWithDefaults() *ReviewDismissedIssueEventDismissedReview`

NewReviewDismissedIssueEventDismissedReviewWithDefaults instantiates a new ReviewDismissedIssueEventDismissedReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ReviewDismissedIssueEventDismissedReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReviewDismissedIssueEventDismissedReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReviewDismissedIssueEventDismissedReview) SetState(v string)`

SetState sets State field to given value.


### GetReviewId

`func (o *ReviewDismissedIssueEventDismissedReview) GetReviewId() int32`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *ReviewDismissedIssueEventDismissedReview) GetReviewIdOk() (*int32, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *ReviewDismissedIssueEventDismissedReview) SetReviewId(v int32)`

SetReviewId sets ReviewId field to given value.


### GetDismissalMessage

`func (o *ReviewDismissedIssueEventDismissedReview) GetDismissalMessage() string`

GetDismissalMessage returns the DismissalMessage field if non-nil, zero value otherwise.

### GetDismissalMessageOk

`func (o *ReviewDismissedIssueEventDismissedReview) GetDismissalMessageOk() (*string, bool)`

GetDismissalMessageOk returns a tuple with the DismissalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalMessage

`func (o *ReviewDismissedIssueEventDismissedReview) SetDismissalMessage(v string)`

SetDismissalMessage sets DismissalMessage field to given value.


### SetDismissalMessageNil

`func (o *ReviewDismissedIssueEventDismissedReview) SetDismissalMessageNil(b bool)`

 SetDismissalMessageNil sets the value for DismissalMessage to be an explicit nil

### UnsetDismissalMessage
`func (o *ReviewDismissedIssueEventDismissedReview) UnsetDismissalMessage()`

UnsetDismissalMessage ensures that no value is present for DismissalMessage, not even an explicit nil
### GetDismissalCommitId

`func (o *ReviewDismissedIssueEventDismissedReview) GetDismissalCommitId() string`

GetDismissalCommitId returns the DismissalCommitId field if non-nil, zero value otherwise.

### GetDismissalCommitIdOk

`func (o *ReviewDismissedIssueEventDismissedReview) GetDismissalCommitIdOk() (*string, bool)`

GetDismissalCommitIdOk returns a tuple with the DismissalCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalCommitId

`func (o *ReviewDismissedIssueEventDismissedReview) SetDismissalCommitId(v string)`

SetDismissalCommitId sets DismissalCommitId field to given value.

### HasDismissalCommitId

`func (o *ReviewDismissedIssueEventDismissedReview) HasDismissalCommitId() bool`

HasDismissalCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


