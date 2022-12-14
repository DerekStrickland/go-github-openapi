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

// PullRequestReviewCommentLinksPullRequest struct for PullRequestReviewCommentLinksPullRequest
type PullRequestReviewCommentLinksPullRequest struct {
	Href string `json:"href"`
}

// NewPullRequestReviewCommentLinksPullRequest instantiates a new PullRequestReviewCommentLinksPullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestReviewCommentLinksPullRequest(href string) *PullRequestReviewCommentLinksPullRequest {
	this := PullRequestReviewCommentLinksPullRequest{}
	this.Href = href
	return &this
}

// NewPullRequestReviewCommentLinksPullRequestWithDefaults instantiates a new PullRequestReviewCommentLinksPullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestReviewCommentLinksPullRequestWithDefaults() *PullRequestReviewCommentLinksPullRequest {
	this := PullRequestReviewCommentLinksPullRequest{}
	return &this
}

// GetHref returns the Href field value
func (o *PullRequestReviewCommentLinksPullRequest) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *PullRequestReviewCommentLinksPullRequest) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *PullRequestReviewCommentLinksPullRequest) SetHref(v string) {
	o.Href = v
}

func (o PullRequestReviewCommentLinksPullRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestReviewCommentLinksPullRequest struct {
	value *PullRequestReviewCommentLinksPullRequest
	isSet bool
}

func (v NullablePullRequestReviewCommentLinksPullRequest) Get() *PullRequestReviewCommentLinksPullRequest {
	return v.value
}

func (v *NullablePullRequestReviewCommentLinksPullRequest) Set(val *PullRequestReviewCommentLinksPullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestReviewCommentLinksPullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestReviewCommentLinksPullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestReviewCommentLinksPullRequest(val *PullRequestReviewCommentLinksPullRequest) *NullablePullRequestReviewCommentLinksPullRequest {
	return &NullablePullRequestReviewCommentLinksPullRequest{value: val, isSet: true}
}

func (v NullablePullRequestReviewCommentLinksPullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestReviewCommentLinksPullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


