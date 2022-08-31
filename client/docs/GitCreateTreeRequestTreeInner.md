# GitCreateTreeRequestTreeInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The file referenced in the tree. | [optional] 
**Mode** | Pointer to **string** | The file mode; one of &#x60;100644&#x60; for file (blob), &#x60;100755&#x60; for executable (blob), &#x60;040000&#x60; for subdirectory (tree), &#x60;160000&#x60; for submodule (commit), or &#x60;120000&#x60; for a blob that specifies the path of a symlink. | [optional] 
**Type** | Pointer to **string** | Either &#x60;blob&#x60;, &#x60;tree&#x60;, or &#x60;commit&#x60;. | [optional] 
**Sha** | Pointer to **NullableString** | The SHA1 checksum ID of the object in the tree. Also called &#x60;tree.sha&#x60;. If the value is &#x60;null&#x60; then the file will be deleted.      **Note:** Use either &#x60;tree.sha&#x60; or &#x60;content&#x60; to specify the contents of the entry. Using both &#x60;tree.sha&#x60; and &#x60;content&#x60; will return an error. | [optional] 
**Content** | Pointer to **string** | The content you want this file to have. GitHub will write this blob out and use that SHA for this entry. Use either this, or &#x60;tree.sha&#x60;.      **Note:** Use either &#x60;tree.sha&#x60; or &#x60;content&#x60; to specify the contents of the entry. Using both &#x60;tree.sha&#x60; and &#x60;content&#x60; will return an error. | [optional] 

## Methods

### NewGitCreateTreeRequestTreeInner

`func NewGitCreateTreeRequestTreeInner() *GitCreateTreeRequestTreeInner`

NewGitCreateTreeRequestTreeInner instantiates a new GitCreateTreeRequestTreeInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateTreeRequestTreeInnerWithDefaults

`func NewGitCreateTreeRequestTreeInnerWithDefaults() *GitCreateTreeRequestTreeInner`

NewGitCreateTreeRequestTreeInnerWithDefaults instantiates a new GitCreateTreeRequestTreeInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GitCreateTreeRequestTreeInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitCreateTreeRequestTreeInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitCreateTreeRequestTreeInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitCreateTreeRequestTreeInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMode

`func (o *GitCreateTreeRequestTreeInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GitCreateTreeRequestTreeInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GitCreateTreeRequestTreeInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GitCreateTreeRequestTreeInner) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetType

`func (o *GitCreateTreeRequestTreeInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitCreateTreeRequestTreeInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitCreateTreeRequestTreeInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitCreateTreeRequestTreeInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSha

`func (o *GitCreateTreeRequestTreeInner) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitCreateTreeRequestTreeInner) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitCreateTreeRequestTreeInner) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitCreateTreeRequestTreeInner) HasSha() bool`

HasSha returns a boolean if a field has been set.

### SetShaNil

`func (o *GitCreateTreeRequestTreeInner) SetShaNil(b bool)`

 SetShaNil sets the value for Sha to be an explicit nil

### UnsetSha
`func (o *GitCreateTreeRequestTreeInner) UnsetSha()`

UnsetSha ensures that no value is present for Sha, not even an explicit nil
### GetContent

`func (o *GitCreateTreeRequestTreeInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitCreateTreeRequestTreeInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitCreateTreeRequestTreeInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitCreateTreeRequestTreeInner) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


