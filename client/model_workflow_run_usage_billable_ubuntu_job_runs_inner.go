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

// WorkflowRunUsageBillableUBUNTUJobRunsInner struct for WorkflowRunUsageBillableUBUNTUJobRunsInner
type WorkflowRunUsageBillableUBUNTUJobRunsInner struct {
	JobId int32 `json:"job_id"`
	DurationMs int32 `json:"duration_ms"`
}

// NewWorkflowRunUsageBillableUBUNTUJobRunsInner instantiates a new WorkflowRunUsageBillableUBUNTUJobRunsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunUsageBillableUBUNTUJobRunsInner(jobId int32, durationMs int32) *WorkflowRunUsageBillableUBUNTUJobRunsInner {
	this := WorkflowRunUsageBillableUBUNTUJobRunsInner{}
	this.JobId = jobId
	this.DurationMs = durationMs
	return &this
}

// NewWorkflowRunUsageBillableUBUNTUJobRunsInnerWithDefaults instantiates a new WorkflowRunUsageBillableUBUNTUJobRunsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunUsageBillableUBUNTUJobRunsInnerWithDefaults() *WorkflowRunUsageBillableUBUNTUJobRunsInner {
	this := WorkflowRunUsageBillableUBUNTUJobRunsInner{}
	return &this
}

// GetJobId returns the JobId field value
func (o *WorkflowRunUsageBillableUBUNTUJobRunsInner) GetJobId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunUsageBillableUBUNTUJobRunsInner) GetJobIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *WorkflowRunUsageBillableUBUNTUJobRunsInner) SetJobId(v int32) {
	o.JobId = v
}

// GetDurationMs returns the DurationMs field value
func (o *WorkflowRunUsageBillableUBUNTUJobRunsInner) GetDurationMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunUsageBillableUBUNTUJobRunsInner) GetDurationMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMs, true
}

// SetDurationMs sets field value
func (o *WorkflowRunUsageBillableUBUNTUJobRunsInner) SetDurationMs(v int32) {
	o.DurationMs = v
}

func (o WorkflowRunUsageBillableUBUNTUJobRunsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["job_id"] = o.JobId
	}
	if true {
		toSerialize["duration_ms"] = o.DurationMs
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunUsageBillableUBUNTUJobRunsInner struct {
	value *WorkflowRunUsageBillableUBUNTUJobRunsInner
	isSet bool
}

func (v NullableWorkflowRunUsageBillableUBUNTUJobRunsInner) Get() *WorkflowRunUsageBillableUBUNTUJobRunsInner {
	return v.value
}

func (v *NullableWorkflowRunUsageBillableUBUNTUJobRunsInner) Set(val *WorkflowRunUsageBillableUBUNTUJobRunsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunUsageBillableUBUNTUJobRunsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunUsageBillableUBUNTUJobRunsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunUsageBillableUBUNTUJobRunsInner(val *WorkflowRunUsageBillableUBUNTUJobRunsInner) *NullableWorkflowRunUsageBillableUBUNTUJobRunsInner {
	return &NullableWorkflowRunUsageBillableUBUNTUJobRunsInner{value: val, isSet: true}
}

func (v NullableWorkflowRunUsageBillableUBUNTUJobRunsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunUsageBillableUBUNTUJobRunsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

