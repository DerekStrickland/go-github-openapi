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

// ActionsOrganizationPermissions struct for ActionsOrganizationPermissions
type ActionsOrganizationPermissions struct {
	EnabledRepositories EnabledRepositories `json:"enabled_repositories"`
	// The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
	SelectedRepositoriesUrl *string `json:"selected_repositories_url,omitempty"`
	AllowedActions *AllowedActions `json:"allowed_actions,omitempty"`
	// The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	SelectedActionsUrl *string `json:"selected_actions_url,omitempty"`
}

// NewActionsOrganizationPermissions instantiates a new ActionsOrganizationPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsOrganizationPermissions(enabledRepositories EnabledRepositories) *ActionsOrganizationPermissions {
	this := ActionsOrganizationPermissions{}
	this.EnabledRepositories = enabledRepositories
	return &this
}

// NewActionsOrganizationPermissionsWithDefaults instantiates a new ActionsOrganizationPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsOrganizationPermissionsWithDefaults() *ActionsOrganizationPermissions {
	this := ActionsOrganizationPermissions{}
	return &this
}

// GetEnabledRepositories returns the EnabledRepositories field value
func (o *ActionsOrganizationPermissions) GetEnabledRepositories() EnabledRepositories {
	if o == nil {
		var ret EnabledRepositories
		return ret
	}

	return o.EnabledRepositories
}

// GetEnabledRepositoriesOk returns a tuple with the EnabledRepositories field value
// and a boolean to check if the value has been set.
func (o *ActionsOrganizationPermissions) GetEnabledRepositoriesOk() (*EnabledRepositories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledRepositories, true
}

// SetEnabledRepositories sets field value
func (o *ActionsOrganizationPermissions) SetEnabledRepositories(v EnabledRepositories) {
	o.EnabledRepositories = v
}

// GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field value if set, zero value otherwise.
func (o *ActionsOrganizationPermissions) GetSelectedRepositoriesUrl() string {
	if o == nil || o.SelectedRepositoriesUrl == nil {
		var ret string
		return ret
	}
	return *o.SelectedRepositoriesUrl
}

// GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsOrganizationPermissions) GetSelectedRepositoriesUrlOk() (*string, bool) {
	if o == nil || o.SelectedRepositoriesUrl == nil {
		return nil, false
	}
	return o.SelectedRepositoriesUrl, true
}

// HasSelectedRepositoriesUrl returns a boolean if a field has been set.
func (o *ActionsOrganizationPermissions) HasSelectedRepositoriesUrl() bool {
	if o != nil && o.SelectedRepositoriesUrl != nil {
		return true
	}

	return false
}

// SetSelectedRepositoriesUrl gets a reference to the given string and assigns it to the SelectedRepositoriesUrl field.
func (o *ActionsOrganizationPermissions) SetSelectedRepositoriesUrl(v string) {
	o.SelectedRepositoriesUrl = &v
}

// GetAllowedActions returns the AllowedActions field value if set, zero value otherwise.
func (o *ActionsOrganizationPermissions) GetAllowedActions() AllowedActions {
	if o == nil || o.AllowedActions == nil {
		var ret AllowedActions
		return ret
	}
	return *o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsOrganizationPermissions) GetAllowedActionsOk() (*AllowedActions, bool) {
	if o == nil || o.AllowedActions == nil {
		return nil, false
	}
	return o.AllowedActions, true
}

// HasAllowedActions returns a boolean if a field has been set.
func (o *ActionsOrganizationPermissions) HasAllowedActions() bool {
	if o != nil && o.AllowedActions != nil {
		return true
	}

	return false
}

// SetAllowedActions gets a reference to the given AllowedActions and assigns it to the AllowedActions field.
func (o *ActionsOrganizationPermissions) SetAllowedActions(v AllowedActions) {
	o.AllowedActions = &v
}

// GetSelectedActionsUrl returns the SelectedActionsUrl field value if set, zero value otherwise.
func (o *ActionsOrganizationPermissions) GetSelectedActionsUrl() string {
	if o == nil || o.SelectedActionsUrl == nil {
		var ret string
		return ret
	}
	return *o.SelectedActionsUrl
}

// GetSelectedActionsUrlOk returns a tuple with the SelectedActionsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsOrganizationPermissions) GetSelectedActionsUrlOk() (*string, bool) {
	if o == nil || o.SelectedActionsUrl == nil {
		return nil, false
	}
	return o.SelectedActionsUrl, true
}

// HasSelectedActionsUrl returns a boolean if a field has been set.
func (o *ActionsOrganizationPermissions) HasSelectedActionsUrl() bool {
	if o != nil && o.SelectedActionsUrl != nil {
		return true
	}

	return false
}

// SetSelectedActionsUrl gets a reference to the given string and assigns it to the SelectedActionsUrl field.
func (o *ActionsOrganizationPermissions) SetSelectedActionsUrl(v string) {
	o.SelectedActionsUrl = &v
}

func (o ActionsOrganizationPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled_repositories"] = o.EnabledRepositories
	}
	if o.SelectedRepositoriesUrl != nil {
		toSerialize["selected_repositories_url"] = o.SelectedRepositoriesUrl
	}
	if o.AllowedActions != nil {
		toSerialize["allowed_actions"] = o.AllowedActions
	}
	if o.SelectedActionsUrl != nil {
		toSerialize["selected_actions_url"] = o.SelectedActionsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableActionsOrganizationPermissions struct {
	value *ActionsOrganizationPermissions
	isSet bool
}

func (v NullableActionsOrganizationPermissions) Get() *ActionsOrganizationPermissions {
	return v.value
}

func (v *NullableActionsOrganizationPermissions) Set(val *ActionsOrganizationPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsOrganizationPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsOrganizationPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsOrganizationPermissions(val *ActionsOrganizationPermissions) *NullableActionsOrganizationPermissions {
	return &NullableActionsOrganizationPermissions{value: val, isSet: true}
}

func (v NullableActionsOrganizationPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsOrganizationPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


