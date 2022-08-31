# GetConsumedLicenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSeatsConsumed** | Pointer to **int32** |  | [optional] 
**TotalSeatsPurchased** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to [**[]GetConsumedLicensesUsersInner**](GetConsumedLicensesUsersInner.md) |  | [optional] 

## Methods

### NewGetConsumedLicenses

`func NewGetConsumedLicenses() *GetConsumedLicenses`

NewGetConsumedLicenses instantiates a new GetConsumedLicenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsumedLicensesWithDefaults

`func NewGetConsumedLicensesWithDefaults() *GetConsumedLicenses`

NewGetConsumedLicensesWithDefaults instantiates a new GetConsumedLicenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSeatsConsumed

`func (o *GetConsumedLicenses) GetTotalSeatsConsumed() int32`

GetTotalSeatsConsumed returns the TotalSeatsConsumed field if non-nil, zero value otherwise.

### GetTotalSeatsConsumedOk

`func (o *GetConsumedLicenses) GetTotalSeatsConsumedOk() (*int32, bool)`

GetTotalSeatsConsumedOk returns a tuple with the TotalSeatsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeatsConsumed

`func (o *GetConsumedLicenses) SetTotalSeatsConsumed(v int32)`

SetTotalSeatsConsumed sets TotalSeatsConsumed field to given value.

### HasTotalSeatsConsumed

`func (o *GetConsumedLicenses) HasTotalSeatsConsumed() bool`

HasTotalSeatsConsumed returns a boolean if a field has been set.

### GetTotalSeatsPurchased

`func (o *GetConsumedLicenses) GetTotalSeatsPurchased() int32`

GetTotalSeatsPurchased returns the TotalSeatsPurchased field if non-nil, zero value otherwise.

### GetTotalSeatsPurchasedOk

`func (o *GetConsumedLicenses) GetTotalSeatsPurchasedOk() (*int32, bool)`

GetTotalSeatsPurchasedOk returns a tuple with the TotalSeatsPurchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeatsPurchased

`func (o *GetConsumedLicenses) SetTotalSeatsPurchased(v int32)`

SetTotalSeatsPurchased sets TotalSeatsPurchased field to given value.

### HasTotalSeatsPurchased

`func (o *GetConsumedLicenses) HasTotalSeatsPurchased() bool`

HasTotalSeatsPurchased returns a boolean if a field has been set.

### GetUsers

`func (o *GetConsumedLicenses) GetUsers() []GetConsumedLicensesUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetConsumedLicenses) GetUsersOk() (*[]GetConsumedLicensesUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetConsumedLicenses) SetUsers(v []GetConsumedLicensesUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetConsumedLicenses) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


