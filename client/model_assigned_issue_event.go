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

// AssignedIssueEvent Assigned Issue Event
type AssignedIssueEvent struct {
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Actor SimpleUser `json:"actor"`
	Event string `json:"event"`
	CommitId NullableString `json:"commit_id"`
	CommitUrl NullableString `json:"commit_url"`
	CreatedAt string `json:"created_at"`
	PerformedViaGithubApp Integration `json:"performed_via_github_app"`
	Assignee SimpleUser `json:"assignee"`
	Assigner SimpleUser `json:"assigner"`
}

// NewAssignedIssueEvent instantiates a new AssignedIssueEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp Integration, assignee SimpleUser, assigner SimpleUser) *AssignedIssueEvent {
	this := AssignedIssueEvent{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.Actor = actor
	this.Event = event
	this.CommitId = commitId
	this.CommitUrl = commitUrl
	this.CreatedAt = createdAt
	this.PerformedViaGithubApp = performedViaGithubApp
	this.Assignee = assignee
	this.Assigner = assigner
	return &this
}

// NewAssignedIssueEventWithDefaults instantiates a new AssignedIssueEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedIssueEventWithDefaults() *AssignedIssueEvent {
	this := AssignedIssueEvent{}
	return &this
}

// GetId returns the Id field value
func (o *AssignedIssueEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssignedIssueEvent) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *AssignedIssueEvent) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *AssignedIssueEvent) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *AssignedIssueEvent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AssignedIssueEvent) SetUrl(v string) {
	o.Url = v
}

// GetActor returns the Actor field value
func (o *AssignedIssueEvent) GetActor() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetActorOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *AssignedIssueEvent) SetActor(v SimpleUser) {
	o.Actor = v
}

// GetEvent returns the Event field value
func (o *AssignedIssueEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AssignedIssueEvent) SetEvent(v string) {
	o.Event = v
}

// GetCommitId returns the CommitId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssignedIssueEvent) GetCommitId() string {
	if o == nil || o.CommitId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitId.Get()
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssignedIssueEvent) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitId.Get(), o.CommitId.IsSet()
}

// SetCommitId sets field value
func (o *AssignedIssueEvent) SetCommitId(v string) {
	o.CommitId.Set(&v)
}

// GetCommitUrl returns the CommitUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssignedIssueEvent) GetCommitUrl() string {
	if o == nil || o.CommitUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitUrl.Get()
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssignedIssueEvent) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitUrl.Get(), o.CommitUrl.IsSet()
}

// SetCommitUrl sets field value
func (o *AssignedIssueEvent) SetCommitUrl(v string) {
	o.CommitUrl.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *AssignedIssueEvent) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AssignedIssueEvent) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetPerformedViaGithubApp returns the PerformedViaGithubApp field value
func (o *AssignedIssueEvent) GetPerformedViaGithubApp() Integration {
	if o == nil {
		var ret Integration
		return ret
	}

	return o.PerformedViaGithubApp
}

// GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetPerformedViaGithubAppOk() (*Integration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerformedViaGithubApp, true
}

// SetPerformedViaGithubApp sets field value
func (o *AssignedIssueEvent) SetPerformedViaGithubApp(v Integration) {
	o.PerformedViaGithubApp = v
}

// GetAssignee returns the Assignee field value
func (o *AssignedIssueEvent) GetAssignee() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetAssigneeOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignee, true
}

// SetAssignee sets field value
func (o *AssignedIssueEvent) SetAssignee(v SimpleUser) {
	o.Assignee = v
}

// GetAssigner returns the Assigner field value
func (o *AssignedIssueEvent) GetAssigner() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Assigner
}

// GetAssignerOk returns a tuple with the Assigner field value
// and a boolean to check if the value has been set.
func (o *AssignedIssueEvent) GetAssignerOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assigner, true
}

// SetAssigner sets field value
func (o *AssignedIssueEvent) SetAssigner(v SimpleUser) {
	o.Assigner = v
}

func (o AssignedIssueEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["commit_id"] = o.CommitId.Get()
	}
	if true {
		toSerialize["commit_url"] = o.CommitUrl.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["performed_via_github_app"] = o.PerformedViaGithubApp
	}
	if true {
		toSerialize["assignee"] = o.Assignee
	}
	if true {
		toSerialize["assigner"] = o.Assigner
	}
	return json.Marshal(toSerialize)
}

type NullableAssignedIssueEvent struct {
	value *AssignedIssueEvent
	isSet bool
}

func (v NullableAssignedIssueEvent) Get() *AssignedIssueEvent {
	return v.value
}

func (v *NullableAssignedIssueEvent) Set(val *AssignedIssueEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedIssueEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedIssueEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedIssueEvent(val *AssignedIssueEvent) *NullableAssignedIssueEvent {
	return &NullableAssignedIssueEvent{value: val, isSet: true}
}

func (v NullableAssignedIssueEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedIssueEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

