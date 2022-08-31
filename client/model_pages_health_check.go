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

// PagesHealthCheck Pages Health Check Status
type PagesHealthCheck struct {
	Domain *PagesHealthCheckDomain `json:"domain,omitempty"`
	AltDomain NullablePagesHealthCheckAltDomain `json:"alt_domain,omitempty"`
}

// NewPagesHealthCheck instantiates a new PagesHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagesHealthCheck() *PagesHealthCheck {
	this := PagesHealthCheck{}
	return &this
}

// NewPagesHealthCheckWithDefaults instantiates a new PagesHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagesHealthCheckWithDefaults() *PagesHealthCheck {
	this := PagesHealthCheck{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PagesHealthCheck) GetDomain() PagesHealthCheckDomain {
	if o == nil || o.Domain == nil {
		var ret PagesHealthCheckDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagesHealthCheck) GetDomainOk() (*PagesHealthCheckDomain, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PagesHealthCheck) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given PagesHealthCheckDomain and assigns it to the Domain field.
func (o *PagesHealthCheck) SetDomain(v PagesHealthCheckDomain) {
	o.Domain = &v
}

// GetAltDomain returns the AltDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PagesHealthCheck) GetAltDomain() PagesHealthCheckAltDomain {
	if o == nil || o.AltDomain.Get() == nil {
		var ret PagesHealthCheckAltDomain
		return ret
	}
	return *o.AltDomain.Get()
}

// GetAltDomainOk returns a tuple with the AltDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagesHealthCheck) GetAltDomainOk() (*PagesHealthCheckAltDomain, bool) {
	if o == nil {
		return nil, false
	}
	return o.AltDomain.Get(), o.AltDomain.IsSet()
}

// HasAltDomain returns a boolean if a field has been set.
func (o *PagesHealthCheck) HasAltDomain() bool {
	if o != nil && o.AltDomain.IsSet() {
		return true
	}

	return false
}

// SetAltDomain gets a reference to the given NullablePagesHealthCheckAltDomain and assigns it to the AltDomain field.
func (o *PagesHealthCheck) SetAltDomain(v PagesHealthCheckAltDomain) {
	o.AltDomain.Set(&v)
}
// SetAltDomainNil sets the value for AltDomain to be an explicit nil
func (o *PagesHealthCheck) SetAltDomainNil() {
	o.AltDomain.Set(nil)
}

// UnsetAltDomain ensures that no value is present for AltDomain, not even an explicit nil
func (o *PagesHealthCheck) UnsetAltDomain() {
	o.AltDomain.Unset()
}

func (o PagesHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.AltDomain.IsSet() {
		toSerialize["alt_domain"] = o.AltDomain.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePagesHealthCheck struct {
	value *PagesHealthCheck
	isSet bool
}

func (v NullablePagesHealthCheck) Get() *PagesHealthCheck {
	return v.value
}

func (v *NullablePagesHealthCheck) Set(val *PagesHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullablePagesHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullablePagesHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagesHealthCheck(val *PagesHealthCheck) *NullablePagesHealthCheck {
	return &NullablePagesHealthCheck{value: val, isSet: true}
}

func (v NullablePagesHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagesHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

