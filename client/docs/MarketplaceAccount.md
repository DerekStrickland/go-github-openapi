# MarketplaceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Id** | **int32** |  | 
**Type** | **string** |  | 
**NodeId** | Pointer to **string** |  | [optional] 
**Login** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**OrganizationBillingEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMarketplaceAccount

`func NewMarketplaceAccount(url string, id int32, type_ string, login string, ) *MarketplaceAccount`

NewMarketplaceAccount instantiates a new MarketplaceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceAccountWithDefaults

`func NewMarketplaceAccountWithDefaults() *MarketplaceAccount`

NewMarketplaceAccountWithDefaults instantiates a new MarketplaceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *MarketplaceAccount) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketplaceAccount) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketplaceAccount) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *MarketplaceAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceAccount) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *MarketplaceAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceAccount) SetType(v string)`

SetType sets Type field to given value.


### GetNodeId

`func (o *MarketplaceAccount) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *MarketplaceAccount) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *MarketplaceAccount) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *MarketplaceAccount) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetLogin

`func (o *MarketplaceAccount) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *MarketplaceAccount) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *MarketplaceAccount) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetEmail

`func (o *MarketplaceAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketplaceAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketplaceAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketplaceAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MarketplaceAccount) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MarketplaceAccount) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOrganizationBillingEmail

`func (o *MarketplaceAccount) GetOrganizationBillingEmail() string`

GetOrganizationBillingEmail returns the OrganizationBillingEmail field if non-nil, zero value otherwise.

### GetOrganizationBillingEmailOk

`func (o *MarketplaceAccount) GetOrganizationBillingEmailOk() (*string, bool)`

GetOrganizationBillingEmailOk returns a tuple with the OrganizationBillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationBillingEmail

`func (o *MarketplaceAccount) SetOrganizationBillingEmail(v string)`

SetOrganizationBillingEmail sets OrganizationBillingEmail field to given value.

### HasOrganizationBillingEmail

`func (o *MarketplaceAccount) HasOrganizationBillingEmail() bool`

HasOrganizationBillingEmail returns a boolean if a field has been set.

### SetOrganizationBillingEmailNil

`func (o *MarketplaceAccount) SetOrganizationBillingEmailNil(b bool)`

 SetOrganizationBillingEmailNil sets the value for OrganizationBillingEmail to be an explicit nil

### UnsetOrganizationBillingEmail
`func (o *MarketplaceAccount) UnsetOrganizationBillingEmail()`

UnsetOrganizationBillingEmail ensures that no value is present for OrganizationBillingEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


