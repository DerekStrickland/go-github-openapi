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

// OrganizationSecretScanningAlert struct for OrganizationSecretScanningAlert
type OrganizationSecretScanningAlert struct {
	// The security alert number.
	Number *int32 `json:"number,omitempty"`
	// The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// The REST API URL of the alert resource.
	Url *string `json:"url,omitempty"`
	// The GitHub URL of the alert resource.
	HtmlUrl *string `json:"html_url,omitempty"`
	// The REST API URL of the code locations for this alert.
	LocationsUrl *string `json:"locations_url,omitempty"`
	State *SecretScanningAlertState `json:"state,omitempty"`
	Resolution NullableSecretScanningAlertResolution `json:"resolution,omitempty"`
	// The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	ResolvedAt NullableTime `json:"resolved_at,omitempty"`
	ResolvedBy NullableNullableSimpleUser `json:"resolved_by,omitempty"`
	// The type of secret that secret scanning detected.
	SecretType *string `json:"secret_type,omitempty"`
	// User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see \"[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security).\"
	SecretTypeDisplayName *string `json:"secret_type_display_name,omitempty"`
	// The secret that was detected.
	Secret *string `json:"secret,omitempty"`
	Repository *SimpleRepository `json:"repository,omitempty"`
	// Whether push protection was bypassed for the detected secret.
	PushProtectionBypassed NullableBool `json:"push_protection_bypassed,omitempty"`
	PushProtectionBypassedBy NullableNullableSimpleUser `json:"push_protection_bypassed_by,omitempty"`
	// The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	PushProtectionBypassedAt NullableTime `json:"push_protection_bypassed_at,omitempty"`
}

// NewOrganizationSecretScanningAlert instantiates a new OrganizationSecretScanningAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSecretScanningAlert() *OrganizationSecretScanningAlert {
	this := OrganizationSecretScanningAlert{}
	return &this
}

// NewOrganizationSecretScanningAlertWithDefaults instantiates a new OrganizationSecretScanningAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSecretScanningAlertWithDefaults() *OrganizationSecretScanningAlert {
	this := OrganizationSecretScanningAlert{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *OrganizationSecretScanningAlert) SetNumber(v int32) {
	o.Number = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrganizationSecretScanningAlert) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *OrganizationSecretScanningAlert) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrganizationSecretScanningAlert) SetUrl(v string) {
	o.Url = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetHtmlUrlOk() (*string, bool) {
	if o == nil || o.HtmlUrl == nil {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl != nil {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *OrganizationSecretScanningAlert) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetLocationsUrl returns the LocationsUrl field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetLocationsUrl() string {
	if o == nil || o.LocationsUrl == nil {
		var ret string
		return ret
	}
	return *o.LocationsUrl
}

// GetLocationsUrlOk returns a tuple with the LocationsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetLocationsUrlOk() (*string, bool) {
	if o == nil || o.LocationsUrl == nil {
		return nil, false
	}
	return o.LocationsUrl, true
}

// HasLocationsUrl returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasLocationsUrl() bool {
	if o != nil && o.LocationsUrl != nil {
		return true
	}

	return false
}

// SetLocationsUrl gets a reference to the given string and assigns it to the LocationsUrl field.
func (o *OrganizationSecretScanningAlert) SetLocationsUrl(v string) {
	o.LocationsUrl = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetState() SecretScanningAlertState {
	if o == nil || o.State == nil {
		var ret SecretScanningAlertState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetStateOk() (*SecretScanningAlertState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given SecretScanningAlertState and assigns it to the State field.
func (o *OrganizationSecretScanningAlert) SetState(v SecretScanningAlertState) {
	o.State = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetResolution() SecretScanningAlertResolution {
	if o == nil || o.Resolution.Get() == nil {
		var ret SecretScanningAlertResolution
		return ret
	}
	return *o.Resolution.Get()
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetResolutionOk() (*SecretScanningAlertResolution, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolution.Get(), o.Resolution.IsSet()
}

// HasResolution returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasResolution() bool {
	if o != nil && o.Resolution.IsSet() {
		return true
	}

	return false
}

// SetResolution gets a reference to the given NullableSecretScanningAlertResolution and assigns it to the Resolution field.
func (o *OrganizationSecretScanningAlert) SetResolution(v SecretScanningAlertResolution) {
	o.Resolution.Set(&v)
}
// SetResolutionNil sets the value for Resolution to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetResolutionNil() {
	o.Resolution.Set(nil)
}

// UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetResolution() {
	o.Resolution.Unset()
}

// GetResolvedAt returns the ResolvedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetResolvedAt() time.Time {
	if o == nil || o.ResolvedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ResolvedAt.Get()
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetResolvedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedAt.Get(), o.ResolvedAt.IsSet()
}

// HasResolvedAt returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasResolvedAt() bool {
	if o != nil && o.ResolvedAt.IsSet() {
		return true
	}

	return false
}

// SetResolvedAt gets a reference to the given NullableTime and assigns it to the ResolvedAt field.
func (o *OrganizationSecretScanningAlert) SetResolvedAt(v time.Time) {
	o.ResolvedAt.Set(&v)
}
// SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetResolvedAtNil() {
	o.ResolvedAt.Set(nil)
}

// UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetResolvedAt() {
	o.ResolvedAt.Unset()
}

