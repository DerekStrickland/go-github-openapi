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

// ActionsWorkflowAccessToRepository struct for ActionsWorkflowAccessToRepository
type ActionsWorkflowAccessToRepository struct {
	// Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means access is only possible from workflows in this repository.
	AccessLevel string `json:"access_level"`
}

// NewActionsWorkflowAccessToRepository instantiates a new ActionsWorkflowAccessToRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsWorkflowAccessToRepository(accessLevel string) *ActionsWorkflowAccessToRepository {
	this := ActionsWorkflowAccessToRepository{}
	this.AccessLevel = accessLevel
	return &this
}

// NewActionsWorkflowAccessToRepositoryWithDefaults instantiates a new ActionsWorkflowAccessToRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsWorkflowAccessToRepositoryWithDefaults() *ActionsWorkflowAccessToRepository {
	this := ActionsWorkflowAccessToRepository{}
	return &this
}

// GetAccessLevel returns the AccessLevel field value
func (o *ActionsWorkflowAccessToRepository) GetAccessLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value
// and a boolean to check if the value has been set.
func (o *ActionsWorkflowAccessToRepository) GetAccessLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevel, true
}

// SetAccessLevel sets field value
func (o *ActionsWorkflowAccessToRepository) SetAccessLevel(v string) {
	o.AccessLevel = v
}

func (o ActionsWorkflowAccessToRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_level"] = o.AccessLevel
	}
	return json.Marshal(toSerialize)
}

type NullableActionsWorkflowAccessToRepository struct {
	value *ActionsWorkflowAccessToRepository
	isSet bool
}

func (v NullableActionsWorkflowAccessToRepository) Get() *ActionsWorkflowAccessToRepository {
	return v.value
}

func (v *NullableActionsWorkflowAccessToRepository) Set(val *ActionsWorkflowAccessToRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsWorkflowAccessToRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsWorkflowAccessToRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsWorkflowAccessToRepository(val *ActionsWorkflowAccessToRepository) *NullableActionsWorkflowAccessToRepository {
	return &NullableActionsWorkflowAccessToRepository{value: val, isSet: true}
}

func (v NullableActionsWorkflowAccessToRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsWorkflowAccessToRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


