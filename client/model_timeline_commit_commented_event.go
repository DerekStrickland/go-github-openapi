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

// TimelineCommitCommentedEvent Timeline Commit Commented Event
type TimelineCommitCommentedEvent struct {
	Event *string `json:"event,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	CommitId *string `json:"commit_id,omitempty"`
	Comments []CommitComment `json:"comments,omitempty"`
}

// NewTimelineCommitCommentedEvent instantiates a new TimelineCommitCommentedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineCommitCommentedEvent() *TimelineCommitCommentedEvent {
	this := TimelineCommitCommentedEvent{}
	return &this
}

// NewTimelineCommitCommentedEventWithDefaults instantiates a new TimelineCommitCommentedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineCommitCommentedEventWithDefaults() *TimelineCommitCommentedEvent {
	this := TimelineCommitCommentedEvent{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TimelineCommitCommentedEvent) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCommitCommentedEvent) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TimelineCommitCommentedEvent) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *TimelineCommitCommentedEvent) SetEvent(v string) {
	o.Event = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *TimelineCommitCommentedEvent) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCommitCommentedEvent) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *TimelineCommitCommentedEvent) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *TimelineCommitCommentedEvent) SetNodeId(v string) {
	o.NodeId = &v
}

// GetCommitId returns the CommitId field value if set, zero value otherwise.
func (o *TimelineCommitCommentedEvent) GetCommitId() string {
	if o == nil || o.CommitId == nil {
		var ret string
		return ret
	}
	return *o.CommitId
}

// GetCommitIdOk returns a tuple with the CommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCommitCommentedEvent) GetCommitIdOk() (*string, bool) {
	if o == nil || o.CommitId == nil {
		return nil, false
	}
	return o.CommitId, true
}

// HasCommitId returns a boolean if a field has been set.
func (o *TimelineCommitCommentedEvent) HasCommitId() bool {
	if o != nil && o.CommitId != nil {
		return true
	}

	return false
}

// SetCommitId gets a reference to the given string and assigns it to the CommitId field.
func (o *TimelineCommitCommentedEvent) SetCommitId(v string) {
	o.CommitId = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *TimelineCommitCommentedEvent) GetComments() []CommitComment {
	if o == nil || o.Comments == nil {
		var ret []CommitComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCommitCommentedEvent) GetCommentsOk() ([]CommitComment, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *TimelineCommitCommentedEvent) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []CommitComment and assigns it to the Comments field.
func (o *TimelineCommitCommentedEvent) SetComments(v []CommitComment) {
	o.Comments = v
}

func (o TimelineCommitCommentedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.CommitId != nil {
		toSerialize["commit_id"] = o.CommitId
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineCommitCommentedEvent struct {
	value *TimelineCommitCommentedEvent
	isSet bool
}

func (v NullableTimelineCommitCommentedEvent) Get() *TimelineCommitCommentedEvent {
	return v.value
}

func (v *NullableTimelineCommitCommentedEvent) Set(val *TimelineCommitCommentedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineCommitCommentedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineCommitCommentedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineCommitCommentedEvent(val *TimelineCommitCommentedEvent) *NullableTimelineCommitCommentedEvent {
	return &NullableTimelineCommitCommentedEvent{value: val, isSet: true}
}

func (v NullableTimelineCommitCommentedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineCommitCommentedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


