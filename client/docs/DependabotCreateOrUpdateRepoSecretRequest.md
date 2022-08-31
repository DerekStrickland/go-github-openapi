# DependabotCreateOrUpdateRepoSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedValue** | Pointer to **string** | Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get a repository public key](https://docs.github.com/rest/reference/dependabot#get-a-repository-public-key) endpoint. | [optional] 
**KeyId** | Pointer to **string** | ID of the key you used to encrypt the secret. | [optional] 

## Methods

### NewDependabotCreateOrUpdateRepoSecretRequest

`func NewDependabotCreateOrUpdateRepoSecretRequest() *DependabotCreateOrUpdateRepoSecretRequest`

NewDependabotCreateOrUpdateRepoSecretRequest instantiates a new DependabotCreateOrUpdateRepoSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependabotCreateOrUpdateRepoSecretRequestWithDefaults

`func NewDependabotCreateOrUpdateRepoSecretRequestWithDefaults() *DependabotCreateOrUpdateRepoSecretRequest`

NewDependabotCreateOrUpdateRepoSecretRequestWithDefaults instantiates a new DependabotCreateOrUpdateRepoSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedValue

`func (o *DependabotCreateOrUpdateRepoSecretRequest) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *DependabotCreateOrUpdateRepoSecretRequest) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *DependabotCreateOrUpdateRepoSecretRequest) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *DependabotCreateOrUpdateRepoSecretRequest) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetKeyId

`func (o *DependabotCreateOrUpdateRepoSecretRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *DependabotCreateOrUpdateRepoSecretRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *DependabotCreateOrUpdateRepoSecretRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *DependabotCreateOrUpdateRepoSecretRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


