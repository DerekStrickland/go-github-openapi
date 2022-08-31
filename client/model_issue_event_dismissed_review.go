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

// IssueEventDismissedReview struct for IssueEventDismissedReview
type IssueEventDismissedReview struct {
	State string `json:"state"`
	ReviewId int32 `json:"review_id"`
	DismissalMessage NullableString `json:"dismissal_message"`
	DismissalCommitId NullableString `json:"dismissal_commit_id,omitempty"`
}

// NewIssueEventDismissedReview instantiates a new IssueEventDismissedReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueEventDismissedReview(state string, reviewId int32, dismissalMessage NullableString) *IssueEventDismissedReview {
	this := IssueEventDismissedReview{}
	this.State = state
	this.ReviewId = reviewId
	this.DismissalMessage = dismissalMessage
	return &this
}

// NewIssueEventDismissedReviewWithDefaults instantiates a new IssueEventDismissedReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueEventDismissedReviewWithDefaults() *IssueEventDismissedReview {
	this := IssueEventDismissedReview{}
	return &this
}

// GetState returns the State field value
func (o *IssueEventDismissedReview) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *IssueEventDismissedReview) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *IssueEventDismissedReview) SetState(v string) {
	o.State = v
}

// GetReviewId returns the ReviewId field value
func (o *IssueEventDismissedReview) GetReviewId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReviewId
}

// GetReviewIdOk returns a tuple with the ReviewId field value
// and a boolean to check if the value has been set.
func (o *IssueEventDismissedReview) GetReviewIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewId, true
}

// SetReviewId sets field value
func (o *IssueEventDismissedReview) SetReviewId(v int32) {
	o.ReviewId = v
}

// GetDismissalMessage returns the DismissalMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IssueEventDismissedReview) GetDismissalMessage() string {
	if o == nil || o.DismissalMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.DismissalMessage.Get()
}

// GetDismissalMessageOk returns a tuple with the DismissalMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEventDismissedReview) GetDismissalMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DismissalMessage.Get(), o.DismissalMessage.IsSet()
}

// SetDismissalMessage sets field value
func (o *IssueEventDismissedReview) SetDismissalMessage(v string) {
	o.DismissalMessage.Set(&v)
}

// GetDismissalCommitId returns the DismissalCommitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEventDismissedReview) GetDismissalCommitId() string {
	if o == nil || o.DismissalCommitId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DismissalCommitId.Get()
}

// GetDismissalCommitIdOk returns a tuple with the DismissalCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEventDismissedReview) GetDismissalCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DismissalCommitId.Get(), o.DismissalCommitId.IsSet()
}

// HasDismissalCommitId returns a boolean if a field has been set.
func (o *IssueEventDismissedReview) HasDismissalCommitId() bool {
	if o != nil && o.DismissalCommitId.IsSet() {
		return true
	}

	return false
}

// SetDismissalCommitId gets a reference to the given NullableString and assigns it to the DismissalCommitId field.
func (o *IssueEventDismissedReview) SetDismissalCommitId(v string) {
	o.DismissalCommitId.Set(&v)
}
// SetDismissalCommitIdNil sets the value for DismissalCommitId to be an explicit nil
func (o *IssueEventDismissedReview) SetDismissalCommitIdNil() {
	o.DismissalCommitId.Set(nil)
}

// UnsetDismissalCommitId ensures that no value is present for DismissalCommitId, not even an explicit nil
func (o *IssueEventDismissedReview) UnsetDismissalCommitId() {
	o.DismissalCommitId.Unset()
}

func (o IssueEventDismissedReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["review_id"] = o.ReviewId
	}
	if true {
		toSerialize["dismissal_message"] = o.DismissalMessage.Get()
	}
	if o.DismissalCommitId.IsSet() {
		toSerialize["dismissal_commit_id"] = o.DismissalCommitId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIssueEventDismissedReview struct {
	value *IssueEventDismissedReview
	isSet bool
}

func (v NullableIssueEventDismissedReview) Get() *IssueEventDismissedReview {
	return v.value
}

func (v *NullableIssueEventDismissedReview) Set(val *IssueEventDismissedReview) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueEventDismissedReview) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueEventDismissedReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueEventDismissedReview(val *IssueEventDismissedReview) *NullableIssueEventDismissedReview {
	return &NullableIssueEventDismissedReview{value: val, isSet: true}
}

func (v NullableIssueEventDismissedReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueEventDismissedReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

