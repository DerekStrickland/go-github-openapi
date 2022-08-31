# GitCreateTreeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tree** | [**[]GitCreateTreeRequestTreeInner**](GitCreateTreeRequestTreeInner.md) | Objects (of &#x60;path&#x60;, &#x60;mode&#x60;, &#x60;type&#x60;, and &#x60;sha&#x60;) specifying a tree structure. | 
**BaseTree** | Pointer to **string** | The SHA1 of an existing Git tree object which will be used as the base for the new tree. If provided, a new Git tree object will be created from entries in the Git tree object pointed to by &#x60;base_tree&#x60; and entries defined in the &#x60;tree&#x60; parameter. Entries defined in the &#x60;tree&#x60; parameter will overwrite items from &#x60;base_tree&#x60; with the same &#x60;path&#x60;. If you&#39;re creating new changes on a branch, then normally you&#39;d set &#x60;base_tree&#x60; to the SHA1 of the Git tree object of the current latest commit on the branch you&#39;re working on. If not provided, GitHub will create a new Git tree object from only the entries defined in the &#x60;tree&#x60; parameter. If you create a new commit pointing to such a tree, then all files which were a part of the parent commit&#39;s tree and were not defined in the &#x60;tree&#x60; parameter will be listed as deleted by the new commit.  | [optional] 

## Methods

### NewGitCreateTreeRequest

`func NewGitCreateTreeRequest(tree []GitCreateTreeRequestTreeInner, ) *GitCreateTreeRequest`

NewGitCreateTreeRequest instantiates a new GitCreateTreeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateTreeRequestWithDefaults

`func NewGitCreateTreeRequestWithDefaults() *GitCreateTreeRequest`

NewGitCreateTreeRequestWithDefaults instantiates a new GitCreateTreeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTree

`func (o *GitCreateTreeRequest) GetTree() []GitCreateTreeRequestTreeInner`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *GitCreateTreeRequest) GetTreeOk() (*[]GitCreateTreeRequestTreeInner, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *GitCreateTreeRequest) SetTree(v []GitCreateTreeRequestTreeInner)`

SetTree sets Tree field to given value.


### GetBaseTree

`func (o *GitCreateTreeRequest) GetBaseTree() string`

GetBaseTree returns the BaseTree field if non-nil, zero value otherwise.

### GetBaseTreeOk

`func (o *GitCreateTreeRequest) GetBaseTreeOk() (*string, bool)`

GetBaseTreeOk returns a tuple with the BaseTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTree

`func (o *GitCreateTreeRequest) SetBaseTree(v string)`

SetBaseTree sets BaseTree field to given value.

### HasBaseTree

`func (o *GitCreateTreeRequest) HasBaseTree() bool`

HasBaseTree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


