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

// OrgMembership Org Membership
type OrgMembership struct {
	Url string `json:"url"`
	// The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	State string `json:"state"`
	// The user's membership type in the organization.
	Role string `json:"role"`
	OrganizationUrl string `json:"organization_url"`
	Organization OrganizationSimple `json:"organization"`
	User NullableNullableSimpleUser `json:"user"`
	Permissions *OrgMembershipPermissions `json:"permissions,omitempty"`
}

// NewOrgMembership instantiates a new OrgMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgMembership(url string, state string, role string, organizationUrl string, organization OrganizationSimple, user NullableNullableSimpleUser) *OrgMembership {
	this := OrgMembership{}
	this.Url = url
	this.State = state
	this.Role = role
	this.OrganizationUrl = organizationUrl
	this.Organization = organization
	this.User = user
	return &this
}

// NewOrgMembershipWithDefaults instantiates a new OrgMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgMembershipWithDefaults() *OrgMembership {
	this := OrgMembership{}
	return &this
}

// GetUrl returns the Url field value
func (o *OrgMembership) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrgMembership) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OrgMembership) SetUrl(v string) {
	o.Url = v
}

// GetState returns the State field value
func (o *OrgMembership) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OrgMembership) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *OrgMembership) SetState(v string) {
	o.State = v
}

// GetRole returns the Role field value
func (o *OrgMembership) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *OrgMembership) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *OrgMembership) SetRole(v string) {
	o.Role = v
}

// GetOrganizationUrl returns the OrganizationUrl field value
func (o *OrgMembership) GetOrganizationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationUrl
}

// GetOrganizationUrlOk returns a tuple with the OrganizationUrl field value
// and a boolean to check if the value has been set.
func (o *OrgMembership) GetOrganizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationUrl, true
}

// SetOrganizationUrl sets field value
func (o *OrgMembership) SetOrganizationUrl(v string) {
	o.OrganizationUrl = v
}

// GetOrganization returns the Organization field value
func (o *OrgMembership) GetOrganization() OrganizationSimple {
	if o == nil {
		var ret OrganizationSimple
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *OrgMembership) GetOrganizationOk() (*OrganizationSimple, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *OrgMembership) SetOrganization(v OrganizationSimple) {
	o.Organization = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *OrgMembership) GetUser() NullableSimpleUser {
	if o == nil || o.User.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgMembership) GetUserOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *OrgMembership) SetUser(v NullableSimpleUser) {
	o.User.Set(&v)
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *OrgMembership) GetPermissions() OrgMembershipPermissions {
	if o == nil || o.Permissions == nil {
		var ret OrgMembershipPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMembership) GetPermissionsOk() (*OrgMembershipPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *OrgMembership) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given OrgMembershipPermissions and assigns it to the Permissions field.
func (o *OrgMembership) SetPermissions(v OrgMembershipPermissions) {
	o.Permissions = &v
}

func (o OrgMembership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["organization_url"] = o.OrganizationUrl
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["user"] = o.User.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableOrgMembership struct {
	value *OrgMembership
	isSet bool
}

func (v NullableOrgMembership) Get() *OrgMembership {
	return v.value
}

func (v *NullableOrgMembership) Set(val *OrgMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgMembership(val *OrgMembership) *NullableOrgMembership {
	return &NullableOrgMembership{value: val, isSet: true}
}

func (v NullableOrgMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


