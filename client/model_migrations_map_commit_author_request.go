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

// MigrationsMapCommitAuthorRequest struct for MigrationsMapCommitAuthorRequest
type MigrationsMapCommitAuthorRequest struct {
	// The new Git author email.
	Email *string `json:"email,omitempty"`
	// The new Git author name.
	Name *string `json:"name,omitempty"`
}

// NewMigrationsMapCommitAuthorRequest instantiates a new MigrationsMapCommitAuthorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationsMapCommitAuthorRequest() *MigrationsMapCommitAuthorRequest {
	this := MigrationsMapCommitAuthorRequest{}
	return &this
}

// NewMigrationsMapCommitAuthorRequestWithDefaults instantiates a new MigrationsMapCommitAuthorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationsMapCommitAuthorRequestWithDefaults() *MigrationsMapCommitAuthorRequest {
	this := MigrationsMapCommitAuthorRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MigrationsMapCommitAuthorRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsMapCommitAuthorRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MigrationsMapCommitAuthorRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MigrationsMapCommitAuthorRequest) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MigrationsMapCommitAuthorRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsMapCommitAuthorRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MigrationsMapCommitAuthorRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MigrationsMapCommitAuthorRequest) SetName(v string) {
	o.Name = &v
}

func (o MigrationsMapCommitAuthorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableMigrationsMapCommitAuthorRequest struct {
	value *MigrationsMapCommitAuthorRequest
	isSet bool
}

func (v NullableMigrationsMapCommitAuthorRequest) Get() *MigrationsMapCommitAuthorRequest {
	return v.value
}

func (v *NullableMigrationsMapCommitAuthorRequest) Set(val *MigrationsMapCommitAuthorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationsMapCommitAuthorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationsMapCommitAuthorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationsMapCommitAuthorRequest(val *MigrationsMapCommitAuthorRequest) *NullableMigrationsMapCommitAuthorRequest {
	return &NullableMigrationsMapCommitAuthorRequest{value: val, isSet: true}
}

func (v NullableMigrationsMapCommitAuthorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationsMapCommitAuthorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


