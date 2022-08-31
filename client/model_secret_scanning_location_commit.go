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

// SecretScanningLocationCommit Represents a 'commit' secret scanning location type. This location type shows that a secret was detected inside a commit to a repository.
type SecretScanningLocationCommit struct {
	// The file path in the repository
	Path string `json:"path"`
	// Line number at which the secret starts in the file
	StartLine float32 `json:"start_line"`
	// Line number at which the secret ends in the file
	EndLine float32 `json:"end_line"`
	// The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	StartColumn float32 `json:"start_column"`
	// The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	EndColumn float32 `json:"end_column"`
	// SHA-1 hash ID of the associated blob
	BlobSha string `json:"blob_sha"`
	// The API URL to get the associated blob resource
	BlobUrl string `json:"blob_url"`
	// SHA-1 hash ID of the associated commit
	CommitSha string `json:"commit_sha"`
	// The API URL to get the associated commit resource
	CommitUrl string `json:"commit_url"`
}

// NewSecretScanningLocationCommit instantiates a new SecretScanningLocationCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretScanningLocationCommit(path string, startLine float32, endLine float32, startColumn float32, endColumn float32, blobSha string, blobUrl string, commitSha string, commitUrl string) *SecretScanningLocationCommit {
	this := SecretScanningLocationCommit{}
	this.Path = path
	this.StartLine = startLine
	this.EndLine = endLine
	this.StartColumn = startColumn
	this.EndColumn = endColumn
	this.BlobSha = blobSha
	this.BlobUrl = blobUrl
	this.CommitSha = commitSha
	this.CommitUrl = commitUrl
	return &this
}

// NewSecretScanningLocationCommitWithDefaults instantiates a new SecretScanningLocationCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretScanningLocationCommitWithDefaults() *SecretScanningLocationCommit {
	this := SecretScanningLocationCommit{}
	return &this
}

// GetPath returns the Path field value
func (o *SecretScanningLocationCommit) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SecretScanningLocationCommit) SetPath(v string) {
	o.Path = v
}

// GetStartLine returns the StartLine field value
func (o *SecretScanningLocationCommit) GetStartLine() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StartLine
}

// GetStartLineOk returns a tuple with the StartLine field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetStartLineOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartLine, true
}

// SetStartLine sets field value
func (o *SecretScanningLocationCommit) SetStartLine(v float32) {
	o.StartLine = v
}

// GetEndLine returns the EndLine field value
func (o *SecretScanningLocationCommit) GetEndLine() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EndLine
}

// GetEndLineOk returns a tuple with the EndLine field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetEndLineOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndLine, true
}

// SetEndLine sets field value
func (o *SecretScanningLocationCommit) SetEndLine(v float32) {
	o.EndLine = v
}

// GetStartColumn returns the StartColumn field value
func (o *SecretScanningLocationCommit) GetStartColumn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StartColumn
}

// GetStartColumnOk returns a tuple with the StartColumn field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetStartColumnOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartColumn, true
}

// SetStartColumn sets field value
func (o *SecretScanningLocationCommit) SetStartColumn(v float32) {
	o.StartColumn = v
}

// GetEndColumn returns the EndColumn field value
func (o *SecretScanningLocationCommit) GetEndColumn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EndColumn
}

// GetEndColumnOk returns a tuple with the EndColumn field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetEndColumnOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndColumn, true
}

// SetEndColumn sets field value
func (o *SecretScanningLocationCommit) SetEndColumn(v float32) {
	o.EndColumn = v
}

// GetBlobSha returns the BlobSha field value
func (o *SecretScanningLocationCommit) GetBlobSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobSha
}

// GetBlobShaOk returns a tuple with the BlobSha field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetBlobShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlobSha, true
}

// SetBlobSha sets field value
func (o *SecretScanningLocationCommit) SetBlobSha(v string) {
	o.BlobSha = v
}

// GetBlobUrl returns the BlobUrl field value
func (o *SecretScanningLocationCommit) GetBlobUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobUrl
}

// GetBlobUrlOk returns a tuple with the BlobUrl field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetBlobUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlobUrl, true
}

// SetBlobUrl sets field value
func (o *SecretScanningLocationCommit) SetBlobUrl(v string) {
	o.BlobUrl = v
}

// GetCommitSha returns the CommitSha field value
func (o *SecretScanningLocationCommit) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value
func (o *SecretScanningLocationCommit) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetCommitUrl returns the CommitUrl field value
func (o *SecretScanningLocationCommit) GetCommitUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitUrl
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value
// and a boolean to check if the value has been set.
func (o *SecretScanningLocationCommit) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitUrl, true
}

// SetCommitUrl sets field value
func (o *SecretScanningLocationCommit) SetCommitUrl(v string) {
	o.CommitUrl = v
}

func (o SecretScanningLocationCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["start_line"] = o.StartLine
	}
	if true {
		toSerialize["end_line"] = o.EndLine
	}
	if true {
		toSerialize["start_column"] = o.StartColumn
	}
	if true {
		toSerialize["end_column"] = o.EndColumn
	}
	if true {
		toSerialize["blob_sha"] = o.BlobSha
	}
	if true {
		toSerialize["blob_url"] = o.BlobUrl
	}
	if true {
		toSerialize["commit_sha"] = o.CommitSha
	}
	if true {
		toSerialize["commit_url"] = o.CommitUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSecretScanningLocationCommit struct {
	value *SecretScanningLocationCommit
	isSet bool
}

func (v NullableSecretScanningLocationCommit) Get() *SecretScanningLocationCommit {
	return v.value
}

func (v *NullableSecretScanningLocationCommit) Set(val *SecretScanningLocationCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretScanningLocationCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretScanningLocationCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretScanningLocationCommit(val *SecretScanningLocationCommit) *NullableSecretScanningLocationCommit {
	return &NullableSecretScanningLocationCommit{value: val, isSet: true}
}

func (v NullableSecretScanningLocationCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretScanningLocationCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

