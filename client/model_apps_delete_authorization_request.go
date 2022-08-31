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

// AppsDeleteAuthorizationRequest struct for AppsDeleteAuthorizationRequest
type AppsDeleteAuthorizationRequest struct {
	// The OAuth access token used to authenticate to the GitHub API.
	AccessToken string `json:"access_token"`
}

// NewAppsDeleteAuthorizationRequest instantiates a new AppsDeleteAuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsDeleteAuthorizationRequest(accessToken string) *AppsDeleteAuthorizationRequest {
	this := AppsDeleteAuthorizationRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewAppsDeleteAuthorizationRequestWithDefaults instantiates a new AppsDeleteAuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsDeleteAuthorizationRequestWithDefaults() *AppsDeleteAuthorizationRequest {
	this := AppsDeleteAuthorizationRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AppsDeleteAuthorizationRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AppsDeleteAuthorizationRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AppsDeleteAuthorizationRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o AppsDeleteAuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	return json.Marshal(toSerialize)
}

type NullableAppsDeleteAuthorizationRequest struct {
	value *AppsDeleteAuthorizationRequest
	isSet bool
}

func (v NullableAppsDeleteAuthorizationRequest) Get() *AppsDeleteAuthorizationRequest {
	return v.value
}

func (v *NullableAppsDeleteAuthorizationRequest) Set(val *AppsDeleteAuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsDeleteAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsDeleteAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsDeleteAuthorizationRequest(val *AppsDeleteAuthorizationRequest) *NullableAppsDeleteAuthorizationRequest {
	return &NullableAppsDeleteAuthorizationRequest{value: val, isSet: true}
}

func (v NullableAppsDeleteAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsDeleteAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


