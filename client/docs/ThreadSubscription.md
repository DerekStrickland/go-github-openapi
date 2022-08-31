# ThreadSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribed** | **bool** |  | 
**Ignored** | **bool** |  | 
**Reason** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 
**Url** | **string** |  | 
**ThreadUrl** | Pointer to **string** |  | [optional] 
**RepositoryUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewThreadSubscription

`func NewThreadSubscription(subscribed bool, ignored bool, reason NullableString, createdAt NullableTime, url string, ) *ThreadSubscription`

NewThreadSubscription instantiates a new ThreadSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadSubscriptionWithDefaults

`func NewThreadSubscriptionWithDefaults() *ThreadSubscription`

NewThreadSubscriptionWithDefaults instantiates a new ThreadSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribed

`func (o *ThreadSubscription) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *ThreadSubscription) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *ThreadSubscription) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.


### GetIgnored

`func (o *ThreadSubscription) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ThreadSubscription) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ThreadSubscription) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.


### GetReason

`func (o *ThreadSubscription) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ThreadSubscription) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ThreadSubscription) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ThreadSubscription) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ThreadSubscription) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetCreatedAt

`func (o *ThreadSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ThreadSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ThreadSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ThreadSubscription) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ThreadSubscription) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUrl

`func (o *ThreadSubscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThreadSubscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThreadSubscription) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetThreadUrl

`func (o *ThreadSubscription) GetThreadUrl() string`

GetThreadUrl returns the ThreadUrl field if non-nil, zero value otherwise.

### GetThreadUrlOk

`func (o *ThreadSubscription) GetThreadUrlOk() (*string, bool)`

GetThreadUrlOk returns a tuple with the ThreadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadUrl

`func (o *ThreadSubscription) SetThreadUrl(v string)`

SetThreadUrl sets ThreadUrl field to given value.

### HasThreadUrl

`func (o *ThreadSubscription) HasThreadUrl() bool`

HasThreadUrl returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *ThreadSubscription) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *ThreadSubscription) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *ThreadSubscription) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *ThreadSubscription) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


