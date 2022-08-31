# GitCreateBlobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The new blob&#39;s content. | 
**Encoding** | Pointer to **string** | The encoding used for &#x60;content&#x60;. Currently, &#x60;\&quot;utf-8\&quot;&#x60; and &#x60;\&quot;base64\&quot;&#x60; are supported. | [optional] [default to "utf-8"]

## Methods

### NewGitCreateBlobRequest

`func NewGitCreateBlobRequest(content string, ) *GitCreateBlobRequest`

NewGitCreateBlobRequest instantiates a new GitCreateBlobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateBlobRequestWithDefaults

`func NewGitCreateBlobRequestWithDefaults() *GitCreateBlobRequest`

NewGitCreateBlobRequestWithDefaults instantiates a new GitCreateBlobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GitCreateBlobRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitCreateBlobRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitCreateBlobRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetEncoding

`func (o *GitCreateBlobRequest) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GitCreateBlobRequest) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GitCreateBlobRequest) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GitCreateBlobRequest) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


