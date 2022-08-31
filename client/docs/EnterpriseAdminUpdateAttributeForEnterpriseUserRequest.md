# EnterpriseAdminUpdateAttributeForEnterpriseUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | The SCIM schema URIs. | 
**Operations** | **[]map[string]interface{}** | Array of [SCIM operations](https://tools.ietf.org/html/rfc7644#section-3.5.2). | 

## Methods

### NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequest

`func NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequest(schemas []string, operations []map[string]interface{}, ) *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest`

NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequest instantiates a new EnterpriseAdminUpdateAttributeForEnterpriseUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequestWithDefaults

`func NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequestWithDefaults() *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest`

NewEnterpriseAdminUpdateAttributeForEnterpriseUserRequestWithDefaults instantiates a new EnterpriseAdminUpdateAttributeForEnterpriseUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetOperations

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest) GetOperations() []map[string]interface{}`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest) GetOperationsOk() (*[]map[string]interface{}, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *EnterpriseAdminUpdateAttributeForEnterpriseUserRequest) SetOperations(v []map[string]interface{})`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


