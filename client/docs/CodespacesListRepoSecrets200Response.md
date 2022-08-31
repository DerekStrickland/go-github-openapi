# CodespacesListRepoSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Secrets** | [**[]RepoCodespacesSecret**](RepoCodespacesSecret.md) |  | 

## Methods

### NewCodespacesListRepoSecrets200Response

`func NewCodespacesListRepoSecrets200Response(totalCount int32, secrets []RepoCodespacesSecret, ) *CodespacesListRepoSecrets200Response`

NewCodespacesListRepoSecrets200Response instantiates a new CodespacesListRepoSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesListRepoSecrets200ResponseWithDefaults

`func NewCodespacesListRepoSecrets200ResponseWithDefaults() *CodespacesListRepoSecrets200Response`

NewCodespacesListRepoSecrets200ResponseWithDefaults instantiates a new CodespacesListRepoSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CodespacesListRepoSecrets200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CodespacesListRepoSecrets200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CodespacesListRepoSecrets200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetSecrets

`func (o *CodespacesListRepoSecrets200Response) GetSecrets() []RepoCodespacesSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CodespacesListRepoSecrets200Response) GetSecretsOk() (*[]RepoCodespacesSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CodespacesListRepoSecrets200Response) SetSecrets(v []RepoCodespacesSecret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


