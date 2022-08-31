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

// GitCreateTreeRequestTreeInner struct for GitCreateTreeRequestTreeInner
type GitCreateTreeRequestTreeInner struct {
	// The file referenced in the tree.
	Path *string `json:"path,omitempty"`
	// The file mode; one of `100644` for file (blob), `100755` for executable (blob), `040000` for subdirectory (tree), `160000` for submodule (commit), or `120000` for a blob that specifies the path of a symlink.
	Mode *string `json:"mode,omitempty"`
	// Either `blob`, `tree`, or `commit`.
	Type *string `json:"type,omitempty"`
	// The SHA1 checksum ID of the object in the tree. Also called `tree.sha`. If the value is `null` then the file will be deleted.      **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
	Sha NullableString `json:"sha,omitempty"`
	// The content you want this file to have. GitHub will write this blob out and use that SHA for this entry. Use either this, or `tree.sha`.      **Note:** Use either `tree.sha` or `content` to specify the contents of the entry. Using both `tree.sha` and `content` will return an error.
	Content *string `json:"content,omitempty"`
}

// NewGitCreateTreeRequestTreeInner instantiates a new GitCreateTreeRequestTreeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitCreateTreeRequestTreeInner() *GitCreateTreeRequestTreeInner {
	this := GitCreateTreeRequestTreeInner{}
	return &this
}

// NewGitCreateTreeRequestTreeInnerWithDefaults instantiates a new GitCreateTreeRequestTreeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitCreateTreeRequestTreeInnerWithDefaults() *GitCreateTreeRequestTreeInner {
	this := GitCreateTreeRequestTreeInner{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GitCreateTreeRequestTreeInner) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitCreateTreeRequestTreeInner) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GitCreateTreeRequestTreeInner) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GitCreateTreeRequestTreeInner) SetPath(v string) {
	o.Path = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GitCreateTreeRequestTreeInner) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitCreateTreeRequestTreeInner) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GitCreateTreeRequestTreeInner) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GitCreateTreeRequestTreeInner) SetMode(v string) {
	o.Mode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GitCreateTreeRequestTreeInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitCreateTreeRequestTreeInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GitCreateTreeRequestTreeInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GitCreateTreeRequestTreeInner) SetType(v string) {
	o.Type = &v
}

// GetSha returns the Sha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitCreateTreeRequestTreeInner) GetSha() string {
	if o == nil || o.Sha.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sha.Get()
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitCreateTreeRequestTreeInner) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha.Get(), o.Sha.IsSet()
}

// HasSha returns a boolean if a field has been set.
func (o *GitCreateTreeRequestTreeInner) HasSha() bool {
	if o != nil && o.Sha.IsSet() {
		return true
	}

	return false
}

// SetSha gets a reference to the given NullableString and assigns it to the Sha field.
func (o *GitCreateTreeRequestTreeInner) SetSha(v string) {
	o.Sha.Set(&v)
}
// SetShaNil sets the value for Sha to be an explicit nil
func (o *GitCreateTreeRequestTreeInner) SetShaNil() {
	o.Sha.Set(nil)
}

// UnsetSha ensures that no value is present for Sha, not even an explicit nil
func (o *GitCreateTreeRequestTreeInner) UnsetSha() {
	o.Sha.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *GitCreateTreeRequestTreeInner) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitCreateTreeRequestTreeInner) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *GitCreateTreeRequestTreeInner) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *GitCreateTreeRequestTreeInner) SetContent(v string) {
	o.Content = &v
}

func (o GitCreateTreeRequestTreeInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Sha.IsSet() {
		toSerialize["sha"] = o.Sha.Get()
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableGitCreateTreeRequestTreeInner struct {
	value *GitCreateTreeRequestTreeInner
	isSet bool
}

func (v NullableGitCreateTreeRequestTreeInner) Get() *GitCreateTreeRequestTreeInner {
	return v.value
}

func (v *NullableGitCreateTreeRequestTreeInner) Set(val *GitCreateTreeRequestTreeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGitCreateTreeRequestTreeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGitCreateTreeRequestTreeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitCreateTreeRequestTreeInner(val *GitCreateTreeRequestTreeInner) *NullableGitCreateTreeRequestTreeInner {
	return &NullableGitCreateTreeRequestTreeInner{value: val, isSet: true}
}

func (v NullableGitCreateTreeRequestTreeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitCreateTreeRequestTreeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


