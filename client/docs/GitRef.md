# GitRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Object** | [**GitRefObject**](GitRefObject.md) |  | 

## Methods

### NewGitRef

`func NewGitRef(ref string, nodeId string, url string, object GitRefObject, ) *GitRef`

NewGitRef instantiates a new GitRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRefWithDefaults

`func NewGitRefWithDefaults() *GitRef`

NewGitRefWithDefaults instantiates a new GitRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GitRef) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GitRef) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GitRef) SetRef(v string)`

SetRef sets Ref field to given value.


### GetNodeId

`func (o *GitRef) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *GitRef) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *GitRef) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *GitRef) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitRef) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitRef) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetObject

`func (o *GitRef) GetObject() GitRefObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitRef) GetObjectOk() (*GitRefObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitRef) SetObject(v GitRefObject)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


