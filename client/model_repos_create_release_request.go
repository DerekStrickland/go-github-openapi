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

// ReposCreateReleaseRequest struct for ReposCreateReleaseRequest
type ReposCreateReleaseRequest struct {
	// The name of the tag.
	TagName string `json:"tag_name"`
	// Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch (usually `master`).
	TargetCommitish *string `json:"target_commitish,omitempty"`
	// The name of the release.
	Name *string `json:"name,omitempty"`
	// Text describing the contents of the tag.
	Body *string `json:"body,omitempty"`
	// `true` to create a draft (unpublished) release, `false` to create a published one.
	Draft *bool `json:"draft,omitempty"`
	// `true` to identify the release as a prerelease. `false` to identify the release as a full release.
	Prerelease *bool `json:"prerelease,omitempty"`
	// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see \"[Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).\"
	DiscussionCategoryName *string `json:"discussion_category_name,omitempty"`
	// Whether to automatically generate the name and body for this release. If `name` is specified, the specified name will be used; otherwise, a name will be automatically generated. If `body` is specified, the body will be pre-pended to the automatically generated notes.
	GenerateReleaseNotes *bool `json:"generate_release_notes,omitempty"`
}

// NewReposCreateReleaseRequest instantiates a new ReposCreateReleaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposCreateReleaseRequest(tagName string) *ReposCreateReleaseRequest {
	this := ReposCreateReleaseRequest{}
	this.TagName = tagName
	var draft bool = false
	this.Draft = &draft
	var prerelease bool = false
	this.Prerelease = &prerelease
	var generateReleaseNotes bool = false
	this.GenerateReleaseNotes = &generateReleaseNotes
	return &this
}

// NewReposCreateReleaseRequestWithDefaults instantiates a new ReposCreateReleaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposCreateReleaseRequestWithDefaults() *ReposCreateReleaseRequest {
	this := ReposCreateReleaseRequest{}
	var draft bool = false
	this.Draft = &draft
	var prerelease bool = false
	this.Prerelease = &prerelease
	var generateReleaseNotes bool = false
	this.GenerateReleaseNotes = &generateReleaseNotes
	return &this
}

// GetTagName returns the TagName field value
func (o *ReposCreateReleaseRequest) GetTagName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagName, true
}

// SetTagName sets field value
func (o *ReposCreateReleaseRequest) SetTagName(v string) {
	o.TagName = v
}

// GetTargetCommitish returns the TargetCommitish field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetTargetCommitish() string {
	if o == nil || o.TargetCommitish == nil {
		var ret string
		return ret
	}
	return *o.TargetCommitish
}

// GetTargetCommitishOk returns a tuple with the TargetCommitish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetTargetCommitishOk() (*string, bool) {
	if o == nil || o.TargetCommitish == nil {
		return nil, false
	}
	return o.TargetCommitish, true
}

// HasTargetCommitish returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasTargetCommitish() bool {
	if o != nil && o.TargetCommitish != nil {
		return true
	}

	return false
}

// SetTargetCommitish gets a reference to the given string and assigns it to the TargetCommitish field.
func (o *ReposCreateReleaseRequest) SetTargetCommitish(v string) {
	o.TargetCommitish = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReposCreateReleaseRequest) SetName(v string) {
	o.Name = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ReposCreateReleaseRequest) SetBody(v string) {
	o.Body = &v
}

// GetDraft returns the Draft field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetDraft() bool {
	if o == nil || o.Draft == nil {
		var ret bool
		return ret
	}
	return *o.Draft
}

// GetDraftOk returns a tuple with the Draft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetDraftOk() (*bool, bool) {
	if o == nil || o.Draft == nil {
		return nil, false
	}
	return o.Draft, true
}

// HasDraft returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasDraft() bool {
	if o != nil && o.Draft != nil {
		return true
	}

	return false
}

// SetDraft gets a reference to the given bool and assigns it to the Draft field.
func (o *ReposCreateReleaseRequest) SetDraft(v bool) {
	o.Draft = &v
}

