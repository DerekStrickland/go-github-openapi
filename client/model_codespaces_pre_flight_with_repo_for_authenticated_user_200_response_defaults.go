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

// CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults struct for CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults
type CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults struct {
	Location string `json:"location"`
	DevcontainerPath NullableString `json:"devcontainer_path"`
}

// NewCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults instantiates a new CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults(location string, devcontainerPath NullableString) *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults {
	this := CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults{}
	this.Location = location
	this.DevcontainerPath = devcontainerPath
	return &this
}

// NewCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaultsWithDefaults instantiates a new CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaultsWithDefaults() *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults {
	this := CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults{}
	return &this
}

// GetLocation returns the Location field value
func (o *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) SetLocation(v string) {
	o.Location = v
}

// GetDevcontainerPath returns the DevcontainerPath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) GetDevcontainerPath() string {
	if o == nil || o.DevcontainerPath.Get() == nil {
		var ret string
		return ret
	}

	return *o.DevcontainerPath.Get()
}

// GetDevcontainerPathOk returns a tuple with the DevcontainerPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) GetDevcontainerPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DevcontainerPath.Get(), o.DevcontainerPath.IsSet()
}

// SetDevcontainerPath sets field value
func (o *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) SetDevcontainerPath(v string) {
	o.DevcontainerPath.Set(&v)
}

func (o CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["devcontainer_path"] = o.DevcontainerPath.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults struct {
	value *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults
	isSet bool
}

func (v NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) Get() *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults {
	return v.value
}

func (v *NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) Set(val *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults(val *CodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) *NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults {
	return &NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults{value: val, isSet: true}
}

func (v NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespacesPreFlightWithRepoForAuthenticatedUser200ResponseDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


