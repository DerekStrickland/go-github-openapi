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

// CodeSearchResultItem Code Search Result Item
type CodeSearchResultItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Url string `json:"url"`
	GitUrl string `json:"git_url"`
	HtmlUrl string `json:"html_url"`
	Repository MinimalRepository `json:"repository"`
	Score float32 `json:"score"`
	FileSize *int32 `json:"file_size,omitempty"`
	Language NullableString `json:"language,omitempty"`
	LastModifiedAt *time.Time `json:"last_modified_at,omitempty"`
	LineNumbers []string `json:"line_numbers,omitempty"`
	TextMatches []SearchResultTextMatchesInner `json:"text_matches,omitempty"`
}

// NewCodeSearchResultItem instantiates a new CodeSearchResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeSearchResultItem(name string, path string, sha string, url string, gitUrl string, htmlUrl string, repository MinimalRepository, score float32) *CodeSearchResultItem {
	this := CodeSearchResultItem{}
	this.Name = name
	this.Path = path
	this.Sha = sha
	this.Url = url
	this.GitUrl = gitUrl
	this.HtmlUrl = htmlUrl
	this.Repository = repository
	this.Score = score
	return &this
}

// NewCodeSearchResultItemWithDefaults instantiates a new CodeSearchResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeSearchResultItemWithDefaults() *CodeSearchResultItem {
	this := CodeSearchResultItem{}
	return &this
}

// GetName returns the Name field value
func (o *CodeSearchResultItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CodeSearchResultItem) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *CodeSearchResultItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CodeSearchResultItem) SetPath(v string) {
	o.Path = v
}

// GetSha returns the Sha field value
func (o *CodeSearchResultItem) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *CodeSearchResultItem) SetSha(v string) {
	o.Sha = v
}

// GetUrl returns the Url field value
func (o *CodeSearchResultItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CodeSearchResultItem) SetUrl(v string) {
	o.Url = v
}

// GetGitUrl returns the GitUrl field value
func (o *CodeSearchResultItem) GetGitUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitUrl
}

// GetGitUrlOk returns a tuple with the GitUrl field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetGitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitUrl, true
}

// SetGitUrl sets field value
func (o *CodeSearchResultItem) SetGitUrl(v string) {
	o.GitUrl = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *CodeSearchResultItem) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *CodeSearchResultItem) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetRepository returns the Repository field value
func (o *CodeSearchResultItem) GetRepository() MinimalRepository {
	if o == nil {
		var ret MinimalRepository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetRepositoryOk() (*MinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *CodeSearchResultItem) SetRepository(v MinimalRepository) {
	o.Repository = v
}

// GetScore returns the Score field value
func (o *CodeSearchResultItem) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CodeSearchResultItem) SetScore(v float32) {
	o.Score = v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *CodeSearchResultItem) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *CodeSearchResultItem) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *CodeSearchResultItem) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeSearchResultItem) GetLanguage() string {
	if o == nil || o.Language.Get() == nil {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeSearchResultItem) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *CodeSearchResultItem) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *CodeSearchResultItem) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *CodeSearchResultItem) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *CodeSearchResultItem) UnsetLanguage() {
	o.Language.Unset()
}

// GetLastModifiedAt returns the LastModifiedAt field value if set, zero value otherwise.
func (o *CodeSearchResultItem) GetLastModifiedAt() time.Time {
	if o == nil || o.LastModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetLastModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedAt == nil {
		return nil, false
	}
	return o.LastModifiedAt, true
}

// HasLastModifiedAt returns a boolean if a field has been set.
func (o *CodeSearchResultItem) HasLastModifiedAt() bool {
	if o != nil && o.LastModifiedAt != nil {
		return true
	}

	return false
}

// SetLastModifiedAt gets a reference to the given time.Time and assigns it to the LastModifiedAt field.
func (o *CodeSearchResultItem) SetLastModifiedAt(v time.Time) {
	o.LastModifiedAt = &v
}

// GetLineNumbers returns the LineNumbers field value if set, zero value otherwise.
func (o *CodeSearchResultItem) GetLineNumbers() []string {
	if o == nil || o.LineNumbers == nil {
		var ret []string
		return ret
	}
	return o.LineNumbers
}

// GetLineNumbersOk returns a tuple with the LineNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetLineNumbersOk() ([]string, bool) {
	if o == nil || o.LineNumbers == nil {
		return nil, false
	}
	return o.LineNumbers, true
}

// HasLineNumbers returns a boolean if a field has been set.
func (o *CodeSearchResultItem) HasLineNumbers() bool {
	if o != nil && o.LineNumbers != nil {
		return true
	}

	return false
}

// SetLineNumbers gets a reference to the given []string and assigns it to the LineNumbers field.
func (o *CodeSearchResultItem) SetLineNumbers(v []string) {
	o.LineNumbers = v
}

// GetTextMatches returns the TextMatches field value if set, zero value otherwise.
func (o *CodeSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner {
	if o == nil || o.TextMatches == nil {
		var ret []SearchResultTextMatchesInner
		return ret
	}
	return o.TextMatches
}

// GetTextMatchesOk returns a tuple with the TextMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeSearchResultItem) GetTextMatchesOk() ([]SearchResultTextMatchesInner, bool) {
	if o == nil || o.TextMatches == nil {
		return nil, false
	}
	return o.TextMatches, true
}

// HasTextMatches returns a boolean if a field has been set.
func (o *CodeSearchResultItem) HasTextMatches() bool {
	if o != nil && o.TextMatches != nil {
		return true
	}

	return false
}

// SetTextMatches gets a reference to the given []SearchResultTextMatchesInner and assigns it to the TextMatches field.
func (o *CodeSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner) {
	o.TextMatches = v
}

func (o CodeSearchResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["git_url"] = o.GitUrl
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["score"] = o.Score
	}
	if o.FileSize != nil {
		toSerialize["file_size"] = o.FileSize
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.LastModifiedAt != nil {
		toSerialize["last_modified_at"] = o.LastModifiedAt
	}
	if o.LineNumbers != nil {
		toSerialize["line_numbers"] = o.LineNumbers
	}
	if o.TextMatches != nil {
		toSerialize["text_matches"] = o.TextMatches
	}
	return json.Marshal(toSerialize)
}

type NullableCodeSearchResultItem struct {
	value *CodeSearchResultItem
	isSet bool
}

func (v NullableCodeSearchResultItem) Get() *CodeSearchResultItem {
	return v.value
}

func (v *NullableCodeSearchResultItem) Set(val *CodeSearchResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeSearchResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeSearchResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeSearchResultItem(val *CodeSearchResultItem) *NullableCodeSearchResultItem {
	return &NullableCodeSearchResultItem{value: val, isSet: true}
}

func (v NullableCodeSearchResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeSearchResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


