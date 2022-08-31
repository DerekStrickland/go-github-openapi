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

// CodespaceRuntimeConstraints struct for CodespaceRuntimeConstraints
type CodespaceRuntimeConstraints struct {
	// The privacy settings a user can select from when forwarding a port.
	AllowedPortPrivacySettings []string `json:"allowed_port_privacy_settings,omitempty"`
}

// NewCodespaceRuntimeConstraints instantiates a new CodespaceRuntimeConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespaceRuntimeConstraints() *CodespaceRuntimeConstraints {
	this := CodespaceRuntimeConstraints{}
	return &this
}

// NewCodespaceRuntimeConstraintsWithDefaults instantiates a new CodespaceRuntimeConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespaceRuntimeConstraintsWithDefaults() *CodespaceRuntimeConstraints {
	this := CodespaceRuntimeConstraints{}
	return &this
}

// GetAllowedPortPrivacySettings returns the AllowedPortPrivacySettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodespaceRuntimeConstraints) GetAllowedPortPrivacySettings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedPortPrivacySettings
}

// GetAllowedPortPrivacySettingsOk returns a tuple with the AllowedPortPrivacySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespaceRuntimeConstraints) GetAllowedPortPrivacySettingsOk() ([]string, bool) {
	if o == nil || o.AllowedPortPrivacySettings == nil {
		return nil, false
	}
	return o.AllowedPortPrivacySettings, true
}

// HasAllowedPortPrivacySettings returns a boolean if a field has been set.
func (o *CodespaceRuntimeConstraints) HasAllowedPortPrivacySettings() bool {
	if o != nil && o.AllowedPortPrivacySettings != nil {
		return true
	}

	return false
}

// SetAllowedPortPrivacySettings gets a reference to the given []string and assigns it to the AllowedPortPrivacySettings field.
func (o *CodespaceRuntimeConstraints) SetAllowedPortPrivacySettings(v []string) {
	o.AllowedPortPrivacySettings = v
}

func (o CodespaceRuntimeConstraints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedPortPrivacySettings != nil {
		toSerialize["allowed_port_privacy_settings"] = o.AllowedPortPrivacySettings
	}
	return json.Marshal(toSerialize)
}

type NullableCodespaceRuntimeConstraints struct {
	value *CodespaceRuntimeConstraints
	isSet bool
}

func (v NullableCodespaceRuntimeConstraints) Get() *CodespaceRuntimeConstraints {
	return v.value
}

func (v *NullableCodespaceRuntimeConstraints) Set(val *CodespaceRuntimeConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespaceRuntimeConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespaceRuntimeConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespaceRuntimeConstraints(val *CodespaceRuntimeConstraints) *NullableCodespaceRuntimeConstraints {
	return &NullableCodespaceRuntimeConstraints{value: val, isSet: true}
}

func (v NullableCodespaceRuntimeConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespaceRuntimeConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

