# CodespacesRepoMachinesForAuthenticatedUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Machines** | [**[]CodespaceMachine**](CodespaceMachine.md) |  | 

## Methods

### NewCodespacesRepoMachinesForAuthenticatedUser200Response

`func NewCodespacesRepoMachinesForAuthenticatedUser200Response(totalCount int32, machines []CodespaceMachine, ) *CodespacesRepoMachinesForAuthenticatedUser200Response`

NewCodespacesRepoMachinesForAuthenticatedUser200Response instantiates a new CodespacesRepoMachinesForAuthenticatedUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesRepoMachinesForAuthenticatedUser200ResponseWithDefaults

`func NewCodespacesRepoMachinesForAuthenticatedUser200ResponseWithDefaults() *CodespacesRepoMachinesForAuthenticatedUser200Response`

NewCodespacesRepoMachinesForAuthenticatedUser200ResponseWithDefaults instantiates a new CodespacesRepoMachinesForAuthenticatedUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CodespacesRepoMachinesForAuthenticatedUser200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CodespacesRepoMachinesForAuthenticatedUser200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CodespacesRepoMachinesForAuthenticatedUser200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetMachines

`func (o *CodespacesRepoMachinesForAuthenticatedUser200Response) GetMachines() []CodespaceMachine`

GetMachines returns the Machines field if non-nil, zero value otherwise.

### GetMachinesOk

`func (o *CodespacesRepoMachinesForAuthenticatedUser200Response) GetMachinesOk() (*[]CodespaceMachine, bool)`

GetMachinesOk returns a tuple with the Machines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachines

`func (o *CodespacesRepoMachinesForAuthenticatedUser200Response) SetMachines(v []CodespaceMachine)`

SetMachines sets Machines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


