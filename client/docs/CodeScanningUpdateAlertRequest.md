# CodeScanningUpdateAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**CodeScanningAlertSetState**](CodeScanningAlertSetState.md) |  | 
**DismissedReason** | Pointer to [**NullableCodeScanningAlertDismissedReason**](CodeScanningAlertDismissedReason.md) |  | [optional] 
**DismissedComment** | Pointer to **NullableString** | The dismissal comment associated with the dismissal of the alert. | [optional] 

## Methods

### NewCodeScanningUpdateAlertRequest

`func NewCodeScanningUpdateAlertRequest(state CodeScanningAlertSetState, ) *CodeScanningUpdateAlertRequest`

NewCodeScanningUpdateAlertRequest instantiates a new CodeScanningUpdateAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningUpdateAlertRequestWithDefaults

`func NewCodeScanningUpdateAlertRequestWithDefaults() *CodeScanningUpdateAlertRequest`

NewCodeScanningUpdateAlertRequestWithDefaults instantiates a new CodeScanningUpdateAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CodeScanningUpdateAlertRequest) GetState() CodeScanningAlertSetState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CodeScanningUpdateAlertRequest) GetStateOk() (*CodeScanningAlertSetState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CodeScanningUpdateAlertRequest) SetState(v CodeScanningAlertSetState)`

SetState sets State field to given value.


### GetDismissedReason

`func (o *CodeScanningUpdateAlertRequest) GetDismissedReason() CodeScanningAlertDismissedReason`

GetDismissedReason returns the DismissedReason field if non-nil, zero value otherwise.

### GetDismissedReasonOk

`func (o *CodeScanningUpdateAlertRequest) GetDismissedReasonOk() (*CodeScanningAlertDismissedReason, bool)`

GetDismissedReasonOk returns a tuple with the DismissedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedReason

`func (o *CodeScanningUpdateAlertRequest) SetDismissedReason(v CodeScanningAlertDismissedReason)`

SetDismissedReason sets DismissedReason field to given value.

### HasDismissedReason

`func (o *CodeScanningUpdateAlertRequest) HasDismissedReason() bool`

HasDismissedReason returns a boolean if a field has been set.

### SetDismissedReasonNil

`func (o *CodeScanningUpdateAlertRequest) SetDismissedReasonNil(b bool)`

 SetDismissedReasonNil sets the value for DismissedReason to be an explicit nil

### UnsetDismissedReason
`func (o *CodeScanningUpdateAlertRequest) UnsetDismissedReason()`

UnsetDismissedReason ensures that no value is present for DismissedReason, not even an explicit nil
### GetDismissedComment

`func (o *CodeScanningUpdateAlertRequest) GetDismissedComment() string`

GetDismissedComment returns the DismissedComment field if non-nil, zero value otherwise.

### GetDismissedCommentOk

`func (o *CodeScanningUpdateAlertRequest) GetDismissedCommentOk() (*string, bool)`

GetDismissedCommentOk returns a tuple with the DismissedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedComment

`func (o *CodeScanningUpdateAlertRequest) SetDismissedComment(v string)`

SetDismissedComment sets DismissedComment field to given value.

### HasDismissedComment

`func (o *CodeScanningUpdateAlertRequest) HasDismissedComment() bool`

HasDismissedComment returns a boolean if a field has been set.

### SetDismissedCommentNil

`func (o *CodeScanningUpdateAlertRequest) SetDismissedCommentNil(b bool)`

 SetDismissedCommentNil sets the value for DismissedComment to be an explicit nil

### UnsetDismissedComment
`func (o *CodeScanningUpdateAlertRequest) UnsetDismissedComment()`

UnsetDismissedComment ensures that no value is present for DismissedComment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


