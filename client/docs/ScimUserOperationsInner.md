# ScimUserOperationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**ScimUserOperationsInnerValue**](ScimUserOperationsInnerValue.md) |  | [optional] 

## Methods

### NewScimUserOperationsInner

`func NewScimUserOperationsInner(op string, ) *ScimUserOperationsInner`

NewScimUserOperationsInner instantiates a new ScimUserOperationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUserOperationsInnerWithDefaults

`func NewScimUserOperationsInnerWithDefaults() *ScimUserOperationsInner`

NewScimUserOperationsInnerWithDefaults instantiates a new ScimUserOperationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ScimUserOperationsInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ScimUserOperationsInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ScimUserOperationsInner) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *ScimUserOperationsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ScimUserOperationsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ScimUserOperationsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ScimUserOperationsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *ScimUserOperationsInner) GetValue() ScimUserOperationsInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScimUserOperationsInner) GetValueOk() (*ScimUserOperationsInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScimUserOperationsInner) SetValue(v ScimUserOperationsInnerValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ScimUserOperationsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


