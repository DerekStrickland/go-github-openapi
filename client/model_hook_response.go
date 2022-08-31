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

// HookResponse struct for HookResponse
type HookResponse struct {
	Code NullableInt32 `json:"code"`
	Status NullableString `json:"status"`
	Message NullableString `json:"message"`
}

// NewHookResponse instantiates a new HookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHookResponse(code NullableInt32, status NullableString, message NullableString) *HookResponse {
	this := HookResponse{}
	this.Code = code
	this.Status = status
	this.Message = message
	return &this
}

// NewHookResponseWithDefaults instantiates a new HookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookResponseWithDefaults() *HookResponse {
	this := HookResponse{}
	return &this
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *HookResponse) GetCode() int32 {
	if o == nil || o.Code.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookResponse) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *HookResponse) SetCode(v int32) {
	o.Code.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HookResponse) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *HookResponse) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HookResponse) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *HookResponse) SetMessage(v string) {
	o.Message.Set(&v)
}

func (o HookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code.Get()
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["message"] = o.Message.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHookResponse struct {
	value *HookResponse
	isSet bool
}

func (v NullableHookResponse) Get() *HookResponse {
	return v.value
}

func (v *NullableHookResponse) Set(val *HookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookResponse(val *HookResponse) *NullableHookResponse {
	return &NullableHookResponse{value: val, isSet: true}
}

func (v NullableHookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

