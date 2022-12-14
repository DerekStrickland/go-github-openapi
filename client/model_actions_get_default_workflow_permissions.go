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

// ActionsGetDefaultWorkflowPermissions struct for ActionsGetDefaultWorkflowPermissions
type ActionsGetDefaultWorkflowPermissions struct {
	DefaultWorkflowPermissions ActionsDefaultWorkflowPermissions `json:"default_workflow_permissions"`
	// Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	CanApprovePullRequestReviews bool `json:"can_approve_pull_request_reviews"`
}

// NewActionsGetDefaultWorkflowPermissions instantiates a new ActionsGetDefaultWorkflowPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsGetDefaultWorkflowPermissions(defaultWorkflowPermissions ActionsDefaultWorkflowPermissions, canApprovePullRequestReviews bool) *ActionsGetDefaultWorkflowPermissions {
	this := ActionsGetDefaultWorkflowPermissions{}
	this.DefaultWorkflowPermissions = defaultWorkflowPermissions
	this.CanApprovePullRequestReviews = canApprovePullRequestReviews
	return &this
}

// NewActionsGetDefaultWorkflowPermissionsWithDefaults instantiates a new ActionsGetDefaultWorkflowPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsGetDefaultWorkflowPermissionsWithDefaults() *ActionsGetDefaultWorkflowPermissions {
	this := ActionsGetDefaultWorkflowPermissions{}
	return &this
}

// GetDefaultWorkflowPermissions returns the DefaultWorkflowPermissions field value
func (o *ActionsGetDefaultWorkflowPermissions) GetDefaultWorkflowPermissions() ActionsDefaultWorkflowPermissions {
	if o == nil {
		var ret ActionsDefaultWorkflowPermissions
		return ret
	}

	return o.DefaultWorkflowPermissions
}

// GetDefaultWorkflowPermissionsOk returns a tuple with the DefaultWorkflowPermissions field value
// and a boolean to check if the value has been set.
func (o *ActionsGetDefaultWorkflowPermissions) GetDefaultWorkflowPermissionsOk() (*ActionsDefaultWorkflowPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultWorkflowPermissions, true
}

// SetDefaultWorkflowPermissions sets field value
func (o *ActionsGetDefaultWorkflowPermissions) SetDefaultWorkflowPermissions(v ActionsDefaultWorkflowPermissions) {
	o.DefaultWorkflowPermissions = v
}

// GetCanApprovePullRequestReviews returns the CanApprovePullRequestReviews field value
func (o *ActionsGetDefaultWorkflowPermissions) GetCanApprovePullRequestReviews() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanApprovePullRequestReviews
}

// GetCanApprovePullRequestReviewsOk returns a tuple with the CanApprovePullRequestReviews field value
// and a boolean to check if the value has been set.
func (o *ActionsGetDefaultWorkflowPermissions) GetCanApprovePullRequestReviewsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanApprovePullRequestReviews, true
}

// SetCanApprovePullRequestReviews sets field value
func (o *ActionsGetDefaultWorkflowPermissions) SetCanApprovePullRequestReviews(v bool) {
	o.CanApprovePullRequestReviews = v
}

func (o ActionsGetDefaultWorkflowPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["default_workflow_permissions"] = o.DefaultWorkflowPermissions
	}
	if true {
		toSerialize["can_approve_pull_request_reviews"] = o.CanApprovePullRequestReviews
	}
	return json.Marshal(toSerialize)
}

type NullableActionsGetDefaultWorkflowPermissions struct {
	value *ActionsGetDefaultWorkflowPermissions
	isSet bool
}

func (v NullableActionsGetDefaultWorkflowPermissions) Get() *ActionsGetDefaultWorkflowPermissions {
	return v.value
}

func (v *NullableActionsGetDefaultWorkflowPermissions) Set(val *ActionsGetDefaultWorkflowPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsGetDefaultWorkflowPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsGetDefaultWorkflowPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsGetDefaultWorkflowPermissions(val *ActionsGetDefaultWorkflowPermissions) *NullableActionsGetDefaultWorkflowPermissions {
	return &NullableActionsGetDefaultWorkflowPermissions{value: val, isSet: true}
}

func (v NullableActionsGetDefaultWorkflowPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsGetDefaultWorkflowPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


