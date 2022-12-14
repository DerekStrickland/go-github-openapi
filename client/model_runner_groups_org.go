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

// RunnerGroupsOrg struct for RunnerGroupsOrg
type RunnerGroupsOrg struct {
	Id float32 `json:"id"`
	Name string `json:"name"`
	Visibility string `json:"visibility"`
	Default bool `json:"default"`
	// Link to the selected repositories resource for this runner group. Not present unless visibility was set to `selected`
	SelectedRepositoriesUrl *string `json:"selected_repositories_url,omitempty"`
	RunnersUrl string `json:"runners_url"`
	Inherited bool `json:"inherited"`
	InheritedAllowsPublicRepositories *bool `json:"inherited_allows_public_repositories,omitempty"`
	AllowsPublicRepositories bool `json:"allows_public_repositories"`
	// If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
	WorkflowRestrictionsReadOnly *bool `json:"workflow_restrictions_read_only,omitempty"`
	// If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	RestrictedToWorkflows *bool `json:"restricted_to_workflows,omitempty"`
	// List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	SelectedWorkflows []string `json:"selected_workflows,omitempty"`
}

// NewRunnerGroupsOrg instantiates a new RunnerGroupsOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunnerGroupsOrg(id float32, name string, visibility string, default_ bool, runnersUrl string, inherited bool, allowsPublicRepositories bool) *RunnerGroupsOrg {
	this := RunnerGroupsOrg{}
	this.Id = id
	this.Name = name
	this.Visibility = visibility
	this.Default = default_
	this.RunnersUrl = runnersUrl
	this.Inherited = inherited
	this.AllowsPublicRepositories = allowsPublicRepositories
	var workflowRestrictionsReadOnly bool = false
	this.WorkflowRestrictionsReadOnly = &workflowRestrictionsReadOnly
	var restrictedToWorkflows bool = false
	this.RestrictedToWorkflows = &restrictedToWorkflows
	return &this
}

// NewRunnerGroupsOrgWithDefaults instantiates a new RunnerGroupsOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunnerGroupsOrgWithDefaults() *RunnerGroupsOrg {
	this := RunnerGroupsOrg{}
	var workflowRestrictionsReadOnly bool = false
	this.WorkflowRestrictionsReadOnly = &workflowRestrictionsReadOnly
	var restrictedToWorkflows bool = false
	this.RestrictedToWorkflows = &restrictedToWorkflows
	return &this
}

// GetId returns the Id field value
func (o *RunnerGroupsOrg) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunnerGroupsOrg) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RunnerGroupsOrg) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RunnerGroupsOrg) SetName(v string) {
	o.Name = v
}

// GetVisibility returns the Visibility field value
func (o *RunnerGroupsOrg) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *RunnerGroupsOrg) SetVisibility(v string) {
	o.Visibility = v
}

// GetDefault returns the Default field value
func (o *RunnerGroupsOrg) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *RunnerGroupsOrg) SetDefault(v bool) {
	o.Default = v
}

// GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field value if set, zero value otherwise.
func (o *RunnerGroupsOrg) GetSelectedRepositoriesUrl() string {
	if o == nil || o.SelectedRepositoriesUrl == nil {
		var ret string
		return ret
	}
	return *o.SelectedRepositoriesUrl
}

// GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetSelectedRepositoriesUrlOk() (*string, bool) {
	if o == nil || o.SelectedRepositoriesUrl == nil {
		return nil, false
	}
	return o.SelectedRepositoriesUrl, true
}

// HasSelectedRepositoriesUrl returns a boolean if a field has been set.
func (o *RunnerGroupsOrg) HasSelectedRepositoriesUrl() bool {
	if o != nil && o.SelectedRepositoriesUrl != nil {
		return true
	}

	return false
}

// SetSelectedRepositoriesUrl gets a reference to the given string and assigns it to the SelectedRepositoriesUrl field.
func (o *RunnerGroupsOrg) SetSelectedRepositoriesUrl(v string) {
	o.SelectedRepositoriesUrl = &v
}

// GetRunnersUrl returns the RunnersUrl field value
func (o *RunnerGroupsOrg) GetRunnersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunnersUrl
}

// GetRunnersUrlOk returns a tuple with the RunnersUrl field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetRunnersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnersUrl, true
}

// SetRunnersUrl sets field value
func (o *RunnerGroupsOrg) SetRunnersUrl(v string) {
	o.RunnersUrl = v
}

// GetInherited returns the Inherited field value
func (o *RunnerGroupsOrg) GetInherited() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetInheritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inherited, true
}

// SetInherited sets field value
func (o *RunnerGroupsOrg) SetInherited(v bool) {
	o.Inherited = v
}

// GetInheritedAllowsPublicRepositories returns the InheritedAllowsPublicRepositories field value if set, zero value otherwise.
func (o *RunnerGroupsOrg) GetInheritedAllowsPublicRepositories() bool {
	if o == nil || o.InheritedAllowsPublicRepositories == nil {
		var ret bool
		return ret
	}
	return *o.InheritedAllowsPublicRepositories
}

// GetInheritedAllowsPublicRepositoriesOk returns a tuple with the InheritedAllowsPublicRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetInheritedAllowsPublicRepositoriesOk() (*bool, bool) {
	if o == nil || o.InheritedAllowsPublicRepositories == nil {
		return nil, false
	}
	return o.InheritedAllowsPublicRepositories, true
}

// HasInheritedAllowsPublicRepositories returns a boolean if a field has been set.
func (o *RunnerGroupsOrg) HasInheritedAllowsPublicRepositories() bool {
	if o != nil && o.InheritedAllowsPublicRepositories != nil {
		return true
	}

	return false
}

