# EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | The SCIM schema URIs. | 
**Operations** | [**[]EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner**](EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner.md) | Array of [SCIM operations](https://tools.ietf.org/html/rfc7644#section-3.5.2). | 

## Methods

### NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequest

`func NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequest(schemas []string, operations []EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner, ) *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest`

NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequest instantiates a new EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequestWithDefaults

`func NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequestWithDefaults() *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest`

NewEnterpriseAdminUpdateAttributeForEnterpriseGroupRequestWithDefaults instantiates a new EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetOperations

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest) GetOperations() []EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest) GetOperationsOk() (*[]EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseGroupRequest) SetOperations(v []EnterpriseAdminUpdateAttributeForEnterpriseGroupRequestOperationsInner)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


