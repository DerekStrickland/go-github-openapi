# CommitSearchResultItemCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**CommitSearchResultItemCommitAuthor**](CommitSearchResultItemCommitAuthor.md) |  | 
**Committer** | [**NullableNullableGitUser**](NullableGitUser.md) |  | 
**CommentCount** | **int32** |  | 
**Message** | **string** |  | 
**Tree** | [**ShortBranchCommit**](ShortBranchCommit.md) |  | 
**Url** | **string** |  | 
**Verification** | Pointer to [**Verification**](Verification.md) |  | [optional] 

## Methods

### NewCommitSearchResultItemCommit

`func NewCommitSearchResultItemCommit(author CommitSearchResultItemCommitAuthor, committer NullableNullableGitUser, commentCount int32, message string, tree ShortBranchCommit, url string, ) *CommitSearchResultItemCommit`

NewCommitSearchResultItemCommit instantiates a new CommitSearchResultItemCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitSearchResultItemCommitWithDefaults

`func NewCommitSearchResultItemCommitWithDefaults() *CommitSearchResultItemCommit`

NewCommitSearchResultItemCommitWithDefaults instantiates a new CommitSearchResultItemCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *CommitSearchResultItemCommit) GetAuthor() CommitSearchResultItemCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitSearchResultItemCommit) GetAuthorOk() (*CommitSearchResultItemCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitSearchResultItemCommit) SetAuthor(v CommitSearchResultItemCommitAuthor)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *CommitSearchResultItemCommit) GetCommitter() NullableGitUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitSearchResultItemCommit) GetCommitterOk() (*NullableGitUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitSearchResultItemCommit) SetCommitter(v NullableGitUser)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *CommitSearchResultItemCommit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *CommitSearchResultItemCommit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil
### GetCommentCount

`func (o *CommitSearchResultItemCommit) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *CommitSearchResultItemCommit) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *CommitSearchResultItemCommit) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetMessage

`func (o *CommitSearchResultItemCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommitSearchResultItemCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommitSearchResultItemCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTree

`func (o *CommitSearchResultItemCommit) GetTree() ShortBranchCommit`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *CommitSearchResultItemCommit) GetTreeOk() (*ShortBranchCommit, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *CommitSearchResultItemCommit) SetTree(v ShortBranchCommit)`

SetTree sets Tree field to given value.


### GetUrl

`func (o *CommitSearchResultItemCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommitSearchResultItemCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommitSearchResultItemCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerification

`func (o *CommitSearchResultItemCommit) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *CommitSearchResultItemCommit) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *CommitSearchResultItemCommit) SetVerification(v Verification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *CommitSearchResultItemCommit) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


