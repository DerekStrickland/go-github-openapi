# ActionsListRepoSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Secrets** | [**[]ActionsSecret**](ActionsSecret.md) |  | 

## Methods

### NewActionsListRepoSecrets200Response

`func NewActionsListRepoSecrets200Response(totalCount int32, secrets []ActionsSecret, ) *ActionsListRepoSecrets200Response`

NewActionsListRepoSecrets200Response instantiates a new ActionsListRepoSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsListRepoSecrets200ResponseWithDefaults

`func NewActionsListRepoSecrets200ResponseWithDefaults() *ActionsListRepoSecrets200Response`

NewActionsListRepoSecrets200ResponseWithDefaults instantiates a new ActionsListRepoSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsListRepoSecrets200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsListRepoSecrets200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsListRepoSecrets200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetSecrets

`func (o *ActionsListRepoSecrets200Response) GetSecrets() []ActionsSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ActionsListRepoSecrets200Response) GetSecretsOk() (*[]ActionsSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ActionsListRepoSecrets200Response) SetSecrets(v []ActionsSecret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


