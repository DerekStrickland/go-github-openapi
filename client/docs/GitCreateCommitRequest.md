# GitCreateCommitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The commit message | 
**Tree** | **string** | The SHA of the tree object this commit points to | 
**Parents** | Pointer to **[]string** | The SHAs of the commits that were the parents of this commit. If omitted or empty, the commit will be written as a root commit. For a single parent, an array of one SHA should be provided; for a merge commit, an array of more than one should be provided. | [optional] 
**Author** | Pointer to [**GitCreateCommitRequestAuthor**](GitCreateCommitRequestAuthor.md) |  | [optional] 
**Committer** | Pointer to [**GitCreateCommitRequestCommitter**](GitCreateCommitRequestCommitter.md) |  | [optional] 
**Signature** | Pointer to **string** | The [PGP signature](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) of the commit. GitHub adds the signature to the &#x60;gpgsig&#x60; header of the created commit. For a commit signature to be verifiable by Git or GitHub, it must be an ASCII-armored detached PGP signature over the string commit as it would be written to the object database. To pass a &#x60;signature&#x60; parameter, you need to first manually create a valid PGP signature, which can be complicated. You may find it easier to [use the command line](https://git-scm.com/book/id/v2/Git-Tools-Signing-Your-Work) to create signed commits. | [optional] 

## Methods

### NewGitCreateCommitRequest

`func NewGitCreateCommitRequest(message string, tree string, ) *GitCreateCommitRequest`

NewGitCreateCommitRequest instantiates a new GitCreateCommitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateCommitRequestWithDefaults

`func NewGitCreateCommitRequestWithDefaults() *GitCreateCommitRequest`

NewGitCreateCommitRequestWithDefaults instantiates a new GitCreateCommitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GitCreateCommitRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitCreateCommitRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitCreateCommitRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTree

`func (o *GitCreateCommitRequest) GetTree() string`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *GitCreateCommitRequest) GetTreeOk() (*string, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *GitCreateCommitRequest) SetTree(v string)`

SetTree sets Tree field to given value.


### GetParents

`func (o *GitCreateCommitRequest) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *GitCreateCommitRequest) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *GitCreateCommitRequest) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *GitCreateCommitRequest) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetAuthor

`func (o *GitCreateCommitRequest) GetAuthor() GitCreateCommitRequestAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GitCreateCommitRequest) GetAuthorOk() (*GitCreateCommitRequestAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GitCreateCommitRequest) SetAuthor(v GitCreateCommitRequestAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GitCreateCommitRequest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitter

`func (o *GitCreateCommitRequest) GetCommitter() GitCreateCommitRequestCommitter`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *GitCreateCommitRequest) GetCommitterOk() (*GitCreateCommitRequestCommitter, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *GitCreateCommitRequest) SetCommitter(v GitCreateCommitRequestCommitter)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *GitCreateCommitRequest) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetSignature

`func (o *GitCreateCommitRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GitCreateCommitRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GitCreateCommitRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *GitCreateCommitRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


