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

// CredentialAuthorization Credential Authorization
type CredentialAuthorization struct {
	// User login that owns the underlying credential.
	Login string `json:"login"`
	// Unique identifier for the credential.
	CredentialId int32 `json:"credential_id"`
	// Human-readable description of the credential type.
	CredentialType string `json:"credential_type"`
	// Last eight characters of the credential. Only included in responses with credential_type of personal access token.
	TokenLastEight *string `json:"token_last_eight,omitempty"`
	// Date when the credential was authorized for use.
	CredentialAuthorizedAt time.Time `json:"credential_authorized_at"`
	// List of oauth scopes the token has been granted.
	Scopes []string `json:"scopes,omitempty"`
	// Unique string to distinguish the credential. Only included in responses with credential_type of SSH Key.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Date when the credential was last accessed. May be null if it was never accessed
	CredentialAccessedAt NullableTime `json:"credential_accessed_at"`
	AuthorizedCredentialId NullableInt32 `json:"authorized_credential_id"`
	// The title given to the ssh key. This will only be present when the credential is an ssh key.
	AuthorizedCredentialTitle NullableString `json:"authorized_credential_title,omitempty"`
	// The note given to the token. This will only be present when the credential is a token.
	AuthorizedCredentialNote NullableString `json:"authorized_credential_note,omitempty"`
	// The expiry for the token. This will only be present when the credential is a token.
	AuthorizedCredentialExpiresAt NullableTime `json:"authorized_credential_expires_at,omitempty"`
}

// NewCredentialAuthorization instantiates a new CredentialAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialAuthorization(login string, credentialId int32, credentialType string, credentialAuthorizedAt time.Time, credentialAccessedAt NullableTime, authorizedCredentialId NullableInt32) *CredentialAuthorization {
	this := CredentialAuthorization{}
	this.Login = login
	this.CredentialId = credentialId
	this.CredentialType = credentialType
	this.CredentialAuthorizedAt = credentialAuthorizedAt
	this.CredentialAccessedAt = credentialAccessedAt
	this.AuthorizedCredentialId = authorizedCredentialId
	return &this
}

// NewCredentialAuthorizationWithDefaults instantiates a new CredentialAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialAuthorizationWithDefaults() *CredentialAuthorization {
	this := CredentialAuthorization{}
	return &this
}

// GetLogin returns the Login field value
func (o *CredentialAuthorization) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *CredentialAuthorization) SetLogin(v string) {
	o.Login = v
}

// GetCredentialId returns the CredentialId field value
func (o *CredentialAuthorization) GetCredentialId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetCredentialIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialId, true
}

// SetCredentialId sets field value
func (o *CredentialAuthorization) SetCredentialId(v int32) {
	o.CredentialId = v
}

// GetCredentialType returns the CredentialType field value
func (o *CredentialAuthorization) GetCredentialType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetCredentialTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialType, true
}

// SetCredentialType sets field value
func (o *CredentialAuthorization) SetCredentialType(v string) {
	o.CredentialType = v
}

// GetTokenLastEight returns the TokenLastEight field value if set, zero value otherwise.
func (o *CredentialAuthorization) GetTokenLastEight() string {
	if o == nil || o.TokenLastEight == nil {
		var ret string
		return ret
	}
	return *o.TokenLastEight
}

// GetTokenLastEightOk returns a tuple with the TokenLastEight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetTokenLastEightOk() (*string, bool) {
	if o == nil || o.TokenLastEight == nil {
		return nil, false
	}
	return o.TokenLastEight, true
}

// HasTokenLastEight returns a boolean if a field has been set.
func (o *CredentialAuthorization) HasTokenLastEight() bool {
	if o != nil && o.TokenLastEight != nil {
		return true
	}

	return false
}

// SetTokenLastEight gets a reference to the given string and assigns it to the TokenLastEight field.
func (o *CredentialAuthorization) SetTokenLastEight(v string) {
	o.TokenLastEight = &v
}

// GetCredentialAuthorizedAt returns the CredentialAuthorizedAt field value
func (o *CredentialAuthorization) GetCredentialAuthorizedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CredentialAuthorizedAt
}

