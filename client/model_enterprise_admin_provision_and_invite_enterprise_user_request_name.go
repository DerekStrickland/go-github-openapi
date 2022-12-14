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

// EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName struct for EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName
type EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName struct {
	// The first name of the user.
	GivenName string `json:"givenName"`
	// The last name of the user.
	FamilyName string `json:"familyName"`
}

// NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName(givenName string, familyName string) *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName {
	this := EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName{}
	this.GivenName = givenName
	this.FamilyName = familyName
	return &this
}

// NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestNameWithDefaults instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestNameWithDefaults() *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName {
	this := EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName{}
	return &this
}

// GetGivenName returns the GivenName field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) GetGivenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) GetFamilyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) SetFamilyName(v string) {
	o.FamilyName = v
}

func (o EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["givenName"] = o.GivenName
	}
	if true {
		toSerialize["familyName"] = o.FamilyName
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName struct {
	value *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName
	isSet bool
}

func (v NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) Get() *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName {
	return v.value
}

func (v *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) Set(val *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName(val *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName {
	return &NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName{value: val, isSet: true}
}

func (v NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


