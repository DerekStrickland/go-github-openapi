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

// PublicUserPlan struct for PublicUserPlan
type PublicUserPlan struct {
	Collaborators int32 `json:"collaborators"`
	Name string `json:"name"`
	Space int32 `json:"space"`
	PrivateRepos int32 `json:"private_repos"`
}

// NewPublicUserPlan instantiates a new PublicUserPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicUserPlan(collaborators int32, name string, space int32, privateRepos int32) *PublicUserPlan {
	this := PublicUserPlan{}
	this.Collaborators = collaborators
	this.Name = name
	this.Space = space
	this.PrivateRepos = privateRepos
	return &this
}

// NewPublicUserPlanWithDefaults instantiates a new PublicUserPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicUserPlanWithDefaults() *PublicUserPlan {
	this := PublicUserPlan{}
	return &this
}

// GetCollaborators returns the Collaborators field value
func (o *PublicUserPlan) GetCollaborators() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Collaborators
}

// GetCollaboratorsOk returns a tuple with the Collaborators field value
// and a boolean to check if the value has been set.
func (o *PublicUserPlan) GetCollaboratorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collaborators, true
}

// SetCollaborators sets field value
func (o *PublicUserPlan) SetCollaborators(v int32) {
	o.Collaborators = v
}

// GetName returns the Name field value
func (o *PublicUserPlan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicUserPlan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicUserPlan) SetName(v string) {
	o.Name = v
}

// GetSpace returns the Space field value
func (o *PublicUserPlan) GetSpace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *PublicUserPlan) GetSpaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *PublicUserPlan) SetSpace(v int32) {
	o.Space = v
}

// GetPrivateRepos returns the PrivateRepos field value
func (o *PublicUserPlan) GetPrivateRepos() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrivateRepos
}

// GetPrivateReposOk returns a tuple with the PrivateRepos field value
// and a boolean to check if the value has been set.
func (o *PublicUserPlan) GetPrivateReposOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateRepos, true
}

// SetPrivateRepos sets field value
func (o *PublicUserPlan) SetPrivateRepos(v int32) {
	o.PrivateRepos = v
}

func (o PublicUserPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["collaborators"] = o.Collaborators
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["space"] = o.Space
	}
	if true {
		toSerialize["private_repos"] = o.PrivateRepos
	}
	return json.Marshal(toSerialize)
}

type NullablePublicUserPlan struct {
	value *PublicUserPlan
	isSet bool
}

func (v NullablePublicUserPlan) Get() *PublicUserPlan {
	return v.value
}

func (v *NullablePublicUserPlan) Set(val *PublicUserPlan) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicUserPlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicUserPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicUserPlan(val *PublicUserPlan) *NullablePublicUserPlan {
	return &NullablePublicUserPlan{value: val, isSet: true}
}

func (v NullablePublicUserPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicUserPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


