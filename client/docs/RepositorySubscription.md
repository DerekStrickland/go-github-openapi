# RepositorySubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribed** | **bool** | Determines if notifications should be received from this repository. | 
**Ignored** | **bool** | Determines if all notifications should be blocked from this repository. | 
**Reason** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**RepositoryUrl** | **string** |  | 

## Methods

### NewRepositorySubscription

`func NewRepositorySubscription(subscribed bool, ignored bool, reason NullableString, createdAt time.Time, url string, repositoryUrl string, ) *RepositorySubscription`

NewRepositorySubscription instantiates a new RepositorySubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySubscriptionWithDefaults

`func NewRepositorySubscriptionWithDefaults() *RepositorySubscription`

NewRepositorySubscriptionWithDefaults instantiates a new RepositorySubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribed

`func (o *RepositorySubscription) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *RepositorySubscription) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *RepositorySubscription) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.


### GetIgnored

`func (o *RepositorySubscription) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *RepositorySubscription) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *RepositorySubscription) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.


### GetReason

`func (o *RepositorySubscription) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RepositorySubscription) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RepositorySubscription) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *RepositorySubscription) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *RepositorySubscription) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetCreatedAt

`func (o *RepositorySubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepositorySubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepositorySubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUrl

`func (o *RepositorySubscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositorySubscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositorySubscription) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRepositoryUrl

`func (o *RepositorySubscription) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *RepositorySubscription) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *RepositorySubscription) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


