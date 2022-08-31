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

// NullableIntegration GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
type NullableIntegration struct {
	// Unique identifier of the GitHub app
	Id int32 `json:"id"`
	// The slug name of the GitHub app
	Slug *string `json:"slug,omitempty"`
	NodeId string `json:"node_id"`
	Owner NullableNullableSimpleUser `json:"owner"`
	// The name of the GitHub app
	Name string `json:"name"`
	Description NullableString `json:"description"`
	ExternalUrl string `json:"external_url"`
	HtmlUrl string `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Permissions IntegrationPermissions `json:"permissions"`
	// The list of events for the GitHub app
	Events []string `json:"events"`
	// The number of installations associated with the GitHub app
	InstallationsCount *int32 `json:"installations_count,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	WebhookSecret NullableString `json:"webhook_secret,omitempty"`
	Pem *string `json:"pem,omitempty"`
}

// NewNullableIntegration instantiates a new NullableIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullableIntegration(id int32, nodeId string, owner NullableNullableSimpleUser, name string, description NullableString, externalUrl string, htmlUrl string, createdAt time.Time, updatedAt time.Time, permissions IntegrationPermissions, events []string) *NullableIntegration {
	this := NullableIntegration{}
	this.Id = id
	this.NodeId = nodeId
	this.Owner = owner
	this.Name = name
	this.Description = description
	this.ExternalUrl = externalUrl
	this.HtmlUrl = htmlUrl
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Permissions = permissions
	this.Events = events
	return &this
}

// NewNullableIntegrationWithDefaults instantiates a new NullableIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullableIntegrationWithDefaults() *NullableIntegration {
	this := NullableIntegration{}
	return &this
}

// GetId returns the Id field value
func (o *NullableIntegration) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NullableIntegration) SetId(v int32) {
	o.Id = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *NullableIntegration) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *NullableIntegration) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *NullableIntegration) SetSlug(v string) {
	o.Slug = &v
}

// GetNodeId returns the NodeId field value
func (o *NullableIntegration) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *NullableIntegration) SetNodeId(v string) {
	o.NodeId = v
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *NullableIntegration) GetOwner() NullableSimpleUser {
	if o == nil || o.Owner.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableIntegration) GetOwnerOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// SetOwner sets field value
func (o *NullableIntegration) SetOwner(v NullableSimpleUser) {
	o.Owner.Set(&v)
}

// GetName returns the Name field value
func (o *NullableIntegration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NullableIntegration) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NullableIntegration) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableIntegration) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *NullableIntegration) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetExternalUrl returns the ExternalUrl field value
func (o *NullableIntegration) GetExternalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetExternalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrl, true
}

// SetExternalUrl sets field value
func (o *NullableIntegration) SetExternalUrl(v string) {
	o.ExternalUrl = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *NullableIntegration) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *NullableIntegration) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NullableIntegration) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NullableIntegration) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NullableIntegration) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NullableIntegration) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetPermissions returns the Permissions field value
func (o *NullableIntegration) GetPermissions() IntegrationPermissions {
	if o == nil {
		var ret IntegrationPermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetPermissionsOk() (*IntegrationPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *NullableIntegration) SetPermissions(v IntegrationPermissions) {
	o.Permissions = v
}

// GetEvents returns the Events field value
func (o *NullableIntegration) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *NullableIntegration) SetEvents(v []string) {
	o.Events = v
}

// GetInstallationsCount returns the InstallationsCount field value if set, zero value otherwise.
func (o *NullableIntegration) GetInstallationsCount() int32 {
	if o == nil || o.InstallationsCount == nil {
		var ret int32
		return ret
	}
	return *o.InstallationsCount
}

// GetInstallationsCountOk returns a tuple with the InstallationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetInstallationsCountOk() (*int32, bool) {
	if o == nil || o.InstallationsCount == nil {
		return nil, false
	}
	return o.InstallationsCount, true
}

// HasInstallationsCount returns a boolean if a field has been set.
func (o *NullableIntegration) HasInstallationsCount() bool {
	if o != nil && o.InstallationsCount != nil {
		return true
	}

	return false
}

// SetInstallationsCount gets a reference to the given int32 and assigns it to the InstallationsCount field.
func (o *NullableIntegration) SetInstallationsCount(v int32) {
	o.InstallationsCount = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *NullableIntegration) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *NullableIntegration) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *NullableIntegration) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *NullableIntegration) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *NullableIntegration) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *NullableIntegration) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetWebhookSecret returns the WebhookSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullableIntegration) GetWebhookSecret() string {
	if o == nil || o.WebhookSecret.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebhookSecret.Get()
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableIntegration) GetWebhookSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookSecret.Get(), o.WebhookSecret.IsSet()
}

// HasWebhookSecret returns a boolean if a field has been set.
func (o *NullableIntegration) HasWebhookSecret() bool {
	if o != nil && o.WebhookSecret.IsSet() {
		return true
	}

	return false
}

// SetWebhookSecret gets a reference to the given NullableString and assigns it to the WebhookSecret field.
func (o *NullableIntegration) SetWebhookSecret(v string) {
	o.WebhookSecret.Set(&v)
}
// SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil
func (o *NullableIntegration) SetWebhookSecretNil() {
	o.WebhookSecret.Set(nil)
}

// UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
func (o *NullableIntegration) UnsetWebhookSecret() {
	o.WebhookSecret.Unset()
}

// GetPem returns the Pem field value if set, zero value otherwise.
func (o *NullableIntegration) GetPem() string {
	if o == nil || o.Pem == nil {
		var ret string
		return ret
	}
	return *o.Pem
}

// GetPemOk returns a tuple with the Pem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableIntegration) GetPemOk() (*string, bool) {
	if o == nil || o.Pem == nil {
		return nil, false
	}
	return o.Pem, true
}

// HasPem returns a boolean if a field has been set.
func (o *NullableIntegration) HasPem() bool {
	if o != nil && o.Pem != nil {
		return true
	}

	return false
}

// SetPem gets a reference to the given string and assigns it to the Pem field.
func (o *NullableIntegration) SetPem(v string) {
	o.Pem = &v
}

func (o NullableIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["owner"] = o.Owner.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["external_url"] = o.ExternalUrl
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if o.InstallationsCount != nil {
		toSerialize["installations_count"] = o.InstallationsCount
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.WebhookSecret.IsSet() {
		toSerialize["webhook_secret"] = o.WebhookSecret.Get()
	}
	if o.Pem != nil {
		toSerialize["pem"] = o.Pem
	}
	return json.Marshal(toSerialize)
}

type NullableNullableIntegration struct {
	value *NullableIntegration
	isSet bool
}

func (v NullableNullableIntegration) Get() *NullableIntegration {
	return v.value
}

func (v *NullableNullableIntegration) Set(val *NullableIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableNullableIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableNullableIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullableIntegration(val *NullableIntegration) *NullableNullableIntegration {
	return &NullableNullableIntegration{value: val, isSet: true}
}

func (v NullableNullableIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullableIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


