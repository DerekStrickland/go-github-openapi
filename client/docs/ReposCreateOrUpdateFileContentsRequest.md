# ReposCreateOrUpdateFileContentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The commit message. | 
**Content** | **string** | The new file content, using Base64 encoding. | 
**Sha** | Pointer to **string** | **Required if you are updating a file**. The blob SHA of the file being replaced. | [optional] 
**Branch** | Pointer to **string** | The branch name. Default: the repository’s default branch (usually &#x60;master&#x60;) | [optional] 
**Committer** | Pointer to [**ReposCreateOrUpdateFileContentsRequestCommitter**](ReposCreateOrUpdateFileContentsRequestCommitter.md) |  | [optional] 
**Author** | Pointer to [**ReposCreateOrUpdateFileContentsRequestAuthor**](ReposCreateOrUpdateFileContentsRequestAuthor.md) |  | [optional] 

## Methods

### NewReposCreateOrUpdateFileContentsRequest

`func NewReposCreateOrUpdateFileContentsRequest(message string, content string, ) *ReposCreateOrUpdateFileContentsRequest`

NewReposCreateOrUpdateFileContentsRequest instantiates a new ReposCreateOrUpdateFileContentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateOrUpdateFileContentsRequestWithDefaults

`func NewReposCreateOrUpdateFileContentsRequestWithDefaults() *ReposCreateOrUpdateFileContentsRequest`

NewReposCreateOrUpdateFileContentsRequestWithDefaults instantiates a new ReposCreateOrUpdateFileContentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ReposCreateOrUpdateFileContentsRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReposCreateOrUpdateFileContentsRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReposCreateOrUpdateFileContentsRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetContent

`func (o *ReposCreateOrUpdateFileContentsRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReposCreateOrUpdateFileContentsRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReposCreateOrUpdateFileContentsRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetSha

`func (o *ReposCreateOrUpdateFileContentsRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ReposCreateOrUpdateFileContentsRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ReposCreateOrUpdateFileContentsRequest) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *ReposCreateOrUpdateFileContentsRequest) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetBranch

`func (o *ReposCreateOrUpdateFileContentsRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ReposCreateOrUpdateFileContentsRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ReposCreateOrUpdateFileContentsRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ReposCreateOrUpdateFileContentsRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommitter

`func (o *ReposCreateOrUpdateFileContentsRequest) GetCommitter() ReposCreateOrUpdateFileContentsRequestCommitter`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *ReposCreateOrUpdateFileContentsRequest) GetCommitterOk() (*ReposCreateOrUpdateFileContentsRequestCommitter, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *ReposCreateOrUpdateFileContentsRequest) SetCommitter(v ReposCreateOrUpdateFileContentsRequestCommitter)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *ReposCreateOrUpdateFileContentsRequest) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetAuthor

`func (o *ReposCreateOrUpdateFileContentsRequest) GetAuthor() ReposCreateOrUpdateFileContentsRequestAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReposCreateOrUpdateFileContentsRequest) GetAuthorOk() (*ReposCreateOrUpdateFileContentsRequestAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReposCreateOrUpdateFileContentsRequest) SetAuthor(v ReposCreateOrUpdateFileContentsRequestAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReposCreateOrUpdateFileContentsRequest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


