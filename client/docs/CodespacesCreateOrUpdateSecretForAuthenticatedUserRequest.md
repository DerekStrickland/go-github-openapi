# CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedValue** | Pointer to **string** | Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get the public key for the authenticated user](https://docs.github.com/rest/reference/codespaces#get-the-public-key-for-the-authenticated-user) endpoint. | [optional] 
**KeyId** | **string** | ID of the key you used to encrypt the secret. | 
**SelectedRepositoryIds** | Pointer to **[]string** | An array of repository ids that can access the user secret. You can manage the list of selected repositories using the [List selected repositories for a user secret](https://docs.github.com/rest/reference/codespaces#list-selected-repositories-for-a-user-secret), [Set selected repositories for a user secret](https://docs.github.com/rest/reference/codespaces#set-selected-repositories-for-a-user-secret), and [Remove a selected repository from a user secret](https://docs.github.com/rest/reference/codespaces#remove-a-selected-repository-from-a-user-secret) endpoints. | [optional] 

## Methods

### NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest

`func NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest(keyId string, ) *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest`

NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest instantiates a new CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequestWithDefaults

`func NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequestWithDefaults() *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest`

NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequestWithDefaults instantiates a new CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedValue

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetKeyId

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetSelectedRepositoryIds

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetSelectedRepositoryIds() []string`

GetSelectedRepositoryIds returns the SelectedRepositoryIds field if non-nil, zero value otherwise.

### GetSelectedRepositoryIdsOk

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetSelectedRepositoryIdsOk() (*[]string, bool)`

GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoryIds

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) SetSelectedRepositoryIds(v []string)`

SetSelectedRepositoryIds sets SelectedRepositoryIds field to given value.

### HasSelectedRepositoryIds

`func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) HasSelectedRepositoryIds() bool`

HasSelectedRepositoryIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


