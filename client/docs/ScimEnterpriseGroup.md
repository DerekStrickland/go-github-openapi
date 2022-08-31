# ScimEnterpriseGroup

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

### NewScimEnterpriseGroup

`func NewScimEnterpriseGroup(schemas []string, id string, ) *ScimEnterpriseGroup`

NewScimEnterpriseGroup instantiates a new ScimEnterpriseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimEnterpriseGroupWithDefaults

`func NewScimEnterpriseGroupWithDefaults() *ScimEnterpriseGroup`

NewScimEnterpriseGroupWithDefaults instantiates a new ScimEnterpriseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimEnterpriseGroup) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimEnterpriseGroup) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimEnterpriseGroup) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimEnterpriseGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimEnterpriseGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimEnterpriseGroup) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ScimEnterpriseGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimEnterpriseGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimEnterpriseGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ScimEnterpriseGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ScimEnterpriseGroup) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ScimEnterpriseGroup) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDisplayName

`func (o *ScimEnterpriseGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimEnterpriseGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimEnterpriseGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimEnterpriseGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMembers

`func (o *ScimEnterpriseGroup) GetMembers() []ScimGroupListEnterpriseResourcesInnerMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ScimEnterpriseGroup) GetMembersOk() (*[]ScimGroupListEnterpriseResourcesInnerMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ScimEnterpriseGroup) SetMembers(v []ScimGroupListEnterpriseResourcesInnerMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ScimEnterpriseGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMeta

`func (o *ScimEnterpriseGroup) GetMeta() ScimGroupListEnterpriseResourcesInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimEnterpriseGroup) GetMetaOk() (*ScimGroupListEnterpriseResourcesInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimEnterpriseGroup) SetMeta(v ScimGroupListEnterpriseResourcesInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimEnterpriseGroup) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