// GetResolvedBy returns the ResolvedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetResolvedBy() NullableSimpleUser {
	if o == nil || o.ResolvedBy.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.ResolvedBy.Get()
}

// GetResolvedByOk returns a tuple with the ResolvedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetResolvedByOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedBy.Get(), o.ResolvedBy.IsSet()
}

// HasResolvedBy returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasResolvedBy() bool {
	if o != nil && o.ResolvedBy.IsSet() {
		return true
	}

	return false
}

// SetResolvedBy gets a reference to the given NullableNullableSimpleUser and assigns it to the ResolvedBy field.
func (o *OrganizationSecretScanningAlert) SetResolvedBy(v NullableSimpleUser) {
	o.ResolvedBy.Set(&v)
}
// SetResolvedByNil sets the value for ResolvedBy to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetResolvedByNil() {
	o.ResolvedBy.Set(nil)
}

// UnsetResolvedBy ensures that no value is present for ResolvedBy, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetResolvedBy() {
	o.ResolvedBy.Unset()
}

// GetSecretType returns the SecretType field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetSecretType() string {
	if o == nil || o.SecretType == nil {
		var ret string
		return ret
	}
	return *o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetSecretTypeOk() (*string, bool) {
	if o == nil || o.SecretType == nil {
		return nil, false
	}
	return o.SecretType, true
}

// HasSecretType returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasSecretType() bool {
	if o != nil && o.SecretType != nil {
		return true
	}

	return false
}

// SetSecretType gets a reference to the given string and assigns it to the SecretType field.
func (o *OrganizationSecretScanningAlert) SetSecretType(v string) {
	o.SecretType = &v
}

// GetSecretTypeDisplayName returns the SecretTypeDisplayName field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetSecretTypeDisplayName() string {
	if o == nil || o.SecretTypeDisplayName == nil {
		var ret string
		return ret
	}
	return *o.SecretTypeDisplayName
}

// GetSecretTypeDisplayNameOk returns a tuple with the SecretTypeDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetSecretTypeDisplayNameOk() (*string, bool) {
	if o == nil || o.SecretTypeDisplayName == nil {
		return nil, false
	}
	return o.SecretTypeDisplayName, true
}

// HasSecretTypeDisplayName returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasSecretTypeDisplayName() bool {
	if o != nil && o.SecretTypeDisplayName != nil {
		return true
	}

	return false
}

// SetSecretTypeDisplayName gets a reference to the given string and assigns it to the SecretTypeDisplayName field.
func (o *OrganizationSecretScanningAlert) SetSecretTypeDisplayName(v string) {
	o.SecretTypeDisplayName = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *OrganizationSecretScanningAlert) SetSecret(v string) {
	o.Secret = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *OrganizationSecretScanningAlert) GetRepository() SimpleRepository {
	if o == nil || o.Repository == nil {
		var ret SimpleRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSecretScanningAlert) GetRepositoryOk() (*SimpleRepository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given SimpleRepository and assigns it to the Repository field.
func (o *OrganizationSecretScanningAlert) SetRepository(v SimpleRepository) {
	o.Repository = &v
}

// GetPushProtectionBypassed returns the PushProtectionBypassed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassed() bool {
	if o == nil || o.PushProtectionBypassed.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PushProtectionBypassed.Get()
}

// GetPushProtectionBypassedOk returns a tuple with the PushProtectionBypassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushProtectionBypassed.Get(), o.PushProtectionBypassed.IsSet()
}

// HasPushProtectionBypassed returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasPushProtectionBypassed() bool {
	if o != nil && o.PushProtectionBypassed.IsSet() {
		return true
	}

	return false
}

// SetPushProtectionBypassed gets a reference to the given NullableBool and assigns it to the PushProtectionBypassed field.
func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassed(v bool) {
	o.PushProtectionBypassed.Set(&v)
}
// SetPushProtectionBypassedNil sets the value for PushProtectionBypassed to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedNil() {
	o.PushProtectionBypassed.Set(nil)
}

