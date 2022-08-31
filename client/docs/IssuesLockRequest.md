# IssuesLockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockReason** | Pointer to **string** | The reason for locking the issue or pull request conversation. Lock will fail if you don&#39;t use one of these reasons:   \\* &#x60;off-topic&#x60;   \\* &#x60;too heated&#x60;   \\* &#x60;resolved&#x60;   \\* &#x60;spam&#x60; | [optional] 

## Methods

### NewIssuesLockRequest

`func NewIssuesLockRequest() *IssuesLockRequest`

NewIssuesLockRequest instantiates a new IssuesLockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesLockRequestWithDefaults

`func NewIssuesLockRequestWithDefaults() *IssuesLockRequest`

NewIssuesLockRequestWithDefaults instantiates a new IssuesLockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockReason

`func (o *IssuesLockRequest) GetLockReason() string`

GetLockReason returns the LockReason field if non-nil, zero value otherwise.

### GetLockReasonOk

`func (o *IssuesLockRequest) GetLockReasonOk() (*string, bool)`

GetLockReasonOk returns a tuple with the LockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockReason

`func (o *IssuesLockRequest) SetLockReason(v string)`

SetLockReason sets LockReason field to given value.

### HasLockReason

`func (o *IssuesLockRequest) HasLockReason() bool`

HasLockReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


