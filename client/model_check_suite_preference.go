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

// CheckSuitePreference Check suite configuration preferences for a repository.
type CheckSuitePreference struct {
	Preferences CheckSuitePreferencePreferences `json:"preferences"`
	Repository MinimalRepository `json:"repository"`
}

// NewCheckSuitePreference instantiates a new CheckSuitePreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckSuitePreference(preferences CheckSuitePreferencePreferences, repository MinimalRepository) *CheckSuitePreference {
	this := CheckSuitePreference{}
	this.Preferences = preferences
	this.Repository = repository
	return &this
}

// NewCheckSuitePreferenceWithDefaults instantiates a new CheckSuitePreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckSuitePreferenceWithDefaults() *CheckSuitePreference {
	this := CheckSuitePreference{}
	return &this
}

// GetPreferences returns the Preferences field value
func (o *CheckSuitePreference) GetPreferences() CheckSuitePreferencePreferences {
	if o == nil {
		var ret CheckSuitePreferencePreferences
		return ret
	}

	return o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value
// and a boolean to check if the value has been set.
func (o *CheckSuitePreference) GetPreferencesOk() (*CheckSuitePreferencePreferences, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preferences, true
}

// SetPreferences sets field value
func (o *CheckSuitePreference) SetPreferences(v CheckSuitePreferencePreferences) {
	o.Preferences = v
}

// GetRepository returns the Repository field value
func (o *CheckSuitePreference) GetRepository() MinimalRepository {
	if o == nil {
		var ret MinimalRepository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CheckSuitePreference) GetRepositoryOk() (*MinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *CheckSuitePreference) SetRepository(v MinimalRepository) {
	o.Repository = v
}

func (o CheckSuitePreference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["preferences"] = o.Preferences
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableCheckSuitePreference struct {
	value *CheckSuitePreference
	isSet bool
}

func (v NullableCheckSuitePreference) Get() *CheckSuitePreference {
	return v.value
}

func (v *NullableCheckSuitePreference) Set(val *CheckSuitePreference) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckSuitePreference) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckSuitePreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckSuitePreference(val *CheckSuitePreference) *NullableCheckSuitePreference {
	return &NullableCheckSuitePreference{value: val, isSet: true}
}

func (v NullableCheckSuitePreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckSuitePreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


