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

// BranchRestrictionPolicyTeamsInner struct for BranchRestrictionPolicyTeamsInner
type BranchRestrictionPolicyTeamsInner struct {
	Id *int32 `json:"id,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	Url *string `json:"url,omitempty"`
	HtmlUrl *string `json:"html_url,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Privacy *string `json:"privacy,omitempty"`
	Permission *string `json:"permission,omitempty"`
	MembersUrl *string `json:"members_url,omitempty"`
	RepositoriesUrl *string `json:"repositories_url,omitempty"`
	Parent NullableString `json:"parent,omitempty"`
}

// NewBranchRestrictionPolicyTeamsInner instantiates a new BranchRestrictionPolicyTeamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchRestrictionPolicyTeamsInner() *BranchRestrictionPolicyTeamsInner {
	this := BranchRestrictionPolicyTeamsInner{}
	return &this
}

// NewBranchRestrictionPolicyTeamsInnerWithDefaults instantiates a new BranchRestrictionPolicyTeamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchRestrictionPolicyTeamsInnerWithDefaults() *BranchRestrictionPolicyTeamsInner {
	this := BranchRestrictionPolicyTeamsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BranchRestrictionPolicyTeamsInner) SetId(v int32) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BranchRestrictionPolicyTeamsInner) SetNodeId(v string) {
	o.NodeId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BranchRestrictionPolicyTeamsInner) SetUrl(v string) {
	o.Url = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetHtmlUrlOk() (*string, bool) {
	if o == nil || o.HtmlUrl == nil {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl != nil {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *BranchRestrictionPolicyTeamsInner) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BranchRestrictionPolicyTeamsInner) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *BranchRestrictionPolicyTeamsInner) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BranchRestrictionPolicyTeamsInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BranchRestrictionPolicyTeamsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *BranchRestrictionPolicyTeamsInner) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *BranchRestrictionPolicyTeamsInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *BranchRestrictionPolicyTeamsInner) UnsetDescription() {
	o.Description.Unset()
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetPrivacy() string {
	if o == nil || o.Privacy == nil {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetPrivacyOk() (*string, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *BranchRestrictionPolicyTeamsInner) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *BranchRestrictionPolicyTeamsInner) SetPermission(v string) {
	o.Permission = &v
}

// GetMembersUrl returns the MembersUrl field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetMembersUrl() string {
	if o == nil || o.MembersUrl == nil {
		var ret string
		return ret
	}
	return *o.MembersUrl
}

// GetMembersUrlOk returns a tuple with the MembersUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetMembersUrlOk() (*string, bool) {
	if o == nil || o.MembersUrl == nil {
		return nil, false
	}
	return o.MembersUrl, true
}

// HasMembersUrl returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasMembersUrl() bool {
	if o != nil && o.MembersUrl != nil {
		return true
	}

	return false
}

// SetMembersUrl gets a reference to the given string and assigns it to the MembersUrl field.
func (o *BranchRestrictionPolicyTeamsInner) SetMembersUrl(v string) {
	o.MembersUrl = &v
}

// GetRepositoriesUrl returns the RepositoriesUrl field value if set, zero value otherwise.
func (o *BranchRestrictionPolicyTeamsInner) GetRepositoriesUrl() string {
	if o == nil || o.RepositoriesUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoriesUrl
}

// GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchRestrictionPolicyTeamsInner) GetRepositoriesUrlOk() (*string, bool) {
	if o == nil || o.RepositoriesUrl == nil {
		return nil, false
	}
	return o.RepositoriesUrl, true
}

// HasRepositoriesUrl returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasRepositoriesUrl() bool {
	if o != nil && o.RepositoriesUrl != nil {
		return true
	}

	return false
}

// SetRepositoriesUrl gets a reference to the given string and assigns it to the RepositoriesUrl field.
func (o *BranchRestrictionPolicyTeamsInner) SetRepositoriesUrl(v string) {
	o.RepositoriesUrl = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BranchRestrictionPolicyTeamsInner) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BranchRestrictionPolicyTeamsInner) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *BranchRestrictionPolicyTeamsInner) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *BranchRestrictionPolicyTeamsInner) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *BranchRestrictionPolicyTeamsInner) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *BranchRestrictionPolicyTeamsInner) UnsetParent() {
	o.Parent.Unset()
}

func (o BranchRestrictionPolicyTeamsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.HtmlUrl != nil {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.MembersUrl != nil {
		toSerialize["members_url"] = o.MembersUrl
	}
	if o.RepositoriesUrl != nil {
		toSerialize["repositories_url"] = o.RepositoriesUrl
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBranchRestrictionPolicyTeamsInner struct {
	value *BranchRestrictionPolicyTeamsInner
	isSet bool
}

func (v NullableBranchRestrictionPolicyTeamsInner) Get() *BranchRestrictionPolicyTeamsInner {
	return v.value
}

func (v *NullableBranchRestrictionPolicyTeamsInner) Set(val *BranchRestrictionPolicyTeamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchRestrictionPolicyTeamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchRestrictionPolicyTeamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchRestrictionPolicyTeamsInner(val *BranchRestrictionPolicyTeamsInner) *NullableBranchRestrictionPolicyTeamsInner {
	return &NullableBranchRestrictionPolicyTeamsInner{value: val, isSet: true}
}

func (v NullableBranchRestrictionPolicyTeamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchRestrictionPolicyTeamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


