# ActionsCreateOrUpdateOrgSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedValue** | Pointer to **string** | Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get an organization public key](https://docs.github.com/rest/reference/actions#get-an-organization-public-key) endpoint. | [optional] 
**KeyId** | Pointer to **string** | ID of the key you used to encrypt the secret. | [optional] 
**Visibility** | **string** | Which type of organization repositories have access to the organization secret. &#x60;selected&#x60; means only the repositories specified by &#x60;selected_repository_ids&#x60; can access the secret. | 
**SelectedRepositoryIds** | Pointer to **[]int32** | An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the &#x60;visibility&#x60; is set to &#x60;selected&#x60;. You can manage the list of selected repositories using the [List selected repositories for an organization secret](https://docs.github.com/rest/reference/actions#list-selected-repositories-for-an-organization-secret), [Set selected repositories for an organization secret](https://docs.github.com/rest/reference/actions#set-selected-repositories-for-an-organization-secret), and [Remove selected repository from an organization secret](https://docs.github.com/rest/reference/actions#remove-selected-repository-from-an-organization-secret) endpoints. | [optional] 

## Methods

### NewActionsCreateOrUpdateOrgSecretRequest

`func NewActionsCreateOrUpdateOrgSecretRequest(visibility string, ) *ActionsCreateOrUpdateOrgSecretRequest`

NewActionsCreateOrUpdateOrgSecretRequest instantiates a new ActionsCreateOrUpdateOrgSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsCreateOrUpdateOrgSecretRequestWithDefaults

`func NewActionsCreateOrUpdateOrgSecretRequestWithDefaults() *ActionsCreateOrUpdateOrgSecretRequest`

NewActionsCreateOrUpdateOrgSecretRequestWithDefaults instantiates a new ActionsCreateOrUpdateOrgSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedValue

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *ActionsCreateOrUpdateOrgSecretRequest) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *ActionsCreateOrUpdateOrgSecretRequest) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetKeyId

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ActionsCreateOrUpdateOrgSecretRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ActionsCreateOrUpdateOrgSecretRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetVisibility

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ActionsCreateOrUpdateOrgSecretRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetSelectedRepositoryIds

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetSelectedRepositoryIds() []int32`

GetSelectedRepositoryIds returns the SelectedRepositoryIds field if non-nil, zero value otherwise.

### GetSelectedRepositoryIdsOk

`func (o *ActionsCreateOrUpdateOrgSecretRequest) GetSelectedRepositoryIdsOk() (*[]int32, bool)`

GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoryIds

`func (o *ActionsCreateOrUpdateOrgSecretRequest) SetSelectedRepositoryIds(v []int32)`

SetSelectedRepositoryIds sets SelectedRepositoryIds field to given value.

### HasSelectedRepositoryIds

`func (o *ActionsCreateOrUpdateOrgSecretRequest) HasSelectedRepositoryIds() bool`

HasSelectedRepositoryIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


