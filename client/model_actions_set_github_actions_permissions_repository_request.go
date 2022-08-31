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

// ActionsSetGithubActionsPermissionsRepositoryRequest struct for ActionsSetGithubActionsPermissionsRepositoryRequest
type ActionsSetGithubActionsPermissionsRepositoryRequest struct {
	// Whether GitHub Actions is enabled on the repository.
	Enabled bool `json:"enabled"`
	AllowedActions *AllowedActions `json:"allowed_actions,omitempty"`
}

// NewActionsSetGithubActionsPermissionsRepositoryRequest instantiates a new ActionsSetGithubActionsPermissionsRepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsSetGithubActionsPermissionsRepositoryRequest(enabled bool) *ActionsSetGithubActionsPermissionsRepositoryRequest {
	this := ActionsSetGithubActionsPermissionsRepositoryRequest{}
	this.Enabled = enabled
	return &this
}

// NewActionsSetGithubActionsPermissionsRepositoryRequestWithDefaults instantiates a new ActionsSetGithubActionsPermissionsRepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsSetGithubActionsPermissionsRepositoryRequestWithDefaults() *ActionsSetGithubActionsPermissionsRepositoryRequest {
	this := ActionsSetGithubActionsPermissionsRepositoryRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAllowedActions returns the AllowedActions field value if set, zero value otherwise.
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetAllowedActions() AllowedActions {
	if o == nil || o.AllowedActions == nil {
		var ret AllowedActions
		return ret
	}
	return *o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) GetAllowedActionsOk() (*AllowedActions, bool) {
	if o == nil || o.AllowedActions == nil {
		return nil, false
	}
	return o.AllowedActions, true
}

// HasAllowedActions returns a boolean if a field has been set.
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) HasAllowedActions() bool {
	if o != nil && o.AllowedActions != nil {
		return true
	}

	return false
}

// SetAllowedActions gets a reference to the given AllowedActions and assigns it to the AllowedActions field.
func (o *ActionsSetGithubActionsPermissionsRepositoryRequest) SetAllowedActions(v AllowedActions) {
	o.AllowedActions = &v
}

func (o ActionsSetGithubActionsPermissionsRepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AllowedActions != nil {
		toSerialize["allowed_actions"] = o.AllowedActions
	}
	return json.Marshal(toSerialize)
}

type NullableActionsSetGithubActionsPermissionsRepositoryRequest struct {
	value *ActionsSetGithubActionsPermissionsRepositoryRequest
	isSet bool
}

func (v NullableActionsSetGithubActionsPermissionsRepositoryRequest) Get() *ActionsSetGithubActionsPermissionsRepositoryRequest {
	return v.value
}

func (v *NullableActionsSetGithubActionsPermissionsRepositoryRequest) Set(val *ActionsSetGithubActionsPermissionsRepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsSetGithubActionsPermissionsRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsSetGithubActionsPermissionsRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsSetGithubActionsPermissionsRepositoryRequest(val *ActionsSetGithubActionsPermissionsRepositoryRequest) *NullableActionsSetGithubActionsPermissionsRepositoryRequest {
	return &NullableActionsSetGithubActionsPermissionsRepositoryRequest{value: val, isSet: true}
}

func (v NullableActionsSetGithubActionsPermissionsRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsSetGithubActionsPermissionsRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

