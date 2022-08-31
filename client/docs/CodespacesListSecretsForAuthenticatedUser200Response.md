# CodespacesListSecretsForAuthenticatedUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Secrets** | [**[]CodespacesSecret**](CodespacesSecret.md) |  | 

## Methods

### NewCodespacesListSecretsForAuthenticatedUser200Response

`func NewCodespacesListSecretsForAuthenticatedUser200Response(totalCount int32, secrets []CodespacesSecret, ) *CodespacesListSecretsForAuthenticatedUser200Response`

NewCodespacesListSecretsForAuthenticatedUser200Response instantiates a new CodespacesListSecretsForAuthenticatedUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesListSecretsForAuthenticatedUser200ResponseWithDefaults

`func NewCodespacesListSecretsForAuthenticatedUser200ResponseWithDefaults() *CodespacesListSecretsForAuthenticatedUser200Response`

NewCodespacesListSecretsForAuthenticatedUser200ResponseWithDefaults instantiates a new CodespacesListSecretsForAuthenticatedUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CodespacesListSecretsForAuthenticatedUser200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CodespacesListSecretsForAuthenticatedUser200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CodespacesListSecretsForAuthenticatedUser200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetSecrets

`func (o *CodespacesListSecretsForAuthenticatedUser200Response) GetSecrets() []CodespacesSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CodespacesListSecretsForAuthenticatedUser200Response) GetSecretsOk() (*[]CodespacesSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CodespacesListSecretsForAuthenticatedUser200Response) SetSecrets(v []CodespacesSecret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


