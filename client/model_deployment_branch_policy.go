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

// DeploymentBranchPolicy Details of a deployment branch policy.
type DeploymentBranchPolicy struct {
	// The unique identifier of the branch policy.
	Id *int32 `json:"id,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	// The name pattern that branches must match in order to deploy to the environment.
	Name *string `json:"name,omitempty"`
}

// NewDeploymentBranchPolicy instantiates a new DeploymentBranchPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentBranchPolicy() *DeploymentBranchPolicy {
	this := DeploymentBranchPolicy{}
	return &this
}

// NewDeploymentBranchPolicyWithDefaults instantiates a new DeploymentBranchPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentBranchPolicyWithDefaults() *DeploymentBranchPolicy {
	this := DeploymentBranchPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeploymentBranchPolicy) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentBranchPolicy) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentBranchPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeploymentBranchPolicy) SetId(v int32) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *DeploymentBranchPolicy) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentBranchPolicy) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *DeploymentBranchPolicy) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *DeploymentBranchPolicy) SetNodeId(v string) {
	o.NodeId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentBranchPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentBranchPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentBranchPolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentBranchPolicy) SetName(v string) {
	o.Name = &v
}

func (o DeploymentBranchPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentBranchPolicy struct {
	value *DeploymentBranchPolicy
	isSet bool
}

func (v NullableDeploymentBranchPolicy) Get() *DeploymentBranchPolicy {
	return v.value
}

func (v *NullableDeploymentBranchPolicy) Set(val *DeploymentBranchPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentBranchPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentBranchPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentBranchPolicy(val *DeploymentBranchPolicy) *NullableDeploymentBranchPolicy {
	return &NullableDeploymentBranchPolicy{value: val, isSet: true}
}

func (v NullableDeploymentBranchPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentBranchPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


