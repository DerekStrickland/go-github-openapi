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

// Manifest struct for Manifest
type Manifest struct {
	// The name of the manifest.
	Name string `json:"name"`
	File *ManifestFile `json:"file,omitempty"`
	// User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Metadata *map[string]Metadata1 `json:"metadata,omitempty"`
	// A collection of resolved package dependencies.
	Resolved *map[string]Dependency `json:"resolved,omitempty"`
}

// NewManifest instantiates a new Manifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManifest(name string) *Manifest {
	this := Manifest{}
	this.Name = name
	return &this
}

// NewManifestWithDefaults instantiates a new Manifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManifestWithDefaults() *Manifest {
	this := Manifest{}
	return &this
}

// GetName returns the Name field value
func (o *Manifest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Manifest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Manifest) SetName(v string) {
	o.Name = v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *Manifest) GetFile() ManifestFile {
	if o == nil || o.File == nil {
		var ret ManifestFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manifest) GetFileOk() (*ManifestFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *Manifest) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given ManifestFile and assigns it to the File field.
func (o *Manifest) SetFile(v ManifestFile) {
	o.File = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Manifest) GetMetadata() map[string]Metadata1 {
	if o == nil || o.Metadata == nil {
		var ret map[string]Metadata1
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manifest) GetMetadataOk() (*map[string]Metadata1, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Manifest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]Metadata1 and assigns it to the Metadata field.
func (o *Manifest) SetMetadata(v map[string]Metadata1) {
	o.Metadata = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *Manifest) GetResolved() map[string]Dependency {
	if o == nil || o.Resolved == nil {
		var ret map[string]Dependency
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Manifest) GetResolvedOk() (*map[string]Dependency, bool) {
	if o == nil || o.Resolved == nil {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *Manifest) HasResolved() bool {
	if o != nil && o.Resolved != nil {
		return true
	}

	return false
}

// SetResolved gets a reference to the given map[string]Dependency and assigns it to the Resolved field.
func (o *Manifest) SetResolved(v map[string]Dependency) {
	o.Resolved = &v
}

func (o Manifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Resolved != nil {
		toSerialize["resolved"] = o.Resolved
	}
	return json.Marshal(toSerialize)
}

type NullableManifest struct {
	value *Manifest
	isSet bool
}

func (v NullableManifest) Get() *Manifest {
	return v.value
}

func (v *NullableManifest) Set(val *Manifest) {
	v.value = val
	v.isSet = true
}

func (v NullableManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManifest(val *Manifest) *NullableManifest {
	return &NullableManifest{value: val, isSet: true}
}

func (v NullableManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


