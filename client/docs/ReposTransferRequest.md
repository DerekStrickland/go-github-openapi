# ReposTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewOwner** | **string** | The username or organization name the repository will be transferred to. | 
**TeamIds** | Pointer to **[]int32** | ID of the team or teams to add to the repository. Teams can only be added to organization-owned repositories. | [optional] 

## Methods

### NewReposTransferRequest

`func NewReposTransferRequest(newOwner string, ) *ReposTransferRequest`

NewReposTransferRequest instantiates a new ReposTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposTransferRequestWithDefaults

`func NewReposTransferRequestWithDefaults() *ReposTransferRequest`

NewReposTransferRequestWithDefaults instantiates a new ReposTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewOwner

`func (o *ReposTransferRequest) GetNewOwner() string`

GetNewOwner returns the NewOwner field if non-nil, zero value otherwise.

### GetNewOwnerOk

`func (o *ReposTransferRequest) GetNewOwnerOk() (*string, bool)`

GetNewOwnerOk returns a tuple with the NewOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwner

`func (o *ReposTransferRequest) SetNewOwner(v string)`

SetNewOwner sets NewOwner field to given value.


### GetTeamIds

`func (o *ReposTransferRequest) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *ReposTransferRequest) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *ReposTransferRequest) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *ReposTransferRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


