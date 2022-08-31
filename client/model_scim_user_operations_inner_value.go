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

// ScimUserOperationsInnerValue - struct for ScimUserOperationsInnerValue
type ScimUserOperationsInnerValue struct {
	ArrayOfInterface{} *[]interface{}
	MapmapOfStringinterface{} *map[string]interface{}
	String *string
}

// []interface{}AsScimUserOperationsInnerValue is a convenience function that returns []interface{} wrapped in ScimUserOperationsInnerValue
func ArrayOfInterface{}AsScimUserOperationsInnerValue(v *[]interface{}) ScimUserOperationsInnerValue {
	return ScimUserOperationsInnerValue{
		ArrayOfInterface{}: v,
	}
}

// map[string]interface{}AsScimUserOperationsInnerValue is a convenience function that returns map[string]interface{} wrapped in ScimUserOperationsInnerValue
func MapmapOfStringinterface{}AsScimUserOperationsInnerValue(v *map[string]interface{}) ScimUserOperationsInnerValue {
	return ScimUserOperationsInnerValue{
		MapmapOfStringinterface{}: v,
	}
}

// stringAsScimUserOperationsInnerValue is a convenience function that returns string wrapped in ScimUserOperationsInnerValue
func StringAsScimUserOperationsInnerValue(v *string) ScimUserOperationsInnerValue {
	return ScimUserOperationsInnerValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ScimUserOperationsInnerValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfInterface{}
	err = newStrictDecoder(data).Decode(&dst.ArrayOfInterface{})
	if err == nil {
		jsonArrayOfInterface{}, _ := json.Marshal(dst.ArrayOfInterface{})
		if string(jsonArrayOfInterface{}) == "{}" { // empty struct
			dst.ArrayOfInterface{} = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfInterface{} = nil
	}

	// try to unmarshal data into MapmapOfStringinterface{}
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringinterface{})
	if err == nil {
		jsonMapmapOfStringinterface{}, _ := json.Marshal(dst.MapmapOfStringinterface{})
		if string(jsonMapmapOfStringinterface{}) == "{}" { // empty struct
			dst.MapmapOfStringinterface{} = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringinterface{} = nil
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
		dst.ArrayOfInterface{} = nil
		dst.MapmapOfStringinterface{} = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ScimUserOperationsInnerValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ScimUserOperationsInnerValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ScimUserOperationsInnerValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfInterface{} != nil {
		return json.Marshal(&src.ArrayOfInterface{})
	}

	if src.MapmapOfStringinterface{} != nil {
		return json.Marshal(&src.MapmapOfStringinterface{})
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ScimUserOperationsInnerValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfInterface{} != nil {
		return obj.ArrayOfInterface{}
	}

	if obj.MapmapOfStringinterface{} != nil {
		return obj.MapmapOfStringinterface{}
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableScimUserOperationsInnerValue struct {
	value *ScimUserOperationsInnerValue
	isSet bool
}

func (v NullableScimUserOperationsInnerValue) Get() *ScimUserOperationsInnerValue {
	return v.value
}

func (v *NullableScimUserOperationsInnerValue) Set(val *ScimUserOperationsInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUserOperationsInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUserOperationsInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUserOperationsInnerValue(val *ScimUserOperationsInnerValue) *NullableScimUserOperationsInnerValue {
	return &NullableScimUserOperationsInnerValue{value: val, isSet: true}
}

func (v NullableScimUserOperationsInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUserOperationsInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

