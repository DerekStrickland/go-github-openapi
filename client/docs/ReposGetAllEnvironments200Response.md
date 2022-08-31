# ReposGetAllEnvironments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The number of environments in this repository | [optional] 
**Environments** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 

## Methods

### NewReposGetAllEnvironments200Response

`func NewReposGetAllEnvironments200Response() *ReposGetAllEnvironments200Response`

NewReposGetAllEnvironments200Response instantiates a new ReposGetAllEnvironments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposGetAllEnvironments200ResponseWithDefaults

`func NewReposGetAllEnvironments200ResponseWithDefaults() *ReposGetAllEnvironments200Response`

NewReposGetAllEnvironments200ResponseWithDefaults instantiates a new ReposGetAllEnvironments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ReposGetAllEnvironments200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReposGetAllEnvironments200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReposGetAllEnvironments200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ReposGetAllEnvironments200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetEnvironments

`func (o *ReposGetAllEnvironments200Response) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ReposGetAllEnvironments200Response) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ReposGetAllEnvironments200Response) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ReposGetAllEnvironments200Response) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


