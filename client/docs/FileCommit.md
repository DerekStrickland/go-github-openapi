# FileCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**NullableFileCommitContent**](FileCommitContent.md) |  | 
**Commit** | [**FileCommitCommit**](FileCommitCommit.md) |  | 

## Methods

### NewFileCommit

`func NewFileCommit(content NullableFileCommitContent, commit FileCommitCommit, ) *FileCommit`

NewFileCommit instantiates a new FileCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCommitWithDefaults

`func NewFileCommitWithDefaults() *FileCommit`

NewFileCommitWithDefaults instantiates a new FileCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FileCommit) GetContent() FileCommitContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileCommit) GetContentOk() (*FileCommitContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileCommit) SetContent(v FileCommitContent)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *FileCommit) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *FileCommit) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCommit

`func (o *FileCommit) GetCommit() FileCommitCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *FileCommit) GetCommitOk() (*FileCommitCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *FileCommit) SetCommit(v FileCommitCommit)`

SetCommit sets Commit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


