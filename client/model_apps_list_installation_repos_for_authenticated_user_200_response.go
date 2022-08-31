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

// AppsListInstallationReposForAuthenticatedUser200Response struct for AppsListInstallationReposForAuthenticatedUser200Response
type AppsListInstallationReposForAuthenticatedUser200Response struct {
	TotalCount int32 `json:"total_count"`
	RepositorySelection *string `json:"repository_selection,omitempty"`
	Repositories []Repository `json:"repositories"`
}

// NewAppsListInstallationReposForAuthenticatedUser200Response instantiates a new AppsListInstallationReposForAuthenticatedUser200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsListInstallationReposForAuthenticatedUser200Response(totalCount int32, repositories []Repository) *AppsListInstallationReposForAuthenticatedUser200Response {
	this := AppsListInstallationReposForAuthenticatedUser200Response{}
	this.TotalCount = totalCount
	this.Repositories = repositories
	return &this
}

// NewAppsListInstallationReposForAuthenticatedUser200ResponseWithDefaults instantiates a new AppsListInstallationReposForAuthenticatedUser200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsListInstallationReposForAuthenticatedUser200ResponseWithDefaults() *AppsListInstallationReposForAuthenticatedUser200Response {
	this := AppsListInstallationReposForAuthenticatedUser200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AppsListInstallationReposForAuthenticatedUser200Response) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetRepositorySelection returns the RepositorySelection field value if set, zero value otherwise.
func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositorySelection() string {
	if o == nil || o.RepositorySelection == nil {
		var ret string
		return ret
	}
	return *o.RepositorySelection
}

// GetRepositorySelectionOk returns a tuple with the RepositorySelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositorySelectionOk() (*string, bool) {
	if o == nil || o.RepositorySelection == nil {
		return nil, false
	}
	return o.RepositorySelection, true
}

// HasRepositorySelection returns a boolean if a field has been set.
func (o *AppsListInstallationReposForAuthenticatedUser200Response) HasRepositorySelection() bool {
	if o != nil && o.RepositorySelection != nil {
		return true
	}

	return false
}

// SetRepositorySelection gets a reference to the given string and assigns it to the RepositorySelection field.
func (o *AppsListInstallationReposForAuthenticatedUser200Response) SetRepositorySelection(v string) {
	o.RepositorySelection = &v
}

// GetRepositories returns the Repositories field value
func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositories() []Repository {
	if o == nil {
		var ret []Repository
		return ret
	}

	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value
// and a boolean to check if the value has been set.
func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositoriesOk() ([]Repository, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repositories, true
}

// SetRepositories sets field value
func (o *AppsListInstallationReposForAuthenticatedUser200Response) SetRepositories(v []Repository) {
	o.Repositories = v
}

func (o AppsListInstallationReposForAuthenticatedUser200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.RepositorySelection != nil {
		toSerialize["repository_selection"] = o.RepositorySelection
	}
	if true {
		toSerialize["repositories"] = o.Repositories
	}
	return json.Marshal(toSerialize)
}

type NullableAppsListInstallationReposForAuthenticatedUser200Response struct {
	value *AppsListInstallationReposForAuthenticatedUser200Response
	isSet bool
}

func (v NullableAppsListInstallationReposForAuthenticatedUser200Response) Get() *AppsListInstallationReposForAuthenticatedUser200Response {
	return v.value
}

func (v *NullableAppsListInstallationReposForAuthenticatedUser200Response) Set(val *AppsListInstallationReposForAuthenticatedUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsListInstallationReposForAuthenticatedUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsListInstallationReposForAuthenticatedUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsListInstallationReposForAuthenticatedUser200Response(val *AppsListInstallationReposForAuthenticatedUser200Response) *NullableAppsListInstallationReposForAuthenticatedUser200Response {
	return &NullableAppsListInstallationReposForAuthenticatedUser200Response{value: val, isSet: true}
}

func (v NullableAppsListInstallationReposForAuthenticatedUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsListInstallationReposForAuthenticatedUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

