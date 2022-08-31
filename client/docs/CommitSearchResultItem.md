# CommitSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Sha** | **string** |  | 
**HtmlUrl** | **string** |  | 
**CommentsUrl** | **string** |  | 
**Commit** | [**CommitSearchResultItemCommit**](CommitSearchResultItemCommit.md) |  | 
**Author** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Committer** | [**NullableNullableGitUser**](NullableGitUser.md) |  | 
**Parents** | [**[]FileCommitCommitParentsInner**](FileCommitCommitParentsInner.md) |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**Score** | **float32** |  | 
**NodeId** | **string** |  | 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 

## Methods

### NewCommitSearchResultItem

`func NewCommitSearchResultItem(url string, sha string, htmlUrl string, commentsUrl string, commit CommitSearchResultItemCommit, author NullableNullableSimpleUser, committer NullableNullableGitUser, parents []FileCommitCommitParentsInner, repository MinimalRepository, score float32, nodeId string, ) *CommitSearchResultItem`

NewCommitSearchResultItem instantiates a new CommitSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitSearchResultItemWithDefaults

`func NewCommitSearchResultItemWithDefaults() *CommitSearchResultItem`

NewCommitSearchResultItemWithDefaults instantiates a new CommitSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CommitSearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommitSearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommitSearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSha

`func (o *CommitSearchResultItem) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CommitSearchResultItem) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CommitSearchResultItem) SetSha(v string)`

SetSha sets Sha field to given value.


### GetHtmlUrl

`func (o *CommitSearchResultItem) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CommitSearchResultItem) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CommitSearchResultItem) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetCommentsUrl

`func (o *CommitSearchResultItem) GetCommentsUrl() string`

GetCommentsUrl returns the CommentsUrl field if non-nil, zero value otherwise.

### GetCommentsUrlOk

`func (o *CommitSearchResultItem) GetCommentsUrlOk() (*string, bool)`

GetCommentsUrlOk returns a tuple with the CommentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsUrl

`func (o *CommitSearchResultItem) SetCommentsUrl(v string)`

SetCommentsUrl sets CommentsUrl field to given value.


### GetCommit

`func (o *CommitSearchResultItem) GetCommit() CommitSearchResultItemCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitSearchResultItem) GetCommitOk() (*CommitSearchResultItemCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitSearchResultItem) SetCommit(v CommitSearchResultItemCommit)`

SetCommit sets Commit field to given value.


### GetAuthor

`func (o *CommitSearchResultItem) GetAuthor() NullableSimpleUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitSearchResultItem) GetAuthorOk() (*NullableSimpleUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitSearchResultItem) SetAuthor(v NullableSimpleUser)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *CommitSearchResultItem) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *CommitSearchResultItem) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCommitter

`func (o *CommitSearchResultItem) GetCommitter() NullableGitUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitSearchResultItem) GetCommitterOk() (*NullableGitUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitSearchResultItem) SetCommitter(v NullableGitUser)`

SetCommitter sets Committer field to given value.


### SetCommitterNil

`func (o *CommitSearchResultItem) SetCommitterNil(b bool)`

 SetCommitterNil sets the value for Committer to be an explicit nil

### UnsetCommitter
`func (o *CommitSearchResultItem) UnsetCommitter()`

UnsetCommitter ensures that no value is present for Committer, not even an explicit nil
### GetParents

`func (o *CommitSearchResultItem) GetParents() []FileCommitCommitParentsInner`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *CommitSearchResultItem) GetParentsOk() (*[]FileCommitCommitParentsInner, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *CommitSearchResultItem) SetParents(v []FileCommitCommitParentsInner)`

SetParents sets Parents field to given value.


### GetRepository

`func (o *CommitSearchResultItem) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CommitSearchResultItem) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CommitSearchResultItem) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetScore

`func (o *CommitSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CommitSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CommitSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetNodeId

`func (o *CommitSearchResultItem) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CommitSearchResultItem) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CommitSearchResultItem) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetTextMatches

`func (o *CommitSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *CommitSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *CommitSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *CommitSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


