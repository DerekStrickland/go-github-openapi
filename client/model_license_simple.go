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

// LicenseSimple License Simple
type LicenseSimple struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url NullableString `json:"url"`
	SpdxId NullableString `json:"spdx_id"`
	NodeId string `json:"node_id"`
	HtmlUrl *string `json:"html_url,omitempty"`
}

// NewLicenseSimple instantiates a new LicenseSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSimple(key string, name string, url NullableString, spdxId NullableString, nodeId string) *LicenseSimple {
	this := LicenseSimple{}
	this.Key = key
	this.Name = name
	this.Url = url
	this.SpdxId = spdxId
	this.NodeId = nodeId
	return &this
}

// NewLicenseSimpleWithDefaults instantiates a new LicenseSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseSimpleWithDefaults() *LicenseSimple {
	this := LicenseSimple{}
	return &this
}

// GetKey returns the Key field value
func (o *LicenseSimple) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *LicenseSimple) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *LicenseSimple) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *LicenseSimple) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LicenseSimple) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LicenseSimple) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LicenseSimple) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseSimple) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *LicenseSimple) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetSpdxId returns the SpdxId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LicenseSimple) GetSpdxId() string {
	if o == nil || o.SpdxId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SpdxId.Get()
}

// GetSpdxIdOk returns a tuple with the SpdxId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseSimple) GetSpdxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpdxId.Get(), o.SpdxId.IsSet()
}

// SetSpdxId sets field value
func (o *LicenseSimple) SetSpdxId(v string) {
	o.SpdxId.Set(&v)
}

// GetNodeId returns the NodeId field value
func (o *LicenseSimple) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *LicenseSimple) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *LicenseSimple) SetNodeId(v string) {
	o.NodeId = v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *LicenseSimple) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSimple) GetHtmlUrlOk() (*string, bool) {
	if o == nil || o.HtmlUrl == nil {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *LicenseSimple) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl != nil {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *LicenseSimple) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

func (o LicenseSimple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["url"] = o.Url.Get()
	}
	if true {
		toSerialize["spdx_id"] = o.SpdxId.Get()
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if o.HtmlUrl != nil {
		toSerialize["html_url"] = o.HtmlUrl
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseSimple struct {
	value *LicenseSimple
	isSet bool
}

func (v NullableLicenseSimple) Get() *LicenseSimple {
	return v.value
}

func (v *NullableLicenseSimple) Set(val *LicenseSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSimple(val *LicenseSimple) *NullableLicenseSimple {
	return &NullableLicenseSimple{value: val, isSet: true}
}

func (v NullableLicenseSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

