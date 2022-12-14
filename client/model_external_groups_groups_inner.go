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

// ExternalGroupsGroupsInner struct for ExternalGroupsGroupsInner
type ExternalGroupsGroupsInner struct {
	// The internal ID of the group
	GroupId int32 `json:"group_id"`
	// The display name of the group
	GroupName string `json:"group_name"`
	// The time of the last update for this group
	UpdatedAt string `json:"updated_at"`
}

// NewExternalGroupsGroupsInner instantiates a new ExternalGroupsGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGroupsGroupsInner(groupId int32, groupName string, updatedAt string) *ExternalGroupsGroupsInner {
	this := ExternalGroupsGroupsInner{}
	this.GroupId = groupId
	this.GroupName = groupName
	this.UpdatedAt = updatedAt
	return &this
}

// NewExternalGroupsGroupsInnerWithDefaults instantiates a new ExternalGroupsGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGroupsGroupsInnerWithDefaults() *ExternalGroupsGroupsInner {
	this := ExternalGroupsGroupsInner{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ExternalGroupsGroupsInner) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ExternalGroupsGroupsInner) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ExternalGroupsGroupsInner) SetGroupId(v int32) {
	o.GroupId = v
}

// GetGroupName returns the GroupName field value
func (o *ExternalGroupsGroupsInner) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *ExternalGroupsGroupsInner) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *ExternalGroupsGroupsInner) SetGroupName(v string) {
	o.GroupName = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ExternalGroupsGroupsInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalGroupsGroupsInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ExternalGroupsGroupsInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ExternalGroupsGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	if true {
		toSerialize["group_name"] = o.GroupName
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGroupsGroupsInner struct {
	value *ExternalGroupsGroupsInner
	isSet bool
}

func (v NullableExternalGroupsGroupsInner) Get() *ExternalGroupsGroupsInner {
	return v.value
}

func (v *NullableExternalGroupsGroupsInner) Set(val *ExternalGroupsGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGroupsGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGroupsGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGroupsGroupsInner(val *ExternalGroupsGroupsInner) *NullableExternalGroupsGroupsInner {
	return &NullableExternalGroupsGroupsInner{value: val, isSet: true}
}

func (v NullableExternalGroupsGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGroupsGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


