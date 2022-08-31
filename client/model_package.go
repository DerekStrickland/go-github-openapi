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

// ModelPackage A software package
type ModelPackage struct {
	// Unique identifier of the package.
	Id int32 `json:"id"`
	// The name of the package.
	Name string `json:"name"`
	PackageType string `json:"package_type"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	// The number of versions of the package.
	VersionCount int32 `json:"version_count"`
	Visibility string `json:"visibility"`
	Owner NullableNullableSimpleUser `json:"owner,omitempty"`
	Repository NullableNullableMinimalRepository `json:"repository,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewModelPackage instantiates a new ModelPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPackage(id int32, name string, packageType string, url string, htmlUrl string, versionCount int32, visibility string, createdAt time.Time, updatedAt time.Time) *ModelPackage {
	this := ModelPackage{}
	this.Id = id
	this.Name = name
	this.PackageType = packageType
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.VersionCount = versionCount
	this.Visibility = visibility
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelPackageWithDefaults instantiates a new ModelPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPackageWithDefaults() *ModelPackage {
	this := ModelPackage{}
	return &this
}

// GetId returns the Id field value
func (o *ModelPackage) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelPackage) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ModelPackage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelPackage) SetName(v string) {
	o.Name = v
}

// GetPackageType returns the PackageType field value
func (o *ModelPackage) GetPackageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetPackageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageType, true
}

// SetPackageType sets field value
func (o *ModelPackage) SetPackageType(v string) {
	o.PackageType = v
}

// GetUrl returns the Url field value
func (o *ModelPackage) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ModelPackage) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *ModelPackage) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *ModelPackage) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetVersionCount returns the VersionCount field value
func (o *ModelPackage) GetVersionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionCount
}

// GetVersionCountOk returns a tuple with the VersionCount field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetVersionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionCount, true
}

// SetVersionCount sets field value
func (o *ModelPackage) SetVersionCount(v int32) {
	o.VersionCount = v
}

// GetVisibility returns the Visibility field value
func (o *ModelPackage) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *ModelPackage) SetVisibility(v string) {
	o.Visibility = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPackage) GetOwner() NullableSimpleUser {
	if o == nil || o.Owner.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPackage) GetOwnerOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelPackage) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableNullableSimpleUser and assigns it to the Owner field.
func (o *ModelPackage) SetOwner(v NullableSimpleUser) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ModelPackage) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ModelPackage) UnsetOwner() {
	o.Owner.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPackage) GetRepository() NullableMinimalRepository {
	if o == nil || o.Repository.Get() == nil {
		var ret NullableMinimalRepository
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPackage) GetRepositoryOk() (*NullableMinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *ModelPackage) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableNullableMinimalRepository and assigns it to the Repository field.
func (o *ModelPackage) SetRepository(v NullableMinimalRepository) {
	o.Repository.Set(&v)
}
// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *ModelPackage) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *ModelPackage) UnsetRepository() {
	o.Repository.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelPackage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelPackage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelPackage) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelPackage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ModelPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["package_type"] = o.PackageType
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["version_count"] = o.VersionCount
	}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelPackage struct {
	value *ModelPackage
	isSet bool
}

func (v NullableModelPackage) Get() *ModelPackage {
	return v.value
}

func (v *NullableModelPackage) Set(val *ModelPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPackage(val *ModelPackage) *NullableModelPackage {
	return &NullableModelPackage{value: val, isSet: true}
}

func (v NullableModelPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


