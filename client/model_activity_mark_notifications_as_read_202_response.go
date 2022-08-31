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

// ActivityMarkNotificationsAsRead202Response struct for ActivityMarkNotificationsAsRead202Response
type ActivityMarkNotificationsAsRead202Response struct {
	Message *string `json:"message,omitempty"`
}

// NewActivityMarkNotificationsAsRead202Response instantiates a new ActivityMarkNotificationsAsRead202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityMarkNotificationsAsRead202Response() *ActivityMarkNotificationsAsRead202Response {
	this := ActivityMarkNotificationsAsRead202Response{}
	return &this
}

// NewActivityMarkNotificationsAsRead202ResponseWithDefaults instantiates a new ActivityMarkNotificationsAsRead202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityMarkNotificationsAsRead202ResponseWithDefaults() *ActivityMarkNotificationsAsRead202Response {
	this := ActivityMarkNotificationsAsRead202Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ActivityMarkNotificationsAsRead202Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityMarkNotificationsAsRead202Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ActivityMarkNotificationsAsRead202Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ActivityMarkNotificationsAsRead202Response) SetMessage(v string) {
	o.Message = &v
}

func (o ActivityMarkNotificationsAsRead202Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableActivityMarkNotificationsAsRead202Response struct {
	value *ActivityMarkNotificationsAsRead202Response
	isSet bool
}

func (v NullableActivityMarkNotificationsAsRead202Response) Get() *ActivityMarkNotificationsAsRead202Response {
	return v.value
}

func (v *NullableActivityMarkNotificationsAsRead202Response) Set(val *ActivityMarkNotificationsAsRead202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityMarkNotificationsAsRead202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityMarkNotificationsAsRead202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityMarkNotificationsAsRead202Response(val *ActivityMarkNotificationsAsRead202Response) *NullableActivityMarkNotificationsAsRead202Response {
	return &NullableActivityMarkNotificationsAsRead202Response{value: val, isSet: true}
}

func (v NullableActivityMarkNotificationsAsRead202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityMarkNotificationsAsRead202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

