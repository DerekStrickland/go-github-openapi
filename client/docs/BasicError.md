# BasicError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewBasicError

`func NewBasicError() *BasicError`

NewBasicError instantiates a new BasicError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicErrorWithDefaults

`func NewBasicErrorWithDefaults() *BasicError`

NewBasicErrorWithDefaults instantiates a new BasicError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BasicError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BasicError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BasicError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BasicError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *BasicError) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *BasicError) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *BasicError) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *BasicError) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetUrl

`func (o *BasicError) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BasicError) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BasicError) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BasicError) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *BasicError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BasicError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BasicError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BasicError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


