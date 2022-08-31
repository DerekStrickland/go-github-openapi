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

// EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner struct for EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner
type EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner struct {
	// The email address.
	Value string `json:"value"`
	// The type of email address.
	Type string `json:"type"`
	// Whether this email address is the primary address.
	Primary bool `json:"primary"`
}

// NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner(value string, type_ string, primary bool) *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner {
	this := EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner{}
	this.Value = value
	this.Type = type_
	this.Primary = primary
	return &this
}

// NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInnerWithDefaults instantiates a new EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInnerWithDefaults() *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner {
	this := EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner{}
	return &this
}

// GetValue returns the Value field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) SetType(v string) {
	o.Type = v
}

// GetPrimary returns the Primary field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) SetPrimary(v bool) {
	o.Primary = v
}

func (o EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["primary"] = o.Primary
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner struct {
	value *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner
	isSet bool
}

func (v NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) Get() *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner {
	return v.value
}

func (v *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) Set(val *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner(val *EnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner {
	return &NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner{value: val, isSet: true}
}

func (v NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAdminProvisionAndInviteEnterpriseUserRequestEmailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


