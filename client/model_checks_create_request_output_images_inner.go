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

// ChecksCreateRequestOutputImagesInner struct for ChecksCreateRequestOutputImagesInner
type ChecksCreateRequestOutputImagesInner struct {
	// The alternative text for the image.
	Alt string `json:"alt"`
	// The full URL of the image.
	ImageUrl string `json:"image_url"`
	// A short image description.
	Caption *string `json:"caption,omitempty"`
}

// NewChecksCreateRequestOutputImagesInner instantiates a new ChecksCreateRequestOutputImagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecksCreateRequestOutputImagesInner(alt string, imageUrl string) *ChecksCreateRequestOutputImagesInner {
	this := ChecksCreateRequestOutputImagesInner{}
	this.Alt = alt
	this.ImageUrl = imageUrl
	return &this
}

// NewChecksCreateRequestOutputImagesInnerWithDefaults instantiates a new ChecksCreateRequestOutputImagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecksCreateRequestOutputImagesInnerWithDefaults() *ChecksCreateRequestOutputImagesInner {
	this := ChecksCreateRequestOutputImagesInner{}
	return &this
}

// GetAlt returns the Alt field value
func (o *ChecksCreateRequestOutputImagesInner) GetAlt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alt
}

// GetAltOk returns a tuple with the Alt field value
// and a boolean to check if the value has been set.
func (o *ChecksCreateRequestOutputImagesInner) GetAltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alt, true
}

// SetAlt sets field value
func (o *ChecksCreateRequestOutputImagesInner) SetAlt(v string) {
	o.Alt = v
}

// GetImageUrl returns the ImageUrl field value
func (o *ChecksCreateRequestOutputImagesInner) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *ChecksCreateRequestOutputImagesInner) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *ChecksCreateRequestOutputImagesInner) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *ChecksCreateRequestOutputImagesInner) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecksCreateRequestOutputImagesInner) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *ChecksCreateRequestOutputImagesInner) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *ChecksCreateRequestOutputImagesInner) SetCaption(v string) {
	o.Caption = &v
}

func (o ChecksCreateRequestOutputImagesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alt"] = o.Alt
	}
	if true {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Caption != nil {
		toSerialize["caption"] = o.Caption
	}
	return json.Marshal(toSerialize)
}

type NullableChecksCreateRequestOutputImagesInner struct {
	value *ChecksCreateRequestOutputImagesInner
	isSet bool
}

func (v NullableChecksCreateRequestOutputImagesInner) Get() *ChecksCreateRequestOutputImagesInner {
	return v.value
}

func (v *NullableChecksCreateRequestOutputImagesInner) Set(val *ChecksCreateRequestOutputImagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChecksCreateRequestOutputImagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChecksCreateRequestOutputImagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecksCreateRequestOutputImagesInner(val *ChecksCreateRequestOutputImagesInner) *NullableChecksCreateRequestOutputImagesInner {
	return &NullableChecksCreateRequestOutputImagesInner{value: val, isSet: true}
}

func (v NullableChecksCreateRequestOutputImagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecksCreateRequestOutputImagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


