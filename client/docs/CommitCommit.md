# CommitCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Author** | [**NullableNullableGitUser**](NullableGitUser.md) |  | 
**Committer** | [**NullableNullableGitUser**](NullableGitUser.md) |  | 
**Message** | **string** |  | 
**CommentCount** | **int32** |  | 
**Tree** | [**CommitCommitTree**](CommitCommitTree.md) |  | 
**Verification** | Pointer to [**Verification**](Verification.md) |  | [optional] 

## Methods

### NewCommitCommit

`func NewCommitCommit(url string, author NullableNullableGitUser, committer NullableNullableGitUser, message string, commentCount int32, tree CommitCommitTree, ) *CommitCommit`

NewCommitCommit instantiates a new CommitCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitCommitWithDefaults

`func NewCommitCommitWithDefaults() *CommitCommit`

NewCommitCommitWithDefaults instantiates a new CommitCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CommitCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommitCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommitCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthor

`func (o *CommitCommit) GetAuthor() NullableGitUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitCommit) GetAuthorOk() (*NullableGitUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitCommit) SetAuthor(v NullableGitUser)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *CommitCommit) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *CommitCommit) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *CommitCommit) GetCommitter() NullableGitUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitCommit) GetCommitterOk() (*NullableGitUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitCommit) SetCommitter(v NullableGitUser)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *CommitCommit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *CommitCommit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil
### GetMessage

`func (o *CommitCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommitCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommitCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCommentCount

`func (o *CommitCommit) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *CommitCommit) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *CommitCommit) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetTree

`func (o *CommitCommit) GetTree() CommitCommitTree`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *CommitCommit) GetTreeOk() (*CommitCommitTree, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *CommitCommit) SetTree(v CommitCommitTree)`

SetTree sets Tree field to given value.


### GetVerification

`func (o *CommitCommit) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *CommitCommit) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *CommitCommit) SetVerification(v Verification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *CommitCommit) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


