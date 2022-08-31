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

// CollaboratorPermissions struct for CollaboratorPermissions
type CollaboratorPermissions struct {
	Pull bool `json:"pull"`
	Triage *bool `json:"triage,omitempty"`
	Push bool `json:"push"`
	Maintain *bool `json:"maintain,omitempty"`
	Admin bool `json:"admin"`
}

// NewCollaboratorPermissions instantiates a new CollaboratorPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollaboratorPermissions(pull bool, push bool, admin bool) *CollaboratorPermissions {
	this := CollaboratorPermissions{}
	this.Pull = pull
	this.Push = push
	this.Admin = admin
	return &this
}

// NewCollaboratorPermissionsWithDefaults instantiates a new CollaboratorPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollaboratorPermissionsWithDefaults() *CollaboratorPermissions {
	this := CollaboratorPermissions{}
	return &this
}

// GetPull returns the Pull field value
func (o *CollaboratorPermissions) GetPull() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pull
}

// GetPullOk returns a tuple with the Pull field value
// and a boolean to check if the value has been set.
func (o *CollaboratorPermissions) GetPullOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pull, true
}

// SetPull sets field value
func (o *CollaboratorPermissions) SetPull(v bool) {
	o.Pull = v
}

// GetTriage returns the Triage field value if set, zero value otherwise.
func (o *CollaboratorPermissions) GetTriage() bool {
	if o == nil || o.Triage == nil {
		var ret bool
		return ret
	}
	return *o.Triage
}

// GetTriageOk returns a tuple with the Triage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollaboratorPermissions) GetTriageOk() (*bool, bool) {
	if o == nil || o.Triage == nil {
		return nil, false
	}
	return o.Triage, true
}

// HasTriage returns a boolean if a field has been set.
func (o *CollaboratorPermissions) HasTriage() bool {
	if o != nil && o.Triage != nil {
		return true
	}

	return false
}

// SetTriage gets a reference to the given bool and assigns it to the Triage field.
func (o *CollaboratorPermissions) SetTriage(v bool) {
	o.Triage = &v
}

// GetPush returns the Push field value
func (o *CollaboratorPermissions) GetPush() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Push
}

// GetPushOk returns a tuple with the Push field value
// and a boolean to check if the value has been set.
func (o *CollaboratorPermissions) GetPushOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Push, true
}

// SetPush sets field value
func (o *CollaboratorPermissions) SetPush(v bool) {
	o.Push = v
}

// GetMaintain returns the Maintain field value if set, zero value otherwise.
func (o *CollaboratorPermissions) GetMaintain() bool {
	if o == nil || o.Maintain == nil {
		var ret bool
		return ret
	}
	return *o.Maintain
}

// GetMaintainOk returns a tuple with the Maintain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollaboratorPermissions) GetMaintainOk() (*bool, bool) {
	if o == nil || o.Maintain == nil {
		return nil, false
	}
	return o.Maintain, true
}

// HasMaintain returns a boolean if a field has been set.
func (o *CollaboratorPermissions) HasMaintain() bool {
	if o != nil && o.Maintain != nil {
		return true
	}

	return false
}

// SetMaintain gets a reference to the given bool and assigns it to the Maintain field.
func (o *CollaboratorPermissions) SetMaintain(v bool) {
	o.Maintain = &v
}

// GetAdmin returns the Admin field value
func (o *CollaboratorPermissions) GetAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Admin
}

// GetAdminOk returns a tuple with the Admin field value
// and a boolean to check if the value has been set.
func (o *CollaboratorPermissions) GetAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Admin, true
}

// SetAdmin sets field value
func (o *CollaboratorPermissions) SetAdmin(v bool) {
	o.Admin = v
}

func (o CollaboratorPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pull"] = o.Pull
	}
	if o.Triage != nil {
		toSerialize["triage"] = o.Triage
	}
	if true {
		toSerialize["push"] = o.Push
	}
	if o.Maintain != nil {
		toSerialize["maintain"] = o.Maintain
	}
	if true {
		toSerialize["admin"] = o.Admin
	}
	return json.Marshal(toSerialize)
}

type NullableCollaboratorPermissions struct {
	value *CollaboratorPermissions
	isSet bool
}

func (v NullableCollaboratorPermissions) Get() *CollaboratorPermissions {
	return v.value
}

func (v *NullableCollaboratorPermissions) Set(val *CollaboratorPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableCollaboratorPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableCollaboratorPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollaboratorPermissions(val *CollaboratorPermissions) *NullableCollaboratorPermissions {
	return &NullableCollaboratorPermissions{value: val, isSet: true}
}

func (v NullableCollaboratorPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollaboratorPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

