# ValidationErrorErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**Index** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to [**ValidationErrorErrorsInnerValue**](ValidationErrorErrorsInnerValue.md) |  | [optional] 

## Methods

### NewValidationErrorErrorsInner

`func NewValidationErrorErrorsInner(code string, ) *ValidationErrorErrorsInner`

NewValidationErrorErrorsInner instantiates a new ValidationErrorErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorErrorsInnerWithDefaults

`func NewValidationErrorErrorsInnerWithDefaults() *ValidationErrorErrorsInner`

NewValidationErrorErrorsInnerWithDefaults instantiates a new ValidationErrorErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ValidationErrorErrorsInner) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ValidationErrorErrorsInner) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ValidationErrorErrorsInner) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ValidationErrorErrorsInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetField

`func (o *ValidationErrorErrorsInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationErrorErrorsInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationErrorErrorsInner) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ValidationErrorErrorsInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationErrorErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationErrorErrorsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *ValidationErrorErrorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationErrorErrorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationErrorErrorsInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetIndex

`func (o *ValidationErrorErrorsInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ValidationErrorErrorsInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ValidationErrorErrorsInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ValidationErrorErrorsInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetValue

`func (o *ValidationErrorErrorsInner) GetValue() ValidationErrorErrorsInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValidationErrorErrorsInner) GetValueOk() (*ValidationErrorErrorsInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValidationErrorErrorsInner) SetValue(v ValidationErrorErrorsInnerValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ValidationErrorErrorsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


