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

// ReposGetContent200Response - struct for ReposGetContent200Response
type ReposGetContent200Response struct {
	ContentFile *ContentFile
	ContentSubmodule *ContentSubmodule
	ContentSymlink *ContentSymlink
	ArrayOfContentDirectoryInner *[]ContentDirectoryInner
}

// ContentFileAsReposGetContent200Response is a convenience function that returns ContentFile wrapped in ReposGetContent200Response
func ContentFileAsReposGetContent200Response(v *ContentFile) ReposGetContent200Response {
	return ReposGetContent200Response{
		ContentFile: v,
	}
}

// ContentSubmoduleAsReposGetContent200Response is a convenience function that returns ContentSubmodule wrapped in ReposGetContent200Response
func ContentSubmoduleAsReposGetContent200Response(v *ContentSubmodule) ReposGetContent200Response {
	return ReposGetContent200Response{
		ContentSubmodule: v,
	}
}

// ContentSymlinkAsReposGetContent200Response is a convenience function that returns ContentSymlink wrapped in ReposGetContent200Response
func ContentSymlinkAsReposGetContent200Response(v *ContentSymlink) ReposGetContent200Response {
	return ReposGetContent200Response{
		ContentSymlink: v,
	}
}

// []ContentDirectoryInnerAsReposGetContent200Response is a convenience function that returns []ContentDirectoryInner wrapped in ReposGetContent200Response
func ArrayOfContentDirectoryInnerAsReposGetContent200Response(v *[]ContentDirectoryInner) ReposGetContent200Response {
	return ReposGetContent200Response{
		ArrayOfContentDirectoryInner: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReposGetContent200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentFile
	err = newStrictDecoder(data).Decode(&dst.ContentFile)
	if err == nil {
		jsonContentFile, _ := json.Marshal(dst.ContentFile)
		if string(jsonContentFile) == "{}" { // empty struct
			dst.ContentFile = nil
		} else {
			match++
		}
	} else {
		dst.ContentFile = nil
	}

	// try to unmarshal data into ContentSubmodule
	err = newStrictDecoder(data).Decode(&dst.ContentSubmodule)
	if err == nil {
		jsonContentSubmodule, _ := json.Marshal(dst.ContentSubmodule)
		if string(jsonContentSubmodule) == "{}" { // empty struct
			dst.ContentSubmodule = nil
		} else {
			match++
		}
	} else {
		dst.ContentSubmodule = nil
	}

	// try to unmarshal data into ContentSymlink
	err = newStrictDecoder(data).Decode(&dst.ContentSymlink)
	if err == nil {
		jsonContentSymlink, _ := json.Marshal(dst.ContentSymlink)
		if string(jsonContentSymlink) == "{}" { // empty struct
			dst.ContentSymlink = nil
		} else {
			match++
		}
	} else {
		dst.ContentSymlink = nil
	}

	// try to unmarshal data into ArrayOfContentDirectoryInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfContentDirectoryInner)
	if err == nil {
		jsonArrayOfContentDirectoryInner, _ := json.Marshal(dst.ArrayOfContentDirectoryInner)
		if string(jsonArrayOfContentDirectoryInner) == "{}" { // empty struct
			dst.ArrayOfContentDirectoryInner = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfContentDirectoryInner = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContentFile = nil
		dst.ContentSubmodule = nil
		dst.ContentSymlink = nil
		dst.ArrayOfContentDirectoryInner = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ReposGetContent200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ReposGetContent200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReposGetContent200Response) MarshalJSON() ([]byte, error) {
	if src.ContentFile != nil {
		return json.Marshal(&src.ContentFile)
	}

	if src.ContentSubmodule != nil {
		return json.Marshal(&src.ContentSubmodule)
	}

	if src.ContentSymlink != nil {
		return json.Marshal(&src.ContentSymlink)
	}

	if src.ArrayOfContentDirectoryInner != nil {
		return json.Marshal(&src.ArrayOfContentDirectoryInner)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReposGetContent200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContentFile != nil {
		return obj.ContentFile
	}

	if obj.ContentSubmodule != nil {
		return obj.ContentSubmodule
	}

	if obj.ContentSymlink != nil {
		return obj.ContentSymlink
	}

	if obj.ArrayOfContentDirectoryInner != nil {
		return obj.ArrayOfContentDirectoryInner
	}

	// all schemas are nil
	return nil
}

type NullableReposGetContent200Response struct {
	value *ReposGetContent200Response
	isSet bool
}

func (v NullableReposGetContent200Response) Get() *ReposGetContent200Response {
	return v.value
}

func (v *NullableReposGetContent200Response) Set(val *ReposGetContent200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReposGetContent200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReposGetContent200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposGetContent200Response(val *ReposGetContent200Response) *NullableReposGetContent200Response {
	return &NullableReposGetContent200Response{value: val, isSet: true}
}

func (v NullableReposGetContent200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposGetContent200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


