# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Sha** | **string** |  | 
**NodeId** | **string** |  | 
**HtmlUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**Commit** | [**CommitCommit**](CommitCommit.md) |  | 
**Author** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Committer** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Parents** | [**[]CommitParentsInner**](CommitParentsInner.md) |  | 
**Stats** | Pointer to [**CommitStats**](CommitStats.md) |  | [optional] 
**Files** | Pointer to [**[]DiffEntry**](DiffEntry.md) |  | [optional] 

## Methods

### NewCommit

`func NewCommit(url string, sha string, nodeId string, htmlUrl string, commentsUrl string, commit CommitCommit, author NullableNullableSimpleUser, committer NullableNullableSimpleUser, parents []CommitParentsInner, ) *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Commit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Commit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Commit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSha

`func (o *Commit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Commit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Commit) SetSha(v string)`

SetSha sets Sha field to given value.


### GetNodeId

`func (o *Commit) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Commit) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Commit) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHtmlUrl

`func (o *Commit) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Commit) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Commit) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCommentsUrl

`func (o *Commit) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *Commit) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *Commit) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCommit

`func (o *Commit) GetCommit() CommitCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *Commit) GetCommitOk() (*CommitCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *Commit) SetCommit(v CommitCommit)`

SetCommit sets Commit field to given value.


### GetAuthor

`func (o *Commit) GetAuthor() NullableSimpleUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Commit) GetAuthorOk() (*NullableSimpleUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Commit) SetAuthor(v NullableSimpleUser)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *Commit) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *Commit) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *Commit) GetCommitter() NullableSimpleUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *Commit) GetCommitterOk() (*NullableSimpleUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *Commit) SetCommitter(v NullableSimpleUser)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *Commit) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *Commit) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil
### GetParents

`func (o *Commit) GetParents() []CommitParentsInner`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Commit) GetParentsOk() (*[]CommitParentsInner, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Commit) SetParents(v []CommitParentsInner)`

SetParents sets Parents field to given value.


### GetStats

`func (o *Commit) GetStats() CommitStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Commit) GetStatsOk() (*CommitStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Commit) SetStats(v CommitStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Commit) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetFiles

`func (o *Commit) GetFiles() []DiffEntry`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Commit) GetFilesOk() (*[]DiffEntry, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Commit) SetFiles(v []DiffEntry)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Commit) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


