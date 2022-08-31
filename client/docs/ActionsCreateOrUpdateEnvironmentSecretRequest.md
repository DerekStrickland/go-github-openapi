# ActionsCreateOrUpdateEnvironmentSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedValue** | **string** | Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get an environment public key](https://docs.github.com/rest/reference/actions#get-an-environment-public-key) endpoint. | 
**KeyId** | **string** | ID of the key you used to encrypt the secret. | 

## Methods

### NewActionsCreateOrUpdateEnvironmentSecretRequest

`func NewActionsCreateOrUpdateEnvironmentSecretRequest(encryptedValue string, keyId string, ) *ActionsCreateOrUpdateEnvironmentSecretRequest`

NewActionsCreateOrUpdateEnvironmentSecretRequest instantiates a new ActionsCreateOrUpdateEnvironmentSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsCreateOrUpdateEnvironmentSecretRequestWithDefaults

`func NewActionsCreateOrUpdateEnvironmentSecretRequestWithDefaults() *ActionsCreateOrUpdateEnvironmentSecretRequest`

NewActionsCreateOrUpdateEnvironmentSecretRequestWithDefaults instantiates a new ActionsCreateOrUpdateEnvironmentSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedValue

`func (o *ActionsCreateOrUpdateEnvironmentSecretRequest) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *ActionsCreateOrUpdateEnvironmentSecretRequest) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *ActionsCreateOrUpdateEnvironmentSecretRequest) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.


### GetKeyId

`func (o *ActionsCreateOrUpdateEnvironmentSecretRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ActionsCreateOrUpdateEnvironmentSecretRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ActionsCreateOrUpdateEnvironmentSecretRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


