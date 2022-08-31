# OrganizationFullPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Space** | **int32** |  | 
**PrivateRepos** | **int32** |  | 
**FilledSeats** | Pointer to **int32** |  | [optional] 
**Seats** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationFullPlan

`func NewOrganizationFullPlan(name string, space int32, privateRepos int32, ) *OrganizationFullPlan`

NewOrganizationFullPlan instantiates a new OrganizationFullPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFullPlanWithDefaults

`func NewOrganizationFullPlanWithDefaults() *OrganizationFullPlan`

NewOrganizationFullPlanWithDefaults instantiates a new OrganizationFullPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationFullPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationFullPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationFullPlan) SetName(v string)`

SetName sets Name field to given value.


### GetSpace

`func (o *OrganizationFullPlan) GetSpace() int32`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *OrganizationFullPlan) GetSpaceOk() (*int32, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *OrganizationFullPlan) SetSpace(v int32)`

SetSpace sets Space field to given value.


### GetPrivateRepos

`func (o *OrganizationFullPlan) GetPrivateRepos() int32`

GetPrivateRepos returns the PrivateRepos field if non-nil, zero value otherwise.

### GetPrivateReposOk

`func (o *OrganizationFullPlan) GetPrivateReposOk() (*int32, bool)`

GetPrivateReposOk returns a tuple with the PrivateRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateRepos

`func (o *OrganizationFullPlan) SetPrivateRepos(v int32)`

SetPrivateRepos sets PrivateRepos field to given value.


### GetFilledSeats

`func (o *OrganizationFullPlan) GetFilledSeats() int32`

GetFilledSeats returns the FilledSeats field if non-nil, zero value otherwise.

### GetFilledSeatsOk

`func (o *OrganizationFullPlan) GetFilledSeatsOk() (*int32, bool)`

GetFilledSeatsOk returns a tuple with the FilledSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledSeats

`func (o *OrganizationFullPlan) SetFilledSeats(v int32)`

SetFilledSeats sets FilledSeats field to given value.

### HasFilledSeats

`func (o *OrganizationFullPlan) HasFilledSeats() bool`

HasFilledSeats returns a boolean if a field has been set.

### GetSeats

`func (o *OrganizationFullPlan) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *OrganizationFullPlan) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *OrganizationFullPlan) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *OrganizationFullPlan) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


