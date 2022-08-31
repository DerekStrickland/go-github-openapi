# InteractionLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | [**InteractionGroup**](InteractionGroup.md) |  | 
**Expiry** | Pointer to [**InteractionExpiry**](InteractionExpiry.md) |  | [optional] 

## Methods

### NewInteractionLimit

`func NewInteractionLimit(limit InteractionGroup, ) *InteractionLimit`

NewInteractionLimit instantiates a new InteractionLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionLimitWithDefaults

`func NewInteractionLimitWithDefaults() *InteractionLimit`

NewInteractionLimitWithDefaults instantiates a new InteractionLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *InteractionLimit) GetLimit() InteractionGroup`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InteractionLimit) GetLimitOk() (*InteractionGroup, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InteractionLimit) SetLimit(v InteractionGroup)`

SetLimit sets Limit field to given value.


### GetExpiry

`func (o *InteractionLimit) GetExpiry() InteractionExpiry`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *InteractionLimit) GetExpiryOk() (*InteractionExpiry, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *InteractionLimit) SetExpiry(v InteractionExpiry)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *InteractionLimit) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