// GetPrerelease returns the Prerelease field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetPrerelease() bool {
	if o == nil || o.Prerelease == nil {
		var ret bool
		return ret
	}
	return *o.Prerelease
}

// GetPrereleaseOk returns a tuple with the Prerelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetPrereleaseOk() (*bool, bool) {
	if o == nil || o.Prerelease == nil {
		return nil, false
	}
	return o.Prerelease, true
}

// HasPrerelease returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasPrerelease() bool {
	if o != nil && o.Prerelease != nil {
		return true
	}

	return false
}

// SetPrerelease gets a reference to the given bool and assigns it to the Prerelease field.
func (o *ReposCreateReleaseRequest) SetPrerelease(v bool) {
	o.Prerelease = &v
}

// GetDiscussionCategoryName returns the DiscussionCategoryName field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetDiscussionCategoryName() string {
	if o == nil || o.DiscussionCategoryName == nil {
		var ret string
		return ret
	}
	return *o.DiscussionCategoryName
}

// GetDiscussionCategoryNameOk returns a tuple with the DiscussionCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetDiscussionCategoryNameOk() (*string, bool) {
	if o == nil || o.DiscussionCategoryName == nil {
		return nil, false
	}
	return o.DiscussionCategoryName, true
}

// HasDiscussionCategoryName returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasDiscussionCategoryName() bool {
	if o != nil && o.DiscussionCategoryName != nil {
		return true
	}

	return false
}

// SetDiscussionCategoryName gets a reference to the given string and assigns it to the DiscussionCategoryName field.
func (o *ReposCreateReleaseRequest) SetDiscussionCategoryName(v string) {
	o.DiscussionCategoryName = &v
}

// GetGenerateReleaseNotes returns the GenerateReleaseNotes field value if set, zero value otherwise.
func (o *ReposCreateReleaseRequest) GetGenerateReleaseNotes() bool {
	if o == nil || o.GenerateReleaseNotes == nil {
		var ret bool
		return ret
	}
	return *o.GenerateReleaseNotes
}

// GetGenerateReleaseNotesOk returns a tuple with the GenerateReleaseNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateReleaseRequest) GetGenerateReleaseNotesOk() (*bool, bool) {
	if o == nil || o.GenerateReleaseNotes == nil {
		return nil, false
	}
	return o.GenerateReleaseNotes, true
}

// HasGenerateReleaseNotes returns a boolean if a field has been set.
func (o *ReposCreateReleaseRequest) HasGenerateReleaseNotes() bool {
	if o != nil && o.GenerateReleaseNotes != nil {
		return true
	}

	return false
}

// SetGenerateReleaseNotes gets a reference to the given bool and assigns it to the GenerateReleaseNotes field.
func (o *ReposCreateReleaseRequest) SetGenerateReleaseNotes(v bool) {
	o.GenerateReleaseNotes = &v
}

func (o ReposCreateReleaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tag_name"] = o.TagName
	}
	if o.TargetCommitish != nil {
		toSerialize["target_commitish"] = o.TargetCommitish
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Draft != nil {
		toSerialize["draft"] = o.Draft
	}
	if o.Prerelease != nil {
		toSerialize["prerelease"] = o.Prerelease
	}
	if o.DiscussionCategoryName != nil {
		toSerialize["discussion_category_name"] = o.DiscussionCategoryName
	}
	if o.GenerateReleaseNotes != nil {
		toSerialize["generate_release_notes"] = o.GenerateReleaseNotes
	}
	return json.Marshal(toSerialize)
}

type NullableReposCreateReleaseRequest struct {
	value *ReposCreateReleaseRequest
	isSet bool
}

func (v NullableReposCreateReleaseRequest) Get() *ReposCreateReleaseRequest {
	return v.value
}

func (v *NullableReposCreateReleaseRequest) Set(val *ReposCreateReleaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposCreateReleaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposCreateReleaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposCreateReleaseRequest(val *ReposCreateReleaseRequest) *NullableReposCreateReleaseRequest {
	return &NullableReposCreateReleaseRequest{value: val, isSet: true}
}

func (v NullableReposCreateReleaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposCreateReleaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

