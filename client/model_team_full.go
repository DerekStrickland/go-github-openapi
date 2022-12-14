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

// TeamFull Groups of organization members that gives permissions on specified repositories.
type TeamFull struct {
	// Unique identifier of the team
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	// URL for the team
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	// Name of the team
	Name string `json:"name"`
	Slug string `json:"slug"`
	Description NullableString `json:"description"`
	// The level of privacy this team should have
	Privacy *string `json:"privacy,omitempty"`
	// Permission that the team will have for its repositories
	Permission string `json:"permission"`
	MembersUrl string `json:"members_url"`
	RepositoriesUrl string `json:"repositories_url"`
	Parent NullableNullableTeamSimple `json:"parent,omitempty"`
	MembersCount int32 `json:"members_count"`
	ReposCount int32 `json:"repos_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Organization TeamOrganization `json:"organization"`
	// Distinguished Name (DN) that team maps to within LDAP environment
	LdapDn *string `json:"ldap_dn,omitempty"`
}

// NewTeamFull instantiates a new TeamFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamFull(id int32, nodeId string, url string, htmlUrl string, name string, slug string, description NullableString, permission string, membersUrl string, repositoriesUrl string, membersCount int32, reposCount int32, createdAt time.Time, updatedAt time.Time, organization TeamOrganization) *TeamFull {
	this := TeamFull{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.Name = name
	this.Slug = slug
	this.Description = description
	this.Permission = permission
	this.MembersUrl = membersUrl
	this.RepositoriesUrl = repositoriesUrl
	this.MembersCount = membersCount
	this.ReposCount = reposCount
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Organization = organization
	return &this
}

// NewTeamFullWithDefaults instantiates a new TeamFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamFullWithDefaults() *TeamFull {
	this := TeamFull{}
	return &this
}

// GetId returns the Id field value
func (o *TeamFull) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamFull) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *TeamFull) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *TeamFull) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *TeamFull) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TeamFull) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *TeamFull) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *TeamFull) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetName returns the Name field value
func (o *TeamFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamFull) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *TeamFull) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *TeamFull) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TeamFull) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamFull) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *TeamFull) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *TeamFull) GetPrivacy() string {
	if o == nil || o.Privacy == nil {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamFull) GetPrivacyOk() (*string, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *TeamFull) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *TeamFull) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetPermission returns the Permission field value
func (o *TeamFull) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *TeamFull) SetPermission(v string) {
	o.Permission = v
}

// GetMembersUrl returns the MembersUrl field value
func (o *TeamFull) GetMembersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MembersUrl
}

// GetMembersUrlOk returns a tuple with the MembersUrl field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetMembersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MembersUrl, true
}

// SetMembersUrl sets field value
func (o *TeamFull) SetMembersUrl(v string) {
	o.MembersUrl = v
}

// GetRepositoriesUrl returns the RepositoriesUrl field value
func (o *TeamFull) GetRepositoriesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoriesUrl
}

// GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetRepositoriesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoriesUrl, true
}

// SetRepositoriesUrl sets field value
func (o *TeamFull) SetRepositoriesUrl(v string) {
	o.RepositoriesUrl = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamFull) GetParent() NullableTeamSimple {
	if o == nil || o.Parent.Get() == nil {
		var ret NullableTeamSimple
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamFull) GetParentOk() (*NullableTeamSimple, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *TeamFull) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableNullableTeamSimple and assigns it to the Parent field.
func (o *TeamFull) SetParent(v NullableTeamSimple) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *TeamFull) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *TeamFull) UnsetParent() {
	o.Parent.Unset()
}

// GetMembersCount returns the MembersCount field value
func (o *TeamFull) GetMembersCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MembersCount
}

// GetMembersCountOk returns a tuple with the MembersCount field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetMembersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MembersCount, true
}

// SetMembersCount sets field value
func (o *TeamFull) SetMembersCount(v int32) {
	o.MembersCount = v
}

// GetReposCount returns the ReposCount field value
func (o *TeamFull) GetReposCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReposCount
}

// GetReposCountOk returns a tuple with the ReposCount field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetReposCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReposCount, true
}

// SetReposCount sets field value
func (o *TeamFull) SetReposCount(v int32) {
	o.ReposCount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TeamFull) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TeamFull) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TeamFull) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TeamFull) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetOrganization returns the Organization field value
func (o *TeamFull) GetOrganization() TeamOrganization {
	if o == nil {
		var ret TeamOrganization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *TeamFull) GetOrganizationOk() (*TeamOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *TeamFull) SetOrganization(v TeamOrganization) {
	o.Organization = v
}

// GetLdapDn returns the LdapDn field value if set, zero value otherwise.
func (o *TeamFull) GetLdapDn() string {
	if o == nil || o.LdapDn == nil {
		var ret string
		return ret
	}
	return *o.LdapDn
}

// GetLdapDnOk returns a tuple with the LdapDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamFull) GetLdapDnOk() (*string, bool) {
	if o == nil || o.LdapDn == nil {
		return nil, false
	}
	return o.LdapDn, true
}

// HasLdapDn returns a boolean if a field has been set.
func (o *TeamFull) HasLdapDn() bool {
	if o != nil && o.LdapDn != nil {
		return true
	}

	return false
}

// SetLdapDn gets a reference to the given string and assigns it to the LdapDn field.
func (o *TeamFull) SetLdapDn(v string) {
	o.LdapDn = &v
}

func (o TeamFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if true {
		toSerialize["members_url"] = o.MembersUrl
	}
	if true {
		toSerialize["repositories_url"] = o.RepositoriesUrl
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if true {
		toSerialize["members_count"] = o.MembersCount
	}
	if true {
		toSerialize["repos_count"] = o.ReposCount
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if o.LdapDn != nil {
		toSerialize["ldap_dn"] = o.LdapDn
	}
	return json.Marshal(toSerialize)
}

type NullableTeamFull struct {
	value *TeamFull
	isSet bool
}

func (v NullableTeamFull) Get() *TeamFull {
	return v.value
}

func (v *NullableTeamFull) Set(val *TeamFull) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamFull) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamFull(val *TeamFull) *NullableTeamFull {
	return &NullableTeamFull{value: val, isSet: true}
}

func (v NullableTeamFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


