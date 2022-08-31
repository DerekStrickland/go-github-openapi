# ScimUserListEnterpriseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** |  | 
**Id** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**ScimUserListEnterpriseResourcesInnerName**](ScimUserListEnterpriseResourcesInnerName.md) |  | [optional] 
**Emails** | Pointer to [**[]ScimUserListEnterpriseResourcesInnerEmailsInner**](ScimUserListEnterpriseResourcesInnerEmailsInner.md) |  | [optional] 
**Groups** | Pointer to [**[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**ScimGroupListEnterpriseResourcesInnerMeta**](ScimGroupListEnterpriseResourcesInnerMeta.md) |  | [optional] 

## Methods

### NewScimUserListEnterpriseResourcesInner

`func NewScimUserListEnterpriseResourcesInner(schemas []string, id string, ) *ScimUserListEnterpriseResourcesInner`

NewScimUserListEnterpriseResourcesInner instantiates a new ScimUserListEnterpriseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUserListEnterpriseResourcesInnerWithDefaults

`func NewScimUserListEnterpriseResourcesInnerWithDefaults() *ScimUserListEnterpriseResourcesInner`

NewScimUserListEnterpriseResourcesInnerWithDefaults instantiates a new ScimUserListEnterpriseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimUserListEnterpriseResourcesInner) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimUserListEnterpriseResourcesInner) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimUserListEnterpriseResourcesInner) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimUserListEnterpriseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimUserListEnterpriseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimUserListEnterpriseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ScimUserListEnterpriseResourcesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimUserListEnterpriseResourcesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimUserListEnterpriseResourcesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ScimUserListEnterpriseResourcesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUserName

`func (o *ScimUserListEnterpriseResourcesInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimUserListEnterpriseResourcesInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimUserListEnterpriseResourcesInner) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ScimUserListEnterpriseResourcesInner) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetName

`func (o *ScimUserListEnterpriseResourcesInner) GetName() ScimUserListEnterpriseResourcesInnerName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimUserListEnterpriseResourcesInner) GetNameOk() (*ScimUserListEnterpriseResourcesInnerName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimUserListEnterpriseResourcesInner) SetName(v ScimUserListEnterpriseResourcesInnerName)`

SetName sets Name field to given value.

### HasName

`func (o *ScimUserListEnterpriseResourcesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmails

`func (o *ScimUserListEnterpriseResourcesInner) GetEmails() []ScimUserListEnterpriseResourcesInnerEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScimUserListEnterpriseResourcesInner) GetEmailsOk() (*[]ScimUserListEnterpriseResourcesInnerEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScimUserListEnterpriseResourcesInner) SetEmails(v []ScimUserListEnterpriseResourcesInnerEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ScimUserListEnterpriseResourcesInner) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetGroups

`func (o *ScimUserListEnterpriseResourcesInner) GetGroups() []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimUserListEnterpriseResourcesInner) GetGroupsOk() (*[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimUserListEnterpriseResourcesInner) SetGroups(v []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ScimUserListEnterpriseResourcesInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetActive

`func (o *ScimUserListEnterpriseResourcesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimUserListEnterpriseResourcesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimUserListEnterpriseResourcesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ScimUserListEnterpriseResourcesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMeta

`func (o *ScimUserListEnterpriseResourcesInner) GetMeta() ScimGroupListEnterpriseResourcesInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimUserListEnterpriseResourcesInner) GetMetaOk() (*ScimGroupListEnterpriseResourcesInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimUserListEnterpriseResourcesInner) SetMeta(v ScimGroupListEnterpriseResourcesInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimUserListEnterpriseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


