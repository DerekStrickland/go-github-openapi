# Thread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**Subject** | [**ThreadSubject**](ThreadSubject.md) |  | 
**Reason** | **string** |  | 
**Unread** | **bool** |  | 
**UpdatedAt** | **string** |  | 
**LastReadAt** | **NullableString** |  | 
**Url** | **string** |  | 
**SubscriptionUrl** | **string** |  | 

## Methods

### NewThread

`func NewThread(id string, repository MinimalRepository, subject ThreadSubject, reason string, unread bool, updatedAt string, lastReadAt NullableString, url string, subscriptionUrl string, ) *Thread`

NewThread instantiates a new Thread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadWithDefaults

`func NewThreadWithDefaults() *Thread`

NewThreadWithDefaults instantiates a new Thread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Thread) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Thread) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Thread) SetId(v string)`

SetId sets Id field to given value.


### GetRepository

`func (o *Thread) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Thread) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Thread) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetSubject

`func (o *Thread) GetSubject() ThreadSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Thread) GetSubjectOk() (*ThreadSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Thread) SetSubject(v ThreadSubject)`

SetSubject sets Subject field to given value.


### GetReason

`func (o *Thread) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Thread) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Thread) SetReason(v string)`

SetReason sets Reason field to given value.


### GetUnread

`func (o *Thread) GetUnread() bool`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *Thread) GetUnreadOk() (*bool, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *Thread) SetUnread(v bool)`

SetUnread sets Unread field to given value.


### GetUpdatedAt

`func (o *Thread) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Thread) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Thread) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastReadAt

`func (o *Thread) GetLastReadAt() string`

GetLastReadAt returns the LastReadAt field if non-nil, zero value otherwise.

### GetLastReadAtOk

`func (o *Thread) GetLastReadAtOk() (*string, bool)`

GetLastReadAtOk returns a tuple with the LastReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadAt

`func (o *Thread) SetLastReadAt(v string)`

SetLastReadAt sets LastReadAt field to given value.


### SetLastReadAtNil

`func (o *Thread) SetLastReadAtNil(b bool)`

 SetLastReadAtNil sets the value for LastReadAt to be an explicit nil

### UnsetLastReadAt
`func (o *Thread) UnsetLastReadAt()`

UnsetLastReadAt ensures that no value is present for LastReadAt, not even an explicit nil
### GetUrl

`func (o *Thread) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Thread) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Thread) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSubscriptionUrl

`func (o *Thread) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *Thread) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *Thread) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


