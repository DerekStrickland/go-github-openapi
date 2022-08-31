# ActivitySetThreadSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ignored** | Pointer to **bool** | Whether to block all notifications from a thread. | [optional] [default to false]

## Methods

### NewActivitySetThreadSubscriptionRequest

`func NewActivitySetThreadSubscriptionRequest() *ActivitySetThreadSubscriptionRequest`

NewActivitySetThreadSubscriptionRequest instantiates a new ActivitySetThreadSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivitySetThreadSubscriptionRequestWithDefaults

`func NewActivitySetThreadSubscriptionRequestWithDefaults() *ActivitySetThreadSubscriptionRequest`

NewActivitySetThreadSubscriptionRequestWithDefaults instantiates a new ActivitySetThreadSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnored

`func (o *ActivitySetThreadSubscriptionRequest) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ActivitySetThreadSubscriptionRequest) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ActivitySetThreadSubscriptionRequest) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ActivitySetThreadSubscriptionRequest) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


