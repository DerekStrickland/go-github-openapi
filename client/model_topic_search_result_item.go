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

// TopicSearchResultItem Topic Search Result Item
type TopicSearchResultItem struct {
	Name string `json:"name"`
	DisplayName NullableString `json:"display_name"`
	ShortDescription NullableString `json:"short_description"`
	Description NullableString `json:"description"`
	CreatedBy NullableString `json:"created_by"`
	Released NullableString `json:"released"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Featured bool `json:"featured"`
	Curated bool `json:"curated"`
	Score float32 `json:"score"`
	RepositoryCount NullableInt32 `json:"repository_count,omitempty"`
	LogoUrl NullableString `json:"logo_url,omitempty"`
	TextMatches []SearchResultTextMatchesInner `json:"text_matches,omitempty"`
	Related []TopicSearchResultItemRelatedInner `json:"related,omitempty"`
	Aliases []TopicSearchResultItemRelatedInner `json:"aliases,omitempty"`
}

// NewTopicSearchResultItem instantiates a new TopicSearchResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicSearchResultItem(name string, displayName NullableString, shortDescription NullableString, description NullableString, createdBy NullableString, released NullableString, createdAt time.Time, updatedAt time.Time, featured bool, curated bool, score float32) *TopicSearchResultItem {
	this := TopicSearchResultItem{}
	this.Name = name
	this.DisplayName = displayName
	this.ShortDescription = shortDescription
	this.Description = description
	this.CreatedBy = createdBy
	this.Released = released
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Featured = featured
	this.Curated = curated
	this.Score = score
	return &this
}

// NewTopicSearchResultItemWithDefaults instantiates a new TopicSearchResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicSearchResultItemWithDefaults() *TopicSearchResultItem {
	this := TopicSearchResultItem{}
	return &this
}

// GetName returns the Name field value
func (o *TopicSearchResultItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TopicSearchResultItem) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TopicSearchResultItem) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// SetDisplayName sets field value
func (o *TopicSearchResultItem) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// GetShortDescription returns the ShortDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TopicSearchResultItem) GetShortDescription() string {
	if o == nil || o.ShortDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShortDescription.Get()
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortDescription.Get(), o.ShortDescription.IsSet()
}

// SetShortDescription sets field value
func (o *TopicSearchResultItem) SetShortDescription(v string) {
	o.ShortDescription.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TopicSearchResultItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *TopicSearchResultItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TopicSearchResultItem) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// SetCreatedBy sets field value
func (o *TopicSearchResultItem) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// GetReleased returns the Released field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TopicSearchResultItem) GetReleased() string {
	if o == nil || o.Released.Get() == nil {
		var ret string
		return ret
	}

	return *o.Released.Get()
}

// GetReleasedOk returns a tuple with the Released field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetReleasedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Released.Get(), o.Released.IsSet()
}

// SetReleased sets field value
func (o *TopicSearchResultItem) SetReleased(v string) {
	o.Released.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *TopicSearchResultItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TopicSearchResultItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TopicSearchResultItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TopicSearchResultItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetFeatured returns the Featured field value
func (o *TopicSearchResultItem) GetFeatured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetFeaturedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Featured, true
}

// SetFeatured sets field value
func (o *TopicSearchResultItem) SetFeatured(v bool) {
	o.Featured = v
}

// GetCurated returns the Curated field value
func (o *TopicSearchResultItem) GetCurated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Curated
}

// GetCuratedOk returns a tuple with the Curated field value
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetCuratedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Curated, true
}

// SetCurated sets field value
func (o *TopicSearchResultItem) SetCurated(v bool) {
	o.Curated = v
}

// GetScore returns the Score field value
func (o *TopicSearchResultItem) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *TopicSearchResultItem) SetScore(v float32) {
	o.Score = v
}

// GetRepositoryCount returns the RepositoryCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicSearchResultItem) GetRepositoryCount() int32 {
	if o == nil || o.RepositoryCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RepositoryCount.Get()
}

// GetRepositoryCountOk returns a tuple with the RepositoryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetRepositoryCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryCount.Get(), o.RepositoryCount.IsSet()
}

// HasRepositoryCount returns a boolean if a field has been set.
func (o *TopicSearchResultItem) HasRepositoryCount() bool {
	if o != nil && o.RepositoryCount.IsSet() {
		return true
	}

	return false
}

// SetRepositoryCount gets a reference to the given NullableInt32 and assigns it to the RepositoryCount field.
func (o *TopicSearchResultItem) SetRepositoryCount(v int32) {
	o.RepositoryCount.Set(&v)
}
// SetRepositoryCountNil sets the value for RepositoryCount to be an explicit nil
func (o *TopicSearchResultItem) SetRepositoryCountNil() {
	o.RepositoryCount.Set(nil)
}

// UnsetRepositoryCount ensures that no value is present for RepositoryCount, not even an explicit nil
func (o *TopicSearchResultItem) UnsetRepositoryCount() {
	o.RepositoryCount.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicSearchResultItem) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *TopicSearchResultItem) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *TopicSearchResultItem) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *TopicSearchResultItem) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *TopicSearchResultItem) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetTextMatches returns the TextMatches field value if set, zero value otherwise.
func (o *TopicSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner {
	if o == nil || o.TextMatches == nil {
		var ret []SearchResultTextMatchesInner
		return ret
	}
	return o.TextMatches
}

// GetTextMatchesOk returns a tuple with the TextMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicSearchResultItem) GetTextMatchesOk() ([]SearchResultTextMatchesInner, bool) {
	if o == nil || o.TextMatches == nil {
		return nil, false
	}
	return o.TextMatches, true
}

// HasTextMatches returns a boolean if a field has been set.
func (o *TopicSearchResultItem) HasTextMatches() bool {
	if o != nil && o.TextMatches != nil {
		return true
	}

	return false
}

// SetTextMatches gets a reference to the given []SearchResultTextMatchesInner and assigns it to the TextMatches field.
func (o *TopicSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner) {
	o.TextMatches = v
}

// GetRelated returns the Related field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicSearchResultItem) GetRelated() []TopicSearchResultItemRelatedInner {
	if o == nil {
		var ret []TopicSearchResultItemRelatedInner
		return ret
	}
	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetRelatedOk() ([]TopicSearchResultItemRelatedInner, bool) {
	if o == nil || o.Related == nil {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *TopicSearchResultItem) HasRelated() bool {
	if o != nil && o.Related != nil {
		return true
	}

	return false
}

// SetRelated gets a reference to the given []TopicSearchResultItemRelatedInner and assigns it to the Related field.
func (o *TopicSearchResultItem) SetRelated(v []TopicSearchResultItemRelatedInner) {
	o.Related = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopicSearchResultItem) GetAliases() []TopicSearchResultItemRelatedInner {
	if o == nil {
		var ret []TopicSearchResultItemRelatedInner
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopicSearchResultItem) GetAliasesOk() ([]TopicSearchResultItemRelatedInner, bool) {
	if o == nil || o.Aliases == nil {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *TopicSearchResultItem) HasAliases() bool {
	if o != nil && o.Aliases != nil {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []TopicSearchResultItemRelatedInner and assigns it to the Aliases field.
func (o *TopicSearchResultItem) SetAliases(v []TopicSearchResultItemRelatedInner) {
	o.Aliases = v
}

func (o TopicSearchResultItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if true {
		toSerialize["short_description"] = o.ShortDescription.Get()
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if true {
		toSerialize["released"] = o.Released.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["featured"] = o.Featured
	}
	if true {
		toSerialize["curated"] = o.Curated
	}
	if true {
		toSerialize["score"] = o.Score
	}
	if o.RepositoryCount.IsSet() {
		toSerialize["repository_count"] = o.RepositoryCount.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.TextMatches != nil {
		toSerialize["text_matches"] = o.TextMatches
	}
	if o.Related != nil {
		toSerialize["related"] = o.Related
	}
	if o.Aliases != nil {
		toSerialize["aliases"] = o.Aliases
	}
	return json.Marshal(toSerialize)
}

type NullableTopicSearchResultItem struct {
	value *TopicSearchResultItem
	isSet bool
}

func (v NullableTopicSearchResultItem) Get() *TopicSearchResultItem {
	return v.value
}

func (v *NullableTopicSearchResultItem) Set(val *TopicSearchResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicSearchResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicSearchResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicSearchResultItem(val *TopicSearchResultItem) *NullableTopicSearchResultItem {
	return &NullableTopicSearchResultItem{value: val, isSet: true}
}

func (v NullableTopicSearchResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicSearchResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


