# HookDeliveryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the webhook delivery. | 
**Guid** | **string** | Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event). | 
**DeliveredAt** | **time.Time** | Time when the webhook delivery occurred. | 
**Redelivery** | **bool** | Whether the webhook delivery is a redelivery. | 
**Duration** | **float32** | Time spent delivering. | 
**Status** | **string** | Describes the response returned after attempting the delivery. | 
**StatusCode** | **int32** | Status code received when delivery was made. | 
**Event** | **string** | The event that triggered the delivery. | 
**Action** | **NullableString** | The type of activity for the event that triggered the delivery. | 
**InstallationId** | **NullableInt32** | The id of the GitHub App installation associated with this event. | 
**RepositoryId** | **NullableInt32** | The id of the repository associated with this event. | 

## Methods

### NewHookDeliveryItem

`func NewHookDeliveryItem(id int32, guid string, deliveredAt time.Time, redelivery bool, duration float32, status string, statusCode int32, event string, action NullableString, installationId NullableInt32, repositoryId NullableInt32, ) *HookDeliveryItem`

NewHookDeliveryItem instantiates a new HookDeliveryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookDeliveryItemWithDefaults

`func NewHookDeliveryItemWithDefaults() *HookDeliveryItem`

NewHookDeliveryItemWithDefaults instantiates a new HookDeliveryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HookDeliveryItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HookDeliveryItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HookDeliveryItem) SetId(v int32)`

SetId sets Id field to given value.


### GetGuid

`func (o *HookDeliveryItem) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *HookDeliveryItem) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *HookDeliveryItem) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetDeliveredAt

`func (o *HookDeliveryItem) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *HookDeliveryItem) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *HookDeliveryItem) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.


### GetRedelivery

`func (o *HookDeliveryItem) GetRedelivery() bool`

GetRedelivery returns the Redelivery field if non-nil, zero value otherwise.

### GetRedeliveryOk

`func (o *HookDeliveryItem) GetRedeliveryOk() (*bool, bool)`

GetRedeliveryOk returns a tuple with the Redelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedelivery

`func (o *HookDeliveryItem) SetRedelivery(v bool)`

SetRedelivery sets Redelivery field to given value.


### GetDuration

`func (o *HookDeliveryItem) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HookDeliveryItem) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HookDeliveryItem) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetStatus

`func (o *HookDeliveryItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HookDeliveryItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HookDeliveryItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusCode

`func (o *HookDeliveryItem) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HookDeliveryItem) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HookDeliveryItem) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetEvent

`func (o *HookDeliveryItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *HookDeliveryItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *HookDeliveryItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAction

`func (o *HookDeliveryItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HookDeliveryItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HookDeliveryItem) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *HookDeliveryItem) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *HookDeliveryItem) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetInstallationId

`func (o *HookDeliveryItem) GetInstallationId() int32`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *HookDeliveryItem) GetInstallationIdOk() (*int32, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *HookDeliveryItem) SetInstallationId(v int32)`

SetInstallationId sets InstallationId field to given value.


### SetInstallationIdNil

`func (o *HookDeliveryItem) SetInstallationIdNil(b bool)`

 SetInstallationIdNil sets the value for InstallationId to be an explicit nil

### UnsetInstallationId
`func (o *HookDeliveryItem) UnsetInstallationId()`

UnsetInstallationId ensures that no value is present for InstallationId, not even an explicit nil
### GetRepositoryId

`func (o *HookDeliveryItem) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *HookDeliveryItem) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *HookDeliveryItem) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.


### SetRepositoryIdNil

`func (o *HookDeliveryItem) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *HookDeliveryItem) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


