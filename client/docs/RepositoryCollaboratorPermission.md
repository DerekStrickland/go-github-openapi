# RepositoryCollaboratorPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** |  | 
**RoleName** | **string** |  | 
**User** | [**NullableNullableCollaborator**](NullableCollaborator.md) |  | 

## Methods

### NewRepositoryCollaboratorPermission

`func NewRepositoryCollaboratorPermission(permission string, roleName string, user NullableNullableCollaborator, ) *RepositoryCollaboratorPermission`

NewRepositoryCollaboratorPermission instantiates a new RepositoryCollaboratorPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryCollaboratorPermissionWithDefaults

`func NewRepositoryCollaboratorPermissionWithDefaults() *RepositoryCollaboratorPermission`

NewRepositoryCollaboratorPermissionWithDefaults instantiates a new RepositoryCollaboratorPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *RepositoryCollaboratorPermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *RepositoryCollaboratorPermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *RepositoryCollaboratorPermission) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetRoleName

`func (o *RepositoryCollaboratorPermission) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *RepositoryCollaboratorPermission) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *RepositoryCollaboratorPermission) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetUser

`func (o *RepositoryCollaboratorPermission) GetUser() NullableCollaborator`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RepositoryCollaboratorPermission) GetUserOk() (*NullableCollaborator, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RepositoryCollaboratorPermission) SetUser(v NullableCollaborator)`

SetUser sets User field to given value.


### SetUserNil

`func (o *RepositoryCollaboratorPermission) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RepositoryCollaboratorPermission) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


