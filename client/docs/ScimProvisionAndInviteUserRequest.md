# ScimProvisionAndInviteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | Configured by the admin. Could be an email, login, or username | 
**DisplayName** | Pointer to **string** | The name of the user, suitable for display to end-users | [optional] 
**Name** | [**ScimProvisionAndInviteUserRequestName**](ScimProvisionAndInviteUserRequestName.md) |  | 
**Emails** | [**[]ScimProvisionAndInviteUserRequestEmailsInner**](ScimProvisionAndInviteUserRequestEmailsInner.md) | user emails | 
**Schemas** | Pointer to **[]string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewScimProvisionAndInviteUserRequest

`func NewScimProvisionAndInviteUserRequest(userName string, name ScimProvisionAndInviteUserRequestName, emails []ScimProvisionAndInviteUserRequestEmailsInner, ) *ScimProvisionAndInviteUserRequest`

NewScimProvisionAndInviteUserRequest instantiates a new ScimProvisionAndInviteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimProvisionAndInviteUserRequestWithDefaults

`func NewScimProvisionAndInviteUserRequestWithDefaults() *ScimProvisionAndInviteUserRequest`

NewScimProvisionAndInviteUserRequestWithDefaults instantiates a new ScimProvisionAndInviteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *ScimProvisionAndInviteUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimProvisionAndInviteUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimProvisionAndInviteUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetDisplayName

`func (o *ScimProvisionAndInviteUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimProvisionAndInviteUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimProvisionAndInviteUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimProvisionAndInviteUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ScimProvisionAndInviteUserRequest) GetName() ScimProvisionAndInviteUserRequestName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimProvisionAndInviteUserRequest) GetNameOk() (*ScimProvisionAndInviteUserRequestName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimProvisionAndInviteUserRequest) SetName(v ScimProvisionAndInviteUserRequestName)`

SetName sets Name field to given value.


### GetEmails

`func (o *ScimProvisionAndInviteUserRequest) GetEmails() []ScimProvisionAndInviteUserRequestEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScimProvisionAndInviteUserRequest) GetEmailsOk() (*[]ScimProvisionAndInviteUserRequestEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScimProvisionAndInviteUserRequest) SetEmails(v []ScimProvisionAndInviteUserRequestEmailsInner)`

SetEmails sets Emails field to given value.


### GetSchemas

`func (o *ScimProvisionAndInviteUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimProvisionAndInviteUserRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimProvisionAndInviteUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimProvisionAndInviteUserRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetExternalId

`func (o *ScimProvisionAndInviteUserRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimProvisionAndInviteUserRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimProvisionAndInviteUserRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ScimProvisionAndInviteUserRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGroups

`func (o *ScimProvisionAndInviteUserRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimProvisionAndInviteUserRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimProvisionAndInviteUserRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ScimProvisionAndInviteUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetActive

`func (o *ScimProvisionAndInviteUserRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimProvisionAndInviteUserRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimProvisionAndInviteUserRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ScimProvisionAndInviteUserRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


