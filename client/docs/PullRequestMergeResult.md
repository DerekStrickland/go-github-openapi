# PullRequestMergeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** |  | 
**Merged** | **bool** |  | 
**Message** | **string** |  | 

## Methods

### NewPullRequestMergeResult

`func NewPullRequestMergeResult(sha string, merged bool, message string, ) *PullRequestMergeResult`

NewPullRequestMergeResult instantiates a new PullRequestMergeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestMergeResultWithDefaults

`func NewPullRequestMergeResultWithDefaults() *PullRequestMergeResult`

NewPullRequestMergeResultWithDefaults instantiates a new PullRequestMergeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *PullRequestMergeResult) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PullRequestMergeResult) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PullRequestMergeResult) SetSha(v string)`

SetSha sets Sha field to given value.


### GetMerged

`func (o *PullRequestMergeResult) GetMerged() bool`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *PullRequestMergeResult) GetMergedOk() (*bool, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *PullRequestMergeResult) SetMerged(v bool)`

SetMerged sets Merged field to given value.


### GetMessage

`func (o *PullRequestMergeResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PullRequestMergeResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PullRequestMergeResult) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


