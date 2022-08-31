# GitTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** |  | 
**Url** | **string** |  | 
**Truncated** | **bool** |  | 
**Tree** | [**[]GitTreeTreeInner**](GitTreeTreeInner.md) | Objects specifying a tree structure | 

## Methods

### NewGitTree

`func NewGitTree(sha string, url string, truncated bool, tree []GitTreeTreeInner, ) *GitTree`

NewGitTree instantiates a new GitTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTreeWithDefaults

`func NewGitTreeWithDefaults() *GitTree`

NewGitTreeWithDefaults instantiates a new GitTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *GitTree) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitTree) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitTree) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *GitTree) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitTree) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitTree) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTruncated

`func (o *GitTree) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GitTree) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GitTree) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetTree

`func (o *GitTree) GetTree() []GitTreeTreeInner`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *GitTree) GetTreeOk() (*[]GitTreeTreeInner, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *GitTree) SetTree(v []GitTreeTreeInner)`

SetTree sets Tree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


