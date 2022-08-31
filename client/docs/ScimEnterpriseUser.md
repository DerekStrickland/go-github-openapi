# ScimEnterpriseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** |  | 
**Id** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**ScimUserListEnterpriseResourcesInnerName**](ScimUserListEnterpriseResourcesInnerName.md) |  | [optional] 
**Emails** | Pointer to [**[]ScimEnterpriseUserEmailsInner**](ScimEnterpriseUserEmailsInner.md) |  | [optional] 
**Groups** | Pointer to [**[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**ScimGroupListEnterpriseResourcesInnerMeta**](ScimGroupListEnterpriseResourcesInnerMeta.md) |  | [optional] 

## Methods

### NewScimEnterpriseUser

`func NewScimEnterpriseUser(schemas []string, id string, ) *ScimEnterpriseUser`

NewScimEnterpriseUser instantiates a new ScimEnterpriseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimEnterpriseUserWithDefaults

`func NewScimEnterpriseUserWithDefaults() *ScimEnterpriseUser`

NewScimEnterpriseUserWithDefaults instantiates a new ScimEnterpriseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimEnterpriseUser) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimEnterpriseUser) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimEnterpriseUser) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimEnterpriseUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimEnterpriseUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimEnterpriseUser) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ScimEnterpriseUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimEnterpriseUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimEnterpriseUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ScimEnterpriseUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUserName

`func (o *ScimEnterpriseUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimEnterpriseUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimEnterpriseUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ScimEnterpriseUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetName

`func (o *ScimEnterpriseUser) GetName() ScimUserListEnterpriseResourcesInnerName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimEnterpriseUser) GetNameOk() (*ScimUserListEnterpriseResourcesInnerName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimEnterpriseUser) SetName(v ScimUserListEnterpriseResourcesInnerName)`

SetName sets Name field to given value.

### HasName

`func (o *ScimEnterpriseUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmails

`func (o *ScimEnterpriseUser) GetEmails() []ScimEnterpriseUserEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScimEnterpriseUser) GetEmailsOk() (*[]ScimEnterpriseUserEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScimEnterpriseUser) SetEmails(v []ScimEnterpriseUserEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ScimEnterpriseUser) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetGroups

`func (o *ScimEnterpriseUser) GetGroups() []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimEnterpriseUser) GetGroupsOk() (*[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimEnterpriseUser) SetGroups(v []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ScimEnterpriseUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetActive

`func (o *ScimEnterpriseUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimEnterpriseUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimEnterpriseUser) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ScimEnterpriseUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMeta

`func (o *ScimEnterpriseUser) GetMeta() ScimGroupListEnterpriseResourcesInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimEnterpriseUser) GetMetaOk() (*ScimGroupListEnterpriseResourcesInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimEnterpriseUser) SetMeta(v ScimGroupListEnterpriseResourcesInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimEnterpriseUser) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


