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

// ReposSetUserAccessRestrictionsRequestOneOf struct for ReposSetUserAccessRestrictionsRequestOneOf
type ReposSetUserAccessRestrictionsRequestOneOf struct {
	// The username for users
	Users []string `json:"users"`
}

// NewReposSetUserAccessRestrictionsRequestOneOf instantiates a new ReposSetUserAccessRestrictionsRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposSetUserAccessRestrictionsRequestOneOf(users []string) *ReposSetUserAccessRestrictionsRequestOneOf {
	this := ReposSetUserAccessRestrictionsRequestOneOf{}
	this.Users = users
	return &this
}

// NewReposSetUserAccessRestrictionsRequestOneOfWithDefaults instantiates a new ReposSetUserAccessRestrictionsRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposSetUserAccessRestrictionsRequestOneOfWithDefaults() *ReposSetUserAccessRestrictionsRequestOneOf {
	this := ReposSetUserAccessRestrictionsRequestOneOf{}
	return &this
}

// GetUsers returns the Users field value
func (o *ReposSetUserAccessRestrictionsRequestOneOf) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ReposSetUserAccessRestrictionsRequestOneOf) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ReposSetUserAccessRestrictionsRequestOneOf) SetUsers(v []string) {
	o.Users = v
}

func (o ReposSetUserAccessRestrictionsRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableReposSetUserAccessRestrictionsRequestOneOf struct {
	value *ReposSetUserAccessRestrictionsRequestOneOf
	isSet bool
}

func (v NullableReposSetUserAccessRestrictionsRequestOneOf) Get() *ReposSetUserAccessRestrictionsRequestOneOf {
	return v.value
}

func (v *NullableReposSetUserAccessRestrictionsRequestOneOf) Set(val *ReposSetUserAccessRestrictionsRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReposSetUserAccessRestrictionsRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReposSetUserAccessRestrictionsRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposSetUserAccessRestrictionsRequestOneOf(val *ReposSetUserAccessRestrictionsRequestOneOf) *NullableReposSetUserAccessRestrictionsRequestOneOf {
	return &NullableReposSetUserAccessRestrictionsRequestOneOf{value: val, isSet: true}
}

func (v NullableReposSetUserAccessRestrictionsRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposSetUserAccessRestrictionsRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

