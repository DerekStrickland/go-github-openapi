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

// ProjectsCreateCardRequestOneOf struct for ProjectsCreateCardRequestOneOf
type ProjectsCreateCardRequestOneOf struct {
	// The project card's note
	Note NullableString `json:"note"`
}

// NewProjectsCreateCardRequestOneOf instantiates a new ProjectsCreateCardRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsCreateCardRequestOneOf(note NullableString) *ProjectsCreateCardRequestOneOf {
	this := ProjectsCreateCardRequestOneOf{}
	this.Note = note
	return &this
}

// NewProjectsCreateCardRequestOneOfWithDefaults instantiates a new ProjectsCreateCardRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsCreateCardRequestOneOfWithDefaults() *ProjectsCreateCardRequestOneOf {
	this := ProjectsCreateCardRequestOneOf{}
	return &this
}

// GetNote returns the Note field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectsCreateCardRequestOneOf) GetNote() string {
	if o == nil || o.Note.Get() == nil {
		var ret string
		return ret
	}

	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsCreateCardRequestOneOf) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// SetNote sets field value
func (o *ProjectsCreateCardRequestOneOf) SetNote(v string) {
	o.Note.Set(&v)
}

func (o ProjectsCreateCardRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["note"] = o.Note.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsCreateCardRequestOneOf struct {
	value *ProjectsCreateCardRequestOneOf
	isSet bool
}

func (v NullableProjectsCreateCardRequestOneOf) Get() *ProjectsCreateCardRequestOneOf {
	return v.value
}

func (v *NullableProjectsCreateCardRequestOneOf) Set(val *ProjectsCreateCardRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsCreateCardRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsCreateCardRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsCreateCardRequestOneOf(val *ProjectsCreateCardRequestOneOf) *NullableProjectsCreateCardRequestOneOf {
	return &NullableProjectsCreateCardRequestOneOf{value: val, isSet: true}
}

func (v NullableProjectsCreateCardRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsCreateCardRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

