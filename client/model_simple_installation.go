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

// SimpleInstallation Simple Installation
type SimpleInstallation struct {
	// The ID of the installation.
	Id int32 `json:"id"`
	// The global node ID of the installation.
	NodeId string `json:"node_id"`
}

// NewSimpleInstallation instantiates a new SimpleInstallation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleInstallation(id int32, nodeId string) *SimpleInstallation {
	this := SimpleInstallation{}
	this.Id = id
	this.NodeId = nodeId
	return &this
}

// NewSimpleInstallationWithDefaults instantiates a new SimpleInstallation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleInstallationWithDefaults() *SimpleInstallation {
	this := SimpleInstallation{}
	return &this
}

// GetId returns the Id field value
func (o *SimpleInstallation) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleInstallation) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleInstallation) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *SimpleInstallation) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *SimpleInstallation) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *SimpleInstallation) SetNodeId(v string) {
	o.NodeId = v
}

func (o SimpleInstallation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleInstallation struct {
	value *SimpleInstallation
	isSet bool
}

func (v NullableSimpleInstallation) Get() *SimpleInstallation {
	return v.value
}

func (v *NullableSimpleInstallation) Set(val *SimpleInstallation) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleInstallation) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleInstallation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleInstallation(val *SimpleInstallation) *NullableSimpleInstallation {
	return &NullableSimpleInstallation{value: val, isSet: true}
}

func (v NullableSimpleInstallation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleInstallation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


