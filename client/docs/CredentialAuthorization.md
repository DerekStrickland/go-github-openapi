# CredentialAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | User login that owns the underlying credential. | 
**CredentialId** | **int32** | Unique identifier for the credential. | 
**CredentialType** | **string** | Human-readable description of the credential type. | 
**TokenLastEight** | Pointer to **string** | Last eight characters of the credential. Only included in responses with credential_type of personal access token. | [optional] 
**CredentialAuthorizedAt** | **time.Time** | Date when the credential was authorized for use. | 
**Scopes** | Pointer to **[]string** | List of oauth scopes the token has been granted. | [optional] 
**Fingerprint** | Pointer to **string** | Unique string to distinguish the credential. Only included in responses with credential_type of SSH Key. | [optional] 
**CredentialAccessedAt** | **NullableTime** | Date when the credential was last accessed. May be null if it was never accessed | 
**AuthorizedCredentialId** | **NullableInt32** |  | 
**AuthorizedCredentialTitle** | Pointer to **NullableString** | The title given to the ssh key. This will only be present when the credential is an ssh key. | [optional] 
**AuthorizedCredentialNote** | Pointer to **NullableString** | The note given to the token. This will only be present when the credential is a token. | [optional] 
**AuthorizedCredentialExpiresAt** | Pointer to **NullableTime** | The expiry for the token. This will only be present when the credential is a token. | [optional] 

## Methods

### NewCredentialAuthorization

`func NewCredentialAuthorization(login string, credentialId int32, credentialType string, credentialAuthorizedAt time.Time, credentialAccessedAt NullableTime, authorizedCredentialId NullableInt32, ) *CredentialAuthorization`

NewCredentialAuthorization instantiates a new CredentialAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialAuthorizationWithDefaults

`func NewCredentialAuthorizationWithDefaults() *CredentialAuthorization`

NewCredentialAuthorizationWithDefaults instantiates a new CredentialAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CredentialAuthorization) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CredentialAuthorization) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CredentialAuthorization) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetCredentialId

`func (o *CredentialAuthorization) GetCredentialId() int32`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *CredentialAuthorization) GetCredentialIdOk() (*int32, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *CredentialAuthorization) SetCredentialId(v int32)`

SetCredentialId sets CredentialId field to given value.


### GetCredentialType

`func (o *CredentialAuthorization) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *CredentialAuthorization) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *CredentialAuthorization) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.


### GetTokenLastEight

`func (o *CredentialAuthorization) GetTokenLastEight() string`

GetTokenLastEight returns the TokenLastEight field if non-nil, zero value otherwise.

### GetTokenLastEightOk

`func (o *CredentialAuthorization) GetTokenLastEightOk() (*string, bool)`

GetTokenLastEightOk returns a tuple with the TokenLastEight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLastEight

`func (o *CredentialAuthorization) SetTokenLastEight(v string)`

SetTokenLastEight sets TokenLastEight field to given value.

### HasTokenLastEight

`func (o *CredentialAuthorization) HasTokenLastEight() bool`

HasTokenLastEight returns a boolean if a field has been set.

### GetCredentialAuthorizedAt

`func (o *CredentialAuthorization) GetCredentialAuthorizedAt() time.Time`

GetCredentialAuthorizedAt returns the CredentialAuthorizedAt field if non-nil, zero value otherwise.

### GetCredentialAuthorizedAtOk

`func (o *CredentialAuthorization) GetCredentialAuthorizedAtOk() (*time.Time, bool)`

GetCredentialAuthorizedAtOk returns a tuple with the CredentialAuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialAuthorizedAt

`func (o *CredentialAuthorization) SetCredentialAuthorizedAt(v time.Time)`

SetCredentialAuthorizedAt sets CredentialAuthorizedAt field to given value.


### GetScopes

`func (o *CredentialAuthorization) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CredentialAuthorization) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CredentialAuthorization) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CredentialAuthorization) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetFingerprint

`func (o *CredentialAuthorization) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CredentialAuthorization) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CredentialAuthorization) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CredentialAuthorization) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCredentialAccessedAt

`func (o *CredentialAuthorization) GetCredentialAccessedAt() time.Time`

GetCredentialAccessedAt returns the CredentialAccessedAt field if non-nil, zero value otherwise.

### GetCredentialAccessedAtOk

`func (o *CredentialAuthorization) GetCredentialAccessedAtOk() (*time.Time, bool)`

GetCredentialAccessedAtOk returns a tuple with the CredentialAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialAccessedAt

`func (o *CredentialAuthorization) SetCredentialAccessedAt(v time.Time)`

SetCredentialAccessedAt sets CredentialAccessedAt field to given value.


### SetCredentialAccessedAtNil

`func (o *CredentialAuthorization) SetCredentialAccessedAtNil(b bool)`

 SetCredentialAccessedAtNil sets the value for CredentialAccessedAt to be an explicit nil

### UnsetCredentialAccessedAt
`func (o *CredentialAuthorization) UnsetCredentialAccessedAt()`

UnsetCredentialAccessedAt ensures that no value is present for CredentialAccessedAt, not even an explicit nil
### GetAuthorizedCredentialId

`func (o *CredentialAuthorization) GetAuthorizedCredentialId() int32`

GetAuthorizedCredentialId returns the AuthorizedCredentialId field if non-nil, zero value otherwise.

