# PullsMergeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitTitle** | Pointer to **string** | Title for the automatic commit message. | [optional] 
**CommitMessage** | Pointer to **string** | Extra detail to append to automatic commit message. | [optional] 
**Sha** | Pointer to **string** | SHA that pull request head must match to allow merge. | [optional] 
**MergeMethod** | Pointer to **string** | Merge method to use. Possible values are &#x60;merge&#x60;, &#x60;squash&#x60; or &#x60;rebase&#x60;. Default is &#x60;merge&#x60;. | [optional] 

## Methods

### NewPullsMergeRequest

`func NewPullsMergeRequest() *PullsMergeRequest`

NewPullsMergeRequest instantiates a new PullsMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsMergeRequestWithDefaults

`func NewPullsMergeRequestWithDefaults() *PullsMergeRequest`

NewPullsMergeRequestWithDefaults instantiates a new PullsMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitTitle

`func (o *PullsMergeRequest) GetCommitTitle() string`

GetCommitTitle returns the CommitTitle field if non-nil, zero value otherwise.

### GetCommitTitleOk

`func (o *PullsMergeRequest) GetCommitTitleOk() (*string, bool)`

GetCommitTitleOk returns a tuple with the CommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTitle

`func (o *PullsMergeRequest) SetCommitTitle(v string)`

SetCommitTitle sets CommitTitle field to given value.

### HasCommitTitle

`func (o *PullsMergeRequest) HasCommitTitle() bool`

HasCommitTitle returns a boolean if a field has been set.

### GetCommitMessage

`func (o *PullsMergeRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *PullsMergeRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *PullsMergeRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *PullsMergeRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetSha

`func (o *PullsMergeRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PullsMergeRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PullsMergeRequest) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *PullsMergeRequest) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetMergeMethod

`func (o *PullsMergeRequest) GetMergeMethod() string`

GetMergeMethod returns the MergeMethod field if non-nil, zero value otherwise.

### GetMergeMethodOk

`func (o *PullsMergeRequest) GetMergeMethodOk() (*string, bool)`

GetMergeMethodOk returns a tuple with the MergeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMethod

`func (o *PullsMergeRequest) SetMergeMethod(v string)`

SetMergeMethod sets MergeMethod field to given value.

### HasMergeMethod

`func (o *PullsMergeRequest) HasMergeMethod() bool`

HasMergeMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


