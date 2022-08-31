# Blob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Encoding** | **string** |  | 
**Url** | **string** |  | 
**Sha** | **string** |  | 
**Size** | **NullableInt32** |  | 
**NodeId** | **string** |  | 
**HighlightedContent** | Pointer to **string** |  | [optional] 

## Methods

### NewBlob

`func NewBlob(content string, encoding string, url string, sha string, size NullableInt32, nodeId string, ) *Blob`

NewBlob instantiates a new Blob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobWithDefaults

`func NewBlobWithDefaults() *Blob`

NewBlobWithDefaults instantiates a new Blob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Blob) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Blob) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Blob) SetContent(v string)`

SetContent sets Content field to given value.


### GetEncoding

`func (o *Blob) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *Blob) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *Blob) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetUrl

`func (o *Blob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Blob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Blob) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSha

`func (o *Blob) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Blob) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Blob) SetSha(v string)`

SetSha sets Sha field to given value.


### GetSize

`func (o *Blob) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Blob) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Blob) SetSize(v int32)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *Blob) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Blob) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetNodeId

`func (o *Blob) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Blob) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Blob) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHighlightedContent

`func (o *Blob) GetHighlightedContent() string`

GetHighlightedContent returns the HighlightedContent field if non-nil, zero value otherwise.

### GetHighlightedContentOk

`func (o *Blob) GetHighlightedContentOk() (*string, bool)`

GetHighlightedContentOk returns a tuple with the HighlightedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightedContent

`func (o *Blob) SetHighlightedContent(v string)`

SetHighlightedContent sets HighlightedContent field to given value.

### HasHighlightedContent

`func (o *Blob) HasHighlightedContent() bool`

HasHighlightedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


