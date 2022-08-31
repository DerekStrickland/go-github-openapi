# ScimUpdateAttributeForUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**Operations** | [**[]ScimUpdateAttributeForUserRequestOperationsInner**](ScimUpdateAttributeForUserRequestOperationsInner.md) | Set of operations to be performed | 

## Methods

### NewScimUpdateAttributeForUserRequest

`func NewScimUpdateAttributeForUserRequest(operations []ScimUpdateAttributeForUserRequestOperationsInner, ) *ScimUpdateAttributeForUserRequest`

NewScimUpdateAttributeForUserRequest instantiates a new ScimUpdateAttributeForUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUpdateAttributeForUserRequestWithDefaults

`func NewScimUpdateAttributeForUserRequestWithDefaults() *ScimUpdateAttributeForUserRequest`

NewScimUpdateAttributeForUserRequestWithDefaults instantiates a new ScimUpdateAttributeForUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimUpdateAttributeForUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimUpdateAttributeForUserRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimUpdateAttributeForUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimUpdateAttributeForUserRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetOperations

`func (o *ScimUpdateAttributeForUserRequest) GetOperations() []ScimUpdateAttributeForUserRequestOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ScimUpdateAttributeForUserRequest) GetOperationsOk() (*[]ScimUpdateAttributeForUserRequestOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ScimUpdateAttributeForUserRequest) SetOperations(v []ScimUpdateAttributeForUserRequestOperationsInner)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


