# ReposDeleteFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The commit message. | 
**Sha** | **string** | The blob SHA of the file being replaced. | 
**Branch** | Pointer to **string** | The branch name. Default: the repository’s default branch (usually &#x60;master&#x60;) | [optional] 
**Committer** | Pointer to [**ReposDeleteFileRequestCommitter**](ReposDeleteFileRequestCommitter.md) |  | [optional] 
**Author** | Pointer to [**ReposDeleteFileRequestAuthor**](ReposDeleteFileRequestAuthor.md) |  | [optional] 

## Methods

### NewReposDeleteFileRequest

`func NewReposDeleteFileRequest(message string, sha string, ) *ReposDeleteFileRequest`

NewReposDeleteFileRequest instantiates a new ReposDeleteFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposDeleteFileRequestWithDefaults

`func NewReposDeleteFileRequestWithDefaults() *ReposDeleteFileRequest`

NewReposDeleteFileRequestWithDefaults instantiates a new ReposDeleteFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ReposDeleteFileRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReposDeleteFileRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReposDeleteFileRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSha

`func (o *ReposDeleteFileRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ReposDeleteFileRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ReposDeleteFileRequest) SetSha(v string)`

SetSha sets Sha field to given value.


### GetBranch

`func (o *ReposDeleteFileRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ReposDeleteFileRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ReposDeleteFileRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ReposDeleteFileRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommitter

`func (o *ReposDeleteFileRequest) GetCommitter() ReposDeleteFileRequestCommitter`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *ReposDeleteFileRequest) GetCommitterOk() (*ReposDeleteFileRequestCommitter, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *ReposDeleteFileRequest) SetCommitter(v ReposDeleteFileRequestCommitter)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *ReposDeleteFileRequest) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetAuthor

`func (o *ReposDeleteFileRequest) GetAuthor() ReposDeleteFileRequestAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReposDeleteFileRequest) GetAuthorOk() (*ReposDeleteFileRequestAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReposDeleteFileRequest) SetAuthor(v ReposDeleteFileRequestAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReposDeleteFileRequest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


