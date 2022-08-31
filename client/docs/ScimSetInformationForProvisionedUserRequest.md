# ScimSetInformationForProvisionedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the user, suitable for display to end-users | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**UserName** | **string** | Configured by the admin. Could be an email, login, or username | 
**Name** | [**ScimProvisionAndInviteUserRequestName**](ScimProvisionAndInviteUserRequestName.md) |  | 
**Emails** | [**[]ScimSetInformationForProvisionedUserRequestEmailsInner**](ScimSetInformationForProvisionedUserRequestEmailsInner.md) | user emails | 

## Methods

### NewScimSetInformationForProvisionedUserRequest

`func NewScimSetInformationForProvisionedUserRequest(userName string, name ScimProvisionAndInviteUserRequestName, emails []ScimSetInformationForProvisionedUserRequestEmailsInner, ) *ScimSetInformationForProvisionedUserRequest`

NewScimSetInformationForProvisionedUserRequest instantiates a new ScimSetInformationForProvisionedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSetInformationForProvisionedUserRequestWithDefaults

`func NewScimSetInformationForProvisionedUserRequestWithDefaults() *ScimSetInformationForProvisionedUserRequest`

NewScimSetInformationForProvisionedUserRequestWithDefaults instantiates a new ScimSetInformationForProvisionedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimSetInformationForProvisionedUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimSetInformationForProvisionedUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimSetInformationForProvisionedUserRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDisplayName

`func (o *ScimSetInformationForProvisionedUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ScimSetInformationForProvisionedUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ScimSetInformationForProvisionedUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *ScimSetInformationForProvisionedUserRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ScimSetInformationForProvisionedUserRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ScimSetInformationForProvisionedUserRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGroups

`func (o *ScimSetInformationForProvisionedUserRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimSetInformationForProvisionedUserRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ScimSetInformationForProvisionedUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetActive

`func (o *ScimSetInformationForProvisionedUserRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ScimSetInformationForProvisionedUserRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ScimSetInformationForProvisionedUserRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserName

`func (o *ScimSetInformationForProvisionedUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScimSetInformationForProvisionedUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *ScimSetInformationForProvisionedUserRequest) GetName() ScimProvisionAndInviteUserRequestName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetNameOk() (*ScimProvisionAndInviteUserRequestName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimSetInformationForProvisionedUserRequest) SetName(v ScimProvisionAndInviteUserRequestName)`

SetName sets Name field to given value.


### GetEmails

`func (o *ScimSetInformationForProvisionedUserRequest) GetEmails() []ScimSetInformationForProvisionedUserRequestEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScimSetInformationForProvisionedUserRequest) GetEmailsOk() (*[]ScimSetInformationForProvisionedUserRequestEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScimSetInformationForProvisionedUserRequest) SetEmails(v []ScimSetInformationForProvisionedUserRequestEmailsInner)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


