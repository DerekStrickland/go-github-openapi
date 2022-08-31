# MarketplacePurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Type** | **string** |  | 
**Id** | **int32** |  | 
**Login** | **string** |  | 
**OrganizationBillingEmail** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**MarketplacePendingChange** | Pointer to [**NullableMarketplacePurchaseMarketplacePendingChange**](MarketplacePurchaseMarketplacePendingChange.md) |  | [optional] 
**MarketplacePurchase** | [**MarketplacePurchaseMarketplacePurchase**](MarketplacePurchaseMarketplacePurchase.md) |  | 

## Methods

### NewMarketplacePurchase

`func NewMarketplacePurchase(url string, type_ string, id int32, login string, marketplacePurchase MarketplacePurchaseMarketplacePurchase, ) *MarketplacePurchase`

NewMarketplacePurchase instantiates a new MarketplacePurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePurchaseWithDefaults

`func NewMarketplacePurchaseWithDefaults() *MarketplacePurchase`

NewMarketplacePurchaseWithDefaults instantiates a new MarketplacePurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *MarketplacePurchase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketplacePurchase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketplacePurchase) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *MarketplacePurchase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplacePurchase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplacePurchase) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MarketplacePurchase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplacePurchase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplacePurchase) SetId(v int32)`

SetId sets Id field to given value.


### GetLogin

`func (o *MarketplacePurchase) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *MarketplacePurchase) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *MarketplacePurchase) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetOrganizationBillingEmail

`func (o *MarketplacePurchase) GetOrganizationBillingEmail() string`

GetOrganizationBillingEmail returns the OrganizationBillingEmail field if non-nil, zero value otherwise.

### GetOrganizationBillingEmailOk

`func (o *MarketplacePurchase) GetOrganizationBillingEmailOk() (*string, bool)`

GetOrganizationBillingEmailOk returns a tuple with the OrganizationBillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationBillingEmail

`func (o *MarketplacePurchase) SetOrganizationBillingEmail(v string)`

SetOrganizationBillingEmail sets OrganizationBillingEmail field to given value.

### HasOrganizationBillingEmail

`func (o *MarketplacePurchase) HasOrganizationBillingEmail() bool`

HasOrganizationBillingEmail returns a boolean if a field has been set.

### GetEmail

`func (o *MarketplacePurchase) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketplacePurchase) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketplacePurchase) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketplacePurchase) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MarketplacePurchase) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MarketplacePurchase) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMarketplacePendingChange

`func (o *MarketplacePurchase) GetMarketplacePendingChange() MarketplacePurchaseMarketplacePendingChange`

GetMarketplacePendingChange returns the MarketplacePendingChange field if non-nil, zero value otherwise.

### GetMarketplacePendingChangeOk

`func (o *MarketplacePurchase) GetMarketplacePendingChangeOk() (*MarketplacePurchaseMarketplacePendingChange, bool)`

GetMarketplacePendingChangeOk returns a tuple with the MarketplacePendingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplacePendingChange

`func (o *MarketplacePurchase) SetMarketplacePendingChange(v MarketplacePurchaseMarketplacePendingChange)`

SetMarketplacePendingChange sets MarketplacePendingChange field to given value.

### HasMarketplacePendingChange

`func (o *MarketplacePurchase) HasMarketplacePendingChange() bool`

HasMarketplacePendingChange returns a boolean if a field has been set.

### SetMarketplacePendingChangeNil

`func (o *MarketplacePurchase) SetMarketplacePendingChangeNil(b bool)`

 SetMarketplacePendingChangeNil sets the value for MarketplacePendingChange to be an explicit nil

### UnsetMarketplacePendingChange
`func (o *MarketplacePurchase) UnsetMarketplacePendingChange()`

UnsetMarketplacePendingChange ensures that no value is present for MarketplacePendingChange, not even an explicit nil
### GetMarketplacePurchase

`func (o *MarketplacePurchase) GetMarketplacePurchase() MarketplacePurchaseMarketplacePurchase`

GetMarketplacePurchase returns the MarketplacePurchase field if non-nil, zero value otherwise.

### GetMarketplacePurchaseOk

`func (o *MarketplacePurchase) GetMarketplacePurchaseOk() (*MarketplacePurchaseMarketplacePurchase, bool)`

GetMarketplacePurchaseOk returns a tuple with the MarketplacePurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplacePurchase

`func (o *MarketplacePurchase) SetMarketplacePurchase(v MarketplacePurchaseMarketplacePurchase)`

SetMarketplacePurchase sets MarketplacePurchase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


