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

// AppsCreateInstallationAccessTokenRequest struct for AppsCreateInstallationAccessTokenRequest
type AppsCreateInstallationAccessTokenRequest struct {
	// List of repository names that the token should have access to
	Repositories []string `json:"repositories,omitempty"`
	// List of repository IDs that the token should have access to
	RepositoryIds []int32 `json:"repository_ids,omitempty"`
	Permissions *AppPermissions `json:"permissions,omitempty"`
}

// NewAppsCreateInstallationAccessTokenRequest instantiates a new AppsCreateInstallationAccessTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsCreateInstallationAccessTokenRequest() *AppsCreateInstallationAccessTokenRequest {
	this := AppsCreateInstallationAccessTokenRequest{}
	return &this
}

// NewAppsCreateInstallationAccessTokenRequestWithDefaults instantiates a new AppsCreateInstallationAccessTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsCreateInstallationAccessTokenRequestWithDefaults() *AppsCreateInstallationAccessTokenRequest {
	this := AppsCreateInstallationAccessTokenRequest{}
	return &this
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *AppsCreateInstallationAccessTokenRequest) GetRepositories() []string {
	if o == nil || o.Repositories == nil {
		var ret []string
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCreateInstallationAccessTokenRequest) GetRepositoriesOk() ([]string, bool) {
	if o == nil || o.Repositories == nil {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *AppsCreateInstallationAccessTokenRequest) HasRepositories() bool {
	if o != nil && o.Repositories != nil {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []string and assigns it to the Repositories field.
func (o *AppsCreateInstallationAccessTokenRequest) SetRepositories(v []string) {
	o.Repositories = v
}

// GetRepositoryIds returns the RepositoryIds field value if set, zero value otherwise.
func (o *AppsCreateInstallationAccessTokenRequest) GetRepositoryIds() []int32 {
	if o == nil || o.RepositoryIds == nil {
		var ret []int32
		return ret
	}
	return o.RepositoryIds
}

// GetRepositoryIdsOk returns a tuple with the RepositoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCreateInstallationAccessTokenRequest) GetRepositoryIdsOk() ([]int32, bool) {
	if o == nil || o.RepositoryIds == nil {
		return nil, false
	}
	return o.RepositoryIds, true
}

// HasRepositoryIds returns a boolean if a field has been set.
func (o *AppsCreateInstallationAccessTokenRequest) HasRepositoryIds() bool {
	if o != nil && o.RepositoryIds != nil {
		return true
	}

	return false
}

// SetRepositoryIds gets a reference to the given []int32 and assigns it to the RepositoryIds field.
func (o *AppsCreateInstallationAccessTokenRequest) SetRepositoryIds(v []int32) {
	o.RepositoryIds = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AppsCreateInstallationAccessTokenRequest) GetPermissions() AppPermissions {
	if o == nil || o.Permissions == nil {
		var ret AppPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCreateInstallationAccessTokenRequest) GetPermissionsOk() (*AppPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AppsCreateInstallationAccessTokenRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given AppPermissions and assigns it to the Permissions field.
func (o *AppsCreateInstallationAccessTokenRequest) SetPermissions(v AppPermissions) {
	o.Permissions = &v
}

func (o AppsCreateInstallationAccessTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Repositories != nil {
		toSerialize["repositories"] = o.Repositories
	}
	if o.RepositoryIds != nil {
		toSerialize["repository_ids"] = o.RepositoryIds
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAppsCreateInstallationAccessTokenRequest struct {
	value *AppsCreateInstallationAccessTokenRequest
	isSet bool
}

func (v NullableAppsCreateInstallationAccessTokenRequest) Get() *AppsCreateInstallationAccessTokenRequest {
	return v.value
}

func (v *NullableAppsCreateInstallationAccessTokenRequest) Set(val *AppsCreateInstallationAccessTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsCreateInstallationAccessTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsCreateInstallationAccessTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsCreateInstallationAccessTokenRequest(val *AppsCreateInstallationAccessTokenRequest) *NullableAppsCreateInstallationAccessTokenRequest {
	return &NullableAppsCreateInstallationAccessTokenRequest{value: val, isSet: true}
}

func (v NullableAppsCreateInstallationAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsCreateInstallationAccessTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


