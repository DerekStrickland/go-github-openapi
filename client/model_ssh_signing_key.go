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

// SshSigningKey A public SSH key used to sign Git commits
type SshSigningKey struct {
	Key string `json:"key"`
	Id int32 `json:"id"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// NewSshSigningKey instantiates a new SshSigningKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshSigningKey(key string, id int32, title string, createdAt time.Time) *SshSigningKey {
	this := SshSigningKey{}
	this.Key = key
	this.Id = id
	this.Title = title
	this.CreatedAt = createdAt
	return &this
}

// NewSshSigningKeyWithDefaults instantiates a new SshSigningKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshSigningKeyWithDefaults() *SshSigningKey {
	this := SshSigningKey{}
	return &this
}

// GetKey returns the Key field value
func (o *SshSigningKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SshSigningKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SshSigningKey) SetKey(v string) {
	o.Key = v
}

// GetId returns the Id field value
func (o *SshSigningKey) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SshSigningKey) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SshSigningKey) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *SshSigningKey) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SshSigningKey) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SshSigningKey) SetTitle(v string) {
	o.Title = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SshSigningKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SshSigningKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SshSigningKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o SshSigningKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSshSigningKey struct {
	value *SshSigningKey
	isSet bool
}

func (v NullableSshSigningKey) Get() *SshSigningKey {
	return v.value
}

func (v *NullableSshSigningKey) Set(val *SshSigningKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSshSigningKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSshSigningKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshSigningKey(val *SshSigningKey) *NullableSshSigningKey {
	return &NullableSshSigningKey{value: val, isSet: true}
}

func (v NullableSshSigningKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshSigningKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


