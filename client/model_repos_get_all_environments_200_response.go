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

// ReposGetAllEnvironments200Response struct for ReposGetAllEnvironments200Response
type ReposGetAllEnvironments200Response struct {
	// The number of environments in this repository
	TotalCount *int32 `json:"total_count,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
}

// NewReposGetAllEnvironments200Response instantiates a new ReposGetAllEnvironments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposGetAllEnvironments200Response() *ReposGetAllEnvironments200Response {
	this := ReposGetAllEnvironments200Response{}
	return &this
}

// NewReposGetAllEnvironments200ResponseWithDefaults instantiates a new ReposGetAllEnvironments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposGetAllEnvironments200ResponseWithDefaults() *ReposGetAllEnvironments200Response {
	this := ReposGetAllEnvironments200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ReposGetAllEnvironments200Response) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGetAllEnvironments200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ReposGetAllEnvironments200Response) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ReposGetAllEnvironments200Response) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ReposGetAllEnvironments200Response) GetEnvironments() []Environment {
	if o == nil || o.Environments == nil {
		var ret []Environment
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposGetAllEnvironments200Response) GetEnvironmentsOk() ([]Environment, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ReposGetAllEnvironments200Response) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *ReposGetAllEnvironments200Response) SetEnvironments(v []Environment) {
	o.Environments = v
}

func (o ReposGetAllEnvironments200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	return json.Marshal(toSerialize)
}

type NullableReposGetAllEnvironments200Response struct {
	value *ReposGetAllEnvironments200Response
	isSet bool
}

func (v NullableReposGetAllEnvironments200Response) Get() *ReposGetAllEnvironments200Response {
	return v.value
}

func (v *NullableReposGetAllEnvironments200Response) Set(val *ReposGetAllEnvironments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReposGetAllEnvironments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReposGetAllEnvironments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposGetAllEnvironments200Response(val *ReposGetAllEnvironments200Response) *NullableReposGetAllEnvironments200Response {
	return &NullableReposGetAllEnvironments200Response{value: val, isSet: true}
}

func (v NullableReposGetAllEnvironments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposGetAllEnvironments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


