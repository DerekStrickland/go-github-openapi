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

// TimelineCommittedEvent Timeline Committed Event
type TimelineCommittedEvent struct {
	Event *string `json:"event,omitempty"`
	// SHA for the commit
	Sha string `json:"sha"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Author GitCommitAuthor `json:"author"`
	Committer GitCommitAuthor `json:"committer"`
	// Message describing the purpose of the commit
	Message string `json:"message"`
	Tree GitCommitTree `json:"tree"`
	Parents []GitCommitParentsInner `json:"parents"`
	Verification GitCommitVerification `json:"verification"`
	HtmlUrl string `json:"html_url"`
}

// NewTimelineCommittedEvent instantiates a new TimelineCommittedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineCommittedEvent(sha string, nodeId string, url string, author GitCommitAuthor, committer GitCommitAuthor, message string, tree GitCommitTree, parents []GitCommitParentsInner, verification GitCommitVerification, htmlUrl string) *TimelineCommittedEvent {
	this := TimelineCommittedEvent{}
	this.Sha = sha
	this.NodeId = nodeId
	this.Url = url
	this.Author = author
	this.Committer = committer
	this.Message = message
	this.Tree = tree
	this.Parents = parents
	this.Verification = verification
	this.HtmlUrl = htmlUrl
	return &this
}

// NewTimelineCommittedEventWithDefaults instantiates a new TimelineCommittedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineCommittedEventWithDefaults() *TimelineCommittedEvent {
	this := TimelineCommittedEvent{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TimelineCommittedEvent) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TimelineCommittedEvent) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *TimelineCommittedEvent) SetEvent(v string) {
	o.Event = &v
}

// GetSha returns the Sha field value
func (o *TimelineCommittedEvent) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *TimelineCommittedEvent) SetSha(v string) {
	o.Sha = v
}

// GetNodeId returns the NodeId field value
func (o *TimelineCommittedEvent) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *TimelineCommittedEvent) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *TimelineCommittedEvent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TimelineCommittedEvent) SetUrl(v string) {
	o.Url = v
}

// GetAuthor returns the Author field value
func (o *TimelineCommittedEvent) GetAuthor() GitCommitAuthor {
	if o == nil {
		var ret GitCommitAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetAuthorOk() (*GitCommitAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *TimelineCommittedEvent) SetAuthor(v GitCommitAuthor) {
	o.Author = v
}

// GetCommitter returns the Committer field value
func (o *TimelineCommittedEvent) GetCommitter() GitCommitAuthor {
	if o == nil {
		var ret GitCommitAuthor
		return ret
	}

	return o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetCommitterOk() (*GitCommitAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Committer, true
}

// SetCommitter sets field value
func (o *TimelineCommittedEvent) SetCommitter(v GitCommitAuthor) {
	o.Committer = v
}

// GetMessage returns the Message field value
func (o *TimelineCommittedEvent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TimelineCommittedEvent) SetMessage(v string) {
	o.Message = v
}

// GetTree returns the Tree field value
func (o *TimelineCommittedEvent) GetTree() GitCommitTree {
	if o == nil {
		var ret GitCommitTree
		return ret
	}

	return o.Tree
}

// GetTreeOk returns a tuple with the Tree field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetTreeOk() (*GitCommitTree, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tree, true
}

// SetTree sets field value
func (o *TimelineCommittedEvent) SetTree(v GitCommitTree) {
	o.Tree = v
}

// GetParents returns the Parents field value
func (o *TimelineCommittedEvent) GetParents() []GitCommitParentsInner {
	if o == nil {
		var ret []GitCommitParentsInner
		return ret
	}

	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetParentsOk() ([]GitCommitParentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parents, true
}

// SetParents sets field value
func (o *TimelineCommittedEvent) SetParents(v []GitCommitParentsInner) {
	o.Parents = v
}

// GetVerification returns the Verification field value
func (o *TimelineCommittedEvent) GetVerification() GitCommitVerification {
	if o == nil {
		var ret GitCommitVerification
		return ret
	}

	return o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetVerificationOk() (*GitCommitVerification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verification, true
}

// SetVerification sets field value
func (o *TimelineCommittedEvent) SetVerification(v GitCommitVerification) {
	o.Verification = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *TimelineCommittedEvent) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *TimelineCommittedEvent) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *TimelineCommittedEvent) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

func (o TimelineCommittedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["author"] = o.Author
	}
	if true {
		toSerialize["committer"] = o.Committer
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["tree"] = o.Tree
	}
	if true {
		toSerialize["parents"] = o.Parents
	}
	if true {
		toSerialize["verification"] = o.Verification
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineCommittedEvent struct {
	value *TimelineCommittedEvent
	isSet bool
}

func (v NullableTimelineCommittedEvent) Get() *TimelineCommittedEvent {
	return v.value
}

func (v *NullableTimelineCommittedEvent) Set(val *TimelineCommittedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineCommittedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineCommittedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineCommittedEvent(val *TimelineCommittedEvent) *NullableTimelineCommittedEvent {
	return &NullableTimelineCommittedEvent{value: val, isSet: true}
}

func (v NullableTimelineCommittedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineCommittedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


