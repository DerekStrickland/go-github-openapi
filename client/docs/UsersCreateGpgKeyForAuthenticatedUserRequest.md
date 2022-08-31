# UsersCreateGpgKeyForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the new key. | [optional] 
**ArmoredPublicKey** | **string** | A GPG key in ASCII-armored format. | 

## Methods

### NewUsersCreateGpgKeyForAuthenticatedUserRequest

`func NewUsersCreateGpgKeyForAuthenticatedUserRequest(armoredPublicKey string, ) *UsersCreateGpgKeyForAuthenticatedUserRequest`

NewUsersCreateGpgKeyForAuthenticatedUserRequest instantiates a new UsersCreateGpgKeyForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersCreateGpgKeyForAuthenticatedUserRequestWithDefaults

`func NewUsersCreateGpgKeyForAuthenticatedUserRequestWithDefaults() *UsersCreateGpgKeyForAuthenticatedUserRequest`

NewUsersCreateGpgKeyForAuthenticatedUserRequestWithDefaults instantiates a new UsersCreateGpgKeyForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArmoredPublicKey

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetArmoredPublicKey() string`

GetArmoredPublicKey returns the ArmoredPublicKey field if non-nil, zero value otherwise.

### GetArmoredPublicKeyOk

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) GetArmoredPublicKeyOk() (*string, bool)`

GetArmoredPublicKeyOk returns a tuple with the ArmoredPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmoredPublicKey

`func (o *UsersCreateGpgKeyForAuthenticatedUserRequest) SetArmoredPublicKey(v string)`

SetArmoredPublicKey sets ArmoredPublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


