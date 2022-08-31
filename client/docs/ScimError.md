# ScimError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** |  | [optional] 
**DocumentationUrl** | Pointer to **NullableString** |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**ScimType** | Pointer to **NullableString** |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 

## Methods

### NewScimError

`func NewScimError() *ScimError`

NewScimError instantiates a new ScimError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimErrorWithDefaults

`func NewScimErrorWithDefaults() *ScimError`

NewScimErrorWithDefaults instantiates a new ScimError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ScimError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScimError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScimError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScimError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ScimError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ScimError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetDocumentationUrl

`func (o *ScimError) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *ScimError) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *ScimError) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *ScimError) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### SetDocumentationUrlNil

`func (o *ScimError) SetDocumentationUrlNil(b bool)`

 SetDocumentationUrlNil sets the value for DocumentationUrl to be an explicit nil

### UnsetDocumentationUrl
`func (o *ScimError) UnsetDocumentationUrl()`

UnsetDocumentationUrl ensures that no value is present for DocumentationUrl, not even an explicit nil
### GetDetail

`func (o *ScimError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ScimError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ScimError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ScimError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ScimError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ScimError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetStatus

`func (o *ScimError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScimError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScimError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScimError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScimType

`func (o *ScimError) GetScimType() string`

GetScimType returns the ScimType field if non-nil, zero value otherwise.

### GetScimTypeOk

`func (o *ScimError) GetScimTypeOk() (*string, bool)`

GetScimTypeOk returns a tuple with the ScimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimType

`func (o *ScimError) SetScimType(v string)`

SetScimType sets ScimType field to given value.

### HasScimType

`func (o *ScimError) HasScimType() bool`

HasScimType returns a boolean if a field has been set.

### SetScimTypeNil

`func (o *ScimError) SetScimTypeNil(b bool)`

 SetScimTypeNil sets the value for ScimType to be an explicit nil

### UnsetScimType
`func (o *ScimError) UnsetScimType()`

UnsetScimType ensures that no value is present for ScimType, not even an explicit nil
### GetSchemas

`func (o *ScimError) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimError) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimError) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimError) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


