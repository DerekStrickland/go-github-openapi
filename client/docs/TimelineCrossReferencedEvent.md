# TimelineCrossReferencedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Actor** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Source** | [**TimelineCrossReferencedEventSource**](TimelineCrossReferencedEventSource.md) |  | 

## Methods

### NewTimelineCrossReferencedEvent

`func NewTimelineCrossReferencedEvent(event string, createdAt time.Time, updatedAt time.Time, source TimelineCrossReferencedEventSource, ) *TimelineCrossReferencedEvent`

NewTimelineCrossReferencedEvent instantiates a new TimelineCrossReferencedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineCrossReferencedEventWithDefaults

`func NewTimelineCrossReferencedEventWithDefaults() *TimelineCrossReferencedEvent`

NewTimelineCrossReferencedEventWithDefaults instantiates a new TimelineCrossReferencedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TimelineCrossReferencedEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineCrossReferencedEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineCrossReferencedEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetActor

`func (o *TimelineCrossReferencedEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TimelineCrossReferencedEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TimelineCrossReferencedEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.

### HasActor

`func (o *TimelineCrossReferencedEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TimelineCrossReferencedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineCrossReferencedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineCrossReferencedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TimelineCrossReferencedEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TimelineCrossReferencedEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TimelineCrossReferencedEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSource

`func (o *TimelineCrossReferencedEvent) GetSource() TimelineCrossReferencedEventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TimelineCrossReferencedEvent) GetSourceOk() (*TimelineCrossReferencedEventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TimelineCrossReferencedEvent) SetSource(v TimelineCrossReferencedEventSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


