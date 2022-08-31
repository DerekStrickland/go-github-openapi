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

// UsersCreateGpgKeyForAuthenticatedUserRequest struct for UsersCreateGpgKeyForAuthenticatedUserRequest
type UsersCreateGpgKeyForAuthenticatedUserRequest struct {
	// A descriptive name for the new key.
	Name *string `json:"name,omitempty"`
	// A GPG key in ASCII-armored format.
	ArmoredPublicKey string `json:"armored_public_key"`
}

// NewUsersCreateGpgKeyForAuthenticatedUserRequest instantiates a new UsersCreateGpgKeyForAuthenticatedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersCreateGpgKeyForAuthenticatedUserRequest(armoredPublicKey string) *UsersCreateGpgKeyForAuthenticatedUserRequest {
	this := UsersCreateGpgKeyForAuthenticatedUserRequest{}
	this.ArmoredPublicKey = armoredPublicKey
	return &this
}

// NewUsersCreateGpgKeyForAuthenticatedUserRequestWithDefaults instantiates a new UsersCreateGpgKeyForAuthenticatedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersCreateGpgKeyForAuthenticatedUserRequestWithDefaults() *UsersCreateGpgKeyForAuthenticatedUserRequest {
	this := UsersCreateGpgKeyForAuthenticatedUserRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) SetName(v string) {
	o.Name = &v
}

// GetArmoredPublicKey returns the ArmoredPublicKey field value
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetArmoredPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArmoredPublicKey
}

// GetArmoredPublicKeyOk returns a tuple with the ArmoredPublicKey field value
// and a boolean to check if the value has been set.
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetArmoredPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArmoredPublicKey, true
}

// SetArmoredPublicKey sets field value
func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) SetArmoredPublicKey(v string) {
	o.ArmoredPublicKey = v
}

func (o UsersCreateGpgKeyForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["armored_public_key"] = o.ArmoredPublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableUsersCreateGpgKeyForAuthenticatedUserRequest struct {
	value *UsersCreateGpgKeyForAuthenticatedUserRequest
	isSet bool
}

func (v NullableUsersCreateGpgKeyForAuthenticatedUserRequest) Get() *UsersCreateGpgKeyForAuthenticatedUserRequest {
	return v.value
}

func (v *NullableUsersCreateGpgKeyForAuthenticatedUserRequest) Set(val *UsersCreateGpgKeyForAuthenticatedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersCreateGpgKeyForAuthenticatedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersCreateGpgKeyForAuthenticatedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersCreateGpgKeyForAuthenticatedUserRequest(val *UsersCreateGpgKeyForAuthenticatedUserRequest) *NullableUsersCreateGpgKeyForAuthenticatedUserRequest {
	return &NullableUsersCreateGpgKeyForAuthenticatedUserRequest{value: val, isSet: true}
}

func (v NullableUsersCreateGpgKeyForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersCreateGpgKeyForAuthenticatedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

