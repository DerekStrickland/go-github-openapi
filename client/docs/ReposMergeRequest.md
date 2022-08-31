# ReposMergeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **string** | The name of the base branch that the head will be merged into. | 
**Head** | **string** | The head to merge. This can be a branch name or a commit SHA1. | 
**CommitMessage** | Pointer to **string** | Commit message to use for the merge commit. If omitted, a default message will be used. | [optional] 

## Methods

### NewReposMergeRequest

`func NewReposMergeRequest(base string, head string, ) *ReposMergeRequest`

NewReposMergeRequest instantiates a new ReposMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposMergeRequestWithDefaults

`func NewReposMergeRequestWithDefaults() *ReposMergeRequest`

NewReposMergeRequestWithDefaults instantiates a new ReposMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *ReposMergeRequest) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ReposMergeRequest) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ReposMergeRequest) SetBase(v string)`

SetBase sets Base field to given value.


### GetHead

`func (o *ReposMergeRequest) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *ReposMergeRequest) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *ReposMergeRequest) SetHead(v string)`

SetHead sets Head field to given value.


### GetCommitMessage

`func (o *ReposMergeRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ReposMergeRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ReposMergeRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *ReposMergeRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


