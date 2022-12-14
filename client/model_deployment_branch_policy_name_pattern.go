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

// DeploymentBranchPolicyNamePattern struct for DeploymentBranchPolicyNamePattern
type DeploymentBranchPolicyNamePattern struct {
	// The name pattern that branches must match in order to deploy to the environment.  Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/_*_/_*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
	Name string `json:"name"`
}

// NewDeploymentBranchPolicyNamePattern instantiates a new DeploymentBranchPolicyNamePattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentBranchPolicyNamePattern(name string) *DeploymentBranchPolicyNamePattern {
	this := DeploymentBranchPolicyNamePattern{}
	this.Name = name
	return &this
}

// NewDeploymentBranchPolicyNamePatternWithDefaults instantiates a new DeploymentBranchPolicyNamePattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentBranchPolicyNamePatternWithDefaults() *DeploymentBranchPolicyNamePattern {
	this := DeploymentBranchPolicyNamePattern{}
	return &this
}

// GetName returns the Name field value
func (o *DeploymentBranchPolicyNamePattern) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeploymentBranchPolicyNamePattern) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeploymentBranchPolicyNamePattern) SetName(v string) {
	o.Name = v
}

func (o DeploymentBranchPolicyNamePattern) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentBranchPolicyNamePattern struct {
	value *DeploymentBranchPolicyNamePattern
	isSet bool
}

func (v NullableDeploymentBranchPolicyNamePattern) Get() *DeploymentBranchPolicyNamePattern {
	return v.value
}

func (v *NullableDeploymentBranchPolicyNamePattern) Set(val *DeploymentBranchPolicyNamePattern) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentBranchPolicyNamePattern) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentBranchPolicyNamePattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentBranchPolicyNamePattern(val *DeploymentBranchPolicyNamePattern) *NullableDeploymentBranchPolicyNamePattern {
	return &NullableDeploymentBranchPolicyNamePattern{value: val, isSet: true}
}

func (v NullableDeploymentBranchPolicyNamePattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentBranchPolicyNamePattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


