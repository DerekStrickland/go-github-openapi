# UserMarketplacePurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycle** | **string** |  | 
**NextBillingDate** | **NullableTime** |  | 
**UnitCount** | **NullableInt32** |  | 
**OnFreeTrial** | **bool** |  | 
**FreeTrialEndsOn** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 
**Account** | [**MarketplaceAccount**](MarketplaceAccount.md) |  | 
**Plan** | [**MarketplaceListingPlan**](MarketplaceListingPlan.md) |  | 

## Methods

### NewUserMarketplacePurchase

`func NewUserMarketplacePurchase(billingCycle string, nextBillingDate NullableTime, unitCount NullableInt32, onFreeTrial bool, freeTrialEndsOn NullableTime, updatedAt NullableTime, account MarketplaceAccount, plan MarketplaceListingPlan, ) *UserMarketplacePurchase`

NewUserMarketplacePurchase instantiates a new UserMarketplacePurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMarketplacePurchaseWithDefaults

`func NewUserMarketplacePurchaseWithDefaults() *UserMarketplacePurchase`

NewUserMarketplacePurchaseWithDefaults instantiates a new UserMarketplacePurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycle

`func (o *UserMarketplacePurchase) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *UserMarketplacePurchase) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *UserMarketplacePurchase) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetNextBillingDate

`func (o *UserMarketplacePurchase) GetNextBillingDate() time.Time`

GetNextBillingDate returns the NextBillingDate field if non-nil, zero value otherwise.

### GetNextBillingDateOk

`func (o *UserMarketplacePurchase) GetNextBillingDateOk() (*time.Time, bool)`

GetNextBillingDateOk returns a tuple with the NextBillingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBillingDate

`func (o *UserMarketplacePurchase) SetNextBillingDate(v time.Time)`

SetNextBillingDate sets NextBillingDate field to given value.


### SetNextBillingDateNil

`func (o *UserMarketplacePurchase) SetNextBillingDateNil(b bool)`

 SetNextBillingDateNil sets the value for NextBillingDate to be an explicit nil

### UnsetNextBillingDate
`func (o *UserMarketplacePurchase) UnsetNextBillingDate()`

UnsetNextBillingDate ensures that no value is present for NextBillingDate, not even an explicit nil
### GetUnitCount

`func (o *UserMarketplacePurchase) GetUnitCount() int32`

GetUnitCount returns the UnitCount field if non-nil, zero value otherwise.

### GetUnitCountOk

`func (o *UserMarketplacePurchase) GetUnitCountOk() (*int32, bool)`

GetUnitCountOk returns a tuple with the UnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCount

`func (o *UserMarketplacePurchase) SetUnitCount(v int32)`

SetUnitCount sets UnitCount field to given value.


### SetUnitCountNil

`func (o *UserMarketplacePurchase) SetUnitCountNil(b bool)`

 SetUnitCountNil sets the value for UnitCount to be an explicit nil

### UnsetUnitCount
`func (o *UserMarketplacePurchase) UnsetUnitCount()`

UnsetUnitCount ensures that no value is present for UnitCount, not even an explicit nil
### GetOnFreeTrial

`func (o *UserMarketplacePurchase) GetOnFreeTrial() bool`

GetOnFreeTrial returns the OnFreeTrial field if non-nil, zero value otherwise.

### GetOnFreeTrialOk

`func (o *UserMarketplacePurchase) GetOnFreeTrialOk() (*bool, bool)`

GetOnFreeTrialOk returns a tuple with the OnFreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFreeTrial

`func (o *UserMarketplacePurchase) SetOnFreeTrial(v bool)`

SetOnFreeTrial sets OnFreeTrial field to given value.


### GetFreeTrialEndsOn

`func (o *UserMarketplacePurchase) GetFreeTrialEndsOn() time.Time`

GetFreeTrialEndsOn returns the FreeTrialEndsOn field if non-nil, zero value otherwise.

### GetFreeTrialEndsOnOk

`func (o *UserMarketplacePurchase) GetFreeTrialEndsOnOk() (*time.Time, bool)`

GetFreeTrialEndsOnOk returns a tuple with the FreeTrialEndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialEndsOn

`func (o *UserMarketplacePurchase) SetFreeTrialEndsOn(v time.Time)`

SetFreeTrialEndsOn sets FreeTrialEndsOn field to given value.


### SetFreeTrialEndsOnNil

`func (o *UserMarketplacePurchase) SetFreeTrialEndsOnNil(b bool)`

 SetFreeTrialEndsOnNil sets the value for FreeTrialEndsOn to be an explicit nil

### UnsetFreeTrialEndsOn
`func (o *UserMarketplacePurchase) UnsetFreeTrialEndsOn()`

UnsetFreeTrialEndsOn ensures that no value is present for FreeTrialEndsOn, not even an explicit nil
### GetUpdatedAt

`func (o *UserMarketplacePurchase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserMarketplacePurchase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserMarketplacePurchase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *UserMarketplacePurchase) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UserMarketplacePurchase) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAccount

`func (o *UserMarketplacePurchase) GetAccount() MarketplaceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserMarketplacePurchase) GetAccountOk() (*MarketplaceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserMarketplacePurchase) SetAccount(v MarketplaceAccount)`

SetAccount sets Account field to given value.


### GetPlan

`func (o *UserMarketplacePurchase) GetPlan() MarketplaceListingPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UserMarketplacePurchase) GetPlanOk() (*MarketplaceListingPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UserMarketplacePurchase) SetPlan(v MarketplaceListingPlan)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


