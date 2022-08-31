# EnterpriseAdminProvisionAndInviteEnterpriseUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | The SCIM schema URIs. | 
**UserName** | **string** | The username for the user. | 
**Name** | [**EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName.md) |  | 
**Emails** | [**[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner.md) | List of user emails. | 
**Groups** | Pointer to [**[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner**](EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner.md) | List of SCIM group IDs the user is a member of. | [optional] 

## Methods

### NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequest

`func NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequest(schemas []string, userName string, name EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName, emails []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner, ) *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest`

NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequest instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestWithDefaults

`func NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestWithDefaults() *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest`

NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestWithDefaults instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetUserName

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetName() EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetNameOk() (*EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) SetName(v EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName)`

SetName sets Name field to given value.


### GetEmails

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetEmails() []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetEmailsOk() (*[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) SetEmails(v []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner)`

SetEmails sets Emails field to given value.


### GetGroups

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetGroups() []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) GetGroupsOk() (*[]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) SetGroups(v []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