// GetCredentialAuthorizedAtOk returns a tuple with the CredentialAuthorizedAt field value
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetCredentialAuthorizedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialAuthorizedAt, true
}

// SetCredentialAuthorizedAt sets field value
func (o *CredentialAuthorization) SetCredentialAuthorizedAt(v time.Time) {
	o.CredentialAuthorizedAt = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *CredentialAuthorization) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *CredentialAuthorization) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *CredentialAuthorization) SetScopes(v []string) {
	o.Scopes = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CredentialAuthorization) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialAuthorization) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CredentialAuthorization) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CredentialAuthorization) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetCredentialAccessedAt returns the CredentialAccessedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CredentialAuthorization) GetCredentialAccessedAt() time.Time {
	if o == nil || o.CredentialAccessedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CredentialAccessedAt.Get()
}

// GetCredentialAccessedAtOk returns a tuple with the CredentialAccessedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialAuthorization) GetCredentialAccessedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialAccessedAt.Get(), o.CredentialAccessedAt.IsSet()
}

// SetCredentialAccessedAt sets field value
func (o *CredentialAuthorization) SetCredentialAccessedAt(v time.Time) {
	o.CredentialAccessedAt.Set(&v)
}

// GetAuthorizedCredentialId returns the AuthorizedCredentialId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CredentialAuthorization) GetAuthorizedCredentialId() int32 {
	if o == nil || o.AuthorizedCredentialId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AuthorizedCredentialId.Get()
}

// GetAuthorizedCredentialIdOk returns a tuple with the AuthorizedCredentialId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialAuthorization) GetAuthorizedCredentialIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedCredentialId.Get(), o.AuthorizedCredentialId.IsSet()
}

// SetAuthorizedCredentialId sets field value
func (o *CredentialAuthorization) SetAuthorizedCredentialId(v int32) {
	o.AuthorizedCredentialId.Set(&v)
}

// GetAuthorizedCredentialTitle returns the AuthorizedCredentialTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialAuthorization) GetAuthorizedCredentialTitle() string {
	if o == nil || o.AuthorizedCredentialTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizedCredentialTitle.Get()
}

// GetAuthorizedCredentialTitleOk returns a tuple with the AuthorizedCredentialTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialAuthorization) GetAuthorizedCredentialTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedCredentialTitle.Get(), o.AuthorizedCredentialTitle.IsSet()
}

// HasAuthorizedCredentialTitle returns a boolean if a field has been set.
func (o *CredentialAuthorization) HasAuthorizedCredentialTitle() bool {
	if o != nil && o.AuthorizedCredentialTitle.IsSet() {
		return true
	}

	return false
}

// SetAuthorizedCredentialTitle gets a reference to the given NullableString and assigns it to the AuthorizedCredentialTitle field.
func (o *CredentialAuthorization) SetAuthorizedCredentialTitle(v string) {
	o.AuthorizedCredentialTitle.Set(&v)
}
// SetAuthorizedCredentialTitleNil sets the value for AuthorizedCredentialTitle to be an explicit nil
func (o *CredentialAuthorization) SetAuthorizedCredentialTitleNil() {
	o.AuthorizedCredentialTitle.Set(nil)
}

// UnsetAuthorizedCredentialTitle ensures that no value is present for AuthorizedCredentialTitle, not even an explicit nil
func (o *CredentialAuthorization) UnsetAuthorizedCredentialTitle() {
	o.AuthorizedCredentialTitle.Unset()
}

// GetAuthorizedCredentialNote returns the AuthorizedCredentialNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialAuthorization) GetAuthorizedCredentialNote() string {
	if o == nil || o.AuthorizedCredentialNote.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizedCredentialNote.Get()
}

// GetAuthorizedCredentialNoteOk returns a tuple with the AuthorizedCredentialNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialAuthorization) GetAuthorizedCredentialNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedCredentialNote.Get(), o.AuthorizedCredentialNote.IsSet()
}

// HasAuthorizedCredentialNote returns a boolean if a field has been set.
func (o *CredentialAuthorization) HasAuthorizedCredentialNote() bool {
	if o != nil && o.AuthorizedCredentialNote.IsSet() {
		return true
	}

	return false
}

