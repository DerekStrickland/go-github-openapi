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

// TeamPermissions struct for TeamPermissions
type TeamPermissions struct {
	Pull bool `json:"pull"`
	Triage bool `json:"triage"`
	Push bool `json:"push"`
	Maintain bool `json:"maintain"`
	Admin bool `json:"admin"`
}

// NewTeamPermissions instantiates a new TeamPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamPermissions(pull bool, triage bool, push bool, maintain bool, admin bool) *TeamPermissions {
	this := TeamPermissions{}
	this.Pull = pull
	this.Triage = triage
	this.Push = push
	this.Maintain = maintain
	this.Admin = admin
	return &this
}

// NewTeamPermissionsWithDefaults instantiates a new TeamPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamPermissionsWithDefaults() *TeamPermissions {
	this := TeamPermissions{}
	return &this
}

// GetPull returns the Pull field value
func (o *TeamPermissions) GetPull() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pull
}

// GetPullOk returns a tuple with the Pull field value
// and a boolean to check if the value has been set.
func (o *TeamPermissions) GetPullOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pull, true
}

// SetPull sets field value
func (o *TeamPermissions) SetPull(v bool) {
	o.Pull = v
}

// GetTriage returns the Triage field value
func (o *TeamPermissions) GetTriage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Triage
}

// GetTriageOk returns a tuple with the Triage field value
// and a boolean to check if the value has been set.
func (o *TeamPermissions) GetTriageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Triage, true
}

// SetTriage sets field value
func (o *TeamPermissions) SetTriage(v bool) {
	o.Triage = v
}

// GetPush returns the Push field value
func (o *TeamPermissions) GetPush() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Push
}

// GetPushOk returns a tuple with the Push field value
// and a boolean to check if the value has been set.
func (o *TeamPermissions) GetPushOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Push, true
}

// SetPush sets field value
func (o *TeamPermissions) SetPush(v bool) {
	o.Push = v
}

// GetMaintain returns the Maintain field value
func (o *TeamPermissions) GetMaintain() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Maintain
}

// GetMaintainOk returns a tuple with the Maintain field value
// and a boolean to check if the value has been set.
func (o *TeamPermissions) GetMaintainOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maintain, true
}

// SetMaintain sets field value
func (o *TeamPermissions) SetMaintain(v bool) {
	o.Maintain = v
}

// GetAdmin returns the Admin field value
func (o *TeamPermissions) GetAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Admin
}

// GetAdminOk returns a tuple with the Admin field value
// and a boolean to check if the value has been set.
func (o *TeamPermissions) GetAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Admin, true
}

// SetAdmin sets field value
func (o *TeamPermissions) SetAdmin(v bool) {
	o.Admin = v
}

func (o TeamPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pull"] = o.Pull
	}
	if true {
		toSerialize["triage"] = o.Triage
	}
	if true {
		toSerialize["push"] = o.Push
	}
	if true {
		toSerialize["maintain"] = o.Maintain
	}
	if true {
		toSerialize["admin"] = o.Admin
	}
	return json.Marshal(toSerialize)
}

type NullableTeamPermissions struct {
	value *TeamPermissions
	isSet bool
}

func (v NullableTeamPermissions) Get() *TeamPermissions {
	return v.value
}

func (v *NullableTeamPermissions) Set(val *TeamPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamPermissions(val *TeamPermissions) *NullableTeamPermissions {
	return &NullableTeamPermissions{value: val, isSet: true}
}

func (v NullableTeamPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


