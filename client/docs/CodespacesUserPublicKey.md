# CodespacesUserPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | The identifier for the key. | 
**Key** | **string** | The Base64 encoded public key. | 

## Methods

### NewCodespacesUserPublicKey

`func NewCodespacesUserPublicKey(keyId string, key string, ) *CodespacesUserPublicKey`

NewCodespacesUserPublicKey instantiates a new CodespacesUserPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesUserPublicKeyWithDefaults

`func NewCodespacesUserPublicKeyWithDefaults() *CodespacesUserPublicKey`

NewCodespacesUserPublicKeyWithDefaults instantiates a new CodespacesUserPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *CodespacesUserPublicKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *CodespacesUserPublicKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *CodespacesUserPublicKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKey

`func (o *CodespacesUserPublicKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CodespacesUserPublicKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CodespacesUserPublicKey) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


