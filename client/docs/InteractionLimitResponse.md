# InteractionLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | [**InteractionGroup**](InteractionGroup.md) |  | 
**Origin** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewInteractionLimitResponse

`func NewInteractionLimitResponse(limit InteractionGroup, origin string, expiresAt time.Time, ) *InteractionLimitResponse`

NewInteractionLimitResponse instantiates a new InteractionLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionLimitResponseWithDefaults

`func NewInteractionLimitResponseWithDefaults() *InteractionLimitResponse`

NewInteractionLimitResponseWithDefaults instantiates a new InteractionLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *InteractionLimitResponse) GetLimit() InteractionGroup`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InteractionLimitResponse) GetLimitOk() (*InteractionGroup, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InteractionLimitResponse) SetLimit(v InteractionGroup)`

SetLimit sets Limit field to given value.


### GetOrigin

`func (o *InteractionLimitResponse) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *InteractionLimitResponse) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *InteractionLimitResponse) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetExpiresAt

`func (o *InteractionLimitResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InteractionLimitResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InteractionLimitResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


