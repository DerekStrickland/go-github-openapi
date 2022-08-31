# GpgKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**PrimaryKeyId** | **NullableInt32** |  | 
**KeyId** | **string** |  | 
**PublicKey** | **string** |  | 
**Emails** | [**[]GpgKeyEmailsInner**](GpgKeyEmailsInner.md) |  | 
**Subkeys** | [**[]GpgKeySubkeysInner**](GpgKeySubkeysInner.md) |  | 
**CanSign** | **bool** |  | 
**CanEncryptComms** | **bool** |  | 
**CanEncryptStorage** | **bool** |  | 
**CanCertify** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | **NullableTime** |  | 
**Revoked** | **bool** |  | 
**RawKey** | **NullableString** |  | 

## Methods

### NewGpgKey

`func NewGpgKey(id int32, primaryKeyId NullableInt32, keyId string, publicKey string, emails []GpgKeyEmailsInner, subkeys []GpgKeySubkeysInner, canSign bool, canEncryptComms bool, canEncryptStorage bool, canCertify bool, createdAt time.Time, expiresAt NullableTime, revoked bool, rawKey NullableString, ) *GpgKey`

NewGpgKey instantiates a new GpgKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpgKeyWithDefaults

`func NewGpgKeyWithDefaults() *GpgKey`

NewGpgKeyWithDefaults instantiates a new GpgKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GpgKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpgKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpgKey) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GpgKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpgKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpgKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GpgKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GpgKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GpgKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrimaryKeyId

`func (o *GpgKey) GetPrimaryKeyId() int32`

GetPrimaryKeyId returns the PrimaryKeyId field if non-nil, zero value otherwise.

### GetPrimaryKeyIdOk

`func (o *GpgKey) GetPrimaryKeyIdOk() (*int32, bool)`

GetPrimaryKeyIdOk returns a tuple with the PrimaryKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyId

`func (o *GpgKey) SetPrimaryKeyId(v int32)`

SetPrimaryKeyId sets PrimaryKeyId field to given value.


### SetPrimaryKeyIdNil

`func (o *GpgKey) SetPrimaryKeyIdNil(b bool)`

 SetPrimaryKeyIdNil sets the value for PrimaryKeyId to be an explicit nil

### UnsetPrimaryKeyId
`func (o *GpgKey) UnsetPrimaryKeyId()`

UnsetPrimaryKeyId ensures that no value is present for PrimaryKeyId, not even an explicit nil
### GetKeyId

`func (o *GpgKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *GpgKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *GpgKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetPublicKey

`func (o *GpgKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GpgKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GpgKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetEmails

`func (o *GpgKey) GetEmails() []GpgKeyEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *GpgKey) GetEmailsOk() (*[]GpgKeyEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *GpgKey) SetEmails(v []GpgKeyEmailsInner)`

SetEmails sets Emails field to given value.


### GetSubkeys

`func (o *GpgKey) GetSubkeys() []GpgKeySubkeysInner`

GetSubkeys returns the Subkeys field if non-nil, zero value otherwise.

### GetSubkeysOk

`func (o *GpgKey) GetSubkeysOk() (*[]GpgKeySubkeysInner, bool)`

GetSubkeysOk returns a tuple with the Subkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubkeys

`func (o *GpgKey) SetSubkeys(v []GpgKeySubkeysInner)`

SetSubkeys sets Subkeys field to given value.


### GetCanSign

`func (o *GpgKey) GetCanSign() bool`

GetCanSign returns the CanSign field if non-nil, zero value otherwise.

### GetCanSignOk

`func (o *GpgKey) GetCanSignOk() (*bool, bool)`

GetCanSignOk returns a tuple with the CanSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSign

`func (o *GpgKey) SetCanSign(v bool)`

SetCanSign sets CanSign field to given value.


### GetCanEncryptComms

`func (o *GpgKey) GetCanEncryptComms() bool`

GetCanEncryptComms returns the CanEncryptComms field if non-nil, zero value otherwise.

### GetCanEncryptCommsOk

`func (o *GpgKey) GetCanEncryptCommsOk() (*bool, bool)`

GetCanEncryptCommsOk returns a tuple with the CanEncryptComms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEncryptComms

`func (o *GpgKey) SetCanEncryptComms(v bool)`

SetCanEncryptComms sets CanEncryptComms field to given value.


### GetCanEncryptStorage

`func (o *GpgKey) GetCanEncryptStorage() bool`

GetCanEncryptStorage returns the CanEncryptStorage field if non-nil, zero value otherwise.

### GetCanEncryptStorageOk

`func (o *GpgKey) GetCanEncryptStorageOk() (*bool, bool)`

GetCanEncryptStorageOk returns a tuple with the CanEncryptStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEncryptStorage

`func (o *GpgKey) SetCanEncryptStorage(v bool)`

SetCanEncryptStorage sets CanEncryptStorage field to given value.


### GetCanCertify

`func (o *GpgKey) GetCanCertify() bool`

GetCanCertify returns the CanCertify field if non-nil, zero value otherwise.

### GetCanCertifyOk

`func (o *GpgKey) GetCanCertifyOk() (*bool, bool)`

GetCanCertifyOk returns a tuple with the CanCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCertify

`func (o *GpgKey) SetCanCertify(v bool)`

SetCanCertify sets CanCertify field to given value.


### GetCreatedAt

`func (o *GpgKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GpgKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GpgKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *GpgKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GpgKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GpgKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *GpgKey) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GpgKey) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRevoked

`func (o *GpgKey) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *GpgKey) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *GpgKey) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### GetRawKey

`func (o *GpgKey) GetRawKey() string`

GetRawKey returns the RawKey field if non-nil, zero value otherwise.

### GetRawKeyOk

`func (o *GpgKey) GetRawKeyOk() (*string, bool)`

GetRawKeyOk returns a tuple with the RawKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawKey

`func (o *GpgKey) SetRawKey(v string)`

SetRawKey sets RawKey field to given value.


### SetRawKeyNil

`func (o *GpgKey) SetRawKeyNil(b bool)`

 SetRawKeyNil sets the value for RawKey to be an explicit nil

### UnsetRawKey
`func (o *GpgKey) UnsetRawKey()`

UnsetRawKey ensures that no value is present for RawKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


