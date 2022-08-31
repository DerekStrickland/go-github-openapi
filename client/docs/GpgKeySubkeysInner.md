# GpgKeySubkeysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PrimaryKeyId** | Pointer to **int32** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to **[]interface{}** |  | [optional] 
**Subkeys** | Pointer to **[]interface{}** |  | [optional] 
**CanSign** | Pointer to **bool** |  | [optional] 
**CanEncryptComms** | Pointer to **bool** |  | [optional] 
**CanEncryptStorage** | Pointer to **bool** |  | [optional] 
**CanCertify** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **NullableString** |  | [optional] 
**RawKey** | Pointer to **NullableString** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 

## Methods

### NewGpgKeySubkeysInner

`func NewGpgKeySubkeysInner() *GpgKeySubkeysInner`

NewGpgKeySubkeysInner instantiates a new GpgKeySubkeysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpgKeySubkeysInnerWithDefaults

`func NewGpgKeySubkeysInnerWithDefaults() *GpgKeySubkeysInner`

NewGpgKeySubkeysInnerWithDefaults instantiates a new GpgKeySubkeysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GpgKeySubkeysInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpgKeySubkeysInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpgKeySubkeysInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GpgKeySubkeysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrimaryKeyId

`func (o *GpgKeySubkeysInner) GetPrimaryKeyId() int32`

GetPrimaryKeyId returns the PrimaryKeyId field if non-nil, zero value otherwise.

### GetPrimaryKeyIdOk

`func (o *GpgKeySubkeysInner) GetPrimaryKeyIdOk() (*int32, bool)`

GetPrimaryKeyIdOk returns a tuple with the PrimaryKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyId

`func (o *GpgKeySubkeysInner) SetPrimaryKeyId(v int32)`

SetPrimaryKeyId sets PrimaryKeyId field to given value.

### HasPrimaryKeyId

`func (o *GpgKeySubkeysInner) HasPrimaryKeyId() bool`

HasPrimaryKeyId returns a boolean if a field has been set.

### GetKeyId

`func (o *GpgKeySubkeysInner) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *GpgKeySubkeysInner) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *GpgKeySubkeysInner) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *GpgKeySubkeysInner) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetPublicKey

`func (o *GpgKeySubkeysInner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GpgKeySubkeysInner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GpgKeySubkeysInner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GpgKeySubkeysInner) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetEmails

`func (o *GpgKeySubkeysInner) GetEmails() []interface{}`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *GpgKeySubkeysInner) GetEmailsOk() (*[]interface{}, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *GpgKeySubkeysInner) SetEmails(v []interface{})`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *GpgKeySubkeysInner) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetSubkeys

`func (o *GpgKeySubkeysInner) GetSubkeys() []interface{}`

GetSubkeys returns the Subkeys field if non-nil, zero value otherwise.

### GetSubkeysOk

`func (o *GpgKeySubkeysInner) GetSubkeysOk() (*[]interface{}, bool)`

GetSubkeysOk returns a tuple with the Subkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubkeys

`func (o *GpgKeySubkeysInner) SetSubkeys(v []interface{})`

SetSubkeys sets Subkeys field to given value.

### HasSubkeys

`func (o *GpgKeySubkeysInner) HasSubkeys() bool`

HasSubkeys returns a boolean if a field has been set.

### GetCanSign

`func (o *GpgKeySubkeysInner) GetCanSign() bool`

GetCanSign returns the CanSign field if non-nil, zero value otherwise.

### GetCanSignOk

`func (o *GpgKeySubkeysInner) GetCanSignOk() (*bool, bool)`

GetCanSignOk returns a tuple with the CanSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSign

`func (o *GpgKeySubkeysInner) SetCanSign(v bool)`

SetCanSign sets CanSign field to given value.

### HasCanSign

`func (o *GpgKeySubkeysInner) HasCanSign() bool`

HasCanSign returns a boolean if a field has been set.

### GetCanEncryptComms

`func (o *GpgKeySubkeysInner) GetCanEncryptComms() bool`

GetCanEncryptComms returns the CanEncryptComms field if non-nil, zero value otherwise.

### GetCanEncryptCommsOk

`func (o *GpgKeySubkeysInner) GetCanEncryptCommsOk() (*bool, bool)`

GetCanEncryptCommsOk returns a tuple with the CanEncryptComms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEncryptComms

`func (o *GpgKeySubkeysInner) SetCanEncryptComms(v bool)`

SetCanEncryptComms sets CanEncryptComms field to given value.

### HasCanEncryptComms

`func (o *GpgKeySubkeysInner) HasCanEncryptComms() bool`

HasCanEncryptComms returns a boolean if a field has been set.

### GetCanEncryptStorage

`func (o *GpgKeySubkeysInner) GetCanEncryptStorage() bool`

GetCanEncryptStorage returns the CanEncryptStorage field if non-nil, zero value otherwise.

### GetCanEncryptStorageOk

`func (o *GpgKeySubkeysInner) GetCanEncryptStorageOk() (*bool, bool)`

GetCanEncryptStorageOk returns a tuple with the CanEncryptStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEncryptStorage

`func (o *GpgKeySubkeysInner) SetCanEncryptStorage(v bool)`

SetCanEncryptStorage sets CanEncryptStorage field to given value.

### HasCanEncryptStorage

`func (o *GpgKeySubkeysInner) HasCanEncryptStorage() bool`

HasCanEncryptStorage returns a boolean if a field has been set.

### GetCanCertify

`func (o *GpgKeySubkeysInner) GetCanCertify() bool`

GetCanCertify returns the CanCertify field if non-nil, zero value otherwise.

### GetCanCertifyOk

`func (o *GpgKeySubkeysInner) GetCanCertifyOk() (*bool, bool)`

GetCanCertifyOk returns a tuple with the CanCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCertify

`func (o *GpgKeySubkeysInner) SetCanCertify(v bool)`

SetCanCertify sets CanCertify field to given value.

### HasCanCertify

`func (o *GpgKeySubkeysInner) HasCanCertify() bool`

HasCanCertify returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GpgKeySubkeysInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GpgKeySubkeysInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GpgKeySubkeysInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GpgKeySubkeysInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GpgKeySubkeysInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GpgKeySubkeysInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GpgKeySubkeysInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GpgKeySubkeysInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GpgKeySubkeysInner) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GpgKeySubkeysInner) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRawKey

`func (o *GpgKeySubkeysInner) GetRawKey() string`

GetRawKey returns the RawKey field if non-nil, zero value otherwise.

### GetRawKeyOk

`func (o *GpgKeySubkeysInner) GetRawKeyOk() (*string, bool)`

GetRawKeyOk returns a tuple with the RawKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawKey

`func (o *GpgKeySubkeysInner) SetRawKey(v string)`

SetRawKey sets RawKey field to given value.

### HasRawKey

`func (o *GpgKeySubkeysInner) HasRawKey() bool`

HasRawKey returns a boolean if a field has been set.

### SetRawKeyNil

`func (o *GpgKeySubkeysInner) SetRawKeyNil(b bool)`

 SetRawKeyNil sets the value for RawKey to be an explicit nil

### UnsetRawKey
`func (o *GpgKeySubkeysInner) UnsetRawKey()`

UnsetRawKey ensures that no value is present for RawKey, not even an explicit nil
### GetRevoked

`func (o *GpgKeySubkeysInner) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *GpgKeySubkeysInner) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *GpgKeySubkeysInner) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *GpgKeySubkeysInner) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


