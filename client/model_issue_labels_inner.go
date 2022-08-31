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

// IssueLabelsInner - struct for IssueLabelsInner
type IssueLabelsInner struct {
	IssueLabelsInnerOneOf *IssueLabelsInnerOneOf
	String *string
}

// IssueLabelsInnerOneOfAsIssueLabelsInner is a convenience function that returns IssueLabelsInnerOneOf wrapped in IssueLabelsInner
func IssueLabelsInnerOneOfAsIssueLabelsInner(v *IssueLabelsInnerOneOf) IssueLabelsInner {
	return IssueLabelsInner{
		IssueLabelsInnerOneOf: v,
	}
}

// stringAsIssueLabelsInner is a convenience function that returns string wrapped in IssueLabelsInner
func StringAsIssueLabelsInner(v *string) IssueLabelsInner {
	return IssueLabelsInner{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IssueLabelsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IssueLabelsInnerOneOf
	err = newStrictDecoder(data).Decode(&dst.IssueLabelsInnerOneOf)
	if err == nil {
		jsonIssueLabelsInnerOneOf, _ := json.Marshal(dst.IssueLabelsInnerOneOf)
		if string(jsonIssueLabelsInnerOneOf) == "{}" { // empty struct
			dst.IssueLabelsInnerOneOf = nil
		} else {
			match++
		}
	} else {
		dst.IssueLabelsInnerOneOf = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IssueLabelsInnerOneOf = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(IssueLabelsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(IssueLabelsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IssueLabelsInner) MarshalJSON() ([]byte, error) {
	if src.IssueLabelsInnerOneOf != nil {
		return json.Marshal(&src.IssueLabelsInnerOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IssueLabelsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IssueLabelsInnerOneOf != nil {
		return obj.IssueLabelsInnerOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableIssueLabelsInner struct {
	value *IssueLabelsInner
	isSet bool
}

func (v NullableIssueLabelsInner) Get() *IssueLabelsInner {
	return v.value
}

func (v *NullableIssueLabelsInner) Set(val *IssueLabelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueLabelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueLabelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueLabelsInner(val *IssueLabelsInner) *NullableIssueLabelsInner {
	return &NullableIssueLabelsInner{value: val, isSet: true}
}

func (v NullableIssueLabelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueLabelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


