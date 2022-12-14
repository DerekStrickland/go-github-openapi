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

// CodeScanningUpdateAlertRequest struct for CodeScanningUpdateAlertRequest
type CodeScanningUpdateAlertRequest struct {
	State CodeScanningAlertSetState `json:"state"`
	DismissedReason NullableCodeScanningAlertDismissedReason `json:"dismissed_reason,omitempty"`
	// The dismissal comment associated with the dismissal of the alert.
	DismissedComment NullableString `json:"dismissed_comment,omitempty"`
}

// NewCodeScanningUpdateAlertRequest instantiates a new CodeScanningUpdateAlertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeScanningUpdateAlertRequest(state CodeScanningAlertSetState) *CodeScanningUpdateAlertRequest {
	this := CodeScanningUpdateAlertRequest{}
	this.State = state
	return &this
}

// NewCodeScanningUpdateAlertRequestWithDefaults instantiates a new CodeScanningUpdateAlertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeScanningUpdateAlertRequestWithDefaults() *CodeScanningUpdateAlertRequest {
	this := CodeScanningUpdateAlertRequest{}
	return &this
}

// GetState returns the State field value
func (o *CodeScanningUpdateAlertRequest) GetState() CodeScanningAlertSetState {
	if o == nil {
		var ret CodeScanningAlertSetState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CodeScanningUpdateAlertRequest) GetStateOk() (*CodeScanningAlertSetState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CodeScanningUpdateAlertRequest) SetState(v CodeScanningAlertSetState) {
	o.State = v
}

// GetDismissedReason returns the DismissedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeScanningUpdateAlertRequest) GetDismissedReason() CodeScanningAlertDismissedReason {
	if o == nil || o.DismissedReason.Get() == nil {
		var ret CodeScanningAlertDismissedReason
		return ret
	}
	return *o.DismissedReason.Get()
}

// GetDismissedReasonOk returns a tuple with the DismissedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeScanningUpdateAlertRequest) GetDismissedReasonOk() (*CodeScanningAlertDismissedReason, bool) {
	if o == nil {
		return nil, false
	}
	return o.DismissedReason.Get(), o.DismissedReason.IsSet()
}

// HasDismissedReason returns a boolean if a field has been set.
func (o *CodeScanningUpdateAlertRequest) HasDismissedReason() bool {
	if o != nil && o.DismissedReason.IsSet() {
		return true
	}

	return false
}

// SetDismissedReason gets a reference to the given NullableCodeScanningAlertDismissedReason and assigns it to the DismissedReason field.
func (o *CodeScanningUpdateAlertRequest) SetDismissedReason(v CodeScanningAlertDismissedReason) {
	o.DismissedReason.Set(&v)
}
// SetDismissedReasonNil sets the value for DismissedReason to be an explicit nil
func (o *CodeScanningUpdateAlertRequest) SetDismissedReasonNil() {
	o.DismissedReason.Set(nil)
}

// UnsetDismissedReason ensures that no value is present for DismissedReason, not even an explicit nil
func (o *CodeScanningUpdateAlertRequest) UnsetDismissedReason() {
	o.DismissedReason.Unset()
}

// GetDismissedComment returns the DismissedComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeScanningUpdateAlertRequest) GetDismissedComment() string {
	if o == nil || o.DismissedComment.Get() == nil {
		var ret string
		return ret
	}
	return *o.DismissedComment.Get()
}

// GetDismissedCommentOk returns a tuple with the DismissedComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeScanningUpdateAlertRequest) GetDismissedCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DismissedComment.Get(), o.DismissedComment.IsSet()
}

// HasDismissedComment returns a boolean if a field has been set.
func (o *CodeScanningUpdateAlertRequest) HasDismissedComment() bool {
	if o != nil && o.DismissedComment.IsSet() {
		return true
	}

	return false
}

// SetDismissedComment gets a reference to the given NullableString and assigns it to the DismissedComment field.
func (o *CodeScanningUpdateAlertRequest) SetDismissedComment(v string) {
	o.DismissedComment.Set(&v)
}
// SetDismissedCommentNil sets the value for DismissedComment to be an explicit nil
func (o *CodeScanningUpdateAlertRequest) SetDismissedCommentNil() {
	o.DismissedComment.Set(nil)
}

// UnsetDismissedComment ensures that no value is present for DismissedComment, not even an explicit nil
func (o *CodeScanningUpdateAlertRequest) UnsetDismissedComment() {
	o.DismissedComment.Unset()
}

func (o CodeScanningUpdateAlertRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if o.DismissedReason.IsSet() {
		toSerialize["dismissed_reason"] = o.DismissedReason.Get()
	}
	if o.DismissedComment.IsSet() {
		toSerialize["dismissed_comment"] = o.DismissedComment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCodeScanningUpdateAlertRequest struct {
	value *CodeScanningUpdateAlertRequest
	isSet bool
}

func (v NullableCodeScanningUpdateAlertRequest) Get() *CodeScanningUpdateAlertRequest {
	return v.value
}

func (v *NullableCodeScanningUpdateAlertRequest) Set(val *CodeScanningUpdateAlertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeScanningUpdateAlertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeScanningUpdateAlertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeScanningUpdateAlertRequest(val *CodeScanningUpdateAlertRequest) *NullableCodeScanningUpdateAlertRequest {
	return &NullableCodeScanningUpdateAlertRequest{value: val, isSet: true}
}

func (v NullableCodeScanningUpdateAlertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeScanningUpdateAlertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


