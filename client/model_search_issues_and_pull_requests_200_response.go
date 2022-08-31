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

// SearchIssuesAndPullRequests200Response struct for SearchIssuesAndPullRequests200Response
type SearchIssuesAndPullRequests200Response struct {
	TotalCount int32 `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items []IssueSearchResultItem `json:"items"`
}

// NewSearchIssuesAndPullRequests200Response instantiates a new SearchIssuesAndPullRequests200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIssuesAndPullRequests200Response(totalCount int32, incompleteResults bool, items []IssueSearchResultItem) *SearchIssuesAndPullRequests200Response {
	this := SearchIssuesAndPullRequests200Response{}
	this.TotalCount = totalCount
	this.IncompleteResults = incompleteResults
	this.Items = items
	return &this
}

// NewSearchIssuesAndPullRequests200ResponseWithDefaults instantiates a new SearchIssuesAndPullRequests200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIssuesAndPullRequests200ResponseWithDefaults() *SearchIssuesAndPullRequests200Response {
	this := SearchIssuesAndPullRequests200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *SearchIssuesAndPullRequests200Response) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SearchIssuesAndPullRequests200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SearchIssuesAndPullRequests200Response) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetIncompleteResults returns the IncompleteResults field value
func (o *SearchIssuesAndPullRequests200Response) GetIncompleteResults() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncompleteResults
}

// GetIncompleteResultsOk returns a tuple with the IncompleteResults field value
// and a boolean to check if the value has been set.
func (o *SearchIssuesAndPullRequests200Response) GetIncompleteResultsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncompleteResults, true
}

// SetIncompleteResults sets field value
func (o *SearchIssuesAndPullRequests200Response) SetIncompleteResults(v bool) {
	o.IncompleteResults = v
}

// GetItems returns the Items field value
func (o *SearchIssuesAndPullRequests200Response) GetItems() []IssueSearchResultItem {
	if o == nil {
		var ret []IssueSearchResultItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SearchIssuesAndPullRequests200Response) GetItemsOk() ([]IssueSearchResultItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SearchIssuesAndPullRequests200Response) SetItems(v []IssueSearchResultItem) {
	o.Items = v
}

func (o SearchIssuesAndPullRequests200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["incomplete_results"] = o.IncompleteResults
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSearchIssuesAndPullRequests200Response struct {
	value *SearchIssuesAndPullRequests200Response
	isSet bool
}

func (v NullableSearchIssuesAndPullRequests200Response) Get() *SearchIssuesAndPullRequests200Response {
	return v.value
}

func (v *NullableSearchIssuesAndPullRequests200Response) Set(val *SearchIssuesAndPullRequests200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIssuesAndPullRequests200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIssuesAndPullRequests200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIssuesAndPullRequests200Response(val *SearchIssuesAndPullRequests200Response) *NullableSearchIssuesAndPullRequests200Response {
	return &NullableSearchIssuesAndPullRequests200Response{value: val, isSet: true}
}

func (v NullableSearchIssuesAndPullRequests200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIssuesAndPullRequests200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


