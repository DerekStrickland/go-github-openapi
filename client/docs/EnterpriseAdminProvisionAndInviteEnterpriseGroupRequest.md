# EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | The SCIM schema URIs. | 
**DisplayName** | **string** | The name of the SCIM group. This must match the GitHub organization that the group maps to. | 
**Members** | Pointer to [**[]EnterpriseAdminProvisionAndInviteEnterpriseGroupRequestMembersInner**](EnterpriseAdminProvisionAndInviteEnterpriseGroupRequestMembersInner.md) |  | [optional] 

## Methods

### NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest

`func NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(schemas []string, displayName string, ) *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest`

NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequestWithDefaults

`func NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequestWithDefaults() *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest`

NewEnterpriseAdminProvisionAndInviteEnterpriseGroupRequestWithDefaults instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetDisplayName

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMembers

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) GetMembers() []EnterpriseAdminProvisionAndInviteEnterpriseGroupRequestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) GetMembersOk() (*[]EnterpriseAdminProvisionAndInviteEnterpriseGroupRequestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) SetMembers(v []EnterpriseAdminProvisionAndInviteEnterpriseGroupRequestMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseGroupRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


