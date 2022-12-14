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

// ActionsCreateOrUpdateRepoSecretRequest struct for ActionsCreateOrUpdateRepoSecretRequest
type ActionsCreateOrUpdateRepoSecretRequest struct {
	// Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get a repository public key](https://docs.github.com/rest/reference/actions#get-a-repository-public-key) endpoint.
	EncryptedValue *string `json:"encrypted_value,omitempty"`
	// ID of the key you used to encrypt the secret.
	KeyId *string `json:"key_id,omitempty"`
}

// NewActionsCreateOrUpdateRepoSecretRequest instantiates a new ActionsCreateOrUpdateRepoSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsCreateOrUpdateRepoSecretRequest() *ActionsCreateOrUpdateRepoSecretRequest {
	this := ActionsCreateOrUpdateRepoSecretRequest{}
	return &this
}

// NewActionsCreateOrUpdateRepoSecretRequestWithDefaults instantiates a new ActionsCreateOrUpdateRepoSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsCreateOrUpdateRepoSecretRequestWithDefaults() *ActionsCreateOrUpdateRepoSecretRequest {
	this := ActionsCreateOrUpdateRepoSecretRequest{}
	return &this
}

// GetEncryptedValue returns the EncryptedValue field value if set, zero value otherwise.
func (o *ActionsCreateOrUpdateRepoSecretRequest) GetEncryptedValue() string {
	if o == nil || o.EncryptedValue == nil {
		var ret string
		return ret
	}
	return *o.EncryptedValue
}

// GetEncryptedValueOk returns a tuple with the EncryptedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateOrUpdateRepoSecretRequest) GetEncryptedValueOk() (*string, bool) {
	if o == nil || o.EncryptedValue == nil {
		return nil, false
	}
	return o.EncryptedValue, true
}

// HasEncryptedValue returns a boolean if a field has been set.
func (o *ActionsCreateOrUpdateRepoSecretRequest) HasEncryptedValue() bool {
	if o != nil && o.EncryptedValue != nil {
		return true
	}

	return false
}

// SetEncryptedValue gets a reference to the given string and assigns it to the EncryptedValue field.
func (o *ActionsCreateOrUpdateRepoSecretRequest) SetEncryptedValue(v string) {
	o.EncryptedValue = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ActionsCreateOrUpdateRepoSecretRequest) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateOrUpdateRepoSecretRequest) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ActionsCreateOrUpdateRepoSecretRequest) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *ActionsCreateOrUpdateRepoSecretRequest) SetKeyId(v string) {
	o.KeyId = &v
}

func (o ActionsCreateOrUpdateRepoSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptedValue != nil {
		toSerialize["encrypted_value"] = o.EncryptedValue
	}
	if o.KeyId != nil {
		toSerialize["key_id"] = o.KeyId
	}
	return json.Marshal(toSerialize)
}

type NullableActionsCreateOrUpdateRepoSecretRequest struct {
	value *ActionsCreateOrUpdateRepoSecretRequest
	isSet bool
}

func (v NullableActionsCreateOrUpdateRepoSecretRequest) Get() *ActionsCreateOrUpdateRepoSecretRequest {
	return v.value
}

func (v *NullableActionsCreateOrUpdateRepoSecretRequest) Set(val *ActionsCreateOrUpdateRepoSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsCreateOrUpdateRepoSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsCreateOrUpdateRepoSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsCreateOrUpdateRepoSecretRequest(val *ActionsCreateOrUpdateRepoSecretRequest) *NullableActionsCreateOrUpdateRepoSecretRequest {
	return &NullableActionsCreateOrUpdateRepoSecretRequest{value: val, isSet: true}
}

func (v NullableActionsCreateOrUpdateRepoSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsCreateOrUpdateRepoSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


