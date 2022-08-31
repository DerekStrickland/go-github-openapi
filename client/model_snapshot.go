/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// Snapshot Create a new snapshot of a repository's dependencies.
type Snapshot struct {
	// The version of the repository snapshot submission.
	Version int32 `json:"version"`
	Job SnapshotJob `json:"job"`
	// The commit SHA associated with this dependency snapshot.
	Sha string `json:"sha"`
	// The repository branch that triggered this snapshot.
	Ref string `json:"ref"`
	Detector SnapshotDetector `json:"detector"`
	// User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Metadata *map[string]Metadata1 `json:"metadata,omitempty"`
	// A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Manifests *map[string]Manifest `json:"manifests,omitempty"`
	// The time at which the snapshot was scanned.
	Scanned time.Time `json:"scanned"`
}

// NewSnapshot instantiates a new Snapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshot(version int32, job SnapshotJob, sha string, ref string, detector SnapshotDetector, scanned time.Time) *Snapshot {
	this := Snapshot{}
	this.Version = version
	this.Job = job
	this.Sha = sha
	this.Ref = ref
	this.Detector = detector
	this.Scanned = scanned
	return &this
}

// NewSnapshotWithDefaults instantiates a new Snapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotWithDefaults() *Snapshot {
	this := Snapshot{}
	return &this
}

// GetVersion returns the Version field value
func (o *Snapshot) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Snapshot) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Snapshot) SetVersion(v int32) {
	o.Version = v
}

// GetJob returns the Job field value
func (o *Snapshot) GetJob() SnapshotJob {
	if o == nil {
		var ret SnapshotJob
		return ret
	}

	return o.Job
}

// GetJobOk returns a tuple with the Job field value
// and a boolean to check if the value has been set.
func (o *Snapshot) GetJobOk() (*SnapshotJob, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Job, true
}

// SetJob sets field value
func (o *Snapshot) SetJob(v SnapshotJob) {
	o.Job = v
}

// GetSha returns the Sha field value
func (o *Snapshot) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *Snapshot) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *Snapshot) SetSha(v string) {
	o.Sha = v
}

// GetRef returns the Ref field value
func (o *Snapshot) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *Snapshot) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *Snapshot) SetRef(v string) {
	o.Ref = v
}

// GetDetector returns the Detector field value
func (o *Snapshot) GetDetector() SnapshotDetector {
	if o == nil {
		var ret SnapshotDetector
		return ret
	}

	return o.Detector
}

// GetDetectorOk returns a tuple with the Detector field value
// and a boolean to check if the value has been set.
func (o *Snapshot) GetDetectorOk() (*SnapshotDetector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detector, true
}

// SetDetector sets field value
func (o *Snapshot) SetDetector(v SnapshotDetector) {
	o.Detector = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Snapshot) GetMetadata() map[string]Metadata1 {
	if o == nil || o.Metadata == nil {
		var ret map[string]Metadata1
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetMetadataOk() (*map[string]Metadata1, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Snapshot) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]Metadata1 and assigns it to the Metadata field.
func (o *Snapshot) SetMetadata(v map[string]Metadata1) {
	o.Metadata = &v
}

// GetManifests returns the Manifests field value if set, zero value otherwise.
func (o *Snapshot) GetManifests() map[string]Manifest {
	if o == nil || o.Manifests == nil {
		var ret map[string]Manifest
		return ret
	}
	return *o.Manifests
}

// GetManifestsOk returns a tuple with the Manifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snapshot) GetManifestsOk() (*map[string]Manifest, bool) {
	if o == nil || o.Manifests == nil {
		return nil, false
	}
	return o.Manifests, true
}

// HasManifests returns a boolean if a field has been set.
func (o *Snapshot) HasManifests() bool {
	if o != nil && o.Manifests != nil {
		return true
	}

	return false
}

// SetManifests gets a reference to the given map[string]Manifest and assigns it to the Manifests field.
func (o *Snapshot) SetManifests(v map[string]Manifest) {
	o.Manifests = &v
}

// GetScanned returns the Scanned field value
func (o *Snapshot) GetScanned() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Scanned
}

// GetScannedOk returns a tuple with the Scanned field value
// and a boolean to check if the value has been set.
func (o *Snapshot) GetScannedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scanned, true
}

// SetScanned sets field value
func (o *Snapshot) SetScanned(v time.Time) {
	o.Scanned = v
}

func (o Snapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["job"] = o.Job
	}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["ref"] = o.Ref
	}
	if true {
		toSerialize["detector"] = o.Detector
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Manifests != nil {
		toSerialize["manifests"] = o.Manifests
	}
	if true {
		toSerialize["scanned"] = o.Scanned
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshot struct {
	value *Snapshot
	isSet bool
}

func (v NullableSnapshot) Get() *Snapshot {
	return v.value
}

func (v *NullableSnapshot) Set(val *Snapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshot(val *Snapshot) *NullableSnapshot {
	return &NullableSnapshot{value: val, isSet: true}
}

func (v NullableSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


