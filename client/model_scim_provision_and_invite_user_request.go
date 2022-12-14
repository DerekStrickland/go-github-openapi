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

// ScimProvisionAndInviteUserRequest struct for ScimProvisionAndInviteUserRequest
type ScimProvisionAndInviteUserRequest struct {
	// Configured by the admin. Could be an email, login, or username
	UserName string `json:"userName"`
	// The name of the user, suitable for display to end-users
	DisplayName *string `json:"displayName,omitempty"`
	Name ScimProvisionAndInviteUserRequestName `json:"name"`
	// user emails
	Emails []ScimProvisionAndInviteUserRequestEmailsInner `json:"emails"`
	Schemas []string `json:"schemas,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewScimProvisionAndInviteUserRequest instantiates a new ScimProvisionAndInviteUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimProvisionAndInviteUserRequest(userName string, name ScimProvisionAndInviteUserRequestName, emails []ScimProvisionAndInviteUserRequestEmailsInner) *ScimProvisionAndInviteUserRequest {
	this := ScimProvisionAndInviteUserRequest{}
	this.UserName = userName
	this.Name = name
	this.Emails = emails
	return &this
}

// NewScimProvisionAndInviteUserRequestWithDefaults instantiates a new ScimProvisionAndInviteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimProvisionAndInviteUserRequestWithDefaults() *ScimProvisionAndInviteUserRequest {
	this := ScimProvisionAndInviteUserRequest{}
	return &this
}

// GetUserName returns the UserName field value
func (o *ScimProvisionAndInviteUserRequest) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *ScimProvisionAndInviteUserRequest) SetUserName(v string) {
	o.UserName = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ScimProvisionAndInviteUserRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScimProvisionAndInviteUserRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ScimProvisionAndInviteUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value
func (o *ScimProvisionAndInviteUserRequest) GetName() ScimProvisionAndInviteUserRequestName {
	if o == nil {
		var ret ScimProvisionAndInviteUserRequestName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetNameOk() (*ScimProvisionAndInviteUserRequestName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScimProvisionAndInviteUserRequest) SetName(v ScimProvisionAndInviteUserRequestName) {
	o.Name = v
}

// GetEmails returns the Emails field value
func (o *ScimProvisionAndInviteUserRequest) GetEmails() []ScimProvisionAndInviteUserRequestEmailsInner {
	if o == nil {
		var ret []ScimProvisionAndInviteUserRequestEmailsInner
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetEmailsOk() ([]ScimProvisionAndInviteUserRequestEmailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *ScimProvisionAndInviteUserRequest) SetEmails(v []ScimProvisionAndInviteUserRequestEmailsInner) {
	o.Emails = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimProvisionAndInviteUserRequest) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetSchemasOk() ([]string, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimProvisionAndInviteUserRequest) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimProvisionAndInviteUserRequest) SetSchemas(v []string) {
	o.Schemas = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ScimProvisionAndInviteUserRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ScimProvisionAndInviteUserRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ScimProvisionAndInviteUserRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ScimProvisionAndInviteUserRequest) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ScimProvisionAndInviteUserRequest) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ScimProvisionAndInviteUserRequest) SetGroups(v []string) {
	o.Groups = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ScimProvisionAndInviteUserRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimProvisionAndInviteUserRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ScimProvisionAndInviteUserRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ScimProvisionAndInviteUserRequest) SetActive(v bool) {
	o.Active = &v
}

func (o ScimProvisionAndInviteUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["userName"] = o.UserName
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableScimProvisionAndInviteUserRequest struct {
	value *ScimProvisionAndInviteUserRequest
	isSet bool
}

func (v NullableScimProvisionAndInviteUserRequest) Get() *ScimProvisionAndInviteUserRequest {
	return v.value
}

func (v *NullableScimProvisionAndInviteUserRequest) Set(val *ScimProvisionAndInviteUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScimProvisionAndInviteUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScimProvisionAndInviteUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimProvisionAndInviteUserRequest(val *ScimProvisionAndInviteUserRequest) *NullableScimProvisionAndInviteUserRequest {
	return &NullableScimProvisionAndInviteUserRequest{value: val, isSet: true}
}

func (v NullableScimProvisionAndInviteUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimProvisionAndInviteUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


