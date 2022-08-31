# DependabotListRepoSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Secrets** | [**[]DependabotSecret**](DependabotSecret.md) |  | 

## Methods

### NewDependabotListRepoSecrets200Response

`func NewDependabotListRepoSecrets200Response(totalCount int32, secrets []DependabotSecret, ) *DependabotListRepoSecrets200Response`

NewDependabotListRepoSecrets200Response instantiates a new DependabotListRepoSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependabotListRepoSecrets200ResponseWithDefaults

`func NewDependabotListRepoSecrets200ResponseWithDefaults() *DependabotListRepoSecrets200Response`

NewDependabotListRepoSecrets200ResponseWithDefaults instantiates a new DependabotListRepoSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *DependabotListRepoSecrets200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DependabotListRepoSecrets200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DependabotListRepoSecrets200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetSecrets

`func (o *DependabotListRepoSecrets200Response) GetSecrets() []DependabotSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *DependabotListRepoSecrets200Response) GetSecretsOk() (*[]DependabotSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *DependabotListRepoSecrets200Response) SetSecrets(v []DependabotSecret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


