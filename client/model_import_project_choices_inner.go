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

// ImportProjectChoicesInner struct for ImportProjectChoicesInner
type ImportProjectChoicesInner struct {
	Vcs *string `json:"vcs,omitempty"`
	TfvcProject *string `json:"tfvc_project,omitempty"`
	HumanName *string `json:"human_name,omitempty"`
}

// NewImportProjectChoicesInner instantiates a new ImportProjectChoicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportProjectChoicesInner() *ImportProjectChoicesInner {
	this := ImportProjectChoicesInner{}
	return &this
}

// NewImportProjectChoicesInnerWithDefaults instantiates a new ImportProjectChoicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportProjectChoicesInnerWithDefaults() *ImportProjectChoicesInner {
	this := ImportProjectChoicesInner{}
	return &this
}

// GetVcs returns the Vcs field value if set, zero value otherwise.
func (o *ImportProjectChoicesInner) GetVcs() string {
	if o == nil || o.Vcs == nil {
		var ret string
		return ret
	}
	return *o.Vcs
}

// GetVcsOk returns a tuple with the Vcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportProjectChoicesInner) GetVcsOk() (*string, bool) {
	if o == nil || o.Vcs == nil {
		return nil, false
	}
	return o.Vcs, true
}

// HasVcs returns a boolean if a field has been set.
func (o *ImportProjectChoicesInner) HasVcs() bool {
	if o != nil && o.Vcs != nil {
		return true
	}

	return false
}

// SetVcs gets a reference to the given string and assigns it to the Vcs field.
func (o *ImportProjectChoicesInner) SetVcs(v string) {
	o.Vcs = &v
}

// GetTfvcProject returns the TfvcProject field value if set, zero value otherwise.
func (o *ImportProjectChoicesInner) GetTfvcProject() string {
	if o == nil || o.TfvcProject == nil {
		var ret string
		return ret
	}
	return *o.TfvcProject
}

// GetTfvcProjectOk returns a tuple with the TfvcProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportProjectChoicesInner) GetTfvcProjectOk() (*string, bool) {
	if o == nil || o.TfvcProject == nil {
		return nil, false
	}
	return o.TfvcProject, true
}

// HasTfvcProject returns a boolean if a field has been set.
func (o *ImportProjectChoicesInner) HasTfvcProject() bool {
	if o != nil && o.TfvcProject != nil {
		return true
	}

	return false
}

// SetTfvcProject gets a reference to the given string and assigns it to the TfvcProject field.
func (o *ImportProjectChoicesInner) SetTfvcProject(v string) {
	o.TfvcProject = &v
}

// GetHumanName returns the HumanName field value if set, zero value otherwise.
func (o *ImportProjectChoicesInner) GetHumanName() string {
	if o == nil || o.HumanName == nil {
		var ret string
		return ret
	}
	return *o.HumanName
}

// GetHumanNameOk returns a tuple with the HumanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportProjectChoicesInner) GetHumanNameOk() (*string, bool) {
	if o == nil || o.HumanName == nil {
		return nil, false
	}
	return o.HumanName, true
}

// HasHumanName returns a boolean if a field has been set.
func (o *ImportProjectChoicesInner) HasHumanName() bool {
	if o != nil && o.HumanName != nil {
		return true
	}

	return false
}

// SetHumanName gets a reference to the given string and assigns it to the HumanName field.
func (o *ImportProjectChoicesInner) SetHumanName(v string) {
	o.HumanName = &v
}

func (o ImportProjectChoicesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vcs != nil {
		toSerialize["vcs"] = o.Vcs
	}
	if o.TfvcProject != nil {
		toSerialize["tfvc_project"] = o.TfvcProject
	}
	if o.HumanName != nil {
		toSerialize["human_name"] = o.HumanName
	}
	return json.Marshal(toSerialize)
}

type NullableImportProjectChoicesInner struct {
	value *ImportProjectChoicesInner
	isSet bool
}

func (v NullableImportProjectChoicesInner) Get() *ImportProjectChoicesInner {
	return v.value
}

func (v *NullableImportProjectChoicesInner) Set(val *ImportProjectChoicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableImportProjectChoicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableImportProjectChoicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportProjectChoicesInner(val *ImportProjectChoicesInner) *NullableImportProjectChoicesInner {
	return &NullableImportProjectChoicesInner{value: val, isSet: true}
}

func (v NullableImportProjectChoicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportProjectChoicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

