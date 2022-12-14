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

// GroupMapping External Groups to be mapped to a team for membership
type GroupMapping struct {
	// Array of groups to be mapped to this team
	Groups []GroupMappingGroupsInner `json:"groups,omitempty"`
}

// NewGroupMapping instantiates a new GroupMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMapping() *GroupMapping {
	this := GroupMapping{}
	return &this
}

// NewGroupMappingWithDefaults instantiates a new GroupMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMappingWithDefaults() *GroupMapping {
	this := GroupMapping{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *GroupMapping) GetGroups() []GroupMappingGroupsInner {
	if o == nil || o.Groups == nil {
		var ret []GroupMappingGroupsInner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMapping) GetGroupsOk() ([]GroupMappingGroupsInner, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupMapping) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupMappingGroupsInner and assigns it to the Groups field.
func (o *GroupMapping) SetGroups(v []GroupMappingGroupsInner) {
	o.Groups = v
}

func (o GroupMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableGroupMapping struct {
	value *GroupMapping
	isSet bool
}

func (v NullableGroupMapping) Get() *GroupMapping {
	return v.value
}

func (v *NullableGroupMapping) Set(val *GroupMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMapping(val *GroupMapping) *NullableGroupMapping {
	return &NullableGroupMapping{value: val, isSet: true}
}

func (v NullableGroupMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


