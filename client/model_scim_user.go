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

// ScimUser SCIM /Users provisioning endpoints
type ScimUser struct {
	// SCIM schema used.
	Schemas []string `json:"schemas"`
	// Unique identifier of an external identity
	Id string `json:"id"`
	// The ID of the User.
	ExternalId NullableString `json:"externalId"`
	// Configured by the admin. Could be an email, login, or username
	UserName NullableString `json:"userName"`
	// The name of the user, suitable for display to end-users
	DisplayName NullableString `json:"displayName,omitempty"`
	Name ScimUserName `json:"name"`
	// user emails
	Emails []ScimUserEmailsInner `json:"emails"`
	// The active status of the User.
	Active bool `json:"active"`
	Meta ScimUserMeta `json:"meta"`
	// The ID of the organization.
	OrganizationId *int32 `json:"organization_id,omitempty"`
	// Set of operations to be performed
	Operations []ScimUserOperationsInner `json:"operations,omitempty"`
	// associated groups
	Groups []ScimUserGroupsInner `json:"groups,omitempty"`
}

// NewScimUser instantiates a new ScimUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimUser(schemas []string, id string, externalId NullableString, userName NullableString, name ScimUserName, emails []ScimUserEmailsInner, active bool, meta ScimUserMeta) *ScimUser {
	this := ScimUser{}
	this.Schemas = schemas
	this.Id = id
	this.ExternalId = externalId
	this.UserName = userName
	this.Name = name
	this.Emails = emails
	this.Active = active
	this.Meta = meta
	return &this
}

// NewScimUserWithDefaults instantiates a new ScimUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimUserWithDefaults() *ScimUser {
	this := ScimUser{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ScimUser) GetSchemas() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ScimUser) GetSchemasOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ScimUser) SetSchemas(v []string) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ScimUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScimUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScimUser) SetId(v string) {
	o.Id = v
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ScimUser) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScimUser) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *ScimUser) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetUserName returns the UserName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ScimUser) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScimUser) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// SetUserName sets field value
func (o *ScimUser) SetUserName(v string) {
	o.UserName.Set(&v)
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScimUser) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScimUser) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScimUser) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ScimUser) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ScimUser) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ScimUser) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetName returns the Name field value
func (o *ScimUser) GetName() ScimUserName {
	if o == nil {
		var ret ScimUserName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScimUser) GetNameOk() (*ScimUserName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScimUser) SetName(v ScimUserName) {
	o.Name = v
}

// GetEmails returns the Emails field value
func (o *ScimUser) GetEmails() []ScimUserEmailsInner {
	if o == nil {
		var ret []ScimUserEmailsInner
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ScimUser) GetEmailsOk() ([]ScimUserEmailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *ScimUser) SetEmails(v []ScimUserEmailsInner) {
	o.Emails = v
}

// GetActive returns the Active field value
func (o *ScimUser) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ScimUser) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ScimUser) SetActive(v bool) {
	o.Active = v
}

// GetMeta returns the Meta field value
func (o *ScimUser) GetMeta() ScimUserMeta {
	if o == nil {
		var ret ScimUserMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ScimUser) GetMetaOk() (*ScimUserMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ScimUser) SetMeta(v ScimUserMeta) {
	o.Meta = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ScimUser) GetOrganizationId() int32 {
	if o == nil || o.OrganizationId == nil {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ScimUser) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *ScimUser) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *ScimUser) GetOperations() []ScimUserOperationsInner {
	if o == nil || o.Operations == nil {
		var ret []ScimUserOperationsInner
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetOperationsOk() ([]ScimUserOperationsInner, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ScimUser) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []ScimUserOperationsInner and assigns it to the Operations field.
func (o *ScimUser) SetOperations(v []ScimUserOperationsInner) {
	o.Operations = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ScimUser) GetGroups() []ScimUserGroupsInner {
	if o == nil || o.Groups == nil {
		var ret []ScimUserGroupsInner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetGroupsOk() ([]ScimUserGroupsInner, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ScimUser) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []ScimUserGroupsInner and assigns it to the Groups field.
func (o *ScimUser) SetGroups(v []ScimUserGroupsInner) {
	o.Groups = v
}

func (o ScimUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if true {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableScimUser struct {
	value *ScimUser
	isSet bool
}

func (v NullableScimUser) Get() *ScimUser {
	return v.value
}

func (v *NullableScimUser) Set(val *ScimUser) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUser) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUser(val *ScimUser) *NullableScimUser {
	return &NullableScimUser{value: val, isSet: true}
}

func (v NullableScimUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


