# InteractionsGetRestrictionsForOrg200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | [**InteractionGroup**](InteractionGroup.md) |  | 
**Origin** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewInteractionsGetRestrictionsForOrg200Response

`func NewInteractionsGetRestrictionsForOrg200Response(limit InteractionGroup, origin string, expiresAt time.Time, ) *InteractionsGetRestrictionsForOrg200Response`

NewInteractionsGetRestrictionsForOrg200Response instantiates a new InteractionsGetRestrictionsForOrg200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionsGetRestrictionsForOrg200ResponseWithDefaults

`func NewInteractionsGetRestrictionsForOrg200ResponseWithDefaults() *InteractionsGetRestrictionsForOrg200Response`

NewInteractionsGetRestrictionsForOrg200ResponseWithDefaults instantiates a new InteractionsGetRestrictionsForOrg200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *InteractionsGetRestrictionsForOrg200Response) GetLimit() InteractionGroup`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InteractionsGetRestrictionsForOrg200Response) GetLimitOk() (*InteractionGroup, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InteractionsGetRestrictionsForOrg200Response) SetLimit(v InteractionGroup)`

SetLimit sets Limit field to given value.


### GetOrigin

`func (o *InteractionsGetRestrictionsForOrg200Response) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *InteractionsGetRestrictionsForOrg200Response) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *InteractionsGetRestrictionsForOrg200Response) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetExpiresAt

`func (o *InteractionsGetRestrictionsForOrg200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InteractionsGetRestrictionsForOrg200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InteractionsGetRestrictionsForOrg200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


