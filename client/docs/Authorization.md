# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Url** | **string** |  | 
**Scopes** | **[]string** | A list of scopes that this authorization is in. | 
**Token** | **string** |  | 
**TokenLastEight** | **NullableString** |  | 
**HashedToken** | **NullableString** |  | 
**App** | [**AuthorizationApp**](AuthorizationApp.md) |  | 
**Note** | **NullableString** |  | 
**NoteUrl** | **NullableString** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Fingerprint** | **NullableString** |  | 
**User** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**Installation** | Pointer to [**NullableNullableScopedInstallation**](NullableScopedInstallation.md) |  | [optional] 
**ExpiresAt** | **NullableTime** |  | 

## Methods

### NewAuthorization

`func NewAuthorization(id int32, url string, scopes []string, token string, tokenLastEight NullableString, hashedToken NullableString, app AuthorizationApp, note NullableString, noteUrl NullableString, updatedAt time.Time, createdAt time.Time, fingerprint NullableString, expiresAt NullableTime, ) *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Authorization) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authorization) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authorization) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Authorization) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Authorization) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Authorization) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetScopes

`func (o *Authorization) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Authorization) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Authorization) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *Authorization) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *Authorization) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetToken

`func (o *Authorization) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Authorization) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Authorization) SetToken(v string)`

SetToken sets Token field to given value.


### GetTokenLastEight

`func (o *Authorization) GetTokenLastEight() string`

GetTokenLastEight returns the TokenLastEight field if non-nil, zero value otherwise.

### GetTokenLastEightOk

`func (o *Authorization) GetTokenLastEightOk() (*string, bool)`

GetTokenLastEightOk returns a tuple with the TokenLastEight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLastEight

`func (o *Authorization) SetTokenLastEight(v string)`

SetTokenLastEight sets TokenLastEight field to given value.


### SetTokenLastEightNil

`func (o *Authorization) SetTokenLastEightNil(b bool)`

 SetTokenLastEightNil sets the value for TokenLastEight to be an explicit nil

### UnsetTokenLastEight
`func (o *Authorization) UnsetTokenLastEight()`

UnsetTokenLastEight ensures that no value is present for TokenLastEight, not even an explicit nil
### GetHashedToken

`func (o *Authorization) GetHashedToken() string`

GetHashedToken returns the HashedToken field if non-nil, zero value otherwise.

### GetHashedTokenOk

`func (o *Authorization) GetHashedTokenOk() (*string, bool)`

GetHashedTokenOk returns a tuple with the HashedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedToken

`func (o *Authorization) SetHashedToken(v string)`

SetHashedToken sets HashedToken field to given value.


### SetHashedTokenNil

`func (o *Authorization) SetHashedTokenNil(b bool)`

 SetHashedTokenNil sets the value for HashedToken to be an explicit nil

### UnsetHashedToken
`func (o *Authorization) UnsetHashedToken()`

UnsetHashedToken ensures that no value is present for HashedToken, not even an explicit nil
### GetApp

`func (o *Authorization) GetApp() AuthorizationApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *Authorization) GetAppOk() (*AuthorizationApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *Authorization) SetApp(v AuthorizationApp)`

SetApp sets App field to given value.


### GetNote

`func (o *Authorization) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Authorization) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Authorization) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *Authorization) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *Authorization) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetNoteUrl

`func (o *Authorization) GetNoteUrl() string`

GetNoteUrl returns the NoteUrl field if non-nil, zero value otherwise.

### GetNoteUrlOk

`func (o *Authorization) GetNoteUrlOk() (*string, bool)`

GetNoteUrlOk returns a tuple with the NoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteUrl

`func (o *Authorization) SetNoteUrl(v string)`

SetNoteUrl sets NoteUrl field to given value.


### SetNoteUrlNil

`func (o *Authorization) SetNoteUrlNil(b bool)`

 SetNoteUrlNil sets the value for NoteUrl to be an explicit nil

### UnsetNoteUrl
`func (o *Authorization) UnsetNoteUrl()`

UnsetNoteUrl ensures that no value is present for NoteUrl, not even an explicit nil
### GetUpdatedAt

`func (o *Authorization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Authorization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Authorization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Authorization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Authorization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Authorization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFingerprint

`func (o *Authorization) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Authorization) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Authorization) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### SetFingerprintNil

`func (o *Authorization) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Authorization) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetUser

`func (o *Authorization) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Authorization) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Authorization) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Authorization) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Authorization) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Authorization) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetInstallation

`func (o *Authorization) GetInstallation() NullableScopedInstallation`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *Authorization) GetInstallationOk() (*NullableScopedInstallation, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *Authorization) SetInstallation(v NullableScopedInstallation)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *Authorization) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### SetInstallationNil

`func (o *Authorization) SetInstallationNil(b bool)`

 SetInstallationNil sets the value for Installation to be an explicit nil

### UnsetInstallation
`func (o *Authorization) UnsetInstallation()`

UnsetInstallation ensures that no value is present for Installation, not even an explicit nil
### GetExpiresAt

`func (o *Authorization) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Authorization) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Authorization) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Authorization) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Authorization) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


