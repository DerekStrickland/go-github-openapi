# GitCommitTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** | SHA for the commit | 
**Url** | **string** |  | 

## Methods

### NewGitCommitTree

`func NewGitCommitTree(sha string, url string, ) *GitCommitTree`

NewGitCommitTree instantiates a new GitCommitTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitTreeWithDefaults

`func NewGitCommitTreeWithDefaults() *GitCommitTree`

NewGitCommitTreeWithDefaults instantiates a new GitCommitTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *GitCommitTree) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitCommitTree) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitCommitTree) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *GitCommitTree) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitCommitTree) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitCommitTree) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


