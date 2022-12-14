/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// RepositoryCollaboratorPermission Repository Collaborator Permission
type RepositoryCollaboratorPermission struct {
	Permission string `json:"permission"`
	RoleName string `json:"role_name"`
	User NullableNullableCollaborator `json:"user"`
}

// NewRepositoryCollaboratorPermission instantiates a new RepositoryCollaboratorPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryCollaboratorPermission(permission string, roleName string, user NullableNullableCollaborator) *RepositoryCollaboratorPermission {
	this := RepositoryCollaboratorPermission{}
	this.Permission = permission
	this.RoleName = roleName
	this.User = user
	return &this
}

// NewRepositoryCollaboratorPermissionWithDefaults instantiates a new RepositoryCollaboratorPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryCollaboratorPermissionWithDefaults() *RepositoryCollaboratorPermission {
	this := RepositoryCollaboratorPermission{}
	return &this
}

// GetPermission returns the Permission field value
func (o *RepositoryCollaboratorPermission) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *RepositoryCollaboratorPermission) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *RepositoryCollaboratorPermission) SetPermission(v string) {
	o.Permission = v
}

// GetRoleName returns the RoleName field value
func (o *RepositoryCollaboratorPermission) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *RepositoryCollaboratorPermission) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *RepositoryCollaboratorPermission) SetRoleName(v string) {
	o.RoleName = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for NullableCollaborator will be returned
func (o *RepositoryCollaboratorPermission) GetUser() NullableCollaborator {
	if o == nil || o.User.Get() == nil {
		var ret NullableCollaborator
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RepositoryCollaboratorPermission) GetUserOk() (*NullableCollaborator, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *RepositoryCollaboratorPermission) SetUser(v NullableCollaborator) {
	o.User.Set(&v)
}

func (o RepositoryCollaboratorPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if true {
		toSerialize["role_name"] = o.RoleName
	}
	if true {
		toSerialize["user"] = o.User.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryCollaboratorPermission struct {
	value *RepositoryCollaboratorPermission
	isSet bool
}

func (v NullableRepositoryCollaboratorPermission) Get() *RepositoryCollaboratorPermission {
	return v.value
}

func (v *NullableRepositoryCollaboratorPermission) Set(val *RepositoryCollaboratorPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryCollaboratorPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryCollaboratorPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryCollaboratorPermission(val *RepositoryCollaboratorPermission) *NullableRepositoryCollaboratorPermission {
	return &NullableRepositoryCollaboratorPermission{value: val, isSet: true}
}

func (v NullableRepositoryCollaboratorPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryCollaboratorPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