// SetAuthorizedCredentialNote gets a reference to the given NullableString and assigns it to the AuthorizedCredentialNote field.
func (o *CredentialAuthorization) SetAuthorizedCredentialNote(v string) {
	o.AuthorizedCredentialNote.Set(&v)
}
// SetAuthorizedCredentialNoteNil sets the value for AuthorizedCredentialNote to be an explicit nil
func (o *CredentialAuthorization) SetAuthorizedCredentialNoteNil() {
	o.AuthorizedCredentialNote.Set(nil)
}

// UnsetAuthorizedCredentialNote ensures that no value is present for AuthorizedCredentialNote, not even an explicit nil
func (o *CredentialAuthorization) UnsetAuthorizedCredentialNote() {
	o.AuthorizedCredentialNote.Unset()
}

// GetAuthorizedCredentialExpiresAt returns the AuthorizedCredentialExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialAuthorization) GetAuthorizedCredentialExpiresAt() time.Time {
	if o == nil || o.AuthorizedCredentialExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AuthorizedCredentialExpiresAt.Get()
}

// GetAuthorizedCredentialExpiresAtOk returns a tuple with the AuthorizedCredentialExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialAuthorization) GetAuthorizedCredentialExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedCredentialExpiresAt.Get(), o.AuthorizedCredentialExpiresAt.IsSet()
}

// HasAuthorizedCredentialExpiresAt returns a boolean if a field has been set.
func (o *CredentialAuthorization) HasAuthorizedCredentialExpiresAt() bool {
	if o != nil && o.AuthorizedCredentialExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetAuthorizedCredentialExpiresAt gets a reference to the given NullableTime and assigns it to the AuthorizedCredentialExpiresAt field.
func (o *CredentialAuthorization) SetAuthorizedCredentialExpiresAt(v time.Time) {
	o.AuthorizedCredentialExpiresAt.Set(&v)
}
// SetAuthorizedCredentialExpiresAtNil sets the value for AuthorizedCredentialExpiresAt to be an explicit nil
func (o *CredentialAuthorization) SetAuthorizedCredentialExpiresAtNil() {
	o.AuthorizedCredentialExpiresAt.Set(nil)
}

// UnsetAuthorizedCredentialExpiresAt ensures that no value is present for AuthorizedCredentialExpiresAt, not even an explicit nil
func (o *CredentialAuthorization) UnsetAuthorizedCredentialExpiresAt() {
	o.AuthorizedCredentialExpiresAt.Unset()
}

func (o CredentialAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["credential_id"] = o.CredentialId
	}
	if true {
		toSerialize["credential_type"] = o.CredentialType
	}
	if o.TokenLastEight != nil {
		toSerialize["token_last_eight"] = o.TokenLastEight
	}
	if true {
		toSerialize["credential_authorized_at"] = o.CredentialAuthorizedAt
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if true {
		toSerialize["credential_accessed_at"] = o.CredentialAccessedAt.Get()
	}
	if true {
		toSerialize["authorized_credential_id"] = o.AuthorizedCredentialId.Get()
	}
	if o.AuthorizedCredentialTitle.IsSet() {
		toSerialize["authorized_credential_title"] = o.AuthorizedCredentialTitle.Get()
	}
	if o.AuthorizedCredentialNote.IsSet() {
		toSerialize["authorized_credential_note"] = o.AuthorizedCredentialNote.Get()
	}
	if o.AuthorizedCredentialExpiresAt.IsSet() {
		toSerialize["authorized_credential_expires_at"] = o.AuthorizedCredentialExpiresAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialAuthorization struct {
	value *CredentialAuthorization
	isSet bool
}

func (v NullableCredentialAuthorization) Get() *CredentialAuthorization {
	return v.value
}

func (v *NullableCredentialAuthorization) Set(val *CredentialAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialAuthorization(val *CredentialAuthorization) *NullableCredentialAuthorization {
	return &NullableCredentialAuthorization{value: val, isSet: true}
}

func (v NullableCredentialAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


