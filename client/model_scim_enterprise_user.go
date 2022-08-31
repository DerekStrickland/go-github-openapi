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

// ScimEnterpriseUser struct for ScimEnterpriseUser
type ScimEnterpriseUser struct {
	Schemas []string `json:"schemas"`
	Id string `json:"id"`
	ExternalId *string `json:"externalId,omitempty"`
	UserName *string `json:"userName,omitempty"`
	Name *ScimUserListEnterpriseResourcesInnerName `json:"name,omitempty"`
	Emails []ScimEnterpriseUserEmailsInner `json:"emails,omitempty"`
	Groups []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner `json:"groups,omitempty"`
	Active *bool `json:"active,omitempty"`
	Meta *ScimGroupListEnterpriseResourcesInnerMeta `json:"meta,omitempty"`
}

// NewScimEnterpriseUser instantiates a new ScimEnterpriseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimEnterpriseUser(schemas []string, id string) *ScimEnterpriseUser {
	this := ScimEnterpriseUser{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewScimEnterpriseUserWithDefaults instantiates a new ScimEnterpriseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimEnterpriseUserWithDefaults() *ScimEnterpriseUser {
	this := ScimEnterpriseUser{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ScimEnterpriseUser) GetSchemas() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetSchemasOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ScimEnterpriseUser) SetSchemas(v []string) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ScimEnterpriseUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScimEnterpriseUser) SetId(v string) {
	o.Id = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ScimEnterpriseUser) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ScimEnterpriseUser) SetUserName(v string) {
	o.UserName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetName() ScimUserListEnterpriseResourcesInnerName {
	if o == nil || o.Name == nil {
		var ret ScimUserListEnterpriseResourcesInnerName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetNameOk() (*ScimUserListEnterpriseResourcesInnerName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given ScimUserListEnterpriseResourcesInnerName and assigns it to the Name field.
func (o *ScimEnterpriseUser) SetName(v ScimUserListEnterpriseResourcesInnerName) {
	o.Name = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetEmails() []ScimEnterpriseUserEmailsInner {
	if o == nil || o.Emails == nil {
		var ret []ScimEnterpriseUserEmailsInner
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetEmailsOk() ([]ScimEnterpriseUserEmailsInner, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []ScimEnterpriseUserEmailsInner and assigns it to the Emails field.
func (o *ScimEnterpriseUser) SetEmails(v []ScimEnterpriseUserEmailsInner) {
	o.Emails = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetGroups() []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner {
	if o == nil || o.Groups == nil {
		var ret []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetGroupsOk() ([]EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner and assigns it to the Groups field.
func (o *ScimEnterpriseUser) SetGroups(v []EnterpriseAdminProvisionAndInviteEnterpriseUserRequestGroupsInner) {
	o.Groups = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ScimEnterpriseUser) SetActive(v bool) {
	o.Active = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimEnterpriseUser) GetMeta() ScimGroupListEnterpriseResourcesInnerMeta {
	if o == nil || o.Meta == nil {
		var ret ScimGroupListEnterpriseResourcesInnerMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimEnterpriseUser) GetMetaOk() (*ScimGroupListEnterpriseResourcesInnerMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimEnterpriseUser) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ScimGroupListEnterpriseResourcesInnerMeta and assigns it to the Meta field.
func (o *ScimEnterpriseUser) SetMeta(v ScimGroupListEnterpriseResourcesInnerMeta) {
	o.Meta = &v
}

func (o ScimEnterpriseUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableScimEnterpriseUser struct {
	value *ScimEnterpriseUser
	isSet bool
}

func (v NullableScimEnterpriseUser) Get() *ScimEnterpriseUser {
	return v.value
}

func (v *NullableScimEnterpriseUser) Set(val *ScimEnterpriseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableScimEnterpriseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableScimEnterpriseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimEnterpriseUser(val *ScimEnterpriseUser) *NullableScimEnterpriseUser {
	return &NullableScimEnterpriseUser{value: val, isSet: true}
}

func (v NullableScimEnterpriseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimEnterpriseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


