# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Commit** | [**ShortBranchCommit**](ShortBranchCommit.md) |  | 
**ZipballUrl** | **string** |  | 
**TarballUrl** | **string** |  | 
**NodeId** | **string** |  | 

## Methods

### NewTag

`func NewTag(name string, commit ShortBranchCommit, zipballUrl string, tarballUrl string, nodeId string, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### GetCommit

`func (o *Tag) GetCommit() ShortBranchCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *Tag) GetCommitOk() (*ShortBranchCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *Tag) SetCommit(v ShortBranchCommit)`

SetCommit sets Commit field to given value.


### GetZipballUrl

`func (o *Tag) GetZipballUrl() string`

GetZipballUrl returns the ZipballUrl field if non-nil, zero value otherwise.

### GetZipballUrlOk

`func (o *Tag) GetZipballUrlOk() (*string, bool)`

GetZipballUrlOk returns a tuple with the ZipballUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipballUrl

`func (o *Tag) SetZipballUrl(v string)`

SetZipballUrl sets ZipballUrl field to given value.


### GetTarballUrl

`func (o *Tag) GetTarballUrl() string`

GetTarballUrl returns the TarballUrl field if non-nil, zero value otherwise.

### GetTarballUrlOk

`func (o *Tag) GetTarballUrlOk() (*string, bool)`

GetTarballUrlOk returns a tuple with the TarballUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarballUrl

`func (o *Tag) SetTarballUrl(v string)`

SetTarballUrl sets TarballUrl field to given value.


### GetNodeId

`func (o *Tag) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Tag) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Tag) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


