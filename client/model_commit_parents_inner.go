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

// CommitParentsInner struct for CommitParentsInner
type CommitParentsInner struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
	HtmlUrl *string `json:"html_url,omitempty"`
}

// NewCommitParentsInner instantiates a new CommitParentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitParentsInner(sha string, url string) *CommitParentsInner {
	this := CommitParentsInner{}
	this.Sha = sha
	this.Url = url
	return &this
}

// NewCommitParentsInnerWithDefaults instantiates a new CommitParentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitParentsInnerWithDefaults() *CommitParentsInner {
	this := CommitParentsInner{}
	return &this
}

// GetSha returns the Sha field value
func (o *CommitParentsInner) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *CommitParentsInner) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *CommitParentsInner) SetSha(v string) {
	o.Sha = v
}

// GetUrl returns the Url field value
func (o *CommitParentsInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CommitParentsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CommitParentsInner) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *CommitParentsInner) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitParentsInner) GetHtmlUrlOk() (*string, bool) {
	if o == nil || o.HtmlUrl == nil {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *CommitParentsInner) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl != nil {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *CommitParentsInner) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

func (o CommitParentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.HtmlUrl != nil {
		toSerialize["html_url"] = o.HtmlUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCommitParentsInner struct {
	value *CommitParentsInner
	isSet bool
}

func (v NullableCommitParentsInner) Get() *CommitParentsInner {
	return v.value
}

func (v *NullableCommitParentsInner) Set(val *CommitParentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitParentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitParentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitParentsInner(val *CommitParentsInner) *NullableCommitParentsInner {
	return &NullableCommitParentsInner{value: val, isSet: true}
}

func (v NullableCommitParentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitParentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


