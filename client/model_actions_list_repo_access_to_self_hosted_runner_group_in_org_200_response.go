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

// ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response struct for ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response
type ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response struct {
	TotalCount float32 `json:"total_count"`
	Repositories []MinimalRepository `json:"repositories"`
}

// NewActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response instantiates a new ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response(totalCount float32, repositories []MinimalRepository) *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response {
	this := ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response{}
	this.TotalCount = totalCount
	this.Repositories = repositories
	return &this
}

// NewActionsListRepoAccessToSelfHostedRunnerGroupInOrg200ResponseWithDefaults instantiates a new ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsListRepoAccessToSelfHostedRunnerGroupInOrg200ResponseWithDefaults() *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response {
	this := ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) SetTotalCount(v float32) {
	o.TotalCount = v
}

// GetRepositories returns the Repositories field value
func (o *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) GetRepositories() []MinimalRepository {
	if o == nil {
		var ret []MinimalRepository
		return ret
	}

	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value
// and a boolean to check if the value has been set.
func (o *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) GetRepositoriesOk() ([]MinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repositories, true
}

// SetRepositories sets field value
func (o *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) SetRepositories(v []MinimalRepository) {
	o.Repositories = v
}

func (o ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["repositories"] = o.Repositories
	}
	return json.Marshal(toSerialize)
}

type NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response struct {
	value *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response
	isSet bool
}

func (v NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) Get() *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response {
	return v.value
}

func (v *NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) Set(val *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response(val *ActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) *NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response {
	return &NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response{value: val, isSet: true}
}

func (v NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsListRepoAccessToSelfHostedRunnerGroupInOrg200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


