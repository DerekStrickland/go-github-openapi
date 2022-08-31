/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"fmt"
)

// ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 Update the source for the repository. Must include the branch name, and may optionally specify the subdirectory `/docs`. Possible values are `\"gh-pages\"`, `\"master\"`, and `\"master /docs\"`.
type ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 string

// List of repos_update_information_about_pages_site_request_source_anyOf_1
const (
	GH_PAGES ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 = "gh-pages"
	MASTER ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 = "master"
	MASTER__DOCS ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 = "master /docs"
)

// All allowed values of ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 enum
var AllowedReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1EnumValues = []ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1{
	"gh-pages",
	"master",
	"master /docs",
}

func (v *ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1(value)
	for _, existing := range AllowedReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1", value)
}

// NewReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1FromValue returns a pointer to a valid ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1FromValue(v string) (*ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1, error) {
	ev := ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1: valid values are %v", v, AllowedReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) IsValid() bool {
	for _, existing := range AllowedReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to repos_update_information_about_pages_site_request_source_anyOf_1 value
func (v ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) Ptr() *ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 {
	return &v
}

type NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 struct {
	value *ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1
	isSet bool
}

func (v NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) Get() *ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 {
	return v.value
}

func (v *NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) Set(val *ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1(val *ReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) *NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1 {
	return &NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1{value: val, isSet: true}
}

func (v NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposUpdateInformationAboutPagesSiteRequestSourceAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

