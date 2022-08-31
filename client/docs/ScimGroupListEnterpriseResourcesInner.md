# ScimGroupListEnterpriseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** |  | 
**Id** | **string** |  | 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]ScimGroupListEnterpriseResourcesInnerMembersInner**](ScimGroupListEnterpriseResourcesInnerMembersInner.md) |  | [optional] 
**Meta** | Pointer to [**ScimGroupListEnterpriseResourcesInnerMeta**](ScimGroupListEnterpriseResourcesInnerMeta.md) |  | [optional] 

## Methods

### NewScimGroupListEnterpriseResourcesInner

`func NewScimGroupListEnterpriseResourcesInner(schemas []string, id string, ) *ScimGroupListEnterpriseResourcesInner`

NewScimGroupListEnterpriseResourcesInner instantiates a new ScimGroupListEnterpriseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimGroupListEnterpriseResourcesInnerWithDefaults

`func NewScimGroupListEnterpriseResourcesInnerWithDefaults() *ScimGroupListEnterpriseResourcesInner`

NewScimGroupListEnterpriseResourcesInnerWithDefaults instantiates a new ScimGroupListEnterpriseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimGroupListEnterpriseResourcesInner) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimGroupListEnterpriseResourcesInner) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimGroupListEnterpriseResourcesInner) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimGroupListEnterpriseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimGroupListEnterpriseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimGroupListEnterpriseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ScimGroupListEnterpriseResourcesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimGroupListEnterpriseResourcesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimGroupListEnterpriseResourcesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ScimGroupListEnterpriseResourcesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ScimGroupListEnterpriseResourcesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ScimGroupListEnterpriseResourcesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDisplayName

`func (o *ScimGroupListEnterpriseResourcesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimGroupListEnterpriseResourcesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimGroupListEnterpriseResourcesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimGroupListEnterpriseResourcesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMembers

`func (o *ScimGroupListEnterpriseResourcesInner) GetMembers() []ScimGroupListEnterpriseResourcesInnerMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScimGroupListEnterpriseResourcesInner) GetMembersOk() (*[]ScimGroupListEnterpriseResourcesInnerMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ScimGroupListEnterpriseResourcesInner) SetMembers(v []ScimGroupListEnterpriseResourcesInnerMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ScimGroupListEnterpriseResourcesInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMeta

`func (o *ScimGroupListEnterpriseResourcesInner) GetMeta() ScimGroupListEnterpriseResourcesInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimGroupListEnterpriseResourcesInner) GetMetaOk() (*ScimGroupListEnterpriseResourcesInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimGroupListEnterpriseResourcesInner) SetMeta(v ScimGroupListEnterpriseResourcesInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimGroupListEnterpriseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


