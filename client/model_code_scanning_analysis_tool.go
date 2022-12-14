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

// CodeScanningAnalysisTool struct for CodeScanningAnalysisTool
type CodeScanningAnalysisTool struct {
	// The name of the tool used to generate the code scanning analysis.
	Name *string `json:"name,omitempty"`
	// The version of the tool used to generate the code scanning analysis.
	Version NullableString `json:"version,omitempty"`
	// The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
	Guid NullableString `json:"guid,omitempty"`
}

// NewCodeScanningAnalysisTool instantiates a new CodeScanningAnalysisTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeScanningAnalysisTool() *CodeScanningAnalysisTool {
	this := CodeScanningAnalysisTool{}
	return &this
}

// NewCodeScanningAnalysisToolWithDefaults instantiates a new CodeScanningAnalysisTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeScanningAnalysisToolWithDefaults() *CodeScanningAnalysisTool {
	this := CodeScanningAnalysisTool{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CodeScanningAnalysisTool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysisTool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CodeScanningAnalysisTool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CodeScanningAnalysisTool) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeScanningAnalysisTool) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeScanningAnalysisTool) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *CodeScanningAnalysisTool) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *CodeScanningAnalysisTool) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *CodeScanningAnalysisTool) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *CodeScanningAnalysisTool) UnsetVersion() {
	o.Version.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeScanningAnalysisTool) GetGuid() string {
	if o == nil || o.Guid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeScanningAnalysisTool) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *CodeScanningAnalysisTool) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *CodeScanningAnalysisTool) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *CodeScanningAnalysisTool) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *CodeScanningAnalysisTool) UnsetGuid() {
	o.Guid.Unset()
}

func (o CodeScanningAnalysisTool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCodeScanningAnalysisTool struct {
	value *CodeScanningAnalysisTool
	isSet bool
}

func (v NullableCodeScanningAnalysisTool) Get() *CodeScanningAnalysisTool {
	return v.value
}

func (v *NullableCodeScanningAnalysisTool) Set(val *CodeScanningAnalysisTool) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeScanningAnalysisTool) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeScanningAnalysisTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeScanningAnalysisTool(val *CodeScanningAnalysisTool) *NullableCodeScanningAnalysisTool {
	return &NullableCodeScanningAnalysisTool{value: val, isSet: true}
}

func (v NullableCodeScanningAnalysisTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeScanningAnalysisTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


