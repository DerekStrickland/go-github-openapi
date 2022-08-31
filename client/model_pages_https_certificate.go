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

// PagesHttpsCertificate struct for PagesHttpsCertificate
type PagesHttpsCertificate struct {
	State string `json:"state"`
	Description string `json:"description"`
	// Array of the domain set and its alternate name (if it is configured)
	Domains []string `json:"domains"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// NewPagesHttpsCertificate instantiates a new PagesHttpsCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagesHttpsCertificate(state string, description string, domains []string) *PagesHttpsCertificate {
	this := PagesHttpsCertificate{}
	this.State = state
	this.Description = description
	this.Domains = domains
	return &this
}

// NewPagesHttpsCertificateWithDefaults instantiates a new PagesHttpsCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagesHttpsCertificateWithDefaults() *PagesHttpsCertificate {
	this := PagesHttpsCertificate{}
	return &this
}

// GetState returns the State field value
func (o *PagesHttpsCertificate) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PagesHttpsCertificate) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PagesHttpsCertificate) SetState(v string) {
	o.State = v
}

// GetDescription returns the Description field value
func (o *PagesHttpsCertificate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PagesHttpsCertificate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PagesHttpsCertificate) SetDescription(v string) {
	o.Description = v
}

// GetDomains returns the Domains field value
func (o *PagesHttpsCertificate) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *PagesHttpsCertificate) GetDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *PagesHttpsCertificate) SetDomains(v []string) {
	o.Domains = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PagesHttpsCertificate) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagesHttpsCertificate) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PagesHttpsCertificate) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PagesHttpsCertificate) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o PagesHttpsCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["domains"] = o.Domains
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullablePagesHttpsCertificate struct {
	value *PagesHttpsCertificate
	isSet bool
}

func (v NullablePagesHttpsCertificate) Get() *PagesHttpsCertificate {
	return v.value
}

func (v *NullablePagesHttpsCertificate) Set(val *PagesHttpsCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullablePagesHttpsCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullablePagesHttpsCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagesHttpsCertificate(val *PagesHttpsCertificate) *NullablePagesHttpsCertificate {
	return &NullablePagesHttpsCertificate{value: val, isSet: true}
}

func (v NullablePagesHttpsCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagesHttpsCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

