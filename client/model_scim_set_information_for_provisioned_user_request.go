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

// ScimSetInformationForProvisionedUserRequest struct for ScimSetInformationForProvisionedUserRequest
type ScimSetInformationForProvisionedUserRequest struct {
	Schemas []string `json:"schemas,omitempty"`
	// The name of the user, suitable for display to end-users
	DisplayName *string `json:"displayName,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Active *bool `json:"active,omitempty"`
	// Configured by the admin. Could be an email, login, or username
	UserName string `json:"userName"`
	Name ScimProvisionAndInviteUserRequestName `json:"name"`
	// user emails
	Emails []ScimSetInformationForProvisionedUserRequestEmailsInner `json:"emails"`
}

// NewScimSetInformationForProvisionedUserRequest instantiates a new ScimSetInformationForProvisionedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimSetInformationForProvisionedUserRequest(userName string, name ScimProvisionAndInviteUserRequestName, emails []ScimSetInformationForProvisionedUserRequestEmailsInner) *ScimSetInformationForProvisionedUserRequest {
	this := ScimSetInformationForProvisionedUserRequest{}
	this.UserName = userName
	this.Name = name
	this.Emails = emails
	return &this
}

// NewScimSetInformationForProvisionedUserRequestWithDefaults instantiates a new ScimSetInformationForProvisionedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimSetInformationForProvisionedUserRequestWithDefaults() *ScimSetInformationForProvisionedUserRequest {
	this := ScimSetInformationForProvisionedUserRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimSetInformationForProvisionedUserRequest) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetSchemasOk() ([]string, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimSetInformationForProvisionedUserRequest) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimSetInformationForProvisionedUserRequest) SetSchemas(v []string) {
	o.Schemas = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ScimSetInformationForProvisionedUserRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScimSetInformationForProvisionedUserRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ScimSetInformationForProvisionedUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ScimSetInformationForProvisionedUserRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ScimSetInformationForProvisionedUserRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ScimSetInformationForProvisionedUserRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ScimSetInformationForProvisionedUserRequest) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ScimSetInformationForProvisionedUserRequest) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ScimSetInformationForProvisionedUserRequest) SetGroups(v []string) {
	o.Groups = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ScimSetInformationForProvisionedUserRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ScimSetInformationForProvisionedUserRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ScimSetInformationForProvisionedUserRequest) SetActive(v bool) {
	o.Active = &v
}

// GetUserName returns the UserName field value
func (o *ScimSetInformationForProvisionedUserRequest) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *ScimSetInformationForProvisionedUserRequest) SetUserName(v string) {
	o.UserName = v
}

// GetName returns the Name field value
func (o *ScimSetInformationForProvisionedUserRequest) GetName() ScimProvisionAndInviteUserRequestName {
	if o == nil {
		var ret ScimProvisionAndInviteUserRequestName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetNameOk() (*ScimProvisionAndInviteUserRequestName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScimSetInformationForProvisionedUserRequest) SetName(v ScimProvisionAndInviteUserRequestName) {
	o.Name = v
}

// GetEmails returns the Emails field value
func (o *ScimSetInformationForProvisionedUserRequest) GetEmails() []ScimSetInformationForProvisionedUserRequestEmailsInner {
	if o == nil {
		var ret []ScimSetInformationForProvisionedUserRequestEmailsInner
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ScimSetInformationForProvisionedUserRequest) GetEmailsOk() ([]ScimSetInformationForProvisionedUserRequestEmailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *ScimSetInformationForProvisionedUserRequest) SetEmails(v []ScimSetInformationForProvisionedUserRequestEmailsInner) {
	o.Emails = v
}

func (o ScimSetInformationForProvisionedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
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
	if true {
		toSerialize["userName"] = o.UserName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	return json.Marshal(toSerialize)
}

type NullableScimSetInformationForProvisionedUserRequest struct {
	value *ScimSetInformationForProvisionedUserRequest
	isSet bool
}

func (v NullableScimSetInformationForProvisionedUserRequest) Get() *ScimSetInformationForProvisionedUserRequest {
	return v.value
}

func (v *NullableScimSetInformationForProvisionedUserRequest) Set(val *ScimSetInformationForProvisionedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScimSetInformationForProvisionedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScimSetInformationForProvisionedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimSetInformationForProvisionedUserRequest(val *ScimSetInformationForProvisionedUserRequest) *NullableScimSetInformationForProvisionedUserRequest {
	return &NullableScimSetInformationForProvisionedUserRequest{value: val, isSet: true}
}

func (v NullableScimSetInformationForProvisionedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimSetInformationForProvisionedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

