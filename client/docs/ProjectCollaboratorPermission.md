# ProjectCollaboratorPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 

## Methods

### NewProjectCollaboratorPermission

`func NewProjectCollaboratorPermission(permission string, user NullableNullableSimpleUser, ) *ProjectCollaboratorPermission`

NewProjectCollaboratorPermission instantiates a new ProjectCollaboratorPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCollaboratorPermissionWithDefaults

`func NewProjectCollaboratorPermissionWithDefaults() *ProjectCollaboratorPermission`

NewProjectCollaboratorPermissionWithDefaults instantiates a new ProjectCollaboratorPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *ProjectCollaboratorPermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ProjectCollaboratorPermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ProjectCollaboratorPermission) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetUser

`func (o *ProjectCollaboratorPermission) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProjectCollaboratorPermission) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProjectCollaboratorPermission) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *ProjectCollaboratorPermission) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ProjectCollaboratorPermission) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


