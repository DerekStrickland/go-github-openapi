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

// IssuesCreateRequest struct for IssuesCreateRequest
type IssuesCreateRequest struct {
	Title IssuesCreateRequestTitle `json:"title"`
	// The contents of the issue.
	Body *string `json:"body,omitempty"`
	// Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_
	Assignee NullableString `json:"assignee,omitempty"`
	Milestone NullableIssuesCreateRequestMilestone `json:"milestone,omitempty"`
	// Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._
	Labels []IssuesCreateRequestLabelsInner `json:"labels,omitempty"`
	// Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._
	Assignees []string `json:"assignees,omitempty"`
}

// NewIssuesCreateRequest instantiates a new IssuesCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesCreateRequest(title IssuesCreateRequestTitle) *IssuesCreateRequest {
	this := IssuesCreateRequest{}
	this.Title = title
	return &this
}

// NewIssuesCreateRequestWithDefaults instantiates a new IssuesCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesCreateRequestWithDefaults() *IssuesCreateRequest {
	this := IssuesCreateRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *IssuesCreateRequest) GetTitle() IssuesCreateRequestTitle {
	if o == nil {
		var ret IssuesCreateRequestTitle
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IssuesCreateRequest) GetTitleOk() (*IssuesCreateRequestTitle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *IssuesCreateRequest) SetTitle(v IssuesCreateRequestTitle) {
	o.Title = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *IssuesCreateRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesCreateRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *IssuesCreateRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *IssuesCreateRequest) SetBody(v string) {
	o.Body = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuesCreateRequest) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuesCreateRequest) GetAssigneeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// HasAssignee returns a boolean if a field has been set.
func (o *IssuesCreateRequest) HasAssignee() bool {
	if o != nil && o.Assignee.IsSet() {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given NullableString and assigns it to the Assignee field.
func (o *IssuesCreateRequest) SetAssignee(v string) {
	o.Assignee.Set(&v)
}
// SetAssigneeNil sets the value for Assignee to be an explicit nil
func (o *IssuesCreateRequest) SetAssigneeNil() {
	o.Assignee.Set(nil)
}

// UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
func (o *IssuesCreateRequest) UnsetAssignee() {
	o.Assignee.Unset()
}

// GetMilestone returns the Milestone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssuesCreateRequest) GetMilestone() IssuesCreateRequestMilestone {
	if o == nil || o.Milestone.Get() == nil {
		var ret IssuesCreateRequestMilestone
		return ret
	}
	return *o.Milestone.Get()
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssuesCreateRequest) GetMilestoneOk() (*IssuesCreateRequestMilestone, bool) {
	if o == nil {
		return nil, false
	}
	return o.Milestone.Get(), o.Milestone.IsSet()
}

// HasMilestone returns a boolean if a field has been set.
func (o *IssuesCreateRequest) HasMilestone() bool {
	if o != nil && o.Milestone.IsSet() {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given NullableIssuesCreateRequestMilestone and assigns it to the Milestone field.
func (o *IssuesCreateRequest) SetMilestone(v IssuesCreateRequestMilestone) {
	o.Milestone.Set(&v)
}
// SetMilestoneNil sets the value for Milestone to be an explicit nil
func (o *IssuesCreateRequest) SetMilestoneNil() {
	o.Milestone.Set(nil)
}

// UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
func (o *IssuesCreateRequest) UnsetMilestone() {
	o.Milestone.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *IssuesCreateRequest) GetLabels() []IssuesCreateRequestLabelsInner {
	if o == nil || o.Labels == nil {
		var ret []IssuesCreateRequestLabelsInner
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesCreateRequest) GetLabelsOk() ([]IssuesCreateRequestLabelsInner, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *IssuesCreateRequest) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []IssuesCreateRequestLabelsInner and assigns it to the Labels field.
func (o *IssuesCreateRequest) SetLabels(v []IssuesCreateRequestLabelsInner) {
	o.Labels = v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *IssuesCreateRequest) GetAssignees() []string {
	if o == nil || o.Assignees == nil {
		var ret []string
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesCreateRequest) GetAssigneesOk() ([]string, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *IssuesCreateRequest) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []string and assigns it to the Assignees field.
func (o *IssuesCreateRequest) SetAssignees(v []string) {
	o.Assignees = v
}

func (o IssuesCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Assignee.IsSet() {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if o.Milestone.IsSet() {
		toSerialize["milestone"] = o.Milestone.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Assignees != nil {
		toSerialize["assignees"] = o.Assignees
	}
	return json.Marshal(toSerialize)
}

type NullableIssuesCreateRequest struct {
	value *IssuesCreateRequest
	isSet bool
}

func (v NullableIssuesCreateRequest) Get() *IssuesCreateRequest {
	return v.value
}

func (v *NullableIssuesCreateRequest) Set(val *IssuesCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesCreateRequest(val *IssuesCreateRequest) *NullableIssuesCreateRequest {
	return &NullableIssuesCreateRequest{value: val, isSet: true}
}

func (v NullableIssuesCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


