/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"fmt"
)

// EnabledRepositories The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
type EnabledRepositories string

// List of enabled-repositories
const (
	ALL EnabledRepositories = "all"
	NONE EnabledRepositories = "none"
	SELECTED EnabledRepositories = "selected"
)

// All allowed values of EnabledRepositories enum
var AllowedEnabledRepositoriesEnumValues = []EnabledRepositories{
	"all",
	"none",
	"selected",
}

func (v *EnabledRepositories) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnabledRepositories(value)
	for _, existing := range AllowedEnabledRepositoriesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnabledRepositories", value)
}

// NewEnabledRepositoriesFromValue returns a pointer to a valid EnabledRepositories
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnabledRepositoriesFromValue(v string) (*EnabledRepositories, error) {
	ev := EnabledRepositories(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnabledRepositories: valid values are %v", v, AllowedEnabledRepositoriesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnabledRepositories) IsValid() bool {
	for _, existing := range AllowedEnabledRepositoriesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enabled-repositories value
func (v EnabledRepositories) Ptr() *EnabledRepositories {
	return &v
}

type NullableEnabledRepositories struct {
	value *EnabledRepositories
	isSet bool
}

func (v NullableEnabledRepositories) Get() *EnabledRepositories {
	return v.value
}

func (v *NullableEnabledRepositories) Set(val *EnabledRepositories) {
	v.value = val
	v.isSet = true
}

func (v NullableEnabledRepositories) IsSet() bool {
	return v.isSet
}

func (v *NullableEnabledRepositories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnabledRepositories(val *EnabledRepositories) *NullableEnabledRepositories {
	return &NullableEnabledRepositories{value: val, isSet: true}
}

func (v NullableEnabledRepositories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnabledRepositories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

