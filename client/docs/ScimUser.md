# ScimUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | SCIM schema used. | 
**Id** | **string** | Unique identifier of an external identity | 
**ExternalId** | **NullableString** | The ID of the User. | 
**UserName** | **NullableString** | Configured by the admin. Could be an email, login, or username | 
**DisplayName** | Pointer to **NullableString** | The name of the user, suitable for display to end-users | [optional] 
**Name** | [**ScimUserName**](ScimUserName.md) |  | 
**Emails** | [**[]ScimUserEmailsInner**](ScimUserEmailsInner.md) | user emails | 
**Active** | **bool** | The active status of the User. | 
**Meta** | [**ScimUserMeta**](ScimUserMeta.md) |  | 
**OrganizationId** | Pointer to **int32** | The ID of the organization. | [optional] 
**Operations** | Pointer to [**[]ScimUserOperationsInner**](ScimUserOperationsInner.md) | Set of operations to be performed | [optional] 
**Groups** | Pointer to [**[]ScimUserGroupsInner**](ScimUserGroupsInner.md) | associated groups | [optional] 

## Methods

### NewScimUser

`func NewScimUser(schemas []string, id string, externalId NullableString, userName NullableString, name ScimUserName, emails []ScimUserEmailsInner, active bool, meta ScimUserMeta, ) *ScimUser`

NewScimUser instantiates a new ScimUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUserWithDefaults

`func NewScimUserWithDefaults() *ScimUser`

NewScimUserWithDefaults instantiates a new ScimUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimUser) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimUser) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimUser) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ScimUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimUser) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ScimUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *ScimUser) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ScimUser) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetUserName

`func (o *ScimUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimUser) SetUserName(v string)`

SetUserName sets UserName field to given value.


### SetUserNameNil

`func (o *ScimUser) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ScimUser) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetDisplayName

`func (o *ScimUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ScimUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ScimUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetName

`func (o *ScimUser) GetName() ScimUserName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimUser) GetNameOk() (*ScimUserName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimUser) SetName(v ScimUserName)`

SetName sets Name field to given value.


### GetEmails

`func (o *ScimUser) GetEmails() []ScimUserEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScimUser) GetEmailsOk() (*[]ScimUserEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScimUser) SetEmails(v []ScimUserEmailsInner)`

SetEmails sets Emails field to given value.


### GetActive

`func (o *ScimUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimUser) SetActive(v bool)`

SetActive sets Active field to given value.


### GetMeta

`func (o *ScimUser) GetMeta() ScimUserMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimUser) GetMetaOk() (*ScimUserMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimUser) SetMeta(v ScimUserMeta)`

SetMeta sets Meta field to given value.


### GetOrganizationId

`func (o *ScimUser) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ScimUser) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ScimUser) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ScimUser) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOperations

`func (o *ScimUser) GetOperations() []ScimUserOperationsInner`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ScimUser) GetOperationsOk() (*[]ScimUserOperationsInner, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ScimUser) SetOperations(v []ScimUserOperationsInner)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ScimUser) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetGroups

`func (o *ScimUser) GetGroups() []ScimUserGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimUser) GetGroupsOk() (*[]ScimUserGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimUser) SetGroups(v []ScimUserGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ScimUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


