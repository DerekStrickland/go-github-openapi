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

// Email Email
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility NullableString `json:"visibility"`
}

// NewEmail instantiates a new Email object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmail(email string, primary bool, verified bool, visibility NullableString) *Email {
	this := Email{}
	this.Email = email
	this.Primary = primary
	this.Verified = verified
	this.Visibility = visibility
	return &this
}

// NewEmailWithDefaults instantiates a new Email object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailWithDefaults() *Email {
	this := Email{}
	return &this
}

// GetEmail returns the Email field value
func (o *Email) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Email) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Email) SetEmail(v string) {
	o.Email = v
}

// GetPrimary returns the Primary field value
func (o *Email) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *Email) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *Email) SetPrimary(v bool) {
	o.Primary = v
}

// GetVerified returns the Verified field value
func (o *Email) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *Email) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *Email) SetVerified(v bool) {
	o.Verified = v
}

// GetVisibility returns the Visibility field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Email) GetVisibility() string {
	if o == nil || o.Visibility.Get() == nil {
		var ret string
		return ret
	}

	return *o.Visibility.Get()
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Email) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visibility.Get(), o.Visibility.IsSet()
}

// SetVisibility sets field value
func (o *Email) SetVisibility(v string) {
	o.Visibility.Set(&v)
}

func (o Email) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["primary"] = o.Primary
	}
	if true {
		toSerialize["verified"] = o.Verified
	}
	if true {
		toSerialize["visibility"] = o.Visibility.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEmail struct {
	value *Email
	isSet bool
}

func (v NullableEmail) Get() *Email {
	return v.value
}

func (v *NullableEmail) Set(val *Email) {
	v.value = val
	v.isSet = true
}

func (v NullableEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmail(val *Email) *NullableEmail {
	return &NullableEmail{value: val, isSet: true}
}

func (v NullableEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


