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

// ScimUpdateAttributeForUserRequestOperationsInnerValue - struct for ScimUpdateAttributeForUserRequestOperationsInnerValue
type ScimUpdateAttributeForUserRequestOperationsInnerValue struct {
	ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf
	ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner *[]ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner
	String *string
}

// ScimUpdateAttributeForUserRequestOperationsInnerValueOneOfAsScimUpdateAttributeForUserRequestOperationsInnerValue is a convenience function that returns ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf wrapped in ScimUpdateAttributeForUserRequestOperationsInnerValue
func ScimUpdateAttributeForUserRequestOperationsInnerValueOneOfAsScimUpdateAttributeForUserRequestOperationsInnerValue(v *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf) ScimUpdateAttributeForUserRequestOperationsInnerValue {
	return ScimUpdateAttributeForUserRequestOperationsInnerValue{
		ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf: v,
	}
}

// []ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1InnerAsScimUpdateAttributeForUserRequestOperationsInnerValue is a convenience function that returns []ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner wrapped in ScimUpdateAttributeForUserRequestOperationsInnerValue
func ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1InnerAsScimUpdateAttributeForUserRequestOperationsInnerValue(v *[]ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) ScimUpdateAttributeForUserRequestOperationsInnerValue {
	return ScimUpdateAttributeForUserRequestOperationsInnerValue{
		ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner: v,
	}
}

// stringAsScimUpdateAttributeForUserRequestOperationsInnerValue is a convenience function that returns string wrapped in ScimUpdateAttributeForUserRequestOperationsInnerValue
func StringAsScimUpdateAttributeForUserRequestOperationsInnerValue(v *string) ScimUpdateAttributeForUserRequestOperationsInnerValue {
	return ScimUpdateAttributeForUserRequestOperationsInnerValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ScimUpdateAttributeForUserRequestOperationsInnerValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf
	err = newStrictDecoder(data).Decode(&dst.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf)
	if err == nil {
		jsonScimUpdateAttributeForUserRequestOperationsInnerValueOneOf, _ := json.Marshal(dst.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf)
		if string(jsonScimUpdateAttributeForUserRequestOperationsInnerValueOneOf) == "{}" { // empty struct
			dst.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf = nil
	}

	// try to unmarshal data into ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner)
	if err == nil {
		jsonArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner, _ := json.Marshal(dst.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner)
		if string(jsonArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) == "{}" { // empty struct
			dst.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner = nil
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
		dst.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf = nil
		dst.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ScimUpdateAttributeForUserRequestOperationsInnerValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ScimUpdateAttributeForUserRequestOperationsInnerValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ScimUpdateAttributeForUserRequestOperationsInnerValue) MarshalJSON() ([]byte, error) {
	if src.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf != nil {
		return json.Marshal(&src.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf)
	}

	if src.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner != nil {
		return json.Marshal(&src.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ScimUpdateAttributeForUserRequestOperationsInnerValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf != nil {
		return obj.ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf
	}

	if obj.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner != nil {
		return obj.ArrayOfScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableScimUpdateAttributeForUserRequestOperationsInnerValue struct {
	value *ScimUpdateAttributeForUserRequestOperationsInnerValue
	isSet bool
}

func (v NullableScimUpdateAttributeForUserRequestOperationsInnerValue) Get() *ScimUpdateAttributeForUserRequestOperationsInnerValue {
	return v.value
}

func (v *NullableScimUpdateAttributeForUserRequestOperationsInnerValue) Set(val *ScimUpdateAttributeForUserRequestOperationsInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUpdateAttributeForUserRequestOperationsInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUpdateAttributeForUserRequestOperationsInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUpdateAttributeForUserRequestOperationsInnerValue(val *ScimUpdateAttributeForUserRequestOperationsInnerValue) *NullableScimUpdateAttributeForUserRequestOperationsInnerValue {
	return &NullableScimUpdateAttributeForUserRequestOperationsInnerValue{value: val, isSet: true}
}

func (v NullableScimUpdateAttributeForUserRequestOperationsInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUpdateAttributeForUserRequestOperationsInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

