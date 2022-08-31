# FileCommitCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**FileCommitCommitAuthor**](FileCommitCommitAuthor.md) |  | [optional] 
**Committer** | Pointer to [**FileCommitCommitAuthor**](FileCommitCommitAuthor.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Tree** | Pointer to [**FileCommitCommitTree**](FileCommitCommitTree.md) |  | [optional] 
**Parents** | Pointer to [**[]FileCommitCommitParentsInner**](FileCommitCommitParentsInner.md) |  | [optional] 
**Verification** | Pointer to [**FileCommitCommitVerification**](FileCommitCommitVerification.md) |  | [optional] 

## Methods

### NewFileCommitCommit

`func NewFileCommitCommit() *FileCommitCommit`

NewFileCommitCommit instantiates a new FileCommitCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCommitCommitWithDefaults

`func NewFileCommitCommitWithDefaults() *FileCommitCommit`

NewFileCommitCommitWithDefaults instantiates a new FileCommitCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *FileCommitCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *FileCommitCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *FileCommitCommit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *FileCommitCommit) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetNodeId

`func (o *FileCommitCommit) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FileCommitCommit) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FileCommitCommit) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *FileCommitCommit) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetUrl

`func (o *FileCommitCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileCommitCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileCommitCommit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileCommitCommit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *FileCommitCommit) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *FileCommitCommit) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *FileCommitCommit) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *FileCommitCommit) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetAuthor

`func (o *FileCommitCommit) GetAuthor() FileCommitCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *FileCommitCommit) GetAuthorOk() (*FileCommitCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *FileCommitCommit) SetAuthor(v FileCommitCommitAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *FileCommitCommit) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitter

`func (o *FileCommitCommit) GetCommitter() FileCommitCommitAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *FileCommitCommit) GetCommitterOk() (*FileCommitCommitAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *FileCommitCommit) SetCommitter(v FileCommitCommitAuthor)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *FileCommitCommit) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetMessage

`func (o *FileCommitCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FileCommitCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FileCommitCommit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FileCommitCommit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTree

`func (o *FileCommitCommit) GetTree() FileCommitCommitTree`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *FileCommitCommit) GetTreeOk() (*FileCommitCommitTree, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *FileCommitCommit) SetTree(v FileCommitCommitTree)`

SetTree sets Tree field to given value.

### HasTree

`func (o *FileCommitCommit) HasTree() bool`

HasTree returns a boolean if a field has been set.

### GetParents

`func (o *FileCommitCommit) GetParents() []FileCommitCommitParentsInner`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *FileCommitCommit) GetParentsOk() (*[]FileCommitCommitParentsInner, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *FileCommitCommit) SetParents(v []FileCommitCommitParentsInner)`

SetParents sets Parents field to given value.

### HasParents

`func (o *FileCommitCommit) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetVerification

`func (o *FileCommitCommit) GetVerification() FileCommitCommitVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *FileCommitCommit) GetVerificationOk() (*FileCommitCommitVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *FileCommitCommit) SetVerification(v FileCommitCommitVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *FileCommitCommit) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