// UnsetPushProtectionBypassed ensures that no value is present for PushProtectionBypassed, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetPushProtectionBypassed() {
	o.PushProtectionBypassed.Unset()
}

// GetPushProtectionBypassedBy returns the PushProtectionBypassedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedBy() NullableSimpleUser {
	if o == nil || o.PushProtectionBypassedBy.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.PushProtectionBypassedBy.Get()
}

// GetPushProtectionBypassedByOk returns a tuple with the PushProtectionBypassedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedByOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushProtectionBypassedBy.Get(), o.PushProtectionBypassedBy.IsSet()
}

// HasPushProtectionBypassedBy returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasPushProtectionBypassedBy() bool {
	if o != nil && o.PushProtectionBypassedBy.IsSet() {
		return true
	}

	return false
}

// SetPushProtectionBypassedBy gets a reference to the given NullableNullableSimpleUser and assigns it to the PushProtectionBypassedBy field.
func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedBy(v NullableSimpleUser) {
	o.PushProtectionBypassedBy.Set(&v)
}
// SetPushProtectionBypassedByNil sets the value for PushProtectionBypassedBy to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedByNil() {
	o.PushProtectionBypassedBy.Set(nil)
}

// UnsetPushProtectionBypassedBy ensures that no value is present for PushProtectionBypassedBy, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetPushProtectionBypassedBy() {
	o.PushProtectionBypassedBy.Unset()
}

// GetPushProtectionBypassedAt returns the PushProtectionBypassedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedAt() time.Time {
	if o == nil || o.PushProtectionBypassedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PushProtectionBypassedAt.Get()
}

// GetPushProtectionBypassedAtOk returns a tuple with the PushProtectionBypassedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSecretScanningAlert) GetPushProtectionBypassedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushProtectionBypassedAt.Get(), o.PushProtectionBypassedAt.IsSet()
}

// HasPushProtectionBypassedAt returns a boolean if a field has been set.
func (o *OrganizationSecretScanningAlert) HasPushProtectionBypassedAt() bool {
	if o != nil && o.PushProtectionBypassedAt.IsSet() {
		return true
	}

	return false
}

// SetPushProtectionBypassedAt gets a reference to the given NullableTime and assigns it to the PushProtectionBypassedAt field.
func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedAt(v time.Time) {
	o.PushProtectionBypassedAt.Set(&v)
}
// SetPushProtectionBypassedAtNil sets the value for PushProtectionBypassedAt to be an explicit nil
func (o *OrganizationSecretScanningAlert) SetPushProtectionBypassedAtNil() {
	o.PushProtectionBypassedAt.Set(nil)
}

// UnsetPushProtectionBypassedAt ensures that no value is present for PushProtectionBypassedAt, not even an explicit nil
func (o *OrganizationSecretScanningAlert) UnsetPushProtectionBypassedAt() {
	o.PushProtectionBypassedAt.Unset()
}

func (o OrganizationSecretScanningAlert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.HtmlUrl != nil {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if o.LocationsUrl != nil {
		toSerialize["locations_url"] = o.LocationsUrl
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Resolution.IsSet() {
		toSerialize["resolution"] = o.Resolution.Get()
	}
	if o.ResolvedAt.IsSet() {
		toSerialize["resolved_at"] = o.ResolvedAt.Get()
	}
	if o.ResolvedBy.IsSet() {
		toSerialize["resolved_by"] = o.ResolvedBy.Get()
	}
	if o.SecretType != nil {
		toSerialize["secret_type"] = o.SecretType
	}
	if o.SecretTypeDisplayName != nil {
		toSerialize["secret_type_display_name"] = o.SecretTypeDisplayName
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.PushProtectionBypassed.IsSet() {
		toSerialize["push_protection_bypassed"] = o.PushProtectionBypassed.Get()
	}
	if o.PushProtectionBypassedBy.IsSet() {
		toSerialize["push_protection_bypassed_by"] = o.PushProtectionBypassedBy.Get()
	}
	if o.PushProtectionBypassedAt.IsSet() {
		toSerialize["push_protection_bypassed_at"] = o.PushProtectionBypassedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationSecretScanningAlert struct {
	value *OrganizationSecretScanningAlert
	isSet bool
}

func (v NullableOrganizationSecretScanningAlert) Get() *OrganizationSecretScanningAlert {
	return v.value
}

func (v *NullableOrganizationSecretScanningAlert) Set(val *OrganizationSecretScanningAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSecretScanningAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSecretScanningAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSecretScanningAlert(val *OrganizationSecretScanningAlert) *NullableOrganizationSecretScanningAlert {
	return &NullableOrganizationSecretScanningAlert{value: val, isSet: true}
}

func (v NullableOrganizationSecretScanningAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSecretScanningAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

