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

// FileCommitContentLinks struct for FileCommitContentLinks
type FileCommitContentLinks struct {
	Self *string `json:"self,omitempty"`
	Git *string `json:"git,omitempty"`
	Html *string `json:"html,omitempty"`
}

// NewFileCommitContentLinks instantiates a new FileCommitContentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileCommitContentLinks() *FileCommitContentLinks {
	this := FileCommitContentLinks{}
	return &this
}

// NewFileCommitContentLinksWithDefaults instantiates a new FileCommitContentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileCommitContentLinksWithDefaults() *FileCommitContentLinks {
	this := FileCommitContentLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *FileCommitContentLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitContentLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *FileCommitContentLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *FileCommitContentLinks) SetSelf(v string) {
	o.Self = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *FileCommitContentLinks) GetGit() string {
	if o == nil || o.Git == nil {
		var ret string
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitContentLinks) GetGitOk() (*string, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *FileCommitContentLinks) HasGit() bool {
	if o != nil && o.Git != nil {
		return true
	}

	return false
}

// SetGit gets a reference to the given string and assigns it to the Git field.
func (o *FileCommitContentLinks) SetGit(v string) {
	o.Git = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *FileCommitContentLinks) GetHtml() string {
	if o == nil || o.Html == nil {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCommitContentLinks) GetHtmlOk() (*string, bool) {
	if o == nil || o.Html == nil {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *FileCommitContentLinks) HasHtml() bool {
	if o != nil && o.Html != nil {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *FileCommitContentLinks) SetHtml(v string) {
	o.Html = &v
}

func (o FileCommitContentLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Html != nil {
		toSerialize["html"] = o.Html
	}
	return json.Marshal(toSerialize)
}

type NullableFileCommitContentLinks struct {
	value *FileCommitContentLinks
	isSet bool
}

func (v NullableFileCommitContentLinks) Get() *FileCommitContentLinks {
	return v.value
}

func (v *NullableFileCommitContentLinks) Set(val *FileCommitContentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCommitContentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCommitContentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCommitContentLinks(val *FileCommitContentLinks) *NullableFileCommitContentLinks {
	return &NullableFileCommitContentLinks{value: val, isSet: true}
}

func (v NullableFileCommitContentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCommitContentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


