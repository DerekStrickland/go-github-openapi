# GitCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** | SHA for the commit | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Author** | [**GitCommitAuthor**](GitCommitAuthor.md) |  | 
**Committer** | [**GitCommitAuthor**](GitCommitAuthor.md) |  | 
**Message** | **string** | Message describing the purpose of the commit | 
**Tree** | [**GitCommitTree**](GitCommitTree.md) |  | 
**Parents** | [**[]GitCommitParentsInner**](GitCommitParentsInner.md) |  | 
**Verification** | [**GitCommitVerification**](GitCommitVerification.md) |  | 
**HtmlUrl** | **string** |  | 

## Methods

### NewGitCommit

`func NewGitCommit(sha string, nodeId string, url string, author GitCommitAuthor, committer GitCommitAuthor, message string, tree GitCommitTree, parents []GitCommitParentsInner, verification GitCommitVerification, htmlUrl string, ) *GitCommit`

NewGitCommit instantiates a new GitCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitWithDefaults

`func NewGitCommitWithDefaults() *GitCommit`

NewGitCommitWithDefaults instantiates a new GitCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *GitCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitCommit) SetSha(v string)`

SetSha sets Sha field to given value.


### GetNodeId

`func (o *GitCommit) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *GitCommit) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *GitCommit) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *GitCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthor

`func (o *GitCommit) GetAuthor() GitCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GitCommit) GetAuthorOk() (*GitCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GitCommit) SetAuthor(v GitCommitAuthor)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *GitCommit) GetCommitter() GitCommitAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *GitCommit) GetCommitterOk() (*GitCommitAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *GitCommit) SetCommitter(v GitCommitAuthor)`

SetCommitter sets Committer field to given value.


### GetMessage

`func (o *GitCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTree

`func (o *GitCommit) GetTree() GitCommitTree`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *GitCommit) GetTreeOk() (*GitCommitTree, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *GitCommit) SetTree(v GitCommitTree)`

SetTree sets Tree field to given value.


### GetParents

`func (o *GitCommit) GetParents() []GitCommitParentsInner`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *GitCommit) GetParentsOk() (*[]GitCommitParentsInner, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *GitCommit) SetParents(v []GitCommitParentsInner)`

SetParents sets Parents field to given value.


### GetVerification

`func (o *GitCommit) GetVerification() GitCommitVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *GitCommit) GetVerificationOk() (*GitCommitVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *GitCommit) SetVerification(v GitCommitVerification)`

SetVerification sets Verification field to given value.


### GetHtmlUrl

`func (o *GitCommit) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GitCommit) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GitCommit) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


