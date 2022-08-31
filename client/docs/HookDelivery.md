# HookDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier of the delivery. | 
**Guid** | **string** | Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event). | 
**DeliveredAt** | **time.Time** | Time when the delivery was delivered. | 
**Redelivery** | **bool** | Whether the delivery is a redelivery. | 
**Duration** | **float32** | Time spent delivering. | 
**Status** | **string** | Description of the status of the attempted delivery | 
**StatusCode** | **int32** | Status code received when delivery was made. | 
**Event** | **string** | The event that triggered the delivery. | 
**Action** | **NullableString** | The type of activity for the event that triggered the delivery. | 
**InstallationId** | **NullableInt32** | The id of the GitHub App installation associated with this event. | 
**RepositoryId** | **NullableInt32** | The id of the repository associated with this event. | 
**Url** | Pointer to **string** | The URL target of the delivery. | [optional] 
**Request** | [**HookDeliveryRequest**](HookDeliveryRequest.md) |  | 
**Response** | [**HookDeliveryResponse**](HookDeliveryResponse.md) |  | 

## Methods

### NewHookDelivery

`func NewHookDelivery(id int32, guid string, deliveredAt time.Time, redelivery bool, duration float32, status string, statusCode int32, event string, action NullableString, installationId NullableInt32, repositoryId NullableInt32, request HookDeliveryRequest, response HookDeliveryResponse, ) *HookDelivery`

NewHookDelivery instantiates a new HookDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookDeliveryWithDefaults

`func NewHookDeliveryWithDefaults() *HookDelivery`

NewHookDeliveryWithDefaults instantiates a new HookDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HookDelivery) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HookDelivery) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HookDelivery) SetId(v int32)`

SetId sets Id field to given value.


### GetGuid

`func (o *HookDelivery) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *HookDelivery) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *HookDelivery) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetDeliveredAt

`func (o *HookDelivery) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *HookDelivery) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *HookDelivery) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.


### GetRedelivery

`func (o *HookDelivery) GetRedelivery() bool`

GetRedelivery returns the Redelivery field if non-nil, zero value otherwise.

### GetRedeliveryOk

`func (o *HookDelivery) GetRedeliveryOk() (*bool, bool)`

GetRedeliveryOk returns a tuple with the Redelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedelivery

`func (o *HookDelivery) SetRedelivery(v bool)`

SetRedelivery sets Redelivery field to given value.


### GetDuration

`func (o *HookDelivery) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HookDelivery) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HookDelivery) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetStatus

`func (o *HookDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HookDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HookDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusCode

`func (o *HookDelivery) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HookDelivery) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HookDelivery) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetEvent

`func (o *HookDelivery) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *HookDelivery) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *HookDelivery) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAction

`func (o *HookDelivery) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HookDelivery) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HookDelivery) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *HookDelivery) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *HookDelivery) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetInstallationId

`func (o *HookDelivery) GetInstallationId() int32`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *HookDelivery) GetInstallationIdOk() (*int32, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *HookDelivery) SetInstallationId(v int32)`

SetInstallationId sets InstallationId field to given value.


### SetInstallationIdNil

`func (o *HookDelivery) SetInstallationIdNil(b bool)`

 SetInstallationIdNil sets the value for InstallationId to be an explicit nil

### UnsetInstallationId
`func (o *HookDelivery) UnsetInstallationId()`

UnsetInstallationId ensures that no value is present for InstallationId, not even an explicit nil
### GetRepositoryId

`func (o *HookDelivery) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *HookDelivery) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *HookDelivery) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.


### SetRepositoryIdNil

`func (o *HookDelivery) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *HookDelivery) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
### GetUrl

`func (o *HookDelivery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HookDelivery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HookDelivery) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HookDelivery) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *HookDelivery) GetRequest() HookDeliveryRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *HookDelivery) GetRequestOk() (*HookDeliveryRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *HookDelivery) SetRequest(v HookDeliveryRequest)`

SetRequest sets Request field to given value.


### GetResponse

`func (o *HookDelivery) GetResponse() HookDeliveryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *HookDelivery) GetResponseOk() (*HookDeliveryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *HookDelivery) SetResponse(v HookDeliveryResponse)`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


