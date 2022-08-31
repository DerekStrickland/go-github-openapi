# ActivitySetRepoSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribed** | Pointer to **bool** | Determines if notifications should be received from this repository. | [optional] 
**Ignored** | Pointer to **bool** | Determines if all notifications should be blocked from this repository. | [optional] 

## Methods

### NewActivitySetRepoSubscriptionRequest

`func NewActivitySetRepoSubscriptionRequest() *ActivitySetRepoSubscriptionRequest`

NewActivitySetRepoSubscriptionRequest instantiates a new ActivitySetRepoSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivitySetRepoSubscriptionRequestWithDefaults

`func NewActivitySetRepoSubscriptionRequestWithDefaults() *ActivitySetRepoSubscriptionRequest`

NewActivitySetRepoSubscriptionRequestWithDefaults instantiates a new ActivitySetRepoSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribed

`func (o *ActivitySetRepoSubscriptionRequest) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *ActivitySetRepoSubscriptionRequest) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *ActivitySetRepoSubscriptionRequest) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *ActivitySetRepoSubscriptionRequest) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.

### GetIgnored

`func (o *ActivitySetRepoSubscriptionRequest) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ActivitySetRepoSubscriptionRequest) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ActivitySetRepoSubscriptionRequest) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ActivitySetRepoSubscriptionRequest) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


