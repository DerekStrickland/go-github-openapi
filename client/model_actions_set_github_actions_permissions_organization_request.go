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

// ActionsSetGithubActionsPermissionsOrganizationRequest struct for ActionsSetGithubActionsPermissionsOrganizationRequest
type ActionsSetGithubActionsPermissionsOrganizationRequest struct {
	EnabledRepositories EnabledRepositories `json:"enabled_repositories"`
	AllowedActions *AllowedActions `json:"allowed_actions,omitempty"`
}

// NewActionsSetGithubActionsPermissionsOrganizationRequest instantiates a new ActionsSetGithubActionsPermissionsOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsSetGithubActionsPermissionsOrganizationRequest(enabledRepositories EnabledRepositories) *ActionsSetGithubActionsPermissionsOrganizationRequest {
	this := ActionsSetGithubActionsPermissionsOrganizationRequest{}
	this.EnabledRepositories = enabledRepositories
	return &this
}

// NewActionsSetGithubActionsPermissionsOrganizationRequestWithDefaults instantiates a new ActionsSetGithubActionsPermissionsOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsSetGithubActionsPermissionsOrganizationRequestWithDefaults() *ActionsSetGithubActionsPermissionsOrganizationRequest {
	this := ActionsSetGithubActionsPermissionsOrganizationRequest{}
	return &this
}

// GetEnabledRepositories returns the EnabledRepositories field value
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetEnabledRepositories() EnabledRepositories {
	if o == nil {
		var ret EnabledRepositories
		return ret
	}

	return o.EnabledRepositories
}

// GetEnabledRepositoriesOk returns a tuple with the EnabledRepositories field value
// and a boolean to check if the value has been set.
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetEnabledRepositoriesOk() (*EnabledRepositories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledRepositories, true
}

// SetEnabledRepositories sets field value
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) SetEnabledRepositories(v EnabledRepositories) {
	o.EnabledRepositories = v
}

// GetAllowedActions returns the AllowedActions field value if set, zero value otherwise.
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetAllowedActions() AllowedActions {
	if o == nil || o.AllowedActions == nil {
		var ret AllowedActions
		return ret
	}
	return *o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) GetAllowedActionsOk() (*AllowedActions, bool) {
	if o == nil || o.AllowedActions == nil {
		return nil, false
	}
	return o.AllowedActions, true
}

// HasAllowedActions returns a boolean if a field has been set.
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) HasAllowedActions() bool {
	if o != nil && o.AllowedActions != nil {
		return true
	}

	return false
}

// SetAllowedActions gets a reference to the given AllowedActions and assigns it to the AllowedActions field.
func (o *ActionsSetGithubActionsPermissionsOrganizationRequest) SetAllowedActions(v AllowedActions) {
	o.AllowedActions = &v
}

func (o ActionsSetGithubActionsPermissionsOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled_repositories"] = o.EnabledRepositories
	}
	if o.AllowedActions != nil {
		toSerialize["allowed_actions"] = o.AllowedActions
	}
	return json.Marshal(toSerialize)
}

type NullableActionsSetGithubActionsPermissionsOrganizationRequest struct {
	value *ActionsSetGithubActionsPermissionsOrganizationRequest
	isSet bool
}

func (v NullableActionsSetGithubActionsPermissionsOrganizationRequest) Get() *ActionsSetGithubActionsPermissionsOrganizationRequest {
	return v.value
}

func (v *NullableActionsSetGithubActionsPermissionsOrganizationRequest) Set(val *ActionsSetGithubActionsPermissionsOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsSetGithubActionsPermissionsOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsSetGithubActionsPermissionsOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsSetGithubActionsPermissionsOrganizationRequest(val *ActionsSetGithubActionsPermissionsOrganizationRequest) *NullableActionsSetGithubActionsPermissionsOrganizationRequest {
	return &NullableActionsSetGithubActionsPermissionsOrganizationRequest{value: val, isSet: true}
}

func (v NullableActionsSetGithubActionsPermissionsOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsSetGithubActionsPermissionsOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


