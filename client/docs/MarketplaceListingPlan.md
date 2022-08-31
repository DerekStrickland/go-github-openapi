# MarketplaceListingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**AccountsUrl** | **string** |  | 
**Id** | **int32** |  | 
**Number** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**MonthlyPriceInCents** | **int32** |  | 
**YearlyPriceInCents** | **int32** |  | 
**PriceModel** | **string** |  | 
**HasFreeTrial** | **bool** |  | 
**UnitName** | **NullableString** |  | 
**State** | **string** |  | 
**Bullets** | **[]string** |  | 

## Methods

### NewMarketplaceListingPlan

`func NewMarketplaceListingPlan(url string, accountsUrl string, id int32, number int32, name string, description string, monthlyPriceInCents int32, yearlyPriceInCents int32, priceModel string, hasFreeTrial bool, unitName NullableString, state string, bullets []string, ) *MarketplaceListingPlan`

NewMarketplaceListingPlan instantiates a new MarketplaceListingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceListingPlanWithDefaults

`func NewMarketplaceListingPlanWithDefaults() *MarketplaceListingPlan`

NewMarketplaceListingPlanWithDefaults instantiates a new MarketplaceListingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *MarketplaceListingPlan) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketplaceListingPlan) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketplaceListingPlan) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAccountsUrl

`func (o *MarketplaceListingPlan) GetAccountsUrl() string`

GetAccountsUrl returns the AccountsUrl field if non-nil, zero value otherwise.

### GetAccountsUrlOk

`func (o *MarketplaceListingPlan) GetAccountsUrlOk() (*string, bool)`

GetAccountsUrlOk returns a tuple with the AccountsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsUrl

`func (o *MarketplaceListingPlan) SetAccountsUrl(v string)`

SetAccountsUrl sets AccountsUrl field to given value.


### GetId

`func (o *MarketplaceListingPlan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceListingPlan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceListingPlan) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *MarketplaceListingPlan) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *MarketplaceListingPlan) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *MarketplaceListingPlan) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetName

`func (o *MarketplaceListingPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceListingPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceListingPlan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MarketplaceListingPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceListingPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceListingPlan) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMonthlyPriceInCents

`func (o *MarketplaceListingPlan) GetMonthlyPriceInCents() int32`

GetMonthlyPriceInCents returns the MonthlyPriceInCents field if non-nil, zero value otherwise.

### GetMonthlyPriceInCentsOk

`func (o *MarketplaceListingPlan) GetMonthlyPriceInCentsOk() (*int32, bool)`

GetMonthlyPriceInCentsOk returns a tuple with the MonthlyPriceInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPriceInCents

`func (o *MarketplaceListingPlan) SetMonthlyPriceInCents(v int32)`

SetMonthlyPriceInCents sets MonthlyPriceInCents field to given value.


### GetYearlyPriceInCents

`func (o *MarketplaceListingPlan) GetYearlyPriceInCents() int32`

GetYearlyPriceInCents returns the YearlyPriceInCents field if non-nil, zero value otherwise.

### GetYearlyPriceInCentsOk

`func (o *MarketplaceListingPlan) GetYearlyPriceInCentsOk() (*int32, bool)`

GetYearlyPriceInCentsOk returns a tuple with the YearlyPriceInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPriceInCents

`func (o *MarketplaceListingPlan) SetYearlyPriceInCents(v int32)`

SetYearlyPriceInCents sets YearlyPriceInCents field to given value.


### GetPriceModel

`func (o *MarketplaceListingPlan) GetPriceModel() string`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *MarketplaceListingPlan) GetPriceModelOk() (*string, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *MarketplaceListingPlan) SetPriceModel(v string)`

SetPriceModel sets PriceModel field to given value.


### GetHasFreeTrial

`func (o *MarketplaceListingPlan) GetHasFreeTrial() bool`

GetHasFreeTrial returns the HasFreeTrial field if non-nil, zero value otherwise.

### GetHasFreeTrialOk

`func (o *MarketplaceListingPlan) GetHasFreeTrialOk() (*bool, bool)`

GetHasFreeTrialOk returns a tuple with the HasFreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFreeTrial

`func (o *MarketplaceListingPlan) SetHasFreeTrial(v bool)`

SetHasFreeTrial sets HasFreeTrial field to given value.


### GetUnitName

`func (o *MarketplaceListingPlan) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *MarketplaceListingPlan) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *MarketplaceListingPlan) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### SetUnitNameNil

`func (o *MarketplaceListingPlan) SetUnitNameNil(b bool)`

 SetUnitNameNil sets the value for UnitName to be an explicit nil

### UnsetUnitName
`func (o *MarketplaceListingPlan) UnsetUnitName()`

UnsetUnitName ensures that no value is present for UnitName, not even an explicit nil
### GetState

`func (o *MarketplaceListingPlan) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MarketplaceListingPlan) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MarketplaceListingPlan) SetState(v string)`

SetState sets State field to given value.


### GetBullets

`func (o *MarketplaceListingPlan) GetBullets() []string`

GetBullets returns the Bullets field if non-nil, zero value otherwise.

### GetBulletsOk

`func (o *MarketplaceListingPlan) GetBulletsOk() (*[]string, bool)`

GetBulletsOk returns a tuple with the Bullets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBullets

`func (o *MarketplaceListingPlan) SetBullets(v []string)`

SetBullets sets Bullets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