// SetInheritedAllowsPublicRepositories gets a reference to the given bool and assigns it to the InheritedAllowsPublicRepositories field.
func (o *RunnerGroupsOrg) SetInheritedAllowsPublicRepositories(v bool) {
	o.InheritedAllowsPublicRepositories = &v
}

// GetAllowsPublicRepositories returns the AllowsPublicRepositories field value
func (o *RunnerGroupsOrg) GetAllowsPublicRepositories() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowsPublicRepositories
}

// GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field value
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetAllowsPublicRepositoriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowsPublicRepositories, true
}

// SetAllowsPublicRepositories sets field value
func (o *RunnerGroupsOrg) SetAllowsPublicRepositories(v bool) {
	o.AllowsPublicRepositories = v
}

// GetWorkflowRestrictionsReadOnly returns the WorkflowRestrictionsReadOnly field value if set, zero value otherwise.
func (o *RunnerGroupsOrg) GetWorkflowRestrictionsReadOnly() bool {
	if o == nil || o.WorkflowRestrictionsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.WorkflowRestrictionsReadOnly
}

// GetWorkflowRestrictionsReadOnlyOk returns a tuple with the WorkflowRestrictionsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetWorkflowRestrictionsReadOnlyOk() (*bool, bool) {
	if o == nil || o.WorkflowRestrictionsReadOnly == nil {
		return nil, false
	}
	return o.WorkflowRestrictionsReadOnly, true
}

// HasWorkflowRestrictionsReadOnly returns a boolean if a field has been set.
func (o *RunnerGroupsOrg) HasWorkflowRestrictionsReadOnly() bool {
	if o != nil && o.WorkflowRestrictionsReadOnly != nil {
		return true
	}

	return false
}

// SetWorkflowRestrictionsReadOnly gets a reference to the given bool and assigns it to the WorkflowRestrictionsReadOnly field.
func (o *RunnerGroupsOrg) SetWorkflowRestrictionsReadOnly(v bool) {
	o.WorkflowRestrictionsReadOnly = &v
}

// GetRestrictedToWorkflows returns the RestrictedToWorkflows field value if set, zero value otherwise.
func (o *RunnerGroupsOrg) GetRestrictedToWorkflows() bool {
	if o == nil || o.RestrictedToWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.RestrictedToWorkflows
}

// GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetRestrictedToWorkflowsOk() (*bool, bool) {
	if o == nil || o.RestrictedToWorkflows == nil {
		return nil, false
	}
	return o.RestrictedToWorkflows, true
}

// HasRestrictedToWorkflows returns a boolean if a field has been set.
func (o *RunnerGroupsOrg) HasRestrictedToWorkflows() bool {
	if o != nil && o.RestrictedToWorkflows != nil {
		return true
	}

	return false
}

// SetRestrictedToWorkflows gets a reference to the given bool and assigns it to the RestrictedToWorkflows field.
func (o *RunnerGroupsOrg) SetRestrictedToWorkflows(v bool) {
	o.RestrictedToWorkflows = &v
}

// GetSelectedWorkflows returns the SelectedWorkflows field value if set, zero value otherwise.
func (o *RunnerGroupsOrg) GetSelectedWorkflows() []string {
	if o == nil || o.SelectedWorkflows == nil {
		var ret []string
		return ret
	}
	return o.SelectedWorkflows
}

// GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerGroupsOrg) GetSelectedWorkflowsOk() ([]string, bool) {
	if o == nil || o.SelectedWorkflows == nil {
		return nil, false
	}
	return o.SelectedWorkflows, true
}

// HasSelectedWorkflows returns a boolean if a field has been set.
func (o *RunnerGroupsOrg) HasSelectedWorkflows() bool {
	if o != nil && o.SelectedWorkflows != nil {
		return true
	}

	return false
}

// SetSelectedWorkflows gets a reference to the given []string and assigns it to the SelectedWorkflows field.
func (o *RunnerGroupsOrg) SetSelectedWorkflows(v []string) {
	o.SelectedWorkflows = v
}

func (o RunnerGroupsOrg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if true {
		toSerialize["default"] = o.Default
	}
	if o.SelectedRepositoriesUrl != nil {
		toSerialize["selected_repositories_url"] = o.SelectedRepositoriesUrl
	}
	if true {
		toSerialize["runners_url"] = o.RunnersUrl
	}
	if true {
		toSerialize["inherited"] = o.Inherited
	}
	if o.InheritedAllowsPublicRepositories != nil {
		toSerialize["inherited_allows_public_repositories"] = o.InheritedAllowsPublicRepositories
	}
	if true {
		toSerialize["allows_public_repositories"] = o.AllowsPublicRepositories
	}
	if o.WorkflowRestrictionsReadOnly != nil {
		toSerialize["workflow_restrictions_read_only"] = o.WorkflowRestrictionsReadOnly
	}
	if o.RestrictedToWorkflows != nil {
		toSerialize["restricted_to_workflows"] = o.RestrictedToWorkflows
	}
	if o.SelectedWorkflows != nil {
		toSerialize["selected_workflows"] = o.SelectedWorkflows
	}
	return json.Marshal(toSerialize)
}

type NullableRunnerGroupsOrg struct {
	value *RunnerGroupsOrg
	isSet bool
}

func (v NullableRunnerGroupsOrg) Get() *RunnerGroupsOrg {
	return v.value
}

func (v *NullableRunnerGroupsOrg) Set(val *RunnerGroupsOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableRunnerGroupsOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableRunnerGroupsOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunnerGroupsOrg(val *RunnerGroupsOrg) *NullableRunnerGroupsOrg {
	return &NullableRunnerGroupsOrg{value: val, isSet: true}
}

func (v NullableRunnerGroupsOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunnerGroupsOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


