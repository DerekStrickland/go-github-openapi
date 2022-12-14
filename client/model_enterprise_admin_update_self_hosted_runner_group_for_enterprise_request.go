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

// EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest struct for EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest
type EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest struct {
	// Name of the runner group.
	Name *string `json:"name,omitempty"`
	// Visibility of a runner group. You can select all organizations or select individual organizations.
	Visibility *string `json:"visibility,omitempty"`
	// Whether the runner group can be used by `public` repositories.
	AllowsPublicRepositories *bool `json:"allows_public_repositories,omitempty"`
	// If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	RestrictedToWorkflows *bool `json:"restricted_to_workflows,omitempty"`
	// List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	SelectedWorkflows []string `json:"selected_workflows,omitempty"`
}

// NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest instantiates a new EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest() *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest {
	this := EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest{}
	var visibility string = "all"
	this.Visibility = &visibility
	var allowsPublicRepositories bool = false
	this.AllowsPublicRepositories = &allowsPublicRepositories
	var restrictedToWorkflows bool = false
	this.RestrictedToWorkflows = &restrictedToWorkflows
	return &this
}

// NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults instantiates a new EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequestWithDefaults() *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest {
	this := EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest{}
	var visibility string = "all"
	this.Visibility = &visibility
	var allowsPublicRepositories bool = false
	this.AllowsPublicRepositories = &allowsPublicRepositories
	var restrictedToWorkflows bool = false
	this.RestrictedToWorkflows = &restrictedToWorkflows
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetName(v string) {
	o.Name = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetVisibility(v string) {
	o.Visibility = &v
}

// GetAllowsPublicRepositories returns the AllowsPublicRepositories field value if set, zero value otherwise.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetAllowsPublicRepositories() bool {
	if o == nil || o.AllowsPublicRepositories == nil {
		var ret bool
		return ret
	}
	return *o.AllowsPublicRepositories
}

// GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetAllowsPublicRepositoriesOk() (*bool, bool) {
	if o == nil || o.AllowsPublicRepositories == nil {
		return nil, false
	}
	return o.AllowsPublicRepositories, true
}

// HasAllowsPublicRepositories returns a boolean if a field has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasAllowsPublicRepositories() bool {
	if o != nil && o.AllowsPublicRepositories != nil {
		return true
	}

	return false
}

// SetAllowsPublicRepositories gets a reference to the given bool and assigns it to the AllowsPublicRepositories field.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetAllowsPublicRepositories(v bool) {
	o.AllowsPublicRepositories = &v
}

// GetRestrictedToWorkflows returns the RestrictedToWorkflows field value if set, zero value otherwise.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetRestrictedToWorkflows() bool {
	if o == nil || o.RestrictedToWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.RestrictedToWorkflows
}

// GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetRestrictedToWorkflowsOk() (*bool, bool) {
	if o == nil || o.RestrictedToWorkflows == nil {
		return nil, false
	}
	return o.RestrictedToWorkflows, true
}

// HasRestrictedToWorkflows returns a boolean if a field has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasRestrictedToWorkflows() bool {
	if o != nil && o.RestrictedToWorkflows != nil {
		return true
	}

	return false
}

// SetRestrictedToWorkflows gets a reference to the given bool and assigns it to the RestrictedToWorkflows field.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetRestrictedToWorkflows(v bool) {
	o.RestrictedToWorkflows = &v
}

// GetSelectedWorkflows returns the SelectedWorkflows field value if set, zero value otherwise.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedWorkflows() []string {
	if o == nil || o.SelectedWorkflows == nil {
		var ret []string
		return ret
	}
	return o.SelectedWorkflows
}

// GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) GetSelectedWorkflowsOk() ([]string, bool) {
	if o == nil || o.SelectedWorkflows == nil {
		return nil, false
	}
	return o.SelectedWorkflows, true
}

// HasSelectedWorkflows returns a boolean if a field has been set.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) HasSelectedWorkflows() bool {
	if o != nil && o.SelectedWorkflows != nil {
		return true
	}

	return false
}

// SetSelectedWorkflows gets a reference to the given []string and assigns it to the SelectedWorkflows field.
func (o *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) SetSelectedWorkflows(v []string) {
	o.SelectedWorkflows = v
}

func (o EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.AllowsPublicRepositories != nil {
		toSerialize["allows_public_repositories"] = o.AllowsPublicRepositories
	}
	if o.RestrictedToWorkflows != nil {
		toSerialize["restricted_to_workflows"] = o.RestrictedToWorkflows
	}
	if o.SelectedWorkflows != nil {
		toSerialize["selected_workflows"] = o.SelectedWorkflows
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest struct {
	value *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest
	isSet bool
}

func (v NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) Get() *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest {
	return v.value
}

func (v *NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) Set(val *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest(val *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) *NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest {
	return &NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest{value: val, isSet: true}
}

func (v NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


