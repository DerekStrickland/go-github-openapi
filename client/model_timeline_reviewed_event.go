/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// TimelineReviewedEvent Timeline Reviewed Event
type TimelineReviewedEvent struct {
	Event string `json:"event"`
	// Unique identifier of the review
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	User SimpleUser `json:"user"`
	// The text of the review.
	Body NullableString `json:"body"`
	State string `json:"state"`
	HtmlUrl string `json:"html_url"`
	PullRequestUrl string `json:"pull_request_url"`
	Links TimelineReviewedEventLinks `json:"_links"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
	// A commit SHA for the review.
	CommitId string `json:"commit_id"`
	BodyHtml *string `json:"body_html,omitempty"`
	BodyText *string `json:"body_text,omitempty"`
	AuthorAssociation AuthorAssociation `json:"author_association"`
}

// NewTimelineReviewedEvent instantiates a new TimelineReviewedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineReviewedEvent(event string, id int32, nodeId string, user SimpleUser, body NullableString, state string, htmlUrl string, pullRequestUrl string, links TimelineReviewedEventLinks, commitId string, authorAssociation AuthorAssociation) *TimelineReviewedEvent {
	this := TimelineReviewedEvent{}
	this.Event = event
	this.Id = id
	this.NodeId = nodeId
	this.User = user
	this.Body = body
	this.State = state
	this.HtmlUrl = htmlUrl
	this.PullRequestUrl = pullRequestUrl
	this.Links = links
	this.CommitId = commitId
	this.AuthorAssociation = authorAssociation
	return &this
}

// NewTimelineReviewedEventWithDefaults instantiates a new TimelineReviewedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineReviewedEventWithDefaults() *TimelineReviewedEvent {
	this := TimelineReviewedEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *TimelineReviewedEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *TimelineReviewedEvent) SetEvent(v string) {
	o.Event = v
}

// GetId returns the Id field value
func (o *TimelineReviewedEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TimelineReviewedEvent) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *TimelineReviewedEvent) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *TimelineReviewedEvent) SetNodeId(v string) {
	o.NodeId = v
}

// GetUser returns the User field value
func (o *TimelineReviewedEvent) GetUser() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetUserOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TimelineReviewedEvent) SetUser(v SimpleUser) {
	o.User = v
}

// GetBody returns the Body field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimelineReviewedEvent) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}

	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineReviewedEvent) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// SetBody sets field value
func (o *TimelineReviewedEvent) SetBody(v string) {
	o.Body.Set(&v)
}

// GetState returns the State field value
func (o *TimelineReviewedEvent) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TimelineReviewedEvent) SetState(v string) {
	o.State = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *TimelineReviewedEvent) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *TimelineReviewedEvent) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetPullRequestUrl returns the PullRequestUrl field value
func (o *TimelineReviewedEvent) GetPullRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PullRequestUrl
}

// GetPullRequestUrlOk returns a tuple with the PullRequestUrl field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetPullRequestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PullRequestUrl, true
}

// SetPullRequestUrl sets field value
func (o *TimelineReviewedEvent) SetPullRequestUrl(v string) {
	o.PullRequestUrl = v
}

// GetLinks returns the Links field value
func (o *TimelineReviewedEvent) GetLinks() TimelineReviewedEventLinks {
	if o == nil {
		var ret TimelineReviewedEventLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetLinksOk() (*TimelineReviewedEventLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TimelineReviewedEvent) SetLinks(v TimelineReviewedEventLinks) {
	o.Links = v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *TimelineReviewedEvent) GetSubmittedAt() time.Time {
	if o == nil || o.SubmittedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || o.SubmittedAt == nil {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *TimelineReviewedEvent) HasSubmittedAt() bool {
	if o != nil && o.SubmittedAt != nil {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *TimelineReviewedEvent) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetCommitId returns the CommitId field value
func (o *TimelineReviewedEvent) GetCommitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitId
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitId, true
}

// SetCommitId sets field value
func (o *TimelineReviewedEvent) SetCommitId(v string) {
	o.CommitId = v
}

// GetBodyHtml returns the BodyHtml field value if set, zero value otherwise.
func (o *TimelineReviewedEvent) GetBodyHtml() string {
	if o == nil || o.BodyHtml == nil {
		var ret string
		return ret
	}
	return *o.BodyHtml
}

// GetBodyHtmlOk returns a tuple with the BodyHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetBodyHtmlOk() (*string, bool) {
	if o == nil || o.BodyHtml == nil {
		return nil, false
	}
	return o.BodyHtml, true
}

// HasBodyHtml returns a boolean if a field has been set.
func (o *TimelineReviewedEvent) HasBodyHtml() bool {
	if o != nil && o.BodyHtml != nil {
		return true
	}

	return false
}

// SetBodyHtml gets a reference to the given string and assigns it to the BodyHtml field.
func (o *TimelineReviewedEvent) SetBodyHtml(v string) {
	o.BodyHtml = &v
}

// GetBodyText returns the BodyText field value if set, zero value otherwise.
func (o *TimelineReviewedEvent) GetBodyText() string {
	if o == nil || o.BodyText == nil {
		var ret string
		return ret
	}
	return *o.BodyText
}

// GetBodyTextOk returns a tuple with the BodyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetBodyTextOk() (*string, bool) {
	if o == nil || o.BodyText == nil {
		return nil, false
	}
	return o.BodyText, true
}

// HasBodyText returns a boolean if a field has been set.
func (o *TimelineReviewedEvent) HasBodyText() bool {
	if o != nil && o.BodyText != nil {
		return true
	}

	return false
}

// SetBodyText gets a reference to the given string and assigns it to the BodyText field.
func (o *TimelineReviewedEvent) SetBodyText(v string) {
	o.BodyText = &v
}

// GetAuthorAssociation returns the AuthorAssociation field value
func (o *TimelineReviewedEvent) GetAuthorAssociation() AuthorAssociation {
	if o == nil {
		var ret AuthorAssociation
		return ret
	}

	return o.AuthorAssociation
}

// GetAuthorAssociationOk returns a tuple with the AuthorAssociation field value
// and a boolean to check if the value has been set.
func (o *TimelineReviewedEvent) GetAuthorAssociationOk() (*AuthorAssociation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorAssociation, true
}

// SetAuthorAssociation sets field value
func (o *TimelineReviewedEvent) SetAuthorAssociation(v AuthorAssociation) {
	o.AuthorAssociation = v
}

func (o TimelineReviewedEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["body"] = o.Body.Get()
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["pull_request_url"] = o.PullRequestUrl
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if o.SubmittedAt != nil {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if true {
		toSerialize["commit_id"] = o.CommitId
	}
	if o.BodyHtml != nil {
		toSerialize["body_html"] = o.BodyHtml
	}
	if o.BodyText != nil {
		toSerialize["body_text"] = o.BodyText
	}
	if true {
		toSerialize["author_association"] = o.AuthorAssociation
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineReviewedEvent struct {
	value *TimelineReviewedEvent
	isSet bool
}

func (v NullableTimelineReviewedEvent) Get() *TimelineReviewedEvent {
	return v.value
}

func (v *NullableTimelineReviewedEvent) Set(val *TimelineReviewedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineReviewedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineReviewedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineReviewedEvent(val *TimelineReviewedEvent) *NullableTimelineReviewedEvent {
	return &NullableTimelineReviewedEvent{value: val, isSet: true}
}

func (v NullableTimelineReviewedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineReviewedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

