/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// RepoCodespacesSecret Set repository secrets for GitHub Codespaces.
type RepoCodespacesSecret struct {
	// The name of the secret.
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewRepoCodespacesSecret instantiates a new RepoCodespacesSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepoCodespacesSecret(name string, createdAt time.Time, updatedAt time.Time) *RepoCodespacesSecret {
	this := RepoCodespacesSecret{}
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewRepoCodespacesSecretWithDefaults instantiates a new RepoCodespacesSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepoCodespacesSecretWithDefaults() *RepoCodespacesSecret {
	this := RepoCodespacesSecret{}
	return &this
}

// GetName returns the Name field value
func (o *RepoCodespacesSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RepoCodespacesSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RepoCodespacesSecret) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RepoCodespacesSecret) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RepoCodespacesSecret) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RepoCodespacesSecret) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RepoCodespacesSecret) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RepoCodespacesSecret) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RepoCodespacesSecret) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RepoCodespacesSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRepoCodespacesSecret struct {
	value *RepoCodespacesSecret
	isSet bool
}

func (v NullableRepoCodespacesSecret) Get() *RepoCodespacesSecret {
	return v.value
}

func (v *NullableRepoCodespacesSecret) Set(val *RepoCodespacesSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableRepoCodespacesSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableRepoCodespacesSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepoCodespacesSecret(val *RepoCodespacesSecret) *NullableRepoCodespacesSecret {
	return &NullableRepoCodespacesSecret{value: val, isSet: true}
}

func (v NullableRepoCodespacesSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepoCodespacesSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


