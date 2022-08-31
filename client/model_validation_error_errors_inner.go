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

// ValidationErrorErrorsInner struct for ValidationErrorErrorsInner
type ValidationErrorErrorsInner struct {
	Resource *string `json:"resource,omitempty"`
	Field *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
	Code string `json:"code"`
	Index *int32 `json:"index,omitempty"`
	Value *ValidationErrorErrorsInnerValue `json:"value,omitempty"`
}

// NewValidationErrorErrorsInner instantiates a new ValidationErrorErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorErrorsInner(code string) *ValidationErrorErrorsInner {
	this := ValidationErrorErrorsInner{}
	this.Code = code
	return &this
}

// NewValidationErrorErrorsInnerWithDefaults instantiates a new ValidationErrorErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorErrorsInnerWithDefaults() *ValidationErrorErrorsInner {
	this := ValidationErrorErrorsInner{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ValidationErrorErrorsInner) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrorsInner) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ValidationErrorErrorsInner) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ValidationErrorErrorsInner) SetResource(v string) {
	o.Resource = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ValidationErrorErrorsInner) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrorsInner) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ValidationErrorErrorsInner) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ValidationErrorErrorsInner) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationErrorErrorsInner) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationErrorErrorsInner) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationErrorErrorsInner) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value
func (o *ValidationErrorErrorsInner) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ValidationErrorErrorsInner) SetCode(v string) {
	o.Code = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ValidationErrorErrorsInner) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrorsInner) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ValidationErrorErrorsInner) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *ValidationErrorErrorsInner) SetIndex(v int32) {
	o.Index = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ValidationErrorErrorsInner) GetValue() ValidationErrorErrorsInnerValue {
	if o == nil || o.Value == nil {
		var ret ValidationErrorErrorsInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrorsInner) GetValueOk() (*ValidationErrorErrorsInnerValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ValidationErrorErrorsInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given ValidationErrorErrorsInnerValue and assigns it to the Value field.
func (o *ValidationErrorErrorsInner) SetValue(v ValidationErrorErrorsInnerValue) {
	o.Value = &v
}

func (o ValidationErrorErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableValidationErrorErrorsInner struct {
	value *ValidationErrorErrorsInner
	isSet bool
}

func (v NullableValidationErrorErrorsInner) Get() *ValidationErrorErrorsInner {
	return v.value
}

func (v *NullableValidationErrorErrorsInner) Set(val *ValidationErrorErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorErrorsInner(val *ValidationErrorErrorsInner) *NullableValidationErrorErrorsInner {
	return &NullableValidationErrorErrorsInner{value: val, isSet: true}
}

func (v NullableValidationErrorErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


