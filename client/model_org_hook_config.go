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

// OrgHookConfig struct for OrgHookConfig
type OrgHookConfig struct {
	Url *string `json:"url,omitempty"`
	InsecureSsl *string `json:"insecure_ssl,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Secret *string `json:"secret,omitempty"`
}

// NewOrgHookConfig instantiates a new OrgHookConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgHookConfig() *OrgHookConfig {
	this := OrgHookConfig{}
	return &this
}

// NewOrgHookConfigWithDefaults instantiates a new OrgHookConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgHookConfigWithDefaults() *OrgHookConfig {
	this := OrgHookConfig{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrgHookConfig) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgHookConfig) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrgHookConfig) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrgHookConfig) SetUrl(v string) {
	o.Url = &v
}

// GetInsecureSsl returns the InsecureSsl field value if set, zero value otherwise.
func (o *OrgHookConfig) GetInsecureSsl() string {
	if o == nil || o.InsecureSsl == nil {
		var ret string
		return ret
	}
	return *o.InsecureSsl
}

// GetInsecureSslOk returns a tuple with the InsecureSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgHookConfig) GetInsecureSslOk() (*string, bool) {
	if o == nil || o.InsecureSsl == nil {
		return nil, false
	}
	return o.InsecureSsl, true
}

// HasInsecureSsl returns a boolean if a field has been set.
func (o *OrgHookConfig) HasInsecureSsl() bool {
	if o != nil && o.InsecureSsl != nil {
		return true
	}

	return false
}

// SetInsecureSsl gets a reference to the given string and assigns it to the InsecureSsl field.
func (o *OrgHookConfig) SetInsecureSsl(v string) {
	o.InsecureSsl = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *OrgHookConfig) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgHookConfig) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *OrgHookConfig) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *OrgHookConfig) SetContentType(v string) {
	o.ContentType = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *OrgHookConfig) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgHookConfig) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *OrgHookConfig) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *OrgHookConfig) SetSecret(v string) {
	o.Secret = &v
}

func (o OrgHookConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.InsecureSsl != nil {
		toSerialize["insecure_ssl"] = o.InsecureSsl
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableOrgHookConfig struct {
	value *OrgHookConfig
	isSet bool
}

func (v NullableOrgHookConfig) Get() *OrgHookConfig {
	return v.value
}

func (v *NullableOrgHookConfig) Set(val *OrgHookConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgHookConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgHookConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgHookConfig(val *OrgHookConfig) *NullableOrgHookConfig {
	return &NullableOrgHookConfig{value: val, isSet: true}
}

func (v NullableOrgHookConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgHookConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