### GetAuthorizedCredentialIdOk

`func (o *CredentialAuthorization) GetAuthorizedCredentialIdOk() (*int32, bool)`

GetAuthorizedCredentialIdOk returns a tuple with the AuthorizedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCredentialId

`func (o *CredentialAuthorization) SetAuthorizedCredentialId(v int32)`

SetAuthorizedCredentialId sets AuthorizedCredentialId field to given value.


### SetAuthorizedCredentialIdNil

`func (o *CredentialAuthorization) SetAuthorizedCredentialIdNil(b bool)`

 SetAuthorizedCredentialIdNil sets the value for AuthorizedCredentialId to be an explicit nil

### UnsetAuthorizedCredentialId
`func (o *CredentialAuthorization) UnsetAuthorizedCredentialId()`

UnsetAuthorizedCredentialId ensures that no value is present for AuthorizedCredentialId, not even an explicit nil
### GetAuthorizedCredentialTitle

`func (o *CredentialAuthorization) GetAuthorizedCredentialTitle() string`

GetAuthorizedCredentialTitle returns the AuthorizedCredentialTitle field if non-nil, zero value otherwise.

### GetAuthorizedCredentialTitleOk

`func (o *CredentialAuthorization) GetAuthorizedCredentialTitleOk() (*string, bool)`

GetAuthorizedCredentialTitleOk returns a tuple with the AuthorizedCredentialTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCredentialTitle

`func (o *CredentialAuthorization) SetAuthorizedCredentialTitle(v string)`

SetAuthorizedCredentialTitle sets AuthorizedCredentialTitle field to given value.

### HasAuthorizedCredentialTitle

`func (o *CredentialAuthorization) HasAuthorizedCredentialTitle() bool`

HasAuthorizedCredentialTitle returns a boolean if a field has been set.

### SetAuthorizedCredentialTitleNil

`func (o *CredentialAuthorization) SetAuthorizedCredentialTitleNil(b bool)`

 SetAuthorizedCredentialTitleNil sets the value for AuthorizedCredentialTitle to be an explicit nil

### UnsetAuthorizedCredentialTitle
`func (o *CredentialAuthorization) UnsetAuthorizedCredentialTitle()`

UnsetAuthorizedCredentialTitle ensures that no value is present for AuthorizedCredentialTitle, not even an explicit nil
### GetAuthorizedCredentialNote

`func (o *CredentialAuthorization) GetAuthorizedCredentialNote() string`

GetAuthorizedCredentialNote returns the AuthorizedCredentialNote field if non-nil, zero value otherwise.

### GetAuthorizedCredentialNoteOk

`func (o *CredentialAuthorization) GetAuthorizedCredentialNoteOk() (*string, bool)`

GetAuthorizedCredentialNoteOk returns a tuple with the AuthorizedCredentialNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCredentialNote

`func (o *CredentialAuthorization) SetAuthorizedCredentialNote(v string)`

SetAuthorizedCredentialNote sets AuthorizedCredentialNote field to given value.

### HasAuthorizedCredentialNote

`func (o *CredentialAuthorization) HasAuthorizedCredentialNote() bool`

HasAuthorizedCredentialNote returns a boolean if a field has been set.

### SetAuthorizedCredentialNoteNil

`func (o *CredentialAuthorization) SetAuthorizedCredentialNoteNil(b bool)`

 SetAuthorizedCredentialNoteNil sets the value for AuthorizedCredentialNote to be an explicit nil

### UnsetAuthorizedCredentialNote
`func (o *CredentialAuthorization) UnsetAuthorizedCredentialNote()`

UnsetAuthorizedCredentialNote ensures that no value is present for AuthorizedCredentialNote, not even an explicit nil
### GetAuthorizedCredentialExpiresAt

`func (o *CredentialAuthorization) GetAuthorizedCredentialExpiresAt() time.Time`

GetAuthorizedCredentialExpiresAt returns the AuthorizedCredentialExpiresAt field if non-nil, zero value otherwise.

### GetAuthorizedCredentialExpiresAtOk

`func (o *CredentialAuthorization) GetAuthorizedCredentialExpiresAtOk() (*time.Time, bool)`

GetAuthorizedCredentialExpiresAtOk returns a tuple with the AuthorizedCredentialExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCredentialExpiresAt

`func (o *CredentialAuthorization) SetAuthorizedCredentialExpiresAt(v time.Time)`

SetAuthorizedCredentialExpiresAt sets AuthorizedCredentialExpiresAt field to given value.

### HasAuthorizedCredentialExpiresAt

`func (o *CredentialAuthorization) HasAuthorizedCredentialExpiresAt() bool`

HasAuthorizedCredentialExpiresAt returns a boolean if a field has been set.

### SetAuthorizedCredentialExpiresAtNil

`func (o *CredentialAuthorization) SetAuthorizedCredentialExpiresAtNil(b bool)`

 SetAuthorizedCredentialExpiresAtNil sets the value for AuthorizedCredentialExpiresAt to be an explicit nil

### UnsetAuthorizedCredentialExpiresAt
`func (o *CredentialAuthorization) UnsetAuthorizedCredentialExpiresAt()`

UnsetAuthorizedCredentialExpiresAt ensures that no value is present for AuthorizedCredentialExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


